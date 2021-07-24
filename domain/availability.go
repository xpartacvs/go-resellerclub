package domain

import "github.com/xpartacvs/go-resellerclub/core"

type DomainRegistration struct {
	Key    core.DomainKey           `json:"classkey"`
	Status DomainRegistrationStatus `json:"status"`
}

type DomainAvailabilities map[string]DomainRegistration
