package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/book/entity"
	"test/book/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, "success")
}

// GetUserBooks 获取用户书籍
// @Summary 获取用户关联书籍
// @Router /users/{id}/books [get]
func (h *UserHandler) GetUserBooks(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	books, err := h.userService.GetUserBooks(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取书籍失败"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// UpdateUserStatus 更新用户状态
// @Summary 更新用户状态
// @Router /users/{id}/status [put]
func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var status struct {
		Status int32 `json:"status"`
	}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.UpdateUserStatus(c.Request.Context(), userID, status.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新状态失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "状态更新成功"})
}
