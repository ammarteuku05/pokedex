package controller

import (
	"pokedex/repository"
	"pokedex/service"
	"pokedex/shared/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB = config.Connection()
	userRepository              = repository.NewUserRepository(DB)
	userService                 = service.NewUserService(userRepository)
	authService                 = service.NewAuthService()
	userControllers             = NewUserController(userService, authService)
	monsterRepository           = repository.NewMonsterRepository(DB)
	monsterService              = service.NewMonsterService(monsterRepository)
	monsterControllers          = NewMonsterController(monsterService)
)

func Route(r *gin.Engine) {
	r.POST("/user/register", userControllers.RegisterUserController)
	r.POST("/admin/register", userControllers.RegisterAdminController)
	r.POST("/login", userControllers.LoginUserController)
	r.POST("/create-monster", Middleware(authService), AdminMiddleware(userRepository), monsterControllers.CreateMonsterController)
	r.PUT("/update-monster/:id", Middleware(authService), AdminMiddleware(userRepository), monsterControllers.UpdateMonsterController)
	r.DELETE("/delete-monster", Middleware(authService), AdminMiddleware(userRepository), monsterControllers.DeleteMonsterController)
	r.GET("getall-types", monsterControllers.FindAllType)
	r.GET("getall-monsters", monsterControllers.FindAllMonsters)
}
