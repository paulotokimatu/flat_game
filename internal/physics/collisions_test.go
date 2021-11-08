package physics_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/internal/physics"
	"flat_game/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockEntity struct {
	entity.BaseEntity
	CollisionHappened bool
}

func (entity *MockEntity) OnCollision(game flat_game.IGame, externalEntity flat_game.IEntity) {
	entity.CollisionHappened = true
}

func TestShouldExecuteCollisionChecks(t *testing.T) {
	entity1 := &MockEntity{
		BaseEntity: *entity.NewBaseEntity(&entity.Config{
			Position: utils.Vec2{
				X: 10,
				Y: 10,
			},
			Size: utils.Vec2{
				X: 10,
				Y: 10,
			},
		}),
	}
	entity2 := &MockEntity{
		BaseEntity: *entity.NewBaseEntity(&entity.Config{
			Position: utils.Vec2{
				X: 30,
				Y: 30,
			},
			Size: utils.Vec2{
				X: 10,
				Y: 10,
			},
		}),
	}
	entity3 := &MockEntity{
		BaseEntity: *entity.NewBaseEntity(&entity.Config{
			Position: utils.Vec2{
				X: 10,
				Y: 10,
			},
			Size: utils.Vec2{
				X: 10,
				Y: 10,
			},
		}),
	}
	entity4 := &MockEntity{
		BaseEntity: *entity.NewBaseEntity(&entity.Config{
			Position: utils.Vec2{
				X: 15,
				Y: 15,
			},
			Size: utils.Vec2{
				X: 10,
				Y: 10,
			},
		}),
	}

	var collisions [][2]flat_game.IEntity

	collisions = append(collisions, [2]flat_game.IEntity{entity1, entity2})
	collisions = append(collisions, [2]flat_game.IEntity{entity3, entity4})

	physics.ExecuteCollisions(nil, collisions)

	assert.False(t, entity1.CollisionHappened, "collision should not have happened")
	assert.False(t, entity2.CollisionHappened, "collision should not have happened")
	assert.True(t, entity3.CollisionHappened, "collision should have happened")
	assert.True(t, entity4.CollisionHappened, "collision should have happened")
}
