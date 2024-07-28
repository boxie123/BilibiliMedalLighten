package main

import (
	"fmt"
	"github.com/boxie123/BilibiliMedalLighten/utils"
	login "github.com/boxie123/GoBilibiliLogin"
	"log"
	"net/http"
)

func main() {
	cookie, csrf, _ := login.Login()
	client := &http.Client{}
	log.Println("开始爬取粉丝勋章列表")
	medalList, err := utils.GetLightedMedalList(client, cookie)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("爬取完成，获取到%d个需要点赞的勋章\n", len(medalList))
	log.Println("开始点赞")
	for _, medalInfo := range medalList {
		err := utils.SendLikeReport(client, cookie, csrf, medalInfo)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	log.Println("点赞完成，按回车退出")
	fmt.Scanln()
}
