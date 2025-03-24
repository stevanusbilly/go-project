package handlers

import (
	"go-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleHandler struct {
	service *services.RoleService
}

func NewRoleHandler(service *services.RoleService) *RoleHandler {
	return &RoleHandler{service}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var request struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedBy   uuid.UUID `json:"created_by"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.CreateRole(request.Name, request.Description, request.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.service.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UpdatedBy   string `json:"updated_by"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedByUUID, err := uuid.Parse(request.UpdatedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID for updated_by"})
		return
	}

	err = h.service.UpdateRole(id, request.Name, request.Description, updatedByUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully"})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		DeletedBy string `json:"deleted_by"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	deletedByUUID, err := uuid.Parse(request.DeletedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID for deleted_by"})
		return
	}

	err = h.service.DeleteRole(id, deletedByUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}