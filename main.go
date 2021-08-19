package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_imagemagick" // 基础词库

	// 以下为内置依赖，勿动
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var content = []string{
	"* OneBot + ZeroBot + Golang ",
	"* Version 1.0.4 - 2021-07-14 14:09:58.581489207 +0800 CST",
	"* Copyright © 2020 - 2021  Kanri, DawnNights, Fumiama, Suika",
	"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
}

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Print(
		"====================[ZeroBot-Plugin]====================",
		"\n", strings.Join(content, "\n"), "\n",
		"========================================================",
	) // 启动打印
	zero.Run(zero.Config{
		NickName:      []string{"椛椛", "ATRI", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "/",

		// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
		// []string{}：通过代码写死的方式添加主人账号
		// os.Args[1:]：通过命令行参数的方式添加主人账号
		SuperUsers: append([]string{"2031605105"}, os.Args[1:]...),

		Driver: []zero.Driver{
			&driver.WSClient{
				// OneBot 正向WS 默认使用 6700 端口
				Url:         "ws://101.91.230.253:16700",
				AccessToken: "124816",
			},
		},
	})

	// 帮助
	zero.OnFullMatchGroup([]string{"help", "/help", ".help", "菜单", "帮助"}, zero.OnlyToMe).SetBlock(false).SetPriority(999).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(
				message.Text(strings.Join(content, "\n")),
			)
		})
	select {}
}
