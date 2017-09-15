package def

import (
	"math/rand"
	"time"
	//	"fmt"
	"exportor/defines"
	"flag"
	"fmt"
	//	"log"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	//	"go/token"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

const (
	GameLibXz = 1
)
const MAX_WEAVE int16 = 4 //最大组合

const GAME_PLAYER int16 = 4     //游戏用户数
const MAX_INDEX int16 = 30      //最大索引
const MAX_COUNT int16 = 14      //最大麻将数
const MAX_REPERTORY int16 = 108 //最大库存
//////////////////////////////////////////////////////////////////////////
//动作定义

//动作标志
//动作标志
const WIK_NULL int16 = 0x00    //没有类型
const WIK_LEFT int16 = 0x01    //左吃类型
const WIK_CENTER int16 = 0x02  //中吃类型
const WIK_RIGHT int16 = 0x04   //右吃类型
const WIK_PENG int16 = 0x08    //碰牌类型
const WIK_GuaFeng int16 = 0x10 //刮风类型
const WIK_XiaYu int16 = 0x20   //下雨类型
const WIK_LISTEN int16 = 0x40  //吃牌类型
const WIK_CHI_HU int16 = 0x80  //吃胡类型
const WIK_ZI_MO int16 = 0x100  //自摸

//////////////////////////////////////////////////////////////////////////
//胡牌定义

//胡牌
const CHK_NULL int16 = 0x00 //非胡类型
//牌型
/*
const CHR_DA_DUI_ZI int16 = 0x00000001   //大对子
const CHR_QI_XIAO_DUI int16 = 0x00000002 //暗七对
const CHR_SHU_FAN int16 = 0x00000010     //素番
//const	CHR_LONG_QI_DUI	int16=	0x00000004	//龙七对
//const	CHR_DANDIAO	int16=	0x00000008	//单吊

//翻型
const CHR_QIANG_GANG int32 = 0x00010000     //抢杠
const CHR_GANG_SHANG_PAO int32 = 0x00020000 //杠上炮
const CHR_GANG_KAI int32 = 0x00020000       //杠上花
const CHR_TIAN_HU int32 = 0x00080000        //天胡
const CHR_DI_HU int32 = 0x00100000          //地胡
const CHR_QING_YI_SE int32 = 0x00200000     //清一色
const CHR_HAIDILAO int32 = 0X00200000       //海底捞
const CHR_DAI_YAO int32 = 0x00800010        //带幺
const CHR_DUAN_YAO int32 = 0x01000000       //断幺
const CHR_MENQING int32 = 0x02000000        //门清
const CHR_JIANG_DUI int32 = 0x02000000      //将对
*/
//游戏状态
const GAME_STATUS_FREE int16 = 0        //空闲状态
const GAME_STATUS_Ready int16 = 1       //换三张
const GAME_STATUS_CHANGETHREE int16 = 2 //换三张
const GAME_STATUS_DingQUE int16 = 4     //定缺
const GAME_STATUS_PLAY int16 = 8        //游戏状态
const GAME_STATUS_SCore int16 = 10      //结算状态
const GAME_STATUS_Repay int16 = 20      //游戏回放
const GAME_STATUS_WAIT int16 = 40       //等待状态
const GAME_STATUS_WINED int16 = 80      //游戏完成
/*
0 发牌
1 出牌
2 碰
3 刮风
4 下雨
5 胡牌*/
const Operator_Draw int=0//摸牌
const Operator_Play int=1//出牌
const Operator_PENG int=2//碰牌
const Operator_GuaFeng int=3//刮风
const Operator_XiaYu   int=4//下雨
const Operator_Hu      int=5//下雨


const HuType_JiePao   int16=2 //接炮
const HuType_ZhiMo    int16=3 //自摸
const HuType_ChaJiao  int16=4 //查叫

const LiuJuStatus_HuaZhu  int=1//流局花猪
const LiuJuStatus_YouJiao int=2//流局有叫
const LiuJustatus_MeiJiao int=3//流局没叫


//--*********************      服务器命令结构       ************--
const SUB_S_GAMERULE int16 = 1100               //发送用户私人房间规则
const SUB_S_OTHERUSERINFO int16 = 1101          //其他用户进入房间
const SUB_S_USER_SITDOWN int16 = 1102           //用户坐下给用户分配坐位号
const SUB_s_User_sitdown_Broadcast int16 = 1103 //用户坐下广播消息
const SUB_S_USER_READY int16 = 1104             //用户准备好
const SUB_S_GAME_START int16 = 1105             //游戏开始
const SUB_S_CHANGETHREE_TUIJIAN int16 = 1106    //推荐用户换的三张牌
const SUB_S_CHANGETHREE int16 = 1107            //换三张。
const SUB_S_COLOR_TUIJIAN int16 = 1108          //发送用户定缺,用户缺的花色
const SUB_S_REPLAY_COLOR int16 = 1109           //发送用户定缺,用户缺的花色
const SUB_S_OUT_CARD int16 = 1110               //出牌命令
const SUB_S_SEND_CARD int16 = 1111              //发送扑克
const SUB_S_SEND_CARD_BroadCast int16 = 1112    //发送扑克广播
const SUB_S_OPERATE_NOTIFY int16 = 1113         //操作提示
const SUB_S_OPERATE_RESULT int16 = 1114         //操作命令
///const SUB_S_GANG_Card		int16=110 //杠牌
///const SUB_S_HU_CARD		int16=110 //胡哪些牌
const SUB_S_HU int16 = 1115 //胡牌
///const SUB_S_HU_Replay		int16=111 //胡牌确认
///const SUB_S_RECORD	    int16=111 //游戏记录
//const SUB_S_GAME_END		int16=1115 //游戏结束
//const SUB_S_GAME_TOTALDRAW	int16=1116 //游戏结束

const SUB_S_HU_BroadCast int16 = 1116   //胡牌
const SUB_S_GAME_END int16 = 1117       //游戏结束
const SUB_S_GAME_TOTALDRAW int16 = 1118 //游戏结束
const SUB_S_GAME_ReenterRoom int16=1119 //重新进入房间
const SUB_S_Position int16 = 1120
const SUB_S_UserLeave_BroadCast int16 = 1121
const SUB_S_GameRecord_Init int16 = 1140     //游戏结束
const SUB_S_GameRecord_Operator int16 = 1141 //游戏结束
const SUB_S_GAMERecord_End int16 = 1142      //游戏结束

const SUB_C_READY int16 = 1130        //用户准备好
const SUB_C_CHANGETHREE int16 = 1131  //--发送换3张
const SUB_C_SEND_COLOR int16 = 1132   //--发送定缺(发送花色)
const SUB_C_OUT_CARD int16 = 1133     //--出牌命令
const SUB_C_OPERATE_CARD int16 = 1134 //--操作扑克
const SUB_C_GpsPosition int16 = 1135  //GPS位置信息

const SUB_C_GetCardlist int16 = 1161
const SUB_S_SendCardList int16 = 1171 //

//////////////////////////////////////////////////////////////////////////

//const ZI_PAI_COUNT int16 = 7
const INVALID_CHAIR int16 = 0xFF //无效椅子

type GameRecord_init struct {
	//GameUserInfo[GAME_PLAYER][3] int //回放中的玩家信息 0-USERid 1-UserName 2-UserImage
	GameUserID    [GAME_PLAYER]int
	GameUserName  [GAME_PLAYER]string
	GameUserImage [GAME_PLAYER]string
	Sice          [2]int                      //回放的色子点数
	BankerUser    int                         //庄家
	GameRule      XzConf                      //回放中的游戏规则;
	Cardlist      [108]int                    //整桌的麻将列表
	UserCard      [GAME_PLAYER][MAX_COUNT]int //用户手中的麻将;
	Change3       [GAME_PLAYER][6]int         //换三张 前3代表换之前的麻将，后三张代表换之后的麻将
	DingQue       [GAME_PLAYER]int            //用户定缺
	RoomID        int                         //房间号
	DrawIndex     int                         //局数
}
type GamePositionRoom struct {
	X     [GAME_PLAYER]float64
	Y     [GAME_PLAYER]float64
	Chair [GAME_PLAYER]int
}
type GameUserLeave struct {
	Chair int
}

type GameRecord_Operator_Detail struct {
	Chairid       int //座位号
	Operator_type int // 操作类型
	Card          int //麻将
	Operatorindex int //操作序号
	ProvideUser   int //供牌用户   //发牌时 (发头牌 发尾牌) 头牌55 尾牌99 其他牌 80
}

type GameRecord_Operator struct {
	GameRecords [150]GameRecord_Operator_Detail
}

type EnterRoomuserInfo struct {
	UserId   int
	ChairID  int
	Name     string
	HeadImg  string
	RoomCard int
	GameGOLD int
	Sex int
	//Tm int64
	Tm time.Time
}
type EnterRoomOtheruserInfo struct {
	UserId   [GAME_PLAYER]int
	ChairID  [GAME_PLAYER]int
	Name     [GAME_PLAYER]string
	HeadImg  [GAME_PLAYER]string
	RoomCard [GAME_PLAYER]int
	GameGOLD [GAME_PLAYER]int
}
type EnterRoomBroadcast struct {
	UserId   int
	ChairID  int
	Name     string
	HeadImg  string
	RoomCard int
	GameGOLD int
}

//类型子项
type tagKindItem struct {
	cbWeaveKind  int16    //组合类型
	cbCenterCard int16    //中心扑克
	cbCardIndex  [3]int16 //扑克索引
	cbValidIndex [3]int16 //实际扑克索引
}

//组合子项
type tagWeaveItem struct {
	CbWeaveKind  int16 //组合类型
	CbCenterCard int16 //中心扑克
	//CbPublicCard int16						//公开标志
	WProvideUser    int16              //供应用户
	BeGangUserCount int16              //被杠多少家
	GangUserList    [GAME_PLAYER]int16 //都有哪些用户
	GangSangPao     bool
	//CbCardData[GAME_PLAYER] int16					//麻将数据
}

//游戏结束
type CMD_S_GameConclude struct {
	LGameScore      [GAME_PLAYER]int              //游戏币
	LRevenue        [GAME_PLAYER]int16            //税收
	LGangScore      [GAME_PLAYER]int16            //本局杠输赢分
	WProvideUser    [GAME_PLAYER]int16            //供应用户 如果 wProvidUser[0] = 0 表示自摸， wProvideUser[0]=1表示 1给玩家0点炮
	CbProvideCard   [GAME_PLAYER]int16            //供应扑克
	CbHandCardData  [GAME_PLAYER][MAX_COUNT]int16 //麻将列表
	CbGenCount      [GAME_PLAYER]int16            //根数录
	WLiuJuStatus    [GAME_PLAYER]int16            //流局标志
	UFanDescBase    [GAME_PLAYER]string           //基本番描述
	UFanDescAddtion [GAME_PLAYER]string           //额外番描述
	UFanBase        [GAME_PLAYER]int16            //基本番描述
	UFanAddtion     [GAME_PLAYER]int16            //额外番描述
	HuType          [GAME_PLAYER]int16            //胡牌类型 //接炮 2 自摸3 查叫4
	Tm 		time.Time
}

//游戏结算界面
type CMD_S_GameTotalScore struct {
	LGameScore [GAME_PLAYER]int16
	LZhiMo     [GAME_PLAYER]int16
	LJiePao    [GAME_PLAYER]int16
	LDianPao   [GAME_PLAYER]int16
	LAnGang    [GAME_PLAYER]int16
	LMingGang  [GAME_PLAYER]int16
	LChaDaJia  [GAME_PLAYER]int16
	LRoomCard  [GAME_PLAYER]int16
}

//杠牌结果
type tagGangCardResult struct {
	cbCardCount int16    //扑克数目
	cbCardData  [4]int16 //扑克数据
}

//分析子项
type tagAnalyseItem struct {
	cbCardEye    int16       //牌眼扑克
	bMagicEye    int16       //牌眼是否是王霸
	cbWeaveKind  [4]int16    //组合类型
	cbCenterCard [4]int16    //中心扑克
	cbCardData   [4][4]int16 //实际扑克
}

//胡牌选项
type tagHuAnalyseItem struct {
	cbCardData [4][3]int16 //实际扑克
	cbJiang    int16       //将
}

//杠牌得分
type tagGangScore struct {
	cbGangCount int16                         //杠个数
	lScore      [MAX_WEAVE][GAME_PLAYER]int32 //每个杠得分
}

//操作提示
type CMD_S_OperateNotify struct {
	WResumeUser  int16 //还原用户
	CbActionMask int16 //动作掩码
	CbActionCard int16 //动作扑克
	NotifyID   int16//提示编号
}

//发送扑克
type CMD_S_SendCard struct {
	CbCardData    int16 //扑克数据
	CbGanData     int16
	CbActionMask  int16 //动作掩码
	WCurrentUser  int16 //当前用户
	WSendCardUser int16 //发牌用户
	//	wReplaceUser  WORD	         	 //补牌用户
	BTail bool //末尾发牌
}

//发送扑克广播
type CMD_S_SendCard_Broadcast struct {
	WCurrentUser int16 //当前用户
}

//操作命令
type CMD_S_OperateResult struct {
	WOperateUser   int16              //操作用户
	CbActionMask   int16              //动作掩码
	WProvideUser   int16              //供应用户
	CbOperateCode  int16              //操作代码
	CbOperateCard  int16              //操作扑克
	CbGuaFengXiaYu bool               //是刮风还是下雨 false 刮风 TRUE 下雨
	GangScore      [GAME_PLAYER]int16 //杠牌得分
	IsQiangGang    bool               //是否抢杠
	OPERATE_NOTIFYID int16            //操作的序号
}

//---定缺
type CMD_S_ReplayColor struct {
	CbUser int16
	Cbcard int16
}

//////换三张
type CMD_CHANGE_THREE struct {
	CbUser int16
	Cbcard [3]int16
}

//////换三张
type CMD_Rceive_THREE struct {

	//CbUser int16
	CbCard [3]int16
}

//////换三张
type CMD_S_CHANGE3 struct {
	CbUser    int16
	Cbcard    [3]int16
	CbNewcard [3]int16
}

type CMD_C_SEND_COLOR struct {
	CbColorData int16 //--定缺的扑克花色
}

//-- 吃胡结构体
type CMD_S_ChiHu struct {
	WChiHuUser    int16 //							-- 吃胡玩家
	WProviderUser int16 //							-- 吃胡提供者
	CbChiHuCard   int16 //							-- 胡牌
	CbCardCount   int16 //							-- 牌数
	LGameScore    int16 //						        -- 获得积分
	CbWinOrder    int16 //							-- 吃胡排名
}

//-- 吃胡广播结构体
type CMD_S_ChiHuBroadCast struct {
	WChiHuUser    int16 //							-- 吃胡玩家
	WProviderUser int16 //							-- 吃胡提供者
	CbChiHuCard   int16 //							-- 胡牌
	CbWinOrder    int16 //							-- 吃胡排名
}

type CMD_S_GameStart struct {
	WBankerUser int16            //当前庄家
	WSice1      int16            //骰子点数
	WSice2      int16            //骰子点数
	CbCardData  [MAX_COUNT]int16 //麻将列表
	Draw        int16            //局数
	//RefChange3[3]     int16        //推荐的换三张
}

//--用户出牌
type CMD_C_OutCard struct {
	CbCardData int16 //--出牌麻将
}

//--操作命令
type CMD_C_OperateCard struct {
	CbOperateCode int16 //操作代码
	CbOperateCard int16 //操作扑克
	NotifyID int16//提示ID
}

//--用户出牌
type CMD_S_OutCard struct {
	WOutCardUser  int16 //--出牌用户
	CbOutCardData int16 //--出牌扑克
}

type XzConf struct {
	PlayerCount       int
	Roomxf            bool
	FanCount          int  //3番 4番 5番 0番
	ZiMoType          bool //加底 加番
	Huan3             bool //换三张
	DaiYaoJiuJiangDui bool //带幺九将对
	MenQing           bool //门清
	TianDiHu          bool //天地胡
	HuJiaoZhuanYi     bool //呼叫转移
	GameJU            int  //游戏局数
}

////////////////////////
////设置执行动作
func SetUserAction(roomid int, action int16, chairID int16) {
	M_cbUserAction[roomid][chairID] = action
	return
}

////////////////////////
/////设置执行动作
func SetPerformAction(roomid int, action int16, chairID int16) {
	M_cbPerformAction[roomid][chairID] = action
	return
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func SaveTheGameRecord(roomid int,ProvideUser int,CurUser int,Operatortype int,Card int)  {
	/*
	Operatortype 0 发牌
	             1 出牌
				 2 碰
				 3 刮风
				 4 下雨
				 5 胡牌
	 */
	if M_GameRecord_Operator_Index[roomid]<149{
		M_GameRecord_Operator[roomid].GameRecords[M_GameRecord_Operator_Index[roomid]].ProvideUser=ProvideUser
		M_GameRecord_Operator[roomid].GameRecords[M_GameRecord_Operator_Index[roomid]].Operator_type=Operatortype
		M_GameRecord_Operator[roomid].GameRecords[M_GameRecord_Operator_Index[roomid]].Chairid=CurUser
		M_GameRecord_Operator[roomid].GameRecords[M_GameRecord_Operator_Index[roomid]].Card=Card
		M_GameRecord_Operator[roomid].GameRecords[M_GameRecord_Operator_Index[roomid]].Operatorindex=M_GameRecord_Operator_Index[roomid]
		M_GameRecord_Operator_Index[roomid]++
	}else{fmt.Println("保存记录出错",Operatortype,CurUser)}
}

func GetFileContentAsStringLines(filePath string) string {
	//fmt.Println("get file content as lines: %v", filePath)
	var result string
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		//fmt.Println("read file: %v error: %v", filePath, err)
		return result
	}
	result = string(b)

	//fmt.Println("get file content as lines: %v, size: %v", filePath, len(result))
	return result
}

//取当前玩家个数
func GetPlayIngUserCount(Roomid int) int16 {
	var counts int16
	counts = 0
	for i := 0; i < int(M_DesktopPlayer[Roomid]); i++ {
		if M_GameStatus[Roomid][i] != GAME_STATUS_WINED {
			counts++
		}
	}
	return counts
}

type LastBeGang struct {
	Card int
	Kind int
	Users int
}
type OutCardstruct struct {
	OutCard int
	ChairID int
}

var M_GameStatus [2000][GAME_PLAYER]int16             //游戏状态
var M_wBankerUser [2000]int16                         //庄家
var M_Sice [2]int16                                   //2个骰子
//var M_bPlayStatus [2000][GAME_PLAYER]bool             //玩家状态
var M_cbCardData [2000][108]int16                     //麻将数据
var M_cbCardIndex [2000][GAME_PLAYER][MAX_INDEX]int16 //玩家麻将索引
var M_cbPosition [2000]int16                          //玩家取牌的位置
var M_wProvideUser [2000]int16
var M_cbChangeThree [2000][GAME_PLAYER][3]int16 //玩家换三张的牌
var M_dingqueColor [2000][GAME_PLAYER]int16
var M_cbUserAction [2000][GAME_PLAYER]int16    //提示玩家执行的操作
var M_cbPerformAction [2000][GAME_PLAYER]int16 //玩家选择的执行操作
//var m_wResumeUser [2000]int16                  //还原用户
//var		m_wCurrentUser	int16	//当前用户
//var		m_wProvideUser	int16	//供应用户
var M_cbProvideCard [2000]int16                      //供应扑克
var BeLastisGang [2000][GAME_PLAYER]int16            //最后一张出牌状态
var BeLastQiangGang [2000][GAME_PLAYER]int16         //最后被抢杠
var M_GameConcludeScore [2000]CMD_S_GameConclude     //游戏结束统计
var M_GameTotalScoreCount [2000]CMD_S_GameTotalScore //几局统计的总分;
var M_GameRecordDraw [2000]int16                     //局数;
var M_GameRoomsDrawinit [2000]int16                  //多少个局数一个大统计

var M_DesktopPlayer [2000]int16        //游戏中的玩家个数
var M_GameMaxFan [2000]int16           //游戏中最大番数
var M_GameZiMoAddtype [2000]bool       //自摸加底，自摸加番设置 1自摸加底 2自摸加番
var M_GameHuan3 [2000]bool             //换三张
var M_GameDaiyaoandJiangDui [2000]bool //带幺九将对
var M_GameMenQingZhongZhang [2000]bool //门清中张
var M_GameTianDiHu [2000]bool          //天地胡
var M_GameHuJiaoZhuanYi [2000]bool     //天地胡
var M_GameXiaoHU[2000]bool             //小胡不能胡
var M_GameRoomFeiYongType [2000]bool   //扣房卡类型

//组合扑克
//protected:
var M_cbWeaveItemCount [2000][GAME_PLAYER]int16                 //组合数目
var M_WeaveItemArray [2000][GAME_PLAYER][MAX_WEAVE]tagWeaveItem //组合扑克
var M_cbSendCardCount [2000][GAME_PLAYER]int16                  //发牌数目
var M_RoomChairPosition [2000]int16                             //取用户座位号的POSITION
var M_RoomChair [2000][GAME_PLAYER]defines.PlayerInfo
var Tableslink_daiyao [2000][GAME_PLAYER]int16 //是否带幺

var M_initGameDrawScore [2000]int16
var M_GangScore [2000][GAME_PLAYER][GAME_PLAYER]int16 //前面为某用户[后面为某一供牌用户]
//var     M_GameScore[2000][GAME_PLAYER] int16
var WinOrder [2000]int16
var M_RoomConf [2000]defines.CreateRoomConf
var M_RoomCount int
var M_RoomID [2000]uint32
var M_NextwBankerUser [2000]int
var M_GameRecord_init [2000]GameRecord_init
var M_GameRecord_Operator [2000]GameRecord_Operator
var M_GameRecord_Operator_Index [2000]int

var M_Position [2000]GamePositionRoom
var M_PositionCount [2000]int
var M_OutCardList[2000][GAME_PLAYER][41] OutCardstruct
var M_OutCardListIndex[2000][GAME_PLAYER] int
var M_DispachUser[2000] int
var M_NotifyIndex[2000] int
var M_NotifyProcessed[2000] int
var M_NotifyMaxAction[2000] int
var M_NotifyUsers[2000][GAME_PLAYER] bool
var M_NotifyUserReceive[2000][GAME_PLAYER] bool
var M_NotifyUserMaxactionReceive[2000][GAME_PLAYER] bool
var M_NotifyReceiveMaxAction[2000] int
var M_NeedReturnRoomCard[2000] bool
var M_RoomCreater[2000] defines.PlayerInfo


func GetRoomChair(RoomID int, UserID uint32) int16 {
	var result_chairid int16
	result_chairid = 0
	for i := 0; i < int(M_DesktopPlayer[RoomID]); i++ {
		if M_RoomChair[RoomID][i].UserId == UserID {
			result_chairid = int16(i)
			//fmt.Println("查找到用户坐位号",i)
		}

	}
	return result_chairid
}

//设置已经定缺的花色
func SetQueColor(Color int16, roomID int, ChairID int16) {
	M_dingqueColor[roomID][ChairID] = Color
}

//取麻将花色
func GetCardColor(card int16) int16 {
	var cardColor int16 = 0
	cardColor = card / 10
	return cardColor
}

////////////////////////////////
///取定缺的花色
func GetQueueColor(RoomID int, chairid int16) int16 {
	var color int16 = 1
	//fmt.Println("取定缺花色 ROOM",RoomID,"座位号",chairid)

	color = M_dingqueColor[RoomID][chairid]

	return color
}
func GetFirstZeroRoom() int{
	var roomid int
	roomid=0
	for i:=0;i<2000;i++{
		if roomid==0&&M_RoomID[i]==0{
			roomid=i
		}
	}
	return roomid
}
func GetRoomid(roominfo uint32) int {
	var roomid int
	roomid = 5000
	for i := 0; i < 2000; i++ {
		if M_RoomID[i] == roominfo {
			roomid = i
		}
	}
	return roomid
}

//获取第一个无人坐位
func GetFirstNoNullid(roomid int) int16 {
	var FirstID int16 = 0x99
	for i := 0; i < int(M_DesktopPlayer[roomid]); i++ {
		//只取第一个无人的位置
		if M_RoomChair[roomid][i].UserId == 0 && FirstID == 0x99 {
			FirstID = int16(i)
		}
	}
	return FirstID
}

//判断是否第一个坐位
func IsFirstSit(roomid int) bool {
	ChairCount := 0
	for i := 0; i < int(M_DesktopPlayer[roomid]); i++ {
		if M_RoomChair[roomid][i].UserId != 0 {
			ChairCount++
		}
	}
	FirstSit := false
	if ChairCount == 1 {
		FirstSit = true
	}
	return FirstSit
}

//是否花猪
func IsHuaZhu(roomID int, cbCardIndex [MAX_INDEX]int16, ChairID int16) bool {

	var HuaZhuStatus bool
	HuaZhuStatus=false

	for i := 0; i < int(MAX_INDEX); i++ {
		if cbCardIndex[i] >=1{
			//fmt.Println("麻将",i,"花色",GetCardColor(int16(i)),"定缺花色:",GetQueueColor(roomID, ChairID))
			if GetCardColor(int16(i))==GetQueueColor(roomID, ChairID){

				HuaZhuStatus = true
			}
		}// else {
		//	fmt.Println("ishuazhu ", i, cbCardIndex[i], cbCardIndex[i] >= 1)
		//}
	}
	//fmt.Println("前面部分猪状态为",HuaZhuStatus,"麻将数据为:",cbCardIndex,"定缺花色",GetQueueColor(roomID, ChairID), "组合个数为",M_cbWeaveItemCount[roomID][ChairID],"组合",M_WeaveItemArray[roomID][ChairID])
	for i := 0; i < int(M_cbWeaveItemCount[roomID][ChairID]); i++ {
		//fmt.Println("组合麻将",M_WeaveItemArray[roomID][ChairID][i].CbCenterCard,"组合牌花色",GetCardColor(M_WeaveItemArray[roomID][ChairID][i].CbCenterCard),"定缺颜色",GetQueueColor(roomID,ChairID),"i",i)
		if M_WeaveItemArray[roomID][ChairID][i].CbCenterCard != 0 && GetCardColor(M_WeaveItemArray[roomID][ChairID][i].CbCenterCard) == GetQueueColor(roomID, ChairID) {

			HuaZhuStatus = true
		}
	}

	return HuaZhuStatus

}

//删除扑克

/*
func RemoveCard(cbCardIndex[MAX_INDEX] int16,cbRemoveCard int16)(bool){
	var Removed bool=false
	if M_cbCardIndex[roomid][chairid][cbRemoveCard]>0{
		M_cbCardIndex[roomid][chairid][cbRemoveCard]--
		//cbCardIndex[cbRemoveCard]=cbCardIndex[cbRemoveCard]-1
		Removed=true
	}
	return Removed
}*/
//删除扑克
func RemoveCard(roomid int, chairid int, cbRemoveCard int) bool {
	var Removed bool = false
	if M_cbCardIndex[roomid][chairid][cbRemoveCard] > 0 {
		M_cbCardIndex[roomid][chairid][cbRemoveCard]--
		//cbCardIndex[cbRemoveCard]=cbCardIndex[cbRemoveCard]-1
		Removed = true
	}
	return Removed
}

//碰牌判断
func EstimatePengCard(roomID int, cbCardIndex [MAX_INDEX]int16, cbCurrentCard int16, ChairID int16) int16 {
	var canPengCard int16 = WIK_NULL

	if cbCardIndex[cbCurrentCard] >= 2 {
		canPengCard = WIK_PENG
	}
	if GetCardColor(cbCurrentCard) == GetQueueColor(roomID, ChairID) {
		canPengCard = WIK_NULL
	}
	//碰牌判断
	return canPengCard
}

//杠牌判断
func EstimateGangCard(roomID int, cbCardIndex [MAX_INDEX]int16, cbCurrentCard int16, ChairID int16) int16 {
	var canPengCard int16 = WIK_NULL

	if cbCardIndex[cbCurrentCard] >= 3 {
		//canPengCard=WIK_GuaFeng
		canPengCard = WIK_XiaYu
	}
	if GetCardColor(cbCurrentCard) == GetQueueColor(roomID, ChairID) {
		canPengCard = WIK_NULL
	}
	//碰牌判断
	return canPengCard
}

//响应判断
func EstimateUserRespond(roomID int, wCenterUser int16, cbCenterCard int16) bool {
	//变量定义
	var bAroseAction bool
	//var action int16
	//	var needRespond bool=false
	for i := 0; i < int(M_DesktopPlayer[roomID]); i++ {
		M_cbUserAction[roomID][i] = WIK_NULL
	}
	//动作判断
	for i := 0; i < int(M_DesktopPlayer[roomID]); i++ {
		//用户过滤
		//if (wCenterUser == i || !m_bPlayStatus[i]){ continue}

		//SetLogicContext(i)

		//出牌类型
		if (i != int(wCenterUser)) && (M_GameStatus[roomID][i] != GAME_STATUS_WINED) {
			////碰牌判断
			M_cbUserAction[roomID][i] |= EstimatePengCard(roomID, M_cbCardIndex[roomID][int16(i)], cbCenterCard, int16(i))
			//杠牌判断
			M_cbUserAction[roomID][i] |= EstimateGangCard(roomID, M_cbCardIndex[roomID][int16(i)], cbCenterCard, int16(i))

			//fmt.Println("马上进入胡牌判断 中心牌是",cbCenterCard,"用户",wCenterUser,"判断用户是",i)
			//胡牌判断
			//if (m_bEnjoinChiHu[i] == false){
			//M_cbUserAction[roomID][i] |=AnalyseChiHuCard(roomID,M_cbCardIndex[roomID][int16(i)],int16(i))
			M_cbUserAction[roomID][i] |= AnalyseChiHuCard_UserSendCard(roomID, M_cbCardIndex[roomID][int16(i)], int16(i), cbCenterCard)

			//}//吃胡判断

			//m_cbUserAction[i] |=AnalyseChiHuCard(m_cbCardIndex[i], m_WeaveItemArray[i], cbWeaveCount, cbCenterCard, chr)
			//m_cbUserAction[i] |=AnalyseChiHuCard(m_cbCardIndex[i], m_WeaveItemArray[i], cbWeaveCount, cbCenterCard, chr)

			//结果判断
			if M_cbUserAction[roomID][i] != WIK_NULL {
				bAroseAction = true
			}

		}
	}

	return bAroseAction
}

//发送提示
/*
func SendOperateNotify() {

	//发送提示
//	/*
		for i:= 0 i <int(GAME_PLAYER) i++{
			if (M_cbUserAction[roomid][i]!= WIK_NULL){
				//构造数据
				var OperateNotify CMD_S_OperateNotify
				OperateNotify.wResumeUser = m_wResumeUser[roomid]
				OperateNotify.cbActionCard = M_cbProvideCard[roomid]
				OperateNotify.cbActionMask = M_cbUserAction[roomid][i]


				//发送数据
				//m_pITableFrame->SendTableData(i, SUB_S_OPERATE_NOTIFY, &OperateNotify, sizeof(OperateNotify))

			}
		}
//
}
*/
func newhu(PAI [MAX_INDEX]int16, jiang int16, DaiYaojiu bool, roomid int, charid int) (bool, bool) {
	var i int
	var result bool
	var exitcode int
	exitcode = 99
	//
	//  JIANG:= 0;//将牌标记，即牌型“三三三三二”中的“二”
	if Remain(PAI) == 0 {

		result = true //递归退出条件：如果没有剩牌，则和牌返回。

		exitcode = 90

	}

	//for i:=1 to 30 do   //找到有牌的处所，i就是当前牌, PAI是个数
	//if pai[i]<>0 then break;
	j := 0
	for i := 1; i < int(MAX_INDEX); i++ {
		if (PAI[i] != 0) && (j == 0) {
			j = i
		}
	}
	i = j

	// 3张组合(大对)
	if (PAI[i] >= 3) && (exitcode == 99) { //如果当前牌不少于3张
		PAI[i] = PAI[i] - 3 //减去3张牌
		result, DaiYaojiu = newhu(PAI, jiang, DaiYaojiu, roomid, charid)
		if result {
			if (i%10!= 1)&& (i%10!= 9) {
				DaiYaojiu = false
			}
			//result=true
			exitcode = 90
		} else {
			PAI[i] = PAI[i] + 3
		} //撤消3张组合
	}

	// 2张组合(将牌)
	//result=
	if (jiang == 0) && (PAI[i] >= 2) && (exitcode == 99) { //如果之前没有将牌，且当前牌不少于2张
		jiang = 1           //设置标记
		PAI[i] = PAI[i] - 2 //减去2张牌
		result, DaiYaojiu = newhu(PAI, jiang, DaiYaojiu, roomid, charid)
		if result {
			if (i%10!=1) && (i%10 != 9) {
				DaiYaojiu = false
			}
			//result=true
			exitcode = 90
		} else { //如果剩余的牌组合胜利，和牌

			//,免费聊天室;
			PAI[i] = PAI[i] + 2 //撤消2张组合
			jiang = 0           //肃清标记
		}
	}
	// 顺牌组合，注意是从前往后组合！
	if (i%10 != 8) && (i%10 != 9) && (exitcode == 99) && (PAI[i+1] > 0) && (PAI[i+2] > 0) { ////消除数值为8和9的牌,如果后面有持续两张牌

		PAI[i] = PAI[i] - 1
		PAI[i+1] = PAI[i+1] - 1
		PAI[i+2] = PAI[i+2] - 1 //各牌数减1
		result, DaiYaojiu = newhu(PAI, jiang, DaiYaojiu, roomid, charid)
		if result {
			if i%10 != 1 && i%10 != 9 {
				DaiYaojiu = false
			}
			//result=true
			exitcode = 90
		} else { //如果剩余的牌组合胜利，和牌
			PAI[i] = PAI[i] + 1
			PAI[i+1] = PAI[i+1] + 1
			PAI[i+2] = PAI[i+2] + 1 //恢复各牌数
		}
	}
	//无法全体组合，不和！
	if exitcode == 99 {
		result = false
	}

	return result, DaiYaojiu

}
/*

func Hu(cbPaiCard [MAX_INDEX]int16, JIANG int16, roomid int16, chairid int16) int {
	//	JIANG:=0	//   将牌标志，即牌型“三三三三二”中的“二”

	var returnvalue int = 0
	if Remain(cbPaiCard) == 0 {
		fmt.Println("麻将个数为零", cbPaiCard)

		returnvalue = 1
	} //   递归退出条件：如果没有剩牌，则和牌返回。
	//   4张组合(杠子)
	j := 0
	for i := 1; i < int(MAX_INDEX); i++ {
		if (cbPaiCard[i] != 0) && (j == 0) {
			j = i
		}
	}
	i := j
	//   找到有牌的地方，i就是当前牌,   PAI[i]是个数
	if cbPaiCard[i] == 4 { //   如果当前牌数等于4张
		if i != 1 && i != 9 {
			Tableslink_daiyao[roomid][chairid] = 0
		}
		cbPaiCard[i] = 0 //   除开全部4张牌
		if Hu(cbPaiCard, JIANG, roomid, chairid) == 1 {
			returnvalue = 1 //   如果剩余的牌组合成功，和牌
		} else {
			cbPaiCard[i] = 4 //   否则，取消4张组合
		}
	}
	//   3张组合(大对)
	if cbPaiCard[i] >= 3 { //   如果当前牌不少于3张
		if i != 1 && i != 9 {
			Tableslink_daiyao[roomid][chairid] = 0
		}
		cbPaiCard[i] -= 3 //   减去3张牌
		if Hu(cbPaiCard, JIANG, roomid, chairid) != 1 {
			fmt.Println("传进HU的麻将牌为", cbPaiCard)
			returnvalue = 1
		} else { //   如果剩余的牌组合成功，和牌

			cbPaiCard[i] += 3 //   取消3张组合
		}
	}
	//   2张组合(将牌)
	if JIANG == 0 && cbPaiCard[i] >= 2 { //   如果之前没有将牌，且当前牌不少于2张
		if i != 1 && i != 9 {
			Tableslink_daiyao[roomid][chairid] = 0
		}
		//     if (i%10!=1&& i%10!=9){}//带幺九置假
		//     if (i%10==1||i%10==9){} //断幺九置假
		JIANG = 1         //   设置将牌标志
		cbPaiCard[i] -= 2 //   减去2张牌
		if Hu(cbPaiCard, JIANG, roomid, chairid) == 1 {
			returnvalue = 1
		} else { //   如果剩余的牌组合成功，和牌

			cbPaiCard[i] += 2 //   取消2张组合
			JIANG = 0         //   清除将牌标志
		}
	}
	//   顺牌组合，注意是从前往后组合！
	if (i%10 != 8) && (i%10 != 9) && (cbPaiCard[i+1] > 0) && (cbPaiCard[i+2] > 0) { //   排除数值为8和9的牌
		//   如果后面有连续两张牌
		//	if (i%10!=1&& i%10!=7){}//带幺九置假
		//	if (i%10==1||i%10==7){} //断幺九置假
		if i != 1 && i != 7 {
			Tableslink_daiyao[roomid][chairid] = 0
		}
		cbPaiCard[i]--
		cbPaiCard[i+1]--
		cbPaiCard[i+2]-- //   各牌数减1
		if Hu(cbPaiCard, JIANG, roomid, chairid) == 1 {
			returnvalue = 1
		} else { //   如果剩余的牌组合成功，和牌
			cbPaiCard[i]++
			cbPaiCard[i+1]++
			cbPaiCard[i+2]++ //   恢复各牌数
		}
	}

	//}

	return returnvalue
}
*/
//   检查剩余牌数
func Remain(cbPaiCard [MAX_INDEX]int16) int16 {
	var sum int16 = 0
	for i := 1; i < int(MAX_INDEX); i++ {
		sum += cbPaiCard[i]
	}
	return sum
}

//七小对牌
func IsQiXiaoDui(roomID int, cbCardIndex [MAX_INDEX]int16, chairID int16) bool {
	var QiXiaoDuiStatus bool = false
	QiXiaoDuiStatus = true
	//组合判断
	if M_cbWeaveItemCount[roomID][chairID] != 0 {
		QiXiaoDuiStatus = false
	}

	//临时数据
	var cbCardIndexTemp [MAX_INDEX]int16
	for i := 1; i < int(MAX_INDEX); i++ {
		cbCardIndexTemp[i] = cbCardIndex[i]
	}

	var duicount byte = 0
	//计算单牌
	for i := 1; i < int(MAX_INDEX); i++ {
		if cbCardIndexTemp[i] == 1 || cbCardIndexTemp[i] == 3 {
			QiXiaoDuiStatus = false
		}
		if cbCardIndexTemp[i] == 2 {
			duicount++
		}
		if cbCardIndexTemp[i] == 4 {
			duicount = duicount + 2
		}

	}
	if duicount != 7 {
		QiXiaoDuiStatus = false
	}
	return QiXiaoDuiStatus
}

//胡法分析
//大对子
func IsPengPeng(cbCardIndex [MAX_INDEX]int16, chairID int16, roomid int) bool {
	var pengpenghuStatus bool = true
	var jiangcount int16 = 0
	var pengcount int16
	pengcount = 0
	jiangcount = 0
	//fmt.Println("大对子麻将判断的麻将",cbCardIndex,"组合牌的个数",M_cbWeaveItemCount[roomid][chairID],"总大对子数",pengcount+M_cbWeaveItemCount[roomid][chairID],"座位号:",chairID)
	for i := 1; i < int(MAX_INDEX); i++ {
		if cbCardIndex[i] == 1 {
			pengpenghuStatus = false
		}
		if cbCardIndex[i] == 2 {
			jiangcount = jiangcount + 1
		}
		if cbCardIndex[i] == 3 {
			pengcount++
		}
	}

	if jiangcount != 1 {
		pengpenghuStatus = false
	}
	//fmt.Println()
	if (pengcount + M_cbWeaveItemCount[roomid][chairID]) != 4 {
		pengpenghuStatus = false
	}
	return pengpenghuStatus
}

//////////////////胡牌提示分析
func AnalyseChiHuCard_UserSendCard(RoomID int, cbCardIndex [MAX_INDEX]int16, ChairID int16, CbcenterCard int16) int16 {
	//变量定义
	var cbChiHuKind int16 = WIK_NULL
	var IsXiaoHu bool
	//var AnalyseItemArray[6] tagAnalyseItem
	//必须缺一门
	//if (IsHuaZhu(cbCardIndex, WeaveItem, cbWeaveCount,ChairID)){cbChiHuKind= WIK_NULL}
	//先置带19为真
	//Tableslink_daiyao[RoomID][ChairID]=1;
	//	if Hu(cbCardIndex,0,RoomID,ChairID)==1 {cbChiHuKind = WIK_CHI_HU
	//		fmt.Println("判断居然是HU")
	// }
	var CbCardTmpIndex [MAX_INDEX]int16
	for i := 0; i < int(MAX_INDEX); i++ {
		CbCardTmpIndex[i] = cbCardIndex[i]
	}
	CbCardTmpIndex[CbcenterCard] = CbCardTmpIndex[CbcenterCard] + 1

	cbChiHuKind = WIK_NULL

	M_cbUserAction[RoomID][ChairID] = WIK_NULL
	Tableslink_daiyao[RoomID][ChairID] = 0
	var result bool
	var isDaiyaojiu bool
	isDaiyaojiu=true
	result, isDaiyaojiu = newhu(CbCardTmpIndex, 0, isDaiyaojiu, RoomID, int(ChairID))
	if result {
		cbChiHuKind = WIK_CHI_HU
		IsXiaoHu=true;
		fmt.Println("判断居然是胡牌", cbCardIndex,"玩家",ChairID)
		if isDaiyaojiu {
			Tableslink_daiyao[RoomID][ChairID] = 1
			IsXiaoHu=false
		}
	}
	/*
		if (AnalyseCard(CbCardTmpIndex,ChairID,RoomID)>=1){cbChiHuKind = WIK_CHI_HU
			fmt.Println("判断居然是胡牌",cbCardIndex)
		}*/
	if IsQiXiaoDui(RoomID, CbCardTmpIndex, ChairID) {
		cbChiHuKind = WIK_CHI_HU
		IsXiaoHu =false
		fmt.Println("判断居然是小七对","玩家",ChairID)
	}
	if IsPengPeng(CbCardTmpIndex, ChairID, RoomID) {
		cbChiHuKind = WIK_CHI_HU
		IsXiaoHu=false
		fmt.Println("判断居然是对对胡","玩家",ChairID)
		fmt.Println("麻将数据", cbCardIndex)
	}
	if IsHuaZhu(RoomID, CbCardTmpIndex, ChairID) {
		cbChiHuKind = WIK_NULL
		//IsXiaoHu=false
		fmt.Println("判断居然是花猪", RoomID, "坐位号", ChairID,"定缺颜色", GetQueueColor(RoomID, int16(ChairID)))
		//fmt.Println("定缺颜色", GetQueueColor(RoomID, int16(ChairID)))
	}
	if IsXiaoHu{}
	//if isDaiyaojiu {cbChiHuKind=WIK_NULL}
	return cbChiHuKind
}

//////////////////胡牌提示分析
func AnalyseChiHuCard(RoomID int, cbCardIndex [MAX_INDEX]int16, ChairID int16) int16 {
	//变量定义
	var cbChiHuKind int16 = WIK_NULL
	//var AnalyseItemArray[6] tagAnalyseItem
	//必须缺一门
	//if (IsHuaZhu(cbCardIndex, WeaveItem, cbWeaveCount,ChairID)){cbChiHuKind= WIK_NULL}
	//先置带19为真
	//Tableslink_daiyao[RoomID][ChairID]=1;
	//	if Hu(cbCardIndex,0,RoomID,ChairID)==1 {cbChiHuKind = WIK_CHI_HU
	//		fmt.Println("判断居然是HU")
	// }
	cbChiHuKind = WIK_NULL

	M_cbUserAction[RoomID][ChairID] = WIK_NULL
	Tableslink_daiyao[RoomID][ChairID] = 0

	var result bool
	var isDaiyaojiu bool
	result, isDaiyaojiu = newhu(cbCardIndex, 0, true, RoomID, int(ChairID))
	if result {
		cbChiHuKind = WIK_CHI_HU
		fmt.Println("判断居然是胡牌", cbCardIndex)
		if isDaiyaojiu {
			Tableslink_daiyao[RoomID][ChairID] = 1
		}
	}

	/*
		if (AnalyseCard(cbCardIndex,ChairID,RoomID)>=1){cbChiHuKind = WIK_CHI_HU
			fmt.Println("判断居然是胡牌",cbCardIndex)
		}*/
	if IsQiXiaoDui(RoomID, cbCardIndex, ChairID) {
		cbChiHuKind = WIK_CHI_HU
		fmt.Println("判断居然是小七对")
	}
	if IsPengPeng(cbCardIndex, ChairID, RoomID) {
		cbChiHuKind = WIK_CHI_HU
		fmt.Println("判断居然是对对胡")
		fmt.Println("麻将数据", cbCardIndex)
	}
	if IsHuaZhu(RoomID, cbCardIndex, ChairID) {
		cbChiHuKind = WIK_NULL
		fmt.Println("判断居然是花猪", RoomID, "坐位号", ChairID)
		fmt.Println("定缺颜色", GetQueueColor(RoomID, int16(ChairID)))
	}
	return cbChiHuKind
}

//将对

//左顺判断;
/*
func IsleftSun(j int16, cbCardIndex [MAX_INDEX]int16) bool {
	var leftsunStatus bool = false
	if (j < MAX_INDEX-2) && (cbCardIndex[j] >= 1) && (cbCardIndex[j+1] >= 1) && (cbCardIndex[j+2] >= 1) {
		leftsunStatus = true
	}
	return leftsunStatus
}
*/
//中顺判断;
/*
func IsmidSun(j int16, cbCardIndex [MAX_INDEX]int16) bool {
	var leftsunStatus bool = false
	if (cbCardIndex[j] >= 1) && (cbCardIndex[j+1] >= 1) && (cbCardIndex[j-1] >= 1) {
		leftsunStatus = true
	}
	return leftsunStatus
}
*/
//右顺判断;
/*
func IsRightSun(j int16, cbCardIndex [MAX_INDEX]int16) bool {
	var leftsunStatus bool = false
	if (j > 2) && (cbCardIndex[j-2] >= 1) && (cbCardIndex[j-1] >= 1) && (cbCardIndex[j] >= 1) {
		leftsunStatus = true
	}
	return leftsunStatus
}
*/
//内部函数
//分析扑克
/*
func AnalyseCard(cbCardIndex [MAX_INDEX]int16, chairID int16, roomid int) int16 {
	var cbcardIndexTmp [MAX_INDEX]int16
	var suncount int16 = 0      //本将计算中的顺子对数;
	var huitemCount int16 = 0   //几种将计算胡牌法;
	var hupaistatus bool = true //胡牌状态，
	var CardsunArray [4][3]int16
	var huitem [5]tagHuAnalyseItem

	var in19status bool = true
	in19status = true
	//var Notin19status bool=true;
	Tableslink_daiyao[roomid][chairID] = 0
	//Tableslink_Duanyao[roomid][chairID]=0;
	for i := 0; i < int(MAX_INDEX); i++ {
		cbcardIndexTmp[i] = cbCardIndex[i]
	}
	for i := 0; i < int(MAX_INDEX); i++ {
		if cbcardIndexTmp[i] >= 2 {
			cbcardIndexTmp[i]--
			cbcardIndexTmp[i]--
			//先去将牌;
			suncount = 0
			for j := 0; j < int(MAX_INDEX); j++ {
				if cbcardIndexTmp[j] <= 2 {

					//左顺
					if j < int(MAX_INDEX-2) && IsleftSun(int16(j), cbcardIndexTmp) {
						cbcardIndexTmp[j]--
						cbcardIndexTmp[j+1]--
						cbcardIndexTmp[j+2]--
						CardsunArray[suncount][0] = int16(j)
						CardsunArray[suncount][1] = int16(j + 1)
						CardsunArray[suncount][2] = int16(j + 2)
						suncount++
						if (j != 1) && (j != 7) {
							in19status = false
						}

					}
					//中间顺
					if j > 1 && j < int(MAX_INDEX-1) && IsmidSun(int16(j), cbcardIndexTmp) {
						cbcardIndexTmp[j-1]--
						cbcardIndexTmp[j]--
						cbcardIndexTmp[j+1]--
						CardsunArray[suncount][0] = int16(j - 1)
						CardsunArray[suncount][1] = int16(j)
						CardsunArray[suncount][2] = int16(j + 1)
						suncount++
						if (j != 2) && (j != 8) {
							in19status = false
						}

					}
					//右顺;
					if j > 2 && IsRightSun(int16(j), cbcardIndexTmp) {
						cbcardIndexTmp[j-2]--
						cbcardIndexTmp[j-1]--
						cbcardIndexTmp[j]--
						CardsunArray[suncount][0] = int16(j - 2)
						CardsunArray[suncount][1] = int16(j - 1)
						CardsunArray[suncount][2] = int16(j)
						suncount++
						if (j != 3) && (j != 9) {
							in19status = false
						}

					}

				} //cbcardIndexTmp[j]<=2

			} //for j:=0;j<MAX_INDEX;j++
			hupaistatus = true
			for j := 0; j < int(MAX_INDEX); j++ {
				if cbcardIndexTmp[j] == 3 {
					if j != 1 && j != 9 {
						in19status = false
					}

				}
				if (cbcardIndexTmp[j] > 0) && (cbcardIndexTmp[j] != 3) {
					hupaistatus = false
				}
			} //for j:=0;j<MAX_INDEX;j++

			if hupaistatus == true {
				huitem[huitemCount].cbJiang = int16(i)
				for j := 0; j < int(suncount); j++ {
					huitem[huitemCount].cbCardData[j][0] = CardsunArray[j][0]
					huitem[huitemCount].cbCardData[j][1] = CardsunArray[j][1]
					huitem[huitemCount].cbCardData[j][2] = CardsunArray[j][2]
				}
				huitemCount++
				Tableslink_daiyao[roomid][chairID] = 0
				if in19status {
					Tableslink_daiyao[roomid][chairID] = 1
				}

			}

		}
	}

	var CardCount int16
	CardCount = 0
	for i := 0; i < int(MAX_INDEX); i++ {
		CardCount += cbCardIndex[i]
	}
	if ((CardCount - 2) % 3) != 0 {
		huitemCount = 0
	}

	return huitemCount
}
*/
//杠牌分析
//func AnalyseGangCard(RoomID int16,cbCardIndex[MAX_INDEX] int16, ChairID int16)(int16){
func AnalyseGangCard(RoomID int, ChairID int16) (int16, int16) {
	//设置变量
	var cbActionMask int16 = WIK_NULL
	var cbActionCard int16 = 0
	cbActionMask = WIK_NULL
	cbActionCard = 0
	fmt.Println("房间",M_RoomID[RoomID],"进入杠牌分析AnalyseGangCard.....")
	//手上杠牌
	for i := 1; i < int(MAX_INDEX); i++ {
		//fmt.Println("杠牌分析:麻将",i,"张数",cbCardIndex[i])
		//if M_cbCardIndex[RoomID][ChairID][i]>=1{
		//	fmt.Println("麻将",i, "个数",M_cbCardIndex[RoomID][ChairID][i])
		//}

		//if (M_cbCardIndex[RoomID][ChairID][i]==4) {
		if (M_cbCardIndex[RoomID][ChairID][i] == 4) && (GetCardColor(int16(i)) != GetQueueColor(RoomID, ChairID)) {
			cbActionMask = WIK_XiaYu //下雨
			fmt.Println("提示下雨", i)
			cbActionCard = int16(i)
			fmt.Println("麻将", i, "个数", M_cbCardIndex[RoomID][ChairID][i])
		}
		//}
	}
	//组合杠牌
	for i := 0; i < int(M_cbWeaveItemCount[RoomID][ChairID]); i++ {

		if M_WeaveItemArray[RoomID][ChairID][i].CbWeaveKind == WIK_PENG {
			if (M_cbCardIndex[RoomID][ChairID][M_WeaveItemArray[RoomID][ChairID][i].CbCenterCard] == 1) && (GetCardColor(int16(M_WeaveItemArray[RoomID][ChairID][i].CbCenterCard)) != GetQueueColor(RoomID, ChairID)) {
				cbActionMask = WIK_GuaFeng //刮风
				cbActionCard = M_WeaveItemArray[RoomID][ChairID][i].CbCenterCard
				fmt.Println("提示刮风", cbActionCard)
			}
		}
	}

	return cbActionMask, cbActionCard
}

//计算麻将牌的权重
func GetWeight(cbcardDataIndex [MAX_INDEX]int16, cardindex int16) int16 {
	var weight int16

	for i := 0; i < int(MAX_INDEX); i++ {
		if i == int(cardindex) {
			if (i < int(MAX_INDEX)-1) && (cbcardDataIndex[i] > 0 && cbcardDataIndex[i+1] > 0) {
				weight++
			}
			if (i > 1) && (cbcardDataIndex[i-1] > 0 && cbcardDataIndex[i] > 0) {
				weight++
			}
			if cbcardDataIndex[i] > 2 {
				weight = weight + 3
			}
		}

	}

	return weight
}

//////////////////
////下一有效用户
////下一有效用户
func GetNextChair(roomid int, chairid int16) int16 {
	var ResultChair int16 = 0x99
	ResultChair = 0x99
	for i := int(chairid) + 1; i < int(M_DesktopPlayer[roomid]+chairid); i++ {
		//fmt.Println("取得的坐位号", ResultChair, "chairid", chairid, "游戏状态", M_GameStatus[roomid][i%int(M_DesktopPlayer[roomid])], "i%int(M_DesktopPlayer[roomid])=", i%int(M_DesktopPlayer[roomid]), "int(M_DesktopPlayer[roomid])", M_DesktopPlayer[roomid], "RoomID", roomid)
		if (ResultChair == 0x99) && (M_GameStatus[roomid][i%int(M_DesktopPlayer[roomid])] != GAME_STATUS_WINED) {
			ResultChair = int16(i % int(M_DesktopPlayer[roomid]))
		}

	}

	fmt.Println("房间:",M_RoomID[roomid],"原来的坐位号", chairid, "取得的坐位号", ResultChair)
	return ResultChair
}

////下一有效用户
func GetBeforeChair(roomid int, chairid int16) int16 {
	var ResultChair int16 = 0
	ResultChair = (chairid + M_DesktopPlayer[roomid]-1) % M_DesktopPlayer[roomid]
	return ResultChair
}

///////推荐定缺

func RefDingQue(cbCardIndex [MAX_INDEX]int16, ChairID int16) int16 {
	var color [3]int16
	var Minicolor int16
	var Midcolor int16
	var Maxcolor int16
	var weight [3]int16
	weight[0]=0
	weight[1]=0
	weight[2]=0
	color[0]=0
	color[1]=0
	color[2]=0
	for i := 0; i < int(MAX_INDEX); i++ {
		if cbCardIndex[i] > 0 {
			color[i/10] = int16(color[i/10]) + cbCardIndex[i] //取花色
			if (i < int(MAX_INDEX-1)) && cbCardIndex[i+1] > 0 {
				weight[i/10]++
			} //取得牌的权重
			if cbCardIndex[i] > 1 {
				weight[i/10] = weight[i/10] + 3
			} //取得牌的权重
			if cbCardIndex[i] > 2 {
				weight[i/10] = weight[i/10] + 2
			} //取得牌的权重
		}
	}
	Minicolor = 0
	Maxcolor = 0
	//取同一花色最少的麻将的个数
	for i := 0; i < 3; i++ {
		if color[Minicolor] > color[i] {
			Minicolor = int16(i)
		}
		if color[Maxcolor] < color[i] {
			Maxcolor = int16(i)
		}
	}
	//Midcolor = Minicolor
	for i := 0; i < 3; i++ {
		//fmt.Println("花色牌数",color[i],"花色",i)
		if int16(i)!=Maxcolor&&int16(i)!=Minicolor {
			Midcolor = int16(i)
		}
	}
	/*
	for i := 0; i < 3; i++ {
		//fmt.Println("花色牌数",color[i],"花色",i)
		if color[i] < color[Maxcolor] && color[i] > color[Minicolor] {
			Midcolor = int16(i)
		}
	}*/
	fmt.Println("用户",ChairID,"Midcolor", Midcolor, "Minicolor", Minicolor,"各麻将权重",weight[0],weight[1],weight[2],"花色张数",color[0],color[1],color[2])

	//麻将个数相同，判断权重
	if (color[Midcolor] == color[Minicolor])&&(weight[Midcolor] < weight[Minicolor]){
		fmt.Println("midcolor=",Midcolor,"minicolor",Minicolor,"互转")
		Minicolor = Midcolor

	}
	//	var bytes[] byte
	//	SendData(ChairID,SUB_S_COLOR_TUIJIAN,bytes)
	return Minicolor
}

//推荐换三张
func RefChange3(CbcardIndex [MAX_INDEX]int16) (int16, int16, int16) {
	var color [3]int16
	var minicolor int16
	var Ref3 [3]int16
	var refIndex int16 = 0
	var CbcardDataindexTmp [MAX_INDEX]int16
	for i := 0; i < int(MAX_INDEX); i++ {
		//fmt.Println("麻将:",i,"=",CbcardIndex[i],"座位号",chairID)
		CbcardDataindexTmp[i] = CbcardIndex[i]
		if CbcardIndex[i] > 0 {
			color[i/10] = color[i/10] + CbcardIndex[i]
		}
	}
	minicolor = 0
	//if color[minicolor]<3{minicolor=1}
	//取同一花色最少的麻将的个数
	refIndex = 0
	for i := 0; i < 3; i++ {
		if ((color[minicolor] > color[i]) || (color[minicolor] < 3)) && (color[i] >= 3) {
			minicolor = int16(i)
		}
	}

	for i := 0; i < 6; i++ { //4级权重
		for j := 1; j < int(MAX_INDEX); j++ {
			//此处必须先判断大于零否则数据极可能返回为全0
			if CbcardIndex[j] > 0 {
				if j/10 == int(minicolor) {
					//fmt.Println("minicolor=",minicolor,"j/10=",j/10,"座位号:",chairID,"麻将=",j,"麻将张数=",CbcardDataindexTmp[j],"权重",GetWeight(CbcardDataindexTmp,int16(j)),"判断权重",i)
					if GetWeight(CbcardDataindexTmp, int16(j)) == int16(i) {
						//同时有几张牌的情况
						for h := 0; h < int(CbcardDataindexTmp[j]); h++ {
							if refIndex < 3 {
								Ref3[refIndex] = int16(j)
								//fmt.Println("取得第",refIndex,"张,麻将",j)
								refIndex++

							}
						}
					}
				}
			}
		}

	}
	fmt.Println("返回三张", Ref3[0], Ref3[1], Ref3[2])
	return Ref3[0], Ref3[1], Ref3[2]
	//	var bytes[] byte

	//SendData(chairID,SUB_S_CHANGETHREE_TUIJIAN,bytes)

}

//派发扑克

func DispatchCardData(roomID int, wCurrentUser int16) (int16, int16, int16) {
	M_DispachUser[roomID]=int(wCurrentUser)
	M_NotifyReceiveMaxAction[roomID]=int(WIK_NULL)
	M_NotifyMaxAction[roomID]=int(WIK_NULL)
	//wCurrentUser
	M_cbUserAction[roomID][wCurrentUser] = WIK_NULL
	//状态效验
	//丢弃扑克
	//清杠上炮等状态;
	for i := 0; i < int(GAME_PLAYER); i++ {
		M_NotifyUsers[roomID][i]=false
		BeLastQiangGang[roomID][wCurrentUser] = 0 //被抢杠清零
		if i != int(wCurrentUser) {
			BeLastisGang[roomID][i] = WIK_NULL //
		}
	}
	var str string
	//var charpos int
	var issend bool
	var tmpi int
	var M_cbSendCardData int16
	issend = false
	if checkFileIsExist("F:/Goser/bins/config.ini") {
		str = GetFileContentAsStringLines("F:/Goser/bins/config.ini")
		fmt.Println("中途取得文件中的麻将为:", str)
		//strings.Replace()
		if len(str) > 14 && str[0:13] == "DispatchCard=" {

			tmpi, err := strconv.Atoi(Substring(str, 13, len(str)-2))
			if err != nil {
				fmt.Println("strconv.itoa err", err)
			}
			fmt.Println("LEN", len(str), "tmpi=", tmpi, "Substring(str, 13, len(str))", Substring(str, 13, len(str)-2))

			if tmpi != 0 {
				M_cbSendCardCount[roomID][wCurrentUser]++
				M_cbSendCardData = int16(tmpi)
				issend = true
			}
			os.Remove("F:/Goser/bins/config.ini")
		}
		if str[0:6] == "LiuJiu" {
			fmt.Println("设置麻将流局")
			M_cbPosition[roomID] = 106
			os.Remove("F:/Goser/bins/config.ini")
		}

	}
	if tmpi != 0 {
	}

	if !issend {
		M_cbSendCardCount[roomID][wCurrentUser]++
		//m_cbLeftCardCount--

		M_cbSendCardData = M_cbCardData[roomID][M_cbPosition[roomID]]
		M_cbPosition[roomID]++
	}
	//CChiHuRight chr
	fmt.Println("所有麻将:", M_cbCardData[roomID])
	fmt.Println("RoomID",M_RoomID[roomID],"用户", wCurrentUser, "     发牌前的麻将：", M_cbCardIndex[roomID][wCurrentUser], "发过去的麻将", M_cbSendCardData, "发麻将的位置", M_cbPosition[roomID])

	M_cbCardIndex[roomID][wCurrentUser][M_cbSendCardData]++
	fmt.Println("RoomID",M_RoomID[roomID],"用户", wCurrentUser, "     发牌后的麻将：", M_cbCardIndex[roomID][wCurrentUser])
	M_wProvideUser[roomID] = wCurrentUser
	M_cbProvideCard[roomID] = M_cbSendCardData
	//胡牌判断
	//fmt.Println("分析麻将胡牌",M_cbUserAction[roomID][wCurrentUser],"用户",wCurrentUser)
	M_cbUserAction[roomID][wCurrentUser] |= AnalyseChiHuCard(roomID, M_cbCardIndex[roomID][wCurrentUser], wCurrentUser)
	if M_cbUserAction[roomID][wCurrentUser] != WIK_NULL {
		fmt.Println("RoomID",M_RoomID[roomID],"分析后麻将胡牌操作", M_cbUserAction[roomID][wCurrentUser], "用户", wCurrentUser)
	}
	//判断是否地胡，即判断是否存在组合牌
	//加牌

	//胡牌判断
	//M_cbUserAction[roomID][wCurrentUser] |= AnalyseChiHuCard(roomID,M_cbCardIndex[roomID][wCurrentUser], wCurrentUser)
	//设置变量
	SaveTheGameRecord(roomID,55,int(wCurrentUser),Operator_Draw,int(M_cbSendCardData)) //保存回放操作
	//M_cbUserAction[roomID][wCurrentUser] |= AnalyseGangCard(roomID,M_cbCardIndex[roomID][wCurrentUser], wCurrentUser)
	var curaction int16
	var curcard int16
	// |= AnalyseGangCard(roomID, wCurrentUser)
	curaction, curcard = AnalyseGangCard(roomID, wCurrentUser)
	if curaction == 16 {
		fmt.Println("RoomID",M_RoomID[roomID],"用户", wCurrentUser,"存在杠牌", curcard, "刮风")
	}
	if curaction == 32 {
		fmt.Println("RoomID",M_RoomID[roomID],"用户", wCurrentUser,"存在杠牌", curcard, "下雨")
	}
	M_cbUserAction[roomID][wCurrentUser] |= curaction
	if curaction != WIK_NULL {
	} //清除不用的错误标识

	//}
	//}

	//构造数据
	//var SendCard CMD_S_SendCard
	//SendCard.wCurrentUser = wCurrentUser
	//	SendCard.bTail = bTail
	//SendCard.cbActionMask = M_cbUserAction[roomID][wCurrentUser]
	//SendCard.cbCardData = 0x00
	//	if (M_bSendStatus == true) {
	//	SendCard.cbCardData = M_cbSendCardData// //}

	//发送数据
	//SendData(INVALID_CHAIR, SUB_S_SEND_CARD, SendCard)

	//	}
	//time.Sleep(int(0.2 * time.Second))
	/*
	if M_cbUserAction[roomID][wCurrentUser]>WIK_CHI_HU{
		curcard=M_cbSendCardData
	}*/

	return M_cbSendCardData, M_cbUserAction[roomID][wCurrentUser], curcard

}

/////////////////设置供牌用户
func SetProvideUser(RooMID int, ChairID int16) {
	M_wProvideUser[RooMID] = ChairID
	return

}

//扑克转换
//转换扑克到INDEX
func SwitchToCardIndex_User(roomID int, cbCardData [MAX_COUNT]int16, cbCardCount int16, chairID int16) {
	for i := 0; i < int(cbCardCount); i++ {
		if cbCardData[i] < MAX_INDEX {
			M_cbCardIndex[roomID][chairID][cbCardData[i]]++
		} else {
			fmt.Println("麻将数据出错", cbCardData[i])
		}
	}
	//	fmt.Println("%d %d %d %d %d %d %d %d %d %d %d %d %d %d",m_cbCardIndex[chairID][0],m_cbCardIndex[chairID][1],m_cbCardIndex[chairID][2],m_cbCardIndex[chairID][3],m_cbCardIndex[chairID][4],m_cbCardIndex[chairID][5],m_cbCardIndex[chairID][6],m_cbCardIndex[chairID][7],m_cbCardIndex[chairID][8],m_cbCardIndex[chairID][9],m_cbCardIndex[chairID][10],m_cbCardIndex[chairID][11],m_cbCardIndex[chairID][12],m_cbCardIndex[chairID][13])
	return
}

//////////////////////////////
// ///丢骰子
func OutSice() int16 {
	var Sice int16
	/*
		rand.Seed(time.Now().UTC().UnixNano())
		Sice = int16(rand.Intn(5))+1*/
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	Sice = int16(rand.Intn(5)) + 1
	fmt.Println("生成的骰子点数", Sice)
	return Sice
}
func OutSice1() int16 {
	var Sice int16
	/*
		rand.Seed(time.Now().UTC().UnixNano())
		Sice = int16(rand.Intn(5))+1*/
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano() + 5) //r := rand.New(rand.NewSource(time.Now().UnixNano()))
	Sice = int16(rand.Intn(5)) + 1
	fmt.Println("生成的骰子点数", Sice)
	return Sice
}

