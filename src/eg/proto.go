package eg

/*
//服务器命令结构

#define SUB_S_GAME_START					100									//游戏开始
#define SUB_S_GAME_CALLBANKER				101									//抢分叫庄
#define SUB_S_GAME_SENDCARD					102									//发送玩家牌
#define SUB_S_GAME_MAKEZHUSTART				103									//叫主开始
#define SUB_S_GAME_CHANGEBOTTOMCARDSTART	103									//玩家交换底牌开始
#define SUB_S_GAME_SHOWFRIENDSTART			104									//显示朋友牌
#define SUB_S_GAME_OUTCARDSTART				105									//出牌开始
#define SUB_S_GAME_OUTCARD					105									//出牌
#define SUB_S_PERROUND_FIRSTUSER			106									//每一轮第一个出牌的玩家
#define SUB_S_RECALL						107									//没人叫，重新叫

struct CMD_S_StatusFree
{
	int			time;								//剩余时间
	BYTE		status;								//当前游戏阶段
};

struct CMD_S_StatusPlay
{
	int			time;								//剩余时间
	BYTE		status;								//当前游戏阶段
	BYTE		banker;								//庄家
	BYTE		bankersocre;						//庄家叫分
	BYTE		color;								//主牌颜色
	BYTE		bottomcard[MAX_BOTTOME_CARD_COUNT]; //一开始分配的底牌
	BYTE		chgbottomcard[MAX_BOTTOME_CARD_COUNT]; //交换后的底牌
	BYTE		firendcard;							//朋友牌
	BYTE		card[MAX_COUNT];					//玩家当前的牌
	BYTE		outcard[GAME_PLAYER];				//这轮出牌
	BYTE		score[GAME_PLAYER];					//这轮得分情况
	BYTE		totalscore[GAME_PLAYER];			//当前总得分
	BYTE		curuser;							//当前出牌用户
	BYTE		firstuser;							//这轮第一个出牌的用户
	BYTE		outlist;							//当前轮出牌
};

//游戏开始,进入叫庄阶段
struct CMD_S_GameStart
{
	BYTE		startbanker;				//一开始随机定的庄，默认显示80
	BYTE		callbanker;					//当前叫庄的人
	BYTE		card[MAX_COUNT];			//玩家手牌
};

//玩家叫庄
struct CMD_S_CallBanker {
	BYTE		callchair;
	BYTE		nextcaller;
	BYTE		bankerscore;
};

//叫主开始
struct CMD_S_MakeZhuStart {
	BYTE		banker;
	BYTE		bankersore;
	BYTE		call;
};

//交换底牌开始
struct CMD_S_Banker_ChangeCardStart {
	BYTE		color;
	BYTE		bottomcard[MAX_BOTTOME_CARD_COUNT];
};

//交朋友开始
struct CMD_S_ShowFriendStart {
	BYTE		bottomcard[MAX_BOTTOME_CARD_COUNT];
};

//交朋友结束，进入出牌阶段
struct CMD_S_OutCardStart {
	BYTE		friencard;
	BYTE		outuser;
	BYTE		firstuser;
};

//玩家出牌
struct CMD_S_UserOutCard {
	BYTE		card;
	BYTE		firstuser;
	BYTE		outuser;
	BYTE		nextuser;
	BYTE		newround;
	BYTE		score;
	BYTE		totalscore;
};

struct CMD_S_GameEnd
{
	BYTE		bankerchair;										//庄家座位号
	BYTE		friendchair;										//朋友座位号
	BYTE		scoreinfo[GAME_PLAYER];								//每个人得分情况
};

//////////////////////////////////////////////////////////////////////////
//客户端命令结构

#define CMD_C_CALLBANKER				1							//玩家抢庄
#define CMD_C_MAKEZHU					2							//玩家选主
#define CMD_C_CHANGEBOTTOMCARD			3							//玩家交换底牌
#define CMD_C_SHOWFRIEND				4							//玩家显示朋友牌
#define CMD_C_OUTCARD					5							//玩家出牌

struct CMD_C_CallBanker {
	BYTE		callscore;
};

struct CMD_C_MakeZhu {
	BYTE		color;
};

struct CMD_C_ChangeBottomCard {
	BYTE		bottomcard[MAX_BOTTOME_CARD_COUNT];
};

struct CMD_C_ShowFriend {
	BYTE		friendcard;
};

struct CMD_C_OutCard {
	BYTE		card;
};

 */
