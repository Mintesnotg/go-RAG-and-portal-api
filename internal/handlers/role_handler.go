package handlers

import (
	"net/http"

	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService services.RoleService
}

type AssignRoleRequest struct {
	UserID string `json:"user_id" binding:"required"`
	RoleID string `json:"role_id" binding:"required"`
}

func NewRoleHandler(roleService services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// AssignRole godoc
// @Summary Assign a role to a user
// @Tags Roles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body AssignRoleRequest true "assign role"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/roles/assign [post]
func (h *RoleHandler) AssignRole(c *gin.Context) {
	var req AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	roles, err := h.roleService.AssignRoleToUser(req.UserID, req.RoleID)
	if err != nil {
		switch err {
		case services.ErrRoleAssignUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		case services.ErrRoleAssignRoleNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not assign role"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}
