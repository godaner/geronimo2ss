package main

import (
	"github.com/godaner/logger"
	gn "github.com/godaner/geronimo/net"
	"github.com/godaner/geronimo2ss/s/cfg"
	loggerfac "github.com/godaner/logger/factory"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
)

func main() {
	cfg.ParseFlag()
	loggerfac.Init("geronimo2ss", loggerfac.SetLevel(logger.Level(cfg.LogLev)), loggerfac.SetLogPath(cfg.LogPath))
	logger := loggerfac.GetLogger("S")
	go http.ListenAndServe(":9999", nil)
	ip, port := addr(cfg.LocalAddr)
	l, err := gn.Listen(&gn.GAddr{
		IP:   ip,
		Port: int(port),
	}, gn.SetOverBose(cfg.OverBose),gn.SetEnc(cfg.Enc))

	if err != nil {
		panic(err)
	}

	for {
		lc, err := l.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}

		go func() {

			rc, err := net.Dial("tcp", cfg.RemoteAddr)
			if err != nil {
				logger.Error(err)
				lc.Close()
				return
			}

			go func() {
				_, err = io.Copy(lc, rc)
				if err != nil {
					logger.Error("io.Copy 1", err)
				}

				rc.Close()
				lc.Close()
				return
			}()

			go func() {
				_, err = io.Copy(rc, lc)
				if err != nil {
					logger.Error("io.Copy 2", err)
				}

				rc.Close()
				lc.Close()
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
