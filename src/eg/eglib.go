package eg

import (
	"exportor/defines"
	"math/rand"
)

const (
	GAME_PLAYER				= 4
)

const (
	MAX_COUNT				=		12									//扑克数目
	MAX_CARD_COUNT			=		54									//扑克总数目
	MAX_BOTTOME_CARD_COUNT	=		6									//底牌数目
	MAX_INDEX				= 		0x45
)

const (
	BANKER_SCORE_DEFAULT	=			0							//服务器默认值，客户端从下面值开始
	BANKER_SCORE_NO			=			1
	BANKER_SCORE_80			=			2
	BANKER_SCORE_85			=			3
	BANKER_SCORE_90			=			4
	BANKER_SCORE_95			=			5
	BANKER_SCORE_100		=			6
	BANKER_SCORE_100_BAO	=			7
	BANKER_SCORE_100_GOU	=			8
)


const (
	GAME_STATUS_FREE				=	201							//游戏闲阶段，有空缺位置
	GAME_STATUS_CALLBANKER			=	202							//游戏叫庄阶段
	GAME_STATUS_MAKEZHU				=	203							//选主阶段
	GAME_STATUS_CHANGEBOTTOMCARD	=	204							//交换底牌阶段
	GAME_STATUS_CALLFRIEND			=	205							//选朋友阶段
	GAME_STATUS_OUTCARD				=	206							//出牌阶段
	GAME_STATUS_CONCLUDE			=	207							//阶段阶段
)

type user struct {
	player 		*defines.PlayerInfo
	seat 		int

	cstatus 	bool
	recall		[]int

	handcard 	[]int
	cardindex	[]int
}

type eglib struct {
	mgr 			defines.IGameManager

	userList 		[]*user
	leftround		int

	bottomcard		[]int
	status 			int
	tm 				int

	banker 			int
	curbanker		int
	bankerscore 	uint64

	callsocre		uint64
	recallcount		int

	color 			int

	fcard 			int
	fchair			int

	firstuser 		int
	curuser			int


}

func newlib() *eglib {
	eg := &eglib{}
	return eg
}

func (lib *eglib) OnInit(manager defines.IGameManager, gamedata interface{}) error {
	return nil
}

func (lib *eglib) OnRelease() {

}

func (lib *eglib) OnGameCreate(info *defines.PlayerInfo, conf *defines.CreateRoomConf) error {
	return nil
}

func (lib *eglib) OnUserEnter(info *defines.PlayerInfo) error {
	return nil
}

func (lib *eglib) OnUserLeave(info *defines.PlayerInfo) {

}

func (lib *eglib) OnUserOffline(info *defines.PlayerInfo) {

}

func (lib *eglib) OnUserMessage(info *defines.PlayerInfo, cmd uint32, data []byte) error {
	return nil
}

func (lib *eglib) OnTimer(id uint32, data interface{}) {

}

func (lib *eglib) userReady() {

}

func (lib *eglib) startGame() {
	card := make([]int, MAX_COUNT)
	randCard(card)

	for i := 0; i < GAME_PLAYER; i++ {
		u := lib.userList[i]
		first := i * MAX_COUNT
		last := (i + 1) * MAX_COUNT
		u.handcard = append(u.handcard, card[first:last]...)
		lib.makeIndex(u)
	}

	lib.bottomcard = append(lib.bottomcard, card[GAME_PLAYER*MAX_COUNT:]...)
	lib.banker = rand.Int() % GAME_PLAYER
	lib.curbanker = (lib.banker + 1) % GAME_PLAYER

	//lib.mgr.BroadcastMessage()

	lib.status = GAME_STATUS_CALLBANKER
	//lib.mgr.SetTimer()
}

func (lib *eglib) makeIndex(u *user) {
	for i := 0; i < MAX_COUNT; i++ {
		u.cardindex[u.handcard[i]]++
	}
}

func (lib *eglib) userCallBanker(p *defines.PlayerInfo, score int) {

}

