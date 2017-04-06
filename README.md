# Blog - 基于 Go-Macaron 编写的极简博客

### 运行依赖如下

```Bash
go get gopkg.in/macaron.v1
go get github.com/go-macaron/gzip
go get github.com/go-macaron/session
go get github.com/jinzhu/gorm
```

建议在 GO >= 1.5 && Linux 环境运行

### 配置文件

```
{
    "debug": false,                  # 开启Debug记录
    "http_addr": "0.0.0.0:80",       # 内网监听端口
    "public_dir": "public",          # 静态文件目录
    "secret_key": "!!!F4CK*Y0U!!!",  # 加密KEY
    "admin": {
        "route": "admin",            # 后台路由
        "username": "admin",         # 初始用户名
        "password": "123456"         # 初始密码
    },                               # 初始值仅在第一次运行建表时生效
    "database": {                    # 支持的数据库(MySQL,SQLite,Postgres)
        "db_type": "postgres",
        "db_host": "localhost",
        "db_port": "5432",
        "db_name": "blog",
        "db_user": "root",
        "db_pass": "root"
    }                                # SQLite 直接把相对路径填在db_name项即可
}
```

## Screenshot

后台模版扒了typecho的模版

![Screenshot][1]
![Screenshot][2]
![Screenshot][3]

[1]: screenshots/browse.png
[2]: screenshots/post.png
[3]: screenshots/admin.png



