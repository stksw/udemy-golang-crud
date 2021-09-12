package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

// docker-compose exec backend shで接続
// go run src/commands/populateUsers.go で実行
func populateProducts() {
	database.Connect()

	for i := 0; i < 30; i++ {
		products := models.Product{
			Title: faker.Name(),
			Description: faker.Name(),
			Image: faker.URL(),
			Price: float64(rand.Intn(90) + 10),

		}

		
		database.DB.Create(&products)
	}
}