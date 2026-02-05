package entity

// demo解释： User不是数据库表 也不是http返回结构。 他是业务中的用户！

type User struct {
    ID       uint
    Username string
    Password string
}