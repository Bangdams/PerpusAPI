//go:build wireinject
// +build wireinject

package main

import (
	"golang-api-ulang/app"
	"golang-api-ulang/controller"
	"golang-api-ulang/repository"
	"golang-api-ulang/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var bookSet = wire.NewSet(
	repository.NewBookRepository,
	service.NewBookService,
	controller.NewBookController,
)

var supplierSet = wire.NewSet(
	repository.NewSupplierRepository,
	service.NewSupplierService,
	controller.NewSupplierController,
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidate,
		bookSet,
		supplierSet,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)
	return nil
}
