package dev

import (
	"fmt"
	"net/url"
)

type CompanyCodePair struct {
	ID    int64
	Name  string
	Alias string
}

type VacancyCodePair struct {
	ID           int64
	URL          string
	CompanyAlias string
}

func googleSearchLogos(companyName string) string {
	values := url.Values{
		"q": {fmt.Sprintf("%q review-logos", companyName)},
	}

	return "https://www.google.com/search?" + values.Encode()
}
