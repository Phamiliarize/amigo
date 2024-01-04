package db

import (
	"context"
	"fmt"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/jackc/pgx/v5"
)

// GetUserPreference is going to need some sort of caching to avoid smacking the DB on every auth'd user workflow
func (d *Database) GetUserPreference(userID string) (*dto.Preference, error) {
	var result dto.Preference

	err := d.pool.QueryRow(context.Background(), "select theme, reading_mode from preference where user_id = $1", userID).Scan(&result.Theme, &result.ReadingMode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no_results")
		}
		return nil, err
	}

	return &result, nil
}

// UpdateUserPreference updates the users preferences
func (d *Database) UpdateUserPreference(userID string, update dto.Preference) error {
	res, err := d.pool.Exec(
		context.Background(),
		"update preference set theme = COALESCE($1, theme), reading_mode = COALESCE($2, reading_mode) where user_id = $3",
		update.Theme,
		update.ReadingMode,
		userID,
	)
	if res.RowsAffected() == 0 {
		return fmt.Errorf("not_found")
	}
	if err != nil {
		return err
	}

	return nil
}
