# geronimo
### Introduction
    Geronimo2ss.
### Features
    nothing
### Usage
##### shadowsocks-local,don't need change
    nohup ./go-shadowsocks2 -c 'ss://AEAD_CHACHA20_POLY1305:123qwe@127.0.0.1:1555' -socks :1080 -verbose  >  ssc.log  2>&1  & 
##### client,change remoteaddr ip to server
    nohup ./c-linux64 -localaddr 127.0.0.1:1555 -remoteaddr xx.xx.xx.xx:2555 -overbose -enc aes-256-cfb@123qwe -loglev=1  >  c.log  2>&1  & 
##### server,don't need change
    nohup ./s-linux64 -localaddr :2555 -remoteaddr 127.0.0.1:3555 -overbose -enc aes-256-cfb@123qwe -loglev=1  >  s.log  2>&1  & 
##### shadowsocks-server,don't need change
    nohup ./go-shadowsocks2 -s 'ss://AEAD_CHACHA20_POLY1305:123qwe@:3555' -verbose >  sss.log  2>&1  & 