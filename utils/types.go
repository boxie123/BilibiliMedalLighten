package utils

// ApiResponseCommon bilibili api 普遍返回数据格式
type ApiResponseCommon struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UserInfo struct {
	RoomID int
	UserID int64
}

type ApiRespenseMedalList struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    MedalData `json:"data"`
}

type MedalData struct {
	List        []MedalInfo   `json:"list"`
	SpecialList []MedalInfo   `json:"special_list"`
	PageInfo    MedalPageInfo `json:"page_info"`
}

type MedalPageInfo struct {
	Number          int  `json:"number"`
	CurrentPage     int  `json:"current_page"`
	HasMore         bool `json:"has_more"`
	NextPage        int  `json:"next_page"`
	NextLightStatus int  `json:"next_light_status"`
	TotalPage       int  `json:"total_page"`
}

type MedalInfo struct {
	Medal      MedalInfoMedal      `json:"medal"`
	AnchorInfo MedalInfoAnchorInfo `json:"anchor_info"`
	RoomInfo   MedalInfoRoomInfo   `json:"room_info"`
}

type MedalInfoMedal struct {
	UID        int64  `json:"uid"`
	TargetID   int64  `json:"target_id"`
	TargetName string `json:"target_name"`
	Level      int    `json:"level"`
	MedalName  string `json:"medal_name"`
	TodayFeed  int    `json:"today_feed"`
	IsLighted  int    `json:"is_lighted"`
	GuardLevel int    `json:"guard_level"`
}

type MedalInfoAnchorInfo struct {
	NickName string `json:"nick_name"`
}

type MedalInfoRoomInfo struct {
	RoomID int `json:"room_id"`
}
