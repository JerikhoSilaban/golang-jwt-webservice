package middlewares

import (
	"DTSGolang/Kelas3/Sesi2Bagian2/database"
	"DTSGolang/Kelas3/Sesi2Bagian2/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(ctx.Param("productID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})

			return
		}

		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		admin, ok := userData["admin"].(bool)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": "Unable to retrieve admin status",
			})
		}

		if !admin {
			userID := uint(userData["id"].(float64))
			if Product.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to acces this data",
				})

				return
			}
		}

		ctx.Next()
	}
}

func ProductAuthorizationAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data",
			})

			return
		}

		ctx.Next()
	}
}
