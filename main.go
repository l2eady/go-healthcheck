package main

// package main

import (
	"flag"
	"fmt"
	"go-healthcheck/internal/app/healthycheck"
	"log"
	"os"
	"sync"
	"time"
)

var (
	wg                   = &sync.WaitGroup{}
	CountSuccess  uint64 = 0
	CountWebsites uint64 = 0
	CountFailure  uint64 = 0
)

func main() {
	start := time.Now()
	fileName := flag.String("filename", "websites.csv", "websites csv file")
	flag.Parse()
	log.Println("Perform website checking...")
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("An error encountered ::", err)
	}
	defer file.Close()
	fmt.Println("====================== IMPLEMENT ================")
	service := healthycheck.NewHealthyCheckService(time.Second * 2)

	result, _ := service.HealthyCheckEndPointFromCSVFile(file)
	fmt.Println("====================== IMPLEMENT ================")
	for _, v := range result {
		if v.Status {

		}
	}
	fmt.Printf("Checked webistes: %v\n", CountWebsites)
	fmt.Printf("Successful websites:%v\n", CountSuccess)
	fmt.Printf("Failure websites: %v\n", CountFailure)
	fmt.Printf("Total times to finished checking website: %v nanoseconds\n", time.Since(start).Nanoseconds())

}

// func HealthyCheckWebservice(reader *csv.Reader) error {
// 	checker := checkers.New(time.Duration(time.Second))
// 	for {
// 		record, err := reader.Read()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			return errors.Wrapf(err, "unable to read")
// 		}

// 		wg.Add(1)
// 		url := record[0]
// 		go func(url string, wg *sync.WaitGroup) {
// 			ok, _ := checker.Ping(url)
// 			defer wg.Done()
// 			calReportHealthyCheck(record[0], ok)
// 		}(url, wg)
// 	}

// 	wg.Wait()
// 	return nil
// }

// func calReportHealthyCheck(url string, success bool) {
// 	atomic.AddUint64(&CountWebsites, 1)
// 	if !success {
// 		atomic.AddUint64(&CountFailure, 1)
// 		return
// 	}
// 	atomic.AddUint64(&CountSuccess, 1)
// }