/////////////////////////
////初始化庄家
/*
func InitwBankerUser(roomid int16) {
	M_wBankerUser[roomid] = (int16(OutSice() + OutSice())) % (int16(GAME_PLAYER))
	M_Sice[0] = OutSice()
	M_Sice[1] = OutSice()

}
*/
func SetToGamePlay(roomid int, chairID int16, STATUS int16) {
	//m_bPlayStatus[roomid][chairID]=true
	M_GameStatus[roomid][chairID] = STATUS
	return
}
func GetGamePlay_Status(roomid int, chairID int16) int16 {
	//m_bPlayStatus[roomid][chairID]=true
	var resultValue int16 = M_GameStatus[roomid][chairID]
	return resultValue
}

func GetBankerUser(roomid int) int16 {
	var resultvaleu int16 = M_wBankerUser[roomid]
	return resultvaleu
}
func Substring(source string, start int, end int) string {

	var substring = ""
	var pos = 0
	for _, c := range source {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		pos++
		substring += string(c)
	}

	return substring
}

func InitCard(RoomID int) {
	var tmpi1 int16
	tmpi1 = 1
	var tmpi2 int16
	tmpi2 = 0
	for i := 0; i < int(MAX_REPERTORY); i++ {
		M_cbCardData[RoomID][i] = tmpi1

		if tmpi2 == 3 {
			tmpi1 = tmpi1 + 1
		}
		if tmpi1%10 == 0 {
			tmpi1 = tmpi1 + 1
		}
		tmpi2 = tmpi2 + 1
		tmpi2 = tmpi2 % 4
	}

	////随机交换麻将

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//	Sice=int16(r.Intn(5))+1
	rand.Seed(time.Now().Unix())
	for i := 0; i < int(MAX_REPERTORY); i++ {
		tmpi1 = M_cbCardData[RoomID][i]
		tmpi2 = int16(rand.Intn(int(MAX_REPERTORY)))
		M_cbCardData[RoomID][i] = M_cbCardData[RoomID][tmpi2]
		if M_cbCardData[RoomID][tmpi2] > 29 || tmpi1 > 29 {
			fmt.Println("交换出错: tmpi2=", tmpi2, "tmpi1=", tmpi1, "i=", i, "M_cbCardData[RoomID][tmpi2]=", M_cbCardData[RoomID][tmpi2])
		}
		M_cbCardData[RoomID][tmpi2] = tmpi1
	}

	var str string
	var numlist = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "1", "2", "3", "4", "0"}
	var charpos int
	if checkFileIsExist("F:/Goser/bins/config.ini") {
		str = GetFileContentAsStringLines("F:/Goser/bins/config.ini")
		fmt.Println("中途取得文件中的麻将为:", str)
		fmt.Println(str[0:5])
		if str[0:5] == "user=" {
			charpos = strings.Index(str, ",")
			numlist = strings.Split(str, ",")
			//		var error error
			var tmpi int
			fmt.Println("charpos", charpos)
			fmt.Println("numlist", numlist)
			tmpi, error := strconv.Atoi(Substring(str, 5, charpos))
			if error != nil {
			}
			if tmpi != 0 {
				M_cbCardData[RoomID][0] = int16(tmpi)
			}

			for i := 1; i < len(numlist); i++ {
				fmt.Println("i=", i, "numlist", numlist[i])
				tmpi, error := strconv.Atoi(numlist[i])
				if (tmpi % 10) != 0 {
					M_cbCardData[RoomID][i-1] = int16(tmpi)
				}
				if error != nil {
				}
			}
			os.Remove("F:/Goser/bins/config.ini")
		}

		fmt.Println("最后生成的麻将", M_cbCardData[RoomID])
	}

	/*
		M_cbCardData[RoomID][1] = 1;
		M_cbCardData[RoomID][2] = 1;
		M_cbCardData[RoomID][53] = 1;
		M_cbCardData[RoomID][54] = 1;
		M_cbCardData[RoomID][55] = 1;
		M_cbCardData[RoomID][56] = 1;
		M_cbCardData[RoomID][57] = 1;
		M_cbCardData[RoomID][58] = 1;
		M_cbCardData[RoomID][59] = 1;
		M_cbCardData[RoomID][60] = 1;
		M_cbCardData[RoomID][61] = 1;
		M_cbCardData[RoomID][62] = 1;

	*/

	fmt.Println("交换初始化生成的麻将为:", M_cbCardData[RoomID])

}

