package main

import (
	"log"
	"net/http"
	"time"
	"znav/backend/config"
	"znav/backend/database"
	"znav/backend/models"
	"znav/backend/routes"
	"znav/backend/utils"
)

func init() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}
	time.Local = loc
}

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库连接
	database.InitDB()
	db := database.GetDB()

	// 确保数据库迁移已经执行
	models.Migrate(db)

	now := time.Now() // 先将当前时间保存到变量中

	// 确保默认的 admin 用户存在
	var user models.User
	if err := db.Where("username = ?", "admin").First(&user).Error; err != nil {
		// 如果找不到用户，则创建默认的 admin 用户
		hashedPassword, err := utils.HashPassword("admin")
		if err != nil {
			log.Fatalf("Failed to hash default admin password: %v", err)
		}
		// 创建默认的 admin 用户
		user = models.User{
			Username: "admin",
			Password: hashedPassword,
			LoginAt:  &now,
			Status:   "enabled", // 默认启用
			IsAdmin:  true,      // 设置为管理员
		}

		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create default admin user: %v", err)
		}
		log.Println("Initialized default admin user with username: admin and password: admin")
	} else {
		// 如果用户已经存在，确保其 IsAdmin 字段为 true
		if !user.IsAdmin {
			user.IsAdmin = true
			if err := db.Save(&user).Error; err != nil {
				log.Fatalf("Failed to update admin user to administrator: %v", err)
			}
			log.Println("Updated existing admin user to have administrator privileges")
		} else {
			log.Println("Admin user already exists and is an administrator")
		}
	}

	// 设置路由
	router := routes.SetupRouter()

	// 启动 HTTP 服务器
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
