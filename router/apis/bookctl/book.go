package bookctl

import (
	"github.com/gin-gonic/gin"
	"reserve/models/book"
	"reserve/service/book_service"
	"reserve/tool/response"
	"reserve/tool/util"
)

func AddBook(c *gin.Context) {

	in := &book.AddBookRequest{}
	if err := util.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := book_service.NewBookService().Create(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func ListBook(c *gin.Context) {
	in := &book.ListBookRequest{}
	if err := util.ValidParams(c, in); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	in.UserId = util.UserId(c)
	date, err := book_service.NewBookService().List(in)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, date)
}
