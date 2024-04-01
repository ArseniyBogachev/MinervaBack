package middleware

import (
	"MinervaServer/db"
	"MinervaServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func TokenAuthMiddleware(c *gin.Context) {
	token := utils.ExtractToken(c)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	jwtToken, err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, err := utils.ExtractTokenClaims(jwtToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("userId", claims.UserId.Hex())

	c.Next()
}

func AdminAuthMiddleware(c *gin.Context) {
	userIdHex := c.GetString("userId")
	userId, err := primitive.ObjectIDFromHex(userIdHex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert user ID"})
		c.Abort()
		return
	}

	user := &db.User{}
	err = user.LoadById(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load user"})
		c.Abort()
		return
	}

	if !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	c.Next()
}

func FormOwnerOnlyMiddleware(c *gin.Context) {
	formIdHex := c.Param("formID")
	formId, err := primitive.ObjectIDFromHex(formIdHex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert form ID"})
		c.Abort()
		return
	}

	form := &db.Form{}

	if err = form.LoadById(formId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load form"})
		c.Abort()
		return
	}

	if form.OwnerID.Hex() != c.GetString("userId") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	c.Set("formId", formIdHex)
	c.Next()
}
