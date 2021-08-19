package hs

import (
	"os"
	"strconv"

	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var hs = `https://hs.fbigame.com/ajax.php`

var param = req.Param{
	"mod":          `get_cards_list`,
	"mode":         `-1`,
	"extend":       `-1`,
	"mutil_extend": ``,
	"hero":         `-1`,
	"rarity":       `-1`,
	"cost":         `-1`,
	"mutil_cost":   ``,
	"techlevel":    `-1`,
	"type":         `-1`,
	"collectible":  `-1`,
	"isbacon":      `-1`,
	"page":         `1`,
	"search_type":  `1`,
	"deckmode":     "normal",
	"hash":         "7f7b68fc",
}
var header = req.Header{
	"user-agent": `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36`,
	"referer":    `https://hs.fbigame.com`,
}

func init() {
	zero.OnRegex(`^搜卡(.+)$`).
		SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
		List := ctx.State["regex_matched"].([]string)[1]
		r, _ := req.Get(hs, header, param, req.Param{"search": List})
		g := r.String()
		im, _ := req.Get(`https://res.fbigame.com/hs/v13/`+
			gjson.Get(g, `list.0.CardID`).String()+
			`.png?auth_key=`+
			gjson.Get(g, `list.0.auth_key`).String(), header)
		im.ToFile("1.png")
		file, _ := os.Open("1.png")
		sg, _ := req.Post("https://pic.sogou.com/pic/upload_pic.jsp", req.FileUpload{
			File:      file,
			FieldName: "image",      // FieldName 是表单字段名
			FileName:  "avatar.png", // Filename 是要上传的文件的名称，我们使用它来猜测mimetype，并将其上传到服务器上
		})
		var tx string
		t := int(gjson.Get(g, `list.#`).Int())
		if t == 0 {
			ctx.SendChain(message.Text("查询为空！"))
			return
		}
		for i := 0; i < t && i < 10; i++ {
			tx += strconv.Itoa(i+1) + ". " +
				gjson.Get(g, `list.`+strconv.Itoa(i)+`.CARDNAME`).String() + "\n"
		}
		ctx.SendChain(
			message.Image(sg.String()),
			message.Text(tx),
		)
	})
}