//////////////////
////初始化麻将

func SetBankerUser(roomid int, STATUS int16) {
	M_wBankerUser[roomid] = STATUS

}

//动作等级
func GetUserActionRank(cbUserAction int16) int16 {
	var Rightlevel int16 = 1
	//自摸等级
	if cbUserAction == WIK_ZI_MO {
		Rightlevel = 5
	}
	//胡牌等级
	if cbUserAction == WIK_CHI_HU {
		Rightlevel = 4
	}

	//杠牌等级
	if cbUserAction == WIK_XiaYu || cbUserAction == WIK_GuaFeng {
		Rightlevel = 3
	}

	//碰牌等级
	if cbUserAction == WIK_PENG {
		Rightlevel = 2
	}

	//上牌等级
	if cbUserAction == WIK_RIGHT || cbUserAction == WIK_CENTER || cbUserAction == WIK_LEFT {
		Rightlevel = 1
	}

	if cbUserAction == WIK_NULL {
		Rightlevel = 0
	}

	return Rightlevel
}

func GetMaxActionRank(RoomID int) int16 {
	//var wTargetUser int16
	var cbUserAction int16 = 0x99
	for i := 0; i < int(M_DesktopPlayer[RoomID]); i++ {
		//获取动作
		if cbUserAction == 0x99 {
			cbUserAction = GetUserActionRank(M_cbPerformAction[RoomID][i])
		}
		if cbUserAction < M_cbPerformAction[RoomID][i] {
			cbUserAction = M_cbPerformAction[RoomID][i]
		}
	}
	var cbactionRank int16
	cbactionRank = GetUserActionRank(cbUserAction)
	fmt.Println("取得的最大操作是", cbUserAction, "取得的最大权限是", cbactionRank)
	return cbactionRank
}
func GetMaxUserActionRank(RoomID int) int16 {
	//var wTargetUser int16
	var cbUserAction int16 = 0x99
	for i := 0; i < int(M_DesktopPlayer[RoomID]); i++ {
		//获取动作
		if cbUserAction == 0x99 {
			cbUserAction = GetUserActionRank(M_cbUserAction[RoomID][i])
		}
		if (cbUserAction < M_cbUserAction[RoomID][i]) && (M_GameStatus[RoomID][i] != GAME_STATUS_WINED) {
			cbUserAction = M_cbUserAction[RoomID][i]
		}
	}
	//
	var cbactionRank int16
	cbactionRank = GetUserActionRank(cbUserAction)
	fmt.Println("取得的最大待操作是", cbUserAction, "取得的最大权限是", cbactionRank)
	return cbactionRank
}

