package controller

import (
	"pokedex/repository"
	"pokedex/service"
	"pokedex/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) == 0 {
			errorResponse := utils.APIResponse(401, "Unauthorized user", gin.H{"error": "unauthorize user (no header)"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		// eksekusi code untuk mengecek apakah token itu valid dari server kita atau tidak
		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			errorResponse := utils.APIResponse(401, "Unauthorized user", gin.H{"error": "unauthorize user (token unvalidated)"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			errorResponse := utils.APIResponse(401, "Unauthorized user", gin.H{"error": "unauthorize user"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}
		userID := int(claim["id"].(float64))

		c.Set("currentUser", userID)
	}

}

func AdminMiddleware(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogin := c.MustGet("currentUser").(int)
		// userId := strconv.Itoa(userLogin)
		user, err := userRepository.FindByID(userLogin)

		if err != nil {
			errorResponse := utils.APIResponse(500, "Internal Server Error", gin.H{"error": "internal server error"})

			c.AbortWithStatusJSON(500, errorResponse)

			return
		}

		if user.Role != "admin" {
			errorResponse := utils.APIResponse(401, "Unauthorized user", gin.H{"error": "user login is not admin"})

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
