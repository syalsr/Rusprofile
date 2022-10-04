package pkg

import (
	"Rusprofile/proto"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"sync"
)

const URL = "https://www.rusprofile.ru/search?query=%v"

func FindKPP(doc *goquery.Document, resp *proto.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	doc.Find(".copy_target").Each(
		func(i int, selection *goquery.Selection) {
			switch i {
			case 2:
				resp.KPP = selection.Text()
			}
		},
	)
}

func FindDirector(doc *goquery.Document, resp *proto.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	doc.Find(".gtm_main_fl").Each(
		func(i int, selection *goquery.Selection) {
			resp.Director = selection.Text()
		},
	)
}

func FindCompanyName(doc *goquery.Document, resp *proto.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	doc.Find(".company-name").Each(
		func(i int, selection *goquery.Selection) {
			resp.CompanyName = selection.Text()
		},
	)
}

func getHtmlPage(INN string) (*http.Response, error) {
	response, err := http.Get(fmt.Sprintf(URL, INN))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf(response.Status)
	}
	return response, nil
}

func FindCompanyByINN(INN string) (*proto.Response, error) {
	resp, err := getHtmlPage(INN)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &proto.Response{INN: INN}
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go FindKPP(doc, response, wg)
	go FindCompanyName(doc, response, wg)
	go FindDirector(doc, response, wg)

	wg.Wait()
	return response, nil
}
