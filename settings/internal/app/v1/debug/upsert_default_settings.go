package debug

import (
	"context"
	"image/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/openpgl/settings/internal/domain"
	desc "github.com/meetmorrowsolonmars/openpgl/settings/internal/pb/api/v1/debug"
)

func (i *Implementation) UpsertDefaultSettings(ctx context.Context, request *desc.UpsertDefaultSettingsRequest) (*desc.UpsertDefaultSettingsResponse, error) {
	// TODO: validate pallet != nil
	// TODO: validate pallet id
	settings := &domain.Pallet{
		ID:     request.Pallet.Id,
		Colors: make([]domain.AltitudeColor, 0, len(request.Pallet.Colors)),
	}

	for _, c := range request.Pallet.Colors {
		settings.Colors = append(settings.Colors, domain.AltitudeColor{
			Altitude: c.Altitude,
			Color: color.RGBA{
				R: uint8(c.R),
				G: uint8(c.G),
				B: uint8(c.B),
				A: uint8(c.A),
			},
		})
	}

	err := i.settingsService.UpsertDefaultSettings(ctx, settings)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't upsert default settings: %s", err)
	}

	return &desc.UpsertDefaultSettingsResponse{}, nil
}
