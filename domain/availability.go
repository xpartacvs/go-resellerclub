package domain

type DomainRegistration struct {
	Key    DomainKey                `json:"classkey"`
	Status DomainRegistrationStatus `json:"status"`
}

type DomainAvailabilities map[string]DomainRegistration