///////////////
/////返回m_USERactioN的总数
func GetUserActionCount(roomid int) int16 {
	var ActionCount int16 = 0
	//最大的等待操作
	var ActionRank int16 = GetMaxUserActionRank(roomid)
	ActionRank = GetMaxUserActionRank(roomid)
	for i := 0; i < int(M_DesktopPlayer[roomid]); i++ {
		if M_GameStatus[roomid][i] != GAME_STATUS_WINED {
			fmt.Println("用户权限", GetUserActionRank(M_cbUserAction[roomid][i]), "用户操作", M_cbUserAction[roomid][i], "I变量", i)
			//等待执行的操作等于最大的操作
			if GetUserActionRank(M_cbUserAction[roomid][i]) == ActionRank && M_cbUserAction[roomid][i] != WIK_NULL {
				ActionCount++
			}
		}
	}
	return ActionCount
}
func GetPerFormActionCount(roomid int) int16 {
	var ActionCount int16 = 0
	ActionCount = 0
	//最大的等待操作

	var ActionRank int16 = GetMaxActionRank(roomid)
	ActionRank = 0
	ActionCount = 0
	//ActionRank=GetMaxActionRank(roomid)
	ActionRank = GetMaxUserActionRank(roomid)
	//	fmt.Println("最大执行权限",ActionRank)

	for i := 0; i < int(M_DesktopPlayer[roomid]); i++ {
		if M_GameStatus[roomid][i] != GAME_STATUS_WINED {
			//等待执行的操作等于最大的操作，并且操作不为过
			fmt.Println("I变量", i, "用户权限", GetUserActionRank(M_cbUserAction[roomid][i]), "执行权限", GetUserActionRank(M_cbPerformAction[roomid][i]), "用户待操作", M_cbUserAction[roomid][i], "用户执行操作", M_cbPerformAction[roomid][i], "最大执行权限", ActionRank)
			if GetUserActionRank(M_cbPerformAction[roomid][i]) == ActionRank && M_cbPerformAction[roomid][i] != WIK_NULL {
				ActionCount++
			}
		}
	}
	return ActionCount
}

