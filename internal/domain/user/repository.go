package user

import "my-gin/internal/domain/user/entity"

type UserRepository interface {
    Create(user *entity.User) error
    GetByUsername(username string) (*entity.User, error)
}
