package pkg

import (
	"Rusprofile/proto"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "https://www.rusprofile.ru/search?query=%v"

func FindKPP(doc string) string {
	result, err := FindElementByRegex(doc, "clip_kpp\">\\d{9}")
	if err != nil {
		return ""
	}
	return result[10:]
}

func FindDirector(doc string) string {
	result, err := FindElementByRegex(doc, "person_ul\"><span>(.*)<")
	if err != nil {
		return "Company was liquidate"
	}
	result = result[17:]
	result = strings.Replace(result, "<span>", "", -1)
	result = strings.Replace(result, "</span>", "", -1)
	result = strings.Replace(result, "</a>", "", -1)
	result = strings.Replace(result, "<", "", -1)
	return result
}

func FindCompanyName(doc string) string {
	result, err := FindElementByRegex(doc, "legalName.?>(.*)<")
	if err != nil {
		return ""
	}
	result = result[11:]
	result = strings.Replace(result, "&quot;", "", -1)
	result = strings.Replace(result, "<", "", -1)
	return result
}

func FindCompanyByINN(INN string) (*proto.Response, error) {
	response, err := http.Get(fmt.Sprintf(URL, INN))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf(response.Status)
	}

	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bodyToParse := string(body)

	kpp := FindKPP(bodyToParse)
	companyName := FindCompanyName(bodyToParse)
	director := FindDirector(bodyToParse)

	return &proto.Response{
			INN:         INN,
			KPP:         kpp,
			CompanyName: companyName,
			Director:    director,
		},
		nil
}
