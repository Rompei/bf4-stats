package stat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (

	// Domain is Domain name of BF4 API
	Domain = "http://battlelog.battlefield.com"
)

// BF4Device is the type of device of BF4
type BF4Device string

const (
	PC      = "pc"
	XBOX360 = "xbox360"
	XBOXONE = "xboxone"
	PS3     = "ps3"
	PS4     = "ps4"
)

type (

	// StatResponse is object of BF4 API.
	StatResponse struct {
		Template      string        `json:"template"`
		GlobalContext GlobalContext `json:"globalContext"`
		Context       Context       `json:"context"`
	}

	// GlobalContext is object of BF4 API.
	GlobalContext struct {
		CurrentURL               string           `json:"currentUrl"`
		ActiveUserPage           string           `json:"activeUserPage"`
		IsAjaxFetch              bool             `json:"isAjaxFetch"`
		CurrentCononicalURL      string           `json:"currentConnonicalURL"`
		UseIdentity2RedirectFlow bool             `json:"useIdentity2RedirectFlow"`
		FlushStrage              bool             `json:"flushStrage"`
		Session                  Session          `json:"session"`
		SystemUtcTime            int64            `json:"systemUtcTime"`
		MutedSounds              int              `json:"mutedSounds"`
		BF4trialbannerActive     bool             `json:"bf4trialbannerActive"`
		StaticPrefix             string           `json:"staticPrefix"`
		Realm                    Realm            `json:"realm"`
		FocusGame                bool             `json:"focusGame"`
		PushConnections          []PushConnection `json:"pushConnections"`
		IsBFHVisible             bool             `json:"isBFHVisible"`
		RequestURI               string           `json:"requestURI"`
		NucleusID                string           `json:"nucleusId"`
		PushOperator             string           `json:"pushOperator"`
		WriteBanUntil            int              `json:"writeBanUntill"`
		ComponentName            string           `json:"componentName"`
		HasWriteBan              bool             `json:"hasWriteBan"`
		ActiveMenu               string           `json:"activeMenu"`
		FilePrefix               string           `json:"filePrefix"`
		IsAjaxNavigation         bool             `json:"isAjaxNavigation"`
		ActionName               string           `json:"actionName"`
		IsSecure                 bool             `json:"isSecure"`
		BanType                  int              `json:"banType"`
		HideComcenter            bool             `json:"hideComcenter"`
		CliendIP                 string           `json:"clientIp"`
		BfhbetabannerActive      bool             `json:"bfhbetabannerActive"`
	}

	// Session is object of BF4 API.
	Session struct {
		HasAdminAccess     bool             `json:"hasAdminAccess"`
		OriginLaunch       int              `json:"originLaunch"`
		UserProducts       []string         `json:"userProducts"`
		NpxExploration     NpxExploration   `json:"npxExploration"`
		Locale             Locale           `json:"locale"`
		IntroSectionBits   int              `json:"introSectionBits"`
		UserID             string           `json:"userId"`
		PostChecksum       string           `json:"postChecksum"`
		Facebook           Facebook         `json:"facebook"`
		User               User             `json:"user"`
		IsLoggedIn         bool             `json:"isLoggedIn"`
		Rollouts           []string         `json:"rollouts"`
		UserGameExpansions map[string][]int `json:"userGameExpansions"`
		UserGames          interface{}      `json:"userGames"`
	}

	// NpxExploration is object of BF4 API.
	NpxExploration struct {
		Quickmatch  bool `json:"quickmatch"`
		Competitive bool `json:"competitive"`
	}

	// Locale is object of BF4 API.
	Locale struct {
		Shordate string `json:"shortdate"`
		Fulldate string `json:"fulldate"`
		Offset   int    `json:"offset"`
		Time     string `json:"time"`
	}

	// Facebook is object of BF4 API.
	Facebook struct {
		HasShared   bool `json:"hasShared"`
		IsSharing   bool `json:"isSharing"`
		IsConnected bool `json:"isConnected"`
	}

	// User is object of BF4 API.
	User struct {
		Username    string      `json:"username"`
		GravatarMd5 interface{} `json:"gravatarMd5"`
		UserID      string      `json:"userId"`
		CreatedAt   int64       `json:"createdAt"`
		Presence    Presence    `json:"presence"`
	}

	// Presence is object of BF4 API.
	Presence struct {
		IsOnline       bool       `json:"isOnline"`
		UserID         string     `json:"userId"`
		UpdatedAt      int64      `json:"updatedAt"`
		PresenceStates string     `json:"presenceStates"`
		OnlineGame     OnlineGame `json:"onlineGame"`
		PlayingMp      PlayingMp  `json:"playingMp"`
		IsPlaying      bool       `json:"isPlaying"`
	}

	// OnlineGame is object of BF4 API.
	OnlineGame struct {
		Platform  int    `json:"platform"`
		Game      int    `json:"game"`
		PersonaID string `json:"personaId"`
	}

	// PlayingMp is object of BF4 API.
	PlayingMp struct {
		GameID         string `json:"gameId"`
		GameExpansions []int  `json:"gameExpansions"`
		GameMode       string `json:"gameMode"`
		ServerGuID     string `json:"serverGuid"`
		Game           int    `json:"game"`
		LevelName      string `json:"levelName"`
		PersonaID      string `json:"personaId"`
		ServerName     string `json:"serverName"`
		Experience     int    `json:"experience"`
		Platform       int    `json:"platform"`
		Role           int    `json:"role"`
	}

	// Realm is object of BF4 API.
	Realm struct {
		Lang    int `json:"lang"`
		Game    int `json:"game"`
		Section int `json:"section"`
	}

	// PushConnection is object of BF4 API.
	PushConnection struct {
		SslAddress      string   `json:"sslAddress"`
		APIKey          string   `json:"apiKey"`
		NodeNumber      int      `json:"nodeNumber"`
		ChannelWithAuth []string `json:"channelWithAuth"`
		SslEnabled      bool     `json:"sslEnabled"`
		SslWsAddress    string   `json:"sslWsAddress"`
		Channels        []string `json:"channels"`
		Token           string   `json:"token"`
		User            string   `json:"user"`
		Address         string   `json:"address"`
		WsAddress       string   `json:"wsAddress"`
		Log             bool     `json:"log"`
	}

	// Context is object of BF4 API.
	Context struct {
		Href                        string                       `json:"_href"`
		UserGameExpansions          map[string][]int             `json:"userGameExpansions"`
		UnreadGameReports           int                          `json:"unreadGameReports"`
		PersonaID                   int64                        `json:"personaId"`
		PersonaName                 string                       `json:"personaName"`
		NanigansPID                 string                       `json:"nanigansPid"`
		Notification                interface{}                  `json:"notification"`
		StatsPersona                StatsPersona                 `json:"statsPersona"`
		PersonaPicture              string                       `json:"personaPicture"`
		CookieConsentEnebled        bool                         `json:"cookieConsentEnebled"`
		UserGames                   map[string][]int             `json:"userGames"`
		SociadSharingEnebled        bool                         `json:"socialSharingEnabled"`
		ProfileCommon               ProfileCommon                `json:"profileCommon"`
		UserMessageBox              interface{}                  `json:"userMessageBox"`
		PlatformInt                 int                          `json:"platformInt"`
		SearchUesrs                 interface{}                  `json:"searchUsers"`
		LogClientSideErrors         bool                         `json:"logClientSideErrors"`
		EnableVideoBackground       string                       `json:"enableVideoBackground"`
		SearchPersonas              interface{}                  `json:"searchPersonas"`
		KitMap                      map[string]map[string]string `json:"kitMap"`
		TitleSection                string                       `json:"titleSection"`
		HasValidEntitlement         bool                         `json:"hasValidEntitlement"`
		FacebookOpenGraphDisableBFH bool                         `json:"facebookOpenGraphDisableBFH"`
	}

	// StatsPersona is object of BF4 API.
	StatsPersona struct {
		Picture                  string            `json:"picture"`
		UserID                   string            `json:"userid"`
		User                     interface{}       `json:"user"`
		UpdatedAt                int64             `json:"updatedAt"`
		FirstPartyID             string            `json:"firstPartyId"`
		PersonaID                string            `json:"personaId"`
		GameLegacy               string            `json:"gameLegacy"`
		NameSpace                string            `json:"nameSpace"`
		GameJSON                 string            `json:"gameJson"`
		Games                    map[string]string `json:"games"`
		ClanTag                  string            `json:"clanTag"`
		Club                     Club              `json:"club"`
		FacebookOpenGraphEnabled bool              `json:"facebookOpenGraphEnabled"`
		UnreadNotifications      int               `json:"unreadNotifications"`
		GameInt                  int               `json:"gameInt"`
	}

	// Club is object of BF4 API.
	Club struct {
		WebSite              string      `json:"webSite"`
		Stats                Stats       `json:"stats"`
		Status               string      `json:"status"`
		PublicWallStatus     int         `json:"publicWallStatus"`
		Emblem               interface{} `json:"emblem"`
		Leaders              interface{} `json:"leaders"`
		FormatedPresentation interface{} `json:"formatedPresentation"`
		Created              int64       `json:"created"`
		MemberCount          int         `json:"memberCount"`
		Presentation         string      `json:"presentation"`
		Tag                  string      `json:"tag"`
		AdminIDs             []string    `json:"adminIds"`
		PrivateWallStatus    int         `json:"privateWallStatus"`
		FounderID            string      `json:"founderId"`
		EmblemPath           string      `json:"emblemPath"`
		Members              interface{} `json:"members"`
		Country              interface{} `json:"country"`
		Founders             interface{} `json:"Founders"`
		Hidden               int         `json:"hidden"`
		Friends              interface{} `json:"friends"`
		ID                   string      `json:"id"`
		Name                 string      `json:"name"`
	}

	// ProfileCommon is object of BF4 API.
	ProfileCommon struct {
		FriendCount       int         `json:"friendCount"`
		VeteranStatus     interface{} `json:"veteranStatus"`
		Club              Club        `json:"club"`
		MutualFriends     interface{} `json:"mutualFriends"`
		BlazeClubs        interface{} `json:"blazeClub"`
		Profiletab        string      `json:"profiletab"`
		BlazeClubsLoaded  bool        `json:"blazeClubLoaded"`
		MutualFriendCount interface{} `json:"mutualFriendCount"`
		FriendStatus      string      `json:"friendStatus"`
		PlatoonFans       interface{} `json:"platoonFans"`
		User              User        `json:"user"`
		Platoons          interface{} `json:"platoons"`
		TenFriends        []Friend    `json:"tenFriends"`
		UserGames         interface{} `json:"userGames"`
		IsChatDisabled    bool        `json:"isChatDisabled"`
		UserStatusMessage interface{} `json:"userStatusMessage"`
		UserInfo          UserInfo    `json:"userInfo"`
	}

	// Stats is object of BF4 API.
	Stats struct {
		ScoreWeighted          int         `json:"scoreWeighted"`
		ClubID                 string      `json:"clubId"`
		Aggregated             bool        `json:"aggregated"`
		SkillAvg               float64     `json:"skillAvg"`
		TimeTotal              string      `json:"timeTotal"`
		CurrentRankScoreNeeded int         `json:"currentRankScoreNeeded"`
		SkillTotal             float64     `json:"skillTotal"`
		ScoreTotal             string      `json:"scoreTotal"`
		NextRankScoreNeeded    int         `json:"nextRankScoreNeeded"`
		TimeWeighted           int         `json:"timeWeighted"`
		ScoreAvg               float64     `json:"scoreAvg"`
		Game                   interface{} `json:"game"`
		Rank                   int         `json:"rank"`
		KillsTotal             string      `json:"killsTotal"`
		TimeAvg                float64     `json:"timeAvg"`
		KdrAvg                 float64     `json:"kdrAvg"`
		KillsAvg               float64     `json:"killsAvg"`
	}

	// Friend is object of BF4 API.
	Friend struct {
		UserID string `json:"userId"`
		User   User   `json:"user"`
	}

	// UserInfo is object of BF4 API.
	UserInfo struct {
		PrivacyFeedAndGameActivity int         `json:"privacyFeedAndGameActivity"`
		TwitchUsername             interface{} `json:"twitchUsername"`
		UserID                     string      `json:"userId"`
		IntroSectionBitmask        int         `json:"introSectionBitmask"`
		FeedHidden                 bool        `json:"feedHidden"`
		PushSettings               int         `json:"pushSettings"`
		ShowDetails                bool        `json:"showDetails"`
		ForumSignature             interface{} `json:"forumSignature"`
		Locality                   interface{} `json:"locality"`
		Location                   string      `json:"location"`
		Icon                       int         `json:"icon"`
		Presentation               interface{} `json:"presentation"`
		ProfileBlocked             int         `json:"profileBlocked"`
		AllowFriendRequests        bool        `json:"allowFriendRequests"`
		ShowFriendsUI              bool        `json:"showFriendsUI"`
		GravatarHidden             bool        `json:"gravatarHidden"`
		PresencePrivacy            int         `json:"presencePrivacy"`
		PresentationHidden         bool        `json:"presentationHidden"`
		LoginCounter               int         `json:"loginCounter"`
		PrivacyShowFriends         int         `json:"privacyShowFriends"`
		ForumSignatureHidden       bool        `json:"forumSignatureHidden"`
		Name                       interface{} `json:"name"`
		Age                        interface{} `json:"age"`
		Birthdate                  interface{} `json:"birthdate"`
		FeedActive                 bool        `json:"feedActive"`
		LastLogin                  int         `json:"lastLogin"`
		PrivacyDetails             int         `json:"privacyDetails"`
	}
)

// GetUserStats is returns user status.
func GetUserStats(userName, userID string, device BF4Device) (*StatResponse, error) {
	paths := []string{Domain, "bf4", "soldier", userName, "stats", userID, string(device)}
	url := strings.Join(paths, "/")

	cli := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-AjaxNavigation", "1")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res StatResponse
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
