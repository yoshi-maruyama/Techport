package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/domain"
)

type BlogHandler struct {
	BlgUsecase domain.BlogUsecase
}

func NewBlogHandler(engine *gin.Engine, blgUcase domain.BlogUsecase) {
	handler := &BlogHandler{
		BlgUsecase: blgUcase,
	}
	engine.GET("api/v1/blogs/:id", handler.GetById)
	engine.POST("api/v1/blogs", handler.CreateBlog)
}

func (h *BlogHandler) GetById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blog, err := h.BlgUsecase.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("cant get blog from id: %d, error is %v", id, err),
		})
	}
	c.JSON(200, blog)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog = domain.CreateBlog{}
	if err := c.Bind(&blog); err != nil {
		fmt.Printf("err:%v", err)
	}
	if err := h.BlgUsecase.CreateBlog(&blog); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("cant get blog by %v", err),
		})
	}
	c.JSON(200, blog)
}