/*

const BYTE zhucard[] =
{
	0, 	0, 	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0x0D,0x0E,0x0F,		//方块 3 - 2
	0, 	0, 	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0x0D,0x1E,0x1F,		//梅花 3 - 2
	0, 	0, 	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0x0D,0x2E,0x2F,		//红桃 3 - 2
	0, 	0, 	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0,	0x0D,0x3E,0x3F,		//黑桃 3 - 2
	0, 	0, 	0,	0x43,	0x44												//小鬼，大鬼
};

const BYTE scorecard[] = { 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 10, 0, 0, 10, 0, 0 };

//////////////////////////////////////////////////////////////////////////

#define IDI_TIMER						1
//////////////////////////////////////////////////////////////////////////

//构造函数
CTableFrameSink::CTableFrameSink()
{
	srand(time(NULL));
	return;
}

CTableFrameSink::~CTableFrameSink(void)
{

}

VOID CTableFrameSink::Release()
{
	delete this;
}

void *  CTableFrameSink::QueryInterface(const IID & Guid, DWORD dwQueryVer)
{
	QUERYINTERFACE(ITableFrameSink,Guid,dwQueryVer);
	QUERYINTERFACE(ITableUserAction,Guid,dwQueryVer);
	QUERYINTERFACE_IUNKNOWNEX(ITableFrameSink,Guid,dwQueryVer);
	return NULL;
}

bool  CTableFrameSink::Initialization(IUnknownEx * pIUnknownEx)
{
	ASSERT(pIUnknownEx!=NULL);
	_tableframe = QUERY_OBJECT_PTR_INTERFACE(pIUnknownEx,ITableFrame);
	if (_tableframe==NULL) return false;

	_tableframe->SetStartMode(START_MODE_ALL_READY);

	_gameoptions = _tableframe->GetGameServiceOption();
	ASSERT(_gameoptions);

	ReadConfigInformation();

	bool bstatus = IsRoomCardScoreType();

	_serino = 0;
	return true;
}

void  CTableFrameSink::RepositionSink()
{
	return;
}

void CTableFrameSink::rest() {

}

bool CTableFrameSink::OnUserScroeNotify(WORD wChairID, IServerUserItem * pIServerUserItem, BYTE cbReason)
{
	return false;
}

bool CTableFrameSink::OnActionUserOffLine(WORD wChairID, IServerUserItem * pIServerUserItem)
{

	return true;
}

bool  CTableFrameSink::OnEventGameStart() {

	_tableframe->SetGameStatus(GAME_STATUS_PLAY);

	if ((_gameoptions->wServerType&GAME_GENRE_PERSONAL) !=0 ) {
		_rule = _tableframe->GetGameRule();
	}

	for (WORD i=0;i<GAME_PLAYER;i++) {

		IServerUserItem *pIServerUserItem=_tableframe->GetTableUserItem(i);
		if (pIServerUserItem==NULL) continue;

		const LONGLONG lUserScore=pIServerUserItem->GetUserScore();
		_userlist[i]._status = USER_GAME_STATUS_PLAY;
	}

	BYTE cardbox[MAX_CARD_COUNT];
	_gamelogic.randcardlist(cardbox, sizeof(cardbox)/sizeof(cardbox[0]));

	for (int i = 0; i < GAME_PLAYER; i++) {
		memcpy(_cardlist[i], cardbox + i * MAX_COUNT, MAX_COUNT);
		MakeCardIndex(i);
	}
	memcpy(_botomcard, cardbox + GAME_PLAYER * MAX_COUNT, MAX_BOTTOME_CARD_COUNT);

	_banker = rand() % GAME_PLAYER;
	_curbanker = (_banker + 1) % GAME_PLAYER;

	CMD_S_GameStart gamestart;
	gamestart.callbanker = _curbanker;
	for (WORD i = 0; i < GAME_PLAYER; i++) {
		memcpy(gamestart.card, _cardlist[i], MAX_COUNT);
		SendGameData(i, SUB_S_GAME_START, &gamestart, sizeof(gamestart));
	}

	_serino++;
	_gamestatus = GAME_STATUS_CALLBANKER;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_CALLBANKERTIMER, 1, _serino);

	_bankerscore = BANKER_SCORE_80;

	return true;
}

void CTableFrameSink::MakeCardIndex(WORD chair) {
	BYTE *cards = _cardlist[chair];
	BYTE *cardindex = _cardindex[chair];
	for (int i = 0; i < MAX_COUNT; i++) {
		cardindex[cards[i]]++;
	}
}

bool CTableFrameSink::CheckOutCard(WORD chair, BYTE card) {

	BYTE firstcard = _outcardlist[_firstuser];
	BYTE firstcolor = firstcard & 0xF0;

	bool zcard = zhucard[firstcard] != 0;
	bool zcolor = firstcolor == _color;
	BYTE ccolor = card & 0xF0;

	if (zcard || zcolor) {

		if (zhucard[card])
			return true;

		if (ccolor == firstcolor)
			return true;

		BYTE zmin = _color << 4 + 0x03;
		BYTE zmax = _color << 4 + 0x0F;
		for (BYTE i = zmin; i <= zmax; i++) {
			if (_cardindex[chair][i])
				return false;
		}
	}
	else {

		if (ccolor == firstcolor)
			return true;

		BYTE fmin = firstcolor << 4 + 0x03;
		BYTE fmax = firstcolor << 4 + 0x0F;
		for (BYTE i = fmin; i <= fmax; i++) {
			if (_cardindex[chair][i])
				return false;
		}
	}

	return true;
}

void CTableFrameSink::CalScore() {

	BYTE totalscore = 0;
	for (int i = 0; i < GAME_PLAYER; i++) {
		BYTE outcolor = _outcardlist[i] & 0xF0;
		if (outcolor == 4) continue;
		if (scorecard[_outcardlist[i]] == 0) continue;
		totalscore += scorecard[_outcardlist[i]];
	}
	if (totalscore == 0)
		return;

	BYTE maxscorechair = GetMaxScoreChair();
	if (maxscorechair == INVALID_BYTE)
		return;

	_winscore[maxscorechair] += totalscore;
}

BYTE CTableFrameSink::GetNextChair(WORD chair) {
	return (chair + 1) % GAME_PLAYER;
}

BYTE CTableFrameSink::GetNextFirstChair() {
	if (_leftround == 1)
		return INVALID_BYTE;
	return GetMaxScoreChair();
}

BYTE CTableFrameSink::GetMaxScoreChair() {
	WORD sortlist[GAME_PLAYER];
	for (int i = 0; i < GAME_PLAYER; i++) {
		BYTE color = _outcardlist[i] & 0xF0;
		WORD value = _outcardlist[i];
		if (color == _color)
			value += 1000;
		sortlist[i] = value;
	}

	WORD maxvalue = sortlist[0];
	for (int i = 1; i < GAME_PLAYER; i++) {
		if (sortlist[i] > maxvalue)
			maxvalue = sortlist[i];
	}

	BYTE firstuser = INVALID_BYTE;
	for (int i = 0; i < GAME_PLAYER; i++) {
		BYTE color = _outcardlist[i] & 0xF0;
		WORD value = _outcardlist[i];
		if (color == _color)
			value += 1000;
		if (maxvalue == value)
			return firstuser;
	}

	return firstuser;
}

BYTE CTableFrameSink::GetRecomemdZhu() {
	BYTE colorcnt[GAME_PLAYER];
	memset(colorcnt, 0, GAME_PLAYER);
	BYTE *cardindex = _cardindex[_banker];
	return rand() % 4;
}

void CTableFrameSink::GetRecomendBottomCard(BYTE *bottomcard) {
	memcpy(bottomcard, _cardlist[_banker], MAX_BOTTOME_CARD_COUNT);
}

BYTE CTableFrameSink::GetRecomemdOutCard(WORD chair) {
	BYTE start = rand() % 100;
	BYTE end = start + MAX_INDEX;
	for (int i = start; i < end; i++) {
		BYTE index = i % MAX_INDEX;
		if (_cardindex[chair][index])
			return index;
	}
}

bool CTableFrameSink::OnEventGameConclude(WORD wChairID, IServerUserItem * pIServerUserItem, BYTE cbReason)
{
	switch (cbReason)
	{
	case GER_NORMAL:
		{
			_gamestatus = GAME_STATUS_CONCLUDE;

			CMD_S_GameEnd gameend;
			ZeroMemory(&gameend, sizeof(gameend));

			gameend.bankerchair = _banker;
			gameend.friendchair = _friendchair;

			BYTE multiple = 1;
			if (_bankerscore == BANKER_SCORE_100_BAO) {
				multiple = 4;
			}
			else if (_bankerscore == BANKER_SCORE_100_GOU) {
				multiple = 2;
			}
			for (int i = 0; i < GAME_PLAYER; i++)
				gameend.scoreinfo[i] = _winscore[i] * multiple;


			_tableframe->ConcludeGame(GAME_STATUS_FREE);

			return true;
		}
	case GER_USER_LEAVE:
	case GER_NETWORK_ERROR:
		{

		}
	case GER_DISMISS:
		{
			return true;
		}
	}

	return false;
}

bool  CTableFrameSink::OnEventSendGameScene(WORD wChairID, IServerUserItem * pIServerUserItem, BYTE cbGameStatus, bool bSendSecret)
{
	switch (cbGameStatus)
	{
	case GAME_STATUS_FREE:
		{
			CMD_S_StatusFree StatusFree;
			ZeroMemory(&StatusFree,sizeof(StatusFree));

			return _tableframe->SendGameScene(pIServerUserItem,&StatusFree,sizeof(StatusFree));
		}
	case GAME_STATUS_PLAY:
		{
			CMD_S_StatusPlay StatusPlay;
			memset(&StatusPlay,0,sizeof(StatusPlay));

			return _tableframe->SendGameScene(pIServerUserItem,&StatusPlay,sizeof(StatusPlay));
		}
	}

	ASSERT(FALSE);

	return false;
}

bool CTableFrameSink::OnUserCallBanker(WORD chair, IServerUserItem * pIServerUserItem, BYTE callscore) {

	if (_gamestatus != GAME_STATUS_CALLBANKER)
		return true;

	if (chair != _curbanker)
		return true;

	//not allowed to call same score
	for (int i = 0; i < GAME_PLAYER; i++) {
		if (_call_status[i] != BANKER_SCORE_DEFAULT && _call_status[i] >= callscore)
			return true;
	}

	_call_status[chair] = callscore;
	if (_call_status[chair] != BANKER_SCORE_NO) {
		_recall_list[chair] = TRUE;
		_recallcount++;
	}
	else if (_recall_list[chair]){
		_recallcount--;
	}

	if (callscore != BANKER_SCORE_100_GOU) {
		BYTE nextcaller = (_curbanker + 1) % GAME_PLAYER;
		if (_call_status[nextcaller] == BANKER_SCORE_DEFAULT) {
			_curbanker = nextcaller;

			CMD_S_CallBanker callbanker_s;
			callbanker_s.bankerscore = callscore;
			callbanker_s.callchair = chair;
			callbanker_s.nextcaller = _curbanker;
			SendGameData(INVALID_CHAIR, SUB_S_GAME_CALLBANKER, &callbanker_s, sizeof(callbanker_s));

			_time = time(NULL);
			_tableframe->SetGameTimer(IDI_TIMER, IDI_CALLBANKERTIMER, 1, _serino);

			return true;
		}
		else {

			if (_recallcount == 1) {
				for (int i = 0; i < GAME_PLAYER; i++) {
					if (_recall_list[i]) {
						_banker = i;
						break;
					}
				}

				_bankerscore = _call_status[_banker];
			}
			else if (_recallcount > 0) {
				for (int c = 0, nbanker = _curbanker + 1; c < GAME_PLAYER; c++, nbanker++) {
					nbanker %= GAME_PLAYER;
					if (_recall_list[nbanker] == FALSE)
						continue;

					_curbanker = nextcaller;

					CMD_S_CallBanker callbanker_s;
					callbanker_s.bankerscore = callscore;
					callbanker_s.callchair = chair;
					callbanker_s.nextcaller = _curbanker;
					SendGameData(INVALID_CHAIR, SUB_S_GAME_CALLBANKER, &callbanker_s, sizeof(callbanker_s));

					_time = time(NULL);
					_tableframe->SetGameTimer(IDI_TIMER, IDI_CALLBANKERTIMER, 1, _serino);
					return true;
				}
			}
		}
	}
	else {
		_banker = _curbanker;
		_bankerscore = callscore;
	}

	CMD_S_MakeZhuStart makezhu;
	makezhu.banker = _banker;
	makezhu.bankersore = _bankerscore;
	for (int i = 0; i < GAME_PLAYER; i++) {
		if (i == _banker) {
			makezhu.call = TRUE;
		}
		else {
			makezhu.call = FALSE;
		}
		SendGameData(i, SUB_S_GAME_CHANGEBOTTOMCARDSTART, &makezhu, sizeof(makezhu));
	}

	_gamestatus = GAME_STATUS_MAKEZHU;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_MAKEZHU, 1, _serino);

	return true;
}

bool CTableFrameSink::OnUserMakeZhu(WORD chair, IServerUserItem * pIServerUserItem, BYTE color) {

	if (chair != _banker)
		return true;

	if (color < COLOR_MIN || color > COLOR_MAX)
		return true;

	_color = color;

	CMD_S_Banker_ChangeCardStart chgcard;
	chgcard.color = _color;
	for (int i = 0; i < GAME_PLAYER; i++) {
		if (i == _banker) {
			memcpy(chgcard.bottomcard, _botomcard, MAX_BOTTOME_CARD_COUNT);
		}
		else {
			memset(chgcard.bottomcard, 0, sizeof(chgcard.bottomcard));
		}
		SendGameData(i, SUB_S_GAME_CHANGEBOTTOMCARDSTART, &chgcard, sizeof(chgcard));
	}

	for (int i = 0; i < MAX_BOTTOME_CARD_COUNT; i++) {
		_cardindex[_banker][_botomcard[i]]++;
	}

	_gamestatus = GAME_STATUS_CHANGEBOTTOMCARD;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_CHAGNEBOTTOMCARD, 1, _serino);

	return true;
}

bool CTableFrameSink::OnUserChangeBottomCard(WORD chair, IServerUserItem * pIServerUserItem, BYTE *card) {
	for (int i = 0; i < MAX_BOTTOME_CARD_COUNT; i++) {
		if (_cardindex[chair][card[i]] == 0)
			return true;
	}

	for (int i = 0; i < MAX_BOTTOME_CARD_COUNT; i++) {
		BYTE value = card[i] & 0x0F;
		if (value == 0x05 || value == 0x0A || value == 0x0D)
			return true;
	}

	for (int i = 0; i < MAX_BOTTOME_CARD_COUNT; i++) {
		_cardindex[chair][_botomcard[i]] = 0;
	}

	CMD_S_ShowFriendStart showfriend;
	memcpy(showfriend.bottomcard, card, MAX_BOTTOME_CARD_COUNT);
	SendGameData(INVALID_CHAIR, SUB_S_GAME_SHOWFRIENDSTART, &showfriend, sizeof(showfriend));

	_gamestatus = GAME_STATUS_CALLFRIEND;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_SHOWFRIEND, 1, _serino);

	return true;
}

bool CTableFrameSink::OnUserShowFriend(WORD chair, IServerUserItem * pIServerUserItem, BYTE card) {

	if (chair != _banker)
		return true;

	if (card > CountArray(zhucard))
		return true;

	if (zhucard[card] == 0)
		return true;

	_friendcard = card;

	for (int i = 0; i < GAME_PLAYER; i++) {
		if (_cardindex[i][card]) {
			_friendchair = i;
			break;
		}
	}

	_curuser = _banker;
	_firstuser = _curuser;
	_leftround = MAX_COUNT;

	CMD_S_OutCardStart outstart;
	outstart.friencard = card;
	outstart.outuser = _banker;
	outstart.firstuser = _banker;
	SendGameData(INVALID_CHAIR, SUB_S_GAME_OUTCARDSTART, &outstart, sizeof(outstart));

	_gamestatus = GAME_STATUS_OUTCARD;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_OUTCARD, 1, _serino);

	return true;
}

bool CTableFrameSink::OnUserOutcard(WORD chair, IServerUserItem * pIServerUserItem, BYTE card) {

	if (chair != _curuser)
		return true;

	if (card < VALUE_MIN || card > VALUE_MAX)
		return true;

	if (_cardindex[chair][card] == 0)
		return true;

	if (CheckOutCard(chair, card) == false)
		return true;

	_cardindex[chair][card] = 0;
	_outcardlist[chair] = card;

	BYTE nextuser = GetNextChair(chair);
	if (nextuser == INVALID_BYTE) {
		CTraceService::TraceString(_TEXT("invalid next user"), TraceLevel_Exception);
		ASSERT(0);
		return true;
	}

	_curuser = nextuser;

	BYTE newround = false;
	if (_curuser == _firstuser) {
		BYTE nextfirstchair = GetNextFirstChair();
		if (nextfirstchair == INVALID_BYTE) {
			return OnEventGameConclude(INVALID_BYTE, 0, GER_NORMAL);
		}

		CalScore();

		_curuser = nextfirstchair;
		_firstuser = _curuser;
		newround = true;
		_leftround--;
		memset(_outcardlist, 0, sizeof(_outcardlist));
	}

	CMD_S_UserOutCard outcard;
	outcard.card = card;
	outcard.firstuser = _firstuser;
	outcard.outuser = chair;
	outcard.nextuser = _curuser;
	outcard.newround = newround;
	SendGameData(INVALID_BYTE, SUB_S_GAME_OUTCARD, &outcard, sizeof(outcard));

	_gamestatus = GAME_STATUS_OUTCARD;
	_time = time(NULL);
	_tableframe->SetGameTimer(IDI_TIMER, IDI_OUTCARD, 1, _serino);

	return true;
}

bool CTableFrameSink::OnTimerMessage(DWORD wTimerID, WPARAM wBindParam)
{
	switch (wTimerID)
	{
		case IDI_TIMER:
		{
			BYTE serino = (BYTE)wBindParam;

			int curtime = time(NULL);
			if (curtime - _time < 1) {
				CTraceService::TraceString(_TEXT("timer pre overcome "), TraceLevel_Warning);
			}

			if (serino != _serino) {
				CTraceService::TraceString(_TEXT("serino delayed"), TraceLevel_Warning);
				return true;
			}

			if (_gamestatus == GAME_STATUS_CALLBANKER) {
				return OnUserCallBanker(_curbanker, NULL, BANKER_SCORE_NO);
			}
			else if (_gamestatus == GAME_STATUS_MAKEZHU) {
				return OnUserMakeZhu(_banker, NULL, GetRecomemdZhu());
			}
			else if (_gamestatus == GAME_STATUS_CHANGEBOTTOMCARD) {
				BYTE bottomcard[MAX_BOTTOME_CARD_COUNT];
				GetRecomendBottomCard(bottomcard);
				return OnUserChangeBottomCard(_banker, NULL, bottomcard);
			}
			else if (_gamestatus == GAME_STATUS_OUTCARD) {
				return OnUserOutcard(_curuser, NULL, GetRecomemdOutCard(_curuser));
			}
		};
	}
	return false;
}

bool  CTableFrameSink::OnGameMessage(WORD wSubCmdID, VOID * pData, WORD wDataSize, IServerUserItem * pIServerUserItem)
{
	WORD chair = pIServerUserItem->GetChairID();

	switch (wSubCmdID)
	{
	case CMD_C_CALLBANKER:
	{
		CMD_C_CallBanker *callbanker = (CMD_C_CallBanker *)pData;
		if (callbanker->callscore < BANKER_SCORE_NO || callbanker->callscore > BANKER_SCORE_100_GOU) {
			return true;
		}
		OnUserCallBanker(chair, pIServerUserItem, callbanker->callscore);
		return true;
	};

	case CMD_C_CHANGEBOTTOMCARD:
	{
		CMD_C_ChangeBottomCard *chgcard = (CMD_C_ChangeBottomCard *)pData;
		if (chair != _banker)
			return true;




	}break;

	default:
		break;
	}
	return false;
}

bool  CTableFrameSink::OnFrameMessage(WORD wSubCmdID, VOID * pData, WORD wDataSize, IServerUserItem * pIServerUserItem)
{
	return false;
}

bool CTableFrameSink::OnActionUserSitDown(WORD wChairID,IServerUserItem * pIServerUserItem, bool bLookonUser)
{
	return true;
}

bool CTableFrameSink::OnActionUserStandUp(WORD wChairID,IServerUserItem * pIServerUserItem, bool bLookonUser)
{
	return true;
}

bool CTableFrameSink::OnActionUserOnReady(WORD wChairID,IServerUserItem * pIServerUserItem, VOID * pData, WORD wDataSize)
{
	if (((_gameoptions->wServerType) & GAME_GENRE_PERSONAL) != 0)
	{
		//cbGameRule[1] 为  2 、3 、4 、5, 0分别对应 2人 、 3人 、 4人 、 5人 、 2-5人 这几种配置
		BYTE *pGameRule = _tableframe->GetGameRule();
		switch(pGameRule[1])
		{
		case 2:
		case 3:
		case 4:
		case 5:
			{
				break;
			}
		case 0:
			{
				break;
			}
		default:
			ASSERT(false);
		}
	}

	return true;
}



void CTableFrameSink::ReadConfigInformation()
{
	tagCustomRule *pCustomRule = (tagCustomRule *)_gameoptions->cbCustomRule;
	ASSERT(pCustomRule);

}
bool CTableFrameSink::IsRoomCardScoreType()
{
	return (_tableframe->GetDataBaseMode() == 1) && (((_gameoptions->wServerType) & GAME_GENRE_PERSONAL) != 0);
}

bool CTableFrameSink::IsRoomCardTreasureType()
{
	return (_tableframe->GetDataBaseMode() == 0) && (((_gameoptions->wServerType) & GAME_GENRE_PERSONAL) != 0);
}

void CTableFrameSink::SendGameData(WORD chair, WORD subcmd, VOID *pData, WORD size) {

}

void CTableFrameSink::SetGameBaseScore(LONG lBaseScore) {

}
*/