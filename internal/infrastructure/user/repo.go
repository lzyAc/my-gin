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
    return db.UserDB.Create(&u).Error
}

func (r *UserRepo) GetByUsername(username string) (*entity.User, error) {
    var u userModeldb.UserModel
    if err := db.UserDB.
        Where("username = ?", username).
        First(&u).Error; err != nil {
        return nil, err
    }

    return &entity.User{
        ID:       u.ID,
        Username: u.Username,
        Password: u.Password,
    }, nil
}

func (r *UserRepo) TestInfo() (*entity.User, error) {
    return &entity.User{
        ID:       1,
        Username: "liziyue",
    }, nil
}
