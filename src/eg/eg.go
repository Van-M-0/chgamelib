package eg

import (
	"exportor/defines"
	"def"
)

func GetLib() defines.GameModule {
	return defines.GameModule{
		Type:     def.GameLibXz,
		Creator:  createEgLib,
		Releaser: releaseEgLib,
		GameData: nil,
	}
}

func createEgLib() defines.IGame {
	return newlib()
}

func releaseEgLib(game defines.IGame) {
	game.OnRelease()
}