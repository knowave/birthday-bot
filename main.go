package main

import (
	"birthday-bot/internal/infrastructure/database"
	"fmt"
	"log"
)

func main() {
	db, err := database.NewDatabase(
		"localhost",  // host
        "3306",       // port
        "root",       // user
        "root",   // password (본인 MySQL 비밀번호로 변경)
        "birthday_bot", // dbname
	)

	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	 fmt.Println("✅ DB 연결 성공!")

    if err := db.AutoMigrate(); err != nil {
        log.Fatalf("마이그레이션 실패: %v", err)
    }

    fmt.Println("✅ 테이블 생성 완료!")
}