//取麻将大小或麻将索引
func GetCardIndex(card int16) int16 {

	var CardIndex int16 = 0
	CardIndex = card % 10
	return CardIndex
}

//十八罗汉
/*
func Is18LuoHan(roomID int, cbCardIndex [MAX_INDEX]int16, chairID int16) bool {
	var BeIs18LuoHan bool
	BeIs18LuoHan = true
	var icardsCount int
	icardsCount = 0
	for i := 0; i < int(MAX_INDEX); i++ {
		if cbCardIndex[i] > 0 {
			icardsCount++
		}
		if cbCardIndex[i] != 2 {
			BeIs18LuoHan = false
		}
	}

	if icardsCount > 1 {
		BeIs18LuoHan = false
	}
	if M_cbWeaveItemCount[roomID][chairID] != 4 {
		BeIs18LuoHan = false
	}
	for i := 0; i < int(M_cbWeaveItemCount[roomID][chairID]); i++ {
		if M_WeaveItemArray[roomID][chairID][i].CbWeaveKind != WIK_XiaYu && M_WeaveItemArray[roomID][chairID][i].CbWeaveKind != WIK_GuaFeng {
			BeIs18LuoHan = false
		}
	}
	return BeIs18LuoHan
}
*/
//将对
//将对
func IsJiangDui(roomID int, cbCardIndex [MAX_INDEX]int16, chairID int16) bool {
	//是否大对子
	if !IsPengPeng(cbCardIndex, chairID, roomID) {
		return false
	}
	var JiangDui bool = true
	//检查牌眼
	for i := 0; i < int(MAX_INDEX); i++ {
		if (GetCardIndex(cbCardIndex[i]) != 2) && (GetCardIndex(cbCardIndex[i]) != 5) && (GetCardIndex(cbCardIndex[i]) != 8) {
			JiangDui = false
		}

	}

	for i := 0; i < int(M_cbWeaveItemCount[roomID][chairID]); i++ {
		if M_WeaveItemArray[roomID][chairID][i].CbCenterCard != 2 || M_WeaveItemArray[roomID][chairID][i].CbCenterCard != 5 || M_WeaveItemArray[roomID][chairID][i].CbCenterCard != 8 {
			JiangDui = false
		}
	}

	return JiangDui
}

