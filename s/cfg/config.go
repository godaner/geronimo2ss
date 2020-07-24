package cfg

import "flag"

var LocalAddr string
var RemoteAddr string
func ParseFlag() {

	localaddr := flag.String("localaddr", ":2555", "local addr")
	remoteaddr := flag.String("remoteaddr", "127.0.0.1:3555", "local addr")
	flag.Parse()

	LocalAddr=*localaddr
	RemoteAddr=*remoteaddr
}
