package xzmj

import (
	"exportor/defines"
	"fmt"
	"def"
	"gamelib"
	"time"
	"math/rand"
//	"strconv"
//	"database/sql/driver"
	"errors"
	"exportor/proto"
)




type xzlib struct {
	mgr 		defines.IGameManager
	conf 		*XzConf
}

type XzConf struct {
	BottomScore 	int
	Fanshu			[]int
}

func getXzConf() []byte {
	conf := &XzConf {
		BottomScore: 100,
		Fanshu: []int{1, 2, 4},
	}
	data, err := gamelib.Marshal(conf)
	if err == nil {
		return data
	} else {
		panic("marshal xz conf err " )
	}
}


func newlib() *xzlib {
	return &xzlib{

	}
}

func (lib *xzlib) OnInit(manager defines.IGameManager, gamedata interface{}) error {
	fmt.Println("xzlib.oninit")
	lib.mgr = manager
	//lib.mgr.SetTimer()
	return nil
}

func (lib *xzlib) OnRelease() {
	var RoomID int
	var AddRoomCard int
	RoomID=def.GetRoomid(lib.mgr.GetRoomId())

	AddRoomCard=0
	fmt.Println("xzlib.onrelease,解散房间",lib.mgr.GetRoomId(),"房间序列号:",RoomID)
	//预防2个人中一人强退的情况
	if def.GetRoomid(lib.mgr.GetRoomId())<1999{
		if def.M_NeedReturnRoomCard[RoomID]{
			switch def.M_GameRoomsDrawinit[RoomID] {
			case 4:{
				AddRoomCard=2

			}
			case 8:{
				AddRoomCard=3
			}
			case 12:{
				AddRoomCard=4
			}

			}


			lib.mgr.UpdateUserInfo(&def.M_RoomCreater[RoomID], &proto.GameUserPpUpdate{
				Item: &proto.ItemProp{
					ItemId: 3,
					Count:AddRoomCard,
				},
			})



		}


		def.M_RoomID[def.GetRoomid(lib.mgr.GetRoomId())]=0
	}

}
/*
func (lib *xzlib)SendCardToUser(roomindex int,SendCardUser int16){
	var Cardpai int16
	var Curaction int16
	var Gangpai int16
	Cardpai,Curaction,Gangpai=def.DispatchCardData(roomindex,SendCardUser)
	fmt.Println("给",SendCardUser,"号用户发牌;","杠牌",Gangpai,";提示有操作:",Curaction,"新发的牌",Cardpai,"总共的牌",def.M_cbCardIndex[roomindex][SendCardUser])
	lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][SendCardUser],uint32(def.SUB_S_SEND_CARD), &def.CMD_S_SendCard{
		CbCardData:Cardpai,
		CbActionMask:Curaction,
		WCurrentUser:SendCardUser,
		WSendCardUser:SendCardUser,
		CbGanData:Gangpai,
	})

	//给告诉所有用户给XX用户发牌;
	for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
		//if def.M_GameStatus
		lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_SEND_CARD_BroadCast), &def.CMD_S_SendCard_Broadcast{
			WCurrentUser:SendCardUser,
		})

	}




}
*/

