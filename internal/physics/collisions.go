package physics

import "github.com/paulotokimatu/flat_game"

func doCollisionsOverlap(entityA flat_game.IEntity, entityB flat_game.IEntity) bool {
	positionA := entityA.Position()
	positionB := entityB.Position()
	sizeA := entityA.Size()
	sizeB := entityB.Size()

	// This is just a naive logic considering two rectangle collision shapes
	collisionX := positionA.X+sizeA.X >= positionB.X &&
		positionB.X+sizeB.X >= positionA.X
	collisionY := positionA.Y+sizeA.Y >= positionB.Y &&
		positionB.Y+sizeB.Y >= positionA.Y

	return collisionX && collisionY
}

func ExecuteCollisions(game flat_game.IGame, entities [][2]flat_game.IEntity) {
	for i := 0; i < len(entities); i++ {
		entityA := entities[i][0]
		entityB := entities[i][1]

		if doCollisionsOverlap(entityA, entityB) {
			entityA.OnCollision(game, entityB)
			entityB.OnCollision(game, entityA)
		}
	}
}
