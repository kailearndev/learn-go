package user

import (
	"kai-shop-be/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// Register user-related routes here
	g := r.Group("/users")
	g.POST("/register", h.RegisterUser)
	g.POST("/login", h.LoginUser)
	g.GET("/:id", h.GetUserByID)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
	g.GET("/", h.ListUsers)

}

func (h *Handler) RegisterUser(c *gin.Context) {
	var req RegisterUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.RegisterUser(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "registration failed")
		return
	}

	response.Created(c, user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var req LoginUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.LoginUser(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}
func (h *Handler) GetUserByID(c *gin.Context) {
	idParam := c.GetString("id")
	user, err := h.service.GetUserByID(uuid.MustParse(idParam))
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found")
		return
	}
	response.Success(c, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	idParam := c.GetString("id")
	var req RegisterUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	updatedUser, err := h.service.UpdateUser(uuid.MustParse(idParam), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to update user")
		return
	}
	response.Success(c, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.GetString("id")
	err := h.service.DeleteUser(uuid.MustParse(idParam))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to delete user")
		return
	}
	response.Success(c, gin.H{"message": "user deleted"})
}


func (h *Handler) ListUsers(c *gin.Context) {
	// Implementation of listing users
	limit := 10
	offset := 0
	if l := c.Query("limit"); l != "" {
		// parse limit
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		// parse offset
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}

	items, total, err := h.service.ListUsers(limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}