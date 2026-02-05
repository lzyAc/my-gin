package main

import (
    "github.com/gin-gonic/gin"
    db "MxiqiGo/internal/infrastructure/db"
    infra "MxiqiGo/internal/infrastructure"
    httpTransport "MxiqiGo/internal/transport/http"
)

func main() {
    // 初始化全局 DB
    db.InitDB()

    // 初始化所有模块 service/repo
    modules := infra.InitModules()

    // 注册所有路由
    r := gin.Default()
    httpTransport.RegisterAllRoutes(r, modules)

    // 启动 HTTP 服务
    r.Run(":8080")
}
