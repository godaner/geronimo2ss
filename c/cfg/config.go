package cfg

import "flag"

var LocalAddr string
var RemoteAddr string
var OverBose bool
var LogPath string
var LogLev int

func ParseFlag() {

	localaddr := flag.String("localaddr", "127.0.0.1:1555", "local addr")
	remoteaddr := flag.String("remoteaddr", "55.55.55.55:2555", "local addr")
	overbose := flag.Bool("overbose", false, "overBose")
	logpath := flag.String("logpath", "", "log path")
	loglev := flag.Int("loglev", 0, "log lev , option is : 0:CRIT,1:ERRO,2:WARN,3:NOTI,4:INFO,5:DEBU")
	flag.Parse()

	LocalAddr = *localaddr
	RemoteAddr = *remoteaddr
	OverBose = *overbose
	LogPath = *logpath
	LogLev = *loglev
}
