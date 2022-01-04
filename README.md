# v2board 签到机器人

```
免责声明:
此库仅本人个人学习开发并维护, 不保证任何可用性。
有问题 提issue
( 悠悠的 baby..🌹 Try it!!! )

  喜欢🥰 就用你的小jiojio 点一个 ⭐️ Star ! thank you. ->
```

- [写代码很容易，写成一坨~😶‍🌫️.. 能用即可](https://github.com/trekhleb/state-of-the-art-shitcode)
- 如果并不合您,请自行用❤️发电！

### UU说~
```
只能私信使用,不要拉到群里
可独立部署 另外申请一个bot即可,不与官方设置BOT冲突,可以两个一起用。不影响面板升级

其他功能 emmm 🧚🏻‍不嫌弃jiu 后续再更新...

这是我第一次用 go 不会写,大哥哥大姐姐🥱 不要喷wo, y～
```
#

<details>
<summary> 展开查看预览</summary>

![](uuBot.png)
</details>

## 如何使用它?

### 使用二进制文件部署(无需修改内容)

```shell
# 下载
将版本压缩包clone 到你的服务器

# 修改配置文件
修改 uuBot.yaml 配置

# 机器人配置
bot:
  name: "MiyaUU Bot"      # 机器人名称,回复信息时使用
  token: "5036:AAEhtXJJW" # 机器人Token @BotFather 申请
  byte: 1024              # 签到可获取的最大值流量,不能奸商模式🔨  单位是MB 1024 为 最多1GB
  
# 数据库配置
database:
  host: "localhost"        # 数据库地址 本地 或 ip
  port: 3306               # 数据库端口
  name: "v2board"          # 数据库名称
  username: "root"         # 数据库用户名
  password: "123123123"    # 数据库密码
```

### 运行

```shell
守护进程运行即可
由于时间关系 凌晨2点了 编程了夜猫子🥶. 先不写很多了,就看大家最最常用 PM2做示例
尽可能使用与面板同一台服务器部署,可以降低bot响应时间.

# 安装pm2 (应该大部分人能看到这的都会安装了
npm install pm2 -g

# 启动 biu~
二进制文件 uuBot 于 uuBot.yaml 要在同一个目录
cd -> 工作目录 pm2 start uuBot

启动成功后 pm2 进程会多一个 uuBot

supervisor 或者 docker 会用的应该也不用看教程了...

```