package controller

import (
	"golang-api-ulang/exception"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/web"
	"golang-api-ulang/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryCreateRequest.Nama = strings.Trim(categoryCreateRequest.Nama, " ")
	categoryCreateRequest.Nama = strings.ToLower(categoryCreateRequest.Nama)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := paramas.ByName("categoryId")

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(categoryId, "")
	var idSlice []string

	for _, item := range checkNumber {
		_, err := strconv.Atoi(item)
		if err == nil {
			idSlice = append(idSlice, item)
		}
	}

	id, err := strconv.Atoi(strings.Join(idSlice, ""))
	if err != nil {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}

	categoryUpdateRequest.Id = int32(id)
	categoryUpdateRequest.Nama = strings.Trim(categoryUpdateRequest.Nama, " ")
	categoryUpdateRequest.Nama = strings.ToLower(categoryUpdateRequest.Nama)

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	categoryId := paramas.ByName("categoryId")

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(categoryId, "")
	var idSlice []string

	for _, item := range checkNumber {
		_, err := strconv.Atoi(item)
		if err == nil {
			idSlice = append(idSlice, item)
		}
	}

	id, err := strconv.Atoi(strings.Join(idSlice, ""))
	if err != nil {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}

	controller.CategoryService.Delete(request.Context(), int32(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	categoryId := paramas.ByName("categoryId")

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(categoryId, "")
	var idSlice []string

	for _, item := range checkNumber {
		_, err := strconv.Atoi(item)
		if err == nil {
			idSlice = append(idSlice, item)
		}
	}

	id, err := strconv.Atoi(strings.Join(idSlice, ""))
	if err != nil {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}

	categoryResponse := controller.CategoryService.FindById(request.Context(), int32(id))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Pagination(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
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

	categoryResponses, currentPage := controller.CategoryService.Pagination(request.Context(), int32(page))

	webResponse := web.WebResponse{
		Code:       200,
		Status:     "OK",
		Data:       categoryResponses,
		Pagination: currentPage,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
