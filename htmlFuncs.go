package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	ISODateFormat      string = "2006-01-02"
	AmericanDateFormat string = "01/02/2006"

	PreparseTimeFormat string = "3:04 PM"
)

type TemplateStruct struct {
	Name             string
	ServiceDirection string
	Date             string
	Time             string
	Origin           string
}

func GetAvailableTimes(response string) []string {
	timesList := make([]string, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	timesAvailable := doc.Find(".bup-time-slots-divisor ")

	relevantDate := timesAvailable.First()

	timeSlots := relevantDate.Find(".bup-time-slots-available-list ")

	timesMap := make(map[string]bool)

	timeSlots.Each(func(index int, s *goquery.Selection) {
		children := s.Children()
		if children.Length() > 0 {
			children.Each(func(index int, li *goquery.Selection) {
				liTime := getListItemTime(li)
				if _, ok := timesMap[liTime]; !ok {
					timesMap[liTime] = true
				}
			})
		}
	})

	for k := range timesMap {
		timesList = append(timesList, k)
	}

	sortTimeList(timesList)

	return timesList
}

func LoadTemplate(formData map[string]string) (string, string) {
	rand.Seed(time.Now().UnixNano())

	template, err := template.ParseFiles("emailTemplate.html")
	if err != nil {
		log.Println(err)
	}

	var messageBody bytes.Buffer

	shuttleTime := strings.Replace(formData["time"], "pm", " pm", 1)
	shuttleTime = strings.Replace(shuttleTime, "am", " am", 1)

	resDate, _ := time.Parse(ISODateFormat, formData["date"])
	newDate := resDate.Format(AmericanDateFormat)

	var serviceDirection, origin string
	switch formData["direction"] {
	case "uptown":
		serviceDirection = "Beren to Wilf Campus"
		origin = "Beren"
	case "downtown":
		serviceDirection = "Wilf to Beren Campus"
		origin = "Wilf"
	}

	templateData := TemplateStruct{
		Name:             formData["name"],
		ServiceDirection: serviceDirection,
		Date:             newDate,
		Time:             shuttleTime,
		Origin:           origin,
	}

	err = template.Execute(&messageBody, templateData)
	if err != nil {
		log.Println(err)
	}

	formattedTextBody := FormatTextBody(templateData)

	htmlMessageBody := messageBody.String()

	return htmlMessageBody, formattedTextBody
}

func FormatTextBody(templateData TemplateStruct) string {
	textBody := UnformattedTextBody

	textBody = fmt.Sprintf(textBody, templateData.Name, templateData.ServiceDirection, templateData.ServiceDirection, templateData.Date, templateData.Time, templateData.Origin)

	return textBody
}

func sortTimeList(timeList []string) {
	sort.Slice(timeList, func(i, j int) bool {
		time1, _ := time.Parse(PreparseTimeFormat, timeList[i])
		time2, _ := time.Parse(PreparseTimeFormat, timeList[j])
		return time1.Before(time2)
	})
}

func getListItemTime(listItem *goquery.Selection) string {
	time := listItem.Find(".bup-timeslot-time")
	timeStr := strings.TrimSpace(time.Text())

	return timeStr
}
