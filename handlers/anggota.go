package handlers

import (
	"errors"
	"net/http"
	"toko_mas_api/domain/anggota"
	"toko_mas_api/helper"
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

func (h *AnggotaHandlers) Login(c *gin.Context) {
	var input anggota.InpLogin

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res, errS := h.service.Login(input)
	if errS != nil {
		c.AbortWithStatusJSON(400, errS)
		return
	}
	if res.ID == 0 {
		c.AbortWithStatusJSON(400, errors.New("user not found"))
		return
	}

	accessToken, errToken := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "ACCESS_TOKEN")
	if errToken != nil {
		c.AbortWithStatusJSON(400, errToken)
		return
	}

	refreshToken, errToken := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "REFRESH_TOKEN")
	if errToken != nil {
		c.AbortWithStatusJSON(400, errToken)
		return
	}

	// set Cookie
	c.SetCookie("refresh_token", refreshToken, 3600*12, "/", "", true, true)

	resData := anggota.LoginResponseFormatter(res, accessToken)
	response := helper.ApiResponse("Berhasil Login", resData)
	c.JSON(http.StatusOK, response)
}
