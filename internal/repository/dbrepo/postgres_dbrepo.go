package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllLinks() ([]*models.Link, error) {
	/**
	* Этот код создаёт новый контекст с временным ограничением (timeout) на основе базового
		(background) контекста. Здесь dbTimeout определяет, как долго этот контекст остается активным перед автоматической отменой.
	* Функция context.WithTimeout возвращает два значения: сам контекст (ctx) и функцию
		отмены (cancel). Этот контекст обычно передаётся в функции, которые должны быть ограничены по времени. Например, это может быть запрос к базе данных или HTTP-запрос.
	* Функция cancel, возвращаемая context.WithTimeout, используется для освобождения ресурсов,
		которые могут быть выделены контекстом. В данном случае defer cancel() гарантирует, что cancel будет вызван независимо от того, как завершится функция, в которой этот код находится. Если вы не вызовете cancel, то ресурсы, связанные с контекстом, не будут освобождены, что может привести к утечкам ресурсов.
	* Итак, использование defer cancel() является хорошей практикой для управления жизненным
		циклом контекста. Оно гарантирует, что любые ресурсы, связанные с ctx, будут освобождены в конце текущей функции.
	*/
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, title, phone, email,
			links, description, coalesce(image, ''),
			created_at, updated_at
		from
		    links
		order by 
		    title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*models.Link

	for rows.Next() {
		var link models.Link
		err := rows.Scan(
			&link.ID,
			&link.Title,
			&link.Phone,
			&link.Email,
			&link.Links,
			&link.Description,
			&link.Image,
			&link.CreatedAt,
			&link.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, &link)
	}

	return links, nil
}
