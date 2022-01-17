package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/denverdino/aliyungo/common"
	dns "github.com/honwen/aliyun-ddns-cli/alidns"
	"github.com/honwen/golibs/cip"
	"github.com/honwen/golibs/domain"
	"github.com/urfave/cli"
)

// AccessKey from https://ak-console.aliyun.com/#/accesskey
type AccessKey struct {
	ID             string
	Secret         string
	client         *dns.Client
	managedDomains []string
}

func (ak *AccessKey) getClient() *dns.Client {
	if len(ak.ID) <= 0 && len(ak.Secret) <= 0 {
		return nil
	}
	if ak.client == nil {
		ak.client = dns.NewClient(ak.ID, ak.Secret)
		ak.client.SetEndpoint(dns.DNSDefaultEndpointNew)
	}
	return ak.client
}

func (ak AccessKey) String() string {
	return fmt.Sprintf("Access Key: [ ID: %s ;\t Secret: %s ]", ak.ID, ak.Secret)
}

func (ak *AccessKey) ListManagedDomains() (domains []string, err error) {
	var resp []dns.DomainType
	resp, err = ak.getClient().DescribeDomains(
		&dns.DescribeDomainsArgs{
			Pagination: common.Pagination{PageSize: 50},
		})
	if err != nil {
		return
	}
	domains = make([]string, len(resp))
	for i, v := range resp {
		domains[i] = v.DomainName
	}
	return
}

func (ak *AccessKey) AutocheckDomainRR(rr, domain string) (r, d string, err error) {
	if contains(ak.managedDomains, domain) {
		return rr, domain, nil
	} else {
		if !strings.Contains(rr, `.`) {
			return "", "", fmt.Errorf("Domain [%s.%s] Not Managed", rr, domain)
		} else {
			rrs := strings.Split(rr, `.`)
			for i := len(rrs) - 1; i > 0; i-- {
				d = strings.Join(append(rrs[i:], domain), `.`)
				if contains(ak.managedDomains, d) {
					r = strings.Join(rrs[:i], `.`)
					return
				}
			}
		}
	}
	return "", "", fmt.Errorf("Domain [%s.%s] Not Managed", rr, domain)
}

func (ak *AccessKey) ListRecord(domain string) (dnsRecords []dns.RecordTypeNew, err error) {
	var resp *dns.DescribeDomainRecordsNewResponse
	for idx := 1; idx <= 99; idx++ {
		resp, err = ak.getClient().DescribeDomainRecordsNew(
			&dns.DescribeDomainRecordsNewArgs{
				DomainName: domain,
				Pagination: common.Pagination{PageNumber: idx, PageSize: 50},
			})
		if err != nil {
			return
		}
		dnsRecords = append(dnsRecords, resp.DomainRecords.Record...)
		if len(dnsRecords) >= resp.PaginationResult.TotalCount {
			return
		}
	}
	return
}

func (ak *AccessKey) DelRecord(rr, domain string) (err error) {
	var target *dns.RecordTypeNew
	if dnsRecords, err := ak.ListRecord(domain); err == nil {
		for i := range dnsRecords {
			if dnsRecords[i].RR == rr {
				target = &dnsRecords[i]
				_, err = ak.getClient().DeleteDomainRecord(
					&dns.DeleteDomainRecordArgs{
						RecordId: target.RecordId,
					},
				)
			}
		}
	} else {
		return err
	}
	return
}

func (ak *AccessKey) UpdateRecord(recordID, rr, dmType, value string, ttl int) (err error) {
	_, err = ak.getClient().UpdateDomainRecord(
		&dns.UpdateDomainRecordArgs{
			RecordId: recordID,
			RR:       rr,
			Value:    value,
			Type:     dmType,
			TTL:      json.Number(fmt.Sprint(ttl)),
		})
	return
}

func (ak *AccessKey) AddRecord(domain, rr, dmType, value string, ttl int) (err error) {
	_, err = ak.getClient().AddDomainRecord(
		&dns.AddDomainRecordArgs{
			DomainName: domain,
			RR:         rr,
			Type:       dmType,
			Value:      value,
			TTL:        json.Number(fmt.Sprint(ttl)),
		})
	return err
}

func (ak *AccessKey) CheckAndUpdateRecord(rr, domain, ipaddr, recordType string, ttl int) (err error) {
	fulldomain := strings.Join([]string{rr, domain}, `.`)
	if reslove(fulldomain) == ipaddr {
		return // Skip
	}
	targetCnt := 0
	var target *dns.RecordTypeNew
	if dnsRecords, err := ak.ListRecord(domain); err == nil {
		for i := range dnsRecords {
			if dnsRecords[i].RR == rr && dnsRecords[i].Type == recordType {
				target = &dnsRecords[i]
				targetCnt++
			}
		}
	} else {
		return err
	}

	if targetCnt > 1 {
		ak.DelRecord(rr, domain)
		target = nil
	}

	if target == nil {
		err = ak.AddRecord(domain, rr, recordType, ipaddr, ttl)
	} else if target.Value != ipaddr {
		if target.Type != recordType {
			return fmt.Errorf("record type error! oldType=%s, targetType=%s", target.Type, recordType)
		}
		err = ak.UpdateRecord(target.RecordId, target.RR, target.Type, ipaddr, ttl)
	}
	if err != nil && strings.Contains(err.Error(), `DomainRecordDuplicate`) {
		ak.DelRecord(rr, domain)
		return ak.CheckAndUpdateRecord(rr, domain, ipaddr, recordType, ttl)
	}
	return err
}

