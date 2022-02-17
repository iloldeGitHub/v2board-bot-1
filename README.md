# V2Board 外置面向用户 Bot

## 免责声明
此库 Fork 自 MiyaUU 删库后的一个 Fork 版本进行修改, 不保证任何可用性。有问题，请提issue，我也不一定会看（

## 注意
只能私信使用，不要拉到群里（反正你拉了也没用）。可独立部署，另外申请一个bot即可，不与官方设置BOT冲突，可以两个一起用。

## 如何使用它?

下载最新的 release ，找一个文件夹放好（比如 `/usr/UUBot/` ），再在当前文件夹新建一个 `uuBot.yaml` 文件，写上以下内容

```shell
# 机器人配置
bot:
  # 机器人名称
  name: ""
  token: ""
  # 签到可获取的最大值流量，单位是MB
  byte: 1024
  
# 数据库配置
database:
  host: "localhost"
  port: 3306
  name: ""
  username: ""
  password: ""
```

根据要求完成填写后保存，给予可执行权限（ `chmod +x` ）后运行即可。

## 关于进程守护
使用 `systemd` `Supervisor` `screen` 等等都可以。