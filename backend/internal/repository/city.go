package repository

// Поиск ID города по имени.
func (r *CityRepository) GetCityId(name string) (int, error) {
	var cityId int
	query := "SELECT id FROM city WHERE name ILIKE $1 LIMIT 1"
	// Используем переданное имя для шаблона поиска
	if err := r.db.Get(&cityId, query, "%"+name+"%"); err != nil {
		return 0, err // Возвращаем ошибку, если город не найден
	}
	return cityId, nil // Возвращаем идентификатор города
}

// SearchCities возвращает список названий городов, соответствующих запросу.
func (r *CityRepository) SearchCities(query string) ([]string, error) {
	var cities []string
	sql := `SELECT name FROM city WHERE name ILIKE $1 ORDER BY name LIMIT 50` // фильтрация по шаблону
	if err := r.db.Select(&cities, sql, "%"+query+"%"); err != nil {
		return nil, err
	}
	return cities, nil
}
