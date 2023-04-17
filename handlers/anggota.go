package handlers

import (
	"net/http"
	"toko_mas_api/domain/anggota"

	"github.com/gin-gonic/gin"
)

type AnggotaHandlers struct {
	service anggota.IService
}

func NewAnggotaHandler(service anggota.IService) *AnggotaHandlers {
	return &AnggotaHandlers{service}
}

func (h *AnggotaHandlers) Register(c *gin.Context) {
	var input anggota.Inputan
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "inputan tidak sesuai",
		})
		return
	}
	result, err := h.service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error service",
		})
		return
	}
	c.JSON(http.StatusCreated, result)
}
