package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiVersion = "/v1"

type UserUsecase interface {
	CreateUser(body []byte) error
	GetUser(userID string) ([]byte, error)
	DeleteUser(userID string) error
	ModifyUser(userID string, body []byte) error
	GetAllUsers() ([]byte, error)
}

func NewHTTPController(usecase UserUsecase) *gin.Engine {
	r := gin.Default()

	r.POST("/api/"+apiVersion+"/users", func(c *gin.Context) {
		// Read the request body
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read request body",
			})
			return
		}
		// Call the createUser function from the usecase package
		err = usecase.CreateUser(body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
		})
	})

	r.GET("/api/"+apiVersion+"/users", func(c *gin.Context) {
		// Call the getAllUsers function from the usecase package
		users, err := usecase.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get all users",
			})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.GET("/api/"+apiVersion+"/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		// Call the getUser function from the usecase package
		user, err := usecase.GetUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get user",
			})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.DELETE("/api/"+apiVersion+"/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		// Call the deleteUser function from the usecase package
		err := usecase.DeleteUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
		})
	})

	r.PUT("/api/"+apiVersion+"/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		// Read the request body
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read request body",
			})
			return
		}
		// Call the modifyUser function from the usecase package
		err = usecase.ModifyUser(userID, body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to modify user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User modified successfully",
		})
	})

	return r
}

func Run(usecase UserUsecase) {
	r := NewHTTPController(usecase)
	r.Run()
}
