package handlers

import (
	"net/http"
	"toko_mas_api/domain/anggota"
	"toko_mas_api/helper"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
)

type anggotaHandlers struct {
	service    anggota.IService
	jwtService middleware.IService
}

func NewAnggotaHandler(service anggota.IService, jwtService middleware.IService) *anggotaHandlers {
	return &anggotaHandlers{service, jwtService}
}

func (h *anggotaHandlers) Register(c *gin.Context) {
	var input anggota.Inputan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponseJson("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res, errS := h.service.Register(input)
	if errS != nil {
		response := helper.ApiResponseJson("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "ACCESS_TOKEN")
	if errS != nil {
		response := helper.ApiResponseJson("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := anggota.FormatRegister(res, token)

	response := helper.ApiResponseJson("Account has been registered", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *anggotaHandlers) Login(c *gin.Context) {
	var input anggota.InputLogin

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponseJson("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, errS := h.service.Login(input)
	if errS != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponseJson("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "ACCESS_TOKEN")
	if errS != nil {
		response := helper.ApiResponseJson("Login failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// accessToken, errToken := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "ACCESS_TOKEN")
	// if errToken != nil {
	// 	c.AbortWithStatusJSON(400, errToken)
	// 	return
	// }

	refreshToken, errToken := h.jwtService.GenerateTokenJWT(res.ID, res.Username, 10, "REFRESH_TOKEN")
	if errToken != nil {
		c.AbortWithStatusJSON(400, errToken)
		return
	}

	// set Cookie
	c.SetCookie("refresh_token", refreshToken, 3600*12, "/", "", true, true)

	formatter := anggota.LoginResponseFormatter(res, token)
	response := helper.ApiResponseJson("Successfully login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *anggotaHandlers) ListAnggota(c *gin.Context) {
	res, err := h.service.GetAll()
	if err != nil {
		jsonResponse := helper.ApiResponse("Internal ServerError", err)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	jsonResponse := helper.ApiResponse("List of Type", res)
	c.JSON(http.StatusOK, jsonResponse)
}
