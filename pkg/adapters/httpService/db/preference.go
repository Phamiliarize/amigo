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

	err := d.pool.QueryRow(context.Background(), "select theme from preference where user_id = $1", userID).Scan(&result.Theme)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no_results")
		}
		return nil, err
	}

	return &result, nil
}
