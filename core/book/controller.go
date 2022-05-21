package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tulioguaraldo/online-bookstore/model"
)

type bookController struct {
	service InterfaceBookService
}

func NewBookController(service InterfaceBookService) bookController {
	return bookController{
		service: service,
	}
}

func (c *bookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.getAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, books)
}

func (c *bookController) GetBookByData(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := model.Book{Id: bookId}

	response := BookResponse{}
	BookToResponse(&book, &response)

	req, err := c.service.show(&book)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, req)
}

func (c *bookController) InsertBook(ctx *gin.Context) {
	bookInsertion := model.Book{}
	if err := ctx.ShouldBindJSON(&bookInsertion); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := BookResponse{}
	BookToResponse(&bookInsertion, &response)

	_, err := c.service.create(&bookInsertion)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "book inserted successfully",
		"book":    response,
	})
}
