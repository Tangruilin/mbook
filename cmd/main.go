package main

import (
	"fmt"
	"github.com/pkg/errors"
	"mbook/global"
	"mbook/internal/model"
	"mbook/internal/routers"
	"net/http"
	"time"
)

func init() {
	err := setupDBEngine()
	if err != nil {
		fmt.Printf("Init error: %+v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//setupDBEngine is used to link the MySQL
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return errors.WithMessage(err, "setupDBEngine error")
	}
	return nil
}
