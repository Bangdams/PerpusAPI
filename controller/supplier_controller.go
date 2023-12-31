package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SupplierControler interface {
	Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Pagination(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
}
