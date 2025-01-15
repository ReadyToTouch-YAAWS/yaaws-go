package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart4() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "limango GmbH — A member of the otto group",
			Website: "https://limango.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9381585,
				Alias:             "limango-gmbh",
				Name:              "limango GmbH - A member of the otto group",
				Followers:         "4K",
				Employees:         "201-500",
				AssociatedMembers: "195",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096830999/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Varonis",
			Website: "https://www.varonis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25403,
				Alias:             "varonis",
				Name:              "Varonis",
				Followers:         "111K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,433",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Python, GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093559031/",
							Date:             mustDate("2024-12-10"),
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "At least 5 years experience as a Software Engineer in Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081442058/",
							Date:                 mustDate("2025-01-03"),
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
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WorldTech IT, LLC",
			Website: "https://wtit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1811897,
				Alias:             "worldtechit",
				Name:              "WorldTech IT, LLC",
				Followers:         "18K",
				Employees:         "11-50",
				AssociatedMembers: "46",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101282630/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Research Solutions",
			Website: "https://www.appliedres.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6017665,
				Alias:             "applied-research-solutions",
				Name:              "Applied Research Solutions",
				Followers:         "6K",
				Employees:         "501-1K",
				AssociatedMembers: "572",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101198898/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WEX",
			Website: "https://www.wexinc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11637,
				Alias:             "wexinc",
				Name:              "WEX",
				Followers:         "92K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,872",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Specialist Java/Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048143482/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Comcast",
			Website: "https://corporate.comcast.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1703,
				Alias:             "comcast",
				Name:              "Comcast",
				Followers:         "648K",
				Employees:         "10K+",
				AssociatedMembers: "59,631",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 48,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4081171024/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Semrush",
			Website: "https://www.semrush.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2821922,
				Alias:             "semrush",
				Name:              "Semrush",
				Followers:         "502K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,091",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "Maroon Team",
							URL:              "https://www.linkedin.com/jobs/view/4063404057/",
							Date:             mustDate("2024-12-13"),
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Go as a primary programming language",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077464276/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pi Square Technologies",
			Website: "https://www.pisquaretech.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10589520,
				Alias:             "pi-square-technology",
				Name:              "Pi Square Technologies",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "226",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4098692466/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DataOn (PT Indodev Niaga Internet)",
			Website: "https://dataon.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                844561,
				Alias:             "dataoncorp",
				Name:              "DataOn (PT Indodev Niaga Internet)",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "468",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4099666536/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiteworks",
			Website: "https://www.kiteworks.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33337,
				Alias:             "kiteworkscgcp",
				Name:              "Kiteworks",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "263",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096373782/",
							Date:             mustDate("2024-12-11"),
						},
						{
							Title:                "Golang Developer / Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123639315/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           false,
							Remote:               true, // Hybrid working model with 2 days office and 3 days remote
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centripetal",
			Website: "https://www.centripetal.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1321511,
				Alias:             "centripetal-ai",
				Name:              "Centripetal",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "102",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Engineer",
							ShortDescription: "CIS Data Services (IE)",
							URL:              "https://www.linkedin.com/jobs/view/4094769985/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Technology Innovation Institute",
			Website: "https://www.tii.ae/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28729816,
				Alias:             "tiiuae",
				Name:              "Technology Innovation Institute",
				Followers:         "162K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,122",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4091567553/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BAUHAUS Deutschland",
			Website: "https://bauhaus.info/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9417424,
				Alias:             "bauhaus-deutschland",
				Name:              "BAUHAUS Deutschland",
				Followers:         "15K",
				Employees:         "10K+",
				AssociatedMembers: "1,633",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049064454/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sun* Vietnam",
			Website: "https://sun-asterisk.vn/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14571208,
				Alias:             "sunasterisk",
				Name:              "Sun* Vietnam",
				Followers:         "11K",
				Employees:         "1K-5K",
				AssociatedMembers: "843",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090103893/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Miracle Software Systems, Inc",
			Website: "https://www.miraclesoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15388,
				Alias:             "miraclesoft",
				Name:              "Miracle Software Systems, Inc",
				Followers:         "115K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,581",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090375762/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Collabera",
			Website: "https://www.collaberadigital.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24440,
				Alias:             "collabera",
				Name:              "Collabera",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "6,048",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Java Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4087446589/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sparq",
			Website: "https://www.teamsparq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                90679425,
				Alias:             "teamsparq",
				Name:              "Sparq",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "790",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049853199/",
							Date:             mustDate("2024-12-03"), // mustDate("2024-12-03")
							WithSalary:       false,
							Remote:           false,
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SMG Swiss Marketplace Group",
			Website: "https://swissmarketplace.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76994292,
				Alias:             "smg-marketplace",
				Name:              "SMG Swiss Marketplace Group",
				Followers:         "23K",
				Employees:         "501-1K",
				AssociatedMembers: "620",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4016921964/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hornetsecurity",
			Website: "https://www.hornetsecurity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                750470,
				Alias:             "hornetsecurity",
				Name:              "Hornetsecurity",
				Followers:         "18K",
				Employees:         "501-1K",
				AssociatedMembers: "398",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "Security Lab",
							URL:              "https://www.linkedin.com/jobs/view/4089330917/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gorilla Logic",
			Website: "https://www.gorillalogic.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68692,
				Alias:             "gorillalogic",
				Name:              "Gorilla Logic",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "524",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085414974/",
							Date:             mustDate("2024-11-26"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ZeptoLab",
			Website: "https://www.zeptolab.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1945804,
				Alias:             "zeptolab",
				Name:              "ZeptoLab",
				Followers:         "32K",
				Employees:         "51-200",
				AssociatedMembers: "143",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102533913/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Getir",
			Website: "https://getir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9866841,
				Alias:             "getir",
				Name:              "Getir",
				Followers:         "415K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,610",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050960562/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Happening",
			Website: "https://www.happening.xyz/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86471246,
				Alias:             "happen-ing",
				Name:              "Happening",
				Followers:         "7K",
				Employees:         "501-1K",
				AssociatedMembers: "514",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102013036/",
							Date:             mustDate("2024-12-16"),
						},
					},
				},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3884244005/",
							Date:             mustDate("2024-11-29"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zeller",
			Website: "https://www.myzeller.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12687202,
				Alias:             "zeller",
				Name:              "Zeller",
				Followers:         "16K",
				Employees:         "201-500",
				AssociatedMembers: "242",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer (GO/Typescript + AWS)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047579348/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flip",
			Website: "https://flip.id/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13234824,
				Alias:             "flip.id",
				Name:              "Flip",
				Followers:         "222K",
				Employees:         "201-500",
				AssociatedMembers: "610",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Backend-GO) — Risk Platform",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096854844/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenX",
			Website: "https://www.openx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42890,
				Alias:             "openx",
				Name:              "OpenX",
				Followers:         "30K",
				Employees:         "201-500",
				AssociatedMembers: "440",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer IV (Java/GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062276145/",
							Date:             mustDate("2024-12-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lenskart.com",
			Website: "https://lenskart.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                213339,
				Alias:             "lenskart-com",
				Name:              "Lenskart.com",
				Followers:         "397K",
				Employees:         "10K+",
				AssociatedMembers: "10,127",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096376928/",
							Date:             mustDate("2024-12-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lowe's Companies, Inc.",
			Website: "https://lowes.com/",
			LinkedInProfile: domain.LinkedInProfile{
				// LinkedIn alias that is hard to encode
				ID:                4128,
				Alias:             "4128", // "lowe's-home-improvement",
				Name:              "Lowe's Companies, Inc.",
				Followers:         "756K",
				Employees:         "10K+",
				AssociatedMembers: "118,402",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Mobile/Backend (Android, Go, Kotlin)",
							URL:              "https://www.linkedin.com/jobs/view/4093511115/",
							Date:             mustDate("2025-01-14"), // mustDate("2024-12-11"),
							WithSalary:       true,                   // $92,300.00 - $175,400.00
							Remote:           true,                   // On-site
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Huntress",
			Website: "https://www.huntress.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10172550,
				Alias:             "huntress-labs",
				Name:              "Huntress",
				Followers:         "71K",
				Employees:         "201-500",
				AssociatedMembers: "542",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Windows SIEM Agent (Go)",
							URL:              "https://www.linkedin.com/jobs/view/4055142455/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1NCE",
			Website: "https://1nce.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18471615,
				Alias:             "1nce",
				Name:              "1NCE",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "289",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088250312/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Waystar",
			Website: "https://www.waystar.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11444741,
				Alias:             "waystar",
				Name:              "Waystar",
				Followers:         "21K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,457",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Advanced Application Engineer (PHP / GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4089563703/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "efood",
			Website: "https://www.e-food.gr/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2410959,
				Alias:             "e-food-gr",
				Name:              "efood",
				Followers:         "78K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,190",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3873641135/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ESET",
			Website: "https://www.eset.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28967,
				Alias:             "eset",
				Name:              "ESET",
				Followers:         "77K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,838",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Cloud Security Applications (.NET Core or GO)",
							URL:              "https://www.linkedin.com/jobs/view/4084785751/",
							Date:             mustDate("2024-12-03"),
						},
						{
							Title:                "Software Engineer — Cloud Security Applications (.NET Core or Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122199318/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           true, // from 2 600 EUR
							Remote:               true, // Hybrid
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Praxent",
			Website: "https://praxent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1365857,
				Alias:             "praxent",
				Name:              "Praxent",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "116",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "LATAM Senior Backend Engineer (Python or GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085449585/",
							Date:             mustDate("2024-11-27"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stacklok",
			Website: "https://stacklok.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                94101540,
				Alias:             "stacklok",
				Name:              "Stacklok",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "39",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer II (Go & Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079391812/",
							Date:             mustDate("2024-11-20"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PT Lion Super Indo",
			Website: "https://www.superindo.co.id/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1104467,
				Alias:             "superindo",
				Name:              "PT Lion Super Indo",
				Followers:         "630K",
				Employees:         "10K+",
				AssociatedMembers: "13,961",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102186636/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ubiquiti Inc.",
			Website: "https://www.ui.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                109341,
				Alias:             "ubiquiti-",
				Name:              "Ubiquiti Inc.",
				Followers:         "103K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,320",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4100868953/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TextNow",
			Website: "https://textnow.com/",
			Careers: "https://careers.textnow.com/careers",
			About:   "https://careers.textnow.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                529693,
				Alias:             "enflick-inc-",
				Name:              "TextNow",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "249",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "Trust & Safety",
							URL:              "https://www.linkedin.com/jobs/view/4096314146/",
							Date:             mustDate("2024-12-11"),
						},
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "Maintain and refactor backend systems to improve our existing detections in languages like Golang, PHP, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4096309405/",
							Date:                 mustDate("2025-01-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Software Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124814543/",
							Date:                 mustDate("2025-01-15"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Altenar",
			Website: "https://altenar.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11063103,
				Alias:             "altenar",
				Name:              "Altenar",
				Followers:         "16K",
				Employees:         "501-1K",
				AssociatedMembers: "299",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4006862895/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VUS — The English Center",
			Website: "https://vus.edu.vn/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33233160,
				Alias:             "vustheenglishcenter",
				Name:              "VUS - The English Center",
				Followers:         "13K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,987",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Developer (Golang)",
							ShortDescription: "3 years exp",
							URL:              "https://www.linkedin.com/jobs/view/4089423781/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exness",
			Website: "https://www.exness.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3291356,
				Alias:             "exness",
				Name:              "Exness",
				Followers:         "92K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,664",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang/Python Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088177442/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Xsolla",
			Website: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                875431,
				Alias:             "xsolla",
				Name:              "Xsolla",
				Followers:         "32K",
				Employees:         "501-1K",
				AssociatedMembers: "944",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4069628874/",
							Date:             mustDate("2025-01-15"), // mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "emerchantpay",
			Website: "https://www.emerchantpay.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1432812,
				Alias:             "emerchantpay-ltd",
				Name:              "emerchantpay",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "408",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Developer with Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4084138762/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Power Factors",
			Website: "https://www.powerfactors.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4748453,
				Alias:             "power-factors-inc",
				Name:              "Power Factors",
				Followers:         "15K",
				Employees:         "501-1K",
				AssociatedMembers: "586",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer, Back-End, Golang (Greece)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102953810/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Airspace",
			Website: "https://www.airspace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10504949,
				Alias:             "airspace-inc",
				Name:              "Airspace",
				Followers:         "9K",
				Employees:         "201-500",
				AssociatedMembers: "360",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101213697/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amartha",
			Website: "https://amartha.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3655032,
				Alias:             "amartha",
				Name:              "Amartha",
				Followers:         "123K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,197",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Golang) — Senior/Principal Level",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102510260/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Creative Software",
			Website: "http://www.creativesoftware.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12688,
				Alias:             "creativesoftware",
				Name:              "Creative Software",
				Followers:         "61K",
				Employees:         "201-500",
				AssociatedMembers: "531",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer(Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101950317/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cisco ThousandEyes",
			Website: "https://www.thousandeyes.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                687352,
				Alias:             "thousandeyes",
				Name:              "Cisco ThousandEyes",
				Followers:         "62K",
				Employees:         "1K-5K",
				AssociatedMembers: "959",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Site Reliability Engineer I (Python/Golang), Agent Ops",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4100798671/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "U.S. Bank",
			Website: "https://usbank.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2532,
				Alias:             "us-bank",
				Name:              "U.S. Bank",
				Followers:         "499K",
				Employees:         "10K+",
				AssociatedMembers: "75,830",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Information Security Engineer (Golang or Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063303097/",
							Date:             mustDate("2024-12-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NetApp",
			Website: "https://www.netapp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2105,
				Alias:             "netapp",
				Name:              "NetApp",
				Followers:         "709K",
				Employees:         "10K+",
				AssociatedMembers: "12,935",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4095534140/",
							Date:             mustDate("2024-12-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trend Micro",
			Website: "https://www.trendmicro.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4312,
				Alias:             "trend-micro",
				Name:              "Trend Micro",
				Followers:         "269K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,565",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Cloud Engineer (Golang/Python, Backend Focus)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063567336/",
							Date:             mustDate("2024-12-11"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Central Retail",
			Website: "https://www.centralretail.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26628492,
				Alias:             "centralretail",
				Name:              "Central Retail",
				Followers:         "67K",
				Employees:         "10K+",
				AssociatedMembers: "1,066",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093163873/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kargo",
			Website: "https://www.kargo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                55830,
				Alias:             "kargo",
				Name:              "Kargo",
				Followers:         "37K",
				Employees:         "201-500",
				AssociatedMembers: "980",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4089015041/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Salesloft",
			Website: "https://www.salesloft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2296178,
				Alias:             "salesloft",
				Name:              "Salesloft",
				Followers:         "110K",
				Employees:         "501-1K",
				AssociatedMembers: "1,253",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072522353/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CJ MORE",
			Website: "https://www.cjexpress.co.th/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10081692,
				Alias:             "cj-express-group",
				Name:              "CJ MORE",
				Followers:         "35K",
				Employees:         "501-1K",
				AssociatedMembers: "669",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088241281/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "H-E-B",
			Website: "https://www.heb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164159,
				Alias:             "heb",
				Name:              "H-E-B",
				Followers:         "273K",
				Employees:         "10K+",
				AssociatedMembers: "42,800",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Sr Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4000104862/",
							Date:             mustDate("2024-12-03"),
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
		//},
	}
}
