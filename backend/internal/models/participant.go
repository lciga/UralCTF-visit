package models

// swagger:model Participant
// Participant представляет участника команды
type Participant struct {
	ID         int    `db:"id" json:"id"`
	TeamID     int    `db:"team_id" json:"team_id"`
	FirstName  string `db:"first_name" json:"first_name" binding:"required,min=2"`
	LastName   string `db:"last_name" json:"last_name" binding:"required,min=2"`
	MiddleName string `db:"middle_name,omitempty" json:"middle_name,omitempty"`
	Telegram   string `db:"telegram" json:"telegram" binding:"required"`
	Phone      string `db:"phone" json:"phone" binding:"required,e164"`
	Course     int    `db:"course" json:"course" binding:"required,min=1,max=6"`
	Email      string `db:"email" json:"email" binding:"required,email"`
	ShirtSize  string `db:"shirt_size" json:"shirt_size" binding:"required,oneof=XS S M L XL XXL"`
	IsCaptain  bool   `db:"is_captain" json:"is_captain"`
	CreatedAt  string `db:"created_at" json:"created_at"`
}
