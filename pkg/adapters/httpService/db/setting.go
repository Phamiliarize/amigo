package db

import (
	"context"
	"fmt"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/jackc/pgx/v5"
)

type GeneralSetting struct {
	Title  string
	Detail dto.GeneralSetting
}

func (d *Database) GetGeneralSetting() (*dto.GeneralSetting, error) {
	var result GeneralSetting

	err := d.pool.QueryRow(context.Background(), "select title, detail from setting where title = 'general'").Scan(&result.Title, &result.Detail)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no_results")
		}
		return nil, err
	}

	return &result.Detail, nil
}
