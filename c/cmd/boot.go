package main

import (
	gologging "github.com/godaner/geronimo/logger/go-logging"
	net2 "github.com/godaner/geronimo/net"
	"github.com/godaner/geronimo2ss/c/cfg"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
)

func main() {
	cfg.ParseFlag()
	gologging.SetLogger("geronimo2ss")
	go http.ListenAndServe(":8888", nil)
	logger := gologging.GetLogger("C")
	l, err := net.Listen("tcp", cfg.LocalAddr)
	if err != nil {
		panic(err)
	}
	for {

		lc, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			ip, port := addr(cfg.RemoteAddr)
			rc, err := net2.Dial(&net2.GAddr{
				IP:   ip,
				Port: int(port),
			})
			if err != nil {
				logger.Error("Dial err", err)
				lc.Close()
				return
			}

			go func() {

				_, err := io.Copy(rc, lc)
				if err != nil {
					logger.Error("io.Copy 1", err)
				}
				lc.Close()
				rc.Close()
				return

			}()

			go func() {
				_, err := io.Copy(lc, rc)
				if err != nil {
					logger.Error("io.Copy 2", err)
				}
				lc.Close()
				rc.Close()
				return

			}()

		}()

	}
}

// addr
func addr(addr string) (ip string, port int64) {
	ss := strings.Split(addr, ":")
	ip = ss[0]
	if len(ss) > 1 {
		port, _ = strconv.ParseInt(ss[1], 10, 64)
	}
	return ip, port
}
