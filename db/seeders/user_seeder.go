package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetUsers() {
	users := []models.User{{
		ID:       1,
		Username: "JBrokenshire",
		Password: "$2a$10$1p7pUkeUIemFrJwgMX3.t.b4YZKk6ZvWlqBOug52ksrZ.ybZKtY.S", // Abcd1234$"
		Admin:    true,
	}}

	for _, user := range users {
		err := s.DB.Where("id = ?", user.ID).FirstOrCreate(&user).Error
		if err != nil {
			log.Printf("Error creating user with id %v in seeder", user.ID)
		}
	}
}
