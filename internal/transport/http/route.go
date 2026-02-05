package http

import (
    "github.com/gin-gonic/gin"
    infra "MxiqiGo/internal/infrastructure"
    userTransport "MxiqiGo/internal/transport/http/user"
    // orgTransport "MxiqiGo/internal/transport/http/org" // 未来扩展
)


// RegisterAllRoutes 统一注册所有模块路由
func RegisterAllRoutes(r *gin.Engine, modules *infra.Modules) {
    // user 模块
    userTransport.RegisterRoutes(r, modules.UserService)

    // org 模块（未来扩展）
    // orgTransport.RegisterRoutes(r, modules.OrgService)
}
