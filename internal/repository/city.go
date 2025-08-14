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