func (lib *xzlib) OnGameCreate(info *defines.PlayerInfo, conf *defines.CreateRoomConf) error {
	fmt.Println("xzlib.ongamecreate", lib.mgr.GetRoomId())



	var xconf def.XzConf
	gamelib.UnMarshal(conf.Conf, &xconf)
	//if def.M_RoomCount>1999{def.M_RoomCount=0}
	//def.M_RoomCount++

var GetRoomID int
	GetRoomID=def.GetFirstZeroRoom()
	def.M_RoomConf[GetRoomID].RoomId=conf.RoomId
	def.M_RoomConf[GetRoomID].Conf=conf.Conf
	def.M_RoomID[GetRoomID]=lib.mgr.GetRoomId()
	fmt.Println("创建房间",info.RoomId,"房间index",def.M_RoomCount,"临时房间ID",lib.mgr.GetRoomId(),"conf房间ID",conf.RoomId,"xconf",xconf.PlayerCount)



  //  var ROOMID int
	//ROOMID=def.GetRoomid(lib.mgr.GetRoomId())
	//def.M_GameRecordDraw[ROOMID]       =int16(xconf.GameJU)
	def.M_DesktopPlayer[GetRoomID]        =int16(xconf.PlayerCount)

	def.M_GameMaxFan[GetRoomID]           =int16(xconf.FanCount)
	if def.M_GameMaxFan[GetRoomID]==0{def.M_GameMaxFan[GetRoomID]=14}
	def.M_GameZiMoAddtype[GetRoomID]      =xconf.ZiMoType
	def.M_GameHuan3[GetRoomID]            =xconf.Huan3 //换三张
	def.M_GameDaiyaoandJiangDui[GetRoomID]=xconf.DaiYaoJiuJiangDui
	def.M_GameMenQingZhongZhang[GetRoomID]=xconf.MenQing
	def.M_GameTianDiHu[GetRoomID]         =xconf.TianDiHu
	def.M_GameHuJiaoZhuanYi[GetRoomID]    =xconf.HuJiaoZhuanYi
    def.M_GameRoomFeiYongType[GetRoomID]  =xconf.Roomxf
	def.M_GameRoomsDrawinit[GetRoomID]    =int16(xconf.GameJU)
	def.M_RoomChairPosition[GetRoomID]    =int16(0)
	def.M_GameXiaoHU[GetRoomID]           =false

	def.M_NeedReturnRoomCard[GetRoomID] =true
	var DelRoom int
	DelRoom=0
	if xconf.GameJU==4{
		DelRoom=-2
	}
	if xconf.GameJU==8{
		DelRoom=-3
	}
	if xconf.GameJU==12{
		DelRoom=-4
	}
	lib.mgr.UpdateUserInfo(info, &proto.GameUserPpUpdate{
		Item: &proto.ItemProp{
			ItemId: 3,
			Count:DelRoom,
		},
	})


	//
	def.M_RoomCreater[GetRoomID].Uid=info.Uid
	def.M_RoomCreater[GetRoomID].UserId=info.UserId
	def.M_RoomCreater[GetRoomID].OpenId=info.OpenId
	def.M_RoomCreater[GetRoomID].Sex=info.Sex
	def.M_RoomCreater[GetRoomID].Items=info.Items
	def.M_RoomCreater[GetRoomID].Score=info.Score
	def.M_RoomCreater[GetRoomID].Account=info.Account
	def.M_RoomCreater[GetRoomID].RoomCard=info.RoomCard
	def.M_RoomCreater[GetRoomID].Diamond=info.Diamond
	def.M_RoomCreater[GetRoomID].Gold=info.Gold
	def.M_RoomCreater[GetRoomID].Name=info.Name
	def.M_RoomCreater[GetRoomID].HeadImg=info.HeadImg
	def.M_RoomCreater[GetRoomID].RoomId=info.RoomId


	def.M_GameRecord_init[GetRoomID].GameRule.FanCount   =int(def.M_GameMaxFan[GetRoomID])
	def.M_GameRecord_init[GetRoomID].GameRule.ZiMoType   =def.M_GameZiMoAddtype[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.FanCount   =int(def.M_GameMaxFan[GetRoomID])
	def.M_GameRecord_init[GetRoomID].GameRule.Huan3      =def.M_GameHuan3[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.DaiYaoJiuJiangDui=def.M_GameDaiyaoandJiangDui[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.MenQing    =def.M_GameMenQingZhongZhang[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.TianDiHu   =def.M_GameTianDiHu[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.HuJiaoZhuanYi   =def.M_GameHuJiaoZhuanYi[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.Roomxf     =def.M_GameRoomFeiYongType[GetRoomID]
	def.M_GameRecord_init[GetRoomID].GameRule.GameJU     =int(def.M_GameRoomsDrawinit[GetRoomID])
	def.M_GameRecord_init[GetRoomID].GameRule.PlayerCount=xconf.PlayerCount


	def.M_RoomChairPosition[GetRoomID]=0
	for i:=0;i<int(def.M_DesktopPlayer[GetRoomID]);i++{
		def.M_RoomChair[GetRoomID][i].UserId=0

		def.M_GameTotalScoreCount[GetRoomID].LGameScore[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LAnGang[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LMingGang[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LChaDaJia[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LZhiMo[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LJiePao[i]=0
		def.M_GameTotalScoreCount[GetRoomID].LDianPao[i]=0



	}

	def.M_NotifyProcessed[GetRoomID]=-1
	def.M_NotifyIndex[GetRoomID]=1

	return nil
}

func (lib *xzlib) OnUserEnter(info *defines.PlayerInfo) error {
	//fmt.Println("xzlib.onuserenter")
	//def.M_RoomChairPosition[def.GetRoomid(info.RoomId)]=def.M_RoomChairPosition[def.GetRoomid(info.RoomId)]%4

    fmt.Println("用户进入房间 User Enter Room",info.UserId,lib.mgr.GetRoomId(),"房间index=",def.GetRoomid(lib.mgr.GetRoomId()))

	var roomid int
	roomid=def.GetRoomid(lib.mgr.GetRoomId())
	chairid:=int(def.GetFirstNoNullid(roomid))
	def.M_RoomChairPosition[roomid]++

	if def.M_RoomChairPosition[roomid]<=def.M_DesktopPlayer[roomid]{

		if chairid==0x99{
			for i:=0;i<int(def.M_DesktopPlayer[roomid]);i++{
				def.M_RoomChair[roomid][i].UserId=0

			}
			chairid=int(def.GetFirstNoNullid(roomid))
			//fmt.Println("")
		}


		fmt.Println("房间共：",def.M_DesktopPlayer[roomid],"用户",def.M_RoomChairPosition[roomid],"坐位号:",chairid)


		fmt.Println("用户",info.UserId,"进入房间:",info.RoomId,"坐位号:",chairid)



/*
	lib.mgr.SendUserMessage(info, uint32(def.SUB_S_USER_SITDOWN), &def.EnterRoomuserInfo{
		UserId: int(info.UserId),
		ChairID:chairid,
		Name:info.Name,
		HeadImg:info.HeadImg,
		RoomCard:100,
		GameGOLD:100,
		Tm:time.Now(),
		//Tm: time.Now().Unix(),

	})*/
	lib.mgr.SendGameMessage(info, uint32(def.SUB_S_USER_SITDOWN), &def.EnterRoomuserInfo{
		UserId: int(info.UserId),
		ChairID:chairid,
		Name:info.Name,
		HeadImg:info.HeadImg,
		RoomCard:info.RoomCard,
		GameGOLD:int(info.Gold),
		Sex:int(info.Sex),
		Tm:time.Now(),
		//Tm: time.Now().Unix(),

	})


	fmt.Println("房间号ROOMID=",lib.mgr.GetRoomId(),"用户",info.UserId,"分配坐位号",chairid)

	def.M_RoomChair[roomid][chairid].Uid=info.Uid
	def.M_RoomChair[roomid][chairid].UserId=info.UserId
	def.M_RoomChair[roomid][chairid].OpenId=info.OpenId
	def.M_RoomChair[roomid][chairid].HeadImg=info.HeadImg
	def.M_RoomChair[roomid][chairid].Name=info.Name
	def.M_RoomChair[roomid][chairid].Account=info.Account
	def.M_RoomChair[roomid][chairid].Diamond=info.Diamond
	def.M_RoomChair[roomid][chairid].Gold=info.Gold
	def.M_RoomChair[roomid][chairid].RoomCard=info.RoomCard
	def.M_RoomChair[roomid][chairid].Sex=info.Sex
	def.M_RoomChair[roomid][chairid].RoomId=info.RoomId

	var OtherUser def.EnterRoomOtheruserInfo
	//OtherUser.ChairID[]

	for i:=0;i<int(def.M_DesktopPlayer[roomid]);i++{
		OtherUser.UserId[i]  =int(def.M_RoomChair[roomid][i].UserId)
		OtherUser.ChairID[i]=int(def.GetRoomChair(roomid,def.M_RoomChair[roomid][i].UserId))
		OtherUser.HeadImg[i] =def.M_RoomChair[roomid][i].HeadImg
		OtherUser.Name[i]    =def.M_RoomChair[roomid][i].Name
		OtherUser.RoomCard[i]=def.M_RoomChair[roomid][i].RoomCard
		OtherUser.GameGOLD[i]=int(def.M_RoomChair[roomid][i].Gold)

	}

	lib.mgr.SendGameMessage(info, uint32(def.SUB_S_OTHERUSERINFO), &OtherUser)


	//if getXzConf()


	def.M_GameRecord_init[roomid].GameUserID[chairid]=int(info.UserId)
	def.M_GameRecord_init[roomid].GameUserName[chairid]=info.Name
	def.M_GameRecord_init[roomid].GameUserImage[chairid]=info.HeadImg

	for i:=0;i<int(def.M_DesktopPlayer[roomid]);i++{

		lib.mgr.SendGameMessage(&def.M_RoomChair[roomid][i],uint32(def.SUB_s_User_sitdown_Broadcast),&def.EnterRoomBroadcast{
			UserId: int(info.UserId),
			ChairID:chairid,
			Name:info.Name,
			HeadImg:info.HeadImg,
            RoomCard:info.RoomCard,
			GameGOLD:int(info.Gold),
		})

	}

	if def.IsFirstSit(roomid){
		def.M_wBankerUser[roomid]=def.INVALID_CHAIR
		def.M_cbPosition[roomid]=0
		//def.M_initGameDrawScore[roomid]=
		def.WinOrder[roomid]=0
		def.M_GameRecordDraw[roomid]=0
		def.M_NextwBankerUser[roomid]=chairid
		//清等待状态
		for i:=0;i<int(def.M_DesktopPlayer[roomid]);i++{
			def.M_GameStatus[roomid][i]=0
			fmt.Println("当前玩家chairid=0",chairid,"房间",roomid,"玩家状态",def.M_GameStatus[roomid][0],def.M_GameStatus[roomid][1],def.M_GameStatus[roomid][2],def.M_GameStatus[roomid][3])
			def.M_cbWeaveItemCount[roomid][i]=int16(0)
			for j:=0;j<int(def.MAX_INDEX);j++{
				def.M_cbCardIndex[roomid][i][j]=0

			}

		}
	}

	fmt.Println("玩家进入房间成功....",info.UserId,"房间:",info.RoomId,"坐位号：",chairid)
	//房间用户位置加一
	//def.M_RoomChairPosition[def.GetRoomid(info.RoomId)]=def.M_RoomChairPosition[def.GetRoomid(info.RoomId)]+1

		return nil
	} else {
		return errors.New("roomfull")
	}
	//fmt.Println("房间各用户信息:",def.M_RoomChair[roomid][0].UserId,def.M_RoomChair[roomid][1].UserId,def.M_RoomChair[roomid][2].UserId,def.M_RoomChair[roomid][3].UserId)

}


func (lib *xzlib) OnUserLeave(info *defines.PlayerInfo) {
	fmt.Println("用户离开",info.UserId,"房间号",def.GetRoomid(info.RoomId))
	//info.UserId
	var roomindex int
	roomindex=def.GetRoomid(info.RoomId)
	var Chairid int
	Chairid=int(def.GetRoomChair(def.GetRoomid(info.RoomId),info.UserId))

	def.M_RoomChair[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),info.UserId)].UserId=0
	def.M_RoomChairPosition[def.GetRoomid(info.RoomId)]--
	var IsAllleave bool
	IsAllleave=true
	for i:=0;i<int(def.M_DesktopPlayer[def.GetRoomid(info.RoomId)]);i++{
		if def.M_RoomChair[def.GetRoomid(info.RoomId)][i].UserId!=0{
			IsAllleave=false

			lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_UserLeave_BroadCast), &def.GameUserLeave{

				Chair:Chairid,
			})
		}
	}
	if IsAllleave{}
	/*
	if IsAllleave{
		def.M_RoomID[def.GetRoomid(info.RoomId)]=0
	}
	*/
	//fmt.Println("xzlib.onuserleave")


}

func (lib *xzlib) OnUserOffline(info *defines.PlayerInfo) {
	fmt.Println("xzlib.onuseroffline")
}

func (lib *xzlib) OnUserReEnter(info *defines.PlayerInfo) {
	fmt.Println("xzlib.onuserreenter")
	//var
	//
	//var
	//
	var roomindex int
	roomindex=def.GetRoomid(info.RoomId)
	var ThisChair int16
	fmt.Println("重新进入房间：",def.M_RoomChair[roomindex][0].UserId,def.M_RoomChair[roomindex][1].UserId)
	ThisChair=def.GetRoomChair(roomindex,info.UserId)
//	for i:=0;i<int(def.MAX_INDEX);i++{

//	}//def.M_cbCardIndex[roomindex][ThisChair]
	//def.M_cbCardIndex[roomindex][ThisChair]

	type WeaveItemSend struct {
		CbWeaveKind  int16    //组合类型
		CbCenterCard int16    //中心扑克
		CbUserS      int16    //被组合的用户数
		CbProvideUser int16   //供牌用户
	}


	type ReEnter struct {
		GameUserID    [def.GAME_PLAYER] uint32
		GameUserName  [def.GAME_PLAYER] string
		GameUserImage [def.GAME_PLAYER] string
		BankerUser    int                         //庄家
		GameRule      def.XzConf                      //回放中的游戏规则;
		UserCard      [def.MAX_COUNT] int 		  //用户手中的麻将;
		DingQue       [def.GAME_PLAYER] int            //用户定缺
		DrawIndex     int                         //局数
		ChairID       int
		LEAVECARDCOUNT int
		GangScore[def.GAME_PLAYER] int
		WeaveItem[def.GAME_PLAYER][4] WeaveItemSend
		Hu[def.GAME_PLAYER] bool//判断用户是否已经胡牌
		UserOutCardList[def.GAME_PLAYER][41] def.OutCardstruct //用户打出去的麻将列表
		GameStatus int //游戏状态
		//RoomCard int
		//GameGOLD int
		DispachUser int//发牌用户
		Card int //麻将数据
		Useraction int//当前需要操作
		UserCardCount[def.GAME_PLAYER] int//用户手中的麻将张数
		HuCard[def.GAME_PLAYER] int//胡牌的麻将
		Huan3[3] int //换三张
	}
	var VReEnter ReEnter
	var CardCount int
	var CardIndex int
	var LastStatus int
	var IsAllStatus  bool
	LastStatus=int(def.M_GameStatus[roomindex][ThisChair])
	IsAllStatus=true
	//断线重连发送手牌
	for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++ {
		if LastStatus>int(def.M_GameStatus[roomindex][i]) && def.M_GameStatus[roomindex][i] != def.GAME_STATUS_WINED {
			IsAllStatus = false
		}
	}
	VReEnter.GameStatus=LastStatus
	fmt.Println("游戏状态:",LastStatus,"各玩家状态:",def.M_GameStatus[roomindex][0],def.M_GameStatus[roomindex][1],def.M_GameStatus[roomindex][2],def.M_GameStatus[roomindex][3])
	if !IsAllStatus{
		switch int16(LastStatus){
		case def.GAME_STATUS_CHANGETHREE:{VReEnter.GameStatus=int(def.GAME_STATUS_Ready)}
		case def.GAME_STATUS_DingQUE    :{
			if def.M_GameHuan3[roomindex]{
				VReEnter.GameStatus=int(def.GAME_STATUS_CHANGETHREE)
			}else {
				VReEnter.GameStatus=int(def.GAME_STATUS_Ready)
			}
		}
		case def.GAME_STATUS_PLAY:{VReEnter.GameStatus=int(def.GAME_STATUS_DingQUE)}
		}

	}

	fmt.Println("最后返回的游戏状态:",VReEnter.GameStatus)
	for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{

		VReEnter.GameUserID[i]   =def.M_RoomChair[roomindex][i].UserId
		VReEnter.GameUserName[i] =def.M_RoomChair[roomindex][i].Name
		VReEnter.GameUserImage[i]=def.M_RoomChair[roomindex][i].HeadImg

		if (VReEnter.GameStatus==int(def.GAME_STATUS_Ready))&&(def.M_GameStatus[roomindex][i]==def.GAME_STATUS_Ready){
			VReEnter.DingQue[i]      =-1
			fmt.Println("游戏用户",i,"值为--1")
		}else{
			VReEnter.DingQue[i]      =int(def.M_dingqueColor[roomindex][i])
		}

		fmt.Println("断线重连中：用户",i,"组合牌张数",def.M_cbWeaveItemCount[roomindex][i],"游戏状态:",def.M_GameStatus[roomindex][i],"整个游戏状态",VReEnter.GameStatus)
		for j:=0;j<int(def.M_cbWeaveItemCount[roomindex][i]);j++{
			VReEnter.WeaveItem[i][j].CbCenterCard  =def.M_WeaveItemArray[roomindex][i][j].CbCenterCard
			VReEnter.WeaveItem[i][j].CbWeaveKind   =def.M_WeaveItemArray[roomindex][i][j].CbWeaveKind
			VReEnter.WeaveItem[i][j].CbProvideUser =def.M_WeaveItemArray[roomindex][i][j].WProvideUser
			VReEnter.WeaveItem[i][j].CbUserS       =def.M_WeaveItemArray[roomindex][i][j].BeGangUserCount

		}
		VReEnter.Hu[i]=false
		if def.M_GameStatus[roomindex][i]==def.GAME_STATUS_WINED{
			VReEnter.Hu[i]=true
		}

		VReEnter.ChairID=int(ThisChair)
		VReEnter.LEAVECARDCOUNT=int(def.MAX_REPERTORY-def.M_cbPosition[roomindex])
		VReEnter.GangScore[i]=0





		for j:=0;j<int(def.M_DesktopPlayer[roomindex]);j++{
			VReEnter.GangScore[i]=VReEnter.GangScore[i]+int(def.M_GangScore[roomindex][i][j])
		}
		for j:=0;j<41;j++{
			VReEnter.UserOutCardList[i][j].OutCard=def.M_OutCardList[roomindex][i][j].OutCard
			VReEnter.UserOutCardList[i][j].ChairID=def.M_OutCardList[roomindex][i][j].ChairID
		}
		CardCount=0
		CardIndex=0
		for j:=0;j<int(def.MAX_INDEX);j++{
			if def.M_cbCardIndex[roomindex][i][j]>0{CardCount=CardCount+int(def.M_cbCardIndex[roomindex][i][j])
				if i==int(ThisChair){
					for tmpj:=0;tmpj<int(def.M_cbCardIndex[roomindex][i][j]);tmpj++{
						VReEnter.UserCard[CardIndex]=j
						CardIndex++
					}
				}
			}

		}
		VReEnter.UserCardCount[i]=CardCount

	}
	VReEnter.Useraction=int(def.M_cbUserAction[roomindex][ThisChair])
	if def.M_GameStatus[roomindex][ThisChair]==def.GAME_STATUS_CHANGETHREE{
		VReEnter.Huan3[0]=int(def.M_cbChangeThree[roomindex][ThisChair][0])
		VReEnter.Huan3[1]=int(def.M_cbChangeThree[roomindex][ThisChair][1])
		VReEnter.Huan3[2]=int(def.M_cbChangeThree[roomindex][ThisChair][2])
	} else {
		VReEnter.Huan3[0]=0
		VReEnter.Huan3[1]=0
		VReEnter.Huan3[2]=0
	}


	VReEnter.GameRule.PlayerCount=int(def.M_DesktopPlayer[roomindex])
	fmt.Println("回放人数",VReEnter.GameRule.PlayerCount)
	VReEnter.GameRule.FanCount   =int(def.M_GameMaxFan[roomindex])
	VReEnter.GameRule.ZiMoType   =def.M_GameZiMoAddtype[roomindex]
	VReEnter.GameRule.FanCount   =int(def.M_GameMaxFan[roomindex])
	VReEnter.GameRule.Huan3      =def.M_GameHuan3[roomindex]
	VReEnter.GameRule.DaiYaoJiuJiangDui=def.M_GameDaiyaoandJiangDui[roomindex]
	VReEnter.GameRule.MenQing    =def.M_GameMenQingZhongZhang[roomindex]
	VReEnter.GameRule.TianDiHu   =def.M_GameTianDiHu[roomindex]
	VReEnter.GameRule.HuJiaoZhuanYi   =def.M_GameHuJiaoZhuanYi[roomindex]
	VReEnter.GameRule.Roomxf     =def.M_GameRoomFeiYongType[roomindex]
	VReEnter.GameRule.GameJU     =int(def.M_GameRoomsDrawinit[roomindex])
	VReEnter.BankerUser=int(def.M_wBankerUser[roomindex])
	//VReEnter.GameStatus=int(def.M_GameStatus[roomindex][ThisChair])
	VReEnter.DispachUser=def.M_DispachUser[roomindex]
	if def.M_cbPosition[roomindex]>1&&def.M_cbPosition[roomindex]<107{
		VReEnter.Card=int(def.M_cbCardData[roomindex][def.M_cbPosition[roomindex]-1])
	}

	//VReEnter.
	//VReEnter.GameRule=

	//lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][ThisChair],uint32(def.SUB_S_GAME_ReenterRoom), &VReEnter)
	def.M_RoomChair[roomindex][ThisChair].Uid=info.Uid
	def.M_RoomChair[roomindex][ThisChair].HeadImg=info.HeadImg
	def.M_RoomChair[roomindex][ThisChair].Name=info.Name
	def.M_RoomChair[roomindex][ThisChair].Gold=info.Gold
	def.M_RoomChair[roomindex][ThisChair].Diamond=info.Diamond
	def.M_RoomChair[roomindex][ThisChair].RoomCard=info.RoomCard
	def.M_RoomChair[roomindex][ThisChair].Account=info.Account
	def.M_RoomChair[roomindex][ThisChair].Score=info.Score
	def.M_RoomChair[roomindex][ThisChair].Sex=info.Sex
	def.M_RoomChair[roomindex][ThisChair].Items=info.Items


	lib.mgr.SendGameMessage(info,uint32(def.SUB_S_GAME_ReenterRoom), &VReEnter)
	fmt.Println("用户",info.UserId,"断线重连","房间",info.RoomId,"用户手牌:",def.M_cbCardIndex[roomindex][ThisChair],"发送数据",VReEnter)


}

//游戏流局
func GameLiuJU(roomid int){
	//已经修改


	var cardid int16


	//fmt.Println("游戏流局 以下为流局数据.................")
	//fmt.Println("游戏流局 以下为流局数据.................")


	var WinedUserCount int
	var YouJiaoUserCount int
	var MeiJiaoUserCount int
	WinedUserCount=0
	YouJiaoUserCount=0
	MeiJiaoUserCount=0
	//取得每个用户手中的牌
	for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomid]);tmpi++{
		cardid=0
		fmt.Println("用户",tmpi,"手中的麻将：",def.M_cbCardIndex[roomid][tmpi])
		for tmpj:=0;tmpj<int(def.MAX_INDEX);tmpj++{
			if def.M_cbCardIndex[roomid][tmpi][tmpj]>0{
				for tmph:=0;tmph<int(def.M_cbCardIndex[roomid][tmpi][tmpj]);tmph++{
					def.M_GameConcludeScore[roomid].CbHandCardData[tmpi][cardid]=int16(tmpj)
					cardid++
				}

			}
		}
		if def.M_GameStatus[roomid][tmpi]==def.GAME_STATUS_WINED{//已经胡牌用户数 判断流局庄家用
			WinedUserCount++
		}
	}

  fmt.Println("流局前各玩家杠分状态:",def.M_GangScore[roomid][0],def.M_GangScore[roomid][1],def.M_GangScore[roomid][2],def.M_GangScore[roomid][3])
  fmt.Println("流局前各玩家游戏积分",def.M_GameConcludeScore[roomid].LGameScore[0],def.M_GameConcludeScore[roomid].LGameScore[1],def.M_GameConcludeScore[roomid].LGameScore[2],def.M_GameConcludeScore[roomid].LGameScore[3])

	//牌满 流局，注明流局状态;
	for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomid]);tmpi++{
		if def.M_GameStatus[roomid][tmpi]!=def.GAME_STATUS_WINED{
			//此处所有玩家均为流局
			var isliujuyoujiao bool//流局有叫
			isliujuyoujiao=false
			for tmpj:=0;tmpj<int(def.MAX_INDEX);tmpj++{
				if def.AnalyseChiHuCard_UserSendCard(roomid,def.M_cbCardIndex[roomid][tmpi],int16(tmpi),int16(tmpj))!=def.WIK_NULL{
					isliujuyoujiao=true
					//有叫
					//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].WLiuJuStatus[tmpi]=2
				}

			}
			if isliujuyoujiao{
				def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi]=int16(def.LiuJuStatus_YouJiao)
				def.M_GameConcludeScore[roomid].HuType[tmpi]=def.HuType_ChaJiao


			}else {
				def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi]=int16(def.LiuJustatus_MeiJiao)
			}
			fmt.Println("玩家:",tmpi,"手中的麻将数据：",def.M_cbCardIndex[roomid][tmpi],"判断用户是否花猪:",roomid)
			if def.IsHuaZhu(roomid,def.M_cbCardIndex[roomid][tmpi],int16(tmpi)){
				fmt.Println("玩家:",tmpi,"手中的麻将数据：",def.M_cbCardIndex[roomid][tmpi],"判断用户是否花猪:")
				def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi]=int16(def.LiuJuStatus_HuaZhu)
			}

			switch def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi] {
			case int16(def.LiuJuStatus_YouJiao):{YouJiaoUserCount++}
			case int16(def.LiuJustatus_MeiJiao):{MeiJiaoUserCount++}

			}
			def.M_GameConcludeScore[roomid].UFanDescBase[tmpi]=""
			def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpi]=""
			def.M_GameConcludeScore[roomid].UFanBase[tmpi]=0
			def.M_GameConcludeScore[roomid].UFanAddtion[tmpi]=0
			//def.UserChaJiao()
		}
	}//tmpi
	fmt.Println("各玩家流局状态：",def.M_GameConcludeScore[roomid].WLiuJuStatus[0],def.M_GameConcludeScore[roomid].WLiuJuStatus[1],def.M_GameConcludeScore[roomid].WLiuJuStatus[2],def.M_GameConcludeScore[roomid].WLiuJuStatus[3])
	fmt.Println("各玩家流局时基本番：",def.M_GameConcludeScore[roomid].UFanBase[0],def.M_GameConcludeScore[roomid].UFanBase[1],def.M_GameConcludeScore[roomid].UFanBase[2],def.M_GameConcludeScore[roomid].UFanBase[3])
	fmt.Println("各玩家流局时额外番：",def.M_GameConcludeScore[roomid].UFanAddtion[0],def.M_GameConcludeScore[roomid].UFanAddtion[1],def.M_GameConcludeScore[roomid].UFanAddtion[2],def.M_GameConcludeScore[roomid].UFanAddtion[3])
	var ChangeBlankUserTO int
	ChangeBlankUserTO=0
	if WinedUserCount==0&&MeiJiaoUserCount==1{ChangeBlankUserTO=3}//无叫仅一个人
	if WinedUserCount==0&&YouJiaoUserCount==1{ChangeBlankUserTO=2}//有叫仅一个人
    var MyliuJuStatus int
	MyliuJuStatus=0
	for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomid]);tmpi++{
		MyliuJuStatus=int(def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi])
		switch MyliuJuStatus {
		case def.LiuJuStatus_HuaZhu:{//花猪
			/*
			var JiFen int
			switch def.M_GameMaxFan[roomid] {
			case 0:{JiFen=1}
			case 1:{JiFen=2}
			case 2:{JiFen=4}
			case 3:{JiFen=8}
			case 4:{JiFen=16}
			case 5:{JiFen=32}
			case 6:{JiFen=64}
			case 7:{JiFen=128}
			case 8:{JiFen=256}
			case 9:{JiFen=512}
			case 10:{JiFen=1024}
			case 11:{JiFen=2048}
			case 12:{JiFen=4096}
			case 13:{JiFen=8192}
			case 14:{JiFen=16384}
			//case 15:{JiFen=32768}
			}
			for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomid]);tmpj++{
				//花猪赔所有有叫用户的满格
				//if def.M_GameStatus[def.GetRoomid(info.RoomId)][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].WLiuJuStatus[tmpj]==2{
				//花猪赔所有未胡牌用户的满格

				if def.M_GameStatus[roomid][tmpj]!=def.GAME_STATUS_WINED{
					def.M_GameConcludeScore[roomid].LGameScore[tmpj]=def.M_GameConcludeScore[roomid].LGameScore[tmpj]+JiFen
					def.M_GameConcludeScore[roomid].LGameScore[tmpi]=def.M_GameConcludeScore[roomid].LGameScore[tmpi]-JiFen
					def.M_GameConcludeScore[roomid].WProvideUser[tmpj]=int16(tmpi)
					def.M_GameConcludeScore[roomid].UFanDescBase[tmpj]="花猪"
					def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]="花猪"

				}
			}//for tmpj
			*/
			fmt.Println("玩家",tmpi,"是花猪")
			for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomid]);tmpj++{
				//无叫人赔所有有叫用户的最大叫
				if def.M_GameStatus[roomid][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpj]==2{
					//if def.M_GameStatus[def.GetRoomid(info.RoomId)][tmpj]!=def.GAME_STATUS_WINED&&tmpi!=tmpj{
					var chajiaofan[2] int16
					var chajiaostrbase string
					var chajiaostradd string
					var JiFen int16

					chajiaofan[0],chajiaofan[1],chajiaostrbase,chajiaostradd,JiFen=def.UserChaJiao(roomid,int16(tmpj))
					fmt.Println("花猪用户原来基本番",def.M_GameConcludeScore[roomid].UFanBase[tmpi],"花猪用户原来额外番",def.M_GameConcludeScore[roomid].UFanAddtion[tmpi],"赔用户",tmpj,"基本番",chajiaofan[0],"额外番",chajiaofan[1])
					def.M_GameConcludeScore[roomid].UFanDescBase[tmpj]=def.M_GameConcludeScore[roomid].UFanDescBase[tmpj]+chajiaostrbase
					def.M_GameConcludeScore[roomid].UFanBase[tmpi]=def.M_GameConcludeScore[roomid].UFanBase[tmpi]-chajiaofan[0]
					def.M_GameConcludeScore[roomid].UFanAddtion[tmpi]=def.M_GameConcludeScore[roomid].UFanAddtion[tmpi]-chajiaofan[1]
					def.M_GameConcludeScore[roomid].UFanBase[tmpj]=chajiaofan[0]
					def.M_GameConcludeScore[roomid].UFanAddtion[tmpj]=chajiaofan[1]
					//def.M_GameConcludeScore[roomid].UFanDescBase[tmpi]=def.M_GameConcludeScore[roomid].UFanDescBase[tmpi]-chajiaostrbase
					def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]=def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]+chajiaostradd
					//def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]=def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]+chajiaostradd





					def.M_GameConcludeScore[roomid].LGameScore[tmpj]=def.M_GameConcludeScore[roomid].LGameScore[tmpj]+int(JiFen)
					def.M_GameConcludeScore[roomid].LGameScore[tmpi]=def.M_GameConcludeScore[roomid].LGameScore[tmpi]-int(JiFen)
					def.M_GameConcludeScore[roomid].WProvideUser[tmpj]=int16(tmpi)
				}



			}//for tmpj

			fmt.Println("所有用户杠分",def.M_GangScore[roomid],"玩家",tmpi,"组合牌数",def.M_cbWeaveItemCount[roomid][tmpi])
			//for tmps:=0;tmps<int()
			var GangedJiFen int
			GangedJiFen=1
			for tmps:=0;tmps<int(def.M_cbWeaveItemCount[roomid][tmpi]);tmps++{
				GangedJiFen=1
				if def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_XiaYu{
					GangedJiFen=2
				}
				if (!def.M_WeaveItemArray[roomid][tmpi][tmps].GangSangPao)&&((def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_GuaFeng)||(def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_XiaYu)){
					if def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser!=int16(tmpi){
						fmt.Println("关1家的刮风下雨")
						def.M_GangScore[roomid][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser][tmpi]=def.M_GangScore[roomid][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser][tmpi]+int16(GangedJiFen)//
						def.M_GangScore[roomid][tmpi][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser]=def.M_GangScore[roomid][tmpi][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser]-int16(GangedJiFen)
					}
					if def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser==int16(tmpi){
						fmt.Println("关多家的刮风下雨")
						for tmpkj:=0;tmpkj<int(def.M_DesktopPlayer[roomid]);tmpkj++{
							fmt.Println("接分玩家",tmpkj,"退分玩家",tmpi,"房间ID",roomid,"杠用户状态",def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList,"杠用户222 状态",def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList[tmpkj],"组合",tmps,"组合牌数据",def.M_WeaveItemArray[roomid][tmpi][tmps])
							if (def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList[tmpkj]==1)&&(tmpkj!=tmpi){
								//if tmpkj!=tmpi{
								fmt.Println("1111接分玩家",tmpkj,"1111退分玩家",tmpi,"原",def.M_GangScore[roomid][tmpkj][tmpi],"原需要退分的用户分",def.M_GangScore[roomid][tmpi][tmpkj])
								def.M_GangScore[roomid][tmpkj][tmpi]=def.M_GangScore[roomid][tmpkj][tmpi]+int16(GangedJiFen)
								def.M_GangScore[roomid][tmpi][tmpkj]=def.M_GangScore[roomid][tmpi][tmpkj]-int16(GangedJiFen)
								fmt.Println("1111接分玩家",tmpkj,"2222退分玩家",tmpi,"接分后",def.M_GangScore[roomid][tmpkj][tmpi],"原需要退分的用户分",def.M_GangScore[roomid][tmpi][tmpkj])
							}
						}
					}

				}





			}


