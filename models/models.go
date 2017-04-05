package models

import (
    "blog/helpers"
    "crypto/md5"
    "fmt"
    "io"
    "log"
    "math/rand"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
    db   *gorm.DB
    conf = helpers.GetConfig()
    dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
    var err error
    db, err = gorm.Open(selectDatabase())
    if err != nil {
        log.Fatalln("Database connection failed")
    }
    if conf.Debug {
        db.LogMode(true)
    }
    db.AutoMigrate(new(User), new(Post))
}

func selectDatabase() (string, string) {
    var connection string
    dbtype := conf.Database.Type
    switch dbtype {
    case "mysql":
        connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
            conf.Database.User,
            conf.Database.Pass,
            conf.Database.Host,
            conf.Database.Port,
            conf.Database.Name,
        )
    case "postgres":
        connection = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
            conf.Database.User,
            conf.Database.Pass,
            conf.Database.Host,
            conf.Database.Port,
            conf.Database.Name,
        )
    default:
        connection = conf.Database.Name
    }
    return dbtype, connection
}

func EncryptPwd(pwd string) string {
    hash := md5.New()
    salt := conf.SecretKey
    io.WriteString(hash, pwd)
    io.WriteString(hash, salt)
    return fmt.Sprintf("%x", hash.Sum(nil))
}

func GenShortId(size int) string {
    rand.Seed(time.Now().UnixNano())
    bytes := make([]byte, size)
    for i := range bytes {
        bytes[i] = dict[rand.Int63()%int64(len(dict))]
    }
    return string(bytes)
}
