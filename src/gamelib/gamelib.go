//go:binary-only-package

package gamelib

import (
	"exportor/defines"
)

func StartGame(modules []defines.GameModule) {
	starter.StartProgram("game", modules)
}

func Marshal(data interface{}) ([]byte, error) {
	return msgpacker.Marshal(data)
}

func UnMarshal(data []byte, p interface{}) error {
	return msgpacker.UnMarshal(data, p)
}