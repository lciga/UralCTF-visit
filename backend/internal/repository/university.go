package repository

// swagger:model University
// University представляет запись об университете в ответе
type University struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// SearchByCity возвращает список университетов по имени города.
func (r *UniversityRepository) SearchByCity(city string) ([]University, error) {
	var list []University
	query := `SELECT u.id, u.name
			  FROM universities u
			  JOIN city c ON u.city_id = c.id
			  WHERE c.name ILIKE $1
			  ORDER BY u.name
			  LIMIT 50`
	if err := r.db.Select(&list, query, "%"+city+"%"); err != nil {
		return nil, err
	}
	return list, nil
}
