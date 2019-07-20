package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofoody/auth-service/pkg/config"
	"github.com/gofoody/auth-service/pkg/ctrl"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.New()

	initLogger(config.GetLogLevel())

	router := mountEndpoints()
	startService(config.GetHttpPort(), router)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func mountEndpoints() *mux.Router {
	r := mux.NewRouter()

	statusCtrl := ctrl.NewStatusCtrl()
	r.HandleFunc("/api/status", statusCtrl.Show)

	loginCtrl := ctrl.NewLoginCtrl()
	r.HandleFunc("/api/login", loginCtrl.Login)

	return r
}

func startService(port int, router *mux.Router) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Infof("auth service running at:%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start auth service, error:%v", err)
	}
}
