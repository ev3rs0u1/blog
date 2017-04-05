package models

type User struct {
    Id       uint   `gorm:"primary_key"`
    Username string `gorm:"not null;unique;size:15"`
    Password string `gorm:"not null;size:32"`
    Author   string `gorm:"size:20"`
    Url      string `gorm:"size:255"`
}

func init() {
    if _, err := GetUser(); err != nil {
        user := new(User)
        user.Username = conf.Admin.Username
        user.Password = EncryptPwd(conf.Admin.Password)
        db.Create(user)
    }
}

func (User) TableName() string {
    return "admin"
}

func (user *User) Update() {
    db.Model(user).Updates(user)
}

func GetUser() (*User, error) {
    user := new(User)
    err := db.First(user).Error
    return user, err
}

func ValidateAccount(uname string, pwd string) *User {
    user, _ := GetUser()
    if user.Username == uname &&
        user.Password == EncryptPwd(pwd) {
        return user
    }
    return nil
}
