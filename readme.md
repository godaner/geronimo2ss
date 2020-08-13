# geronimo
### Introduction
    Geronimo2ss.
### Features
    nothing
### Usage
##### shadowsocks-local,don't need change
    ./go-shadowsocks2 -c 'ss://AEAD_CHACHA20_POLY1305:123qwe@127.0.0.1:1555' -socks :1080 -verbose
##### client,change remoteaddr ip to server
    ./c-linux64 -localaddr 127.0.0.1:1555 -remoteaddr 127.0.0.1:2555 -verbose
##### server,don't need change
    ./s-linux64 -localaddr :2555 -remoteaddr 127.0.0.1:3555 -verbose
##### shadowsocks-server,don't need change
    ./go-shadowsocks2 -s 'ss://AEAD_CHACHA20_POLY1305:123qwe@:3555' -verbose