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
	// Parse JSON request body into Inplogin struct
	var loginReq anggota.InpLogin
	err := c.BindJSON(&loginReq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Check if username and password are valid
	if loginReq.Username == "myusername" && loginReq.Password == "mypassword" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
