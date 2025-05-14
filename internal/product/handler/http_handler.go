package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gomicro/internal/product/service"
)

// ProductHTTPHandler handles HTTP requests for products
type ProductHTTPHandler struct {
	service service.ProductService
}

// NewProductHTTPHandler creates a new HTTP handler for products
func NewProductHTTPHandler(service service.ProductService) *ProductHTTPHandler {
	return &ProductHTTPHandler{
		service: service,
	}
}

// RegisterRoutes registers the HTTP routes for products
func (h *ProductHTTPHandler) RegisterRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/:id", h.GetProduct)
		products.POST("/", h.CreateProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
		products.GET("/", h.ListProducts)
	}
}

// GetProduct handles GET /products/:id
func (h *ProductHTTPHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetProduct(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST /products
func (h *ProductHTTPHandler) CreateProduct(c *gin.Context) {
	var product struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Stock       int     `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.service.CreateProduct(c.Request.Context(), product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct handles PUT /products/:id
func (h *ProductHTTPHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.service.UpdateProduct(c.Request.Context(), uint(id), product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles DELETE /products/:id
func (h *ProductHTTPHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.DeleteProduct(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListProducts handles GET /products
func (h *ProductHTTPHandler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
} 