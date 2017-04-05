package helpers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

type Config struct {
    Debug     bool
    Addr      string `json:"http_addr"`
    Public    string `json:"public_dir"`
    SecretKey string `json:"secret_key"`
    Admin     AdminConfig
    Database  DatabaseConfig
}

type AdminConfig struct {
    Route    string `json:"route"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type DatabaseConfig struct {
    Type string `json:"db_type"`
    Host string `json:"db_host"`
    Port string `json:"db_port"`
    Name string `json:"db_name"`
    User string `json:"db_user"`
    Pass string `json:"db_pass"`
}

var (
    config *Config
)

func init() {
    config = new(Config)
    bytes := ReadFile("/conf/conf.json")
    err := json.Unmarshal(bytes, config)
    if err != nil {
        log.Fatalln("JSON parse failed")
    }
}

func ReadFile(path string) []byte {
    pwd, _ := os.Getwd()
    bytes, _ := ioutil.ReadFile(pwd + path)
    return bytes
}

func GetConfig() *Config {
    return config
}
