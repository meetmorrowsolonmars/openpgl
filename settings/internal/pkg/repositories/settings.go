package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/meetmorrowsolonmars/openpgl/settings/internal/domain"
)

type SettingsRepository struct {
	client *mongo.Client

	dbName                        string
	settingsCollectionName        string
	defaultSettingsCollectionName string
}

func NewSettingsRepository(client *mongo.Client) *SettingsRepository {
	return &SettingsRepository{
		client:                        client,
		dbName:                        "settings_db",
		settingsCollectionName:        "pallets_settings",
		defaultSettingsCollectionName: "default_pallets_settings",
	}
}

func (s *SettingsRepository) UpsertDefaultSettings(ctx context.Context, settings *domain.Pallet) error {
	db := s.client.Database(s.dbName)
	collection := db.Collection(s.defaultSettingsCollectionName)

	filter := bson.M{"id": settings.ID}
	_, err := collection.ReplaceOne(ctx, filter, settings, options.Replace().SetUpsert(true))

	return err
}
