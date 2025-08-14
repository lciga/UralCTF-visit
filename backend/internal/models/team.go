package models

type Team struct {
	ID             int           `db:"id" json:"id"`
	Name           string        `db:"name" json:"name"`
	City           int           `db:"city" json:"city"`
	CityName       string        `db:"city_name" json:"city_name"`
	UniversityID   int           `db:"university_id" json:"university_id"`
	UniversityName string        `db:"university_name" json:"university_name"`
	CreatedAt      string        `db:"created_at" json:"created_at"`
	Participants   []Participant `db:"participants" json:"participants"`
}