/*
			for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomid]);tmpj++{
				//if def.M_GameStatus[def.GetRoomid(info.RoomId)][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].WLiuJuStatus[tmpj]==2 {
				if def.M_GameStatus[roomid][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpj]!=1{
					//某某供牌用户的杠分;
					def.M_GangScore[roomid][tmpi][tmpj]=0
				}
			}//for tmpj

*/

		}
		case def.LiuJuStatus_YouJiao:{//有叫
             if ChangeBlankUserTO==2{def.M_NextwBankerUser[roomid]=tmpi}//如果只有一家人有叫则有叫玩家当庄

		}
		case def.LiuJustatus_MeiJiao:{//没叫
			if ChangeBlankUserTO==3{def.M_NextwBankerUser[roomid]=tmpi}//如果只有一家人没叫则没叫玩家当庄

			for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomid]);tmpj++{
				fmt.Println("用户",tmpi,"流局状态为",def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi],"无叫 原来的杠分为",def.M_GangScore[roomid][tmpi],"流局用户",tmpj,"流局状态",def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpj])
				//无叫人赔所有有叫用户的最大叫
				if def.M_GameStatus[roomid][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpj]==2{
					//if def.M_GameStatus[def.GetRoomid(info.RoomId)][tmpj]!=def.GAME_STATUS_WINED&&tmpi!=tmpj{
					var chajiaofan[2] int16
					var chajiaostrbase string
					var chajiaostradd string
					var JiFen int16
					chajiaofan[0],chajiaofan[1],chajiaostrbase,chajiaostradd,JiFen=def.UserChaJiao(roomid,int16(tmpj))
					fmt.Println("用户",tmpj,"有叫 查叫番数",chajiaofan[0],chajiaofan[1],chajiaostrbase,chajiaostradd,"查叫积分为:",JiFen,"原有叫用户积分",def.M_GameConcludeScore[roomid].LGameScore[tmpj],"无叫用户积分",def.M_GameConcludeScore[roomid].LGameScore[tmpi],"无叫用户",tmpi,"无叫用户流局状态",def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpi])
					def.M_GameConcludeScore[roomid].UFanDescBase[tmpj]=chajiaostrbase
					def.M_GameConcludeScore[roomid].UFanBase[tmpi]=def.M_GameConcludeScore[roomid].UFanBase[tmpi]-chajiaofan[0]
					def.M_GameConcludeScore[roomid].UFanAddtion[tmpi]=def.M_GameConcludeScore[roomid].UFanAddtion[tmpi]-chajiaofan[1]
					def.M_GameConcludeScore[roomid].UFanBase[tmpj]=chajiaofan[0]
					def.M_GameConcludeScore[roomid].UFanAddtion[tmpj]=chajiaofan[1]

					def.M_GameConcludeScore[roomid].UFanDescAddtion[tmpj]=chajiaostradd

					fmt.Println("用户",tmpj,"有叫 查叫番数",chajiaofan[0],chajiaofan[1],chajiaostrbase,chajiaostradd,"查叫积分为:",JiFen)


					def.M_GameConcludeScore[roomid].LGameScore[tmpj]=def.M_GameConcludeScore[roomid].LGameScore[tmpj]+int(JiFen)
					def.M_GameConcludeScore[roomid].LGameScore[tmpi]=def.M_GameConcludeScore[roomid].LGameScore[tmpi]-int(JiFen)
					//def.M_GangScore[roomid][tmpi][tmpj]



					def.M_GameConcludeScore[roomid].WProvideUser[tmpj]=int16(tmpi)
				}



			}//for tmpj
			fmt.Println("所有用户杠分",def.M_GangScore[roomid],"玩家",tmpi,"组合牌数",def.M_cbWeaveItemCount[roomid][tmpi])
			//for tmps:=0;tmps<int()
			var GangedJiFen int
			GangedJiFen=1
			for tmps:=0;tmps<int(def.M_cbWeaveItemCount[roomid][tmpi]);tmps++{
				GangedJiFen=1
				if def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_XiaYu{
					GangedJiFen=2
				}
				if (def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_GuaFeng)||(def.M_WeaveItemArray[roomid][tmpi][tmps].CbWeaveKind==def.WIK_XiaYu){
					if def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser!=int16(tmpi){
						fmt.Println("关1家的刮风下雨")
						def.M_GangScore[roomid][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser][tmpi]=def.M_GangScore[roomid][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser][tmpi]+int16(GangedJiFen)//
						def.M_GangScore[roomid][tmpi][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser]=def.M_GangScore[roomid][tmpi][def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser]-int16(GangedJiFen)
					}
					if def.M_WeaveItemArray[roomid][tmpi][tmps].WProvideUser==int16(tmpi){
                        fmt.Println("关多家的刮风下雨")
						for tmpkj:=0;tmpkj<int(def.M_DesktopPlayer[roomid]);tmpkj++{
							fmt.Println("接分玩家",tmpkj,"退分玩家",tmpi,"房间ID",roomid,"杠用户状态",def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList,"杠用户222 状态",def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList[tmpkj],"组合",tmps,"组合牌数据",def.M_WeaveItemArray[roomid][tmpi][tmps])
							if (def.M_WeaveItemArray[roomid][tmpi][tmps].GangUserList[tmpkj]==1)&&(tmpkj!=tmpi){
							//if tmpkj!=tmpi{
								fmt.Println("1111接分玩家",tmpkj,"要退的积分",GangedJiFen,"1111退分玩家",tmpi,"原",def.M_GangScore[roomid][tmpkj][tmpi],"原需要退分的用户分",def.M_GangScore[roomid][tmpi][tmpkj])
								def.M_GangScore[roomid][tmpkj][tmpi]=def.M_GangScore[roomid][tmpkj][tmpi]+int16(GangedJiFen)
								def.M_GangScore[roomid][tmpi][tmpkj]=def.M_GangScore[roomid][tmpi][tmpkj]-int16(GangedJiFen)
								fmt.Println("1111接分玩家",tmpkj,"2222退分玩家",tmpi,"接分后",def.M_GangScore[roomid][tmpkj][tmpi],"原需要退分的用户分",def.M_GangScore[roomid][tmpi][tmpkj])
							}
						}
					}

				}





			}
			/*

			for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomid]);tmpj++{

				//if def.M_GameStatus[def.GetRoomid(info.RoomId)][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].WLiuJuStatus[tmpj]==2 {
				//if def.M_GameStatus[roomid][tmpj]!=def.GAME_STATUS_WINED&&tmpj!=tmpi&&def.M_GameConcludeScore[roomid].WLiuJuStatus[tmpj]==2{
				if tmpj!=tmpi{




					//某某供牌用户的杠分;
					if def.M_GangScore[roomid][tmpi][tmpj]>0{
						fmt.Println("原来用户:",tmpi,"收到",tmpj,"用户雨钱 ，现退雨钱:",def.M_GangScore[roomid][tmpi][tmpj],"退雨后",def.M_GangScore[roomid][tmpj][tmpi]+def.M_GangScore[roomid][tmpi][tmpj])
						def.M_GangScore[roomid][tmpj][tmpi]=def.M_GangScore[roomid][tmpj][tmpi]+def.M_GangScore[roomid][tmpi][tmpj]
						def.M_GangScore[roomid][tmpi][tmpj]=0


					}





				}
			}//for tmpj
			*/


		} //case
		}
		if def.M_GameStatus[roomid][tmpi]!=def.GAME_STATUS_WINED{
		def.M_GameTotalScoreCount[roomid].LChaDaJia[tmpi]++//被查大叫的次数;
		}
		fmt.Println("流局判断用户",tmpi,"流局中各玩家游戏积分",def.M_GameConcludeScore[roomid].LGameScore[0],def.M_GameConcludeScore[roomid].LGameScore[1],def.M_GameConcludeScore[roomid].LGameScore[2],def.M_GameConcludeScore[roomid].LGameScore[3])
	}//for tmpi;

	fmt.Println("流局退税后杠分状态:",def.M_GangScore[roomid][0],def.M_GangScore[roomid][1],def.M_GangScore[roomid][2],def.M_GangScore[roomid][3])
	fmt.Println("流局后各玩家游戏积分",def.M_GameConcludeScore[roomid].LGameScore[0],def.M_GameConcludeScore[roomid].LGameScore[1],def.M_GameConcludeScore[roomid].LGameScore[2],def.M_GameConcludeScore[roomid].LGameScore[3])
	for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomid]);tmpi++ {
		def.M_GameConcludeScore[roomid].LGangScore[tmpi]=0
		for tmpj := 0; tmpj < int(def.M_DesktopPlayer[roomid]); tmpj++ {
			def.M_GameConcludeScore[roomid].LGangScore[tmpi] = def.M_GameConcludeScore[roomid].LGangScore[tmpi] + def.M_GangScore[roomid][tmpi][tmpj]
		}
	}
	def.M_GameRecordDraw[roomid]++


}






