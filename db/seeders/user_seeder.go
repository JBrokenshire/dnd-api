package seeders

import "dnd-api/db/models"

func (s *Seeder) SetUsers() {
	users := []models.User{
		{
			ID:       1,
			Email:    "matt@purplevisits.com",
			Name:     "Matt Nelson",
			Password: "$2a$10$1p7pUkeUIemFrJwgMX3.t.b4YZKk6ZvWlqBOug52ksrZ.ybZKtY.S", // Abcd1234$
		},
	}
	for _, user := range users {
		result := s.DB.Where("id = ?", user.ID).FirstOrCreate(&user)
		if result.RowsAffected == 1 {
			// Link to admin role
			s.DB.Model(user).Association("Roles").Append([]models.Role{{ID: 1}})
		}
	}
}
