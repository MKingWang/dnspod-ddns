# ddns by dnspod
功能说明：动态域名解析，针对有动态外网ip的家庭宽带用户，检测宽带外网ip，发生改变的情况下，将ip重新绑定到自己的dnspod的指定域名下

## 0x00 安装
build后的二进制文件和配置文件，放在任意位置（看各位客官的喜好）

## 0x01 配置说明
```ini
[Dnspod] ;Dnspod配置信息头，不能改动
token = id,token ;dnspod上申请的id和token组合
format = json  ;返回格式，默认json
domainid = 域名的id; 可以通过记录列表获取 参考手册：https://www.dnspod.cn/docs/records.html#record-list
recordid = 记录id;可以通过记录列表获取 参考手册：https://www.dnspod.cn/docs/records.html#record-list 
subdomain = www;记录值 默认www, 留空为@
```

## 0x02 crontab配置
```bash
#例如：将二进制文件和配置文件放在,/opt/dnspod
*/10 * * * * cd /opt/dnspod;./dnspod-dns
```
