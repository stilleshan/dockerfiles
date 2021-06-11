package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/honwen/golibs/cip"
	"github.com/honwen/ip2loc"
)

var funcs = map[string]interface{}{
	"myip":    cip.MyIPv4,
	"reslove": cip.ResloveIPv4,
}

func ip2locCN(ip string) (str string) {
	if strings.Count(ip, `.`) < 3 {
		return
	}
	if loc, err := ip2loc.IP2loc(ip); err != nil {
		log.Printf("%+v", err)
	} else {
		str = fmt.Sprintf("[%s %s %s %s]", loc.CountryName, loc.RegionName, loc.CityName, loc.IspDomain)
		for strings.Contains(str, " ]") {
			str = strings.ReplaceAll(str, " ]", "]")
		}
		for strings.Contains(str, "  ") {
			str = strings.ReplaceAll(str, "  ", " ")
		}
	}
	return
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func myip() (ip string) {
	if result, err := Call(funcs, "myip"); err == nil {
		for _, r := range result {
			return r.String()
		}
	}
	return
}

func reslove(domain string) (ip string) {
	if result, err := Call(funcs, "reslove", domain); err == nil {
		for _, r := range result {
			return r.String()
		}
	}
	return
}
