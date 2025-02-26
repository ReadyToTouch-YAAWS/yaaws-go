package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart7() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Checkmarx",
			Website: "https://checkmarx.com/",
			Careers: "https://checkmarx.com/company/careers/",
			About:   "https://checkmarx.com/company/about-checkmarx/",
			Blog:    "https://checkmarx.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                223472,
				Alias:             "checkmarx",
				Name:              "Checkmarx",
				Followers:         "109K",
				Employees:         "501-1K",
				AssociatedMembers: "902",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "checkmarx",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "checkmarx",
				Employees: "820",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Checkmarx-EI_IE554001.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Checkmarx-Reviews-E554001.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Checkmarx-Jobs-E554001.htm",
				Jobs:        "26",
				Reviews:     "353",
				Salaries:    "442",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "With Golang as the primary technology, you will cross skill yourself in other programming languages like Java and TypeScript/JavaScript and technologies required in the project",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114108035/",
							Date:                 mustDate("2025-01-02"),
							WithSalary:           false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Visa",
			Website: "https://visa.com/",
			Careers: "https://corporate.visa.com/en/careers.html",
			About:   "https://corporate.visa.com/en/about-visa.html",
			Blog:    "https://community.developer.visa.com/t5/Blogs/bg-p/Developer_Blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2190,
				Alias:             "visa",
				Name:              "Visa",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "28,261",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "visa",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "visa",
				Employees:   "10,000+",
				Salary:      "$45K ~ $253K a year",
				Reviews:     "1,217",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "visa",
				Employees: "21,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Visa-Inc-EI_IE3035.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Visa-Inc-Reviews-E3035.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Visa-Inc-Jobs-E3035.htm",
				Jobs:        "872",
				Reviews:     "7.6K",
				Salaries:    "19K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Systems Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4056869486/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (C++ / Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121157867/",
							Date:                 mustDate("2025-01-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (C++ / Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151356731/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Engineer (C++/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168438476/",
							Date:                 mustDate("2025-02-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4167397544/",
							Date:                 mustDate("2025-02-25"),
							WithSalary:           true, // $105.800 - $149.600 per year
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4103097753/",
							Date:                 mustDate("2025-02-09"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Datazip",
			Website: "https://datazip.io/",
			Careers: "",
			About:   "https://datazip.io/about-us",
			Blog:    "https://datazip.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80069364,
				Alias:             "datazipio",
				Name:              "Datazip",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "datazip-inc",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Founding Backend Developer (Golang)",
							ShortDescription:     "Write clean, efficient, and well-documented code in Golang (must) and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114492171/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Composable Lakehouse Platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Baker Hughes",
			Website: "https://www.bakerhughes.com/",
			Careers: "https://careers.bakerhughes.com/",
			About:   "https://www.bakerhughes.com/company/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4334,
				Alias:             "bakerhughes",
				Name:              "Baker Hughes",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "64,211",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "BakerHughes",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Baker-Hughes",
				Employees:   "10,000+",
				Salary:      "$75K ~ $224K a year",
				Reviews:     "30",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "baker-hughes",
				Employees: "57,510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Baker-Hughes-EI_IE1902699.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Baker-Hughes-Reviews-E1902699.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Baker-Hughes-Jobs-E1902699.htm",
				Jobs:        "940",
				Reviews:     "6.3K",
				Salaries:    "9.7K",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4106957208/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Baker Hughes is an energy technology company that provides solutions for energy and industrial customers worldwide",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/BKR/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/BKR:NASDAQ",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centrica",
			Website: "https://www.centrica.com/",
			Careers: "https://www.lifeatcentrica.com/jobs/",
			About:   "https://www.centrica.com/who-we-are/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2766,
				Alias:             "centrica",
				Name:              "Centrica",
				Followers:         "143K",
				Employees:         "10K+",
				AssociatedMembers: "19,604",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Centrica",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "4.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "centrica",
				Employees: "20,406",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Centrica-EI_IE6731.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Centrica-Reviews-E6731.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Centrica-Jobs-E6731.htm",
				Jobs:        "108",
				Reviews:     "1.1K",
				Salaries:    "2.4K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4076005752/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           true, // Enjoy a generous market salary of up to £60,000
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132872793/",
							Date:                 mustDate("2025-01-24"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Centrica is a leading energy services and solutions provider founded on a 200-year heritage of serving people",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mural Health",
			Website: "https://www.muralhealth.com/",
			Careers: "https://www.muralhealth.com/careers",
			About:   "https://www.muralhealth.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89290648,
				Alias:             "muralhealth",
				Name:              "Mural Health",
				Followers:         "23K",
				Employees:         "11-50",
				AssociatedMembers: "36",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MuralHealth",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "Assist in designing, developing, and maintaining efficient and scalable backend systems and APIs using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115784363/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Mural Health is a clinical trials technology company that is dedicated to making it easy to be a participant in clinical trials",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Green-Got",
			Website: "https://green-got.com/",
			Careers: "https://greengotteam.notion.site/On-recrute-dc1e9c1805e54db1ac1b2390a4ae00a8",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42124932,
				Alias:             "greengot",
				Name:              "Green-Got",
				Followers:         "104K",
				Employees:         "11-50",
				AssociatedMembers: "481",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "green-got",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "The stack is mostly Rust and PostgreSQL on AWS",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115290108/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           true, // €50K/yr - €65K/yr
							Remote:               true,
						},
						{
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120580858/",
							Date:                 mustDate("2025-01-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "In payments",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4141808768/",
							Date:                 mustDate("2025-02-05"),
							WithSalary:           true, // €50.000 - €65.000 per year
							Remote:               true,
						},
						{
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "In payments",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161923630/",
							Date:                 mustDate("2025-02-24"),
							WithSalary:           true, // €50.000 - €65.000 per year
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Template short
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:                0,
		//		Alias:             "",
		//		Name:              "",
		//		Followers:         "",
		//		Employees:         "",
		//		AssociatedMembers: "",
		//		Verified:          false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:    "",
		//		Verified: false,
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies: []domain.Vacancy{
		//				{
		//					Title:                "",
		//					ShortDescription:     "",
		//					SwitchingOpportunity: "",
		//					URL:                  "",
		//					Date:                 mustDate(""),
		//					WithSalary:           false,
		//					Remote:               false,
		//				},
		//			},
		//		},
		//		domain.Rust:    {},
		//		domain.Zig:     {},
		//		domain.Scala:   {},
		//		domain.Elixir:  {},
		//		domain.Clojure: {},
		//		domain.Haskell: {},
		//	},
		//	ShortDescription: "",
		//	Industries:       []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{
		//		domain.Ukraine,
		//		domain.Czechia,
		//	},
		//},
	}
}
