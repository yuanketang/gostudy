package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"log"
	"net"
	"net/http"
)

type WebSocket struct {
	Conn       net.Conn
	IsMasked   bool
	MaskingKey []byte
}

func (ws *WebSocket) getKey(key string) string {
	h := sha1.New()
	h.Write([]byte(key))
	h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (ws *WebSocket) HandShake(writer http.ResponseWriter, request *http.Request) error {
	secKey := request.Header.Get("Sec-WebSocket-Key")
	if secKey == "" {
		return errors.New("请求头错误 Sec-WebSocket-Key")
	}

	// 向客户端写入我们的相应头
	// HTTP/1.1 101 Switching Protocols
	// Connection: Upgrade
	// Upgrade: websocket
	// Sec-WebSocket-Accept: EiM9oIF3ryxCrPxNnH2XmQhMLnyp0RnG=

	p := []byte{}
	p = append(p, "HTTP/1.1 101 Switching Protocols\r\n"...)
	p = append(p, "Connection: Upgrade\r\n"...)
	p = append(p, "Upgrade: websocket\r\n"...)
	p = append(p, "Sec-WebSocket-Accept: "...)
	p = append(p, ws.getKey(secKey)...)
	p = append(p, "\r\n"...)
	p = append(p, "\r\n"...)

	h, ok := writer.(http.Hijacker)
	if !ok {
		return errors.New("数据劫持失败")
	}
	conn, _, err := h.Hijack()
	if err != nil {
		return errors.New("数据劫持失败")
	}
	ws.Conn = conn
	if _, err = ws.Conn.Write(p); err != nil {
		ws.Conn.Close()
		return err
	}
	return nil
}

func (ws *WebSocket) ReceiveMessage() (data []byte, err error) {
	n := make([]byte, 2)
	_, _ = ws.Conn.Read(n)

	// 1001 0001 1000 0000
	fin := n[0]&1<<7 != 0
	if rsv := n[0] & (1<<6 | 1<<5 | 1<<4); rsv != 0 {
		return nil, errors.New("不支持自定义扩展协议")
	}
	// 1001 0001 0000 1111
	opcode := int(n[0] & 0xf)

	ws.IsMasked = n[1]&1<<7 != 0
	// 0000 0011 0111 1111
	// <= 125 情况
	payloadLen := int(n[1] & 0x7f)

	// TODO: 作业
	payload := make([]byte, payloadLen)
	if ws.IsMasked {
		_, _ = ws.Conn.Read(ws.MaskingKey)
	}
	_, _ = ws.Conn.Read(payload)
	originBytes := make([]byte, payloadLen)
	// 32 8
	if ws.IsMasked {
		//  j = i MOD 4 transfromed-octed-i = original-octet-i XOR masking-key-octet-j
		// 转换后数据 = 原始数据[i] ^ 掩码数据[i mod 4]
		for i := 0; i < payloadLen; i++ {
			originBytes[i] = payload[i] ^ ws.MaskingKey[i%4]
		}
	} else {
		originBytes = payload
	}
	log.Printf("n: %b\n", n)
	log.Printf("fin: %t\n", fin)
	log.Println("opcode", opcode)
	log.Printf("payloadLen: %d", payloadLen)
	log.Printf("masking key: %b", ws.MaskingKey)

	if fin {
		data = originBytes
		return data, nil
	}

	// TODO: 作业
	return originBytes, nil
}

func main() {
	mu := http.NewServeMux()

	mu.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./advance/chapter3/lesson1/public/index.html")
	})

	mu.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		// 1 实现握手
		ws := &WebSocket{
			MaskingKey: make([]byte, 4),
		}
		err := ws.HandShake(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer ws.Conn.Close()
		// 2 实现对数据报文解析
		for {
			data, err := ws.ReceiveMessage()
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			if data != nil {
				log.Println("Data:", string(data))
			}
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
