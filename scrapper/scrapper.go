package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

//Scrape Indeed by a term
func Scrape(term string) {
	baseURL := "https://de.indeed.com/jobs?q=" + term + "&limit=50&start="
	startTime := time.Now()
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	fmt.Println("total pages are", totalPages)
	for i := 0; i < totalPages; i++ {
		go getPage(baseURL, i, c)

	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	endTime := time.Since(startTime)
	fmt.Println("done in", endTime)
}

//request jobcards from page then extract jobcard to several information
func getPage(baseURL string, page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + strconv.Itoa(page*50)
	fmt.Println("Requesting ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem.fs-unmask")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

//extract informations from jobcard and save to our struct extractedJob
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".jobTitle").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	summary := CleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

//get numbers of pages we should scrap
func getPages(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	//close the body after it's done
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

//write csv file
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	c := make(chan []string)
	errC := make(chan error)
	w := csv.NewWriter(file)

	//put w in file
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)
	for _, job := range jobs {
		go writeJob(job, c)
	}
	for i := 0; i < len(jobs); i++ {
		go func(c chan []string, errC chan<- error) {
			err := w.Write(<-c)
			errC <- err
		}(c, errC)
		// jwErr := w.Write(<-c)
		err := <-errC
		go checkErr(err)
	}

}

func writeJob(job extractedJob, c chan<- []string) {
	c <- []string{"https://de.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.summary}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with status code ", res.StatusCode)
	}
}
