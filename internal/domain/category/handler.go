package category

import (
	"net/http"
	"strconv"

	"kai-shop-be/pkg/response"

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
	// Register product-related routes here
	g := r.Group("/categories")
	{
		g.GET("/", h.ListCategories)
		g.POST("/", h.CreateCategory)
		g.GET("/:id", h.GetCategoryByID)
		g.PUT("/:id", h.UpdateCategory)
		g.DELETE("/:id", h.DeleteCategory)
	}
}

func (h *Handler) ListCategories(c *gin.Context) {
	// Implementation of listing categories
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

	items, total, err := h.service.ListCategories(limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": total,
	})

}

func (h *Handler) CreateCategory(c *gin.Context) {

	var category CategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	categories, err := h.service.CreateCategory(category)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, categories)
}

func (h *Handler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	category, err := h.service.GetCategoryByID(uuid.MustParse(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return

	}
	response.Success(c, category)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category CategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.service.UpdateCategory(uuid.MustParse(id), category)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(uuid.MustParse(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "category deleted"})
}
