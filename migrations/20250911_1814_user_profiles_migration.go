package migrations

import (
	"fmt"
	"squadify-app/models"

	"gorm.io/gorm"
)

func MigrateUserProfiles(db *gorm.DB) error {
	// Step 1: Create the new table
	if err := db.AutoMigrate(&models.UserProfile{}); err != nil {
		return fmt.Errorf("failed to migrate user_profiles: %w", err)
	}

	// Step 2: Move data from old users columns into user_profiles
	// (adjust column names as per your old schema)
	if err := db.Exec(`
        INSERT INTO user_profiles (user_id, first_name, middle_name, last_name, latitude, longitude, address, gender)
        SELECT id, first_name, middle_name, last_name, latitude, longitude, address, gender FROM users
    `).Error; err != nil {
		return fmt.Errorf("failed to migrate old user data to user_profiles: %w", err)
	}

	// Step 3: Drop old columns from users table
	if err := db.Exec(`
        ALTER TABLE users
        DROP COLUMN first_name,
        DROP COLUMN middle_name,
        DROP COLUMN last_name,
        DROP COLUMN latitude,
        DROP COLUMN longitude,
        DROP COLUMN address,
        DROP COLUMN gender
    `).Error; err != nil {
		return fmt.Errorf("failed to drop old columns: %w", err)
	}

	return nil
}
