package service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"os"
	"time"
)

type User struct {
	Id                uint
	TelegramId        uint
	Email             string
	Token             string
	U                 int64
	D                 int64
	PlanId            int64
	Balance           int64
	TransferEnable    int64
	CommissionBalance int64
	ExpiredAt         int64
	CreatedAt         int64
}

type Plan struct {
	Id   uint
	Name string
}

type UUBot struct {
	Id             uint `gorm:"primaryKey"`
	UserId         uint `gorm:"unique"`
	TelegramId     uint `gorm:"unique" `
	CheckinTraffic int64
	CheckinAt      int64
	NextAt         int64
}

var DB *gorm.DB
var c Conf

func init() {
	c.GetConfig()
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "v2_",
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("连接数据库失败... \n错误信息: %v", err)
		os.Exit(1)
	}
	if err = db.AutoMigrate(&UUBot{}); err != nil {
		fmt.Printf("数据库导入失败... \n错误信息: %v", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	DB = db
	return db
}

func QueryPlan(planId int) Plan {
	var plan Plan
	DB.Where("id = ?", planId).First(&plan)
	return plan
}

func QueryUser(tgId int64) User {
	var user User
	DB.Where("telegram_id = ?", tgId).First(&user)
	return user
}

func BindUser(token string, tgId int64) User {
	var user User
	DB.Where("token = ?", token[6:]).First(&user)
	if user.Id <= 0 {
		return user
	}
	if user.TelegramId <= 0 {
		DB.Model(&user).Update("telegram_id", tgId)
	}
	return user
}

func unbindUser(tgId int64) User {
	var user User
	DB.Where("telegram_id = ?", tgId).First(&user)
	if user.Id > 0 {
		DB.Model(&user).Update("telegram_id", nil)
		return user
	}
	return user
}

func CheckinTime(tgId int64) bool {
	var uu UUBot
	DB.Where("telegram_id = ?", tgId).First(&uu)
	if time.Now().Unix() < uu.NextAt {
		return false
	}
	return true
}

func checkinUser(tgId int64) UUBot {
	var user User
	var uu UUBot
	DB.Where("telegram_id = ?", tgId).First(&user)
	DB.Where("telegram_id = ?", tgId).First(&uu)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := r.Int63n(c.Bot.Byte)
	CheckIns := b * 1024 * 1024
	T := user.TransferEnable + CheckIns

	if uu.Id <= 0 {
		newUU := UUBot{
			UserId:         user.Id,
			TelegramId:     user.TelegramId,
			CheckinAt:      time.Now().Unix(),
			NextAt:         time.Now().Unix() + 86400,
			CheckinTraffic: 0,
		}
		DB.Create(&newUU)
	}

	DB.Model(&uu).Updates(UUBot{
		CheckinAt:      time.Now().Unix(),
		NextAt:         time.Now().Unix() + 86400,
		CheckinTraffic: CheckIns,
	})
	DB.Model(&user).Update("transfer_enable", T)

	return uu
}