func (lib *xzlib) OnUserMessage(info *defines.PlayerInfo, cmd uint32, data []byte) error {

//	fmt.Println("收到用户命令",cmd)
	var roomindex int
	roomindex=def.GetRoomid(info.RoomId)

	switch int16(cmd) {

	case def.SUB_C_GpsPosition:
		{
			type GamePosition struct {
				X float64
				Y float64
			}
			var Position GamePosition
			gamelib.UnMarshal(data, &Position)
			var chairposition int16
			chairposition=def.GetRoomChair(roomindex,info.UserId)
			def.M_Position[roomindex].X[chairposition]=Position.X
			def.M_Position[roomindex].Y[chairposition]=Position.Y
			def.M_Position[roomindex].Chair[chairposition]=int(chairposition)
			def.M_PositionCount[roomindex]++
			fmt.Println("接收到坐标位置",Position.X,Position.Y," 已经收到坐标人数：",def.M_PositionCount[roomindex])

            if def.M_PositionCount[roomindex]>=int(def.M_DesktopPlayer[roomindex]){

				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
					fmt.Println("返回坐标 x",def.M_Position[roomindex].X,"Y",def.M_Position[roomindex].Y)

					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_Position), &def.GamePositionRoom{
						X:def.M_Position[roomindex].X,
						Y:def.M_Position[roomindex].Y,
						Chair:def.M_Position[roomindex].Chair,
					})

				}
			}

		}


	case def.SUB_C_READY: //用户准备好
		{
			type readySt struct {

				Fsd 	int
			}

			var r readySt
			gamelib.UnMarshal(data, &r)
			//fmt.Println("r", r, data)

			type readyStatus struct {

				Chairid 	int
			}

//			lib.mgr.BroadcastMessage(uint32(def.SUB_S_USER_READY), &readyStatus{
//				Chairid:int(def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))),
//			})
			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
				//if i!=int(def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))){
				if def.M_RoomChair[roomindex][i].UserId!=0{
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_USER_READY), &readyStatus{
						Chairid:int(def.GetRoomChair(roomindex,info.UserId)),
					})

				}

				//}

			}


			fmt.Println("接收到用户准备好...........",info.UserId,roomindex,def.GetRoomChair(roomindex,info.UserId))
			def.SetToGamePlay(roomindex,def.GetRoomChair(roomindex,info.UserId),def.GAME_STATUS_Ready)



			fmt.Println("玩家状态",def.M_GameStatus[roomindex][0],def.M_GameStatus[roomindex][1],def.M_GameStatus[roomindex][2],def.M_GameStatus[roomindex][3])
			var GameStarttmp bool=true
			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
				if  def.GetGamePlay_Status(roomindex,int16(i))!=def.GAME_STATUS_Ready {
					GameStarttmp = false

				}
				for j:=0;j<int(def.MAX_INDEX);j++{
					//清除原来的麻将
					def.M_cbCardIndex[roomindex][i][j]=0

				}
				for j:=0;j<int(def.MAX_COUNT);j++{
				//清原来玩家手牌数据;
				def.M_GameConcludeScore[roomindex].CbHandCardData[i][j]=0
				}
				//fmt.Println("玩家数据....清手牌成功")
				//清除原来的组合牌
				def.M_cbWeaveItemCount[roomindex][i]=0
				//清原来结算的番型
				def.M_GameConcludeScore[roomindex].UFanDescAddtion[i]=""
				def.M_GameConcludeScore[roomindex].UFanDescBase[i]=""
				def.M_GameConcludeScore[roomindex].UFanAddtion[i]=0
				def.M_GameConcludeScore[roomindex].UFanBase[i]=0
				def.M_GameConcludeScore[roomindex].LGameScore[i]=0


				//清除杠牌得分
				for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
					def.M_GangScore[roomindex][i][tmpi]=0
					def.M_GameConcludeScore[roomindex].LGameScore[i]=0
					def.M_OutCardListIndex[roomindex][tmpi]=0
					def.M_cbChangeThree[roomindex][tmpi][0]=0
					def.M_cbChangeThree[roomindex][tmpi][1]=0
					def.M_cbChangeThree[roomindex][tmpi][2]=0
					def.M_GameConcludeScore[roomindex].HuType[tmpi]=0
					def.M_GameConcludeScore[roomindex].CbProvideCard[tmpi]=0
					for tmpj:=0;tmpj<41;tmpj++{
						def.M_OutCardList[roomindex][tmpi][tmpj].OutCard=0
					}

				}

				def.WinOrder[roomindex]=0
				def.M_GameRecord_Operator_Index[roomindex]=0
				//初始化原来积分
				def.M_initGameDrawScore[roomindex]=1
				def.M_PositionCount[roomindex]=0



			}

			if GameStarttmp{
				def.M_NeedReturnRoomCard[roomindex]=false
				/*
				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{

					//fmt.Println("发送GAMESTART BEFORE")
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.Sub_S_Position), &def.CMD_S_GameStart{
						WBankerUser:def.GetBankerUser(roomindex),
						WSice1:def.M_Sice[0],
						WSice2:def.M_Sice[1],
						CbCardData:CbCardData,
						Draw:def.M_GameRecordDraw[roomindex],
					})

				}
               */
				if def.GetBankerUser(roomindex)==int16(def.INVALID_CHAIR){
					fmt.Println("进入定庄家，丢骰子")
					rand.Seed(time.Now().Unix())


					def.SetBankerUser(roomindex,int16(rand.Intn(int(def.M_DesktopPlayer[roomindex]))))
					def.M_Sice[0]=def.OutSice()
					def.M_Sice[1]=def.OutSice1()
					def.M_NextwBankerUser[roomindex]=int(def.M_wBankerUser[roomindex])
				}else{
					def.SetBankerUser(roomindex,int16(def.M_NextwBankerUser[roomindex]))
				}
				def.M_GameRecord_init[roomindex].Sice[0]=int(def.M_Sice[0])
				def.M_GameRecord_init[roomindex].Sice[1]=int(def.M_Sice[1])

				fmt.Println("进入待发送状态")
				def.InitCard(roomindex)
				def.M_cbPosition[roomindex]=0
				//游戏记录麻将保存
				for i:=0;i<int(def.MAX_REPERTORY);i++{
					def.M_GameRecord_init[roomindex].Cardlist[i]=int(def.M_cbCardData[roomindex][i])
				}

				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
					var CbCardData[def.MAX_COUNT] int16
					var SendCards int16=13
					if i==int(def.GetBankerUser(roomindex)){
						SendCards=14
					}else {
						SendCards=13
					}
					//笨办法复制麻将数据
					fmt.Println("发送麻将",SendCards,"麻将位置",def.M_cbPosition[roomindex])
					for j:=0;j<int(SendCards);j++{
						CbCardData[j]=def.M_cbCardData[roomindex][j+int(def.M_cbPosition[roomindex])]
						def.M_GameRecord_init[roomindex].UserCard[i][j]=int(def.M_cbCardData[roomindex][j+int(def.M_cbPosition[roomindex])])
						//fmt.Println("麻将",j,CbCardData[j])
					}
					def.M_GameRecord_init[roomindex].DrawIndex=int(def.M_GameRecordDraw[roomindex])
                    //fmt.Println("发送GAMESTART BEFORE")
					if def.M_RoomChair[roomindex][i].UserId!=0{


						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_GAME_START), &def.CMD_S_GameStart{
							WBankerUser:def.GetBankerUser(roomindex),
							WSice1:def.M_Sice[0],
							WSice2:def.M_Sice[1],
							CbCardData:CbCardData,
							Draw:def.M_GameRecordDraw[roomindex],
						})
					}
					fmt.Println("骰子点数",def.M_Sice[0],def.M_Sice[1])
					//fmt.Println("发送GAMESTART AFTER")
					def.M_cbPosition[roomindex]+=SendCards
					def.SetProvideUser(roomindex,def.GetRoomChair(roomindex,info.UserId))
					def.SwitchToCardIndex_User(roomindex,CbCardData,SendCards,int16(i))
					//fmt.Println("发送SwitchToCardIndex_User After")

                    if def.M_GameHuan3[roomindex]{

						//发送推荐换三张的消息
						var change3[3] int16

						change3[0],change3[1],change3[2]=def.RefChange3(def.M_cbCardIndex[roomindex][i])
						//fmt.Println("发送换三张",)
						fmt.Println("发送推荐的三张麻将",change3[0],change3[1],change3[2])
						if def.M_RoomChair[roomindex][i].UserId!=0 {
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_CHANGETHREE_TUIJIAN), &def.CMD_CHANGE_THREE{
								CbUser: def.GetRoomChair(roomindex, info.UserId),
								Cbcard: change3,
							})
						}
						//保存换三张
						def.M_cbChangeThree[roomindex][def.GetRoomChair(roomindex,info.UserId)][0]=change3[0]
						def.M_cbChangeThree[roomindex][def.GetRoomChair(roomindex,info.UserId)][1]=change3[1]
						def.M_cbChangeThree[roomindex][def.GetRoomChair(roomindex,info.UserId)][2]=change3[2]

					}else{
                       //否则没有选择换三张 则直接推荐定缺
						//var cardcount int16
						//if i==int(def.M_wBankerUser[roomindex]){cardcount=14}
						//发送推荐定缺
						if def.M_RoomChair[roomindex][i].UserId!=0 {
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_COLOR_TUIJIAN), &def.CMD_C_SEND_COLOR{
								def.RefDingQue(def.M_cbCardIndex[roomindex][int16(i)], int16(i)),
							})
						}
					}





				}

			}
		}

	case def.SUB_C_CHANGETHREE://用户换三张
		{

			var CHANGE_THREE def.CMD_Rceive_THREE
			var Ch3Chair int
			Ch3Chair=int(def.GetRoomChair(roomindex,info.UserId))
			err := gamelib.UnMarshal(data, &CHANGE_THREE)
			if err != nil {
				fmt.Println(err)

			}

			fmt.Println("接收用户发送过来的换三张",CHANGE_THREE.CbCard[0],CHANGE_THREE.CbCard[1],CHANGE_THREE.CbCard[2])
			//记录换的三张牌;

			//
			def.M_cbChangeThree[roomindex][Ch3Chair][0]=CHANGE_THREE.CbCard[0]
			def.M_cbChangeThree[roomindex][Ch3Chair][1]=CHANGE_THREE.CbCard[1]
			def.M_cbChangeThree[roomindex][Ch3Chair][2]=CHANGE_THREE.CbCard[2]

			def.M_GameStatus[roomindex][Ch3Chair]=def.GAME_STATUS_CHANGETHREE

			//判断是否所有用户已经换三张状态
			var GameChangeThree bool
			GameChangeThree=true
			fmt.Println("换三张状态..",def.M_GameStatus[roomindex][0],def.M_GameStatus[roomindex][1],def.M_GameStatus[roomindex][2],def.M_GameStatus[roomindex][3])
			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{

				if  def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_CHANGETHREE {
					fmt.Println("换三张状态置假",i,"换三张状态应该为",def.GAME_STATUS_CHANGETHREE)
					GameChangeThree = false
				}
			}


			if GameChangeThree {

				var change3[3] int16


				//var cardcount int16
				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{

					fmt.Println(i,"用户   换三张前麻将",def.M_cbCardIndex[roomindex][i])
					change3[0]=def.M_cbChangeThree[roomindex][i][0]
					change3[1]=def.M_cbChangeThree[roomindex][i][1]
					change3[2]=def.M_cbChangeThree[roomindex][i][2]
					var newchange3[3] int16
					newchange3[0]=def.M_cbChangeThree[roomindex][def.GetBeforeChair(roomindex,int16(i))][0]
					newchange3[1]=def.M_cbChangeThree[roomindex][def.GetBeforeChair(roomindex,int16(i))][1]
					newchange3[2]=def.M_cbChangeThree[roomindex][def.GetBeforeChair(roomindex,int16(i))][2]
					def.M_GameRecord_init[roomindex].Change3[i][0]=int(change3[0])
					def.M_GameRecord_init[roomindex].Change3[i][1]=int(change3[1])
					def.M_GameRecord_init[roomindex].Change3[i][2]=int(change3[2])
					def.M_GameRecord_init[roomindex].Change3[i][3]=int(newchange3[0])
					def.M_GameRecord_init[roomindex].Change3[i][4]=int(newchange3[1])
					def.M_GameRecord_init[roomindex].Change3[i][5]=int(newchange3[2])

					def.M_cbCardIndex[roomindex][i][change3[0]]--
					def.M_cbCardIndex[roomindex][i][change3[1]]--
					def.M_cbCardIndex[roomindex][i][change3[2]]--
					def.M_cbCardIndex[roomindex][i][newchange3[0]]++
					def.M_cbCardIndex[roomindex][i][newchange3[1]]++
					def.M_cbCardIndex[roomindex][i][newchange3[2]]++



					fmt.Println("换三张处理结果：用户",i,change3[0],change3[1],change3[2],newchange3[0],newchange3[1],newchange3[2])
					//发送换三张结果;
					if def.M_RoomChair[roomindex][i].UserId!=0{
						lib.mgr.SendGameMessage(&def.M_RoomChair[def.GetRoomid(info.RoomId)][i], uint32(def.SUB_S_CHANGETHREE), &def.CMD_S_CHANGE3{
							CbUser: int16(i),
							Cbcard:change3,
							CbNewcard:newchange3,

						})
					}




					//if i==int(def.M_wBankerUser[roomindex]){cardcount=14}


					//发送推荐定缺
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_COLOR_TUIJIAN), &def.CMD_C_SEND_COLOR{
						def.RefDingQue(def.M_cbCardIndex[roomindex][int16(i)],int16(i)),
					})
					fmt.Println("用户推荐定缺:",def.RefDingQue(def.M_cbCardIndex[roomindex][int16(i)],int16(i)),"用户:",i)

				}




			}




		}
	case def.SUB_C_SEND_COLOR://用户定缺
		{
			//var ReceiveColor int16=0
			var ReceiveColor  def.CMD_C_SEND_COLOR
			err := gamelib.UnMarshal(data, &ReceiveColor)
			if err != nil {
				fmt.Println(err)

			}
            ///发送给所有用户自己的定缺花色

			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
				if def.M_RoomChair[roomindex][i].UserId!=0 {
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_REPLAY_COLOR), &def.CMD_S_ReplayColor{
						CbUser: def.GetRoomChair(roomindex, info.UserId),
						Cbcard: ReceiveColor.CbColorData,
					})
				}

			}

            fmt.Println("发送用户:",def.GetRoomChair(roomindex,info.UserId),"定缺",ReceiveColor.CbColorData)
			def.SetQueColor(int16(ReceiveColor.CbColorData),roomindex,def.GetRoomChair(roomindex,info.UserId))
			def.M_GameStatus[roomindex][def.GetRoomChair(roomindex,info.UserId)]=def.GAME_STATUS_DingQUE

			var GameDingQuestatus bool=true
			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
				def.M_GameRecord_init[roomindex].DingQue[i]=int(def.M_dingqueColor[roomindex][i])

				if  def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_DingQUE {
					GameDingQuestatus = false
				}
			}
			var m_cbUserAction[def.GAME_PLAYER] int16
			if GameDingQuestatus {
				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
					if i==int(def.M_wBankerUser[roomindex]){

						fmt.Println("杠牌分析房间号",roomindex,"麻将：",def.M_cbCardIndex[roomindex][def.M_wBankerUser[roomindex]],"用户",i)
						//m_cbUserAction[i] |=def.AnalyseGangCard(def.GetRoomid(info.RoomId),def.M_cbCardIndex[def.GetRoomid(info.RoomId)][i],int16(i))
						var curaction int16
						var curCard int16
						curaction,curCard=def.AnalyseGangCard(roomindex,int16(i))
						m_cbUserAction[i] |=curaction
						//判断杠
						m_cbUserAction[i] |= def.AnalyseChiHuCard(roomindex,def.M_cbCardIndex[roomindex][i],int16(i))
						//设置变量
						//发送提示
						//fmt.Println("Before 发送提示操作",m_cbUserAction[i])
						def.M_wProvideUser[roomindex]=int16(i)
						if m_cbUserAction[i]!=def.CHK_NULL&&def.M_RoomChair[roomindex][i].UserId!=0{
							fmt.Println("发送提示操作",m_cbUserAction[i],"给用户:",i,"房间号:",def.GetRoomid(info.RoomId),"庄家",def.M_wBankerUser[roomindex],"麻将",curCard)
							//如果有提示则提示
							//if def.M_RoomChair[roomindex][i].UserId!=0{
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_NOTIFY), &def.CMD_S_OperateNotify{

								CbActionMask:m_cbUserAction[i],
								CbActionCard:curCard,

								NotifyID:int16(def.M_NotifyIndex[roomindex]),

							})
							def.M_NotifyIndex[roomindex]++

						}//m_cbUserAction[i]!=def.CHK_NULL
						//如果有需要待执行的动作，则发送提示操作

					}

				}

			}


		}
	case def.SUB_C_OUT_CARD:// 用户出牌
		{


			var receiveCard  def.CMD_C_OutCard
			err := gamelib.UnMarshal(data, &receiveCard)
			if err != nil {
				fmt.Println(err)

			}
			fmt.Println("房间",info.RoomId,"收到用户出牌命令 坐位号",def.GetRoomChair(roomindex,info.UserId),"出牌麻将",int(receiveCard.CbCardData))
            var tmpcardcount int
			tmpcardcount=0
			var ThisChair int
			ThisChair=int(def.GetRoomChair(roomindex,info.UserId))
			for i:=0;i<int(def.MAX_INDEX);i++{

				if def.M_cbCardIndex[roomindex][ThisChair][i]>0{
					tmpcardcount+=int(def.M_cbCardIndex[roomindex][ThisChair][i])
				}
			}
			//for i:=0;i<int(def.M_WeaveItemArray[roomindex][def.GetRoomChair(roomindex,info.UserId)]);i++{

			//}
			fmt.Println("房间",info.RoomId,"房间索引",roomindex,"房间ID：",info.RoomId,"用户ID",info.UserId,"坐位号:",ThisChair,"麻将张数",tmpcardcount,"组合合计",def.M_cbWeaveItemCount[roomindex][ThisChair],"麻将数据:",def.M_cbCardIndex[roomindex][ThisChair],"房间用户信息",def.M_RoomChair[roomindex][0].UserId,def.M_RoomChair[roomindex][1].UserId,def.M_RoomChair[roomindex][2].UserId,def.M_RoomChair[roomindex][3].UserId)
			tmpcardcount=tmpcardcount+int(def.M_cbWeaveItemCount[roomindex][ThisChair])*3

			if tmpcardcount!=14{fmt.Println("房间号",info.RoomId,"用户",ThisChair,"取得的麻将张数出错:",tmpcardcount)}

		    //向所有用户提示用户出牌
            if tmpcardcount==14{
				//def.GetRoomChair(roomindex,info.UserId)
			//fmt.Println("发送广播消息用户出牌",def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId)))

				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
					if def.M_RoomChair[roomindex][i].UserId!=0 {
						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_OUT_CARD), &def.CMD_S_OutCard{
							WOutCardUser:  def.GetRoomChair(roomindex, info.UserId),
							CbOutCardData: receiveCard.CbCardData,
						})
					}
					def.M_NotifyUserReceive[roomindex][i]=false
					def.M_NotifyUserMaxactionReceive[roomindex][i]=false

				}
				def.M_NotifyReceiveMaxAction[roomindex]=int(def.WIK_NULL)
				def.M_NotifyMaxAction[roomindex]=int(def.WIK_NULL)
				//var OperatorIndex int
				//OperatorIndex=def.M_GameRecord_Operator_Index[roomindex]

				def.SaveTheGameRecord(roomindex,80,int(def.GetRoomChair(roomindex,info.UserId)),def.Operator_Play,int(receiveCard.CbCardData))
				/*
				if OperatorIndex<145{
				def.M_GameRecord_Operator[roomindex].GameRecords[OperatorIndex].ProvideUser=80
				def.M_GameRecord_Operator[roomindex].GameRecords[OperatorIndex].Operator_type=1
				def.M_GameRecord_Operator[roomindex].GameRecords[OperatorIndex].Chairid=int(def.GetRoomChair(roomindex,info.UserId))
				def.M_GameRecord_Operator[roomindex].GameRecords[OperatorIndex].Card=int(receiveCard.CbCardData)
				def.M_GameRecord_Operator[roomindex].GameRecords[OperatorIndex].Operatorindex=def.M_GameRecord_Operator_Index[roomindex]
				def.M_GameRecord_Operator_Index[roomindex]++
				}*/
				fmt.Println("房间",info.RoomId,"发送广播消息用户出牌结束",def.GetRoomChair(roomindex,info.UserId),"移除麻将",int(receiveCard.CbCardData))

				//var M_OutCardList[2000][GAME_PLAYER][41] OutCardstruct
				//var M_OutCardListIndex[2000][GAME_PLAYER] int
				//用户打出去的麻将堆
				def.M_OutCardList[roomindex][ThisChair][def.M_OutCardListIndex[roomindex][ThisChair]].ChairID=int(def.GetRoomChair(roomindex,info.UserId))
				def.M_OutCardList[roomindex][ThisChair][def.M_OutCardListIndex[roomindex][ThisChair]].OutCard=int(receiveCard.CbCardData)
				def.M_OutCardListIndex[roomindex][ThisChair]++
				def.SetProvideUser(roomindex,def.GetRoomChair(roomindex,info.UserId))
				fmt.Println("房间:",info.RoomId,"设置供牌用户为",def.GetRoomChair(roomindex,info.UserId),"出牌麻将为:",receiveCard.CbCardData)
				def.M_cbProvideCard[roomindex]=receiveCard.CbCardData
			//def.RemoveCard(def.M_cbCardIndex[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))], receiveCard.CbCardData)
			def.RemoveCard(roomindex,int(def.GetRoomChair(roomindex,info.UserId)),int(receiveCard.CbCardData))

			//fmt.Println("移除麻将",def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId)),"麻将",receiveCard.CbCardData)
			fmt.Println("房间",info.RoomId,"用户",ThisChair,"出牌后  现行麻将数据是",def.M_cbCardIndex[roomindex][def.GetRoomChair(roomindex,info.UserId)])
			for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{def.M_GameStatus[roomindex][def.GetRoomChair(roomindex,info.UserId)]=def.GAME_STATUS_PLAY}

			def.SetUserAction(roomindex,def.WIK_NULL,def.GetRoomChair(roomindex,info.UserId))
			def.SetPerformAction(def.GetRoomid(info.RoomId),def.WIK_NULL,def.GetRoomChair(roomindex,info.UserId))




			//用户切换
			var m_wCurrentUser int16=def.GetNextChair(roomindex, def.GetRoomChair(roomindex,info.UserId))

			//响应判断
			def.M_NotifyMaxAction[roomindex]=int(def.WIK_NULL)
			var bAroseAction bool=def.EstimateUserRespond(roomindex,def.GetRoomChair(roomindex,info.UserId),receiveCard.CbCardData)
            //如果有需要提示的用户，则向用户发送提示操作;
				if bAroseAction{
					//def.M_NotifyUsers[roomindex]=0
					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
						def.M_NotifyUsers[roomindex][i]=false
						def.M_cbPerformAction[roomindex][i]=def.WIK_NULL
						if def.M_cbUserAction[roomindex][i]>=def.WIK_PENG&&int(def.M_cbUserAction[roomindex][i])>=def.M_NotifyMaxAction[roomindex]{
							def.M_NotifyMaxAction[roomindex]=int(def.WIK_PENG)
						}
						if def.M_cbUserAction[roomindex][i]>=def.WIK_XiaYu&&int(def.M_cbUserAction[roomindex][i])>=def.M_NotifyMaxAction[roomindex]{
							def.M_NotifyMaxAction[roomindex]=int(def.WIK_XiaYu)
						}
						if def.M_cbUserAction[roomindex][i]>=def.WIK_CHI_HU&&int(def.M_cbUserAction[roomindex][i])>=def.M_NotifyMaxAction[roomindex]{
							def.M_NotifyMaxAction[roomindex]=int(def.WIK_CHI_HU)
						}

						//清空用户执行权限
						if def.M_cbUserAction[roomindex][i]!=def.WIK_NULL&&def.M_RoomChair[roomindex][i].UserId!=0{

							fmt.Println("用户出牌后发送给",i,"用户 出牌用户为:",def.GetRoomChair(roomindex,info.UserId)," 提示命令为",def.M_cbUserAction[roomindex][i],"提示","麻将数据为",def.M_cbCardIndex[roomindex][i])
							//if def.M_RoomChair[roomindex][i].UserId!=0{
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_NOTIFY), &def.CMD_S_OperateNotify{

								CbActionMask:def.M_cbUserAction[roomindex][i],
								CbActionCard:int16(receiveCard.CbCardData),


								NotifyID:int16(def.M_NotifyIndex[roomindex]),

							})
							//def.M_NotifyUsers[roomindex]++
							def.M_NotifyUsers[roomindex][i]=true

						}
					}
					def.M_NotifyIndex[roomindex]++

				}

				//派发扑克
				if !bAroseAction {

					var cardpai int16
					var curaction int16
					var gangpai int16
					if def.M_cbPosition[roomindex] == def.MAX_REPERTORY-1 {
						//

						GameLiuJU(roomindex)

                        //保存回放数据
						type SendRecord struct {
							Inithead def.GameRecord_init
							Operate def.GameRecord_Operator
							GameConclude def.CMD_S_GameConclude
						}
						var BsendRec SendRecord
						BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
						BsendRec.Inithead=def.M_GameRecord_init[roomindex]
						BsendRec.Operate=def.M_GameRecord_Operator[roomindex]

						type Headstruct struct {
							Roomid int
							Draw int
							Userlst[def.GAME_PLAYER] string
							Gamescore[def.GAME_PLAYER] int
							Gamekind int
							Tm 		time.Time

						}
						var heads Headstruct
						heads.Roomid=int(info.RoomId)
						//heads.Draw=int(def.M_GameRecordDraw[roomindex])
						heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
						fmt.Println("游戏局数:",heads.Draw)
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
							heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
						}
						heads.Tm = time.Now()
						heads.Gamekind=1

						header, _ := gamelib.Marshal(&heads)
						content, _ := gamelib.Marshal(&BsendRec)

						var id int
						id=lib.mgr.SaveGameRecord(header, content)
						for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
							//lib.mgr.SaveUserRecord(int(info.UserId), id)
							lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)

						}




						//提示所有玩家游戏结束
						fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
						for tmpi := 0; tmpi < int(def.M_DesktopPlayer[roomindex]); tmpi++ {
							if def.M_RoomChair[roomindex][tmpi].UserId!=0{
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
								LGameScore:      def.M_GameConcludeScore[roomindex].LGameScore,
								LRevenue:        def.M_GameConcludeScore[roomindex].LRevenue,
								LGangScore:      def.M_GameConcludeScore[roomindex].LGangScore,
								WProvideUser:    def.M_GameConcludeScore[roomindex].WProvideUser,
								CbProvideCard:   def.M_GameConcludeScore[roomindex].CbProvideCard,
								CbHandCardData:  def.M_GameConcludeScore[roomindex].CbHandCardData,
								CbGenCount:      def.M_GameConcludeScore[roomindex].CbGenCount,
								WLiuJuStatus:    def.M_GameConcludeScore[roomindex].WLiuJuStatus,
								UFanDescBase:    def.M_GameConcludeScore[roomindex].UFanDescBase,
								UFanDescAddtion: def.M_GameConcludeScore[roomindex].UFanDescAddtion,
								UFanBase:        def.M_GameConcludeScore[roomindex].UFanBase,
								UFanAddtion:     def.M_GameConcludeScore[roomindex].UFanAddtion,
							})
							}
							def.M_GameStatus[roomindex][tmpi] = def.GAME_STATUS_WINED
						} //For

						//几局的统计结算界面
						if def.M_GameRecordDraw[roomindex] >= def.M_GameRoomsDrawinit[roomindex] {
							fmt.Println("初始化局数", def.M_GameRoomsDrawinit[roomindex])
							for tmpi := 0; tmpi < int(def.M_DesktopPlayer[roomindex]); tmpi++ {
								if def.M_RoomChair[roomindex][tmpi].UserId!=0 {
									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
										LGameScore: def.M_GameTotalScoreCount[roomindex].LGameScore,
										LZhiMo:     def.M_GameTotalScoreCount[roomindex].LZhiMo,
										LJiePao:    def.M_GameTotalScoreCount[roomindex].LJiePao,
										LDianPao:   def.M_GameTotalScoreCount[roomindex].LDianPao,
										LAnGang:    def.M_GameTotalScoreCount[roomindex].LAnGang,
										LMingGang:  def.M_GameTotalScoreCount[roomindex].LMingGang,
										LChaDaJia:  def.M_GameTotalScoreCount[roomindex].LChaDaJia,
										LRoomCard:  def.M_GameTotalScoreCount[roomindex].LRoomCard,

									})
								}

							}
						}

					}
						if m_wCurrentUser < 4 && def.M_cbPosition[roomindex] < (def.MAX_REPERTORY-1) {


							cardpai, curaction, gangpai = def.DispatchCardData(roomindex, m_wCurrentUser)
							if curaction != def.WIK_NULL {
								fmt.Println("发牌给", m_wCurrentUser, "用户", m_wCurrentUser, "麻将", cardpai, "杠牌", gangpai, "操作", curaction)
							}
							if def.M_RoomChair[roomindex][m_wCurrentUser].UserId!=0 {
								//给用户发牌;
								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][m_wCurrentUser], uint32(def.SUB_S_SEND_CARD), &def.CMD_S_SendCard{

									CbCardData:    cardpai,
									CbActionMask:  curaction,
									WCurrentUser:  m_wCurrentUser,
									WSendCardUser: m_wCurrentUser,
									CbGanData:     gangpai,
								})
							}
							//给告诉所有用户给XX用户发牌;

							fmt.Println("房间",info.RoomId,"通知所有人发牌给", m_wCurrentUser, "用户")
							for i := 0; i < int(def.M_DesktopPlayer[roomindex]); i++ {
								if (def.M_RoomChair[roomindex][i].UserId!=0)&&(i!=int(m_wCurrentUser)) {
									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i], uint32(def.SUB_S_SEND_CARD_BroadCast), &def.CMD_S_SendCard_Broadcast{

										WCurrentUser: m_wCurrentUser,
									})
								}

							}
						} else {
							fmt.Println("房间",info.RoomId,"派发麻将出错不应出现的用户", m_wCurrentUser)
						}

				}

				//发送扑克

			}

			//return true


		}
	case def.SUB_C_GetCardlist:
		{
			type Cardlist struct{
			 User0Card[def.MAX_INDEX] int
			 User1Card[def.MAX_INDEX] int
			 User2Card[def.MAX_INDEX] int
			 User3Card[def.MAX_INDEX] int
			 AllCardLst[def.MAX_REPERTORY] int
			 Position int
		     }
			var SendallCardlst[def.MAX_REPERTORY] int
			var U0card[def.MAX_INDEX] int
			var U1card[def.MAX_INDEX] int
			var U2card[def.MAX_INDEX] int
			var U3card[def.MAX_INDEX] int

			for i:=0;i<int(def.MAX_REPERTORY);i++{
				SendallCardlst[i]=int(def.M_cbCardData[roomindex][i])
			}
			for i:=0;i<int(def.MAX_INDEX);i++{
				U0card[i]=int(def.M_cbCardIndex[roomindex][0][i])
				U1card[i]=int(def.M_cbCardIndex[roomindex][1][i])
				U2card[i]=int(def.M_cbCardIndex[roomindex][2][i])
				U3card[i]=int(def.M_cbCardIndex[roomindex][3][i])
			}

			lib.mgr.SendGameMessage(info,uint32(def.SUB_S_SendCardList), &Cardlist{
				User0Card:U0card,
				User1Card:U1card,
				User2Card:U2card,
				User3Card:U3card,
				AllCardLst:SendallCardlst,
				Position:int(def.M_cbPosition[roomindex]),

			})

		}
	case def.SUB_C_OPERATE_CARD://用户操作牌:
		{
			var receiveOPerate  def.CMD_C_OperateCard
			err := gamelib.UnMarshal(data, &receiveOPerate)
			if err != nil {
				fmt.Println(err)

			}
			var ThisChair int

			ThisChair=int(def.GetRoomChair(roomindex,info.UserId))
			m_NotifyType:=false //接收的是本人模牌的操作 还是其他人打出来的牌的操作
			if def.M_wProvideUser[roomindex]==int16(ThisChair){m_NotifyType=true}
			fmt.Println("收到用户操作麻将的信息 操作",receiveOPerate.CbOperateCode,"麻将",receiveOPerate.CbOperateCard,"用户",def.GetRoomChair(def.GetRoomid(info.RoomId),info.UserId),"操作人类型",m_NotifyType)
			if def.M_cbUserAction[roomindex][ThisChair]!=def.WIK_NULL{
				def.M_cbPerformAction[roomindex][ThisChair] = receiveOPerate.CbOperateCode
			}else{def.M_cbPerformAction[roomindex][ThisChair]=def.WIK_NULL}

			//如果用户选择过则跳过本操作
			//if receiveOPerate.CbOperateCode==def.WIK_NULL {def.M_cbUserAction[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]=def.WIK_NULL}
			fmt.Println("用户",ThisChair,"修改前用户操作",def.M_cbUserAction[roomindex][ThisChair])

			needSendCard:=false
			needSendCard=false
			var SendCardUser int16
			//解决抢杠胡不能过的问题
			//如果抢杠胡 提示 玩家点了过则给杠牌玩家发牌。
			if def.M_cbUserAction[roomindex][ThisChair]==def.WIK_CHI_HU&&receiveOPerate.CbOperateCode==def.WIK_NULL&&def.BeLastQiangGang[roomindex][def.M_wProvideUser[roomindex]]==1{//抢杠胡时点了过给杠牌玩家发牌
				needSendCard=true
				def.BeLastQiangGang[roomindex][def.M_wProvideUser[roomindex]]=0
				SendCardUser=def.M_wProvideUser[roomindex]
			}
			/*
			if def.M_cbUserAction[roomindex][ThisChair]!=def.WIK_NULL{
			def.M_cbUserAction[roomindex][ThisChair]=receiveOPerate.CbOperateCode
			}*/
            //if def.M_cbUserAction[roomindex][def.GetRoomChair(roomindex,info.UserId)]

			fmt.Println("用户",ThisChair,"修改后用户操作",def.M_cbUserAction[roomindex][ThisChair])


			if m_NotifyType {
				var ThisChair int
                ThisChair=int(def.GetRoomChair(roomindex,info.UserId))
				fmt.Println("保存刮风下雨结果")
				var GuaFengXiaYu=false
				var GangScore[def.GAME_PLAYER] int16
				if receiveOPerate.CbOperateCode==def.WIK_XiaYu{GuaFengXiaYu=true}
				for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
					if (i!=ThisChair)&&(def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_WINED){
						if receiveOPerate.CbOperateCode==def.WIK_XiaYu{
							GangScore[i]=2*def.M_initGameDrawScore[roomindex]
						}
						if receiveOPerate.CbOperateCode==def.WIK_GuaFeng{
							GangScore[i]=def.M_initGameDrawScore[roomindex]
						}
					}

				}
				fmt.Println("保存刮风下雨结果完成")



				switch receiveOPerate.CbOperateCode {
				case def.WIK_CHI_HU:{
					var fans[2] int16
					//如果本人为第一个赢牌的玩家，则本人当庄
					if def.WinOrder[roomindex]==int16(0){
						def.M_NextwBankerUser[roomindex]=ThisChair
					}
					def.M_GameConcludeScore[roomindex].UFanDescBase[ThisChair]=""
					def.M_GameConcludeScore[roomindex].UFanDescAddtion[ThisChair]=""
					var CurJiFen int16

					fans[0],fans[1],CurJiFen=def.UserHuPai(roomindex,int16(ThisChair),int(receiveOPerate.CbOperateCard),true)
					fmt.Println("原来所有玩家的积分",def.M_GameConcludeScore[roomindex].LGameScore)
					def.M_GameConcludeScore[roomindex].HuType[ThisChair]=def.HuType_ZhiMo
					for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
						if (tmpi!=ThisChair)&&(def.M_GameStatus[roomindex][tmpi]!=def.GAME_STATUS_WINED){
							def.M_GameConcludeScore[roomindex].LGameScore[ThisChair]=def.M_GameConcludeScore[roomindex].LGameScore[ThisChair]+int(CurJiFen)
							def.M_GameConcludeScore[roomindex].LGameScore[tmpi]=def.M_GameConcludeScore[roomindex].LGameScore[tmpi]-int(CurJiFen)
						}
					}
					//fmt.Println("积分：",CurJiFen,"所有玩家的积分",def.M_GameConcludeScore[roomindex].LGameScore,"得分玩家",ThisChair,"出分玩家",def.M_wProvideUser[roomindex])
					//def.M_GameConcludeScore[roomindex].LGameScore[ThisChair]=def.M_GameConcludeScore[roomindex].LGameScore[ThisChair]+int(CurJiFen)
					//def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]=def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]-int(CurJiFen)
					fmt.Println("积分：",CurJiFen,"所有玩家的积分",def.M_GameConcludeScore[roomindex].LGameScore,"得分玩家",ThisChair,"出分玩家",def.M_wProvideUser[roomindex])
					//给下一玩家发牌
					SendCardUser =def.GetNextChair(roomindex,int16(ThisChair))

					fmt.Println("向所有用户发送操作结果",receiveOPerate.CbOperateCode,"是否需要发牌",needSendCard)
					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
						if def.M_RoomChair[roomindex][i].UserId!=0{



						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_RESULT), &def.CMD_S_OperateResult{
							WOperateUser:SendCardUser,
							CbActionMask:receiveOPerate.CbOperateCode,
							WProvideUser:SendCardUser,
							CbOperateCode:receiveOPerate.CbOperateCode,
							CbOperateCard:receiveOPerate.CbOperateCard,
							CbGuaFengXiaYu:GuaFengXiaYu,
							GangScore:GangScore,
						})
						}

					}





					for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++ {
						def.M_GameConcludeScore[roomindex].LGangScore[tmpi]=0
						for tmpj := 0; tmpj < int(def.M_DesktopPlayer[roomindex]); tmpj++ {
							def.M_GameConcludeScore[roomindex].LGangScore[tmpi] = def.M_GameConcludeScore[roomindex].LGangScore[tmpi] + def.M_GangScore[roomindex][tmpi][tmpj]
						}
					}
                    //点炮的用户
					def.M_GameConcludeScore[roomindex].WProvideUser[ThisChair]=def.M_wProvideUser[roomindex]
					//被点炮的麻将
					def.M_GameConcludeScore[roomindex].CbProvideCard[ThisChair]=receiveOPerate.CbOperateCard

					//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].CbHandCardData[def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]=def.M_cbCardIndex[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]
					def.M_GameConcludeScore[roomindex].WLiuJuStatus[ThisChair]=4
					def.M_GameConcludeScore[roomindex].UFanBase[ThisChair]=fans[0]
					def.M_GameConcludeScore[roomindex].UFanAddtion[ThisChair]=fans[1]

					def.M_GameTotalScoreCount[roomindex].LZhiMo[ThisChair]++
					def.M_GameStatus[roomindex][ThisChair]=def.GAME_STATUS_WINED

					fmt.Println("摸牌模式吃胡(即自摸) 番数",fans,"供牌用户",def.M_wProvideUser[roomindex],"取胜用户",def.GetRoomChair(roomindex,info.UserId),"CbWinOrder",def.WinOrder[roomindex],"基本番描述",def.M_GameConcludeScore[roomindex].UFanDescBase[def.GetRoomChair(roomindex,info.UserId)],"取胜用户模式2",ThisChair,"额外番描述",def.M_GameConcludeScore[roomindex].UFanDescAddtion[ThisChair],"状态",def.M_GameStatus[roomindex][ThisChair],"游戏积分:",def.M_GameConcludeScore[roomindex].LGameScore[ThisChair],"供牌用户积分",def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]],"所有积分",def.M_GameConcludeScore[roomindex].LGameScore,"所有杠",def.M_GangScore[roomindex][0],def.M_GangScore[roomindex][1],def.M_GangScore[roomindex][2],def.M_GangScore[roomindex][3],"胡到的麻将为",def.M_GameConcludeScore[roomindex].CbProvideCard[ThisChair],"所有人的胡牌麻将为",def.M_GameConcludeScore[roomindex].CbProvideCard)
					def.WinOrder[roomindex]++

					//保存流回放数据
					def.SaveTheGameRecord(roomindex,ThisChair,ThisChair,def.Operator_Hu,int(receiveOPerate.CbOperateCard))
					/*
					if def.M_GameRecord_Operator_Index[roomindex]<149{
					def.M_GameRecord_Operator[roomindex].GameRecords[def.M_GameRecord_Operator_Index[roomindex]].ProvideUser=ThisChair
					def.M_GameRecord_Operator[roomindex].GameRecords[def.M_GameRecord_Operator_Index[roomindex]].Operator_type=5
					def.M_GameRecord_Operator[roomindex].GameRecords[def.M_GameRecord_Operator_Index[roomindex]].Chairid=ThisChair
					def.M_GameRecord_Operator[roomindex].GameRecords[def.M_GameRecord_Operator_Index[roomindex]].Card=int(receiveOPerate.CbOperateCard)
					def.M_GameRecord_Operator[roomindex].GameRecords[def.M_GameRecord_Operator_Index[roomindex]].Operatorindex=def.M_GameRecord_Operator_Index[roomindex]
					def.M_GameRecord_Operator_Index[roomindex]++
					}*/


					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][ThisChair],uint32(def.SUB_S_HU), &def.CMD_S_ChiHu{

						WChiHuUser:int16(ThisChair),
						WProviderUser:def.M_wProvideUser[roomindex],
						CbChiHuCard:def.M_cbProvideCard[roomindex],
						LGameScore:int16(def.M_GameConcludeScore[roomindex].LGameScore[ThisChair]),
						CbWinOrder:def.WinOrder[roomindex],
					})


					needSendCard =true
                    //如果打到最后只有一名玩家了。游戏结束 显示游戏得分
					if def.GetPlayIngUserCount(roomindex)<2{
						needSendCard=false
						//提示所有玩家游戏结束
						def.GameRecordDrawScore(roomindex)
						def.M_GameConcludeScore[roomindex].Tm=time.Now()
						fmt.Println("自摸模式下杠分",def.M_GameConcludeScore[roomindex].LGangScore, def.M_GameConcludeScore[roomindex].Tm,"胡牌类型", def.M_GameConcludeScore[roomindex].HuType)
						//存至川江服务器 回放的数据
						type SendRecord struct {
							Inithead def.GameRecord_init
							Operate def.GameRecord_Operator
							GameConclude def.CMD_S_GameConclude
						}
						//
						var BsendRec SendRecord
                        BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
						BsendRec.Inithead=def.M_GameRecord_init[roomindex]
						BsendRec.Operate=def.M_GameRecord_Operator[roomindex]
						fmt.Println("传递的数据为",BsendRec)
						//gamelib.Marshal(BsendRec)
						//gamelib.Marshal(info.RoomId)
						type Headstruct struct {
							Roomid int
							Draw int
							Userlst[def.GAME_PLAYER] string
							Gamescore[def.GAME_PLAYER] int
							Gamekind int
                            Tm 		time.Time

						}
						var heads Headstruct
                        heads.Roomid=int(info.RoomId)
						//heads.Draw=int(def.M_GameRecordDraw[roomindex])
						heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
						fmt.Println("游戏局数:",heads.Draw)
						//heads.draw=
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
							heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
						}
                        heads.Tm = time.Now()
						heads.Gamekind=1

						header, _ := gamelib.Marshal(&heads)
						content, _ := gamelib.Marshal(&BsendRec)

                        var id int
						id=lib.mgr.SaveGameRecord(header, content)
						//lib.mgr.SaveUserRecord(int(info.UserId),id)


						for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
							//lib.mgr.SaveUserRecord(int(info.UserId), id)
							lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
						}


							//fmt.Println("各玩家手中麻将为",def.M_GameConcludeScore[roomindex].CbHandCardData[0]," ",def.M_GameConcludeScore[roomindex].CbHandCardData[1])
						fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
						for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
							LGameScore:def.M_GameConcludeScore[roomindex].LGameScore,
							LRevenue:def.M_GameConcludeScore[roomindex].LRevenue,
							LGangScore:def.M_GameConcludeScore[roomindex].LGangScore,
							WProvideUser:def.M_GameConcludeScore[roomindex].WProvideUser,
							CbProvideCard:def.M_GameConcludeScore[roomindex].CbProvideCard,
							CbHandCardData:def.M_GameConcludeScore[roomindex].CbHandCardData,
							CbGenCount:def.M_GameConcludeScore[roomindex].CbGenCount,
							WLiuJuStatus:def.M_GameConcludeScore[roomindex].WLiuJuStatus,
							UFanDescBase:def.M_GameConcludeScore[roomindex].UFanDescBase,
							UFanDescAddtion:def.M_GameConcludeScore[roomindex].UFanDescAddtion,
							UFanBase:def.M_GameConcludeScore[roomindex].UFanBase,
							UFanAddtion:def.M_GameConcludeScore[roomindex].UFanAddtion,
							HuType:def.M_GameConcludeScore[roomindex].HuType,
							Tm:def.M_GameConcludeScore[roomindex].Tm,
						})
							def.M_GameStatus[roomindex][tmpi]=def.GAME_STATUS_WINED
							def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi]=int16(def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGangScore[tmpi])

						}//For

						fmt.Println("总计游戏积分为：",def.M_GameTotalScoreCount[roomindex].LGameScore[0],def.M_GameTotalScoreCount[roomindex].LGameScore[1],def.M_GameTotalScoreCount[roomindex].LGameScore[2],def.M_GameTotalScoreCount[roomindex].LGameScore[3])

						//几局的统计结算界面
						if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
							fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)
							for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
									LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
									LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
									LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
									LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
									LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
									LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
									LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
									LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

								})

							}
							lib.mgr.ReleaseRoom()
						}

					}


					//发送扑克
				}

				case def.WIK_GuaFeng:{
					var beQiangGang bool
					beQiangGang=false
					//判断是否用户可以抢杠胡牌
					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++ {
					if i!=int(def.GetRoomChair(roomindex,info.UserId)){
						def.M_cbUserAction[roomindex][i]=def.WIK_NULL

						def.M_cbUserAction[roomindex][i]=def.AnalyseChiHuCard_UserSendCard(roomindex,def.M_cbCardIndex[roomindex][int16(i)],int16(i),receiveOPerate.CbOperateCard)
							if def.M_cbUserAction[roomindex][i]!=def.WIK_NULL{
								beQiangGang=true
								def.BeLastQiangGang[roomindex][ThisChair]=1
								def.M_wProvideUser[roomindex]=int16(ThisChair)
								fmt.Println("抢杠胡   供牌用户",ThisChair,"胡牌用户",i,"抢杠的牌:",receiveOPerate.CbOperateCard)
								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_NOTIFY), &def.CMD_S_OperateNotify{
									CbActionMask:def.M_cbUserAction[roomindex][i],
									CbActionCard:int16(receiveOPerate.CbOperateCard),
									NotifyID:int16(def.M_NotifyIndex[roomindex]),

								})
								def.M_NotifyIndex[roomindex]++


							}
						}
					}
					def.RemoveCard(roomindex,ThisChair,int(receiveOPerate.CbOperateCard))
					//def.RemoveCard(def.M_cbCardIndex[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))],receiveOPerate.CbOperateCard)
					//beGang=true
                    if !beQiangGang{
                        //有多少用户被刮风
						var GameBeGangUsers int16
						GameBeGangUsers=0
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++ {

							if def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_WINED&&(i!=ThisChair){
								GameBeGangUsers++


								//fmt.Println("房间号：",info.RoomId,"roomindex",roomindex,"组合合计：",def.M_cbWeaveItemCount[roomindex][ThisChair],"被计数用户",i,"用户值",def.M_WeaveItemArray[roomindex][ThisChair][def.M_cbWeaveItemCount[roomindex][ThisChair]].GangUserList[i],"刮风用户",ThisChair,"所有用户杠状态",def.M_WeaveItemArray[roomindex][ThisChair][def.M_cbWeaveItemCount[roomindex][ThisChair]].GangUserList)
							}
						}

						//GameBeGangUsers--
						for m:=0;m<int(def.M_cbWeaveItemCount[roomindex][ThisChair]);m++{
							if def.M_WeaveItemArray[roomindex][ThisChair][m].CbWeaveKind==def.WIK_PENG&&def.M_WeaveItemArray[roomindex][ThisChair][m].CbCenterCard==receiveOPerate.CbOperateCard{
								def.M_WeaveItemArray[roomindex][ThisChair][m].CbWeaveKind=def.WIK_GuaFeng
								def.M_WeaveItemArray[roomindex][ThisChair][m].WProvideUser=int16(ThisChair)
								def.M_WeaveItemArray[roomindex][ThisChair][m].GangSangPao=false
								def.M_WeaveItemArray[roomindex][ThisChair][m].BeGangUserCount=GameBeGangUsers
								for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++ {
									if i!=ThisChair&&def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_WINED{
										def.M_WeaveItemArray[roomindex][ThisChair][m].GangUserList[i]=1
									}

								}



							}
						}
						fmt.Println("摸牌模式，刮风 用户为:",ThisChair,"刮风麻将为：",receiveOPerate.CbOperateCard,"刮风",GameBeGangUsers,"家")

						var iGuaFenJiFen int16
						iGuaFenJiFen=0

						for m:=0;m<int(def.M_DesktopPlayer[roomindex]);m++{

							if def.GetRoomChair(roomindex,info.UserId)!=int16(m)&&def.M_GameStatus[roomindex][m]!=def.GAME_STATUS_WINED{
								def.M_GangScore[roomindex][m][ThisChair]-=def.M_initGameDrawScore[roomindex]
								def.M_GangScore[roomindex][ThisChair][m]+=def.M_initGameDrawScore[roomindex]
								iGuaFenJiFen++
							}

							def.BeLastisGang[roomindex][m]=def.WIK_NULL

						}//计算刮风所得
						fmt.Println("杠操作：模牌模式玩家 ",ThisChair,"获得刮风积分:",iGuaFenJiFen,"刮风麻将为:",receiveOPerate.CbOperateCard,"刮风家数",GameBeGangUsers,"各玩家杠分情况",def.M_GangScore[roomindex][0],def.M_GangScore[roomindex][1],def.M_GangScore[roomindex][2],def.M_GangScore[roomindex][3])
						needSendCard=true

						def.BeLastisGang[roomindex][ThisChair]=def.WIK_GuaFeng


						SendCardUser=int16(ThisChair)
						fmt.Println("向所有用户发送操作结果",receiveOPerate.CbOperateCode)

						def.SaveTheGameRecord(roomindex,ThisChair,ThisChair,def.Operator_GuaFeng,int(receiveOPerate.CbOperateCard))

						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_RESULT), &def.CMD_S_OperateResult{
								WOperateUser:SendCardUser,
								CbActionMask:receiveOPerate.CbOperateCode,
								WProvideUser:SendCardUser,
								CbOperateCode:receiveOPerate.CbOperateCode,
								CbOperateCard:receiveOPerate.CbOperateCard,
								CbGuaFengXiaYu:GuaFengXiaYu,
								GangScore:GangScore,
								IsQiangGang:false,
							})

						}
						def.M_GameTotalScoreCount[roomindex].LMingGang[SendCardUser]++
						//明杠增加
					}else{//否则被人抢杠则 显示发送给所有玩家被人抢杠成功
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_RESULT), &def.CMD_S_OperateResult{
								WOperateUser:int16(ThisChair),
								CbActionMask:receiveOPerate.CbOperateCode,
								WProvideUser:int16(ThisChair),
								CbOperateCode:receiveOPerate.CbOperateCode,
								CbOperateCard:receiveOPerate.CbOperateCard,
								CbGuaFengXiaYu:GuaFengXiaYu,
								GangScore:GangScore,
								IsQiangGang:true,
							})

						}
					}


				}

				case def.WIK_XiaYu:{
					fmt.Println("移除麻将",receiveOPerate.CbOperateCard)
					def.RemoveCard(roomindex,ThisChair,int(receiveOPerate.CbOperateCard))
					def.RemoveCard(roomindex,ThisChair,int(receiveOPerate.CbOperateCard))
					def.RemoveCard(roomindex,ThisChair,int(receiveOPerate.CbOperateCard))
					def.RemoveCard(roomindex,ThisChair,int(receiveOPerate.CbOperateCard))

					fmt.Println("麻将组合牌增加")
					//有多少用户被刮风
					var GameBeGangUsers int16
					var ChairWeaveCount int
					GameBeGangUsers=0
					ChairWeaveCount=0
					ChairWeaveCount=int(def.M_cbWeaveItemCount[roomindex][ThisChair])
					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++ {
						def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].GangUserList[i]=0
						if (def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_WINED)&&(i!=ThisChair){
							GameBeGangUsers++
							def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].GangUserList[i]=1
							fmt.Println("房间号：",info.RoomId,"roomindex",roomindex,"组合合计：",ChairWeaveCount,"被计数用户",i,"用户值",def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].GangUserList[i],"下雨用户",ThisChair,"所有用户杠状态",def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].GangUserList)
						}
					}
					//GameBeGangUsers--
					def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].CbWeaveKind =def.WIK_XiaYu
					def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].CbCenterCard = receiveOPerate.CbOperateCard
					def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].WProvideUser = int16(ThisChair)
					def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].BeGangUserCount=GameBeGangUsers
					def.M_WeaveItemArray[roomindex][ThisChair][ChairWeaveCount].GangSangPao=false


					fmt.Println("摸牌模式，下雨 用户为:",ThisChair)
					def.M_cbWeaveItemCount[roomindex][def.GetRoomChair(roomindex,info.UserId)]++
					var IxiaYuJiFen int16
					IxiaYuJiFen=0
					for m:=0;m<int(def.M_DesktopPlayer[roomindex]);m++{

						if def.GetRoomChair(roomindex,info.UserId)!=int16(m)&&def.M_GameStatus[roomindex][m]!=def.GAME_STATUS_WINED{
							def.M_GangScore[roomindex][m][ThisChair]-=2*def.M_initGameDrawScore[roomindex]
							def.M_GangScore[roomindex][ThisChair][m]+=2*def.M_initGameDrawScore[roomindex]
							IxiaYuJiFen=IxiaYuJiFen+2
						}
						def.BeLastisGang[roomindex][m]=def.WIK_NULL
					}//计算刮风所得
					//fmt.Println("记最后操作为杠")
					def.BeLastisGang[roomindex][ThisChair]=def.WIK_XiaYu
					needSendCard=true
					SendCardUser=int16(ThisChair)
					fmt.Println("向所有用户发送操作结果",receiveOPerate.CbOperateCode,"模牌模式玩家",ThisChair,"获得下雨积分:",IxiaYuJiFen,"各玩家杠分情况",def.M_GangScore[roomindex][0],def.M_GangScore[roomindex][1],def.M_GangScore[roomindex][2],def.M_GangScore[roomindex][3])
					//保存下雨结果

					def.SaveTheGameRecord(roomindex,ThisChair,ThisChair,def.Operator_XiaYu,int(receiveOPerate.CbOperateCard))

					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][i],uint32(def.SUB_S_OPERATE_RESULT), &def.CMD_S_OperateResult{
							WOperateUser:SendCardUser,
							CbActionMask:receiveOPerate.CbOperateCode,
							WProvideUser:SendCardUser,
							CbOperateCode:receiveOPerate.CbOperateCode,
							CbOperateCard:receiveOPerate.CbOperateCard,
							CbGuaFengXiaYu:GuaFengXiaYu,
							GangScore:GangScore,
							IsQiangGang:false,
						})

					}
					//增加暗杠
					def.M_GameTotalScoreCount[roomindex].LAnGang[SendCardUser]++


				}

				case def.WIK_NULL:{

				}
				}


				if needSendCard{
					var cardpai int16
					var curaction int16
					var gangPai int16
					fmt.Println("needSendCard......")
					if def.M_cbPosition[roomindex]==def.MAX_REPERTORY-1 {

						GameLiuJU(roomindex)

						type SendRecord struct {
							Inithead def.GameRecord_init
							Operate def.GameRecord_Operator
							GameConclude def.CMD_S_GameConclude
						}
						var BsendRec SendRecord
						BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
						BsendRec.Inithead=def.M_GameRecord_init[roomindex]
						BsendRec.Operate=def.M_GameRecord_Operator[roomindex]

						type Headstruct struct {
							Roomid int
							Draw int
							Userlst[def.GAME_PLAYER] string
							Gamescore[def.GAME_PLAYER] int
							Gamekind int
							Tm 		time.Time

						}
						var heads Headstruct
						heads.Roomid=int(info.RoomId)
						//heads.Draw=int(def.M_GameRecordDraw[roomindex])
						heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
						fmt.Println("游戏局数:",heads.Draw)
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
							heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
						}
						heads.Tm = time.Now()
						heads.Gamekind=1

						header, _ := gamelib.Marshal(&heads)
						content, _ := gamelib.Marshal(&BsendRec)

						var id int
						id=lib.mgr.SaveGameRecord(header, content)
						//lib.mgr.SaveUserRecord(int(info.UserId),id)

						for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
							//lib.mgr.SaveUserRecord(int(info.UserId), id)
							lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
						}

						fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
						//提示所有玩家游戏结束
						for tmpi := 0; tmpi < int(def.M_DesktopPlayer[roomindex]); tmpi++ {
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
								LGameScore:      def.M_GameConcludeScore[roomindex].LGameScore,
								LRevenue:        def.M_GameConcludeScore[roomindex].LRevenue,
								LGangScore:      def.M_GameConcludeScore[roomindex].LGangScore,
								WProvideUser:    def.M_GameConcludeScore[roomindex].WProvideUser,
								CbProvideCard:   def.M_GameConcludeScore[roomindex].CbProvideCard,
								CbHandCardData:  def.M_GameConcludeScore[roomindex].CbHandCardData,
								CbGenCount:      def.M_GameConcludeScore[roomindex].CbGenCount,
								WLiuJuStatus:    def.M_GameConcludeScore[roomindex].WLiuJuStatus,
								UFanDescBase:    def.M_GameConcludeScore[roomindex].UFanDescBase,
								UFanDescAddtion: def.M_GameConcludeScore[roomindex].UFanDescAddtion,
								UFanBase:        def.M_GameConcludeScore[roomindex].UFanBase,
								UFanAddtion:     def.M_GameConcludeScore[roomindex].UFanAddtion,
							})
							def.M_GameStatus[roomindex][tmpi] = def.GAME_STATUS_WINED
						} //For

						//几局的统计结算界面
						if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
							fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)

							for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
									LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
									LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
									LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
									LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
									LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
									LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
									LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
									LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

								})

							}
						}
					}//

                    if def.M_cbPosition[roomindex]<(def.MAX_REPERTORY-1){
					cardpai,curaction,gangPai=def.DispatchCardData(roomindex,SendCardUser)

                    fmt.Println("操作后给",SendCardUser,"号用户发牌;","杠牌",gangPai,";提示有操作:",curaction,"新发的牌",cardpai,"总共的牌",def.M_cbCardIndex[roomindex][def.GetRoomChair(roomindex,info.UserId)])
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][SendCardUser],uint32(def.SUB_S_SEND_CARD), &def.CMD_S_SendCard{
						CbCardData:cardpai,
						CbActionMask:curaction,
						WCurrentUser:SendCardUser,
						WSendCardUser:SendCardUser,
						CbGanData:gangPai,
					})

					//给告诉所有用户给XX用户发牌;
					for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
						if (def.M_RoomChair[roomindex][tmpi].UserId!=0)&&(tmpi!=int(SendCardUser)) {
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_SEND_CARD_BroadCast), &def.CMD_S_SendCard_Broadcast{
								WCurrentUser: SendCardUser,
							})
						}

					}

					}

				}//needSendCard


			} else {//出牌模式:

				//def.M_NotifyUsers

                //def.M_NotifyUsers
				var NeedProcess bool
				var AllReceive bool
				var BigerOtheraction bool
				var AllWikNull bool
				AllWikNull=true
				NeedProcess=false
				AllReceive=true
				BigerOtheraction=true
				//如果本用户没执行过操作 或执行顺序在本列之内
				fmt.Println("用户接收情况",def.M_NotifyUserReceive[roomindex],"用户提示情况",def.M_NotifyUsers[roomindex],"当前坐位:",ThisChair,"已经处理操作",def.M_NotifyProcessed[roomindex],"当前提示操作序号:",def.M_NotifyIndex[roomindex],"最大提示操作",def.M_NotifyMaxAction[roomindex])
				if (!def.M_NotifyUserReceive[roomindex][ThisChair])&&(def.M_NotifyProcessed[roomindex]<=(def.M_NotifyIndex[roomindex]-1)){

					if int(receiveOPerate.CbOperateCode)==def.M_NotifyMaxAction[roomindex]{
						fmt.Println("==def.M_NotifyMaxAction[roomindex],",def.M_NotifyUserMaxactionReceive[roomindex])
						//AllWikNull=false
						def.M_NotifyUserMaxactionReceive[roomindex][ThisChair]=true
					}

					//if int(receiveOPerate.CbOperateCode)==def.M_NotifyMaxAction[roomindex]
					def.M_NotifyUserReceive[roomindex][ThisChair]=true
					def.M_cbPerformAction[roomindex][ThisChair]=receiveOPerate.CbOperateCode
					fmt.Println("def.M_NotifyUserReceive[roomindex][ThisChair]=true")
					//接收到的最大操作
					if int(receiveOPerate.CbOperateCode)>=def.M_NotifyReceiveMaxAction[roomindex]{
						def.M_NotifyReceiveMaxAction[roomindex]=int(receiveOPerate.CbOperateCode)
					}
					AllWikNull=true
					for tmpallreceive:=0;tmpallreceive<int(def.M_DesktopPlayer[roomindex]);tmpallreceive++{
						//是否全部大的操作接收到
						if def.M_NotifyUsers[roomindex][tmpallreceive]{
							if (!def.M_NotifyUserReceive[roomindex][tmpallreceive])&&(int(def.M_cbUserAction[roomindex][tmpallreceive])>=def.M_NotifyReceiveMaxAction[roomindex]){
								AllReceive=false
								fmt.Println("AllReceive=false")
							}
						}
						//是否接收到最大操作
						if tmpallreceive!=ThisChair{
								if receiveOPerate.CbOperateCode<=def.M_cbUserAction[roomindex][tmpallreceive]{
									BigerOtheraction=false
									fmt.Println("BigerOtheraction=false")
							}
						}

					}
					for tmpallreceive:=0;tmpallreceive<int(def.M_DesktopPlayer[roomindex]);tmpallreceive++{
						if def.M_NotifyUsers[roomindex][tmpallreceive]&&def.M_cbPerformAction[roomindex][tmpallreceive]!=def.WIK_NULL{
							AllWikNull=false
						}
					}
					if BigerOtheraction{
						def.M_NotifyUserMaxactionReceive[roomindex][ThisChair]=true
					}
					//执行了最大操作
					if BigerOtheraction||AllReceive{
						NeedProcess=true
						fmt.Println("BigerOtheraction||AllReceive")
					}
					//所有用户都点了过
					if AllWikNull{
						NeedProcess=true
						fmt.Println("AllWikNull")
					}
					
                    //fmt.Println("")
					if NeedProcess{
						fmt.Println("NeedProcess")
						    var fans[2] int16
							var DianPaoUserCount int
							DianPaoUserCount=0

								for i := def.M_wProvideUser[roomindex] ;i < (def.M_wProvideUser[roomindex] + def.M_DesktopPlayer[roomindex]) ;i++ {
									//存在一炮多响以及一杠一胡的提示；
									var curUser int = 0
									//取当前玩家
									curUser = int(i % def.M_DesktopPlayer[roomindex])
									fmt.Println("房间",info.RoomId,"当前用户",curUser,"执行操作",def.M_cbPerformAction[roomindex],"最大操作:",def.M_NotifyReceiveMaxAction[roomindex])
									if int(def.M_cbPerformAction[roomindex][curUser])==def.M_NotifyReceiveMaxAction[roomindex]{
										switch def.M_cbPerformAction[roomindex][curUser] {
										case def.WIK_CHI_HU:{
											//点炮几家
											if def.WinOrder[roomindex]==int16(0){
												def.M_NextwBankerUser[roomindex]=curUser
											}
											DianPaoUserCount++

											var CurJiFen int16
											CurJiFen=0
											fans[0],fans[1],CurJiFen=def.UserHuPai(roomindex,int16(curUser),int(receiveOPerate.CbOperateCard),false)
											def.M_cbCardIndex[roomindex][curUser][receiveOPerate.CbOperateCard]++
											//添加胡的那张麻将
											fmt.Println("RoomID",info.RoomId,"出牌模式吃胡 番数",fans,"供牌用户",def.M_wProvideUser[roomindex],"CbWinOrder",def.WinOrder[roomindex],"当前积分",CurJiFen,"当前用户",curUser)

											def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_Hu,int(receiveOPerate.CbOperateCard))

											lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][curUser],uint32(def.SUB_S_HU), &def.CMD_S_ChiHu{
												WChiHuUser:int16(curUser),
												WProviderUser:def.M_wProvideUser[roomindex],
												CbChiHuCard:def.M_cbProvideCard[roomindex],
												LGameScore:CurJiFen,
												CbWinOrder:def.WinOrder[roomindex],
											})

											for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomindex]);tmpj++{

												lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpj],uint32(def.SUB_S_HU_BroadCast), &def.CMD_S_ChiHuBroadCast{
													WChiHuUser:int16(curUser),
													WProviderUser:def.M_wProvideUser[roomindex],
													CbChiHuCard:def.M_cbProvideCard[roomindex],
													CbWinOrder:def.WinOrder[roomindex],
												})
											}

											def.M_GameConcludeScore[roomindex].HuType[curUser]=def.HuType_JiePao
											def.M_GameConcludeScore[roomindex].LGameScore[curUser]=def.M_GameConcludeScore[roomindex].LGameScore[curUser]+int(CurJiFen)
											def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]=def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]-int(CurJiFen)
											fmt.Println("所有玩家游戏积分",def.M_GameConcludeScore[roomindex].LGameScore,"当前积分",CurJiFen,"当前玩家:",curUser)
											//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].LGangScore=def.M_GangScore[def.GetRoomid(info.RoomId)]
											for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++ {
												if def.M_GameStatus[roomindex][curUser]!=def.GAME_STATUS_WINED{
													def.M_GameConcludeScore[roomindex].LGangScore[tmpi]=0
													for tmpj := 0; tmpj < int(def.M_DesktopPlayer[roomindex]); tmpj++ {
														def.M_GameConcludeScore[roomindex].LGangScore[tmpi] = def.M_GameConcludeScore[roomindex].LGangScore[tmpi] + def.M_GangScore[roomindex][tmpi][tmpj]
													}
												}
											}




											def.M_GameConcludeScore[roomindex].WProvideUser[curUser]=def.M_wProvideUser[roomindex]
											def.M_GameConcludeScore[roomindex].CbProvideCard[curUser]=receiveOPerate.CbOperateCard
											//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].CbHandCardData[def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]=def.M_cbCardIndex[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]
											def.M_GameConcludeScore[roomindex].WLiuJuStatus[curUser]=4
											def.M_GameConcludeScore[roomindex].UFanBase[curUser]=fans[0]
											def.M_GameConcludeScore[roomindex].UFanAddtion[curUser]=fans[1]
											SendCardUser =int16(def.GetNextChair(roomindex,int16(curUser)))
											def.WinOrder[roomindex]++
											def.M_GameStatus[roomindex][curUser]=def.GAME_STATUS_WINED
											needSendCard = true

											fmt.Println("仅剩一名玩家，游戏结束各玩家得分",def.M_GameConcludeScore[roomindex].LGameScore,"各玩家杠得分",def.M_GameConcludeScore[roomindex].LGangScore,"基本番型描述",def.M_GameConcludeScore[roomindex].UFanDescBase,"额外番型描述",def.M_GameConcludeScore[roomindex].UFanDescAddtion)
											def.M_GameTotalScoreCount[roomindex].LJiePao[curUser]++
											def.M_GameTotalScoreCount[roomindex].LDianPao[def.M_wProvideUser[roomindex]]++
											//如果打到最后只有一名玩家了。游戏结束 显示游戏得分
											if def.GetPlayIngUserCount(roomindex)<2{
												needSendCard=false
												def.GameRecordDrawScore(roomindex)
												def.M_GameConcludeScore[roomindex].Tm=time.Now()
												fmt.Println("出牌模式下杠分",def.M_GameConcludeScore[roomindex].LGangScore, def.M_GameConcludeScore[roomindex].Tm,"胡牌类型", def.M_GameConcludeScore[roomindex].HuType)
												type SendRecord struct {
													Inithead def.GameRecord_init
													Operate def.GameRecord_Operator
													GameConclude def.CMD_S_GameConclude
												}
												var BsendRec SendRecord
												BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
												BsendRec.Inithead=def.M_GameRecord_init[roomindex]
												fmt.Println("传递的数据为：",BsendRec)
												BsendRec.Operate=def.M_GameRecord_Operator[roomindex]
												//gamelib.Marshal(BsendRec)
												//gamelib.Marshal(info.RoomId)
												type Headstruct struct {
													Roomid int
													Draw int
													Userlst[def.GAME_PLAYER] string
													Gamescore[def.GAME_PLAYER] int
													Gamekind int
													Tm 		time.Time

												}
												var heads Headstruct
												heads.Roomid=int(info.RoomId)
												//heads.draw=
												heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
												fmt.Println("游戏局数:",heads.Draw)
												for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
													heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
													heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
												}
												heads.Tm = time.Now()
												heads.Gamekind=1


												header, _ := gamelib.Marshal(&heads)
												content, _ := gamelib.Marshal(&BsendRec)

												var id int
												id=lib.mgr.SaveGameRecord(header, content)
												//lib.mgr.SaveUserRecord(int(info.UserId),id)
												for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
													//lib.mgr.SaveUserRecord(int(info.UserId), id)
													lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
												}

												//发送游戏结算信息
												//fmt.Println("各玩家手中麻将为",def.M_GameConcludeScore[roomindex].CbHandCardData[0]," ",def.M_GameConcludeScore[roomindex].CbHandCardData[1])
												fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
												for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
													fmt.Println("玩家TMPI",tmpi,"积分",def.M_GameConcludeScore[roomindex].LGameScore)
													lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
														LGameScore:def.M_GameConcludeScore[roomindex].LGameScore,
														LRevenue:def.M_GameConcludeScore[roomindex].LRevenue,
														LGangScore:def.M_GameConcludeScore[roomindex].LGangScore,
														WProvideUser:def.M_GameConcludeScore[roomindex].WProvideUser,
														CbProvideCard:def.M_GameConcludeScore[roomindex].CbProvideCard,
														CbHandCardData:def.M_GameConcludeScore[roomindex].CbHandCardData,
														CbGenCount:def.M_GameConcludeScore[roomindex].CbGenCount,
														WLiuJuStatus:def.M_GameConcludeScore[roomindex].WLiuJuStatus,
														UFanDescBase:def.M_GameConcludeScore[roomindex].UFanDescBase,
														UFanDescAddtion:def.M_GameConcludeScore[roomindex].UFanDescAddtion,
														UFanBase:def.M_GameConcludeScore[roomindex].UFanBase,
														UFanAddtion:def.M_GameConcludeScore[roomindex].UFanAddtion,
														HuType:def.M_GameConcludeScore[roomindex].HuType,
														Tm:def.M_GameConcludeScore[roomindex].Tm,

													})
													def.M_GameStatus[roomindex][tmpi]=def.GAME_STATUS_WINED
													//def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi]=int16(def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGameScore[tmpi])
													def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi]=int16(def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGangScore[tmpi])

												}//For
												fmt.Println("总计游戏积分为：",def.M_GameTotalScoreCount[roomindex].LGameScore[0],def.M_GameTotalScoreCount[roomindex].LGameScore[1],def.M_GameTotalScoreCount[roomindex].LGameScore[2],def.M_GameTotalScoreCount[roomindex].LGameScore[3])

												//几局的统计结算界面
												if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
													fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)

													for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

														lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
															LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
															LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
															LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
															LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
															LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
															LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
															LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
															LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

														})

													}

													lib.mgr.ReleaseRoom()
												}
											}

										}

										case def.WIK_XiaYu:{
											def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
											def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
											def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbWeaveKind =def. WIK_XiaYu
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbCenterCard = receiveOPerate.CbOperateCard
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].WProvideUser = def.M_wProvideUser[roomindex]
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].GangSangPao=false
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].BeGangUserCount =1
											def.M_cbWeaveItemCount[roomindex][curUser]++
											def.M_GangScore[roomindex][curUser][def.M_wProvideUser[roomindex]] += 2 * def.M_initGameDrawScore[roomindex]
											def.M_GangScore[roomindex][def.M_wProvideUser[roomindex]][curUser] -= 2 * def.M_initGameDrawScore[roomindex]
											SendCardUser = int16(curUser)
											fmt.Println("出牌模式下下雨发牌用户",SendCardUser,"当前用户",curUser,"I的值 ",i,"供牌用户为:",def.M_wProvideUser[roomindex],"获得",2 * def.M_initGameDrawScore[roomindex],"积分")
											for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
												def.BeLastisGang[roomindex][tmpi]=def.WIK_NULL
											}
											def.BeLastisGang[roomindex][curUser]=def.WIK_XiaYu
											def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_XiaYu,int(receiveOPerate.CbOperateCard))

											needSendCard = true
											def.M_GameTotalScoreCount[roomindex].LAnGang[curUser]++
											//

										}
										case def.WIK_PENG:{

											def.M_DispachUser[roomindex]=curUser
											fmt.Println("房间",info.RoomId,"用户",curUser,"碰牌前:",def.M_cbCardIndex[roomindex][curUser])
											def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
											def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbWeaveKind = def.WIK_PENG
											def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbCenterCard = receiveOPerate.CbOperateCard
											def.M_cbWeaveItemCount[roomindex][curUser]++
											fmt.Println("房间",info.RoomId,"用户",curUser,"碰牌后:",def.M_cbCardIndex[roomindex][curUser])
                                            needSendCard=false
											def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_PENG,int(receiveOPerate.CbOperateCard))
											//fmt.Println("用户,")
											fmt.Println("出牌模式下碰,碰牌用户",curUser,"供牌用户",def.M_wProvideUser[roomindex],needSendCard)

										}
										case def.WIK_NULL:{
											fmt.Println("收到出牌模式下的过")


										}

										}//switch


									}//如果为最大操作

									//def.M_cbPerformAction[def.GetRoomid(info.RoomId)][curUser]=def.WIK_NULL
									//def.M_cbUserAction[def.GetRoomid(info.RoomId)][curUser]=def.WIK_NULL

								}//for 用户循环

						        def.M_NotifyProcessed[roomindex]=def.M_NotifyIndex[roomindex]

					            }//if needprocess 需要执行的保管


					if AllWikNull{

						needSendCard=true
						// SendCardUser=def.GetNextChair(roomindex,def.GetRoomChair(roomindex,def.M_wProvideUser[roomindex]))
						SendCardUser=def.GetNextChair(roomindex,def.M_wProvideUser[roomindex])
						fmt.Println("出牌模式下所有玩家点了过，下一发牌玩家",SendCardUser,"供牌玩家",def.M_wProvideUser[roomindex],"原来的坐位号",def.M_wProvideUser[roomindex])
						//fmt.Println("出牌模式下所有玩家点了过，下一发牌玩家",SendCardUser,"供牌玩家",def.M_wProvideUser[roomindex],"原来的坐位号",def.GetNextChair(roomindex,def.GetRoomChair(roomindex,def.M_wProvideUser[roomindex])))

					}


					if needSendCard{
						//for
						fmt.Println("NeedSendCard")


						var cardpai int16
						var curaction int16
						var gangpai int16
						if def.M_cbPosition[roomindex]==def.MAX_REPERTORY-1{
							//已经修改查叫

							GameLiuJU(roomindex)

							type SendRecord struct {
								Inithead def.GameRecord_init
								Operate def.GameRecord_Operator
								GameConclude def.CMD_S_GameConclude
							}
							var BsendRec SendRecord
							BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
							BsendRec.Inithead=def.M_GameRecord_init[roomindex]
							BsendRec.Operate=def.M_GameRecord_Operator[roomindex]

							type Headstruct struct {
								Roomid int
								Draw int
								Userlst[def.GAME_PLAYER] string
								Gamescore[def.GAME_PLAYER] int
								Gamekind int
								Tm 		time.Time

							}
							var heads Headstruct
							heads.Roomid=int(info.RoomId)
							//heads.Draw=int(def.M_GameRecordDraw[roomindex])
							heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
							fmt.Println("游戏局数:",heads.Draw)
							for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
								heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
								heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
							}
							heads.Tm = time.Now()
							heads.Gamekind=1

							header, _ := gamelib.Marshal(&heads)
							content, _ := gamelib.Marshal(&BsendRec)

							var id int
							id=lib.mgr.SaveGameRecord(header, content)
							//lib.mgr.SaveUserRecord(int(info.UserId),id)
							for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
								//lib.mgr.SaveUserRecord(int(info.UserId), id)
								lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
							}

							fmt.Println("游戏局数:",heads.Draw)
							//提示所有玩家游戏结束
							fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
							for tmpi := 0; tmpi < int(def.M_DesktopPlayer[roomindex]); tmpi++ {
								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
									LGameScore:      def.M_GameConcludeScore[roomindex].LGameScore,
									LRevenue:        def.M_GameConcludeScore[roomindex].LRevenue,
									LGangScore:      def.M_GameConcludeScore[roomindex].LGangScore,
									WProvideUser:    def.M_GameConcludeScore[roomindex].WProvideUser,
									CbProvideCard:   def.M_GameConcludeScore[roomindex].CbProvideCard,
									CbHandCardData:  def.M_GameConcludeScore[roomindex].CbHandCardData,
									CbGenCount:      def.M_GameConcludeScore[roomindex].CbGenCount,
									WLiuJuStatus:    def.M_GameConcludeScore[roomindex].WLiuJuStatus,
									UFanDescBase:    def.M_GameConcludeScore[roomindex].UFanDescBase,
									UFanDescAddtion: def.M_GameConcludeScore[roomindex].UFanDescAddtion,
									UFanBase:        def.M_GameConcludeScore[roomindex].UFanBase,
									UFanAddtion:     def.M_GameConcludeScore[roomindex].UFanAddtion,
								})
								def.M_GameStatus[roomindex][tmpi] = def.GAME_STATUS_WINED
							} //For

							//几局的统计结算界面
							if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
								fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)
								for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
										LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
										LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
										LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
										LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
										LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
										LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
										LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
										LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

									})

								}
							}

						}
						if def.M_cbPosition[roomindex]<(def.MAX_REPERTORY-1)&&SendCardUser<153{

							cardpai,curaction,gangpai=def.DispatchCardData(roomindex,SendCardUser)
							fmt.Println("用户发牌",cardpai,"存在杠牌",gangpai,"用户",SendCardUser)
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][SendCardUser],uint32(def.SUB_S_SEND_CARD), &def.CMD_S_SendCard{
								CbCardData:cardpai,
								CbActionMask:curaction,
								WCurrentUser:SendCardUser,
								WSendCardUser:SendCardUser,

								CbGanData:gangpai,
							})

							//给告诉所有用户给XX用户发牌;
							for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
								if (def.M_RoomChair[roomindex][tmpi].UserId!=0)&&(tmpi!=int(SendCardUser)) {
									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_SEND_CARD_BroadCast), &def.CMD_S_SendCard_Broadcast{
										WCurrentUser: SendCardUser,
									})
								}

							}

						}
					}//needSendCard







				}////如果本用户没执行过操作 或执行顺序在本列之内


					}// else 出牌模式


				}


























