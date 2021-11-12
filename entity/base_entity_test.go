package entity_test

import (
	"testing"

	"github.com/paulotokimatu/flat_game/entity"

	"github.com/stretchr/testify/assert"
)

func TestBaseEntCanAlwaysTick(t *testing.T) {
	entity := entity.NewBaseEntity(&entity.Config{Name: "foo"})

	assert.True(t, entity.CanTick(nil), "canTick should not return false")
}
