<div align="center">

[![BilibiliMedalLighten](https://user-images.githubusercontent.com/74038190/212747903-e9bdf048-2dc8-41f9-b973-0e72ff07bfba.gif)](https://boxie123.github.io/Bilibili-GetReceivedGiftStream/)

# BilibiliMedalLighten

替你在b站上班

</div>

## 功能介绍
b站更改了粉丝勋章点亮规则，如果想维持已有的牌子持续点亮，需要每天完成任一任务：“50个点赞、观看20分钟直播、发10条弹幕、投喂电池”，每天做任务如同在上班。

本程序爬取已登录账号中需要做任务维持点亮的粉丝勋章列表，并给每个直播间点50个赞，在不影响主播弹幕的同时维持牌子点亮。

## 使用方法
### 1. 使用已构建好的可执行文件（推荐）
下载[`release`](https://github.com/boxie123/BilibiliMedalLighten/releases/latest)中的`.exe`文件，双击运行，根据文字提示操作。

### 2. 手动构建（不推荐）
下载本仓库，并运行：
```shell
go run ./main.go
```

## 注意事项
1. 每人每天点赞有上限（大概是几万，没有仔细测），达到上限后无法为任何直播间点赞，所以本程序**最多维持数百个牌子点亮**，如你的需求量更大，推荐移步隔壁自动发弹幕脚本。
2. **程序自动生成的`bzcookie.json`为重要账号登录信息，请谨慎保管、切勿泄漏，否则可能导致账号被盗用等后果。**

> [!CAUTION]
> 请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！
>
> 本项目遵守 MIT 协议，滥用造成的一切不良后果与本人无关！

## 支持与反馈

- 右上角点一下💫Star支持作者吧
- 如果在使用过程中有任何问题，欢迎移步[`issues`](https://github.com/boxie123/BilibiliMedalLighten/issues/new)反馈