//清一色牌
func IsQingYiSe(roomid int, cbCardIndex [MAX_INDEX]int16, cbCurrentCard int16, chairID int16) bool {
	var cardcolor int16 = 0xff
	var QingYiSeStatus bool = true
	for i := 0; i < int(MAX_INDEX); i++ {
		if (cbCardIndex[i] >= 1) && (cardcolor == 0xff) {
			cardcolor = GetCardColor(int16(i))
		}
		if (cbCardIndex[i] > 0) && (cardcolor != GetCardColor(int16(i))) {
			QingYiSeStatus = false
		}
		if cardcolor == GetQueueColor(roomid, chairID) {
			QingYiSeStatus = false
		}
	}
	for i := 0; i < int(M_cbWeaveItemCount[roomid][chairID]); i++ {
		if cardcolor != GetCardColor(int16(M_WeaveItemArray[roomid][chairID][i].CbCenterCard)) {
			QingYiSeStatus = false
		}
	}
	if cardcolor != GetCardColor(cbCurrentCard) && (cbCurrentCard != 0) {
		QingYiSeStatus = false
	}

	return QingYiSeStatus
}

//带幺
//bool CGameLogic::IsDaiYao(const BYTE cbCardIndex[MAX_INDEX],const tagAnalyseItem *pAnalyseItem )
func IsDaiYao(roomid int, chairID int16) bool {

	var isindaiyao bool = true
	if Tableslink_daiyao[roomid][chairID] == 0 {
		isindaiyao = false
	}
	for i := 0; i < int(M_cbWeaveItemCount[roomid][chairID]); i++ {
		if (M_WeaveItemArray[roomid][chairID][i].CbCenterCard%10 != 1)&&(M_WeaveItemArray[roomid][chairID][i].CbCenterCard%10!= 9) {
			isindaiyao = false
		}
	}

	return isindaiyao
}

