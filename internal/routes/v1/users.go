package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", h.GetUsers)
	}
}

func (h *Handler) GetUsers(c *gin.Context) {

}
