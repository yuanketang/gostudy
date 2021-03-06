![从零开始自学Go语言 - 进阶篇](https://cdn.yuanketang.cn/images/courses/02/cover.jpg)

### 第一章 Web编程基础

---
#### 1. Web工作原理

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

```bash
# localhost name resolution is handled within DNS itself.
127.0.0.1       localhost
::1             localhost
```

3.如果本地hosts有匹配，发起系统调用。

4.如果匹配项是内网地址则DNS内部处理直接进行内网解析。

5.如果匹配项不是内网地址则发起正常的DNS查询。

6.查询失败那么返回DNS解析失败。

7.查询成功则拿到服务器ip和端口号请求服务器对应ip和端口绑定的应用程序。（web服务一般默认80端口或443端口）

8.进行URL解析。

9.处理请求并返回结果给用户。

##### DNS解析
 <div style="text-align:center;margin: 0 auto;">
    <img src="images/dns.png" alt="DNS解析" />
</div>

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

```shell
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

```shell
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
 * `OPTIONS` 一般用于检测服务器所支持的请求方法。
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
    <img src="images/tcp.png" alt="三次握手，4次挥手" />
</div>

---

#### 2. 路由和路由处理器

##### 什么是Web路由？

> 我们常说的Web路由，简单的说就是一种URL和函数的映射关系。


```mermaid
flowchart TD
subgraph 图2    
D[注册路由] --> E[关联路由处理函数并实现相应的功能逻辑] --> F[返回结果给终端浏览器用户]
end

subgraph 图1
A[URL解析] --> B[响应并处理请求] --> C[返回结果给终端浏览器用户]
end

```

   左图是我们一个浏览器请求Web服务响应流程图，对应到我们右图就是我们开发人员需要做的3件事：
  
  > 绑定一个特定的URL
  关联一个特定的路由处理函数
  完成功能并返回结果


  ##### 路由一般分为以下2种：
 
  <b style="font-size:14px">静态路由</b>
   
   ```
   GET  /users   #获取用户列表
   POST /users   #创建一个新用户
   PUT  /users   #更新用户信息
   ...
   ```

   <b style="font-size:14px">动态路由</b>

   ```
   DELETE /users/:id   #比如id为1时表示删除ID为1的用户
   GET    /users/:id   #比如id为1时表示获取ID为1的用户信息
   ...

   ```

   以上这种语义化的路由方式，我们也称为是一种 `REST API` (Representational State Transfer)，也可以理解为我们常说的 `CRUD` （增删改查）

   > C(Create) 增
   R(Read)     查
   U(Update)   改
   D(Delete)   删
  
  ##### 路由处理器
  
  关联某个特定路由的函数或方法，一般实现具体某些功能和业务逻辑并组合数据呈现给用户，这类函数或方法我们就称为 `路由处理程序`。（对应我们HTTP中的请求处理 `Request` 和响应 `Response`）

  http.Request
  > .Host 获取主机名称和端口号
	.Header 获取请求头
	.Method 获取请求方法类型，如GET, POST
	.URL.RawQuery 获取query-string
	.URL.Fragment 获取anchor
	.Referer() 获取请求来源
	.UserAgent() 获取浏览器特征
  .ParseForm() 将表单或query-string解析成map结构
	.ParseMultipartForm() 一般用于处理文件上传
	.Form.Encode() 对表单进行Url编码，对空格和特殊字符进行编码
  .Form.Get(key) 使用键名获取参数值
	.FormValue(key) 使用键名获取参数值

  http.ResponseWrite
  > .WriteHeader(code) 设置响应头 HTTP 状态码
	.Write([]byte) 设置响应体

  http.StatusXXX定义了完整的HTTP状态码
  > http.StatusOk -> 200
  ...

  方式一:

  ```shell
  package main

  import "net/http"

  // 实现一个自定义http.Handler
  type MyHandle struct{}

  // 实现接口中定义的方法
func (h MyHandle) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
}

  func main() {

    // 注册路由，这里'/'表示匹配web服务根路径，它是一个默认资源路径
    // 使用自定义类型方式
    http.Handle("/", &MyHandle{})
    // 使用匿名函数方式
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    })
    // 启动服务并监听8080端口，这里:8080表示绑定所有可用接口，而不仅仅是具有给定主机地址的接口。
    http.ListenAndServe(":8080", nil)

    // 绑定本地网卡IP内网IP地址为192.167.1.2接口的8080端口（具有给定主机地址）
    // http.ListenAndServe("192.168.1.2:8080", nil)
  }
  ```

  `DefaultServeMux`

  ServeMux默认实现，需要注意的是DefaultServeMux是一个全局变量，容易被污染。
  ```shell
  http.ListenAndServe(":8080", nil)
  ```

  正确的使用方式：
  ```shell
  mu := http.NewServeMux()

  // mu.handelFunc(path, handler)
  //...
  // 监听任意地址的的8080端口，并绑定路由处理器
  http.ListenAndServe(":8080", mu)
  ```

  #### 3. 表单进阶

  * JSON处理
  示例代码：
  ```shell
  // 定义结构体并添加标签
  type User struct {
    Id int        `json:"id"`
    Name string   `json:"name"`
  }
  
  // 读取请求体内容
  bytes, err := ioutil.ReadAll(r.Body)
  if err != nil {
    // 处理异常
  }
  // 防止内存溢出
  defer r.Body.Close()
  // ...
  var t User
  err = json.Unmarshal(bytes, &t) // 注意第二个参数是一个指针
  if err != nil {
    // 处理异常
  }
  // 其他业务逻辑
  
  ```
  * 表单验证
  
  ```shell
  # 安装依赖
  go get -u github.com/go-playground/validator/v10
  ```

  ```shell
  // 定义结构体并添加标签和验证方式
  type User struct {
    Id int        `json:"id" validate:"gt=0"`
    Name string   `json:"name" validate:"required"`
    Password string   `json:"password" validate:"required"`
  }
  
  // 读取请求体内容
  bytes, err := ioutil.ReadAll(r.Body)
  if err != nil {
    // 处理异常
  }
  // 防止内存泄露
  defer r.Body.Close()
  // ...
  var user User
  err = json.Unmarshal(bytes, &user) // 注意第二个参数是一个指针
  if err != nil {
    // 处理异常
  }
  // 表单验证
  // 初始化一个validator
  validator := validator.New()
  err = validator.struct(t)
  if err != nil {
    // 验证失败处理逻辑
    // ...
  }
  // 验证通过，正常逻辑
  ```

  > API文档：https://pkg.go.dev/gopkg.in/go-playground/validator.v10
  
  
  * 文件上传

  ```{.line-numbers}
  // 10M
  const maxMemory = 10 << 20
  // 解析multipart-data, 注意第二个参数为字节数
  err := r.ParseMultipartForm(maxMemory)
  // r.MultipartForm.File 返回是一个map[string]*FileHeader类型，内部实现的
  headers := r.MultipartForm.File
  // 遍历Map
  for _, fh := range headers {
    // 遍历切片 []*FileHeader
    for _, file := range fh {
      // 设置文件保存路径（这里使用绝对路径） './'代表应用程序根路径
      fileSavePath := filepath.Join("./", file.Filename)

      // 打开原文件作为输入
      src, err := file.Open()
      if err != nil {
        panic(err)
      }
      defer src.Close()

      // 创建一个空文件作为输出
      dst, err := os.Create(fileSavePath)
      if err != nil {
        panic(err)
      }
      defer dst.Close()

      // io.Copy使用字节拷贝方式将输入写入输出
      _, err = io.Copy(dst, src)
      if err != nil {
        panic(err)
      }
    }
  }
  ```

  #### 4. Web服务的认证与鉴权（上）
  
  ##### 为什么需要认证和鉴权？

  在实际的工作和生活中，我们往往需要对我们使用的资源进行合理的分配和安全隔离，因此我们就需要某种认证和鉴权的机制。比如，在公司我们需要请假，需要上我们的上级领导申请或同时也需要向人事部门申请，这就是一个需要认证和鉴权的场景。
  
  而对于的Web服务来说，我们经常使用以下几种认证和鉴权的方式：
  
  ##### 账户和密码

  > 概念

  是一种通过为不同的账户分配密码的方式而采取的授权策略。（适合非常简单的应用）

  ##### OAuth（Open Authorization）2.0

  > 概念

  `OAuth` 2.0是用于授权的安全、开放的行业标准协议。是一种无需使用用户名和密码而采用 `用户` `授权` 的方式，允许 `第三方应用` 访问用户在另外的 `服务提供者`上的存储的信息。这里的2.0是它的协议版本。

  比如哔哩哔哩的App有微信登录功能，在这里 `第三方应用` 就是指哔哩哔哩，`用户` 就是指使用者，哔哩哔哩要获取你的微信头像和昵称等信息就要获得你的授权，然后从微信服务器获得数据，微信服务器就是这里的 `服务提供者`。

  那么这些常见的如第三方App微信登录、支付宝登录、新浪微博登录等等，这些都属于采用OAuth协议来进行认证和鉴权的应用。

  > 授权模式

  在OAuth标准中定义了以下几种授权模式：
  
  <b>`Authorization Code` 授权码</b>
  <b>`Client Credentials` 客户端凭证</b>
  `PKCE` 授权码模式的增强版（用于防范跨站攻击）
  `Device Code` 设备代码
  `Refresh Token` 刷新令牌
  `Legacy: Implicit Flow` 隐式流
  `Legacy: Password Grant` 密码

  > 授权码模式

  ```mermaid
  sequenceDiagram
    autonumber
    哔哩哔哩 ->> 用户: 请求使用微信登录
    用户 -->> 微信认证服务器: 请求授权码
    activate 微信认证服务器
    微信认证服务器 -->> 哔哩哔哩: 返回授权码
    deactivate 微信认证服务器
    哔哩哔哩 ->> 微信认证服务器: 使用授权码申请令牌
    activate 微信认证服务器
    微信认证服务器 -->> 哔哩哔哩: 返回令牌
    deactivate 微信认证服务器
    哔哩哔哩 ->> 微信数据服务器: 使用令牌请求数据
    activate 微信数据服务器
    微信数据服务器 -->> 哔哩哔哩: 返回数据
    deactivate 微信数据服务器
  ```

  > 客户端凭证

  ```mermaid
  sequenceDiagram
    autonumber
    哔哩哔哩 ->> 阿里云认证服务器: 使用app_id和app_secret请求认证
    activate 阿里云认证服务器
    阿里云认证服务器 -->> 哔哩哔哩: 返回令牌
    deactivate 阿里云认证服务器
    哔哩哔哩 ->> 阿里云服务: 使用令牌请求数据
    activate 阿里云服务
    阿里云服务 -->> 哔哩哔哩: 返回数据
    deactivate 阿里云服务
  ```

  ##### JWT(Json web token)

  > 概念

  `JWT` 是一种基于JSON的开放标准，是一种不需要在服务端去保留用户的认证信息或者会话信息的对用户信息进行认证方式。

  ```mermaid
  sequenceDiagram
    autonumber
    哔哩哔哩 ->> 用户: 要求用户登录
    用户 ->> 哔哩哔哩认证服务器: 使用账户和密码登录
    activate 哔哩哔哩认证服务器
    哔哩哔哩认证服务器 -->> 哔哩哔哩: 返回令牌
    deactivate 哔哩哔哩认证服务器
    哔哩哔哩 ->> 哔哩哔哩数据服务器: 携带令牌请求数据
    activate 哔哩哔哩数据服务器
    哔哩哔哩数据服务器 -->> 哔哩哔哩: 返回数据
    deactivate 哔哩哔哩数据服务器
  ```

  `JWT` 分为三部分：

  > 1-请求头（Header）
  
  用于请求时携带令牌，一般采用Bearer token方式，令牌需要客户端保存，比如存储在浏览器Cookie或LocalStorage中
  
  ```shell
  # base64编码
  Content-Type: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ5dWFua2V0YW5nIiwiaWF0IjoiMjAyMS0xMS0wNSAwMDowMDowMCIsImV4cCI6IjIwMjEtMTEtMjcgMDA6MDA6MDAiLCJpZCI6IjEifQ.G--jJTGq7g7uIbl-_nlW6HkpGrhWMK0y2w6W6KuIELI
  ```

  > 2-载荷（Payload）（可以理解为请求体）也就是实际存储的有效信息

  公共部分

  `iss`: 签发机构。
  `iat`: 签发时间。
  `exp`: 过期时间，这个过期时间必须要大于签发时间。
  `nbf`: 定义在什么时间之前是不可用的。
  `jti`: 唯一身份标识，主要用防范攻击。

  私有部分

  添加自定义数据，比如一个UUID

  ```json
  {
    "iss": "yuanketang",
    "iat": "2021-11-05 00:00:00",
    "exp": "2021-11-06 00:00:00",
    "uuid": "068bbf4e-2bb5-4a4b-9d6f-9dc8ee4939c3"
  }
  ```

  > 3-签名（Signature）
  
  header (base64后的)
  payload (base64后的)
  secret 随机加密值

  [在线生成JWT](https://tooltt.com/jwt-encode/ "在线生成JWT")

  #### 5. Web服务的认证与鉴权（中）

  需要安装的依赖
  ```bash
  go get -u golang.org/x/oauth2
  go get -u github.com/go-oauth2/oauth2/v4
  ```
  
  ##### 授权码模式的实现

  > 服务端实现

  ```gotemplate
  // 步骤一：初始化一个Manager并配置
  // Manager 认证管理器
  manager := manage.NewDefaultManager()
  // 设置Token默认时间
  manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
  // Token存储位置
  manager.MustTokenStorage(store.NewMemoryTokenStore())
  // Token的生成方式
  manager.MapAccessGenerate(generates.NewAccessGenerate)

  // 步骤二：初始化一个ClientStore并纳入Manager管理
  // 客户端信息配置
  clientStore := store.NewClientStore()
  _ = clientStore.Set("每个单独的应用编号", &models.Client{
    ID:     "客户端ID",
    Secret: "客户端Secret",
    Domain: "客户端地址",
  })
  // 将客户端信息纳入Manager管理
  manager.MapClientStorage(clientStore)
  
  // 步骤三：初始化一个认证服务器并配置
  // OAuth 认证服务器
  // 参数1 初始化配置
  // 参数2 认证管理器
  server := server.NewServer(server.NewConfig(), manager)

  // 是否允许使用GET方式传递token
  srv.SetAllowGetAccessRequest(true)

  // 处理颁发Token（重要）必须实现
  // 参数1 http.ResponseWriter
  // 参数2 *http.Request
  srv.HandleTokenRequest(w http.ResponseWriter, r *http.Request)

  // 处理验证Token（重要）
  // 参数 *http.Request
  srv.ValidationBearerToken(r *http.Request)
  ```

  > 客户端实现

  ```gotemplate
  config = oauth2.Config{
    ClientID:     "客户端ID",
    ClientSecret: "客户端Secret",
    Scopes:       []string{"自定义scope"},
    RedirectURL:  "获取code回调地址",
    Endpoint: oauth2.Endpoint{
      AuthURL:  "授权请求地址", // 授权码模式需要配置
      TokenURL: "申请令牌地址", // 必须配置
    },
  }

  // 请求鉴权的URL
  config.AuthCodeURL("自定义state")

  // 用授权码换取Token
  token, err := config.Exchange(context.Background(), code)
  ```

  ##### 客户端凭证模式的实现

  ```gotemplate
  clientCfg := clientcredentials.Config{
    ClientID:     "客户端ID",
    ClientSecret: "客户端ID",
    TokenURL:     "申请令牌地址",
  }
  
  token, err := clientCfg.Token(context.Background())
  ```

#### 6. Web服务的认证与鉴权（下）
> 安装依赖

```bash
go get github.com/dgrijalva/jwt-go
```

> 创建JWT Token

```gotemplate
// 使用claim创建一个JWT Token
// 参数一：加密方式
// 参数二：claim 可以理解为payload
token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
  IssuedAt: time.Now().Unix(),
  ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
  Issuer: "yuanketang",
})

// 对token进行签名
tokenStr := token.SignedString(秘钥)
// ...
```

> 解析JWT Token

```gotemplate
// 解析Token
// 参数一：JWT Token String
// 参数二：claim
// 参数三：回调函数
token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *Token) (interface{}, error){
  return 秘钥, nil
})

// 捕获特定类型的错误
if err != nil {
  if jwtErr, ok := err.(*jwt.ValidationError); ok {
        if jwtErr.Errors == jwt.ValidationErrorExpired {
            http.Error(writer, "Token过期了", http.StatusBadRequest)
            return
        }
    }
}}

// 判断token是否正确
token.Valid
// ...
```

### 第二章 Go数据库编程

---

#### 1. MySQL

安装依赖

```bash
go get -u github.com/go-sql-driver/mysql
```

#### 2. MongoDB

安装依赖

```bash
go get -u go.mongodb.org/mongo-driver/mongo
```

#### 3. Redis

安装依赖

```bash
go get -u github.com/go-redis/redis
```

### 第三章 Go网络编程

---

#### 1. WebSocket（上）

`WebSocket` 出现的背景

早期，很多网站为了实现推送技术，所用的技术都是轮询（也叫短轮询）。轮询是指由浏览器每隔一段时间向服务器发出 HTTP 请求，然后服务器返回最新的数据给客户端。

<div style="text-align:center;margin: 0 auto;">
    <img src="E:\projects\gostudy\images\poll.png" alt="FRAME" />
</div>

```javascript
// 短轮询
setInterval(() => {
  // Ajax  
}, 5000)
```

```gotemplate
// 长轮询
for {
    // TODO: 检查是否有数据要发送
    time.sleep(time.Second)
}
```

这种传统的模式带来很明显的缺点，即浏览器需要不断的向服务器发出请求，然而 HTTP 请求与响应可能会包含较长的头部，其中真正有效的数据可能只是很小的一部分，所以这样会消耗很多服务器资源。

`WebSocket` 是一种在TCP连接上进行全双工通信的基于 `帧 [Frame]` 的传输协议。(HTTP是基于 `流 [stream]`)

`WebSocket` 使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据。在WebSocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输。

`WebSocket` 的握手过程，符合其 "Web" 的特性，是利用 HTTP 本身的 "协议升级" 来实现。

```
// 客户端
GET ws://localhost/ws HTTP/1.1
Connection: Upgrade
Upgrade: websocket
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: 6mKRFe1lkSpXca0s3paFuijBGTgIgvjK==

// 服务端 
HTTP/1.1 101 Switching Protocols
Connection: Upgrade
Upgrade: websocket
Sec-WebSocket-Accept: EiM9oIF3ryxCrPxNnH2XmQhMLnyp0RnG=
```

Sec-WebSocket-Accept 计算方式

```gotemplate
// GUID "258EAFA5-E914-47DA-95CA-C5AB0DC85B11" (RFC4122)
// Base64(SHA1(Sec-WebSocket-KeyGUID))
// secKey 客户端发来的Sec-WebSocket-Key
h := sha1.New()
h.Write([]byte(secKey))
h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
return base64.StdEncoding.EncodeToString(h.Sum(nil))
```

`WebSocket` 数据帧格式

```
      0                   1                   2                   3
      0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
     +-+-+-+-+-------+-+-------------+-------------------------------+
     |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
     |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
     |N|V|V|V|       |S|             |   (if payload len==126/127)   |
     | |1|2|3|       |K|             |                               |
     +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
     |     Extended payload length continued, if payload len == 127  |
     + - - - - - - - - - - - - - - - +-------------------------------+
     |                               |Masking-key, if MASK set to 1  |
     +-------------------------------+-------------------------------+
     | Masking-key (continued)       |          Payload Data         |
     +-------------------------------- - - - - - - - - - - - - - - - +
     :                     Payload Data continued ...                :
     + - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
     |                     Payload Data continued ...                |
     +---------------------------------------------------------------+
```


FIN：1位，用来表明这是一个消息的是否为最后的消息片断，第一个消息片断也可能是最后的一个消息片断。

RSV1, RSV2, RSV3: 分别都是1位，如果双方之间没有约定自定义协议，那么这几位的值都必须为0，否则必须断掉WebSocket连接。

Opcode：4位操作码，定义有效负载数据，如果收到了一个未知的操作码，连接也必须断掉，以下是定义的操作码：
*  %x0 表示连续消息
*  %x1 表示文本消息
*  %x2 表未二进制消息
*  %x3-7 为将来的非控制消息保留的操作码
*  %x8 表示连接关闭
*  %x9 表示心跳检查的ping
*  %xA 表示心跳检查的pong
*  %xB-F 为将来的控制消息的保留操作码

Mask: 1位，定义传输的数据是否有加掩码，如果设置为1，掩码键必须放在masking-key区域。

Payload length: 传输数据的长度，以字节的形式表示，共有3种形式： 7位、7+16位、或者7+64位。

> 如果这个值以字节表示是0-125这个范围，那这个值就表示传输数据的长度；

> 如果这个值是126，则随后的两个字节表示的是一个16进制无符号数，用来表示传输数据的长度；

> 如果这个值是127,则随后的是8个字节表示的一个64位无符合数，这个数用来表示传输数据的长度。

Masking-key: 0或4个字节，32位，掩码键只有在掩码位设置为1的时候存在。

Payload data: (x+y)位，载荷数据为扩展数据及应用数据长度之和。

Extension data: x位，如果客户端与服务端之间没有特殊约定，那么扩展数据的长度始终为0，任何的扩展都必须指定扩展数据的长度，或者长度的计算方式，以及在握手时如何确定正确的握手方式。如果存在扩展数据，则扩展数据就会包括在负载数据的长度之内。

Application data: y位，任意的应用数据，放在扩展数据之后，应用数据的长度=载荷数据的长度-扩展数据的长度。

掩码计算方式
j = i MOD 4
transfromed-octed-i = original-octet-i XOR masking-key-octet-j
