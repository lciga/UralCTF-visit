package repository

// Метод для логирования отправленных писем в базу данных.
func (r *MailRepository) LogMail(teamID int, email, subject, status, sendErr string) error {
	query := `INSERT INTO emails_log (team_id, email, subject, status, error)
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, teamID, email, subject, status, sendErr)
	if err != nil {
		return err
	}
	return nil
}
