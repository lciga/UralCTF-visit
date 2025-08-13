package models

type Team struct {
	ID           int           `db:"id" json:"id"`
	Name         string        `db:"name" json:"name"`
	City         string        `db:"city" json:"city"`
	University   string        `db:"university" json:"university"`
	CreatedAt    string        `db:"created_at" json:"created_at"`
	Participants []Participant `db:"participants,omitempty" json:"participants,omitempty"`
}
