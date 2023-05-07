package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	bentukbarang "toko_mas_api/domain/bentuk_barang"
	"toko_mas_api/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type BentukBarangHandlers struct {
	service bentukbarang.IService
}

func NewBentukBarangHandler(service bentukbarang.IService) *BentukBarangHandlers {
	return &BentukBarangHandlers{service}
}

func (h *BentukBarangHandlers) AddNewShape(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var input bentukbarang.Inputan

	err = c.ShouldBindJSON(&input)
	if err != nil {
		jsonResponse := helper.ApiResponse("Bad Request", err)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	// ======================================================================
	authHeader := c.GetHeader("Authorization")
	// Periksa skema autentikasi
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Ambil token JWT dari nilai Authorization
	tokenString := authHeader[7:]
	// Parse token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Isi dengan secret key JWT Anda
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})
	// Periksa kesalahan parsing token
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Periksa apakah token valid dan tidak kadaluarsa
	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Ambil nilai payload dari token
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	fmt.Println(id)

	// ====================================================================

	res, errService := h.service.Add(input, int(id))
	if errService != nil {
		jsonResponse := helper.ApiResponse("Internal Server Error", errService)
		c.JSON(http.StatusInternalServerError, jsonResponse)
		return
	}

	jsonResponse := helper.ApiResponse("Created", res)
	c.JSON(http.StatusCreated, jsonResponse)
}

func (h *BentukBarangHandlers) ListBentukBarang(c *gin.Context) {
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

	// formatter := jenisbarang.ResponseFormatter()
	jsonResponse := helper.ApiResponse("List of Type", res)
	c.JSON(http.StatusOK, jsonResponse)
}

func (h *BentukBarangHandlers) UpdateBentukBarang(c *gin.Context) {
	var input bentukbarang.Inputan

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonResponse := helper.ApiResponse("err in atoi ", err)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	errBinding := c.ShouldBindJSON(&input)
	if errBinding != nil {
		jsonResponse := helper.ApiResponse("err in binding", errBinding)
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
