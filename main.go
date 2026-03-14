package main

import (
    "log"

    "github.com/gin-gonic/gin"

    "birthday-bot/internal/controller"
    "birthday-bot/internal/infrastructure/database"
    "birthday-bot/internal/infrastructure/slack"
    "birthday-bot/internal/repository"
    "birthday-bot/internal/scheduler"
    "birthday-bot/internal/service"
)

func main() {
    // 1. DB 연결
    db, err := database.NewDatabase(
        "host.docker.internal",
        "3306",
        "root",
        "root",
        "birthday_bot",
    )
    if err != nil {
        log.Fatalf("DB 연결 실패: %v", err)
    }
    log.Println("✅ DB 연결 성공")

    // migration
    if err := db.AutoMigrate(); err != nil {
        log.Fatalf("migration 실패: %v", err)
    }
    log.Println("✅ migration 완료")

    // Repository
    configRepo := repository.NewConfigRepository(db.GetDB())
    slackUserRepo := repository.NewSlackUserRepository(db.GetDB())

    // Service
    configService := service.NewConfigService(configRepo)
    slackUserService := service.NewSlackUserService(slackUserRepo)

    // Slack Client (DB에서 토큰 조회)
    slackToken, err := configService.GetSlackBotToken()
    if err != nil {
        log.Println("⚠️ Slack 토큰이 없습니다. /api/configs로 등록해주세요")
        slackToken = ""
    }
    slackClient := slack.NewClient(slackToken)

    // Scheduler
    birthdayScheduler := scheduler.NewBirthdayScheduler(
        slackUserService,
        configService,
        slackClient,
    )
    if err := birthdayScheduler.Start(); err != nil {
        log.Fatalf("스케줄러 시작 실패: %v", err)
    }
    defer birthdayScheduler.Stop()

    // Gin 라우터
    r := gin.Default()
    api := r.Group("/api")

    // Controller 등록
    healthController := controller.NewHealthController()
    configController := controller.NewConfigController(configService)
    slackUserController := controller.NewSlackUserController(slackUserService)

    healthController.RegisterRoutes(api)
    configController.RegisterRoutes(api)
    slackUserController.RegisterRoutes(api)

    // 서버 시작
    log.Println("🚀 서버 시작: http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("서버 시작 실패: %v", err)
    }
}