func NoDaiYao_WithCard(roomid int, chairID int16, cbcard int16) bool {
	var resultvalue bool
	var tmpcbcard [MAX_INDEX]int16
	resultvalue = true
	for i := 0; i < int(MAX_INDEX); i++ {
		tmpcbcard[i] = M_cbCardIndex[roomid][chairID][i]
	}
	tmpcbcard[cbcard]++
	for i := 0; i < int(MAX_INDEX); i++ {
		if tmpcbcard[i] > 0 {
			if (i%10 == 1) || (i%10 == 9) {
				resultvalue = false
			}
		}
	}
	for i := 0; i < int(M_cbWeaveItemCount[roomid][chairID]); i++ {
		if (M_WeaveItemArray[roomid][chairID][i].CbCenterCard%10 == 1) || (M_WeaveItemArray[roomid][chairID][i].CbCenterCard%10 == 9) {
			resultvalue = false
		}
	}
	return resultvalue
}

//用户查叫
func UserChaJiao(roomid int, ChairID int16) (int16, int16, string, string, int16) {
	//m_FanShu[]
	var m_FanShu [2]int16
	var m_fanshustr string
	var m_fanshustr2 string
	var cbcardtmp [MAX_INDEX]int16
	var maxFan [2]int16
	var maxFanstr string
	var maxFanstr2 string

	maxFanstr = ""
	maxFanstr2 = ""
	maxFan[0] = 0
	maxFan[1] = 0

	for j := 0; j < int(MAX_INDEX); j++ {
		//maxFan[0] = 0
		//maxFan[1] = 0

		for i := 0; i < int(MAX_INDEX); i++ {
			cbcardtmp[i] = M_cbCardIndex[roomid][ChairID][i]
		}
		//临时增加一张麻将
		if cbcardtmp[j] < 4 {
			if AnalyseChiHuCard_UserSendCard(roomid, cbcardtmp, ChairID, int16(j)) == WIK_CHI_HU {

				cbcardtmp[j]++

				if IsPengPeng(cbcardtmp, ChairID, roomid) {
					m_FanShu[0] = 1
					m_fanshustr = "大对子,"
				}
				if IsQiXiaoDui(roomid, cbcardtmp, int16(ChairID)) {
					m_FanShu[0] = 2
					m_fanshustr = m_fanshustr + "七对,"
				}
				if IsJiangDui(roomid, cbcardtmp, int16(ChairID)) && (M_GameDaiyaoandJiangDui[roomid]) {
					m_FanShu[0] =3
					m_fanshustr = m_fanshustr + "将对,"
				}
				var GenCount int16 = 0
				GenCount = 0
				m_FanShu[1]=0
				//	m_bPlayStatus[ChairID]=false
				//计算手中的根 或四归一
				for i := 0; i < int(MAX_INDEX); i++ {
					if cbcardtmp[i] == 4 {
						GenCount++
					}
					if cbcardtmp[i] == 1 {
						for tmpj := 0; tmpj < int(M_cbWeaveItemCount[roomid][ChairID]); tmpj++ {
							if M_WeaveItemArray[roomid][ChairID][tmpj].CbWeaveKind == WIK_PENG {
								fmt.Println("需要算根的已经碰牌", M_WeaveItemArray[roomid][ChairID][tmpj].CbCenterCard)
							}
							if M_WeaveItemArray[roomid][ChairID][tmpj].CbCenterCard == int16(i) && M_WeaveItemArray[roomid][ChairID][tmpj].CbWeaveKind == WIK_PENG {
								GenCount++
							}
						}

					}
				}
				for i := 0; i < int(M_cbWeaveItemCount[roomid][ChairID]); i++ {
					if M_WeaveItemArray[roomid][ChairID][i].CbWeaveKind == WIK_GuaFeng || M_WeaveItemArray[roomid][ChairID][i].CbWeaveKind == WIK_XiaYu {
						GenCount++
					}
				}
				//根的番数
				if GenCount > 0 {
					m_FanShu[1] = GenCount
					m_fanshustr2 = "(" + strconv.Itoa(int(GenCount)) + ")根,"
				} //根的番数
				if IsQingYiSe(roomid, cbcardtmp, 0, ChairID) {
					m_FanShu[1] += 2
					m_fanshustr2 = m_fanshustr2 + "清一色,"
				} //清一色
				//带幺
				if IsDaiYao(roomid, ChairID) && M_GameDaiyaoandJiangDui[roomid] {
					m_FanShu[1] += 2
					m_fanshustr2 = m_fanshustr2 + "带幺九,"
				} //带幺九
				//中张的意思即为断幺九
				if NoDaiYao_WithCard(roomid, ChairID, int16(j)) && M_GameMenQingZhongZhang[roomid] {
					m_FanShu[1] += 2
					m_fanshustr2 = m_fanshustr2 + "中张(断幺九),"
				} //断幺九
				//if NoDaiYao(roomid,ChairID)&&M_GameMenQingZhongZhang[roomid]{m_FanShu[1]+=2
				//	m_fanshustr2=m_fanshustr2+"断幺九,"
				//}//断幺九
				//如果有门清中张规则就判断门清
				if (M_cbWeaveItemCount[roomid][ChairID] == 0) && (M_GameMenQingZhongZhang[roomid]) {
					m_FanShu[1] += 1
					m_fanshustr2 = m_fanshustr2 + "门清,"
				} //门清

			}
		}
		if maxFan[0]+maxFan[1] < m_FanShu[0]+m_FanShu[1] {
			maxFan[0] = m_FanShu[0]
			maxFan[1] = m_FanShu[1]
			maxFanstr = m_fanshustr
			maxFanstr2 = m_fanshustr2
           fmt.Println("基本番",m_fanshustr,"额外番",m_fanshustr2)
			m_fanshustr = ""
			m_fanshustr2=""
		}

	}
	var FanTotal int16
	var JiFen int16
	FanTotal = maxFan[0] + maxFan[1]
	if FanTotal > M_GameMaxFan[roomid] {
		FanTotal = M_GameMaxFan[roomid]
	}
	switch FanTotal {

	case 0:
		{
			JiFen = 1
		}
	case 1:
		{
			JiFen = 2
		}
	case 2:
		{
			JiFen = 4
		}
	case 3:
		{
			JiFen = 8
		}
	case 4:
		{
			JiFen = 16
		}
	case 5:
		{
			JiFen = 32
		}
	case 6:
		{
			JiFen = 64
		}
	case 7:
		{
			JiFen = 128
		}
	case 8:
		{
			JiFen = 256
		}
	case 9:
		{
			JiFen = 512
		}
	case 10:
		{
			JiFen = 1024
		}
	case 11:
		{
			JiFen = 2048
		}
	case 12:
		{
			JiFen = 4096
		}
	case 13:
		{
			JiFen = 8192
		}
	case 14:
		{
			JiFen = 16384
		}
		//case 15:{JiFen=32768}

	}
	return maxFan[0], maxFan[1], maxFanstr, maxFanstr2, JiFen
}

//用户最终胡牌

