package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.camelit.com/walofz/cicsupport-api/config"
	"gitlab.camelit.com/walofz/cicsupport-api/controller"
	"gitlab.camelit.com/walofz/cicsupport-api/middleware"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
	"gitlab.camelit.com/walofz/cicsupport-api/service"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB           = config.SetupDB()
	jwtService service.JWTService = service.NewJWTService()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, jwtService)

	areagroupRepository repository.AreagroupRepository = repository.NewAreaGroupRepository(db)
	areagroupService    service.AreagroupService       = service.NewAreaGroupService(areagroupRepository)
	areagroupController controller.AreagroupController = controller.NewAreagroupController(areagroupService, jwtService)
)

func main() {
	defer config.CloseDBConn(db)

	server := gin.Default()

	//cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	//corsConfig.AllowHeaders = []string{"Content-Type", "application/json"}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	authRoute := server.Group("api/auth", authController.Login)
	{
		authRoute.POST("/login")
	}

	userRoute := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoute.GET("/profile", userController.Profile)
	}

	areagroupRoute := server.Group("api/areagroup", middleware.AuthorizeJWT(jwtService))
	{
		areagroupRoute.GET("/listall", areagroupController.Listall)
		areagroupRoute.POST("/create", areagroupController.Create)
		areagroupRoute.PUT("update/:id", areagroupController.Update)
		areagroupRoute.DELETE("/delete/:id", areagroupController.Delete)
	}

	server.Run(":1223")
}
