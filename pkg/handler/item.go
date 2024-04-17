package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {

}

func (h *Handler) getAllItems(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
