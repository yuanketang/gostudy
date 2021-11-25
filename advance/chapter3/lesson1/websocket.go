package main

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

type WebSocket struct {
	Conn     net.Conn
	MaskKey  []byte
	IsMasked bool
}

func (ws *WebSocket) Close() error {
	err := ws.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ws *WebSocket) HandleShake(writer http.ResponseWriter, request *http.Request) error {
	// 检查请求头
	secKey := request.Header.Get("Sec-Websocket-Key")
	if secKey == "" {
		return errors.New("请求头错误")
	}
	log.Println("Sec-Websocket-Key", secKey)

	// TODO: 跨域检查

	// 劫持http请求
	h, ok := writer.(http.Hijacker)
	if !ok {
		return errors.New("请求错误")
	}

	conn, _, err := h.Hijack()
	if err != nil {
		return err
	}

	var buf []byte
	p := buf[:0]
	p = append(p, "HTTP/1.1 101 Switching Protocols\r\n"...)
	p = append(p, "Upgrade: websocket\r\n"...)
	p = append(p, "Connection: Upgrade\r\n"...)
	p = append(p, "Sec-WebSocket-Accept: "...)
	p = append(p, ws.generateKey(secKey)...)
	p = append(p, "\r\n"...)

	// 请求体和请求头中间的换行
	p = append(p, "\r\n"...)
	if _, err = conn.Write(p); err != nil {
		conn.Close()
		return err
	}
	ws.Conn = conn
	return nil
}

func (ws *WebSocket) generateKey(secKey string) string {
	// 构造规则：BASE64(SHA1(Sec-WebSocket-KeyGUID))
	// GUID(RFC4122)
	h := sha1.New()
	h.Write([]byte(secKey))
	h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (ws *WebSocket) WriteMessage(data []byte) error {
	payloadLen := len(data)
	// * 实例仅考虑paylaod长度小于126情况
	if payloadLen > 125 {
		return errors.New("数据长度超出范围")
	}
	// 写第1个字节 1000 0001
	ws.Conn.Write([]byte{0x81})
	// 写第2个字节 0000 0000 | 0000 0111
	ws.Conn.Write([]byte{byte(0x00) | byte(payloadLen)})
	// 处理掩码
	if ws.IsMasked {
		maskedData := make([]byte, payloadLen)
		for i := 0; i < payloadLen; i++ {
			maskedData[i] = data[i] ^ ws.MaskKey[i%4]
		}
		ws.Conn.Write(maskedData)
	}
	ws.Conn.Write(data)
	return nil
}

func (ws *WebSocket) ReceiveMessage() (data []byte, err error) {
	// 读取前二个字节
	n := make([]byte, 2)
	_, _ = ws.Conn.Read(n)
	fin := n[0]&1<<7 != 0
	opcode := int(n[0] & 0xf)
	if rsv := n[0] & (1<<6 | 1<<5 | 1<<4); rsv != 0 {
		return nil, errors.New("不支持自定义扩展协议")
	}

	ws.IsMasked = n[1]&1<<7 != 0
	payloadLen := int(n[1] & 0x7f)
	// TODO 需要判断下数据长度，是否有扩展数据

	// 获取掩码
	if ws.IsMasked {
		_, _ = ws.Conn.Read(ws.MaskKey)
	}

	payload := make([]byte, payloadLen)
	_, _ = ws.Conn.Read(payload)
	dataByte := make([]byte, payloadLen)
	if ws.IsMasked {
		// 掩码算法 转换后数据[i] = 原始数据[i] ^ 掩码数据[i%4]
		for i := 0; i < payloadLen; i++ {
			dataByte[i] = payload[i] ^ ws.MaskKey[i%4]
		}
	} else {
		dataByte = payload
	}
	log.Printf("%b\n", n)
	log.Printf("fin %t\n", fin)
	log.Printf("opcode %d\n", opcode)
	log.Printf("payloadLen %d\n", payloadLen)
	log.Printf("mask %t\n", ws.IsMasked)
	log.Printf("maskData %b\n", ws.MaskKey)

	// 如果是最后一帧
	if fin {
		data = dataByte
		return data, nil
	}

	nextData, err := ws.ReceiveMessage()
	if err != nil {
		return nil, err
	}
	data = append(data, nextData...)
	return data, nil
}

func main() {
	mu := http.NewServeMux()

	mu.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./advance/chapter3/lesson1/public/index.html")
	})

	mu.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws := &WebSocket{
			MaskKey: make([]byte, 4),
		}
		// 处理握手
		err := ws.HandleShake(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		//go func() {
		//	for {
		//		err = ws.WriteMessage([]byte("heartbeat...."))
		//		time.Sleep(time.Second)
		//		if err != nil {
		//			break
		//		}
		//	}
		//}()
		// 消息处理
		for {
			data, err := ws.ReceiveMessage()
			if err != nil {
				ws.Close()
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				break
			}
			log.Println("Data:", data)
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