func UserHuPai(roomid int, ChairID int16,  Cbcard int,Isself bool) (int16, int16, int16) {
	//m_FanShu[]
	var JiFen int16
	var m_FanShu [GAME_PLAYER][2]int16
	var cbcardtmp [MAX_INDEX]int16
	for i := 0; i < int(MAX_INDEX); i++ {
		cbcardtmp[i] = 0
		cbcardtmp[i] = M_cbCardIndex[roomid][ChairID][i]
	}
	if !Isself{cbcardtmp[Cbcard]++}

	JiFen = 0
	fmt.Println("胡牌时手中的麻将为:", cbcardtmp, "用户:", ChairID,"原来手中麻将:",M_cbCardIndex[roomid][ChairID])
	if IsPengPeng(cbcardtmp, ChairID, roomid) {
		m_FanShu[ChairID][0] = 1
		M_GameConcludeScore[roomid].UFanDescBase[ChairID] = M_GameConcludeScore[roomid].UFanDescBase[ChairID] + "大对子,"
	}
	if IsQiXiaoDui(roomid, cbcardtmp, int16(ChairID)) {
		m_FanShu[ChairID][0] = 2
		M_GameConcludeScore[roomid].UFanDescBase[ChairID] = M_GameConcludeScore[roomid].UFanDescBase[ChairID] + "七对,"
	}
	if IsJiangDui(roomid, cbcardtmp, int16(ChairID)) && M_GameDaiyaoandJiangDui[roomid] {
		m_FanShu[ChairID][0] = 3
		M_GameConcludeScore[roomid].UFanDescBase[ChairID] = M_GameConcludeScore[roomid].UFanDescBase[ChairID] + "将对,"
	}
	if m_FanShu[ChairID][0] == 0 {
		M_GameConcludeScore[roomid].UFanDescBase[ChairID] = M_GameConcludeScore[roomid].UFanDescBase[ChairID] + "平胡,"
	}

	//if m_FanShu[ChairID][0]==0{m_FanShu[ChairID][0]=1}
	var GenCount int16 = 0
	GenCount = 0
	//	m_bPlayStatus[ChairID]=false
	m_FanShu[ChairID][1] = 0
	//计算手中的根 或四归一
	fmt.Println("手中麻将牌为:", cbcardtmp, "座位号:", ChairID)
	for i := 0; i < int(MAX_INDEX); i++ {
		if cbcardtmp[i] == 4 {
			GenCount++
		}
		if cbcardtmp[i] == 1 {
			for tmpj := 0; tmpj < int(M_cbWeaveItemCount[roomid][ChairID]); tmpj++ {
				if M_WeaveItemArray[roomid][ChairID][tmpj].CbWeaveKind == WIK_PENG {
					fmt.Println("已经碰牌", M_WeaveItemArray[roomid][ChairID][tmpj].CbCenterCard)
				}
				if M_WeaveItemArray[roomid][ChairID][tmpj].CbCenterCard == int16(i) && M_WeaveItemArray[roomid][ChairID][tmpj].CbWeaveKind == WIK_PENG {
					GenCount++
				}
			}

		}
	}

	//计算已经刮风下雨中的根
	for i := 0; i < int(M_cbWeaveItemCount[roomid][ChairID]); i++ {
		if M_WeaveItemArray[roomid][ChairID][i].CbWeaveKind == WIK_GuaFeng || M_WeaveItemArray[roomid][ChairID][i].CbWeaveKind == WIK_XiaYu {
			GenCount++
		}
	}
	//根的番数
	if GenCount > 0 {
		m_FanShu[ChairID][1] += GenCount
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = "(" + strconv.Itoa(int(GenCount)) + ")根,"
	} //根的番数

	//清一色
	if IsQingYiSe(roomid, cbcardtmp, 0, ChairID) {
		m_FanShu[ChairID][1] += 2
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "清一色,"
	} //清一色
	//带幺九
	if IsDaiYao(roomid, ChairID) && M_GameDaiyaoandJiangDui[roomid] {
		m_FanShu[ChairID][1] += 2
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "带幺九,"
	} //带幺九

	if NoDaiYao_WithCard(roomid, ChairID, int16(Cbcard)) && M_GameMenQingZhongZhang[roomid] {
		m_FanShu[ChairID][1] += 2
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "中张(断幺九),"
	} //断幺九
	//门清
	if M_cbWeaveItemCount[roomid][ChairID] == 0 && M_GameMenQingZhongZhang[roomid] {
		m_FanShu[ChairID][1] += 1
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "门清,"
	} //门清
	if BeLastQiangGang[roomid][M_wProvideUser[roomid]] == 1 {
		m_FanShu[ChairID][1] += 1
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "抢杠胡,"
	}
	//自摸
	var OnlyZhiMo bool
	var CurUserCount int
	CurUserCount = 0
	OnlyZhiMo = false
	if M_wProvideUser[roomid] == ChairID {
		//自摸
		//天胡
		//SetFanXing(CHR_TIAN_HU,m_wBankerUser)

		OnlyZhiMo = true
		if (ChairID == M_wBankerUser[roomid]) && (M_cbSendCardCount[roomid][ChairID] == 0) && M_GameTianDiHu[roomid] {
			//SetFanXing(CHR_TIAN_HU,m_wBankerUser)
			OnlyZhiMo = false
			m_FanShu[ChairID][1] += 5
			M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "天胡,"
		}

		//地胡
		/*
			if (M_cbSendCardCount[roomid][ChairID]==0)&&(M_wProvideUser[roomid]==M_wBankerUser[roomid]&&M_GameTianDiHu[roomid]&&ChairID!=M_wBankerUser[roomid]){
			OnlyZhiMo=false
			m_FanShu[ChairID][1]+=5
			M_GameConcludeScore[roomid].UFanDescAddtion[ChairID]=M_GameConcludeScore[roomid].UFanDescAddtion[ChairID]+"地胡,"
			}
		*/
		fmt.Println("麻将个数", M_cbSendCardCount[roomid][ChairID], M_cbSendCardCount[roomid])
		if M_cbSendCardCount[roomid][ChairID] == 1 && M_wProvideUser[roomid] == int16(ChairID) && M_GameTianDiHu[roomid] && int16(ChairID) != M_wBankerUser[roomid] {
			var isDihu bool = true
			for i := 0; i < int(GAME_PLAYER); i++ {
				if M_cbWeaveItemCount[roomid][i] > 0 {
					isDihu = false
				}
			}
			if isDihu { //SetFanXing(CHR_DI_HU,int16(ChairID))
				OnlyZhiMo = false
				m_FanShu[ChairID][1] += 5
				M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "地胡,"
			}

		}
		if BeLastisGang[roomid][ChairID] == WIK_GuaFeng || BeLastisGang[roomid][ChairID] == WIK_XiaYu {
			//杠上花;
			m_FanShu[ChairID][1] += 1
			M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "杠上花,"
			fmt.Println("当前结果操作结果杠上花", M_GameConcludeScore[roomid].UFanDescAddtion[ChairID], "最后杠牌", BeLastisGang[roomid], "供牌用户", M_wProvideUser[roomid], "当前胡牌用户", ChairID)
			BeLastisGang[roomid][ChairID]=WIK_NULL //清除防止下家直接胡牌就显示杠上炮。
			OnlyZhiMo = false
		}
		for tmpij := 0; tmpij < int(M_DesktopPlayer[roomid]); tmpij++ {
			if M_GameStatus[roomid][tmpij] != GAME_STATUS_WINED {
				CurUserCount++
			}
		}

		if OnlyZhiMo {
			fmt.Println("自摸类型：", OnlyZhiMo, "type", M_GameZiMoAddtype[roomid], "原来番数", m_FanShu[ChairID][1])
			//自摸家数去掉自己
			CurUserCount--
			if M_GameZiMoAddtype[roomid] {
				JiFen++
			} else {
				m_FanShu[ChairID][1] = m_FanShu[ChairID][1] + 1
			}
			M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "自摸(" + strconv.Itoa(int(CurUserCount)) + ")家"

		}

	}

	//fmt.Println("杠上花，杠上炮判断",BeLastisGang[roomid],"供牌用户",M_wProvideUser[roomid],"胡牌用户",ChairID)
	//如果供牌用户最后一次操作是杠
	if BeLastisGang[roomid][M_wProvideUser[roomid]] == WIK_GuaFeng || BeLastisGang[roomid][M_wProvideUser[roomid]] == WIK_XiaYu && M_wProvideUser[roomid] != ChairID {

		//杠上炮
		m_FanShu[ChairID][1] += 1
		//杠上炮转雨
		var gangBeishu int16
		gangBeishu = 1
		var GangsanPaoindex int
		GangsanPaoindex=int(M_cbWeaveItemCount[roomid][M_wProvideUser[roomid]]-1)
		if BeLastisGang[roomid][M_wProvideUser[roomid]] == WIK_GuaFeng{
			for tmpsj:=0;tmpsj<int(M_cbWeaveItemCount[roomid][M_wProvideUser[roomid]]);tmpsj++{
			 if M_WeaveItemArray[roomid][M_wProvideUser[roomid]][tmpsj].CbCenterCard==M_cbProvideCard[roomid]{
				 GangsanPaoindex=tmpsj
			 }
			}

		}

		fmt.Println("供牌用户", M_wProvideUser[roomid], "最后杠操作的INDEX", M_cbWeaveItemCount[roomid][M_wProvideUser[roomid]]-1, "杠牌种类", M_WeaveItemArray[roomid][M_wProvideUser[roomid]][M_cbWeaveItemCount[roomid][M_wProvideUser[roomid]]-1].CbWeaveKind)
		if M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].CbWeaveKind == WIK_XiaYu {
			gangBeishu = 2
		}
		var ZhuangYU int
		var ZhuangYuUsers int
		ZhuangYU = 0
		ZhuangYuUsers = 0
		fmt.Println("转雨前杠分情况:", M_GangScore[roomid][0], M_GangScore[roomid][1], M_GangScore[roomid][2], M_GangScore[roomid][3], "转几家雨:", M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].BeGangUserCount, "倍数", gangBeishu)
		if M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].WProvideUser == M_wProvideUser[roomid] { //自己摸的杠牌


			ZhuangYuUsers = int(M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].BeGangUserCount)

			M_GangScore[roomid][ChairID][M_wProvideUser[roomid]] = M_GangScore[roomid][ChairID][M_wProvideUser[roomid]] + int16(ZhuangYuUsers)*gangBeishu*M_initGameDrawScore[roomid]
			M_GangScore[roomid][M_wProvideUser[roomid]][ChairID] = M_GangScore[roomid][M_wProvideUser[roomid]][ChairID] - int16(ZhuangYuUsers)*gangBeishu*M_initGameDrawScore[roomid]
			fmt.Println("转的钱", int16(ZhuangYuUsers)*gangBeishu*M_initGameDrawScore[roomid], "得分用户积分", M_GangScore[roomid][ChairID][M_wProvideUser[roomid]])
			ZhuangYU = int(M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].BeGangUserCount * gangBeishu * M_initGameDrawScore[roomid])

		} else { //其他人点杠

			M_GangScore[roomid][M_wProvideUser[roomid]][M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].WProvideUser] -= gangBeishu * M_initGameDrawScore[roomid]
			M_GangScore[roomid][ChairID][M_WeaveItemArray[roomid][M_wProvideUser[roomid]][GangsanPaoindex].WProvideUser] += gangBeishu * M_initGameDrawScore[roomid]
			M_WeaveItemArray[roomid][M_wProvideUser[roomid]][M_cbWeaveItemCount[roomid][M_wProvideUser[roomid]]-1].GangSangPao = true
			ZhuangYU = int(gangBeishu * M_initGameDrawScore[roomid])
			fmt.Println("杠上炮为他人点杠   点杠人")
			ZhuangYuUsers++
		}
		//杠上炮转雨
		M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] = M_GameConcludeScore[roomid].UFanDescAddtion[ChairID] + "杠上炮,"
		fmt.Println("杠操作：当前结果操作结果杠上炮", M_GameConcludeScore[roomid].UFanDescAddtion[ChairID], "最后杠牌", BeLastisGang[roomid], "供牌用户", M_wProvideUser[roomid], "当前胡牌用户", ChairID, "转雨钱:", ZhuangYU, "转雨：", ZhuangYuUsers, "家   杠牌分情况", M_GangScore[roomid][0], M_GangScore[roomid][1], M_GangScore[roomid][2], M_GangScore[roomid][3])
	}


	var totalfan int16
	totalfan = m_FanShu[ChairID][0] + m_FanShu[ChairID][1]
	fmt.Println("总共番型为：", totalfan, "最大番数", M_GameMaxFan[roomid])
	if totalfan > M_GameMaxFan[roomid] {
		totalfan = M_GameMaxFan[roomid]
		if (OnlyZhiMo) && (!M_GameZiMoAddtype[roomid]) {
			totalfan = M_GameMaxFan[roomid] + 1
		}
	}
	switch totalfan {
	case 0:
		{
			JiFen += 1
		}
	case 1:
		{
			JiFen += 2
		}
	case 2:
		{
			JiFen += 4
		}
	case 3:
		{
			JiFen += 8
		}
	case 4:
		{
			JiFen += 16
		}
	case 5:
		{
			JiFen += 32
		}
	case 6:
		{
			JiFen += 64
		}
	case 7:
		{
			JiFen += 128
		}
	case 8:
		{
			JiFen += 256
		}
	case 9:
		{
			JiFen += 512
		}
	case 10:
		{
			JiFen += 1024
		}
	case 11:
		{
			JiFen += 2048
		}
	case 12:
		{
			JiFen += 4096
		}
	case 13:
		{
			JiFen += 8192
		}
	case 14:
		{
			JiFen += 16384
		}
		//case 15:{JiFen=32768}
	}
	return m_FanShu[ChairID][0], m_FanShu[ChairID][1], JiFen
}
func GameRecordDrawScore(roomid int) {
	var cardid int16
	//取得每个用户手中的牌

	for tmpi := 0; tmpi < int(M_DesktopPlayer[roomid]); tmpi++ {
		cardid = 0
		fmt.Println("游戏结束:玩家", tmpi, "手中的麻将数据为:", M_cbCardIndex[roomid][tmpi])
		for tmpj := 0; tmpj < int(MAX_INDEX); tmpj++ {
			if M_cbCardIndex[roomid][tmpi][tmpj] > 0 {
				for tmph := 0; tmph < int(M_cbCardIndex[roomid][tmpi][tmpj]); tmph++ {
					M_GameConcludeScore[roomid].CbHandCardData[tmpi][cardid] = int16(tmpj)
					cardid++
				}

			}
		}
	}
	fmt.Println("游戏玩家：0手牌",M_GameConcludeScore[roomid].CbHandCardData[0],"游戏玩家：1手牌",M_GameConcludeScore[roomid].CbHandCardData[1])

	for tmpi := 0; tmpi < int(M_DesktopPlayer[roomid]); tmpi++ {
		//M_GameConcludeScore[roomid].LGameScore[tmpi]=int(M_GameScore[roomid][tmpi])
		M_GameConcludeScore[roomid].LGangScore[tmpi] = 0
		for tmpj := 0; tmpj < int(M_DesktopPlayer[roomid]); tmpj++ {
			M_GameConcludeScore[roomid].LGangScore[tmpi] = M_GameConcludeScore[roomid].LGangScore[tmpi] + M_GangScore[roomid][tmpi][tmpj]
		}

	}
	M_GameRecordDraw[roomid]++
}
