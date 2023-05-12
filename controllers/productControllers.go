package controllers

import (
	"DTSGolang/Kelas3/Sesi2Bagian2/database"
	"DTSGolang/Kelas3/Sesi2Bagian2/helpers"
	"DTSGolang/Kelas3/Sesi2Bagian2/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	errA := db.First(&User, "id = ?", userID).Error
	if errA != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "User Not Found",
			"message": errA.Error(),
		})

		return
	}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.User = &User

	err := db.Debug().Create(&Product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": Product,
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}

	productID, _ := strconv.Atoi(ctx.Param("productID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productID)

	err := db.Model(&Product).Where("id = ?", productID).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func GetProductById(ctx *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := db.Debug().Preload("User").First(&Product, "id = ?", uint(productID)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Book Not Found",
				"message": err.Error(),
			})

			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func GetProducts(ctx *gin.Context) {
	db := database.GetDB()

	Products := []models.Product{}

	err := db.Debug().Find(&Products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()

	Product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := db.Debug().Where("id = ?", uint(productID)).Delete(&Product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted product",
	})
}
