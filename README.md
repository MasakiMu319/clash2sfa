# clash2sfa

![last-commit](https://img.shields.io/github/last-commit/MasakiMu319/clash2sfa?style=for-the-badge)
![license](https://img.shields.io/github/license/MasakiMu319/clash2sfa?style=for-the-badge)

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
- hysteria 1/2
- tuic5

## Future Development Plans
- Refactor the front-end display page to provide more customization options; 
- Provide custom filtering options and support persistent operations; 
- Allow online editing and persistent storage of custom configuration templates; 
- Redesign the subscription update mechanism, considering the availability of nodes, there is no good idea for the time being.

## Change Log
### 0.3.2
- Add website icons designed by ChatGPT.
### 0.3.1
- Introduce log output(The logging feature is temporarily available and may be ~~refactored in a later version~~).
### 0.3.0
- Add support for hysteria2;
- Fix obfs field in ss;
- Rename some dependencies.
### 0.2.3
- Updated node filters to accommodate more airport node naming.
### 0.2.2
- Adjusted logic for self-built nodes.Now, your self-build nodes will be shown at the beginning of select && urltest. 
### 0.2.1
- Fix the problem that adding groups caused the original HK and other regional groups to be lost.
### 0.2.0
- Added default country triage label, etc: HK,TW...;
- added delayed test connection for Apple service;
- rewrote filter unnecessary node judgment logic.