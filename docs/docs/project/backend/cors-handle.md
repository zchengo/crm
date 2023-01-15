# 跨域请求处理

您可以在前端、服务端或者 Nginx 服务器上处理跨域请求。

## 什么是跨域请求？

当您在浏览器中从一个域名的网页去请求另一个域名的资源时，域名、端口、协议中的任一不同，就会出现请求跨域。只要协议、域名和端口中的任何一个不同，都会被当作是不同的域，之间的请求就是跨域请求。

## 在服务端处理跨域请求

请参考如下代码：

```go
// Cors 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id, Uid, Ver")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
```

将 Cors 作为 Gin 的中间件，用户每次发送 API 请求时，都会经过此中间件（即每个跨域请求都会被处理）。

处理跨域请求的方式是在这个请求中设置跨域相关的请求头：

| 跨域相关的请求头 | 说明 |
| --- | -- |
| Access-Control-Allow-Origin | 允许跨域的域名，指定了该响应的资源是否被允许与给定的来源（origin）共享。 |
| Access-Control-Allow-Headers | 允许跨域的请求头, 用于预检请求中，列出了将会在正式请求的 ```Access-Control-Request-Headers``` 字段中出现的首部信息。 |
| Access-Control-Allow-Methods | 允许跨域请求的方法，在对预检请求的应答中明确了客户端所要访问的资源允许使用的方法或方法列表。 |
| Access-Control-Expose-Headers | 允许服务器指示那些响应标头可以暴露给浏览器中运行的脚本，以响应跨域请求。 |
| Access-Control-Allow-Credentials | 允许客户端携带验证信息，用于在请求要求包含 ```credentials``` 时，告知浏览器是否可以将对请求的响应暴露给前端 JavaScript 代码。Credentials 可以是 cookies 或 authorization headers。 |

## 在 Vue + Vite 项目中配置允许跨域请求

在 ```vite.config.js``` 配置文件中添加请求代理配置，请参考如下配置：

```js
export default defineConfig({
  server: {
    host: '127.0.0.1',
    port: 8060,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000/',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
      }
    }
  },
  plugins: [vue()],
})
```

## 在 Nginx 中配置允许跨域请求

修改 ```nginx.conf``` 配置文件，请参考如下：

```bash
server {
    
    listen       80;
    server_name  zocrm.cloud;

    location / {
        root   html/dist;
        index  index.html;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   html;
    }

	# api proxy
	location /api {
        # 配置跨域
        add_header 'Access-Control-Allow-Origin' *;
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Allow-Methods' *;
        add_header 'Access-Control-Allow-Headers' *;
        proxy_pass http://127.0.0.1:8000;
    }
}
```

## 如何选择？

以上三种跨域请求处理方式，比较推荐在服务端或者 Nginx 服务器上处理跨域请求，好处是跨域请求相关的配置可以由我们自己来决定。






