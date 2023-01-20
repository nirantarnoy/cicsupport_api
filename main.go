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

	areazoneRepository repository.AreazoneRepository = repository.NewAreazoneRepository(db)
	areazoneService    service.AreazoneService       = service.NewAreazoneService(areazoneRepository)
	areazoneController controller.AreazoneController = controller.NewAreazoneController(areazoneService, jwtService)

	topicItemRepository repository.TopiceitemRepository = repository.NewTopicitemRepository(db)
	topicItemService    service.TopicItemService        = service.NewTopicItemService(topicItemRepository)
	topicItemController controller.TopicItemController  = controller.NewTopicItemController(topicItemService, jwtService)

	planRepository repository.PlanRepository = repository.NewPlanRepository(db)
	planService    service.PlanService       = service.NewPlanService(planRepository)
	planController controller.PlanController = controller.NewPlanController(planService, jwtService)

	teaminspectionitemRepository repository.TeaminspectionitemRepository = repository.NewTeaminspectionitemRepository(db)
	teaminspectionitemService    service.TeaminspectionitemService       = service.NewTeaminspectionitemService(teaminspectionitemRepository)
	teaminspectionitemController controller.TeaminspectionitemController = controller.NewIteminspectionitemController(teaminspectionitemService, jwtService)

	carinspectRepository repository.CarinspecRepository = repository.NewCarinspecRepository(db)
	carinspectService    service.CarinspecSevice        = service.NewCarinspecService(carinspectRepository)
	carinspectController controller.CarinspecController = controller.NewCarinspecController(carinspectRepository, jwtService)

	carRepository repository.CarRepository = repository.NewCarRepository(db)
	carService    service.CarService       = service.NewCarService(carRepository)
	carController controller.CarController = controller.NewCarController(carService, jwtService)

	teamNotifyRepository repository.TeamNotifyRepository = repository.NewTeamnotifyRepository(db)
	teamNotifyServicer   service.TeamNotifyService       = service.NewTeamnotifyService(teamNotifyRepository)
	teamNotifyController controller.TeamNofiryController = controller.NewTeamNotifyController(teamNotifyServicer, jwtService)
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
		userRoute.GET("/profile/:id", userController.Profile)
		userRoute.GET("/teammember/:id", userController.FindUserTeam)

	}

	areagroupRoute := server.Group("api/areagroup", middleware.AuthorizeJWT(jwtService))
	{
		areagroupRoute.GET("/listall", areagroupController.Listall)
		areagroupRoute.POST("/create", areagroupController.Create)
		areagroupRoute.PUT("update/:id", areagroupController.Update)
		areagroupRoute.DELETE("/delete/:id", areagroupController.Delete)
	}

	areazoneRoute := server.Group("api/areazone", middleware.AuthorizeJWT(jwtService))
	{
		areazoneRoute.GET("/listall", areazoneController.Findall)
		areazoneRoute.POST("/create", areazoneController.Create)
		areazoneRoute.PUT("/update/:id", areazoneController.Update)
		areazoneRoute.DELETE("/delete/:id", areazoneController.Delete)
	}

	topicItemRoute := server.Group("api/topicitem", middleware.AuthorizeJWT(jwtService))
	{
		topicItemRoute.GET("/findtopicbyplan/:id", topicItemController.FindTopicByPlan)
	}

	planRoute := server.Group("api/plan", middleware.AuthorizeJWT(jwtService))
	{
		planRoute.GET("/findplan/:id", planController.FindPlanByTeam)
		planRoute.POST("/addinspection", planController.AddInspection)
	}

	teaminspectionitemRoute := server.Group("api/teaminspectionitem", middleware.AuthorizeJWT(jwtService))
	{
		teaminspectionitemRoute.GET("/findbyteam/:id", teaminspectionitemController.FindInspectionItem)
		teaminspectionitemRoute.POST("/findtransbyemp", teaminspectionitemController.FindTransByEmp)
		teaminspectionitemRoute.GET("/findtranshistorybyemp/:id", teaminspectionitemController.FindTransHistoryByEmp)
	}

	carinspectionRoute := server.Group("api/carinspection", middleware.AuthorizeJWT(jwtService))
	{
		carinspectionRoute.GET("/findall", carinspectController.NonconformList)
	}
	carRoute := server.Group("api/car", middleware.AuthorizeJWT(jwtService))
	{
		carRoute.POST("/createcar", carController.CreateCar)
		carRoute.GET("/listcarbyemp/:id", carController.ListCarByEmpId)
		carRoute.POST("/closecar", carController.CloseCar)
	}
	teamNotifyRoute := server.Group("api/teamnotify", middleware.AuthorizeJWT(jwtService))
	{
		teamNotifyRoute.GET("/findempnotify/:id", teamNotifyController.FindEmpNotify)
	}

	server.Run(":1223")
}
