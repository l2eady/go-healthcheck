package main

// package main

import (
	"flag"
	"go-healthcheck/configs"
	"go-healthcheck/internal/app/healthycheck"
	"log"
	"os"
	"sync"
	"time"
)

var (
	wg = &sync.WaitGroup{}
)

func main() {
	state := flag.String("state", "local", "set working environment")
	configPath := flag.String("config", "configs", "set configs path, default as: 'configs'")

	fileName := flag.String("filename", "example.csv", "csv filename for healthycheck")
	pingTimeOut := flag.Int64("ping_timeout_in_second", 2, "http timeout for ping")
	maxWorker := flag.Int("max_worker", 50, "maximum of worker for ping service")
	flag.Parse()

	conf, err := configs.New(*configPath, *state)
	if err != nil {
		log.Panicf("failed to load config, err: %v", err)
	}

	log.Println("Perform website checking...")
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("An error encountered ::", err)
	}
	defer file.Close()

	service := healthycheck.NewHealthyCheckService(time.Duration(*pingTimeOut)*time.Second, conf)
	report := service.HealthyCheckEndPointFromCSVFile(file, *maxWorker)

	log.Printf("Checked webistes: %v\n", len(report.Data))
	log.Printf("Successful websites:%v\n", report.TotalSuccess)
	log.Printf("Failure websites: %v\n", report.TotalFailure)
	log.Printf("Total times to finished checking website: %v nanoseconds\n", report.TotalTimeUsedInNano())

}
