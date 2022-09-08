package controller

import (
	"context"

	"github.com/token-refresher/ent"
	"github.com/token-refresher/ent/apikey"
)

type ApiKeyController struct {
	client *ent.Client
}

func (controller ApiKeyController) GetFreshAPIKey(ctx context.Context) (*ent.ApiKey, error) {
	current, err := controller.client.ApiKey.
		Query().Where(apikey.StatusIn(apikey.StatusInUse)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return current, nil
}
