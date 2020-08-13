package cfg

import "flag"

var LocalAddr string
var RemoteAddr string
var OverBose bool
var LogPath string
var LogLev int
var Enc string

func ParseFlag() {

	localaddr := flag.String("localaddr", "127.0.0.1:1555", "local addr")
	remoteaddr := flag.String("remoteaddr", "55.55.55.55:2555", "local addr")
	overbose := flag.Bool("overbose", false, "overBose")
	logpath := flag.String("logpath", "", "log path")
	loglev := flag.Int("loglev", 0, "log lev , option is : 0:CRIT,1:ERRO,2:WARN,3:NOTI,4:INFO,5:DEBU")
	enc := flag.String("enc", "", "encryption : method@password , method option is : aes-128-cfb,aes-192-cfb,aes-256-cfb,aes-128-ctr,aes-192-ctr,aes-256-ctr,des-cfb,bf-cfb,cast5-cfb,rc4-md5,rc4-md5-6,chacha20,chacha20-ietf,salsa20")
	flag.Parse()

	LocalAddr = *localaddr
	RemoteAddr = *remoteaddr
	OverBose = *overbose
	LogPath = *logpath
	LogLev = *loglev
	Enc = *enc
}