var (
	accessKey     AccessKey
	VersionString = "MISSING build version [git hash]"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := cli.NewApp()
	app.Name = "aliddns"
	app.Usage = "aliyun-ddns-cli"
	app.Version = fmt.Sprintf("Git:[%s] (%s)", strings.ToUpper(VersionString), runtime.Version())
	app.Commands = []cli.Command{
		{
			Name:     "list",
			Category: "DDNS",
			Usage:    "List AliYun's DNS DomainRecords Record",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Specific `DomainName`. like aliyun.com",
				},
			},
			Action: func(c *cli.Context) error {
				if err := appInit(c, true); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", accessKey, c.String("domain"))
				domain := c.String("domain")
				if !contains(accessKey.managedDomains, domain) {
					return fmt.Errorf("Domain [%s] Not Managed", domain)
				}
				if dnsRecords, err := accessKey.ListRecord(domain); err != nil {
					fmt.Printf("%+v", err)
				} else {
					for _, v := range dnsRecords {
						fmt.Printf("%20s   %-16s %s\n", v.RR+`.`+v.DomainName, fmt.Sprintf("%s(TTL:%4s)", v.Type, v.TTL), v.Value)
					}
				}
				return nil
			},
		},
		{
			Name:     "delete",
			Category: "DDNS",
			Usage:    "Delete AliYun's DNS DomainRecords Record",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Specific `FullDomainName`. like ddns.aliyun.com",
				},
			},
			Action: func(c *cli.Context) error {
				if err := appInit(c, true); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", accessKey, c.String("domain"))
				rr, domain, err := accessKey.AutocheckDomainRR(domain.SplitDomainToRR(c.String("domain")))
				if err != nil {
					return err
				}
				if err := accessKey.DelRecord(rr, domain); err != nil {
					fmt.Printf("%+v", err)
				} else {
					fmt.Println(c.String("domain"), "Deleted")
				}
				return nil
			},
		},
		{
			Name:     "update",
			Category: "DDNS",
			Usage:    "Update AliYun's DNS DomainRecords Record, Create Record if not exist",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Specific `DomainName`. like ddns.aliyun.com",
				},
				cli.StringFlag{
					Name:  "ipaddr, i",
					Usage: "Specific `IP`. like 1.2.3.4",
				},
				cli.IntFlag{
					Name:  "ttl, t",
					Value: 600,
					Usage: "The resolution effective time (in `seconds`)",
				},
			},
			Action: func(c *cli.Context) error {
				if err := appInit(c, true); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", accessKey, c.String("domain"), c.String("ipaddr"))
				rr, domain, err := accessKey.AutocheckDomainRR(domain.SplitDomainToRR(c.String("domain")))
				if err != nil {
					return err
				}
				recordType := "A"
				if c.GlobalBool("ipv6") {
					recordType = "AAAA"
				}
				if err := accessKey.CheckAndUpdateRecord(rr, domain, c.String("ipaddr"), recordType, c.Int("ttl")); err != nil {
					log.Printf("%+v", err)
				} else {
					log.Println(c.String("domain"), c.String("ipaddr"), ip2locCN(c.String("ipaddr")))
				}
				return nil
			},
		},
		{
			Name:     "auto-update",
			Category: "DDNS",
			Usage:    "Auto-Update AliYun's DNS DomainRecords Record, Get IP using its getip",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Specific `DomainName`. like ddns.aliyun.com",
				},
				cli.StringFlag{
					Name:  "redo, r",
					Value: "",
					Usage: "redo Auto-Update, every N `Seconds`; Disable if N less than 10; End with [Rr] enable random delay: [N, 2N]",
				},
				cli.IntFlag{
					Name:  "ttl, t",
					Value: 600,
					Usage: "The resolution effective time (in `seconds`)",
				},
			},
			Action: func(c *cli.Context) error {
				if err := appInit(c, true); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", accessKey, c.String("domain"), c.Int64("redo"))
				rr, domain, err := accessKey.AutocheckDomainRR(domain.SplitDomainToRR(c.String("domain")))
				if err != nil {
					return err
				}
				recordType := "A"
				if c.GlobalBool("ipv6") {
					recordType = "AAAA"
				}
				redoDurtionStr := c.String("redo")
				if len(redoDurtionStr) > 0 && !regexp.MustCompile(`\d+[Rr]?$`).MatchString(redoDurtionStr) {
					return errors.New(`redo format: [0-9]+[Rr]?$`)
				}
				randomDelay := regexp.MustCompile(`\d+[Rr]$`).MatchString(redoDurtionStr)
				redoDurtion := 0
				if randomDelay {
					redoDurtion, _ = strconv.Atoi(redoDurtionStr[:len(redoDurtionStr)-1])
				} else {
					redoDurtion, _ = strconv.Atoi(redoDurtionStr)
				}
				// Print Version if exist
				if redoDurtion > 0 && !strings.HasPrefix(VersionString, "MISSING") {
					fmt.Fprintf(os.Stderr, "%s %s\n", strings.ToUpper(c.App.Name), c.App.Version)
				}
				for {
					autoip := myip()
					if len(autoip) == 0 {
						log.Printf("# Err-CheckAndUpdateRecord: [%s]", "IP is empty, PLZ check network")
					} else {
						if err := accessKey.CheckAndUpdateRecord(rr, domain, autoip, recordType, c.Int("ttl")); err != nil {
							log.Printf("# Err-CheckAndUpdateRecord: [%+v]", err)
						} else {
							log.Println(c.String("domain"), autoip, ip2locCN(autoip))
						}
					}
					if redoDurtion < 10 {
						break // Disable if N less than 10
					}
					if randomDelay {
						time.Sleep(time.Duration(redoDurtion+rand.Intn(redoDurtion)) * time.Second)
					} else {
						time.Sleep(time.Duration(redoDurtion) * time.Second)
					}
				}
				return nil
			},
		},
		{
			Name:     "getip",
			Category: "GET-IP",
			Usage:    fmt.Sprintf("      Get IP Combine 10+ different Web-API"),
			Action: func(c *cli.Context) error {
				if err := appInit(c, false); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", c.Command.Usage)
				ip := myip()
				fmt.Println(ip, ip2locCN(ip))
				return nil
			},
		},
		{
			Name:     "resolve",
			Category: "GET-IP",
			Usage:    fmt.Sprintf("      Get DNS-IPv4 Combine 4+ DNS Upstream"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain, d",
					Usage: "Specific `DomainName`. like ddns.aliyun.com",
				},
			},
			Action: func(c *cli.Context) error {
				if err := appInit(c, false); err != nil {
					return err
				}
				// fmt.Println(c.Command.Name, "task: ", c.Command.Usage)
				ip := reslove(c.String("domain"))
				fmt.Println(ip, ip2locCN(ip))
				return nil
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "access-key-id, id",
			Usage: "AliYun's Access Key ID",
		},
		cli.StringFlag{
			Name:  "access-key-secret, secret",
			Usage: "AliYun's Access Key Secret",
		},
		cli.StringSliceFlag{
			Name:  "ipapi, api",
			Usage: "Web-API to Get IP, like: http://myip.ipip.net",
		},
		cli.BoolFlag{
			Name:  "ipv6, 6",
			Usage: "IPv6",
		},
	}
	app.Action = func(c *cli.Context) error {
		return appInit(c, true)
	}
	app.Run(os.Args)
}

