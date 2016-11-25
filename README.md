## damdo/docker-dtdns

Dockerized dtdns client written in Go

#### FEATURES
- Lightweight Image - ~3.67MB (FROM scratch, no distro)
- Solid Cron Routine - based on github.com/robfig/cron
- Efficient - Updates only if ip has changed
- Flexible - Customizable with a set of parameters

#### USAGE
```sh
docker run -d -e DNS_HOSTNAME=yourhost.dtdns.domain -e DNS_PASSD=yourpasswd damdo/docker-dtdns
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
