package user

import (
    "MxiqiGo/internal/domain/user/entity"
    "MxiqiGo/internal/infrastructure/db"
    userModeldb "MxiqiGo/internal/infrastructure/user/db"
)

type UserRepo struct{}

func (r *UserRepo) Create(user *entity.User) error {
    u := userModeldb.UserModel{
        Username: user.Username,
        Password: user.Password,
    }
    return db.DB.Create(&u).Error
}

func (r *UserRepo) GetByUsername(username string) (*entity.User, error) {
    var u userModeldb.UserModel
    if err := db.DB.Where("username = ?", username).First(&u).Error; err != nil {
        return nil, err
    }
    return &entity.User{
        ID:       u.ID,
        Username: u.Username,
        Password: u.Password,
    }, nil
}
