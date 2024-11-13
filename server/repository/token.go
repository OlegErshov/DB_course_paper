package repository

import (
	"DB_course_paper/server/entity"
	"context"
	"database/sql"
	"time"
)

func (r repository) CreateTokens(ctx context.Context, token entity.Token) (int, error) {
	query := `
        INSERT INTO tokens (user_id, access_token, refresh_token, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		token.UserID,
		token.AccessToken,
		token.RefreshToken,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r repository) DeleteTokens(ctx context.Context, tokenId int) error {
	query := `
        DELETE FROM tokens 
        WHERE id = $1
    `

	result, err := r.db.ExecContext(ctx, query, tokenId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Ошибка, если токен с указанным ID не найден
	}

	return nil
}

func (r repository) UpdateTokens(ctx context.Context, token entity.Token) (int, error) {
	query := `
        UPDATE tokens 
        SET access_token = $1, refresh_token = $2, updated_at = $3
        WHERE id = $4
        RETURNING id
    `

	var updatedID int
	err := r.db.QueryRowContext(ctx, query,
		token.AccessToken,
		token.RefreshToken,
		time.Now(), // обновляем поле updated_at на текущее время
		token.ID,   // ID токена, который обновляем
	).Scan(&updatedID)

	if err != nil {
		return 0, err
	}

	return updatedID, nil
}

func (r repository) GetTokens(ctx context.Context, tokenId int) (entity.Token, error) {
	query := `
        SELECT  access_token, refresh_token
        FROM tokens
        WHERE id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, tokenId)
	if err != nil {
		return entity.Token{}, err
	}
	defer rows.Close()

	var tokens entity.Token
	for rows.Next() {
		var token entity.Token
		err := rows.Scan(
			&token.AccessToken,
			&token.RefreshToken,
		)
		if err != nil {
			return entity.Token{}, err
		}
		tokens.AccessToken = token.AccessToken
		tokens.RefreshToken = token.RefreshToken
	}

	if err = rows.Err(); err != nil {
		return entity.Token{}, err
	}

	return tokens, nil
}
