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
    // 初始化 repo和构建 service
    repo := &userRepo.UserRepo{}
    userSvc := &userApp.UserService{Repo: repo}

    // 3 返回模块集合
    return &Modules{
        UserService: userSvc,
    }
}
