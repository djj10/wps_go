package service_discovery

import (
	"github.com/hashicorp/consul/api"
)

func RegisterService(serviceName string, port int) error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceName
	registration.Name = serviceName
	registration.Port = port
	registration.Check = &api.AgentServiceCheck{
		HTTP:     "http://localhost:8080/health",
		Interval: "10s",
		Timeout:  "5s",
	}

	return client.Agent().ServiceRegister(registration)
}

func DeregisterService(serviceName string) error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	return client.Agent().ServiceDeregister(serviceName)
}
