package services

import (
	"BananaStream.API/db/models"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InitializeAdmin(db *gorm.DB) {
	var adminRole models.Role
	if err := db.FirstOrCreate(&adminRole, models.Role{Name: "admin"}).Error; err != nil {
		log.Infof("Admin role already exist: %s", err.Error())
	}

	var adminUser models.User
	if err := db.First(&adminUser, models.User{RoleID: adminRole.ID}).Error; err != nil {
		log.Infof("Admin user does not exist, start creation")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to generate password: %s", err.Error())
		}

		adminUser = models.User{
			Login:    "admin",
			Password: string(hashedPassword),
			RoleID:   adminRole.ID,
		}

		if err := db.Create(&adminUser).Error; err != nil {
			log.Warnw("Admin user creation failed, error: %s", err.Error())
		}

		return
	}
	log.Warnw("Admin user already exist")
}
