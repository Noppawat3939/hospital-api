package migration

import (
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

func MigrationCreateStaff(db *gorm.DB) error {
	if err := db.Migrator().AutoMigrate(&model.Staff{}); err != nil {
		return err
	}

	indexName := "idx_staff_hospital_username"

	if !db.Migrator().HasIndex(&model.Staff{}, indexName) {
		if err := db.Migrator().CreateIndex(&model.Staff{}, "hospital_id_username"); err != nil {
			return err
		}
	}
	return nil
}
