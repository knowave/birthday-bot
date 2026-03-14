package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "birthday-bot/internal/service"
)

type ConfigController struct {
    configService *service.ConfigService
}

func NewConfigController(configService *service.ConfigService) *ConfigController {
    return &ConfigController{configService: configService}
}

func (c *ConfigController) RegisterRoutes(r *gin.RouterGroup) {
    config := r.Group("/configs")
    {
        config.GET("/:key", c.GetByKey)
        config.POST("", c.Create)
    }
}

func (c *ConfigController) GetByKey(ctx *gin.Context) {
    key := ctx.Param("key")

    value, err := c.configService.GetValue(key)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "설정을 찾을 수 없습니다"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "key":   key,
        "value": value,
    })
}

type CreateConfigRequest struct {
    Key   string `json:"key" binding:"required"`
    Value string `json:"value" binding:"required"`
}

func (c *ConfigController) Create(ctx *gin.Context) {
    var req CreateConfigRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := c.configService.SetValue(req.Key, req.Value)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "설정 저장 실패"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "key":   req.Key,
        "value": req.Value,
    })
}