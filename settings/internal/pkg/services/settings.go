package services

import (
	"context"

	"github.com/meetmorrowsolonmars/openpgl/settings/internal/domain"
)

type SettingsRepository interface {
	UpsertDefaultSettings(ctx context.Context, settings *domain.Pallet) error
}

type SettingsService struct {
	repository SettingsRepository
}

func NewSettingsService(repository SettingsRepository) *SettingsService {
	return &SettingsService{
		repository: repository,
	}
}

func (s *SettingsService) UpsertDefaultSettings(ctx context.Context, settings *domain.Pallet) error {
	// TODO: add validation
	settings.ID = 1
	return s.repository.UpsertDefaultSettings(ctx, settings)
}
