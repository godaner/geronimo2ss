# geronimo
### Introduction
    Geronimo2ss.
### Features
    nothing
### Usage
##### shadowsocks-local,don't need change
    ./go-shadowsocks2 -c 'ss://AEAD_CHACHA20_POLY1305:123qwe@127.0.0.1:1555' -socks :1080 -verbose
##### client,change remoteaddr ip to server
    ./c -localaddr 127.0.0.1:1555 -remoteaddr 127.0.0.1:2555
##### server,don't need change
    ./s -localaddr :2555 -remoteaddr 127.0.0.1:3555
##### shadowsocks-server,don't need change
    ./go-shadowsocks2 -s 'ss://AEAD_CHACHA20_POLY1305:123qwe@:3555' -verbose