package xzmj

import "exportor/defines"

type xzlib struct {

}

func newlib() *xzlib {
	return &xzlib{

	}
}

func (lib *xzlib) OnInit(manager defines.IGameManager, gamedata interface{}) error {
	return nil
}

func (lib *xzlib) OnRelease() {

}

func (lib *xzlib) OnGameCreate(info *defines.PlayerInfo, conf *defines.CreateRoomConf) error {
	return nil
}

func (lib *xzlib) OnUserEnter(info *defines.PlayerInfo) error {
	return nil
}

func (lib *xzlib) OnUserLeave(info *defines.PlayerInfo) {

}

func (lib *xzlib) OnUserOffline(info *defines.PlayerInfo) {

}

func (lib *xzlib) OnUserMessage(info *defines.PlayerInfo, cmd uint32, data []byte) error {
	return nil
}

func (lib *xzlib) OnTimer(id uint32, data interface{}) {

}
