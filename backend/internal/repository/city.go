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

type CitySearchResult struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// SearchCities возвращает список названий городов, соответствующих запросу.
func (r *CityRepository) SearchCities(query string) ([]CitySearchResult, error) {
	var cities []CitySearchResult
	// фильтрация по префиксу, выдаем до 50 результатов
	// возвращаем только города, имя которых начинается с запроса
	prefix := query + "%"
	sql := `SELECT id, name FROM city
			WHERE name ILIKE $1
			ORDER BY name
			LIMIT 50`
	if err := r.db.Select(&cities, sql, prefix); err != nil {
		return nil, err
	}
	return cities, nil
}
