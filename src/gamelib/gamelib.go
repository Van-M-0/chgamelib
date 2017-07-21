package gamelib

import (
	"starter"
	"exportor/defines"
)

func StartGame(modules []defines.GameModule) {
	starter.StartProgram("game", modules)
}