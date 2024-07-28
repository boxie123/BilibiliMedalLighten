package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ApiGetMedalList = "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/panel?page=%d&page_size=50"

// GetLightedMedalList 获取所有已经点亮，且需要任务维持的牌子
func GetLightedMedalList(client *http.Client, cookie string) ([]MedalInfo, error) {
	page := 1
	var allLightedMedalList []MedalInfo // 需要点赞的牌子
	for {
		medalData, err := getMedalList(client, cookie, page)
		if err != nil {
			return nil, err
		}
		for _, medalInfo := range medalData.SpecialList {
			if medalInfo.Medal.IsLighted == 1 && medalInfo.Medal.GuardLevel == 0 { // 已经灰了的和在舰的不需要点赞
				allLightedMedalList = append(allLightedMedalList, medalInfo)
			}
			if medalInfo.Medal.IsLighted != 1 {
				break
			}
		}
		for _, medalInfo := range medalData.List {
			if medalInfo.Medal.IsLighted == 1 && medalInfo.Medal.GuardLevel == 0 { // 已经灰了的和在舰的不需要点赞
				allLightedMedalList = append(allLightedMedalList, medalInfo)
			}
			if medalInfo.Medal.IsLighted != 1 {
				return allLightedMedalList, nil
			}
		}
		if !medalData.PageInfo.HasMore {
			break
		}
		page += 1
	}
	return allLightedMedalList, nil
}

// 获取一页牌子列表
func getMedalList(client *http.Client, cookie string, page int) (MedalData, error) {
	apiUrl := fmt.Sprintf(ApiGetMedalList, page)
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return MedalData{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return MedalData{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var medalResp ApiRespenseMedalList
	err = json.NewDecoder(resp.Body).Decode(&medalResp)
	if err != nil {
		return MedalData{}, err
	}

	if medalResp.Code != 0 {
		return MedalData{}, fmt.Errorf("post response error: %s", medalResp.Message)
	}
	return medalResp.Data, nil
}
