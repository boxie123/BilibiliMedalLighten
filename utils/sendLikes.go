package utils

import (
	"log"
	"net/http"
)

const ApiLikeReport = "https://api.live.bilibili.com/xlive/app-ucenter/v1/like_info_v3/like/likeReportV3"

// SendLikeReport 为特定直播间点赞 50 次
func SendLikeReport(client *http.Client, cookie string, csrf string, medalInfo MedalInfo) error {
	paramsMap := map[string]interface{}{
		"uid":        medalInfo.Medal.UID,
		"room_id":    medalInfo.RoomInfo.RoomID,
		"anchor_id":  medalInfo.Medal.TargetID,
		"csrf":       csrf,
		"csrf_token": csrf,
		"click_time": 50,
	}

	_, err := PostApiResponseData(client, cookie, ApiLikeReport, paramsMap)
	if err != nil {
		return err
	}

	log.Printf("已为%s-%d 房间点赞50次", medalInfo.AnchorInfo.NickName, medalInfo.RoomInfo.RoomID)
	return nil
}
