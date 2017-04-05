#!/bin/bash
# tar -zcf blog.tar.gz ./*
WIDTH=$(tput cols)
NUMBER=$((WIDTH - 12))
DEPLOY="" # Deploy dir like `/home/deploy/workspace/blog`
SYSTEM=$2
SUFFIX=""

build() {
    case "$SYSTEM" in
        win)
            SUFFIX=".exe"
            GOOS=windows GOARCH=amd64
            printf "[*] %*s" "-$NUMBER" "building for windows ..."
            ;;
        mac)
            GOOS=darwin GOARCH=amd64
            printf "[*] %*s" "-$NUMBER" "building for macOS ..."
            ;;
        *)
            GOOS=linux GOARCH=amd64
            printf "[*] %*s" "-$NUMBER" "building for linux ..."
            ;;
    esac
    go build -o "blog$SUFFIX" -ldflags "-s -w" main.go
    printf "[\033[32;1m OK \033[0m]\n"
}

push() {
    printf "[*] %*s" "-$NUMBER" "push file ..."
    cp "blog$SUFFIX" $DEPLOY
    cp -fr conf views public $DEPLOY
    printf "[\033[32;1m OK \033[0m]\n"
}

start() {
    build
    push
    export GIN_MODE=release
    nohup "$DEPLOY/blog" 1>"$DEPLOY/access.log" 2>"$DEPLOY/error.log" &
    printf "[+] %*s[\033[34;1m DONE \033[0m]\n" "-$NUMBER" "success"
}

stop() {
    PID=$(ps -eaf | grep blog | grep -v grep | head -1 | awk {'print $2'})
    if [[ "" != "$PID" ]]; then
        printf "[*] %*s" "-$NUMBER" "kill process $PID ..."
        kill -9 $PID
        printf "[\033[32;1m OK \033[0m]\n"
    else
        printf "[-] %*s[\033[31;1m ERR \033[0m]\n" "-$NUMBER" "pid not found"
    fi
}

restart() {
    stop
    start
}

usage() {
    echo "Usage: $0 {[-b|--build]<win,linux,mac>|start|stop|restart|status}"
}

case "$1" in
    -b | --build) build || exit 1 ;;
    start) start || exit 1 ;;
    stop) stop || exit 1 ;;
    restart) restart || exit 1 ;;
    status) ps -eaf | grep blog | grep -v grep | head -1 ;;
    *) usage ;;
esac