package user

import (
    "github.com/gin-gonic/gin"
    "my-gin/internal/application/user"
    "net/http"
)

func RegisterRoutes(r *gin.Engine, svc *user.UserService) {
    r.POST("/user/register", func(c *gin.Context) {
        var req struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := svc.Register(req.Username, req.Password); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "registered"})
    })

    r.POST("/user/login", func(c *gin.Context) {
        var req struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        ok, _ := svc.Login(req.Username, req.Password)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "login success"})
    })

    r.GET("/user/info", func(c *gin.Context) {
        // 这里可以从 service 或 session 获取用户信息
        // 为测试方便，先返回固定数据
        c.JSON(http.StatusOK, gin.H{
            "username": "liziyue",
            "role":     "admin",
        })
    })
}
