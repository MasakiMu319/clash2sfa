# clash2sfa
Used to convert subscription links in Clash.Meta format to sing-box format (tested to work properly on Android, iOS, SFM).

## Deploy
The `port` environment variable controls the port on which the program runs. If not set, it opens on port 8080 by default.

## docker
```
docker volume create clash2sfa    
docker run -d -p 8080:8080 -v clash2sfa:/server/db jiumumu/clash2singbox
```
## Usage
- After launching, use your browser to visit http://ip:port

- The New Profile in the sing-box Profiles fills in the remote link and allows you to switch nodes in Groups by starting a subscription.

## Template Profile
Most changes to the profile template will be preserved, as will adding nodes to the outbounds in the template.

## Support Protocol
- shadowsocks （Only support v2ray-plugin, obfs and shadow-tls plugin）
- shadowsocksR
- vmess
- vless (Include reality)
- trojan
- socks5
- http
- hysteria
- tuic5

## Change Log
### 0.2.2
- Adjusted logic for self-built nodes.Now, your self-build nodes will be shown at the beginning of select && urltest. 
### 0.2.1
- Fix the problem that adding groups caused the original HK and other regional groups to be lost.
### 0.2.0
- Added default country triage label, etc: HK,TW...;
- added delayed test connection for Apple service;
- rewrote filter unnecessary node judgment logic.