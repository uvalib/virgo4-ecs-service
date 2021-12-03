package main

import (
	"fmt"
)

//
// main entry point
//
func main() {

	// Get config params and use them to init service context. Any issues are fatal
	cfg := LoadConfiguration()
	var err error

	if cfg.StartService {
		fmt.Printf("Starting %s...\n", cfg.ServiceName)
		// start the service
		err = serviceStart(cfg.ClusterName, cfg.ServiceName)
	} else {
		fmt.Printf("Stopping %s...\n", cfg.ServiceName)
		// stop the service
		err = serviceStop(cfg.ClusterName, cfg.ServiceName)
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}
}

//
// end of file
//
