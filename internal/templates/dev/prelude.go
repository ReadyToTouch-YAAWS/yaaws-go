package dev

import "github.com/readytotouch/readytotouch/internal/domain"

type Company = domain.CompanyProfile

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