/*


				if ((def.M_NotifyProcessed[roomindex]<def.M_NotifyIndex[roomindex])&&(def.M_NotifyProcessed[roomindex]<=int(receiveOPerate.NotifyID)))||(receiveOPerate.CbOperateCode==def.WIK_CHI_HU){
				fmt.Println("ROOMID",def.M_RoomID[roomindex],"已经处理了的操作ID",def.M_NotifyProcessed[roomindex],"提示ID",def.M_NotifyIndex[roomindex],"接收到的操作ID",receiveOPerate.NotifyID)
				//fmt.Println("出牌模式下操作:::",receiveOPerate.CbOperateCode)
				var ThisChair int
				ThisChair=int(def.GetRoomChair(roomindex,info.UserId))
				needSendCard=false

				var isReceiveAllMaxAction bool=false

				var fans[2] int16
				isReceiveAllMaxAction=false
				//清除已经不需要执行操作的人的操作

				if def.M_cbUserAction[roomindex][ThisChair]==def.WIK_NULL{
					fmt.Println("清除掉不需要执行的操作,用户",ThisChair)
					//needSendCard=false
					def.M_cbPerformAction[roomindex][ThisChair]=def.WIK_NULL
					receiveOPerate.CbOperateCode=def.WIK_NULL
				}


				if def.GetPerFormActionCount(roomindex)==def.GetUserActionCount(roomindex){isReceiveAllMaxAction=true}
				//如果所有最大权限的用户点了操作
                fmt.Println("出牌模式下 用户信息",info.UserId,"是否接收到最大操作",isReceiveAllMaxAction,"用户最大操作",def.GetPerFormActionCount(roomindex),"待执行的最大操作",def.GetUserActionCount(roomindex),"待操作列表",def.M_cbUserAction[roomindex],"执行操作列表",def.M_cbPerformAction[roomindex],"用户",def.GetRoomChair(roomindex,info.UserId))

				if isReceiveAllMaxAction&&(def.GetPerFormActionCount(roomindex)>0) {
					//清除上一张保存的麻将(需要显示出牌不用清除)

					//清除其他不需要的等待操作

					for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
						if (def.GetUserActionRank(def.M_cbUserAction[roomindex][i])<def.GetUserActionRank(def.M_cbPerformAction[roomindex][ThisChair]))&&(def.M_GameStatus[roomindex][i]!=def.GAME_STATUS_WINED){
							def.M_cbUserAction[roomindex][i]=def.WIK_NULL
							fmt.Println("清除掉不需要执行的操作  用户待执行的操作",def.M_cbUserAction[roomindex],"执行",def.M_cbPerformAction[roomindex])
						}
					}

					var DianPaoUserCount int
					DianPaoUserCount=0
					for i := def.M_wProvideUser[roomindex] ;i < def.M_wProvideUser[roomindex] + def.GAME_PLAYER ;i++ {
						//存在一炮多响以及一杠一胡的提示；
						var curUser int = 0

						//取当前玩家
						curUser = int(i % 4)
						//fmt.Println("用户",curUser,"操作",ef.M_cbUserAction[roomindex][curUser])


						//如果给用户提示操作并且用户选择了基本的一项非过的操作
						if (def.M_cbUserAction[roomindex][curUser] !=def.WIK_NULL)&&(def.M_cbPerformAction[roomindex][curUser] !=def.WIK_NULL)&&(def.GetUserActionRank(def.M_cbPerformAction[roomindex][curUser])==def.GetMaxActionRank(roomindex))&&def.M_GameStatus[roomindex][curUser]!=def.GAME_STATUS_WINED{




								GuaFengXiaYu:=false
								var GangScore[def.GAME_PLAYER] int16

							    GangScore[0]=0
							    GangScore[1]=0
							    GangScore[2]=0
							    GangScore[3]=0
								if receiveOPerate.CbOperateCode==def.WIK_XiaYu{GuaFengXiaYu=true
								GangScore[def.M_wProvideUser[roomindex]]=2*def.M_initGameDrawScore[roomindex]
								}
								fmt.Println("发送操作结果:SUB_S_OPERATE_RESULT ",receiveOPerate.CbOperateCode,"麻将",receiveOPerate.CbOperateCard,"刮风下雨",GuaFengXiaYu,"得分",GangScore,"用户",curUser)
							var QiangGangHU bool
							QiangGangHU=false
								for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
									if def.BeLastQiangGang[roomindex][tmpi]==1{
										QiangGangHU=true
									}
								}

								for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_OPERATE_RESULT), &def.CMD_S_OperateResult{
										WOperateUser:int16(curUser),
										CbActionMask:receiveOPerate.CbOperateCode,
										WProvideUser:def.M_wProvideUser[roomindex],
										//CbOperateCode:receiveOPerate.CbOperateCode,
										CbOperateCode:def.M_cbPerformAction[roomindex][curUser],
										CbOperateCard:receiveOPerate.CbOperateCard,
										CbGuaFengXiaYu:GuaFengXiaYu,
										GangScore:GangScore,
										IsQiangGang:QiangGangHU,
									})
								}

							if QiangGangHU{}


							switch def.M_cbPerformAction[roomindex][curUser] {
							case def.WIK_CHI_HU:{
								//点炮几家
								if def.WinOrder[roomindex]==int16(0){
									def.M_NextwBankerUser[roomindex]=curUser
								}
								DianPaoUserCount++

								var CurJiFen int16
								CurJiFen=0
								fans[0],fans[1],CurJiFen=def.UserHuPai(roomindex,int16(curUser),int(receiveOPerate.CbOperateCard),false)
								def.M_cbCardIndex[roomindex][curUser][receiveOPerate.CbOperateCard]++
								//添加胡的那张麻将
								fmt.Println("RoomID",info.RoomId,"出牌模式吃胡 番数",fans,"供牌用户",def.M_wProvideUser[roomindex],"CbWinOrder",def.WinOrder[roomindex],"当前积分",CurJiFen,"当前用户",curUser)

								def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_Hu,int(receiveOPerate.CbOperateCard))

								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][curUser],uint32(def.SUB_S_HU), &def.CMD_S_ChiHu{
									WChiHuUser:int16(curUser),
									WProviderUser:def.M_wProvideUser[roomindex],
									CbChiHuCard:def.M_cbProvideCard[roomindex],
									LGameScore:CurJiFen,
									CbWinOrder:def.WinOrder[roomindex],
								})

								for tmpj:=0;tmpj<int(def.M_DesktopPlayer[roomindex]);tmpj++{

									lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpj],uint32(def.SUB_S_HU_BroadCast), &def.CMD_S_ChiHuBroadCast{
										WChiHuUser:int16(curUser),
										WProviderUser:def.M_wProvideUser[roomindex],
										CbChiHuCard:def.M_cbProvideCard[roomindex],
										CbWinOrder:def.WinOrder[roomindex],
									})
								}

								def.M_GameConcludeScore[roomindex].HuType[curUser]=def.HuType_JiePao
								def.M_GameConcludeScore[roomindex].LGameScore[curUser]=def.M_GameConcludeScore[roomindex].LGameScore[curUser]+int(CurJiFen)
								def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]=def.M_GameConcludeScore[roomindex].LGameScore[def.M_wProvideUser[roomindex]]-int(CurJiFen)
								fmt.Println("所有玩家游戏积分",def.M_GameConcludeScore[roomindex].LGameScore,"当前积分",CurJiFen,"当前玩家:",curUser)
								//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].LGangScore=def.M_GangScore[def.GetRoomid(info.RoomId)]
								for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++ {
									if def.M_GameStatus[roomindex][curUser]!=def.GAME_STATUS_WINED{
									def.M_GameConcludeScore[roomindex].LGangScore[tmpi]=0
									for tmpj := 0; tmpj < int(def.M_DesktopPlayer[roomindex]); tmpj++ {
										def.M_GameConcludeScore[roomindex].LGangScore[tmpi] = def.M_GameConcludeScore[roomindex].LGangScore[tmpi] + def.M_GangScore[roomindex][tmpi][tmpj]
									}
									}
								}




								def.M_GameConcludeScore[roomindex].WProvideUser[curUser]=def.M_wProvideUser[roomindex]
								def.M_GameConcludeScore[roomindex].CbProvideCard[curUser]=receiveOPerate.CbOperateCard
								//def.M_GameConcludeScore[def.GetRoomid(info.RoomId)].CbHandCardData[def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]=def.M_cbCardIndex[def.GetRoomid(info.RoomId)][def.GetRoomChair(def.GetRoomid(info.RoomId),int16(info.UserId))]
								def.M_GameConcludeScore[roomindex].WLiuJuStatus[curUser]=4
								def.M_GameConcludeScore[roomindex].UFanBase[curUser]=fans[0]
								def.M_GameConcludeScore[roomindex].UFanAddtion[curUser]=fans[1]
								SendCardUser =int16(def.GetNextChair(roomindex,int16(curUser)))
								def.WinOrder[roomindex]++
								def.M_GameStatus[roomindex][curUser]=def.GAME_STATUS_WINED
								needSendCard = true

								fmt.Println("仅剩一名玩家，游戏结束各玩家得分",def.M_GameConcludeScore[roomindex].LGameScore,"各玩家杠得分",def.M_GameConcludeScore[roomindex].LGangScore,"基本番型描述",def.M_GameConcludeScore[roomindex].UFanDescBase,"额外番型描述",def.M_GameConcludeScore[roomindex].UFanDescAddtion)
								def.M_GameTotalScoreCount[roomindex].LJiePao[curUser]++
								def.M_GameTotalScoreCount[roomindex].LDianPao[def.M_wProvideUser[roomindex]]++
								//如果打到最后只有一名玩家了。游戏结束 显示游戏得分
								if def.GetPlayIngUserCount(roomindex)<2{
									needSendCard=false
									def.GameRecordDrawScore(roomindex)
									def.M_GameConcludeScore[roomindex].Tm=time.Now()
									fmt.Println("出牌模式下杠分",def.M_GameConcludeScore[roomindex].LGangScore, def.M_GameConcludeScore[roomindex].Tm,"胡牌类型", def.M_GameConcludeScore[roomindex].HuType)
									type SendRecord struct {
										Inithead def.GameRecord_init
										Operate def.GameRecord_Operator
										GameConclude def.CMD_S_GameConclude
									}
									var BsendRec SendRecord
									BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
									BsendRec.Inithead=def.M_GameRecord_init[roomindex]
									fmt.Println("传递的数据为：",BsendRec)
									BsendRec.Operate=def.M_GameRecord_Operator[roomindex]
									//gamelib.Marshal(BsendRec)
									//gamelib.Marshal(info.RoomId)
									type Headstruct struct {
										Roomid int
										Draw int
										Userlst[def.GAME_PLAYER] string
										Gamescore[def.GAME_PLAYER] int
										Gamekind int
										Tm 		time.Time

									}
									var heads Headstruct
									heads.Roomid=int(info.RoomId)
									//heads.draw=
									for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
										heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
										heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
									}
									heads.Tm = time.Now()
									heads.Gamekind=1


									header, _ := gamelib.Marshal(&heads)
									content, _ := gamelib.Marshal(&BsendRec)

									var id int
									id=lib.mgr.SaveGameRecord(header, content)
									//lib.mgr.SaveUserRecord(int(info.UserId),id)
									for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
										//lib.mgr.SaveUserRecord(int(info.UserId), id)
										lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
									}

                                    //发送游戏结算信息
									//fmt.Println("各玩家手中麻将为",def.M_GameConcludeScore[roomindex].CbHandCardData[0]," ",def.M_GameConcludeScore[roomindex].CbHandCardData[1])
									fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
									for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
										fmt.Println("玩家TMPI",tmpi,"积分",def.M_GameConcludeScore[roomindex].LGameScore)
										lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
											LGameScore:def.M_GameConcludeScore[roomindex].LGameScore,
											LRevenue:def.M_GameConcludeScore[roomindex].LRevenue,
											LGangScore:def.M_GameConcludeScore[roomindex].LGangScore,
											WProvideUser:def.M_GameConcludeScore[roomindex].WProvideUser,
											CbProvideCard:def.M_GameConcludeScore[roomindex].CbProvideCard,
											CbHandCardData:def.M_GameConcludeScore[roomindex].CbHandCardData,
											CbGenCount:def.M_GameConcludeScore[roomindex].CbGenCount,
											WLiuJuStatus:def.M_GameConcludeScore[roomindex].WLiuJuStatus,
											UFanDescBase:def.M_GameConcludeScore[roomindex].UFanDescBase,
											UFanDescAddtion:def.M_GameConcludeScore[roomindex].UFanDescAddtion,
											UFanBase:def.M_GameConcludeScore[roomindex].UFanBase,
											UFanAddtion:def.M_GameConcludeScore[roomindex].UFanAddtion,
											HuType:def.M_GameConcludeScore[roomindex].HuType,
											Tm:def.M_GameConcludeScore[roomindex].Tm,

										})
										def.M_GameStatus[roomindex][tmpi]=def.GAME_STATUS_WINED
										//def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi]=int16(def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGameScore[tmpi])
										def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi]=int16(def.M_GameTotalScoreCount[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGameScore[tmpi])+int16(def.M_GameConcludeScore[roomindex].LGangScore[tmpi])

									}//For
									fmt.Println("总计游戏积分为：",def.M_GameTotalScoreCount[roomindex].LGameScore[0],def.M_GameTotalScoreCount[roomindex].LGameScore[1],def.M_GameTotalScoreCount[roomindex].LGameScore[2],def.M_GameTotalScoreCount[roomindex].LGameScore[3])

									//几局的统计结算界面
									if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
										fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)

										for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

											lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
												LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
												LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
												LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
												LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
												LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
												LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
												LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
												LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

											})

										}

										lib.mgr.ReleaseRoom()
									}
								}

							}

							case def.WIK_XiaYu:{
								def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
								def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
								def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbWeaveKind =def. WIK_XiaYu
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbCenterCard = receiveOPerate.CbOperateCard
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].WProvideUser = def.M_wProvideUser[roomindex]
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].GangSangPao=false
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].BeGangUserCount =1
								def.M_cbWeaveItemCount[roomindex][curUser]++
								def.M_GangScore[roomindex][curUser][def.M_wProvideUser[roomindex]] += 2 * def.M_initGameDrawScore[roomindex]
								def.M_GangScore[roomindex][def.M_wProvideUser[roomindex]][curUser] -= 2 * def.M_initGameDrawScore[roomindex]
								SendCardUser = int16(curUser)
								fmt.Println("出牌模式下下雨发牌用户",SendCardUser,"当前用户",curUser,"I的值 ",i,"供牌用户为:",def.M_wProvideUser[roomindex],"获得",2 * def.M_initGameDrawScore[roomindex],"积分")
								for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
									def.BeLastisGang[roomindex][tmpi]=def.WIK_NULL
								}
								def.BeLastisGang[roomindex][curUser]=def.WIK_XiaYu
								def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_XiaYu,int(receiveOPerate.CbOperateCard))

								needSendCard = true
								def.M_GameTotalScoreCount[roomindex].LAnGang[curUser]++
								//

							}
							case def.WIK_PENG:{

								def.M_DispachUser[roomindex]=curUser
								def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
								def.RemoveCard(roomindex,curUser,int(receiveOPerate.CbOperateCard))
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbWeaveKind = def.WIK_PENG
								def.M_WeaveItemArray[roomindex][curUser][def.M_cbWeaveItemCount[roomindex][curUser]].CbCenterCard = receiveOPerate.CbOperateCard
								def.M_cbWeaveItemCount[roomindex][curUser]++

								def.SaveTheGameRecord(roomindex,int(def.M_wProvideUser[roomindex]),curUser,def.Operator_PENG,int(receiveOPerate.CbOperateCard))
								//fmt.Println("用户,")
								fmt.Println("出牌模式下碰,碰牌用户",curUser,"供牌用户",def.M_wProvideUser[roomindex],needSendCard)

							}
							case def.WIK_NULL:{
								fmt.Println("收到出牌模式下的过")


							}

							}


						}

						//def.M_cbPerformAction[def.GetRoomid(info.RoomId)][curUser]=def.WIK_NULL
						//def.M_cbUserAction[def.GetRoomid(info.RoomId)][curUser]=def.WIK_NULL

					}//for
					if DianPaoUserCount>1{def.M_NextwBankerUser[roomindex]=int(def.M_wProvideUser[roomindex])}

				}

				//清除所有待执行操作的数据;
				//if


				if def.GetUserActionRank(def.M_cbPerformAction[roomindex][ThisChair]) ==def.GetMaxActionRank(roomindex){
					def.M_cbPerformAction[roomindex][ThisChair]=receiveOPerate.CbOperateCode
					def.M_cbUserAction[roomindex][ThisChair]=receiveOPerate.CbOperateCode
				}

				//如果所有玩家取消;
				var isallSelectwiknull bool
				isallSelectwiknull=true
				for gamei:=0;gamei<int(def.M_DesktopPlayer[roomindex]);gamei++{
					if def.M_cbPerformAction[roomindex][gamei]!=def.WIK_NULL&&def.M_cbUserAction[roomindex][gamei]!=def.WIK_NULL{
						isallSelectwiknull=false
					}
				}
				if !isReceiveAllMaxAction{isallSelectwiknull=false}
                //如果所有玩家取消则下一家发牌;
				if isallSelectwiknull{

					needSendCard=true
					// SendCardUser=def.GetNextChair(roomindex,def.GetRoomChair(roomindex,def.M_wProvideUser[roomindex]))
					SendCardUser=def.GetNextChair(roomindex,def.M_wProvideUser[roomindex])
					fmt.Println("出牌模式下所有玩家点了过，下一发牌玩家",SendCardUser,"供牌玩家",def.M_wProvideUser[roomindex],"原来的坐位号",def.M_wProvideUser[roomindex])
					//fmt.Println("出牌模式下所有玩家点了过，下一发牌玩家",SendCardUser,"供牌玩家",def.M_wProvideUser[roomindex],"原来的坐位号",def.GetNextChair(roomindex,def.GetRoomChair(roomindex,def.M_wProvideUser[roomindex])))

				}
				if isReceiveAllMaxAction{
					for gamei:=0;gamei<int(def.M_DesktopPlayer[roomindex]);gamei++{
						def.M_cbPerformAction[roomindex][gamei]=def.WIK_NULL
						def.M_cbUserAction[roomindex][gamei]=def.WIK_NULL
					}
				}

				if needSendCard{
					//for
					fmt.Println("NeedSendCard")

					var cardpai int16
					var curaction int16
					var gangpai int16
					if def.M_cbPosition[roomindex]==def.MAX_REPERTORY-1{
                     //已经修改查叫

						GameLiuJU(roomindex)

						type SendRecord struct {
							Inithead def.GameRecord_init
							Operate def.GameRecord_Operator
							GameConclude def.CMD_S_GameConclude
						}
						var BsendRec SendRecord
						BsendRec.GameConclude=def.M_GameConcludeScore[roomindex]
						BsendRec.Inithead=def.M_GameRecord_init[roomindex]
						BsendRec.Operate=def.M_GameRecord_Operator[roomindex]

						type Headstruct struct {
							Roomid int
							Draw int
							Userlst[def.GAME_PLAYER] string
							Gamescore[def.GAME_PLAYER] int
							Gamekind int
							Tm 		time.Time

						}
						var heads Headstruct
						heads.Roomid=int(info.RoomId)
						//heads.Draw=int(def.M_GameRecordDraw[roomindex])
						heads.Draw=def.M_GameRecord_init[roomindex].DrawIndex
						fmt.Println("游戏局数:",heads.Draw)
						for i:=0;i<int(def.M_DesktopPlayer[roomindex]);i++{
							heads.Userlst[i]=def.M_RoomChair[roomindex][i].Name
							heads.Gamescore[i]=def.M_GameConcludeScore[roomindex].LGameScore[i]
						}
						heads.Tm = time.Now()
						heads.Gamekind=1

						header, _ := gamelib.Marshal(&heads)
						content, _ := gamelib.Marshal(&BsendRec)

						var id int
						id=lib.mgr.SaveGameRecord(header, content)
						//lib.mgr.SaveUserRecord(int(info.UserId),id)
						for tmpsavei := 0; tmpsavei < int(def.M_DesktopPlayer[roomindex]); tmpsavei++ {
							//lib.mgr.SaveUserRecord(int(info.UserId), id)
							lib.mgr.SaveUserRecord(int(def.M_RoomChair[roomindex][tmpsavei].UserId), id)
						}

						fmt.Println("游戏局数:",heads.Draw)
						//提示所有玩家游戏结束
						fmt.Println("ROOMID",def.M_RoomID[roomindex],"游戏结算界面：",def.M_GameConcludeScore[roomindex].CbProvideCard,"手牌:",def.M_GameConcludeScore[roomindex].CbHandCardData[0],def.M_GameConcludeScore[roomindex].CbHandCardData[1],def.M_GameConcludeScore[roomindex].CbHandCardData[2],def.M_GameConcludeScore[roomindex].CbHandCardData[3])
						for tmpi := 0; tmpi < int(def.M_DesktopPlayer[roomindex]); tmpi++ {
							lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi], uint32(def.SUB_S_GAME_END), &def.CMD_S_GameConclude{
								LGameScore:      def.M_GameConcludeScore[roomindex].LGameScore,
								LRevenue:        def.M_GameConcludeScore[roomindex].LRevenue,
								LGangScore:      def.M_GameConcludeScore[roomindex].LGangScore,
								WProvideUser:    def.M_GameConcludeScore[roomindex].WProvideUser,
								CbProvideCard:   def.M_GameConcludeScore[roomindex].CbProvideCard,
								CbHandCardData:  def.M_GameConcludeScore[roomindex].CbHandCardData,
								CbGenCount:      def.M_GameConcludeScore[roomindex].CbGenCount,
								WLiuJuStatus:    def.M_GameConcludeScore[roomindex].WLiuJuStatus,
								UFanDescBase:    def.M_GameConcludeScore[roomindex].UFanDescBase,
								UFanDescAddtion: def.M_GameConcludeScore[roomindex].UFanDescAddtion,
								UFanBase:        def.M_GameConcludeScore[roomindex].UFanBase,
								UFanAddtion:     def.M_GameConcludeScore[roomindex].UFanAddtion,
							})
							def.M_GameStatus[roomindex][tmpi] = def.GAME_STATUS_WINED
						} //For

						//几局的统计结算界面
						if def.M_GameRecordDraw[roomindex]>=def.M_GameRoomsDrawinit[roomindex]{
							fmt.Println("初始化局数",def.M_GameRoomsDrawinit[roomindex],"游戏积分:",def.M_GameTotalScoreCount[roomindex].LGameScore,"自摸情况",def.M_GameTotalScoreCount[roomindex].LZhiMo,"点炮情况",def.M_GameTotalScoreCount[roomindex].LDianPao)
							for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{

								lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_GAME_TOTALDRAW), &def.CMD_S_GameTotalScore{
									LGameScore:def.M_GameTotalScoreCount[roomindex].LGameScore,
									LZhiMo:def.M_GameTotalScoreCount[roomindex].LZhiMo,
									LJiePao:def.M_GameTotalScoreCount[roomindex].LJiePao,
									LDianPao:def.M_GameTotalScoreCount[roomindex].LDianPao,
									LAnGang:def.M_GameTotalScoreCount[roomindex].LAnGang,
									LMingGang:def.M_GameTotalScoreCount[roomindex].LMingGang,
									LChaDaJia:def.M_GameTotalScoreCount[roomindex].LChaDaJia,
									LRoomCard:def.M_GameTotalScoreCount[roomindex].LRoomCard,

								})

							}
						}

					}
		if def.M_cbPosition[roomindex]<(def.MAX_REPERTORY-1)&&SendCardUser<153{

					cardpai,curaction,gangpai=def.DispatchCardData(roomindex,SendCardUser)
					fmt.Println("用户发牌",cardpai,"存在杠牌",gangpai,"用户",SendCardUser)
					lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][SendCardUser],uint32(def.SUB_S_SEND_CARD), &def.CMD_S_SendCard{
						CbCardData:cardpai,
						CbActionMask:curaction,
						WCurrentUser:SendCardUser,
						WSendCardUser:SendCardUser,

						CbGanData:gangpai,
					})

					//给告诉所有用户给XX用户发牌;
					for tmpi:=0;tmpi<int(def.M_DesktopPlayer[roomindex]);tmpi++{
						lib.mgr.SendGameMessage(&def.M_RoomChair[roomindex][tmpi],uint32(def.SUB_S_SEND_CARD_BroadCast), &def.CMD_S_SendCard_Broadcast{
							WCurrentUser:SendCardUser,
						})

					}

					}
				}//needSendCard
				def.M_NotifyProcessed[roomindex]=def.M_NotifyIndex[roomindex]//防止执行多次操作

			}//if def.M_NotifyProcessed[GetRoomID]<def.M_NotifyIndex
			}


		}*/
	//case

	}
	//se
	return nil
}

func (lib *xzlib) OnTimer(id uint32, data interface{}) {

}

func (lib *xzlib) GetPlayerCount() int {
	roomid := lib.mgr.GetRoomId()
	roomindex :=def.GetRoomid(roomid)
	var PlayerUserCount int
	PlayerUserCount=4
	if roomindex<1999{
		PlayerUserCount=int(def.M_DesktopPlayer[roomindex])
	}else{fmt.Println("房间超标",def.GetRoomid(roomid),"房间ID",roomid)}
	return PlayerUserCount
}