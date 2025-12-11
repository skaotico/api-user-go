package user

import (
	serviceUser "api-user-go/internal/service/user"
	"api-user-go/internal/service/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service serviceUser.UserService
}

func NewUserHandler(s serviceUser.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserRequest  true  "Create User Request"
// @Success      200   {object}  map[string]dto.UserCreateResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {

	var userDTO dto.CreateUserRequest

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userGenerated, err := h.service.CreateUser(&userDTO, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"user": userGenerated})
	c.Set("user", userGenerated)
}
