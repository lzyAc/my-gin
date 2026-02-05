package user

import "MxiqiGo/internal/domain/user/entity"

type UserRepository interface {
    Create(user *entity.User) error
    GetByUsername(username string) (*entity.User, error)
    TestInfo()(*entity.User, error)
}
