package app

import (
	"golang-api-ulang/controller"
	"golang-api-ulang/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController, supplierController controller.SupplierControler, categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// Router Books
	router.GET("/api/books", bookController.Pagination)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	// Router History Books
	router.GET("/api/report-books", bookController.ReportPagination)

	// Router Suppliers
	router.GET("/api/suppliers", supplierController.Pagination)
	router.GET("/api/suppliers/:supplierId", supplierController.FindById)
	router.POST("/api/suppliers", supplierController.Create)
	router.PUT("/api/suppliers/:supplierId", supplierController.Update)
	router.DELETE("/api/suppliers/:supplierId", supplierController.Delete)

	// Router Categories
	router.GET("/api/categories", categoryController.Pagination)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
