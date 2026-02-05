package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "MxiqiGo/internal/pkg/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            c.Abort()
            return
        }

        claims, err := jwt.ParseToken(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}

func main() {
    r := gin.Default()

    // 登录请求直接转发到 HTTP 服务
    r.POST("/user/login", func(c *gin.Context) {
        c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/user/login")
    })

    // 需要 JWT 的接口
    auth := r.Group("/", JWTMiddleware())
    {
        auth.Any("/user/*", proxyToHTTP)
    }

    r.Run(":8000") // 网关端口
}

func proxyToHTTP(c *gin.Context) {
    // 这里只是示例，你可以使用 reverse proxy 或 http client 转发
    c.JSON(200, gin.H{"msg": "Request would be forwarded to HTTP service"})
}
