# cmd
#### shadowsocks-local
    go-shadowsocks2 -c 'ss://AEAD_CHACHA20_POLY1305:123qwe@127.0.0.1:1555' \
    -verbose -socks :1080
#### client
    ./c -localaddr 127.0.0.1:1555 -remoteaddr 127.0.0.1:2555
#### server
    ./s -localaddr :2555 -remoteaddr 127.0.0.1:3555
#### shadowsocks-server
    go-shadowsocks2 -s 'ss://AEAD_CHACHA20_POLY1305:123qwe@:3555' -verbose