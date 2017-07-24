package xzmj

import (
	"exportor/defines"
	"fmt"
)

type xzlib struct {

}

func newlib() *xzlib {
	return &xzlib{

	}
}

func (lib *xzlib) OnInit(manager defines.IGameManager, gamedata interface{}) error {
	fmt.Println("xzlib.oninit")
	return nil
}

func (lib *xzlib) OnRelease() {
	fmt.Println("xzlib.onrelease")
}

func (lib *xzlib) OnGameCreate(info *defines.PlayerInfo, conf *defines.CreateRoomConf) error {
	fmt.Println("xzlib.ongamecreate")
	return nil
}

func (lib *xzlib) OnUserEnter(info *defines.PlayerInfo) error {
	fmt.Println("xzlib.onuserenter")
	return nil
}

func (lib *xzlib) OnUserLeave(info *defines.PlayerInfo) {
	fmt.Println("xzlib.onuserleave")
}

func (lib *xzlib) OnUserOffline(info *defines.PlayerInfo) {
	fmt.Println("xzlib.onuseroffline")
}

func (lib *xzlib) OnUserMessage(info *defines.PlayerInfo, cmd uint32, data []byte) error {
	fmt.Println("xzlib.onusermessage")
	return nil
}

func (lib *xzlib) OnTimer(id uint32, data interface{}) {

}
