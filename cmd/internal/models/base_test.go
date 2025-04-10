package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBaseModel(t *testing.T) {
	now := time.Now()

	base := Base{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: gorm.DeletedAt{},
	}

	assert.Equal(t, uint(1), base.ID, "ID should be set to 1")
	assert.WithinDuration(t, now, base.CreatedAt, time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, now, base.UpdatedAt, time.Second, "UpdatedAt should be close to now")
	assert.False(t, base.DeletedAt.Valid, "DeletedAt should be initially invalid")
}

func TestBaseModel_DefaultValues(t *testing.T) {
	base := Base{}

	assert.Equal(t, uint(0), base.ID, "Default ID should be 0")
	assert.True(t, base.CreatedAt.IsZero(), "Default CreatedAt should be zero initially")
	assert.True(t, base.UpdatedAt.IsZero(), "Default UpdatedAt should be zero initially")
	assert.False(t, base.DeletedAt.Valid, "Default DeletedAt should be invalid")
}

func TestBaseModel_SetValues(t *testing.T) {
	createdAt := time.Now().Add(-time.Hour)
	updatedAt := time.Now()
	deletedAt := gorm.DeletedAt{Time: time.Now(), Valid: true}

	base := Base{
		ID:        10,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}

	assert.Equal(t, uint(10), base.ID, "ID should be set correctly")
	assert.WithinDuration(t, createdAt, base.CreatedAt, time.Second, "CreatedAt should be set correctly")
	assert.WithinDuration(t, updatedAt, base.UpdatedAt, time.Second, "UpdatedAt should be set correctly")
	assert.True(t, base.DeletedAt.Valid, "DeletedAt should be valid")
	assert.WithinDuration(t, deletedAt.Time, base.DeletedAt.Time, time.Second, "DeletedAt Time should be set correctly")
}
