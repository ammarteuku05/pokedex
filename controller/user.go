package controller

import (
	"pokedex/service"
	"pokedex/shared/dto"
	"pokedex/utils"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserServices
	authService service.AuthService
}

func NewUserController(userService service.UserServices, authService service.AuthService) *userController {
	return &userController{userService, authService}
}

func (h *userController) RegisterUserController(c *gin.Context) {
	var inputUser *dto.UserInputRegister

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := utils.SplitErrorInformation(err)
		responseError := utils.APIResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	inputUser.Role = "user"

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := utils.APIResponse(201, "Create new user succeed", newUser)
	c.JSON(201, response)
}

func (h *userController) RegisterAdminController(c *gin.Context) {
	var inputUser *dto.UserInputRegister

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := utils.SplitErrorInformation(err)
		responseError := utils.APIResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	inputUser.Role = "admin"

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := utils.APIResponse(201, "Create new user succeed", newUser)
	c.JSON(201, response)
}

func (h *userController) LoginUserController(c *gin.Context) {
	var inputLoginUser *dto.UserInputLogin

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := utils.SplitErrorInformation(err)
		responseError := utils.APIResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := utils.APIResponse(401, "Input data error", gin.H{"errors": err.Error()})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateToken(userData.Id)
	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}
	response := utils.APIResponse(200, "Login user succeed", gin.H{"token": token, "user": userData})
	c.JSON(200, response)
}
