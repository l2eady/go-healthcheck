package checkers

import (
	"encoding/csv"
	"io"
	"log"
)

type CheckerHandler interface {
	HealthyCheckFromCSVFile(io.Reader) error
}

type checkerHandlerImpl struct {
	CheckerService
}

// NewCheckerHandler will createa a new checker handler
func NewCheckerHandler(checkerService CheckerService) CheckerHandler {
	return &checkerHandlerImpl{
		CheckerService: checkerService,
	}
}

func (h *checkerHandlerImpl) HealthyCheckFromCSVFile(reader io.Reader) error {
	csvReader := csv.NewReader(reader)
	for {
		record, err := csvReader.Read()
		if err != nil {
			switch err {
			case io.EOF:
				break
			default:
				log.Printf("Read file error: %#v \n", err)
				return err
			}
		}
		h.CheckerService.Ping(record[0])
	}
	return nil
}
