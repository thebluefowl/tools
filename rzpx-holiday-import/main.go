package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var errLogger *log.Logger

func init() {
	errLogger = log.New(os.Stderr, "", 0)
}

type Event struct {
	Subject     string `csv:"Subject"`
	StartDate   string `csv:"Start Date"`
	StartTime   string
	EndDate     string
	EndTime     string
	IsAllDay    bool `csv:"All Day Event"`
	Description string
	Location    string
	Private     bool
}

func main() {
	file, err := os.Open("./source.txt")
	if err != nil {
		errLogger.Fatal(err)
	}

	events := []Event{}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e, err := toEvent(scanner.Text())
		if err != nil {
			errLogger.Fatal("failed to process line")
		}
		events = append(events, *e)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	csv, err := gocsv.MarshalBytes(events)
	if err != nil {
		errLogger.Fatal(err)
	}
	err = os.WriteFile("./calendar.csv", csv, 0644)
	if err != nil {
		errLogger.Fatal(err)
	}
}

func toEvent(line string) (*Event, error) {
	fmt.Println(line)
	parts := strings.Split(line, " - ")
	dateStr := strings.TrimSpace(parts[0])
	textStr := strings.TrimSpace(parts[1])

	startDate, err := parseDate(dateStr)
	if err != nil {
		return nil, err
	}
	event := new(Event)
	event.StartDate = startDate.Format("01/02/2006")
	event.Subject = parseText(textStr)
	event.IsAllDay = true

	return event, nil
}

func parseText(textStr string) string {
	return strings.Split(textStr, ".")[0]
}

func parseDate(dateStr string) (*time.Time, error) {
	replacer := strings.NewReplacer("nd", "", "rd", "", "th", "")
	dateStr = replacer.Replace(dateStr)

	layout := "Mon, 2 Jan"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	t = t.AddDate(time.Now().Year(), 0, 0)
	return &t, nil
}
