package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	_ "project/src/models"
	"project/src/router"
	_ "project/src/router"
	"project/src/setting"
	_ "project/src/setting"
)
func main() {
	go func() {
		// 开启服务
		s := &http.Server{
			Addr:           setting.GlobalConfig.Serve.ServerAddr + ":" + setting.GlobalConfig.Serve.ServerPort,
			Handler:        router.Engine,
			ReadTimeout:    60 * time.Second,
			WriteTimeout:   60 * time.Second,
			MaxHeaderBytes: 1 << 20, // 10M
		}
		logrus.Println("server listen on: ", s.Addr)
		//if setting.GlobalConfig.Serve.IsTLS { // 开启tls
		//	TLSConfig := &tls.Config{
		//		MinVersion:               tls.VersionTLS11,
		//		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		//		PreferServerCipherSuites: true,
		//		CipherSuites: []uint16{
		//			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		//			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		//			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		//			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		//		},
		//	}
		//
		//	TLSProto := make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
		//
		//	s.TLSConfig = TLSConfig
		//	s.TLSNextProto = TLSProto
		//
		//	if err := s.ListenAndServeTLS(setting.GlobalConfig.Serve.TLS.Cert, setting.GlobalConfig.Serve.TLS.Key); err != nil {
		//		logrus.Fatal(err)
		//	}
		//}
			if err := s.ListenAndServe(); err != nil {
				logrus.Fatal(err)
			}

	}()

	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
EXIT:
	for {
		sig := <-sc
		fmt.Printf("signal[%s]\n", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("exit")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
