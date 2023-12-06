package controller

import (
	"golang-api-ulang/helper"
	"golang-api-ulang/model/web"
	"golang-api-ulang/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type SupplierControllerImpl struct {
	SupplierService service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) SupplierControler {
	return &SupplierControllerImpl{
		SupplierService: supplierService,
	}
}

func (controller *SupplierControllerImpl) Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	supplierCreateRequest := web.SupplierCreateRequest{}
	helper.ReadFromRequestBody(request, &supplierCreateRequest)
	supplierResponse := controller.SupplierService.Create(request.Context(), supplierCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   supplierResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	supplierUpdateRequest := web.SupplierUpdateRequest{}
	helper.ReadFromRequestBody(request, &supplierUpdateRequest)

	supplierId := paramas.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)

	supplierUpdateRequest.Id = int32(id)

	supplierResponse := controller.SupplierService.Update(request.Context(), supplierUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   supplierResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	supplierId := paramas.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)

	controller.SupplierService.Delete(request.Context(), int32(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	supplierId := paramas.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)

	supplierResponse := controller.SupplierService.FindById(request.Context(), int32(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   supplierResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	supplierResponses := controller.SupplierService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   supplierResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Pagination(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	pageRequest := request.URL.Query().Get("page")

	// * Check if pageRequest is null
	if pageRequest == "" {
		pageRequest = "1"
	}

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(pageRequest, "")
	var idSlice []string

	for _, item := range checkNumber {
		_, err := strconv.Atoi(item)
		if err == nil {
			idSlice = append(idSlice, item)
		}
	}
	page, err := strconv.Atoi(strings.Join(idSlice, ""))
	if err != nil {
		page = 1
	}

	supplierResponses, currentPage := controller.SupplierService.Pagination(request.Context(), int32(page))

	webResponse := web.WebResponse{
		Code:       200,
		Status:     "OK",
		Data:       supplierResponses,
		Pagination: currentPage,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
