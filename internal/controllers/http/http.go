package http

import (
	docs "github.com/antony-ramos/cop-orm/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var usecase UserUsecase

// @BasePath /api/v1
type UserUsecase interface {
	CreateUser(body []byte) error
	GetUser(userID string) ([]byte, error)
	DeleteUser(userID string) error
	ModifyUser(userID string, body []byte) error
	GetAllUsers() ([]byte, error)
	CreateGroup(body []byte) error
}

func NewHTTPController(uc UserUsecase) *gin.Engine {
	usecase = uc

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/users")
		{
			eg.POST("/", CreateUserHandler)

			eg.GET("/", GetUsersHandler)

			eg.GET("/:id", GetUserHandler)

			eg.DELETE("/:id", DeleteUserHandler)

			eg.PUT("/:id", ModifyUserHandler)
		}
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func Run(usecase UserUsecase) {
	r := NewHTTPController(usecase)
	r.Run()
}