func appInit(c *cli.Context, checkAccessKey bool) error {
	akids := []string{c.GlobalString("access-key-id"), os.Getenv("AKID"), os.Getenv("AccessKeyID")}
	akscts := []string{c.GlobalString("access-key-secret"), os.Getenv("AKSCT"), os.Getenv("AccessKeySecret")}
	sort.Sort(sort.Reverse(sort.StringSlice(akids)))
	sort.Sort(sort.Reverse(sort.StringSlice(akscts)))
	accessKey.ID = akids[0]
	accessKey.Secret = akscts[0]
	if checkAccessKey && accessKey.getClient() == nil {
		cli.ShowAppHelp(c)
		return errors.New("access-key is empty")
	}
	if domains, err := accessKey.ListManagedDomains(); err == nil {
		// log.Println(domains)
		accessKey.managedDomains = domains
	} else {
		cli.ShowAppHelp(c)
		return errors.New("No Managed Domains")
	}

	if c.GlobalBool("ipv6") {
		funcs["myip"] = cip.MyIPv6
		funcs["reslove"] = cip.ResloveIPv6
	}

	ipapi := []string{}
	for _, api := range c.GlobalStringSlice("ipapi") {
		if !regexp.MustCompile(`^https?://.*`).MatchString(api) {
			api = "http://" + api
		}
		if regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`).MatchString(api) {
			ipapi = append(ipapi, api)
		}
	}
	if len(ipapi) > 0 {
		regx := regexp.MustCompile(cip.RegxIPv4)
		if c.GlobalBoolT("ipv6") {
			regx = regexp.MustCompile(cip.RegxIPv6)
		}
		funcs["myip"] = func() string {
			return cip.FastWGetWithVailder(ipapi, func(s string) string {
				return regx.FindString((s))
			})
		}
	}

	return nil
}
