package sentry

type installation struct {
}

func (i *installation) Authorization(code, clientID, clientSecret string) (string, string, error) {}

func (i *installation) Confirm(installationID string) error {}
