// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"golang-api-ulang/app"
	"golang-api-ulang/controller"
	"golang-api-ulang/repository"
	"golang-api-ulang/service"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	bookRepository := repository.NewBookRepository()
	db := app.NewDB()
	validate := NewValidate()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	supplierRepository := repository.NewSupplierRepository()
	supplierService := service.NewSupplierService(supplierRepository, db, validate)
	supplierControler := controller.NewSupplierController(supplierService)
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(bookController, supplierControler, categoryController)
	server := NewServer(router)
	return server
}

// injector.go:

var bookSet = wire.NewSet(repository.NewBookRepository, service.NewBookService, controller.NewBookController)

var supplierSet = wire.NewSet(repository.NewSupplierRepository, service.NewSupplierService, controller.NewSupplierController)

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, controller.NewCategoryController)
