package xzmj

import (
	"exportor/defines"
	"gamedef"
)

func GetLib() defines.GameModule {
	return defines.GameModule{
		Type: gamedef.GameLibXz,
		Creator: createXzLib,
		Releaser: releaseXzLib,
		GameData: nil,
	}
}

func createXzLib() defines.IGame {
	return newlib()
}

func releaseXzLib(game defines.IGame) {
	game.OnRelease()
}
