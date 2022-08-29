package main

import (
	"context"
	"emall/components"
	"emall/providers"
	"emall/task"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var port int
	var envPath string
	wd, err := os.Getwd()
	ctx, cancel := context.WithCancel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	flag.IntVar(&port, "p", 8080, "http server port")
	flag.StringVar(&envPath, "c", wd+".env", "config file path")
	//init components
	rdsCFG := components.RedisConfig{
		ConnectionName: "order-redis",
		Port:           6379,
		PoolSize:       100,
	}
	providers.RedisConnector, err = components.NewRedisClient(&rdsCFG)
	if err != nil {
		log.Fatal(err)
	}

	dbCFG := components.GormConfig{
		DBName:   "zt_audit",
		Host:     "rm-2zezu40s4z8q11w0h.mysql.rds.aliyuncs.com",
		Port:     "3306",
		Password: "xie",
		UserName: "uLHbO1WpsQMgnwrY",
	}
	providers.DBconnector, err = components.NewGormDB(&dbCFG)
	if err != nil {
		log.Fatal(err)
	}
	//async tasks
	cleanTasks := task.Start(ctx, cancel)
	//http server

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: registerRouters(),
		// ReadTimeout:    0,
		// WriteTimeout:   0,
		// IdleTimeout:    0,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	cleanTasks()
	//优雅退出
	log.Println("Every thing done, bye bitch!")

}
