package models

import (
	"gopkg.in/redis.v3"
)

type QueryStatus struct {
	Status string
	Domain string
	Ipaddr string
}

type AgentStatus struct {
	Uuid    string
	Tag     string
	Ipaddr  string
	Version string
	Status  string
}

func GetAgentStatus(redisaddr string) (*AgentStatus, error) {
	agentstatus := new(AgentStatus)
	redisClient := redis.NewClient(&redis.Options{Addr: redisaddr})
	stucmd, err := redisClient.SMembers(agents).Result()
	if err != nil {
		return nil, err
	}
	return agentstatus, nil
}
