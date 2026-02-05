package user

import (
    "github.com/gin-gonic/gin"
    "MxiqiGo/internal/application/user"
    "net/http"
    "MxiqiGo/pkg/logger"
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
        userInfo, err := svc.TestInfo() // 未来可以替换为真实 Service 获取
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        logger.Write("user","测试全流程到db",userInfo)
        c.JSON(http.StatusOK, gin.H{
            "id":       userInfo.ID,
            "username": userInfo.Username,
            "role":     "admin", // 可以扩展为实际角色
        })
    })
}
