package controller

import (
	"books/dto"
	"books/entity"
	"books/helper"
	"books/service"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookService(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c *bookController) All(context *gin.Context) {
	if !Authentication(context.GetHeader("X-API-KEY")) {
		fmt.Println("Unauthenticated")
		response := helper.BuildErrorResponse("Unauthorized", "You are not allowed to access this", nil)
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	var book []entity.Book = c.bookService.All()
	response := helper.BuildResponse(true, "OK", book)
	context.JSON(http.StatusOK, response)
}

func (c *bookController) FindByID(context *gin.Context) {
	if !Authentication(context.GetHeader("X-API-KEY")) {
		fmt.Println("Unauthenticated")
		response := helper.BuildErrorResponse("Unauthorized", "You are not allowed to access this", nil)
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		fmt.Println(err.Error())
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}

	var book entity.Book = c.bookService.FindByID(id)
	if book == (entity.Book{}) {
		response := helper.BuildErrorResponse("No book was found", "No data with the id was found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		response := helper.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, response)
	}

}

func (c *bookController) Create(context *gin.Context) {
	var createBookDto dto.BookRequest
	if !Authentication(context.GetHeader("X-API-KEY")) {
		fmt.Println("Unauthenticated")
		response := helper.BuildErrorResponse("Unauthorized", "You are not allowed to access this", nil)
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	errDto := context.ShouldBind(&createBookDto)
	if errDto != nil {
		fmt.Println(errDto.Error())
		response := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} else {
		result := c.bookService.Create(createBookDto)
		if !result {
			response := helper.BuildErrorResponse("Failed to create book", "Failed to create book", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.BuildResponse(true, "OK", "Success create book")
		context.JSON(http.StatusOK, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var updateBookDto dto.UpdateBookRequest
	if !Authentication(context.GetHeader("X-API-KEY")) {
		fmt.Println("Unauthenticated")
		response := helper.BuildErrorResponse("Unauthorized", "You are not allowed to access this", nil)
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	paramId := context.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("id not found on parameter", "id not found on parameter", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	}
	errDto := context.ShouldBind(&updateBookDto)
	if errDto != nil {
		fmt.Println(errDto.Error())
		response := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	var book entity.Book = c.bookService.FindByID(id)
	if book == (entity.Book{}) {
		response := helper.BuildErrorResponse("No book was found", "No data with the id was found", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		result := c.bookService.Update(id, updateBookDto)
		if !result {
			response := helper.BuildErrorResponse("Failed to update book", "Failed to update book", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.BuildResponse(true, "Success update book", updateBookDto)
		context.JSON(http.StatusOK, response)
	}
}

func (c *bookController) Delete(context *gin.Context) {
	if !Authentication(context.GetHeader("X-API-KEY")) {
		fmt.Println("Unauthenticated")
		response := helper.BuildErrorResponse("Unauthorized", "You are not allowed to access this", nil)
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		fmt.Println(err.Error())
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}

	var book entity.Book = c.bookService.FindByID(id)
	if book == (entity.Book{}) {
		response := helper.BuildErrorResponse("No book was found", "No data with the id was found", book)
		context.JSON(http.StatusNotFound, response)
	} else {
		result := c.bookService.Delete(id)
		if !result {
			response := helper.BuildErrorResponse("Failed to delete book", "Failed to delete book", book)
			context.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.BuildResponse(true, "Success delete book", book)
		context.JSON(http.StatusOK, response)
	}
}

func Authentication(apiKey string) bool {
	if apiKey != os.Getenv("SECRET_KEY") {
		return false
	}
	return true
}
