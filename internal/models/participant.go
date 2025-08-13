package models

type Participant struct {
	ID         int    `db:"id" json:"id"`
	TeamID     int    `db:"team_id" json:"team_id"`
	FirstName  string `db:"first_name" json:"first_name"`
	LastName   string `db:"last_name" json:"last_name"`
	MiddleName string `db:"middle_name,omitempty" json:"middle_name,omitempty"`
	Telegram   string `db:"telegram" json:"telegram"`
	Phone      string `db:"phone" json:"phone"`
	Course     int    `db:"course" json:"course"`
	Email      string `db:"email" json:"email"`
	ShirtSize  string `db:"shirt_size" json:"shirt_size"`
	IsCaptain  bool   `db:"is_captain" json:"is_captain"`
	CreatedAt  string `db:"created_at" json:"created_at"`
}
