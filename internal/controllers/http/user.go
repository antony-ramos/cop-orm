package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler creates a new user.
// @Summary Create a new user
// @Description Create a new user with the provided request body
// @Tags User
// @Accept json
// @Produce json
// @Param body body string true "User object"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to read request body",
		})
		return
	}
	// Call the createUser function from the usecase package
	err = usecase.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User created successfully",
	})
}

// GetUserHandler retrieves a user by ID.
// @Summary Get a user by ID
// @Description Get a user by the provided ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	userID := c.Param("id")
	// Call the getUser function from the usecase package
	user, err := usecase.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsersHandler retrieves all users.
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetUsersHandler(c *gin.Context) {
	// Call the getAllUsers function from the usecase package
	users, err := usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get users",
		})
		return
	}
	c.JSON(http.StatusOK, string(users))
}

// DeleteUserHandler deletes a user by ID.
// @Summary Delete a user by ID
// @Description Delete a user by the provided ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	// Call the deleteUser function from the usecase package
	err := usecase.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// ModifyUserHandler modifies a user by ID.
// @Summary Modify a user by ID
// @Description Modify a user by the provided ID and request body
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body string true "User object"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func ModifyUserHandler(c *gin.Context) {
	userID := c.Param("id")
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to read request body",
		})
		return
	}
	// Call the modifyUser function from the usecase package
	err = usecase.ModifyUser(userID, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to modify user",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User modified successfully",
	})
}
