package main

import (
	"regexp"
	"testing"

	"github.com/honwen/golibs/cip"
	"github.com/honwen/golibs/domain"
	"github.com/stretchr/testify/assert"
)

func TestIp2locCN(t *testing.T) {
	assert.Equal(t, ip2locCN("202.96.128.86"), "[中国 广东 广州 电信]")
	assert.Equal(t, ip2locCN("202.96.209.133"), "[中国 上海 上海 电信]")
	assert.Equal(t, ip2locCN("219.141.136.10"), "[中国 北京 北京 电信]")

	assert.Equal(t, ip2locCN("210.22.70.3"), "[中国 上海 上海 联通]")
	assert.Equal(t, ip2locCN("123.123.123.123"), "[中国 北京 北京 联通]")

	assert.Equal(t, ip2locCN("223.87.238.22"), "[中国 四川 成都 移动]")

	assert.Equal(t, ip2locCN("101.6.6.6"), "[中国 北京 北京 教育网]")

	assert.Equal(t, ip2locCN("168.95.1.1"), "[中国 台湾 中华电信]")
	assert.Equal(t, ip2locCN("202.67.240.222"), "[中国 香港]")

	assert.Equal(t, ip2locCN("203.189.136.148"), "[柬埔寨 柬埔寨]")
	assert.Equal(t, ip2locCN("203.112.2.4"), "[日本 日本]")
	assert.Equal(t, ip2locCN("80.80.80.80"), "[荷兰 荷兰]")
	assert.Equal(t, ip2locCN("74.82.42.42"), "[美国 加利福尼亚州]")

}

func TestGetIPv4(t *testing.T) {
	funcs["myip"] = cip.MyIPv4
	ip4 := myip()
	assert.True(t, regexp.MustCompile(cip.RegxIPv4).MatchString(ip4) || len(ip4) == 0)
}

func TestGetIPv6(t *testing.T) {
	funcs["myip"] = cip.MyIPv6
	ip6 := myip()
	assert.True(t, regexp.MustCompile(cip.RegxIPv6).MatchString(ip6) || len(ip6) == 0)
}

func TestResloveIPv4(t *testing.T) {
	funcs["reslove"] = cip.ResloveIPv4
	assert.Contains(t, []string{"8.8.8.8", "8.8.4.4"}, reslove("dns.google"))
	assert.Contains(t, []string{"223.6.6.6", "223.5.5.5"}, reslove("dns.alidns.com"))
}

func TestResloveIPv6(t *testing.T) {
	funcs["reslove"] = cip.ResloveIPv6
	assert.Contains(t, []string{"2001:4860:4860::8844", "2001:4860:4860::8888"}, reslove("dns.google"))
	assert.Contains(t, []string{"2400:3200::1", "2400:3200:baba::1"}, reslove("dns.alidns.com"))
}

func TestSplitDomain001(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("a.example.com")

	assert.Equal(t, rr, "a")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain002(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("example.com")

	assert.Equal(t, rr, "@")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain003(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("*.example.com")

	assert.Equal(t, rr, "*")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain004(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("*.a.example.com")

	assert.Equal(t, rr, "*.a")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain005(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("*.b.a.example.com")

	assert.Equal(t, rr, "*.b.a")
	assert.Equal(t, domain, "example.com")
}
func TestSplitDomain006(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("a.example.co.kr")

	assert.Equal(t, rr, "a")
	assert.Equal(t, domain, "example.co.kr")
}

func TestSplitDomain007(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("*.a.example.co.kr")

	assert.Equal(t, rr, "*.a")
	assert.Equal(t, domain, "example.co.kr")
}

func TestSplitDomain008(t *testing.T) {
	rr, domain := domain.SplitDomainToRR("example.co.kr")

	assert.Equal(t, rr, "@")
	assert.Equal(t, domain, "example.co.kr")
}
