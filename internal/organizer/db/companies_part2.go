package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart2() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trusting Social",
			Website: "https://www.trustingsocial.com/",
			Careers: "https://trustingsocial.com/careers/",
			About:   "https://trustingsocial.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3506386,
				Alias:             "trustingsocial",
				Name:              "Trusting Social",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "321",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrustingSocial-EI_IE741535.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrustingSocial-Reviews-E741535.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Back End (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4042510486/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4118641348/",
							Date:             mustDate("2025-01-09"),
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
			ShortDescription: "AI FinTech company revolutionizing credit scoring using big data technology",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "McAfee",
			Website: "https://www.mcafee.com/",
			Careers: "https://careers.mcafee.com/global/",
			About:   "https://www.mcafee.com/en-us/consumer-corporate/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2336,
				Alias:             "mcafee",
				Name:              "McAfee",
				Followers:         "320K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,531",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mcafee",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mcafee",
				Employees: "7,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McAfee-EI_IE2244.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McAfee-Reviews-E2244.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            ".Net / Golang with Postgres Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4041686516/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Antivirus software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hadrian",
			Website: "https://www.hadrian.co/",
			Careers: "https://www.hadrian.co/careers",
			About:   "https://www.hadrian.co/mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71668100,
				Alias:             "hadrianspace",
				Name:              "Hadrian",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "182",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hadrian-EI_IE5364733.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hadrian-Reviews-E5364733.htm",
			},
			OttaProfileSlug:   "Hadrian-2",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/63inUEkF",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Building autonomous factories for aerospace components",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scaleway",
			Website: "https://www.scaleway.com/",
			Careers: "https://www.scaleway.com/en/careers/",
			About:   "https://www.scaleway.com/en/about-us/",
			Blog:    "https://www.scaleway.com/en/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9427185,
				Alias:             "scaleway",
				Name:              "Scaleway",
				Followers:         "33K",
				Employees:         "501-1K",
				AssociatedMembers: "561",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "scaleway",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scaleway-EI_IE1861876.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scaleway-Reviews-E1861876.htm",
			},
			OttaProfileSlug:   "Scaleway",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/dFlIR1ow",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Cloud provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Demandbase",
			Website: "https://www.demandbase.com/",
			Careers: "https://www.demandbase.com/about-us/careers/job-openings/",
			About:   "https://www.demandbase.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89759,
				Alias:             "demandbase",
				Name:              "Demandbase",
				Followers:         "68K",
				Employees:         "501-1K",
				AssociatedMembers: "981",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Demandbase",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "demandbase",
				Employees: "930",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Demandbase-EI_IE271272.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Demandbase-Reviews-E271272.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Demandbase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056224841/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134421375/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           true, // $160.000 — $180.000 per year
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Go-To-Market engagement & marketing software",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Overwolf",
			Website: "https://www.overwolf.com/",
			Careers: "https://www.overwolf.com/careers/",
			About:   "https://www.overwolf.com/about-overwolf/",
			Blog:    "https://blog.overwolf.com/tag/developers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                519584,
				Alias:             "overwolf.com",
				Name:              "Overwolf",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "overwolf",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Overwolf-EI_IE1963582.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Overwolf-Reviews-E1963582.htm",
			},
			OttaProfileSlug:   "Overwolf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/duQ9FWwK",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Gaming app & mod development framework",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitfount",
			Website: "https://www.bitfount.com/",
			Careers: "https://www.bitfount.com/company/careers",
			About:   "https://www.bitfount.com/company/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69183486,
				Alias:             "bitfount",
				Name:              "Bitfount",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "25",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bitfount",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Bitfount",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full Stack Developer",
							ShortDescription: "Development experience with Rust or a desire to learn this on the job",
							URL:              "https://www.linkedin.com/jobs/view/4044732400/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Federated AI & distributed data science platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hinge",
			Website: "https://hinge.co/",
			Careers: "https://hinge.co/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2603651,
				Alias:             "hinge-app",
				Name:              "Hinge",
				Followers:         "46K",
				Employees:         "51-200",
				AssociatedMembers: "431",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Hinge",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hinge",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hinge-EI_IE871901.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hinge-Reviews-E871901.htm",
			},
			OttaProfileSlug:   "Hinge",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer",
							ShortDescription: "Design and build distributed systems in Go",
							URL:              "https://www.linkedin.com/jobs/view/4043430809/",
							Date:             mustDate("2024-12-05"),
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
			ShortDescription:          "Dating app for meaningful relationships",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Firework",
			Website: "https://firework.com/",
			Careers: "https://firework.com/careers/",
			About:   "https://firework.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18700473,
				Alias:             "fireworkhq",
				Name:              "Firework",
				Followers:         "26K",
				Employees:         "201-500",
				AssociatedMembers: "371",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firework-EI_IE2300706.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firework-Reviews-E2300706.htm",
			},
			OttaProfileSlug:   "Firework",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/srNX_4cO",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Video-based eCommerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Breakroom",
			Website: "https://www.breakroom.cc/",
			Careers: "https://www.breakroom.cc/en-gb/careers",
			About:   "https://www.breakroom.cc/en-gb/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18458137,
				Alias:             "breakroom",
				Name:              "Breakroom",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "breakroom",
				Verified: true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Breakroom-EI_IE8002334.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Breakroom-Reviews-E8002334.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Breakroom",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Our platform is built on Elixir, using the Phoenix web framework",
							URL:              "https://www.linkedin.com/jobs/view/4052987616/",
							Date:             mustDate("2024-10-21"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "The people-powered job comparison site",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Ansa",
			Website: "https://www.ansa.dev/",
			Careers: "",
			About:   "https://www.ansa.dev/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91158109,
				Alias:             "getansa",
				Name:              "Ansa",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "27",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Ansa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/EjVpa6gl",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Branded customer wallet infrastructure",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1GLOBAL",
			Website: "https://www.1global.com/",
			Careers: "https://www.1global.com/careers",
			About:   "https://www.1global.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                39711,
				Alias:             "1global",
				Name:              "1GLOBAL",
				Followers:         "61K",
				Employees:         "501-1K",
				AssociatedMembers: "463",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1GLOBAL-EI_IE9805631.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1GLOBAL-Reviews-E9805631.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048972208/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Salesforce Developer — .NET / Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4117970233/",
							Date:             mustDate("2025-01-08"),
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
			ShortDescription: "eSIM",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Eco-Movement",
			Website: "https://eco-movement.com/",
			Careers: "https://www.eco-movement.com/careers-job-vacancy/",
			About:   "https://www.eco-movement.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15223273,
				Alias:             "eco-movement-charge-point-data",
				Name:              "Eco-Movement",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "50",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eco-Movement-EI_IE5353422.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eco-Movement-Reviews-E5353422.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4046511967/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "Charging station data platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Network Optix",
			Website: "https://www.networkoptix.com/",
			Careers: "https://www.networkoptix.com/company/careers",
			About:   "https://www.networkoptix.com/company/about-us",
			Blog:    "https://www.networkoptix.com/blog/tag/developers",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1896041,
				Alias:             "network-optix",
				Name:              "Network Optix",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "183",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "networkoptix",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "network-optix",
				Employees: "157",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Network-Optix-EI_IE1103945.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Network-Optix-Reviews-E1103945.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior C++/Golang Engineer (Cloud Services)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4028418131/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Convert video data into practical insights for businesses and industries of all kinds",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cinemo",
			Website: "https://www.cinemo.com/",
			Careers: "https://www.cinemo.com/about/careers/",
			About:   "https://www.cinemo.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                564661,
				Alias:             "cinemo",
				Name:              "Cinemo",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "262",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cinemo-EI_IE2189711.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cinemo-Reviews-E2189711.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer / Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050001754/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "(Senior) Software Engineer – React & Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3921611198/",
							Date:             mustDate("2025-01-23"), // mustDate("2024-12-21"),
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
			ShortDescription: "Make every screen an opportunity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Readdle",
			Website: "https://readdle.com/",
			Careers: "https://readdle.com/careers",
			About:   "https://readdle.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                629551,
				Alias:             "readdle",
				Name:              "Readdle",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "280",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "readdle",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "readdle",
				Employees: "190",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Readdle-EI_IE613771.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Readdle-Reviews-E613771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Back End Engineer — Spark team",
							ShortDescription: "Build distributed, scalable, and fault-tolerant systems in the cloud using Golang",
							URL:              "https://www.linkedin.com/jobs/view/3971228819/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Creating apps",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ipsos",
			Website: "https://www.ipsos.com/",
			Careers: "https://www.ipsos.com/en/careers",
			About:   "https://www.ipsos.com/en/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4318,
				Alias:             "ipsos",
				Name:              "Ipsos",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "17,733",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ipsos",
				Employees: "14,170",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ipsos-EI_IE13063.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ipsos-Reviews-E13063.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Back-end Developer Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053349624/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Market research company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bedrock Streaming",
			Website: "https://bedrockstreaming.com/",
			Careers: "https://bedrockstreaming.com/career/",
			About:   "https://bedrockstreaming.com/company/",
			Blog:    "https://tech.bedrockstreaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64853943,
				Alias:             "bedrock-streaming",
				Name:              "Bedrock Streaming",
				Followers:         "30K",
				Employees:         "201-500",
				AssociatedMembers: "315",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bedrockstreaming",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bedrock-streaming",
				Employees: "360",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bedrock-Streaming-EI_IE5955934.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bedrock-Streaming-Reviews-E5955934.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer — Backend Go / PHP",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049210489/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Software Developer —  Backend Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079483983/",
							Date:             mustDate("2025-01-31"), // mustDate("2024-12-13"), // mustDate("2024-12-11")
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
			ShortDescription:          "Bedrock creates and operates full-scope streaming platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Whalebone",
			Website: "https://www.whalebone.io/",
			Careers: "https://whalebone.recruitee.com/",
			About:   "https://www.whalebone.io/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10961729,
				Alias:             "whaleboneio",
				Name:              "Whalebone",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "118",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "whalebone",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whalebone-EI_IE4889910.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whalebone-Reviews-E4889910.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4007694502/",
							Date:             mustDate("2025-02-06"), // mustDate("2025-01-08"), // mustDate("2024-12-20"),
						},
						{
							Title:                "Backend Developer — PHP, Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132899360/",
							Date:                 mustDate("2025-02-05"), // mustDate("2025-01-25"),
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OneSignal",
			Website: "https://onesignal.com/",
			Careers: "https://onesignal.com/careers",
			About:   "https://onesignal.com/about",
			Blog:    "https://onesignal.com/blog/tag/development/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6424376,
				Alias:             "onesignal",
				Name:              "OneSignal",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "164",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "OneSignal",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "onesignal",
				Employees: "100",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSignal-EI_IE1952551.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSignal-Reviews-E1952551.htm",
			},
			OttaProfileSlug:   "OneSignal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Analytics Team",
							ShortDescription: "Using languages such as, Rust and Golang",
							URL:              "https://www.linkedin.com/jobs/view/4082610714/",
							Date:             mustDate("2024-12-05"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 32,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Analytics Team",
							ShortDescription: "Using languages such as, Rust and Golang",
							URL:              "https://www.linkedin.com/jobs/view/4082610714/",
							Date:             mustDate("2024-12-05"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Customer messaging and engagement platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Metabase",
			Website: "https://www.metabase.com/",
			Careers: "https://www.metabase.com/jobs",
			About:   "",
			Blog:    "https://www.metabase.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6460313,
				Alias:             "metabase",
				Name:              "Metabase",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "94",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "metabase",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "metabase",
				Employees: "45",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Metabase-CA-EI_IE6095700.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Metabase-CA-Reviews-E6095700.htm",
			},
			OttaProfileSlug:   "Metabase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:     {},
				domain.Rust:   {},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 30,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Engineering Manager",
							ShortDescription: "Strong technical background in one of: Clojure or Frontend (React + Typescript) ecosystems",
							URL:              "https://app.welcometothejungle.com/jobs/K21IaTJ2",
							Date:             mustDate("2024-12-27"), // Approximate date
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Open-source Business Intelligence tool",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Anjuna Security",
			Website: "https://www.anjuna.io/",
			Careers: "https://www.anjuna.io/careers",
			About:   "https://www.anjuna.io/our-team",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18438300,
				Alias:             "anjuna-security",
				Name:              "Anjuna Security",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "anjuna-security",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "anjuna",
				Employees: "70",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anjuna-Security-EI_IE3150191.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anjuna-Security-Reviews-E3150191.htm",
			},
			OttaProfileSlug:   "Anjuna",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Experience with Go, Python, and Rust",
							URL:              "https://www.linkedin.com/jobs/view/4052028718/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Experience with Go, Python, and Rust",
							URL:              "https://www.linkedin.com/jobs/view/4052028718/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Public-cloud network security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aqua Security",
			Website: "https://www.aquasec.com/",
			Careers: "https://www.aquasec.com/about-us/careers/",
			About:   "https://www.aquasec.com/about-us/",
			Blog:    "https://www.aquasec.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10034420,
				Alias:             "aquasecteam",
				Name:              "Aqua Security",
				Followers:         "63K",
				Employees:         "501-1K",
				AssociatedMembers: "639",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "aquasecurity",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "aqua-security",
				Employees: "751",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aqua-Security-Software-EI_IE1785939.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aqua-Security-Software-Reviews-E1785939.htm",
			},
			OttaProfileSlug:   "Aqua-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 100,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Expert Software Engineer (Golang Developer)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4045761623/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Cloud native security platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Verve",
			Website: "https://verve.com/",
			Careers: "https://verve.com/careers/",
			About:   "https://verve.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                108605,
				Alias:             "verve-ad-solutions",
				Name:              "Verve",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "470",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verve-EI_IE4432646.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verve-Reviews-E4432646.htm",
			},
			OttaProfileSlug:   "Verve",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Backend Engineer — Java/Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4015809270/",
							Date:             mustDate("2024-09-25"),
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
			ShortDescription: "Real time omnichannel ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wargaming",
			Website: "https://wargaming.com/",
			Careers: "https://wargaming.com/careers/",
			About:   "https://wargaming.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                127309,
				Alias:             "wargaming-net",
				Name:              "Wargaming",
				Followers:         "131K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,921",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "wgnet",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "wargaming",
				Employees: "4,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wargaming-EI_IE381713.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wargaming-Reviews-E381713.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Python/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077489876/",
							Date:                 mustDate("2024-11-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Python/Go)",
							ShortDescription:     "World of Warships, PC",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4147156645/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Software Engineer (Platform)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077490300/",
							Date:             mustDate("2024-11-17"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Game developer and publisher",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Napier AI",
			Website: "https://www.napier.ai/",
			Careers: "https://napier.pinpointhq.com/",
			About:   "https://www.napier.ai/post/napier-mastercard-framl",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15197985,
				Alias:             "napier",
				Name:              "Napier AI",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "229",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "napier-ai",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "napier-technologies",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Napier-Technologies-EI_IE4370088.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Napier-Technologies-Reviews-E4370088.htm",
			},
			OttaProfileSlug:   "Napier",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Back-end Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4044948110/",
							Date:             mustDate("2024-12-18"),
						},
						{
							Title:            "Back-end Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106841758/",
							Date:             mustDate("2024-12-24"),
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
			ShortDescription: "AI-driven AML solutions transform financial crime compliance from legal obligation to competitive edge",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kaseya",
			Website: "https://www.kaseya.com/",
			Careers: "https://www.kaseya.com/careers/jobs/",
			About:   "https://www.kaseya.com/company/",
			Blog:    "https://www.kaseya.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                21377,
				Alias:             "kaseya",
				Name:              "Kaseya",
				Followers:         "157K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,836",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "kaseya",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kaseya",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kaseya-EI_IE262966.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kaseya-Reviews-E262966.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3908613207/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Provider of AI-powered cybersecurity and IT management software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Omio",
			Website: "https://www.omio.com/",
			Careers: "https://www.omio.com/corporate/jobs/",
			About:   "https://www.omio.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2754440,
				Alias:             "omio",
				Name:              "Omio",
				Followers:         "50K",
				Employees:         "201-500",
				AssociatedMembers: "361",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "omio-labs",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "omio",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omio-EI_IE2855962.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omio-Reviews-E2855962.htm",
			},
			OttaProfileSlug:   "Omio",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4025460474/",
							Date:             mustDate("2024-09-25"),
						},
						{
							Title:                "Software Engineer — Go/Node/Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124743101/",
							Date:                 mustDate("2025-01-24"),
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
			ShortDescription: "Comparison & booking website for public transport",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Liebherr",
			Website: "https://www.liebherr.com/",
			Careers: "https://www.liebherr.com/en-int/careers/careers-5370600",
			About:   "https://www.liebherr.com/en-int/group/start-page-3935069",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11057,
				Alias:             "liebherr",
				Name:              "Liebherr Group",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "10,005",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "liebherr-group",
				Employees: "49,611",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Liebherr-Group-EI_IE787369.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Liebherr-Group-Reviews-E787369.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Cloud Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3821037210/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Construction equipment manufacturer",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fiserv",
			Website: "https://www.fiserv.com/",
			Careers: "https://www.careers.fiserv.com/",
			About:   "https://www.fiserv.com/en/about-fiserv.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3364,
				Alias:             "fiserv",
				Name:              "Fiserv",
				Followers:         "755K",
				Employees:         "10K+",
				AssociatedMembers: "35,296",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fiserv",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fiserv",
				Employees: "31,290",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fiserv-EI_IE1384.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fiserv-Reviews-E1384.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4026685486/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4131803257/",
							Date:             mustDate("2025-01-22"),
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
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gett",
			Website: "https://www.gett.com/",
			Careers: "https://www.gett.com/careers/",
			About:   "https://www.gett.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1514929,
				Alias:             "gettaxi",
				Name:              "Gett",
				Followers:         "49K",
				Employees:         "501-1K",
				AssociatedMembers: "862",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "GettEngineering",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gett",
				Employees: "930",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gett-EI_IE810731.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gett-Reviews-E810731.htm",
			},
			OttaProfileSlug:   "Gett",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:            "https://www.linkedin.com/jobs/view/3819936157/",
							ShortDescription: "We are looking for a talented Go developer",
							URL:              "https://www.linkedin.com/jobs/view/3819936157/",
							Date:             mustDate("2024-02-26"),
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
			ShortDescription: "SaaS platform for mobility providers",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stairwell",
			Website: "https://stairwell.com/",
			Careers: "https://stairwell.com/about/careers/",
			About:   "https://stairwell.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68606343,
				Alias:             "stairwell-inc",
				Name:              "Stairwell",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "65",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stairwell-inc",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "stairwell",
				Employees: "58",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Stairwell",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Strong background in software development in one or more general purpose programming languages including but not limited to: Go, Java, or C++",
							URL:              "https://www.linkedin.com/jobs/view/3949883320/",
							Date:             mustDate("2024-06-25"),
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
			ShortDescription: "Threat detection & cybersecurity platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Solo.io",
			Website: "https://www.solo.io/",
			Careers: "https://www.solo.io/company/careers",
			About:   "https://www.solo.io/company/about-us",
			Blog:    "https://www.solo.io/blog?topic=Technical",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11074869,
				Alias:             "solo.io",
				Name:              "Solo.io",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "248",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "solo-io",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-solo-io-EI_IE5382785.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/solo-io-Reviews-E5382785.htm",
			},
			OttaProfileSlug:   "Solo-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 87,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/vuvd0QAu",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Delivering API infrastructure software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Splash Damage",
			Website: "https://www.splashdamage.com/",
			Careers: "https://careers.splashdamage.com/",
			About:   "https://careers.splashdamage.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47756,
				Alias:             "splash-damage",
				Name:              "Splash Damage",
				Followers:         "47K",
				Employees:         "201-500",
				AssociatedMembers: "329",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "splash-damage",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splash-Damage-EI_IE470252.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splash-Damage-Reviews-E470252.htm",
			},
			OttaProfileSlug:   "SplashDamage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/A7daqq7I",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Game development studio",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Ditto",
			Website: "https://ditto.live/",
			Careers: "https://job-boards.greenhouse.io/dittoliveincorporated",
			About:   "https://ditto.live/about",
			Blog:    "https://ditto.live/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18702497,
				Alias:             "dittolive",
				Name:              "Ditto",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "143",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "getditto",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dittolive",
				Employees: "60",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "ditto",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 30,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "1+ year working with Rust",
							URL:              "https://www.linkedin.com/jobs/view/4050904655/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Data sync platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pinecone",
			Website: "https://www.pinecone.io/",
			Careers: "https://www.pinecone.io/careers/",
			About:   "https://www.pinecone.io/company/",
			Blog:    "https://www.pinecone.io/blog/?category=Engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20299330,
				Alias:             "pinecone-io",
				Name:              "Pinecone",
				Followers:         "63K",
				Employees:         "51-200",
				AssociatedMembers: "178",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pinecone-io",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pinecone",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinecone-Systems-EI_IE6085804.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinecone-Systems-Reviews-E6085804.htm",
			},
			OttaProfileSlug:   "Pinecone",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Rust: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Platform Engineer",
							ShortDescription: "We focus on building robust Rust and Python applications to drive our core functionalities",
							URL:              "https://www.linkedin.com/jobs/view/4002028518/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Vector database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Beyond Identity",
			Website: "https://www.beyondidentity.com/",
			Careers: "https://www.beyondidentity.com/careers",
			About:   "https://www.beyondidentity.com/company/about-us",
			Blog:    "https://www.beyondidentity.com/collections/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64665737,
				Alias:             "beyond-identity-inc",
				Name:              "Beyond Identity",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "140",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gobeyondidentity",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "beyond-identity",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Beyond-Identity-EI_IE3403008.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Beyond-Identity-Reviews-E3403008.htm",
			},
			OttaProfileSlug:   "Beyond-Identity",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "We develop most of our server-side code in Rust and it's a plus if you already love Rust, or want a chance to become a Rust expert",
							URL:              "https://www.linkedin.com/jobs/view/4068292201/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Passwordless identity management solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Nightfall AI",
			Website: "https://www.nightfall.ai/",
			Careers: "https://www.nightfall.ai/careers",
			About:   "https://www.nightfall.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19031754,
				Alias:             "nightfall-ai",
				Name:              "Nightfall AI",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "110",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "nightfallai",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nightfall",
				Employees: "60",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nightfall-AI-EI_IE3097406.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nightfall-AI-Reviews-E3097406.htm",
			},
			OttaProfileSlug:   "Nightfall-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer",
							ShortDescription: "Experience in programming with Go, C++, Java or a related language",
							URL:              "https://www.linkedin.com/jobs/view/4068295195/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Cloud-native data protection platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Cyberhaven",
			Website: "https://www.cyberhaven.com/",
			Careers: "https://www.cyberhaven.com/careers",
			About:   "https://www.cyberhaven.com/about",
			Blog:    "https://www.cyberhaven.com/engineering-blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10639445,
				Alias:             "cyberhaven",
				Name:              "Cyberhaven",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "181",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "CyberhavenInc",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cyberhaven",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cyberhaven-EI_IE2985068.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cyberhaven-Reviews-E2985068.htm",
			},
			OttaProfileSlug:   "Cyberhaven",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer",
							ShortDescription: "You get to work with a modern and constantly evolving microservices-based software stack which includes Go, Kubernetes, BigTable, BigQuery, Spanner, Redis, Docker, etcd",
							URL:              "https://www.linkedin.com/jobs/view/4056670384/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Data detection & response platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Fonoa",
			Website: "https://www.fonoa.com/",
			Careers: "https://www.fonoa.com/company/careers",
			About:   "https://www.fonoa.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30129363,
				Alias:             "fonoa",
				Name:              "Fonoa",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "140",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Fonoa-Tech",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fonoa-EI_IE4125781.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fonoa-Reviews-E4125781.htm",
			},
			OttaProfileSlug:   "Fonoa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer",
							ShortDescription: "For the backend, we use Node and Go",
							URL:              "https://www.linkedin.com/jobs/view/3960607071/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Tax automation for the internet economy",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Payrails",
			Website: "https://www.payrails.com/",
			Careers: "https://www.payrails.com/careers",
			About:   "https://www.payrails.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75580859,
				Alias:             "payrails",
				Name:              "Payrails",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "68",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Payrails",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "You show vast experience in developing software in large-scale production systems, ideally using Golang",
							URL:              "https://www.linkedin.com/jobs/view/4026258568/",
							Date:             mustDate("2024-09-25"),
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
			ShortDescription: "Operating system for payments and financial services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teleport",
			Website: "https://goteleport.com/",
			Careers: "https://goteleport.com/careers/",
			About:   "https://goteleport.com/about/",
			Blog:    "https://goteleport.com/blog/tags/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7941233,
				Alias:             "go-teleport",
				Name:              "Teleport",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gravitational",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teleport",
				Employees: "100",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teleport-EI_IE1967263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teleport-Reviews-E1967263.htm",
			},
			OttaProfileSlug:   "Teleport",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 116,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Software Engineer",
							ShortDescription: "Teleport is an open source project, written in Golang with web-based UIs in JavaScript with React",
							URL:              "https://www.linkedin.com/jobs/view/4000249330/",
							Date:             mustDate("2024-12-05"),
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
			ShortDescription: "Identity-native infrastructure access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Bigblue",
			Website: "https://www.bigblue.co/",
			Careers: "https://www.bigblue.co/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18776044,
				Alias:             "bigblue-co",
				Name:              "Bigblue",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "198",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bigblue-EI_IE3379658.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bigblue-Reviews-E3379658.htm",
			},
			OttaProfileSlug:   "Bigblue",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/UU1oL2xl",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Developer of eCommerce logistics tools",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Taxfix",
			Website: "https://taxfix.de/en/",
			Careers: "https://taxfix.de/en/careers/",
			About:   "https://taxfix.de/en/what-is-taxfix/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10791106,
				Alias:             "taxfix",
				Name:              "Taxfix",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "339",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "taxfix",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "taxfix",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Taxfix-EI_IE1301207.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Taxfix-Reviews-E1301207.htm",
			},
			OttaProfileSlug:   "Taxfix",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang backend",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123164602/",
							Date:                 mustDate("2025-01-15"),
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
			ShortDescription: "Tax filing app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Clerk",
			Website: "https://clerk.com/",
			Careers: "https://jobs.ashbyhq.com/clerk",
			About:   "https://clerk.com/company",
			Blog:    "https://clerk.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69336355,
				Alias:             "clerkinc",
				Name:              "Clerk.com",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "147",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "clerk",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Clerk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Identity tool for React applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ClearScore",
			Website: "https://clearscore.com/",
			Careers: "https://www.clearscore.com/careers",
			About:   "https://www.clearscore.com/about-us",
			Blog:    "https://medium.com/clearscore",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9320086,
				Alias:             "clearscore",
				Name:              "ClearScore",
				Followers:         "46K",
				Employees:         "201-500",
				AssociatedMembers: "380",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ClearScore",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ClearScore-EI_IE1046600.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ClearScore-Reviews-E1046600.htm",
			},
			OttaProfileSlug:   "ClearScore",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 2,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/RmSslgyY",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Financial wellbeing marketplace",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "OneSchema",
			Website: "https://www.oneschema.co/",
			Careers: "https://www.oneschema.co/careers",
			About:   "",
			Blog:    "https://www.oneschema.co/blog?category=Engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74704820,
				Alias:             "oneschema",
				Name:              "OneSchema",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "21",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "oneschema",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSchema-EI_IE7681431.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSchema-Reviews-E7681431.htm",
			},
			OttaProfileSlug:   "OneSchema",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Our stack consists of Ruby, Rust, TypeScript, and React. We use Postgres, Redis, and S3 for data storage",
							URL:              "https://www.linkedin.com/jobs/view/3969153699/",
							Date:             mustDate("2024-07-26"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Cloud-based data importing and validation platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "StrongDM",
			Website: "https://www.strongdm.com/",
			Careers: "https://www.strongdm.com/careers",
			About:   "https://www.strongdm.com/about",
			Blog:    "https://www.strongdm.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9446266,
				Alias:             "strongdm",
				Name:              "StrongDM",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "145",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "strongdm",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "strongdm",
				Employees: "170",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-strongDM-EI_IE4514684.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/strongDM-Reviews-E4514684.htm",
			},
			OttaProfileSlug:   "strongDM",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 15,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Design and implementation of highly concurrent, scalable and high-performance distributed systems in Go",
							URL:              "https://www.linkedin.com/jobs/view/3978025966/",
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
			ShortDescription: "Manages auditable database access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JMA Wireless",
			Website: "https://jmawireless.com/",
			Careers: "https://jmawireless.com/careers/",
			About:   "https://jmawireless.com/we-are-jma/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3105986,
				Alias:             "jmawireless",
				Name:              "JMA Wireless",
				Followers:         "28K",
				Employees:         "501-1K",
				AssociatedMembers: "730",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "jma-wireless",
				Employees: "1,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JMA-Wireless-EI_IE875877.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JMA-Wireless-Reviews-E875877.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Backend Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048251309/",
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
			ShortDescription:          "JMA Wireless designs and delivers cutting-edge wireless technology",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ON2IT Cybersecurity",
			Website: "https://on2it.net/",
			Careers: "https://on2it.recruitee.com/",
			About:   "https://on2it.net/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                489607,
				Alias:             "on2it-cybersecurity",
				Name:              "ON2IT Cybersecurity",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "96",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "on2itsecurity",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ON2IT-EI_IE1075463.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ON2IT-Reviews-E1075463.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior PHP/Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3513136315/",
							Date:             mustDate("2024-12-20"),
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
			ShortDescription: "CyberSecurity service provider",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cycloid",
			Website: "https://www.cycloid.io/",
			Careers: "https://www.cycloid.io/careers",
			About:   "https://www.cycloid.io/platform-engineering",
			Blog:    "https://blog.cycloid.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10071522,
				Alias:             "cycloid",
				Name:              "Cycloid - Sustainable Platform Engineering",
				Followers:         "11K",
				Employees:         "11-50",
				AssociatedMembers: "25",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cycloidio",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cycloid",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cycloid-io-EI_IE1562422.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cycloid-io-Reviews-E1562422.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Cycloid",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 40,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075626268/",
							Date:             mustDate("2024-12-12"), // mustDate("2024-11-17"),
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
			ShortDescription: "Hybrid cloud DevOps platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Circuit",
			Website: "https://circuit.ai/",
			Careers: "https://jobs.lever.co/getcircuit",
			About:   "https://circuit.ai/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                36014538,
				Alias:             "circuitknowledge",
				Name:              "Circuit",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "57",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Circuit-TX-EI_IE10004565.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Circuit-TX-Reviews-E10004565.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Go Backend Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3984171089/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:                "Principal Go Backend Engineer",
							ShortDescription:     "5+ years of experience with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115788127/",
							Date:                 mustDate("2025-02-03"), // mustDate("2025-01-03"),
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
			ShortDescription: "AI-powered platform transforms information into actionable knowledge",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aylo",
			Website: "https://www.aylo.com/",
			Careers: "https://www.aylo.com/careers/category-details/",
			About:   "https://www.aylo.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98389423,
				Alias:             "ayloservices",
				Name:              "Aylo",
				Followers:         "22K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,520",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MindGeekOSS",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aylo-EI_IE9049055.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aylo-Reviews-E9049055.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go — Senior Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3968183921/",
							Date:             mustDate("2025-01-06"), // mustDate("2024-12-20"),
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4004286857/",
							Date:                 mustDate("2025-02-07"), // mustDate("2025-01-16"),
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
			ShortDescription: "Adult content platforms",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FactSet",
			Website: "https://www.factset.com/",
			Careers: "https://www.factset.com/careers",
			About:   "https://www.factset.com/our-company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163755,
				Alias:             "factset",
				Name:              "FactSet",
				Followers:         "243K",
				Employees:         "10K+",
				AssociatedMembers: "13,540",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "factset",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "factset",
				Employees: "10,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FactSet-EI_IE6066.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FactSet-Reviews-E6066.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Factset",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Full Stack Engineering — Must Have Go (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4018413250/",
							Date:             mustDate("2024-12-05"),
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
			ShortDescription: "Providing instant access to financial data and analytics that investors use to make decisions",
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
			Name:    "Voodoo",
			Website: "https://voodoo.io/",
			Careers: "https://voodoo.io/careers",
			About:   "https://voodoo.io/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5353630,
				Alias:             "voodoo.io",
				Name:              "Voodoo",
				Followers:         "77K",
				Employees:         "501-1K",
				AssociatedMembers: "694",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "VoodooTeam",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voodoo-EI_IE1889746.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voodoo-Reviews-E1889746.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3966903286/",
							Date:             mustDate("2025-01-14"), // mustDate("2024-12-04"),
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
			ShortDescription: "Company that creates mobile games and apps",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "home24",
			Website: "https://www.home24.com/",
			Careers: "https://home24.career.softgarden.de/en/",
			About:   "https://www.home24.com/websites/homevierundzwanzig/English/1100/who-we-are.html",
			Blog:    "https://home24.tech.blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2653180,
				Alias:             "home24",
				Name:              "home24",
				Followers:         "11K",
				Employees:         "1K-5K",
				AssociatedMembers: "447",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "home24",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-home24-SE-EI_IE3166011.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/home24-SE-Reviews-E3166011.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3872092645/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Home & living e-commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Australian Broadcasting Corporation",
			Website: "https://www.abc.net.au/",
			Careers: "https://www.abc.net.au/careers",
			About:   "https://www.abc.net.au/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2242,
				Alias:             "australian-broadcasting-corporation",
				Name:              "Australian Broadcasting Corporation",
				Followers:         "342K",
				Employees:         "1K-5K",
				AssociatedMembers: "8,087",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "abcnews",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Australian-Broadcasting-EI_IE416033.11,34.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Australian-Broadcasting-Reviews-E416033.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Middle Software Engineer | Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050919710/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Independent public broadcaster",
			Industries:       []domain.Industry{
				// Media
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Nitro",
			Website: "https://www.gonitro.com/",
			Careers: "https://www.gonitro.com/about/careers",
			About:   "https://www.gonitro.com/about/our-story",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76281,
				Alias:             "go-nitro",
				Name:              "Nitro Software",
				Followers:         "16K",
				Employees:         "201-500",
				AssociatedMembers: "325",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "nitro",
				Verified: true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nitro-EI_IE402100.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nitro-Reviews-E402100.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4042171556/",
							Date:             mustDate("2025-01-12"), // mustDate("2024-11-08"),
						},
						{
							Title:            "Scala Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4133037478/",
							Date:             mustDate("2025-01-23"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Document productivity SaaS company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Appodeal",
			Website: "https://appodeal.com/",
			Careers: "https://appodeal.com/career/",
			About:   "https://appodeal.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19139082,
				Alias:             "appodeal",
				Name:              "Appodeal",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "199",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "appodeal",
				Verified: true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Appodeal-EI_IE4606703.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Appodeal-Reviews-E4606703.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Engineer / Real-Time Bidding",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4068403021/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4068296805/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tala",
			Website: "https://tala.co/",
			Careers: "https://tala.co/careers/",
			About:   "https://tala.co/about/",
			Blog:    "https://tala.co/blog/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2647513,
				Alias:             "talamobile",
				Name:              "Tala",
				Followers:         "52K",
				Employees:         "501-1K",
				AssociatedMembers: "792",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "inventure",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tala",
				Employees: "480",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tala-EI_IE1264165.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tala-Reviews-E1264165.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Tala",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048262224/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Financial platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MacroHealth",
			Website: "https://macrohealth.com/",
			Careers: "https://macrohealth.com/careers/",
			About:   "https://macrohealth.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19012089,
				Alias:             "macrohealth",
				Name:              "MacroHealth",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "120",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "macrohealth",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MacroHealth-EI_IE3277977.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MacroHealth-Reviews-E3277977.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4051275869/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Health Market-as-a-Service platform that enables Payers, Providers, and Health Market Partners",
			Industries:       []domain.Industry{
				// HealthTech, eCommerce?
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			// Part of DoorDash
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Wolt",
			Website: "https://wolt.com/",
			Careers: "https://careers.wolt.com/en",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3995846,
				Alias:             "wolt-oy",
				Name:              "Wolt",
				Followers:         "177K",
				Employees:         "10K+",
				AssociatedMembers: "9,722",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "woltapp",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "wolt",
				Employees: "5,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wolt-EI_IE1800227.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wolt-Reviews-E1800227.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4008937766/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Company making it easy to discover and get restaurants and shops delivered home & to the office",
			Industries:       []domain.Industry{
				// Food delivery
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bynder",
			Website: "https://www.bynder.com/en/",
			Careers: "https://careers.bynder.com/",
			About:   "https://www.bynder.com/en/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2427738,
				Alias:             "bynder",
				Name:              "Bynder",
				Followers:         "38K",
				Employees:         "501-1K",
				AssociatedMembers: "603",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bynder",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bynder",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bynder-EI_IE1003066.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bynder-Reviews-E1003066.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Bynder",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Software Engineer (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4011968287/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Digital asset management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "OVO",
			Website: "https://company.ovo.com/",
			Careers: "https://careers.ovo.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                643101,
				Alias:             "ovoenergy",
				Name:              "OVO",
				Followers:         "70K",
				Employees:         "5K-10K",
				AssociatedMembers: "3,258",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ovotech",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OVO-EI_IE767881.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OVO-Reviews-E767881.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 29,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4051591843/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Clean affordable energy",
			Industries:       []domain.Industry{
				// Energy
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Avetta",
			Website: "https://www.avetta.com/",
			Careers: "https://www.avetta.com/careers",
			About:   "https://www.avetta.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7584447,
				Alias:             "avetta",
				Name:              "Avetta",
				Followers:         "60K",
				Employees:         "501-1K",
				AssociatedMembers: "886",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "avetta",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "avetta",
				Employees: "830",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Avetta-EI_IE1182913.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Avetta-Reviews-E1182913.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Software Engineer — Backend",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049869381/",
							Date:             mustDate("2025-01-16"), // mustDate("2024-11-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Contractor risk management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Adverty",
			Website: "https://adverty.com/",
			Careers: "https://adverty.com/join-us/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15233801,
				Alias:             "adverty",
				Name:              "Adverty",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "40",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adverty-EI_IE7208426.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adverty-Reviews-E7208426.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Scala Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053026269/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "In-game ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unzer",
			Website: "https://www.unzer.com/en/", // https://www.youpaylater.com/en/",
			Careers: "https://www.unzer.com/en/jobs/",
			About:   "https://www.unzer.com/en/about-unzer/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67571244,
				Alias:             "unzer",
				Name:              "Unzer",
				Followers:         "11K",
				Employees:         "501-1K",
				AssociatedMembers: "605",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "unzerdev",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unzer-EI_IE2871480.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unzer-Reviews-E2871480.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Unzer",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4016971905/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "All-in-one payment solutions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HERE Technologies",
			Website: "https://www.here.com/",
			Careers: "https://www.here.com/about/careers",
			About:   "https://www.here.com/company/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3237134,
				Alias:             "here",
				Name:              "HERE Technologies",
				Followers:         "238K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,155",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "heremaps",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "here-technologies",
				Employees: "8,140",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HERE-Technologies-EI_IE753677.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HERE-Technologies-Reviews-E753677.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal / Lead Software Engineer - Rust",
							ShortDescription:     "Algorithmic and Mathematics",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119774316/",
							Date:                 mustDate("2025-02-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Scala/Java)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053090497/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Location platform company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "UBS",
			Website: "https://www.ubs.com/",
			Careers: "https://www.ubs.com/global/en/careers.html",
			About:   "https://www.ubs.com/global/en/assetmanagement/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1214,
				Alias:             "ubs",
				Name:              "UBS",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "115,948",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ubs",
				Employees: "72,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-UBS-EI_IE3419.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/UBS-Reviews-E3419.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Data Engineer — Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034024007/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Data Engineer (Spark, Databricks, Azure, Scala & Java)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4097642374/",
							Date:             mustDate("2025-01-20"),
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
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "DroneShield",
			Website: "https://www.droneshield.com/",
			Careers: "https://www.droneshield.com/careers",
			About:   "https://www.droneshield.com/about",
			Blog:    "https://www.droneshield.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7575557,
				Alias:             "droneshield",
				Name:              "DroneShield",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "201",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DroneShield-EI_IE5175947.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DroneShield-Reviews-E5175947.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4032752026/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Go Backend Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4128907541/",
							Date:             mustDate("2025-01-18"),
						},
						{
							Title:                "Full Stack Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4150851032/",
							Date:                 mustDate("2025-02-11"),
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
			ShortDescription: "Counter-UAS solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Automox",
			Website: "https://www.automox.com/",
			Careers: "https://www.automox.com/careers",
			About:   "https://www.automox.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7941257,
				Alias:             "automox",
				Name:              "Automox",
				Followers:         "11K",
				Employees:         "201-500",
				AssociatedMembers: "252",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "automox",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "automox",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Automox-EI_IE2236326.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Automox-Reviews-E2236326.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Automox",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4076698129/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Cloud-native automated patch management platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps, // Programming experience in Golang is recommended
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Darktrace",
			Website: "https://darktrace.com/",
			Careers: "https://darktrace.com/careers",
			About:   "https://darktrace.com/company",
			Blog:    "https://darktrace.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5013440,
				Alias:             "darktrace",
				Name:              "Darktrace",
				Followers:         "217K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,741",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "darktrace",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "darktrace",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Darktrace-EI_IE1059406.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Darktrace-Reviews-E1059406.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Darktrace",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Respond (Python & Rust)",
							ShortDescription: "Predominately working with core software modules written in Python and Rust",
							URL:              "https://www.linkedin.com/jobs/view/4056914662/",
							Date:             mustDate("2024-12-05"),
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
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Poppi Technologies",
			Website: "https://www.poppitechnologies.com/",
			Careers: "https://www.poppitechnologies.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93580213,
				Alias:             "poppi-technologies",
				Name:              "Poppi Technologies",
				Followers:         "351",
				Employees:         "11-50",
				AssociatedMembers: "7",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102305954/",
							Date:             mustDate("2024-12-19"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Purpose driven investing hedge fund operating solely to effect positive change across various sectors and industries",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Booz Allen Hamilton",
			Website: "https://www.boozallen.com/",
			Careers: "https://www.boozallen.com/careers.html",
			About:   "https://www.boozallen.com/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1403,
				Alias:             "booz-allen-hamilton",
				Name:              "Booz Allen Hamilton",
				Followers:         "716K",
				Employees:         "10K+",
				AssociatedMembers: "39,622",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "boozallen",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "booz-allen-hamilton",
				Employees: "27,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Booz-Allen-Hamilton-EI_IE2735.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Booz-Allen-Hamilton-Reviews-E2735.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Systems Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056310581/",
							Date:             mustDate("2024-10-25"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flexera",
			Website: "https://www.flexera.com/",
			Careers: "https://www.flexera.com/about-us/careers",
			About:   "https://www.flexera.com/about-us",
			Blog:    "https://www.flexera.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                574962,
				Alias:             "flexera",
				Name:              "Flexera",
				Followers:         "111K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,669",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "flexera-public",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flexera",
				Employees: "1,060",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flexera-EI_IE304392.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flexera-Reviews-E304392.htm",
				Verified:    true, // openCompany ???
			},
			OttaProfileSlug:   "Flexera",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 32,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Design, implement and run microservices written in Go and C#",
							URL:              "https://www.linkedin.com/jobs/view/4033986551/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "IT management solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Flowhub",
			Website: "https://flowhub.com/",
			Careers: "https://flowhub.com/careers",
			About:   "https://flowhub.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4812785,
				Alias:             "flowhub",
				Name:              "Flowhub",
				Followers:         "36K",
				Employees:         "51-200",
				AssociatedMembers: "81",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flowhub",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flowhub-EI_IE1000410.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flowhub-Reviews-E1000410.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Backend",
							ShortDescription: "Comfortable working in or learning React.js, Golang, GraphQL, PostgreSQL",
							URL:              "https://www.linkedin.com/jobs/view/4053457618/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Cannabis retail platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Teachable",
			Website: "https://teachable.com/",
			Careers: "https://teachable.com/careers",
			About:   "https://teachable.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4815896,
				Alias:             "teachable",
				Name:              "Teachable",
				Followers:         "37K",
				Employees:         "51-200",
				AssociatedMembers: "294",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "usefedora",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teachable",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teachable-EI_IE1614390.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teachable-Reviews-E1614390.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Teachable",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer II, Golang — Mobile",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055915782/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Online teaching platform",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lloyds Banking Group",
			Website: "https://www.lloydsbankinggroup.com/",
			Careers: "https://www.lloydsbankinggroup.com/careers.html",
			About:   "https://www.lloydsbankinggroup.com/who-we-are/group-overview.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                417361,
				Alias:             "lloyds-banking-group",
				Name:              "Lloyds Banking Group",
				Followers:         "485K",
				Employees:         "10K+",
				AssociatedMembers: "60,748",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lloyds-banking-group",
				Employees: "53,970",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer and DevOps Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056705575/",
							Date:             mustDate("2024-10-25"),
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
			Name:    "CyberArk",
			Website: "https://www.cyberark.com/",
			Careers: "https://www.cyberark.com/careers/",
			About:   "",
			Blog:    "https://developer.cyberark.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26630,
				Alias:             "cyber-ark-software",
				Name:              "CyberArk",
				Followers:         "224K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,825",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cyberark",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cyberark",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CyberArk-EI_IE30042.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CyberArk-Reviews-E30042.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 22,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077219379/",
							Date:             mustDate("2024-11-19"),
						},
						{
							Title:            "Senior Full Stack Engineer — Flows Team (Go & Angular)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102677424/",
							Date:             mustDate("2025-02-07"), // mustDate("2025-01-10"), // mustDate("2024-12-19"),
						},
						{
							Title:                "Software Engineer — Java/Golang",
							ShortDescription:     "Machine Identity Security",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134711176/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           true, // $101.000 — $140.000 per year plus commissions or discretionary bonus, which will be based on the employee’s performance.
							Remote:               false,
						},
						{
							Title:                "Software Engineer — Java/Golang",
							ShortDescription:     "Machine Identity Security",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4150919090/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           true, // $101.000 — $160.000 per year
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
			ShortDescription: "Machine Identity Security",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "PayNearMe",
			Website: "https://home.paynearme.com/",
			Careers: "https://home.paynearme.com/about/careers/",
			About:   "https://home.paynearme.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1137260,
				Alias:             "paynearme",
				Name:              "PayNearMe",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "227",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "paynearme",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "paynearme",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PayNearMe-EI_IE954212.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PayNearMe-Reviews-E954212.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058466151/",
							Date:             mustDate("2024-10-25"),
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Creative Fabrica",
			Website: "https://www.creativefabrica.com/",
			Careers: "https://www.creativefabrica.com/tag/careers/",
			About:   "https://www.creativefabrica.com/product/about-me/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10854260,
				Alias:             "creative-fabrica",
				Name:              "Creative Fabrica",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "380",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "creativefabrica",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "creative-fabrica",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Creative-Fabrica-EI_IE3892703.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Creative-Fabrica-Reviews-E3892703.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4001594627/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101753682/",
							Date:             mustDate("2025-01-13"),
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
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "CAST AI",
			Website: "https://cast.ai/",
			Careers: "https://castai.teamtailor.com/jobs",
			About:   "https://cast.ai/about-us/",
			Blog:    "https://cast.ai/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30767670,
				Alias:             "cast-ai",
				Name:              "CAST AI",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "218",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "castai",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cast-ai",
				Employees: "75",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cast-AI-EI_IE3133117.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cast-AI-Reviews-E3133117.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "CAST-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 26,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Wire Team",
							ShortDescription: "Golang is our language",
							URL:              "https://www.linkedin.com/jobs/view/3982723810/",
							Date:             mustDate("2024-07-26"),
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
			ShortDescription: "Cloud cost optimization & automation",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Grasshopper",
			Website: "https://grasshopperasia.com/",
			Careers: "https://grasshopperasia.com/careers/",
			About:   "https://grasshopperasia.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14373895,
				Alias:             "grasshopperasia",
				Name:              "Grasshopper",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "110",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ghpr-asia",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "grasshopper",
				Employees: "75",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grasshopper-EI_IE1145365.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grasshopper-Reviews-E1145365.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3848418692/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Quantitative trading technology provide",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Swish Analytics",
			Website: "https://swishanalytics.com/",
			Careers: "https://swishanalytics.com/careers",
			About:   "https://swishanalytics.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3886282,
				Alias:             "swish-analytics",
				Name:              "Swish Analytics",
				Followers:         "9K",
				Employees:         "11-50",
				AssociatedMembers: "122",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "swishanalytics",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "swish-analytics",
				Employees: "45",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Swish-Analytics-EI_IE2104030.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Swish-Analytics-Reviews-E2104030.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4054776468/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Sports Betting Solutions",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Abbott",
			Website: "https://www.abbott.com/",
			Careers: "https://www.abbott.com/careers.html",
			About:   "https://www.abbott.com/about-abbott.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1612,
				Alias:             "abbott-",
				Name:              "Abbott",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "164,356",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "abbott",
				Employees: "103,530",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Abbott-EI_IE12.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Abbott-Reviews-E12.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Cloud (Go/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4076610998/",
							Date:                 mustDate("2024-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Cloud (Go/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128066961/",
							Date:                 mustDate("2025-02-12"), // mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Go/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3850663049/",
							Date:                 mustDate("2025-02-04"),
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dealer Tire",
			Website: "https://www.dealertire.com/",
			Careers: "https://www.dealertire.com/careers/",
			About:   "https://www.dealertire.com/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                21738,
				Alias:             "dealer-tire",
				Name:              "Dealer Tire",
				Followers:         "53K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,434",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dealer-Tire-EI_IE39809.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dealer-Tire-Reviews-E39809.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full Stack Golang Web Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053146259/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "Distributor of tires and parts",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JPMorganChase",
			Website: "https://www.jpmorganchase.com/",
			Careers: "https://www.jpmorganchase.com/careers",
			About:   "https://www.jpmorganchase.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1068,
				Alias:             "jpmorganchase",
				Name:              "JPMorganChase",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "301,391",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "jpmorganchase",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "jpmorgan-chase",
				Employees: "290,730",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JPMorganChase-EI_IE5224839.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JPMorganChase-Reviews-E5224839.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer III — Golang, AWS, Terraform",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050118492/",
							Date:             mustDate("2024-11-03"),
						},
						{
							Title:            "Software Engineer II —  Core Engineering —  Go / Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105782609/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:            "Software Engineer II —  Core Engineering —  Go / Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4130260943/",
							Date:             mustDate("2025-02-10"), // mustDate("2025-01-19"),
						},
						{
							Title:                "Lead Software Engineer — Core Engineering — Go / Python",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133921095/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           false,
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
							Title:            "Lead Software Engineer — Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055220537/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
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
			Name:    "Outreach",
			Website: "https://www.outreach.io/",
			Careers: "https://www.outreach.io/company/working-at-outreach",
			About:   "https://www.outreach.io/company/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3559595,
				Alias:             "outreach-saas",
				Name:              "Outreach",
				Followers:         "220K",
				Employees:         "501-1K",
				AssociatedMembers: "1,317",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "getoutreach",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "outreach",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Outreach-EI_IE1276320.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Outreach-Reviews-E1276320.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Outreach",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 63,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer (Backend — Golang)",
							ShortDescription: "We primarily use Go language for our next gen applications",
							URL:              "https://www.linkedin.com/jobs/view/4074568575/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Sales execution platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tucows",
			Website: "https://www.tucows.com/",
			Careers: "https://www.tucows.com/careers/",
			About:   "https://www.tucows.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166530,
				Alias:             "tucows",
				Name:              "Tucows",
				Followers:         "36K",
				Employees:         "501-1K",
				AssociatedMembers: "841",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tucows",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tucows",
				Employees: "630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tucows-EI_IE6018.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tucows-Reviews-E6018.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Tucows",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Backend",
							ShortDescription: "Proficiency in Python or Golang programming languages",
							URL:              "https://www.linkedin.com/jobs/view/4004609791/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Domain registrar",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ING",
			Website: "https://www.ing.com/",
			Careers: "",
			About:   "https://www.ing.com/About-us.htm",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2594164,
				IDs:               []int{2594164, 299201, 1523089, 387202, 71625717, 3702090, 3622172, 40711, 516340, 18071794, 2038590, 698107, 2105579, 2924419, 2722195, 215713},
				Alias:             "ing",
				Name:              "ING",
				Followers:         "559K",
				Employees:         "10K+",
				AssociatedMembers: "66,484",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ing-bank",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ing",
				Employees: "65,050",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ING-EI_IE4264.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ING-Reviews-E4264.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Scala Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053346980/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Bank",
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
			Name:    "Remitly",
			Website: "https://www.remitly.com/",
			Careers: "https://careers.remitly.com/",
			About:   "https://www.remitly.com/gb/en/home/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2261199,
				Alias:             "remitly",
				Name:              "Remitly",
				Followers:         "59K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,947",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "remitly",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "remitly",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Remitly-EI_IE1044836.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Remitly-Reviews-E1044836.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Remitly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Development Engineer with Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050285116/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Payments company",
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
			Name:    "Level All",
			Website: "https://www.levelall.com/",
			Careers: "https://www.levelall.com/categories/careers",
			About:   "https://www.levelall.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79136550,
				Alias:             "level-all",
				Name:              "Level All",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "88",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Level-All",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Level-All",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "We’re looking for a Senior Elixir Software Engineer to contribute to our core web platform and supporting applications built in Elixir, Phoenix, LiveView, TypeScript, and Tailwind, AlpineJS backed with a Postgres database",
							URL:              "https://www.linkedin.com/jobs/view/4092211710/",
							Date:             mustDate("2024-12-06"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Personalized college & career guidance",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LittleLives",
			Website: "https://www.littlelives.com/",
			Careers: "https://www.littlelives.com/about/join-the-team",
			About:   "https://www.littlelives.com/about/why-littlelives",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6577681,
				Alias:             "littlelives",
				Name:              "LittleLives",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "76",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "teamlittlelives",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LittleLives-EI_IE959760.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LittleLives-Reviews-E959760.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Elixir + PHP",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4046403243/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "School Management System",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BILL",
			Website: "https://www.bill.com/",
			Careers: "https://www.bill.com/careers",
			About:   "https://www.bill.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                113254,
				Alias:             "bill",
				Name:              "BILL",
				Followers:         "60K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,137",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "billcom",
				Employees: "900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BILL-EI_IE801594.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BILL-Reviews-E801594.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3906024674/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Financial automation software",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Community",
			Website: "https://www.community.com/",
			Careers: "https://www.community.com/careers",
			About:   "https://www.community.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18945668,
				Alias:             "communitydotcom",
				Name:              "Community",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "81",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "communitycom",
				Employees: "100",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-community-com-EI_IE3005175.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/community-com-Reviews-E3005175.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Community",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software & Platform Engineer (Elixir-Focused)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4041216496/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "SMS marketing platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "BigPay",
			Website: "https://bigpayme.com/",
			Careers: "https://bigpayme.com/more/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1067442,
				Alias:             "bigpayme",
				Name:              "BigPay",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "190",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bigpay",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BigPay-EI_IE2497022.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BigPay-Reviews-E2497022.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4052948882/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Fintech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rabbet",
			Website: "https://rabbet.com/",
			Careers: "https://rabbet.com/career/",
			About:   "https://rabbet.com/about-us/",
			Blog:    "https://rabbet.com/blog/category/developers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18937694,
				Alias:             "rabbet",
				Name:              "Rabbet",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Rabbet",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rabbet",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rabbet-EI_IE1805037.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rabbet-Reviews-E1805037.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Full Stack Engineer (Elixir / React)",
							ShortDescription: "Back-end is primarily Elixir/Phoenix/Absinthe and hosted in GCP",
							URL:              "https://www.linkedin.com/jobs/view/4041341979/",
							Date:             mustDate("2024-10-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Agentero",
			Website: "https://agentero.com/",
			Careers: "https://agentero.com/carriers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17946876,
				Alias:             "agentero",
				Name:              "Agentero",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "42",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "agentero",
				Employees: "31",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Agentero-EI_IE2215731.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Agentero-Reviews-E2215731.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer",
							ShortDescription: "You will write clean, robust, and performant software using Go Programming Language",
							URL:              "https://www.linkedin.com/jobs/view/4050442624/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "Digital insurance network",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Marbill Technologies",
			Website: "https://marbill.com/",
			Careers: "https://www.marbill.com/people/carreer",
			About:   "https://www.marbill.com/company/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18914675,
				Alias:             "marbill-technologies",
				Name:              "Marbill Technologies",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "95",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "5+ years experience with Go and engineering software platforms",
							URL:              "https://www.linkedin.com/jobs/view/4057281091/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Platform for the e-commerce industry",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BiTaksi",
			Website: "https://bitaksi.com/",
			Careers: "https://www.bitaksi.com/career",
			About:   "https://www.bitaksi.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2995287,
				Alias:             "bitaksi",
				Name:              "BiTaksi",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "118",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "BiTaksi",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BiTaksi-EI_IE2994794.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BiTaksi-Reviews-E2994794.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4053855198/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "Taxi-hailing app",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Oddin.gg",
			Website: "https://oddin.gg/",
			Careers: "https://oddin.gg/careers",
			About:   "https://oddin.gg/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19137419,
				Alias:             "oddingg",
				Name:              "Oddin.gg",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "73",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "oddin-gg",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oddin-EI_IE5084636.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oddin-Reviews-E5084636.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4052970416/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Esports betting ecosystem",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mimecast",
			Website: "https://www.mimecast.com/",
			Careers: "https://careers.mimecast.com/",
			About:   "https://www.mimecast.com/why-mimecast/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                55895,
				Alias:             "mimecast",
				Name:              "Mimecast",
				Followers:         "102K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,608",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mimecast",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mimecast",
				Employees: "1,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mimecast-EI_IE221309.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mimecast-Reviews-E221309.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Mimecast",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3987632672/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Cloud cybersecurity services",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Sentry",
			Website: "https://sentry.io/",
			Careers: "https://sentry.io/careers/",
			About:   "https://sentry.io/about/",
			Blog:    "https://blog.sentry.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6424460,
				Alias:             "getsentry",
				Name:              "Sentry",
				Followers:         "20K",
				Employees:         "201-500",
				AssociatedMembers: "392",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "getsentry",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sentry",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sentry-CA-EI_IE1622271.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sentry-CA-Reviews-E1622271.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Sentry",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 22,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go), SDK",
							ShortDescription:     "Work in the client infrastructure team to improve and evolve our https://github.com/getsentry/sentry-go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4042362901/",
							Date:                 mustDate("2025-01-03"), // mustDate("2024-11-16"),
							WithSalary:           true,                   // The salary that Sentry offers you is adjusted to the current market situation and our starting point is EUR 95,000 gross per year
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Go), SDK",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127771201/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           true, // 53,592 - 79,000 eur / year
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 67,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Core Product",
							ShortDescription: "Familiarity with Rust",
							URL:              "https://www.linkedin.com/jobs/view/3997074001/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Error tracking and app monitoring software",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Track24",
			Website: "https://www.track24.com/",
			Careers: "https://www.track24.com/join-us/",
			About:   "https://www.track24.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                315653,
				Alias:             "track24",
				Name:              "Track24",
				Followers:         "9K",
				Employees:         "11-50",
				AssociatedMembers: "33",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Track24-EI_IE1463252.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Track24-Reviews-E1463252.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Track24",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Elixir Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055781039/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Operational risk management and communications solutions",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hiive",
			Website: "https://www.hiive.com/",
			Careers: "https://www.hiive.com/careers",
			About:   "https://www.hiive.com/our-story",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76581616,
				Alias:             "hiive",
				Name:              "Hiive",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "105",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hiive",
				Employees: "98",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hiive-Canada-EI_IE6926718.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hiive-Canada-Reviews-E6926718.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4060816724/",
							Date:                 mustDate("2025-01-16"), // mustDate("2024-10-29"), // mustDate("2024-12-31"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "All-in-one liquidity platform for private companies and their shareholders",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ExpressVPN",
			Website: "https://www.expressvpn.com/",
			Careers: "https://www.expressvpn.com/jobs",
			About:   "https://www.expressvpn.com/about-us",
			Blog:    "https://www.expressvpn.com/blog/author/dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14488856,
				Alias:             "expressvpn",
				Name:              "ExpressVPN",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "237",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "expressvpn",
				Verified: true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "expressvpn",
				Employees: "240",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ExpressVPN-EI_IE3355140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ExpressVPN-Reviews-E3355140.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "ExpressVPN",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058715430/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Shuttle",
			Website: "https://www.shuttle.dev/",
			Careers: "https://www.shuttle.dev/about#careers",
			About:   "https://www.shuttle.dev/about",
			Blog:    "https://www.shuttle.dev/blog/tags/all",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67174759,
				Alias:             "shuttle-yc",
				Name:              "Shuttle",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "shuttle-hq",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "shuttle",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 27,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://www.ycombinator.com/companies/shuttle/jobs/pEYmbIL-senior-software-engineer",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Serverless Rust platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Second Spectrum",
			Website: "https://www.secondspectrum.com/",
			Careers: "https://www.geniussports.com/careers/",
			About:   "https://www.geniussports.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3562222,
				Alias:             "second-spectrum",
				Name:              "Second Spectrum",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "181",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "second-spectrum",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Realtime Systems, Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3930229478/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Worldpay",
			Website: "https://www.worldpay.com/",
			Careers: "https://jobs.worldpay.com/",
			About:   "https://corporate.worldpay.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                260387,
				Alias:             "worldpay",
				Name:              "Worldpay",
				Followers:         "221K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,605",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Scala & AWS",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3982972845/",
							Date:             mustDate("2024-11-06"),
						},
						{
							Title:            "DevOps Engineer (Go & Python)",
							ShortDescription: "Kubernetes, Terraform & Scripting",
							URL:              "https://www.linkedin.com/jobs/view/4130478145/",
							Date:             mustDate("2025-01-20"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exabeam",
			Website: "https://www.exabeam.com/",
			Careers: "https://www.exabeam.com/company/careers/",
			About:   "https://www.exabeam.com/company/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3181474,
				Alias:             "exabeam",
				Name:              "Exabeam",
				Followers:         "42K",
				Employees:         "501-1K",
				AssociatedMembers: "880",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "exabeam",
				Employees: "660",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Java/Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048214298/",
							Date:             mustDate("2024-11-06"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Selby Jennings",
			Website: "https://www.selbyjennings.co.uk/",
			Careers: "https://www.selbyjennings.co.uk/jobs",
			About:   "https://www.selbyjennings.co.uk/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106584,
				Alias:             "selby-jennings",
				Name:              "Selby Jennings",
				Followers:         "789K",
				Employees:         "1K-5K",
				AssociatedMembers: "833",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Developer | Digital Asset Trading",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059312901/",
							Date:             mustDate("2025-01-20"), // mustDate("2024-10-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DataBeacon",
			Website: "https://databeacon.aero/",
			Careers: "https://databeacon.aero/careers/",
			About:   "https://databeacon.aero/about/",
			Blog:    "https://datascience.aero/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37898796,
				Alias:             "databeaconaero",
				Name:              "DataBeacon",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "8",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Software Engineer, Rust & Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4060807594/",
							Date:             mustDate("2024-11-01"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MNTN",
			Website: "https://mountain.com/",
			Careers: "https://mountain.com/about/careers/",
			About:   "https://mountain.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                480382,
				Alias:             "wearemntn",
				Name:              "MNTN",
				Followers:         "112K",
				Employees:         "201-500",
				AssociatedMembers: "663",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mntn",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Rust Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4064837533/",
							Date:             mustDate("2024-11-04"),
						},
						{
							Title:                "Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4150806532/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Writer",
			Website: "https://writer.com/",
			Careers: "https://writer.com/company/careers/",
			About:   "https://writer.com/company/about/",
			Blog:    "https://writer.com/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67088679,
				Alias:             "getwriter",
				Name:              "Writer",
				Followers:         "49K",
				Employees:         "501-1K",
				AssociatedMembers: "1,351",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "writer",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Backend (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3972187342/",
							Date:             mustDate("2025-01-21"), // mustDate("2025-01-13"), // mustDate("2024-11-04"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "dLocal",
			Website: "https://www.dlocal.com/",
			Careers: "https://jobs.lever.co/dlocal",
			About:   "https://www.dlocal.com/company/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15156062,
				Alias:             "dlocal",
				Name:              "dLocal",
				Followers:         "140K",
				Employees:         "501-1K",
				AssociatedMembers: "1,129",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dlocal",
				Employees: "535",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Full Stack Engineer (Golang and Node)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4070887973/",
							Date:             mustDate("2024-11-07"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Clearstory.build",
			Website: "https://www.clearstory.build/",
			Careers: "https://www.clearstory.build/careers",
			About:   "https://www.clearstory.build/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18591580,
				Alias:             "clearstory-build",
				Name:              "Clearstory.build",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "61",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Golang Developer — Philippines",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034614195/",
							Date:             mustDate("2024-11-02"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Millennium",
			Website: "https://www.mlp.com/",
			Careers: "https://www.mlp.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164987,
				Alias:             "millennium-partners",
				Name:              "Millennium",
				Followers:         "258K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,236",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "millennium",
				Employees: "3,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3943478838/",
							Date:             mustDate("2024-11-03"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Mondoo",
			Website: "https://mondoo.com/",
			Careers: "https://jobs.ashbyhq.com/mondoo",
			About:   "https://mondoo.com/about-us",
			Blog:    "https://mondoo.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75622407,
				Alias:             "mondoo",
				Name:              "Mondoo",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "53",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mondo",
				Employees: "640",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4065503885/",
							Date:             mustDate("2024-11-01"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "fullinfo",
			Website: "https://www.fullinfo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71773445,
				Alias:             "fullinfo",
				Name:              "fullinfo",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "13",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4065504018/",
							Date:             mustDate("2024-11-01"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bloomberg",
			Website: "https://www.bloomberg.com/",
			Careers: "https://www.bloomberg.com/company/early-careers/",
			About:   "https://www.bloomberg.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2494,
				Alias:             "bloomberg",
				Name:              "Bloomberg",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "24,725",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bloomberg",
				Employees: "23,260",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062165305/",
							Date:             mustDate("2025-01-07"), // mustDate("2024-11-01"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sainsbury's",
			Website: "https://www.sainsburys.co.uk/",
			Careers: "https://sainsburys.jobs/",
			About:   "https://about.sainsburys.co.uk/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5015,
				Alias:             "sainsburys",
				Name:              "Sainsbury's",
				Followers:         "388K",
				Employees:         "10K+",
				AssociatedMembers: "61,975",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sainsburys",
				Employees: "111,900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Fullstack Engineer — React & Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062980858/",
							Date:             mustDate("2024-11-04"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stratascale — An SHI Company",
			Website: "https://www.stratascale.com/",
			Careers: "https://careers-stratascale.icims.com/jobs",
			About:   "https://www.stratascale.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4791832,
				Alias:             "stratascale-llc",
				Name:              "Stratascale – An SHI Company",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "233",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4005692040/",
							Date:             mustDate("2025-01-16"), // mustDate("2024-11-01"),
							WithSalary:       true,                   // $132K — $240K per year
							Remote:           true,
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gartner",
			Website: "https://www.gartner.com/",
			Careers: "https://jobs.gartner.com/",
			About:   "https://www.gartner.com/en/about",
			Blog:    "https://www.gartner.com/en/insights",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2078,
				Alias:             "gartner",
				Name:              "Gartner",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "24,462",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gartner",
				Employees: "18,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Backend Software Engineer (Golang) — Peer Experiences",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4037373961/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Extreme Networks",
			Website: "https://www.extremenetworks.com/",
			Careers: "https://www.extremenetworks.com/about-extreme-networks/career",
			About:   "https://www.extremenetworks.com/about-extreme-networks",
			Blog:    "https://www.extremenetworks.com/resources/blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4761,
				Alias:             "extreme-networks",
				Name:              "Extreme Networks",
				Followers:         "165K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,875",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "extreme-networks",
				Employees: "3,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang — Cloud Networking Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055078837/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:                "Senior Back End Developer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134262544/",
							Date:                 mustDate("2025-01-24"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Yalo",
			Website: "https://www.yalo.ai/",
			Careers: "https://www.yalo.ai/careers",
			About:   "https://www.yalo.ai/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10251920,
				Alias:             "yalo",
				Name:              "Yalo",
				Followers:         "106K",
				Employees:         "201-500",
				AssociatedMembers: "310",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "yalo",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4037222694/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "reMarkable",
			Website: "https://remarkable.com/",
			Careers: "https://careers.remarkable.com/",
			About:   "https://remarkable.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10785500,
				Alias:             "remarkable-as",
				Name:              "reMarkable",
				Followers:         "44K",
				Employees:         "201-500",
				AssociatedMembers: "565",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "remarkable",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4061088133/",
							Date:             mustDate("2025-02-11"), // mustDate("2024-11-01"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Synopsys",
			Website: "https://synopsys.com/",
			Careers: "https://careers.synopsys.com/",
			About:   "https://www.synopsys.com/company.html",
			Blog:    "https://www.synopsys.com/blogs.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2457,
				Alias:             "synopsys",
				Name:              "Synopsys",
				Followers:         "737K",
				Employees:         "10K+",
				AssociatedMembers: "19,480",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "synopsys",
				Employees: "12,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior AI Platform Engineer — Golang and Kubernetes",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4027571222/",
							Date:             mustDate("2024-11-01"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PubMatic",
			Website: "https://pubmatic.com/",
			Careers: "https://pubmatic.com/careers/",
			About:   "https://pubmatic.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167624,
				Alias:             "pubmatic",
				Name:              "PubMatic",
				Followers:         "335K",
				Employees:         "501-1K",
				AssociatedMembers: "1,229",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pubmatic",
				Employees: "810",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4044830862/",
							Date:             mustDate("2024-11-01"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fanatics",
			Website: "https://www.fanaticsinc.com/",
			Careers: "https://www.fanaticsinc.com/careers",
			About:   "https://www.fanaticsinc.com/about-fanatics",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68543,
				Alias:             "fanatics-inc-",
				Name:              "Fanatics",
				Followers:         "198K",
				Employees:         "10K+",
				AssociatedMembers: "9,825",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fanatics",
				Employees: "4,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer III",
							ShortDescription: "Our stack primarily uses Elixir with the Phoenix Framework",
							URL:              "https://www.linkedin.com/jobs/view/4011218533/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kantox",
			Website: "https://www.kantox.com/",
			Careers: "https://www.kantox.com/careers",
			About:   "https://www.kantox.com/about-kantox",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1871617,
				Alias:             "kantox",
				Name:              "Kantox",
				Followers:         "39K",
				Employees:         "201-500",
				AssociatedMembers: "291",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kantox",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Elixir Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072221760/",
							Date:             mustDate("2024-10-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmartRent",
			Website: "https://smartrent.com/",
			Careers: "https://smartrent.com/careers/",
			About:   "https://smartrent.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16165131,
				Alias:             "smartrentdotcom",
				Name:              "SmartRent",
				Followers:         "16K",
				Employees:         "501-1K",
				AssociatedMembers: "434",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "smartrent",
				Employees: "640",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071960242/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Home automation solutions for property managers and renters",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Denim",
			Website: "https://www.denim.com/",
			Careers: "https://www.denim.com/careers",
			About:   "https://www.denim.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11026766,
				Alias:             "join-denim",
				Name:              "Denim",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "203",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "denim-social",
				Employees: "45",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Deep experience building large applications with Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4064330407/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Thryve.Earth",
			Website: "https://www.thryve.earth/",
			Careers: "",
			About:   "https://www.thryve.earth/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82124834,
				Alias:             "thryve-earth",
				Name:              "Thryve.Earth",
				Followers:         "11K",
				Employees:         "2-10",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full-Stack Developer (Elixir Phoenix)",
							ShortDescription: "Develop and maintain web applications using Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4058317241/",
							Date:             mustDate("2024-10-28"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Zubale",
			Website: "https://www.zubale.com/",
			Careers: "https://www.zubale.com/careers/",
			About:   "https://www.zubale.com/about-zubale/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33183626,
				Alias:             "zubale",
				Name:              "Zubale",
				Followers:         "104K",
				Employees:         "501-1K",
				AssociatedMembers: "1,152",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Elixir/Scala/Rust Backend Engineer",
							ShortDescription: "More than 3 years of experience in Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4076675943/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Priceline",
			Website: "https://www.priceline.com/",
			Careers: "https://careers.priceline.com/",
			About:   "https://press.priceline.com/our-story/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7451,
				Alias:             "priceline-com",
				Name:              "Priceline",
				Followers:         "72K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,859",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072949018/",
							Date:             mustDate("2025-01-20"), // mustDate("2024-11-14"),
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134499604/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           false,
							Remote:               false,
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "New Relic",
			Website: "https://newrelic.com/",
			Careers: "https://newrelic.com/careers",
			About:   "https://newrelic.com/about",
			Blog:    "https://newrelic.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                426253,
				Alias:             "new-relic-inc-",
				Name:              "New Relic",
				Followers:         "148K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,926",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "new-relic",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058040204/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Data for Engineers",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Riskified",
			Website: "https://www.riskified.com/",
			Careers: "https://www.riskified.com/careers/",
			About:   "https://www.riskified.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2821604,
				Alias:             "riskified",
				Name:              "Riskified",
				Followers:         "53K",
				Employees:         "501-1K",
				AssociatedMembers: "720",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "riskified",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer — App Infra Team",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047670627/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Omilia",
			Website: "https://omilia.com/",
			Careers: "https://omilia.com/careers/",
			About:   "https://omilia.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1057955,
				Alias:             "omilia-ltd",
				Name:              "Omilia - Conversational Intelligence",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "384",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071728951/",
							Date:             mustDate("2024-11-14"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rakuten Viber",
			Website: "https://www.viber.com/",
			Careers: "https://www.viber.com/careers/",
			About:   "https://www.viber.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1710293,
				Alias:             "rakuten-viber",
				Name:              "Rakuten Viber",
				Followers:         "41K",
				Employees:         "201-500",
				AssociatedMembers: "710",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4076696523/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106769274/",
							Date:             mustDate("2024-12-20"),
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "4+ years of experience in server-side development using Go language",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4113717946/",
							Date:                 mustDate("2024-12-31"),
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nebius",
			Website: "https://nebius.com/",
			Careers: "https://nebius.com/careers",
			About:   "https://group.nebius.com/",
			Blog:    "https://nebius.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89802307,
				Alias:             "nebius",
				Name:              "Nebius",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "499",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nebius",
				Employees: "600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4074905433/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102045400/",
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
			ShortDescription: "Cloud platform specifically designed to train AI models",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zoom",
			Website: "https://www.zoom.com/",
			Careers: "https://careers.zoom.us/",
			About:   "https://www.zoom.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2532259,
				Alias:             "zoom",
				Name:              "Zoom",
				Followers:         "588K",
				Employees:         "5K-10K",
				AssociatedMembers: "11,159",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "zoom",
				Employees: "4,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Security Engineer (PHP Laravel & Go) — Workvivo",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075517057/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LexisNexis Risk Solutions",
			Website: "https://risk.lexisnexis.com/",
			Careers: "https://risk.lexisnexis.com/global/en/about-us/careers",
			About:   "https://risk.lexisnexis.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10371941,
				Alias:             "lexisnexis-risk-solutions",
				Name:              "LexisNexis Risk Solutions",
				Followers:         "245K",
				Employees:         "10K+",
				AssociatedMembers: "8,927",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lexisnexis-risk-solutions",
				Employees: "10,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Full Stack Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077459247/",
							Date:             mustDate("2024-11-17"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "SKIDOS",
			Website: "https://www.skidos.com/",
			Careers: "",
			About:   "https://www.skidos.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9246535,
				Alias:             "skidos-games",
				Name:              "SKIDOS",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "33",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075801293/",
							Date:             mustDate("2024-11-16"),
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4112778941/",
							Date:                 mustDate("2024-12-31"),
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SingleStore",
			Website: "https://www.singlestore.com/",
			Careers: "https://www.singlestore.com/careers/",
			About:   "https://www.singlestore.com/company/",
			Blog:    "https://www.singlestore.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71851558,
				Alias:             "singlestore",
				Name:              "SingleStore",
				Followers:         "40K",
				Employees:         "201-500",
				AssociatedMembers: "520",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "singlestore",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Nova Platform (Go & K8s)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4039359993/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Cloud-native database",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Diabolocom",
			Website: "https://www.diabolocom.com/",
			Careers: "https://www.diabolocom.com/jobs/",
			About:   "https://www.diabolocom.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                902343,
				Alias:             "diabolocom",
				Name:              "Diabolocom",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "119",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "diabolocom",
				Employees: "90",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059794979/",
							Date:             mustDate("2024-11-17"),
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
			ShortDescription:          "Contact Center",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FlexAI",
			Website: "https://www.flex.ai/",
			Careers: "",
			About:   "https://www.flex.ai/#about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100770131,
				Alias:             "flexaihq",
				Name:              "FlexAI",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "63",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer — Golang (Senior)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059611387/",
							Date:             mustDate("2025-01-13"), // mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Stockbit",
			Website: "https://stockbit.com/",
			Careers: "https://careers.stockbit.com/",
			About:   "https://stockbit.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6398473,
				Alias:             "stockbit",
				Name:              "Stockbit",
				Followers:         "31K",
				Employees:         "501-1K",
				AssociatedMembers: "454",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "stockbit",
				Employees: "690",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058347229/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teaching Strategies",
			Website: "https://teachingstrategies.com/",
			Careers: "https://teachingstrategiesatwork.com/",
			About:   "https://teachingstrategies.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                56801,
				Alias:             "teaching-strategies-llc",
				Name:              "Teaching Strategies, LLC",
				Followers:         "98K",
				Employees:         "201-500",
				AssociatedMembers: "1,061",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teaching-strategies",
				Employees: "720",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4023232816/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sophos",
			Website: "https://www.sophos.com/",
			Careers: "https://www.sophos.com/en-us/careers",
			About:   "https://www.sophos.com/en-us/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5053,
				Alias:             "sophos",
				Name:              "Sophos",
				Followers:         "481K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,769",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sophos",
				Employees: "4,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang Developer (Network security domain)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4042181078/",
							Date:             mustDate("2024-11-15"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Macquarie Group",
			Website: "https://www.macquarie.com/",
			Careers: "https://www.macquarie.com/careers.html",
			About:   "https://www.macquarie.com/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3537,
				Alias:             "macquariegroup",
				Name:              "Macquarie Group",
				Followers:         "649K",
				Employees:         "10K+",
				AssociatedMembers: "20,814",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "macquarie-group",
				Employees: "17,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Python/Golang",
							ShortDescription: "We use Golang and Python",
							URL:              "https://www.linkedin.com/jobs/view/4022294274/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
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
			Name:    "Ayoconnect",
			Website: "https://www.ayoconnect.com/",
			Careers: "https://www.ayoconnect.com/career",
			About:   "https://www.ayoconnect.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7598856,
				Alias:             "ayoconnect",
				Name:              "Ayoconnect",
				Followers:         "71K",
				Employees:         "201-500",
				AssociatedMembers: "167",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Software Engineer (Golang)",
							ShortDescription: "Design, develop, and maintain software applications using Golang",
							URL:              "https://www.linkedin.com/jobs/view/4057794419/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Payments and Banking Enablement for Indonesia",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Camping World",
			Website: "https://www.campingworld.com/",
			Careers: "https://www.campingworldcareers.com/",
			About:   "https://www.campingworld.com/helpcenter#aboutUs",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                23441,
				Alias:             "campingworld",
				Name:              "Camping World",
				Followers:         "85K",
				Employees:         "10K+",
				AssociatedMembers: "6,875",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Back End Engineer (TypeScript/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056215700/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Shopee",
			Website: "https://shopee.com/",
			Careers: "https://careers.shopee.sg/",
			About:   "https://careers.shopee.sg/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6451760,
				Alias:             "shopee",
				Name:              "Shopee",
				Followers:         "2M",
				Employees:         "5K-10K",
				AssociatedMembers: "47,126",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "shopee",
				Employees: "34,130",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Engineer, Backend (Python/ Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4045101256/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "E-commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Guardsquare",
			Website: "https://www.guardsquare.com/",
			Careers: "https://www.guardsquare.com/careers",
			About:   "https://www.guardsquare.com/about-us",
			Blog:    "https://www.guardsquare.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5012731,
				Alias:             "guardsquare",
				Name:              "Guardsquare",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "163",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3777868160/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Web Engineer (Go/React)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4020630528/",
							Date:             mustDate("2025-02-07"), // mustDate("2025-02-04"), // mustDate("2025-01-20"), // mustDate("2024-12-18"),
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
			ShortDescription: "Mobile application protection",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JFrog",
			Website: "https://jfrog.com/",
			Careers: "https://join.jfrog.com/",
			About:   "https://jfrog.com/about/",
			Blog:    "https://jfrog.com/blog/?pagenum=1&category=devops",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                455737,
				Alias:             "jfrog-ltd",
				Name:              "JFrog",
				Followers:         "66K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,958",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "jfrog",
				Employees: "960",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077432516/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Software supply chain platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HP",
			Website: "https://www.hp.com/",
			Careers: "https://www.hp.com/jobs/",
			About:   "https://www.hp.com/us-en/hp-information.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5390798,
				Alias:             "hp",
				Name:              "HP",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "122,744",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4078297154/",
							Date:             mustDate("2024-11-17"),
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
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mirantis",
			Website: "https://www.mirantis.com/",
			Careers: "https://www.mirantis.com/careers/",
			About:   "https://www.mirantis.com/about/",
			Blog:    "https://www.mirantis.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7793,
				Alias:             "mirantis",
				Name:              "Mirantis",
				Followers:         "70K",
				Employees:         "501-1K",
				AssociatedMembers: "589",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mirantis",
				Employees: "960",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang/Kubernetes)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075754584/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Open source infrastructure for containers and virtual machines",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zscaler",
			Website: "https://www.zscaler.com/",
			Careers: "https://www.zscaler.com/careers",
			About:   "https://www.zscaler.com/company/about-zscaler",
			Blog:    "https://www.zscaler.com/blogs?type=product-insights",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                234625,
				Alias:             "zscaler",
				Name:              "Zscaler",
				Followers:         "387K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,436",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "zscaler",
				Employees: "3,500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer, Data Forwarding (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077097120/",
							Date:             mustDate("2025-01-08"), // mustDate("2024-11-15"),
							WithSalary:       true,                   // $147K — $210K per year
							Remote:           true,
						},
						{
							Title:                "Senior Staff Software Development Engineer",
							ShortDescription:     "Golang, Kubernetes, eBPF",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4023494848/",
							Date:                 mustDate("2025-01-13"),
							WithSalary:           true, // $147K — $210K per year
							Remote:               true,
						},
						{
							Title:                "Staff Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128638661/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           true, // $122.500 — $175K per year
							Remote:               false,
						},
						{
							Title:                "Staff Software Development Engineer",
							ShortDescription:     "Golang, Kubernetes, eBPF",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133963749/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           true, // $122.500 — $175.000 per year
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pure Storage",
			Website: "https://www.purestorage.com/",
			Careers: "https://www.purestorage.com/company/careers.html",
			About:   "https://www.purestorage.com/company.html",
			Blog:    "https://blog.purestorage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1632202,
				Alias:             "purestorage",
				Name:              "Pure Storage",
				Followers:         "493K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,341",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pure-storage",
				Employees: "4,250",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer, FlashArray",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4072597478/",
							Date:                 mustDate("2024-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Engineer, Flash Array",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125984657/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Engineer, Flash Array",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142328113/",
							Date:                 mustDate("2025-02-05"),
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
			ShortDescription: "Data storage",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teya",
			Website: "https://www.teya.com/",
			Careers: "https://www.teya.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93621491,
				Alias:             "teya-global",
				Name:              "Teya",
				Followers:         "65K",
				Employees:         "1K-5K",
				AssociatedMembers: "735",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teya",
				Employees: "1,500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4073753061/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4107304437/",
							Date:             mustDate("2024-12-24"),
						},
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4116504410/",
							Date:             mustDate("2025-01-16"),
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
			Name:    "TomTom",
			Website: "https://www.tomtom.com/",
			Careers: "https://www.tomtom.com/careers/",
			About:   "https://www.tomtom.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166474,
				Alias:             "tomtom",
				Name:              "TomTom",
				Followers:         "177K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,196",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tomtom",
				Employees: "5,500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full-Stack Developer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075571548/",
							Date:             mustDate("2025-01-16"), // mustDate("2024-11-14"),
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
			ShortDescription: "Mapping and location technology",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Postscript",
			Website: "https://postscript.io/",
			Careers: "https://postscript.io/careers",
			About:   "https://postscript.io/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18771172,
				Alias:             "postscriptio",
				Name:              "Postscript",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "295",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "postscript",
				Employees: "150",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071761682/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription:          "SMS marketing platform for ecommerce companies",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Impossible Cloud",
			Website: "https://www.impossiblecloud.com/",
			Careers: "https://www.impossiblecloud.com/careers",
			About:   "https://www.impossiblecloud.com/about-us",
			Blog:    "https://www.impossiblecloud.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76137251,
				Alias:             "impossiblecloud",
				Name:              "Impossible Cloud",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "51",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Engineer — Golang & Distributed Systems",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079088056/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Cloud storage",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Arista Networks",
			Website: "https://www.arista.com/",
			Careers: "https://www.arista.com/careers",
			About:   "https://www.arista.com/company/company-overview",
			Blog:    "https://blogs.arista.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80069,
				Alias:             "arista-networks-inc",
				Name:              "Arista Networks",
				Followers:         "414K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,697",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "arista-networks",
				Employees: "3,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Java/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077123852/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Data-Driven Networking",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PaveNow",
			Website: "https://www.pavenow.io/",
			Careers: "",
			About:   "https://www.pavenow.io/en/aboutus",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96200958,
				Alias:             "pavenow",
				Name:              "PaveNow",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "15",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077178228/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Financial management for small and medium-sized businesses",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LemFi",
			Website: "https://lemfi.com/",
			Careers: "https://lemonade-technology-inc.breezy.hr/",
			About:   "https://lemfi.com/en-es/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68838211,
				Alias:             "lemfi",
				Name:              "LemFi",
				Followers:         "12K",
				Employees:         "11-50",
				AssociatedMembers: "239",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4080247545/",
							Date:             mustDate("2024-11-19"),
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Template
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	Blog:    "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "",
		//		Verified:          false,
		//	},
		//	BlindProfile: domain.BlindProfile{
		//		Alias:       "",
		//		Employees:   "",
		//		Salary:      "",
		//		Reviews:     "",
		//		ReviewsRate: "",
		//	},
		//	LevelsFyiProfile: domain.LevelsFyiProfile{
		//		Alias:     "",
		//		Employees: "",
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//		Verified:    false,
		//	},
		//	IndeedProfile: domain.IndeedProfile{
		//		Alias: "",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	GoMainLanguage:    false,
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{
		//				{
		//					Title:            "",
		//					ShortDescription: "",
		//					URL:              "",
		//					Date:             mustDate(""),
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
		//	ShortDescription:          "",
		//	DealroomURL:               "",
		//	CrunchbaseURL:             "",
		//	PitchbookURL:              "",
		//	YahooFinanceURL:           "",
		//	GoogleFinanceURL:          "",
		//	YCombinatorURL:            "",
		//	Industries:                []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{},
		//},
	}
}
