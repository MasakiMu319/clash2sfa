# clash2sfa
用于将 Clash.Meta 格式的订阅链接转换为 sing-box 格式（测试在 Android、iOS、SFM 正常工作）。

## 部署
环境变量 `port` 控制程序运行所在的端口，若未设置默认开放在 8080 端口。

## docker
```
docker volume create clash2sfa    
docker run -d -p 8080:8080 -v clash2sfa:/server/db jiumumu/clash2singbox
```
## 使用
启动后使用浏览器访问 http://ip:port

sing-box Profiles 中 New Profile填入 remote 链接，可以通过启动订阅在 Groups 切换节点。

## 配置文件模板
对配置文件模板中大多数修改都将被保留，在模板中的 outbounds 中增加节点也会被保留。

## 可转换的协议
- shadowsocks （仅包含 v2ray-plugin, obfs 和 shadow-tls 插件）
- shadowsocksR
- vmess
- vless (含 reality)
- trojan
- socks5
- http
- hysteria
- tuic5