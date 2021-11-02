![从零开始自学Go语言 - 进阶篇](https://cdn.yuanketang.cn/images/courses/01/cover02.jpg)

### 第一章 Web编程基础
---
1.Web工作原理

我们常说的Web服务一般指B/S架构，B（Browser）指的就是我们的浏览器充当客户端，S（Server）指我们的Web服务器，作为服务器提供相应的服务，比如网页托管，文件托管等等。

我们看下在浏览器中输入一个网址后发生了什么？
```mermaid
 flowchart TD

    subgraph 图1

    A[输入网址] --> Y{查询浏览器缓存};
    Y -- 否 --> B[查询本地hosts文件];
    Y -- 是 --> X[返回本地缓存结果];
    B --> C{是否匹配};
    C -- 是 --> D{是否为内网地址};
    D -- 是 --> E[DNS内部直接处理];
    D -- 否 --> F[外网DNS查询];
    C -- 否 --> F
    F -- 是 --> H[请求Web服务器80端口或443端口];
    F -- 否 --> Z[返回DNS解析失败];
    H -- 图2 --> J[URL解析]
    subgraph 2    
    J --> I[响应并处理请求]
    I --> K[终端浏览器用户]
    end
    end
    
```

观察上图其实在浏览器输入一个网址后，简单的说会触发以下关键流程：

1.查询本地浏览器缓存，如果查询成功则返回本地缓存结果给用户。

2.查询本地缓存失败，会访问本地hosts文件，并检查是否有指定的匹配项。

```shell
# localhost name resolution is handled within DNS itself.
#       127.0.0.1       localhost
#       ::1             localhost
```

3.如果本地hosts有匹配，发起系统调用。

4.如果匹配项是内网地址则DNS内部处理直接进行内网解析。

5.如果匹配项不是内网地址则发起正常的DNS查询。

6.查询失败那么返回DNS解析失败。

7.查询成功则拿到服务器ip和端口号请求服务器对应ip和端口绑定的应用程序。（web服务一般默认80端口或443端口）

8.进行URL解析。

9.返回结果给用户。

##### URL(我们常说的网址)（URI 统一资源标识符 URL是URI的子集）

URL(Uniform Resource Locator)是 `统一资源定位符` 的英文缩写，用于描述一个网络上的资源，基本格式如下:
```
scheme://host[:port#]/path/.../[?query-string][#anchor]
scheme         指定使用的协议(例如：http, https, ftp)
host           服务器的IP地址或者域名
port#          HTTP服务器的默认端口是80，HTTPS服务器默认端口是443，这种情况下端口号可以省略。如果使用了别的端口，必须指明，例如 http://www.bilibili.com:8080/
path           访问资源的路径 '/'一般指服务器的根路径
query-string   发送给http服务器的数据
anchor         锚点（hash哈希）
```

`HTTP`（Hyper Text Transfer Protocol）也叫 `超文本传输协议`。HTTP是一种让Web服务器与浏览器(客户端)通过Internet发送与接收数据的协议，它建立在TCP协议之上，一般采用TCP的80端口。

`HTTPS`（Hyper Text Transfer Protocol over SecureSocket Layer）也叫 `超文本传输安全协议`，一般采用TCP的443端口。

##### HTTP协议由以下3部分组成：

1.请求行

2.请求头

3.请求体

##### HTTP 请求示例

```http
#请求行
GET https://www.bilibili.com/ HTTP/1.1
#请求头
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36
# 自定义请求头
x-request-id: 123456
// 空行
// 请求体
```

##### HTTP 响应示例

```http
#响应行
HTTP/1.1 200 OK
#响应头
cache-control: no-cache
content-encoding: gzip
content-type: text/html; charset=utf-8
date: Mon, 01 Nov 2021 12:55:08 GMT
expires: Mon, 01 Nov 2021 12:55:07 GMT
#自定义响应头
x-cache-time: 0
x-cache-webcdn: MISS from blzone02
// 空行
// 返回内容 html...或其他
```

##### HTTP请求方法

 * `GET` 读取数据，一般语义表示读取数据。
 * `POST` 提交数据（例如提交表单或者上传文件），数据被包含在请求体中。一般语义表示增加数据
 * `PUT` 提交数据，一般语义表示修改数据。
 * `DELETE` 一般语义表示删除数据。
 * `OPTION` 一般用于检测服务器所支持的请求方法。
 * PATCH 提交数据，一般语义表示修改部分数据。
 * HEAD 读取数据但不返回请求体。
 * TRACE 主要用于测试或诊断。
 * CONNECT 一般用于创建代理。

##### 比较重要的状态码

* 1XX 提示信息 - 表示收到请求，需要继续执行操作。
* 2XX 成功 - 表示请求被成功接收并处理。
* 3XX 重定向 - 要完成请求必须进行更进一步的处理。
* 4XX 客户端错误 - 请求有语法错误或请求无法实现。
* 5XX 服务器端错误 - 在处理请求的过程中发生了错误。

##### 其他相关：

1. HTTP协议常见版本：
  * 1.0（早期，很少使用）
  * 1.1（普遍）
  * 2.0（普及中）
  * 3（未来）

2. HTTP协议是一种无状态、无连接的协议
 * 无状态（服务器中没有保存客户端的状态）
 * 无连接（每次处理一个请求。服务器处理完请求，并收到客户端的应答后，即断开连接）

 `但HTTP 1.1协议增加 `keep-alive` 请求头，可以有效实现连接复用。（免去每次创建连接、TCP3次握手的性能损失）`

##### 三次握手，4次挥手
 <div style="text-align:center;margin: 0 auto;">
    <img src="../../images/tcp.png" alt="三次握手，4次挥手" />
</div>