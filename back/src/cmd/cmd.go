package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"ramada/api/src/auth"
	"ramada/api/src/db"
	"ramada/api/src/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

func Flow() {
	var (
		joined = strings.Join(os.Args, " ")
	)

	if strings.Contains(joined, "get-secret") {
		fmt.Println(auth.GetSecret())

		return
	}

	if strings.Contains(joined, "migrate") {
		Migrate()
	}

	if strings.Contains(joined, "seed") {
		Seed()
	}

}

func Migrate() {
	var db gorm.DB = *db.GetDB()

	db.Migrator().DropTable(models.User{})
	db.AutoMigrate(models.User{})

	db.Migrator().DropTable(models.Product{})
	db.AutoMigrate(models.Product{})
}

func seedUsers(count int) {
	var (
		db   gorm.DB = *db.GetDB()
		user models.User
	)

	password, _error := auth.HashIt("Password@123")

	if _error != nil {
		log.Fatal("seeding users error")
		return
	}

	for i := 1; i < count; i++ {
		user = models.User{
			ID:        uint64(i),
			Name:      fmt.Sprintf("User %v", i),
			Username:  fmt.Sprintf("user%v", i),
			Email:     fmt.Sprintf("user%v@example.com", i),
			Password:  string(password),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		db.Create(&user)
	}
}

func seedProducts(count int) {
	var (
		db      gorm.DB = *db.GetDB()
		product models.Product
	)

	var categories = []string{
		"Category 1",
		"Category 2",
		"Category 3",
		"Category 4",
	}

	for i := 1; i < count; i++ {
		product = models.Product{
			ID:          uint64(i),
			Name:        fmt.Sprintf("Product %v", i),
			Description: fmt.Sprintf("Description %v", i),
			Price:       fmt.Sprintf("%v", rand.Float64()*1000),
			Category:    categories[rand.Intn(len(categories))],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		db.Create(&product)
	}
}

func Seed() {
	seedProducts(20)
	seedUsers(10)
}
