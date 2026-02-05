package infrastructure

import (
    userApp "MxiqiGo/internal/application/user"
    userRepo "MxiqiGo/internal/infrastructure/user"
)

// Modules 保存模块的 service 实例
type Modules struct {
    UserService *userApp.UserService
    // OrgService *orgApp.OrgService // 未来扩展
}

// InitModules 初始化各模块，返回 Modules
func InitModules() *Modules {
    // 1初始化 repo
    repo := &userRepo.UserRepo{}

    // 2 构建 service
    svc := &userApp.UserService{Repo: repo}

    // 3 返回模块集合
    return &Modules{
        UserService: svc,
    }
}
