package repository

// Поиск ID города по имени.
// GetUniversityByCity возвращает список университетов по имени города или региона.
func (r *UniversityRepository) GetUniversityByCity(city string) ([]string, error) {
	var universities []string
	// Join city and region to filter by city or region name
	query := `
       SELECT u.name
       FROM universities u
       JOIN city c ON u.city_id = c.id
       JOIN region_city rc ON c.id = rc.city_id
       JOIN region r ON rc.region_id = r.id
       WHERE c.name ILIKE $1
          OR r.name ILIKE $1
       ORDER BY u.name
   `
	if err := r.db.Select(&universities, query, "%"+city+"%"); err != nil {
		return nil, err
	}
	return universities, nil
}
