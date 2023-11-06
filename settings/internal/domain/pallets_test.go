package domain_test

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/meetmorrowsolonmars/openpgl/settings/internal/domain"
)

func TestColorSettings_Bson(t *testing.T) {
	settings := domain.ColorSettings{
		UserID: 10,
		Pallets: []domain.Pallet{
			{
				ID: 1,
				Colors: []domain.AltitudeColor{
					{
						Altitude: 1.0,
						Color: color.RGBA{
							R: 255,
							G: 0,
							B: 0,
							A: 255,
						},
					},
					{
						Altitude: 0.5,
						Color: color.RGBA{
							R: 0,
							G: 255,
							B: 0,
							A: 255,
						},
					},
				},
			},
			{
				ID: 2,
				Colors: []domain.AltitudeColor{
					{
						Altitude: 0.2,
						Color: color.RGBA{
							R: 255,
							G: 255,
							B: 0,
							A: 255,
						},
					},
				},
			},
		},
	}

	payload, err := bson.Marshal(settings)
	require.NoError(t, err)

	var loadedSettings domain.ColorSettings
	err = bson.Unmarshal(payload, &loadedSettings)
	require.NoError(t, err)

	require.Equal(t, settings, loadedSettings)
}
