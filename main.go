package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/robfig/cron"
)

var IP_API_URL = Getenv("IP_API_URL", "http://api.ipify.org").(string)
var DNS_API_URL = Getenv("DNS_API_URL", "http://www.dtdns.com/api/autodns.cfm?").(string)
var DNS_HOSTNAME = Getenv("DNS_HOSTNAME", "").(string)
var DNS_PASSWD = Getenv("DNS_PASSWD", "").(string)
var UPDATE_INTERVAL = Getenv("UPDATE_INTERVAL", "*/1").(string)
var oldip string

func main() {

	c := cron.New()
	c.AddFunc("0 "+UPDATE_INTERVAL+" * * * *", routine)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func routine() {
	ip, err := getIp()
	if err != nil {
		log.Println(err)
	}
	log.Println(ip)

	if ip != oldip {
		log.Println("Ip changed to " + ip + ", updating DNS...")
		err := updateDns()
		if err != nil {
			log.Println(err)
		} else {
			oldip = ip
			log.Println("DNS Updated successully")
		}
	} else {
		log.Println("Ip not changed")
	}
}

func getIp() (string, error) {
	res, err := http.Get(IP_API_URL)
	if err != nil {
		return "", err
	}

	byteip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	ip := string(byteip[:])
	return ip, nil
}

func updateDns() error {
	res, err := http.Get(DNS_API_URL + "id=" + DNS_HOSTNAME + "&pw=" + DNS_PASSWD + "")
	if err != nil {
		return err
	}
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	body := string(bodyByte[:])
	if strings.Contains(body, "failed") || strings.Contains(body, "not valid") {
		return errors.New("DtDNS server: Hostname not valid or too many req failed")
	}
	return nil
}

func Getenv(key string, fallback interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	switch fallback.(type) {
	case string:
		var nw string
		nw = value
		return nw
	case uint:
		var nw uint64
		nw, _ = strconv.ParseUint(value, 10, 32)
		return uint(nw)
	case bool:
		var nw bool
		nw, _ = strconv.ParseBool(value)
		return nw
	case int:
		var nw int64
		nw, _ = strconv.ParseInt(value, 10, 32)
		return int(nw)
	default:
		panic("unrecognized escape character")
		return value
	}
}
