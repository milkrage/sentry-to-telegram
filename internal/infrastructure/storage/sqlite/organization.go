package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/milkrage/sentry-to-telegram/internal/entity"
)

type Organization struct {
	db *sql.DB
}

func NewOrganization(db *sql.DB) *Organization {
	return &Organization{
		db: db,
	}
}

func (o *Organization) Create(slug, installationID, refreshToken string) (*entity.Organization, error) {
	query := "INSERT INTO organizations (slug, installation_id, refresh_token) VALUES (?, ?, ?)"

	_, err := o.db.Exec(query, slug, installationID, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	organization := &entity.Organization{
		Slug:           slug,
		InstallationID: installationID,
		RefreshToken:   refreshToken,
	}

	return organization, nil
}
