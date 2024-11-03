package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `gorm:"type:char(36);primary_key;not null;unique" json:"id"`
	FirstName string    `gorm:"type:varchar(155);not null" json:"first_name" validate:"required"`
	LastName  string    `gorm:"type:varchar(155);not null" json:"last_name" validate:"required"`
	FullName  string    `gorm:"type:varchar(310);not null" json:"full_name"`
	Email     string    `gorm:"type:varchar(155);not null" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:varchar(155);not null" json:"password"`
	Active    bool      `gorm:"not null;default:true" json:"active"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		ID        string    `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  u.FullName,
		Email:     u.Email,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})
}
