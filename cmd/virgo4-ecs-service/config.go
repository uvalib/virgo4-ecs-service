package main

import (
	"flag"
	"log"
)

// ServiceConfig defines all of the service configuration parameters
type ServiceConfig struct {
	ClusterName  string // the cluster name
	ServiceName  string // the service name
	StartService bool   // do we start the service
}

// LoadConfiguration will load the service configuration from env/cmdline
// and return a pointer to it. Any failures are fatal.
func LoadConfiguration() *ServiceConfig {

	var cfg ServiceConfig

	flag.StringVar(&cfg.ClusterName, "cluster", "", "The cluster name")
	flag.StringVar(&cfg.ServiceName, "service", "", "The service name")
	flag.BoolVar(&cfg.StartService, "start", false, "Start the service")

	flag.Parse()

	if len(cfg.ClusterName) == 0 {
		log.Fatalf("ClusterName cannot be blank")
	}

	if len(cfg.ServiceName) == 0 {
		log.Fatalf("ServiceName cannot be blank")
	}

	return &cfg
}
