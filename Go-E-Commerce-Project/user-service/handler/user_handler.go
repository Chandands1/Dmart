package handler


import(
	"net/http"

	"github.com/gin-gonic/gin"

)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (h *UserHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register user"})
}


