# Web 框架

`Web` 框架一般会提供如下核心能力：

- 路由（Routing）：将请求映射到处理函数，需要支持动态路由，如 `/hello/:name`；
- 插件（Plugins）：以插件的形式在请求前后进行处理，如认证鉴权、访问日志等；
- 工具集（Utilities）：提供对 `Cookies`、`Headers` 等处理的工具；
- 模板（Templates）：使用内置模板引擎提供模板渲染机制；
- ...