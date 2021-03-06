package models

import (
	"github.com/astaxie/beego"
	"github.com/miekg/dns"
	"gopkg.in/redis.v3"
	"strings"
	"time"
)

type Qdns struct {
	Domain string
	Answer []string
}

func ResolveFromDNS(domain string) (string, error) {
	answer := make([]string, 0)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	c := new(dns.Client)
	in, _, err := c.Exchange(m, beego.AppConfig.String("dns"))
	if err != nil {
		return strings.Join(answer, ","), err
	}
	for _, ain := range in.Answer {
		if a, ok := ain.(*dns.A); ok {
			answer = append(answer, a.A.String())
		}
	}
	answerstr := strings.Join(answer, ",")
	if err := cacheToRedis(domain, answerstr); err != nil {
		return answerstr, nil
	}
	return strings.Join(answer, ","), err
}

func ResolveFromDNS2(domain string) (*Qdns, error) {
	answer := make([]string, 0)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	c := new(dns.Client)
	in, _, err := c.Exchange(m, beego.AppConfig.String("dns"))
	if err != nil {
		qdns := &Qdns{Domain: domain, Answer: answer}
		return qdns, err
	}
	for _, ain := range in.Answer {
		if a, ok := ain.(*dns.A); ok {
			answer = append(answer, a.A.String())
		}
	}
	qdns := &Qdns{Domain: domain, Answer: answer}
	answerstr := strings.Join(answer, ",")
	if err := cacheToRedis(domain, answerstr); err != nil {
		// return answerstr, nil
		return qdns, err
	}
	return qdns, nil
}

func ResolveFromRedis2(domain string) (*Qdns, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: beego.AppConfig.String("redis")})
	strcmd, err := redisClient.Get(domain).Result()
	answer := strings.Split(strcmd, ",")
	if err == nil {
		qdns := &Qdns{Domain: domain, Answer: answer}
		return qdns, nil
	}
	str, errs := ResolveFromDNS2(domain)
	if errs == nil {
		return str, errs
	}
	return str, errs
}

func ResolveFromRedis(domain string) (string, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: beego.AppConfig.String("redis")})
	strcmd, err := redisClient.Get(domain).Result()
	if err == nil {
		return strcmd, nil
	}
	str, errs := ResolveFromDNS(domain)
	if errs == nil {
		return str, errs
	}
	return str, errs
}

func cacheToRedis(domain, dnsstr string) error {
	redisClient := redis.NewClient(&redis.Options{Addr: beego.AppConfig.String("redis")})
	err := redisClient.Set(domain, dnsstr, time.Duration(60)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}
