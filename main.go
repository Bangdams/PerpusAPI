package main

import (
	"golang-api-ulang/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(Handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: Handler,
	}
}

func NewValidate() *validator.Validate {
	return validator.New()
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
