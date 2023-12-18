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

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFromRequestBody(request, &bookCreateRequest)

	bookCreateRequest.Nama = strings.ToLower(bookCreateRequest.Nama)
	bookCreateRequest.Nama = strings.Trim(bookCreateRequest.Nama, " ")

	book := controller.BookService.Create(request.Context(), bookCreateRequest)
	bookResponse := controller.BookService.FindByName(request.Context(), book.Nama)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookUpdateRequest)

	bookId := paramas.ByName("bookId")
	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(bookId, "")
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

	bookUpdateRequest.Id = int32(id)
	bookUpdateRequest.Nama = strings.ToLower(bookUpdateRequest.Nama)
	bookUpdateRequest.Nama = strings.Trim(bookUpdateRequest.Nama, " ")

	book := controller.BookService.Update(request.Context(), bookUpdateRequest)
	bookResponse := controller.BookService.FindByName(request.Context(), book.Nama)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	bookId := paramas.ByName("bookId")

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(bookId, "")
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

	controller.BookService.Delete(request.Context(), int32(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	bookId := paramas.ByName("bookId")

	// * Check if there are numbers in Qoury
	checkNumber := strings.Split(bookId, "")
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

	bookResponse := controller.BookService.FindById(request.Context(), int32(id))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Pagination(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params) {
	pageRequest := request.URL.Query().Get("page")
	nameQuery := request.URL.Query().Get("name")

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

	bookResponses, currentPage := controller.BookService.Pagination(request.Context(), int32(page), nameQuery)

	webResponse := web.WebResponse{
		Code:       200,
		Status:     "OK",
		Data:       bookResponses,
		Pagination: currentPage,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
