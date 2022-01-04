package service

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strings"
	"time"
)

var Bot *tb.Bot

func Start() {
	var err error
	Bot, err = tb.NewBot(tb.Settings{
		URL:    "https://api.telegram.org",
		Token:  c.Bot.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Printf("Bot 启动失败啦...... \n当前Token [ %s ] \n错误信息:  %s", c.Bot.Token, err)
		os.Exit(1)
	}

	setHandle()
	Bot.Start()
}

func setHandle() {
	Bot.Handle("/start", startCmdCtr)
	Bot.Handle("/help", startCmdCtr)
	Bot.Handle("/checkin", checkinCmdCtr)
	Bot.Handle("/account", accountCmdCtr)
	Bot.Handle("/bind", bindCmdCtr)
	Bot.Handle("/unbind", unbindCmdCtr)
}

func startCmdCtr(m *tb.Message) {
	menu := &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	CheckinBtn := menu.Text("🌈每日签到")
	AccountBtn := menu.Text("🧚🏻‍账户信息")
	BindBtn := menu.Text("💍绑定账户")
	UnbindBtn := menu.Text("🪖解绑账户")

	menu.Reply(
		menu.Row(CheckinBtn, AccountBtn),
		menu.Row(BindBtn, UnbindBtn),
	)

	Bot.Handle(&CheckinBtn, checkinCmdCtr)
	Bot.Handle(&AccountBtn, accountCmdCtr)
	Bot.Handle(&BindBtn, bindCmdCtr)
	Bot.Handle(&UnbindBtn, unbindCmdCtr)

	msg := fmt.Sprintf("%s机器人🤖️\n为你提供以下服务:\n\n每日签到 /checkin\n账户信息 /account\n绑定账户 /bind\n解绑账户 /unbind", c.Bot.Name)
	_, _ = Bot.Send(m.Chat, msg, menu)
}

func checkinCmdCtr(m *tb.Message) {
	user := QueryUser(m.Chat.ID)
	if user.Id <= 0 {
		msg := "⛔️当前未绑定账户\n请发送 /bind <订阅地址> 绑定账户\n\n#示例\n/bind https://域名/api/v1/client/subscribe?token=c09a65fd29cb8453926642c0db2e74c0"
		_, _ = Bot.Send(m.Sender, msg)
		return
	}
	if user.PlanId <= 0 {
		msg := "⛔当前暂无订阅计划,请购买后才能签到赚取流量😯..."
		_, _ = Bot.Send(m.Sender, msg)
		return
	}

	cc := CheckinTime(m.Chat.ID)
	if cc == false {
		msg := fmt.Sprintf("🥳今天已经签到过啦...")
		_, _ = Bot.Send(m.Sender, msg)
		return
	}

	uu := checkinUser(m.Chat.ID)

	msg := fmt.Sprintf("💍签到成功\n本次签到获得 %s 流量\n下次签到时间: %s", ByteSize(uu.CheckinTraffic), UnixToStr(uu.NextAt))
	_, _ = Bot.Send(m.Sender, msg)
}

func accountCmdCtr(m *tb.Message) {
	user := QueryUser(m.Chat.ID)
	if user.Id <= 0 {
		msg := "⛔️当前未绑定账户\n请发送 /bind <订阅地址> 绑定账户\n\n#示例\n/bind https://域名/api/v1/client/subscribe?token=c09a65fd29cb8453926642c0db2e74c0"
		_, _ = Bot.Send(m.Sender, msg)
		return
	}
	p := QueryPlan(int(user.PlanId))
	Email := user.Email
	CreatedAt := UnixToStr(user.CreatedAt)
	Balance := user.Balance / 100
	CommissionBalance := user.CommissionBalance / 100
	PlanName := p.Name
	ExpiredAt := UnixToStr(user.ExpiredAt)
	TransferEnable := ByteSize(user.TransferEnable)
	U := ByteSize(user.U)
	D := ByteSize(user.D)
	S := ByteSize(user.TransferEnable - (user.U + user.D))
	if user.PlanId <= 0 {
		msg := fmt.Sprintf("🧚🏻账户信息概况:\n\n当前绑定账户: %s\n注册时间: %s\n账户余额: %d元\n佣金余额: %d元\n\n当前订阅: 当前暂无订阅计划", Email, CreatedAt, Balance, CommissionBalance)
		_, _ = Bot.Send(m.Sender, msg)
		return
	}

	msg := fmt.Sprintf("🧚🏻账户信息概况:\n\n当前绑定账户: %s\n注册时间: %s\n账户余额: %d元\n佣金余额: %d元\n\n当前订阅: %s\n到期时间: %s\n订阅流量: %s\n已用上行: %s\n已用下行: %s\n剩余可用: %s", Email, CreatedAt, Balance, CommissionBalance, PlanName, ExpiredAt, TransferEnable, U, D, S)
	_, _ = Bot.Send(m.Sender, msg)

}

func bindCmdCtr(m *tb.Message) {
	user := QueryUser(m.Chat.ID)
	if user.Id > 0 {
		_, _ = Bot.Send(m.Sender, fmt.Sprintf("⭐您当前绑定账户: %s\n若需要修改绑定,请先解绑当前账户！", user.Email))
		return
	}

	format := strings.Index(m.Text, "token=")
	if format <= 0 {
		_, _ = Bot.Send(m.Sender, "⭐️️账户绑定格式: /bind <订阅地址>\n\n 发送示例：\n/bind https://域名/api/v1/client/subscribe?token=c09a65fd29cb8453926642c0db2e74c0")
		return
	}

	b := BindUser(m.Text[format:], m.Chat.ID)
	if b.Id <= 0 {
		_, _ = Bot.Send(m.Sender, "❌订阅无效,请前往官网复制最新订阅地址!")
		return
	}

	if b.TelegramId != uint(m.Chat.ID) {
		_, _ = Bot.Send(m.Sender, "❌账户绑定失败,请稍后再试")
	}
	_, _ = Bot.Send(m.Sender, fmt.Sprintf("💍账户绑定成功: %s", b.Email))
}

func unbindCmdCtr(m *tb.Message) {
	user := unbindUser(m.Chat.ID)
	if user.Id <= 0 {
		_, _ = Bot.Send(m.Sender, "⛔️当前未绑定账户")
		return
	}
	if user.TelegramId > 0 {
		_, _ = Bot.Send(m.Sender, "❌账户解绑失败,请稍后再试...")
		return
	}
	_, _ = Bot.Send(m.Sender, "🪖账户解绑成功")
}

func UnixToStr(unix int64) string {
	u := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	return u
}

func ByteSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%.2fB", float64(size)/float64(1))
	} else if size < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
	} else if size < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fTB", float64(size)/float64(1024*1024*1024*1024))
	}
}
