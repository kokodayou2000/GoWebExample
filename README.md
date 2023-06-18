# GoWebExample
写了一些go web的简单例子

官网 https://gowebexamples.com/

包含
- Hello World
  - net/http 简单使用
- HTTP Server
  - 简单的http请求，可以指定文件路径
- Routing
  - 使用 `github.com/gorilla/mux` 实现路由器
- MySQL Database
  - 连接数据库，简单增删改查
- Templates
  - html模板，使用 `template.Must`解析html
- Assets and Files
  - 文件服务
- Forms
  - 表单提交，使用的还是`template.Must`解析html，该html能发送post请求
- Middleware(Basic) 
  - 让多个handle有log日志的功能
- Middleware(Advanced)
  - 高级点的功能，使用链式调用，关于链式调用，写了一个demo，可以看看
- Sessions
  - 使用的 `github.com/gorilla/sessions`
    - 登录 获取cookie
    - 获取隐私数据 必须携带 cookie
    - 结束登录 清除cookie
- JSON
  - json的编码与解码
- Websockets 
  - 网络套接字 使用的`github.com/gorilla/websocket`
  - 包含设置写缓存，读缓存等操作
- Password Hashing
  - 密码加密和解密
  - `golang.org/x/crypto/bcrypt`

