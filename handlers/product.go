package handlers

import (
	"net/http"
	daftarproduk "toko_mas_api/domain/daftar_produk"
	"toko_mas_api/helper"

	"github.com/gin-gonic/gin"
)

type daftarProductHandlers struct {
	service daftarproduk.Service
}

func NewDaftarProdukHandler(service daftarproduk.Service) *daftarProductHandlers {
	return &daftarProductHandlers{service}
}

func (h *daftarProductHandlers) AddProduk(c *gin.Context) {
	var input daftarproduk.ProdukInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// jsonResponse := helper.ApiResponse("Bad Request", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, errService := h.service.Add(input)
	if errService != nil {
		jsonResponse := helper.ApiResponse("Internal Server Error", errService)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}
	response := helper.ApiResponseJson("Insert Data failed", http.StatusCreated, "success", res)
	c.JSON(http.StatusCreated, response)

}

func (h *daftarProductHandlers) ListProduk(c *gin.Context) {

	res, err := h.service.GetAll()
	if err != nil {
		jsonResponse := helper.ApiResponse("Internal ServerError", err)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	jsonResponse := helper.ApiResponse("List of Type", res)
	c.JSON(http.StatusOK, jsonResponse)
}
