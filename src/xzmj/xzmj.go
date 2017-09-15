package xzmj

import (
	"exportor/defines"
	"def"
	"fmt"
)

func GetLib() defines.GameModule {
	return defines.GameModule{
		Type:     def.GameLibXz,
		Creator:  createXzLib,
		Releaser: releaseXzLib,
		GameData: nil,
		PlayerCount: 4,
	}
}

func createXzLib() defines.IGame {
	fmt.Println("create xz lib")
	return newlib()
}

func releaseXzLib(game defines.IGame) {
	fmt.Println("release xz lib")
	game.OnRelease()
}

