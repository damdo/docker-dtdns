## damdo/docker-dtdns
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/damdo/docker-dtdns/blob/master/LICENSE)
<br>
dockerized dtDNS client written in Go

#### FEATURES
- Lightweight Image - ~3.67MB (FROM scratch, no distro)
- Solid Cron Routine - based on github.com/robfig/cron
- Efficient - Updates only if ip has changed
- Flexible - Customizable with a set of parameters

#### USAGE
```sh
docker run -d -e DNS_HOSTNAME=yourhost.dtdns.domain -e DNS_PASSD=yourpasswd damdo/docker-dtdns
```
by default will check ip every minute

##### Optional Parameters
```sh
-e UPDATE_INTERVAL=*/n 
# where n is the amout of minutes between refresh

-e IP_API_URL=ip.api.url 
# where a different "get current ip" api can be specified

-e DNS_API_URL=new.dnsapi.url 
# where a different dns service api can be specified
```

#### HOW IT WORKS

It follows the guidelines defined by the DtDNS specification:
http://www.dtdns.com/dtsite/updatespec

<img src="img/docker-dtdns.png" height="400px" />

#### BUILDING FROM SOURCE

##### In DOCKER building
```sh
make IMAGENAME=desiredimagename
docker run -d -e DNS_HOSTNAME=yourhost.dtdns.domain -e DNS_PASSD=yourpasswd desiredimagename
```

##### Local building (go compiler required)
```sh
make localbuilder IMAGENAME=desiredimagename
docker run -d -e DNS_HOSTNAME=yourhost.dtdns.domain -e DNS_PASSD=yourpasswd desiredimagename
```

This project uses third party libraries that are distributed under their own terms (See LICENSE-3RD-PARTY).
