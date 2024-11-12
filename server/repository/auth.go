package repository

import "context"

func (r repository) LogOutUser(ctx context.Context, userId int) error {
	query := `UPDATE tokens
              SET access_token = NULL, refresh_token = NULL
              WHERE user_id = $1`

	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
