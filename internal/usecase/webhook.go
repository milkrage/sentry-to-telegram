package usecase

import (
	"fmt"

	"github.com/milkrage/sentry-to-telegram/internal/entity"
)

type OrganizationRepository interface {
	Create(slug, installationID, refreshToken string) (*entity.Organization, error)
}

type SentryClient interface {
	Authorization(code, clientID, clientSecret string) (string, error)
	Confirm(installationID string) error
}

type Webhook struct {
	organizationRepository OrganizationRepository
	sentryClient           SentryClient
	clientID               string
	clientSecret           string
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (w *Webhook) Setup(code, installationID, slug string) (*entity.Organization, error) {
	refreshToken, err := w.sentryClient.Authorization(code, w.clientID, w.clientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize in sentry api: %w", err)
	}

	organization, err := w.organizationRepository.Create(slug, installationID, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to save information about organization to the database: %w", err)
	}

	err = w.sentryClient.Confirm(installationID)
	if err != nil {
		return nil, fmt.Errorf("failed to confirm installation in sentry api: %w", err)
	}

	return organization, nil
}
