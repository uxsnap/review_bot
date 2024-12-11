package migrator

import (
	db "github.com/uxsnap/review_bot/internal/client/database"
	"github.com/uxsnap/review_bot/internal/entity"
)

func Migrate(dbClient db.DbClient) error {
	var migrationEntity = []interface{}{
		&entity.Category{},
		&entity.User{},
		&entity.ProblematicQuestion{},
		&entity.Question{},
		&entity.Statistic{},
		&entity.Test{},
		&entity.UsefulLink{},
	}

	err := dbClient.DB().AutoMigrate(migrationEntity...)

	if err != nil {
		return err
	}

	return nil
}
