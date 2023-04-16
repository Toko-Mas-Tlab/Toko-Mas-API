package handlers

import (
	"net/http"
	"strconv"
	jenisbarang "toko_mas_api/domain/jenis_barang"
	"toko_mas_api/helper"

	"github.com/gin-gonic/gin"
)

type JenisBarangHandlers struct {
	service jenisbarang.IService
}

func NewJenisBarangHandler(service jenisbarang.IService) *JenisBarangHandlers {
	return &JenisBarangHandlers{service}
}

func (h *JenisBarangHandlers) AddNewType(c *gin.Context) {
	var input jenisbarang.Inputan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		jsonResponse := helper.ApiResponse("Bad Request", err)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	res, errService := h.service.Add(input)
	if errService != nil {
		jsonResponse := helper.ApiResponse("Internal Server Error", errService)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	jsonResponse := helper.ApiResponse("Created", res)
	c.JSON(http.StatusCreated, jsonResponse)
}

func (h *JenisBarangHandlers) ListJenisBarang(c *gin.Context) {
	sort := c.Query("sort")
	if sort == "" {
		sort = "desc"
	}

	res, err := h.service.GetAll(sort)
	if err != nil {
		jsonResponse := helper.ApiResponse("Internal ServerError", err)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	formatter := jenisbarang.ResponseFormatter(res)
	jsonResponse := helper.ApiResponse("List of Type", formatter)
	c.JSON(http.StatusOK, jsonResponse)
}

func (h *JenisBarangHandlers) UpdateJenisBarang(c *gin.Context) {
	var input jenisbarang.Inputan

	id, err := strconv.Atoi(c.Param("id_jenis_barang"))
	if err != nil {
		jsonResponse := helper.ApiResponse("Bad Request", err)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	errBinding := c.ShouldBindJSON(&input)
	if errBinding != nil {
		jsonResponse := helper.ApiResponse("Bad Request", errBinding)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	res, errUpdate := h.service.Update(id, input)
	if err != nil {
		jsonResponse := helper.ApiResponse("Internal Server Error", errUpdate)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}
	if res.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "ID Not Found",
		})
		return
	}

	jsonResponse := helper.ApiResponse("Updated", res)
	c.JSON(http.StatusCreated, jsonResponse)
}
