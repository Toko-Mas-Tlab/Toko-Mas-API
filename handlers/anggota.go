package handlers

import (
	"net/http"
	"toko_mas_api/domain/anggota"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
)

type AnggotaHandlers struct {
	service    anggota.IService
	jwtService middleware.IService
}

func NewAnggotaHandler(service anggota.IService, jwtService middleware.IService) *AnggotaHandlers {
	return &AnggotaHandlers{service, jwtService}
}

func (h *AnggotaHandlers) Register(c *gin.Context) {
	var input anggota.Inputan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, errS := h.service.Register(input)
	if errS != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errS)
		return
	}
	c.JSON(http.StatusCreated, res)
}
