package user

import (
    "my-gin/internal/domain/user/entity"
    userRepo "my-gin/internal/domain/user"
)

type UserService struct {
    Repo userRepo.UserRepository
}

func (s *UserService) Register(username, password string) error {
    user := &entity.User{
        Username: username,
        Password: password,
    }
    return s.Repo.Create(user)
}

func (s *UserService) Login(username, password string) (bool, error) {
    u, err := s.Repo.GetByUsername(username)
    if err != nil {
        return false, err
    }
    return u.Password == password, nil
}
