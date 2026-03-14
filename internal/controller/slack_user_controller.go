package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"birthday-bot/internal/dto"
	"birthday-bot/internal/service"
)

type SlackUserController struct {
    slackUserService *service.SlackUserService
}

func NewSlackUserController(slackUserService *service.SlackUserService) *SlackUserController {
    return &SlackUserController{slackUserService: slackUserService}
}

func (c *SlackUserController) RegisterRoutes(r *gin.RouterGroup) {
    users := r.Group("/slack-users")
    {
        users.GET("", c.GetAll)
        users.GET("/:id", c.GetByID)
        users.GET("/today-birthdays", c.GetTodayBirthdays)
        users.POST("", c.Create)
        users.POST("/batch", c.CreateBatch)
        users.PUT("/:id", c.Update)
        users.DELETE("/:id", c.Delete)
    }
}

func (c *SlackUserController) GetAll(ctx *gin.Context) {
    users, err := c.slackUserService.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "조회 실패"})
        return
    }

    ctx.JSON(http.StatusOK, users)
}

func (c *SlackUserController) GetByID(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID"})
        return
    }

    user, err := c.slackUserService.GetByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "사용자를 찾을 수 없습니다"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *SlackUserController) GetTodayBirthdays(ctx *gin.Context) {
    users, err := c.slackUserService.GetTodayBirthdays()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "조회 실패"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "date":  time.Now().Format("2006-01-02"),
        "count": len(users),
        "users": users,
    })
}

func (c *SlackUserController) Create(ctx *gin.Context) {
    var req dto.CreateSlackUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.slackUserService.Create(req.Name, req.Email, req.SlackUserID, req.Birthday)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "생성 실패"})
        return
    }

    ctx.JSON(http.StatusCreated, user)
}

func (c *SlackUserController) CreateBatch(ctx *gin.Context) {
    var requests []dto.CreateSlackUserRequest
    if err := ctx.ShouldBindJSON(&requests); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := c.slackUserService.CreateAll(requests)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "일괄 생성 실패"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "생성 완료",
        "count":   len(requests),
    })
}

type UpdateSlackUserRequest struct {
    Name     string    `json:"name" binding:"required"`
    Email    string    `json:"email" binding:"required"`
    Birthday time.Time `json:"birthday" binding:"required"`
}

func (c *SlackUserController) Update(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID"})
        return
    }

    var req UpdateSlackUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.slackUserService.Update(uint(id), req.Name, req.Email, req.Birthday)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "수정 실패"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *SlackUserController) Delete(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID"})
        return
    }

    err = c.slackUserService.Delete(uint(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "삭제 실패"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "삭제 완료"})
}