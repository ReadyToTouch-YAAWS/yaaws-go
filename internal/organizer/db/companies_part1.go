package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart1() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		// Favorites
		// Favorites | ReadyToTouch
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ReadyToTouch",
			Website: "https://readytotouch.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    97909464,
				IDs:   nil,
				Alias: "readytotouch",
				Name:  "ReadyToTouch",
			},
			GitHubProfile:    domain.GitHubProfile{},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Service for simplifying job search",
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
			Ignore: false,
		},

		// Favorites | DocHQ
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DocHQ",
			Website: "https://dochq.co.uk/",
			Careers: "",
			About:   "https://dochq.co.uk/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14136494,
				Alias:             "dochq",
				Name:              "DocHQ",
				Followers:         "1.5k",
				Employees:         "51-200",
				AssociatedMembers: "20+",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "dochq",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
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
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Health Tech company",
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
			},
			Ignore: false,
		},

		// BigTech
		// BigTech | Google
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Google",
			Website: "https://www.google.com/",
			Careers: "https://www.google.com/about/careers/",
			About:   "https://about.google/",
			Blog:    "https://blog.google/technology/developers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1441,
				IDs:               []int{1441, 16140, 791962, 10440912, 17876832, 89982912},
				Alias:             "google",
				Name:              "Google",
				Followers:         "36M",
				Employees:         "10K+",
				AssociatedMembers: "303,369",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "google",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "google",
				Employees:   "10,000+",
				Salary:      "$60K ~ $357K a year",
				Reviews:     "10,557",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "google",
				Employees: "258,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Google-EI_IE9079.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Google-Reviews-E9079.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Google-Jobs-E9079.htm",
				Jobs:        "4.7K",
				Reviews:     "61K",
				Salaries:    "172K",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Google",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 227,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer III, Google Kubernetes Engine",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4087466433/",
							Date:             mustDate("2024-12-04"),
							WithSalary:       false,
							Remote:           false,
						},
						{
							Title:                "Senior Golang Developer Relations Engineer, AI Agents",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115836242/",
							Date:                 mustDate("2025-01-25"), // mustDate("2025-01-03"),
							WithSalary:           true,                   // $142,000-$211,000
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer Relations Engineer, AI Agents",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115837092/",
							Date:                 mustDate("2025-01-27"),
							WithSalary:           true, // $142K — $211K per year
							Remote:               false,
						},
						{
							Title:                "Golang Back-End Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4164764188/",
							Date:                 mustDate("2025-02-25"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 80,
					Vacancies:               nil,
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational technology company",
			DealroomURL:      "https://app.dealroom.co/companies/google",
			CrunchbaseURL:    "https://www.crunchbase.com/organization/google",
			PitchbookURL:     "https://pitchbook.com/profiles/company/10453-33",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/GOOG/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/GOOG:NASDAQ",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
			RustFoundationMember: true,
		},

		// BigTech | Mozilla
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mozilla",
			Website: "https://www.mozilla.org/",
			Careers: "https://www.mozilla.org/careers/",
			About:   "https://www.mozilla.org/about/",
			Blog:    "https://developer.mozilla.org/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13948,
				Alias:             "mozilla-corporation",
				Name:              "Mozilla",
				Followers:         "415K",
				Employees:         "501-1K",
				AssociatedMembers: "1,798",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mozilla",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Mozilla",
				Employees:   "501 to 1,000",
				Salary:      "$66K ~ $292K a year",
				Reviews:     "95",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mozilla",
				Employees: "750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mozilla-EI_IE19129.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mozilla-Reviews-E19129.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mozilla-Jobs-E19129.htm",
				Jobs:        "1",
				Reviews:     "521",
				Salaries:    "916",
				ReviewsRate: "2.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Mozilla-Foundation",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 25,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer, Ads",
							ShortDescription: "Expertise in one of the following Back-End technologies: Golang, Rust or Java",
							URL:              "https://www.linkedin.com/jobs/view/4068076458/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 62,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer, Ads",
							ShortDescription: "Expertise in one of the following Back-End technologies: Golang, Rust or Java",
							URL:              "https://www.linkedin.com/jobs/view/4068076458/",
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
			ShortDescription: "Organization dedicated to making the web better",
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
			RustFoundationMember: true,
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Meta",
			Website: "http://www.facebook.com/",
			Careers: "http://www.facebook.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10667,
				IDs:               []int{10667, 23769, 289891, 2289109, 2763277, 16159097, 27046884, 27104390, 76987811},
				Alias:             "meta",
				Name:              "Meta",
				Followers:         "11M",
				Employees:         "10K+",
				AssociatedMembers: "124,561",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "facebook",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Meta",
				Employees:   "10,000+",
				Salary:      "$72K ~ $350K a year",
				Reviews:     "9,766",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "facebook",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Meta-EI_IE40772.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Meta-Reviews-E40772.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Meta-Jobs-E40772.htm",
				Jobs:        "3.4K",
				Reviews:     "21K",
				Salaries:    "84K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Meta-dd1502f2",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
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
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
			Ignore:                    false,
		},

		// BigTech | Discord
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Discord",
			Website: "https://discord.com/",
			Careers: "https://discord.com/careers",
			About:   "https://discord.com/company",
			Blog:    "https://discord.com/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3765675,
				Alias:             "discord",
				Name:              "Discord",
				Followers:         "427K",
				Employees:         "501-1K",
				AssociatedMembers: "3,028",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "discord",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "discord",
				Employees:   "501 to 1,000",
				Salary:      "$60K ~ $300K a year",
				Reviews:     "141",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "discord",
				Employees: "600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Discord-EI_IE910317.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Discord-Reviews-E910317.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Discord-Jobs-E910317.htm",
				Jobs:        "48",
				Reviews:     "278",
				Salaries:    "650",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Discord",
			},
			OttaProfileSlug:   "Discord",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 36,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Persistence Infrastructure",
							ShortDescription: "Building and operating large-scale, reliable and performant systems with Rust, ScyllaDB, PostgreSQL, ElasticSearch, and Linux",
							URL:              "https://www.linkedin.com/jobs/view/4082697010/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 20,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer, Realtime Infrastructure",
							ShortDescription: "Our tech stack is Elixir, Python and Rust",
							URL:              "https://www.linkedin.com/jobs/view/4102671396/",
							Date:             mustDate("2024-12-20"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Chat",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// BigTech | Figma
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Figma",
			Website: "https://www.figma.com/",
			Careers: "https://www.figma.com/careers/",
			About:   "https://www.figma.com/about/",
			Blog:    "https://www.figma.com/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3650502,
				Alias:             "figma",
				Name:              "Figma",
				Followers:         "2M",
				Employees:         "1K-5K",
				AssociatedMembers: "2,386",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "figma",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "figma",
				Employees:   "51 to 200",
				Salary:      "$75K ~ $205K a year",
				Reviews:     "136",
				ReviewsRate: "4.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "figma",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Figma-EI_IE1537286.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Figma-Reviews-E1537286.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Figma-Jobs-E1537286.htm",
				Jobs:        "",
				Reviews:     "174",
				Salaries:    "463",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Figma",
			},
			OttaProfileSlug:   "Figma",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 15,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — FigFile Platform",
							ShortDescription: "We own services written in Rust, Ruby, TypeScript, and Go",
							URL:              "https://www.linkedin.com/jobs/view/3906200700/",
							Date:             mustDate("2024-12-12"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — FigFile Platform",
							ShortDescription: "We own services written in Rust, Ruby, TypeScript, and Go",
							URL:              "https://www.linkedin.com/jobs/view/3906200700/",
							Date:             mustDate("2024-12-12"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Design platform for product teams",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Microsoft
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Microsoft",
			Website: "https://www.microsoft.com/",
			Careers: "https://careers.microsoft.com/v2/global/en/home.html",
			About:   "https://www.microsoft.com/about",
			Blog:    "https://devblogs.microsoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1035,
				Alias:             "microsoft",
				Name:              "Microsoft",
				Followers:         "24M",
				Employees:         "10K+",
				AssociatedMembers: "244,487",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "microsoft",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "microsoft",
				Employees:   "10,000+",
				Salary:      "$33K ~ $299K a year",
				Reviews:     "10,803",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "microsoft",
				Employees: "182,268",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Microsoft-EI_IE1651.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Microsoft-Reviews-E1651.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Microsoft-Jobs-E1651.htm",
				Jobs:        "1.9K",
				Reviews:     "61K",
				Salaries:    "184K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Microsoft",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Microsoft",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 112,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3905989512/",
							Date:             mustDate("2024-04-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 75,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Product Manager – Rust Tooling",
							ShortDescription: "Join us and play a pivotal role in building the next generation of developer tools and services, enhancing productivity for all Rust developers!",
							URL:              "https://www.linkedin.com/jobs/view/4038788830/",
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
			ShortDescription: "American multinational technology company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
			RustFoundationMember: true,
		},

		// BigTech | Amazon
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amazon",
			Website: "https://www.aboutamazon.com/",
			Careers: "https://www.amazon.jobs/",
			About:   "https://www.aboutamazon.com/about-us",
			Blog:    "https://aws.amazon.com/blogs/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1586,
				IDs:               []int{1586, 12227, 14951, 16551, 17411, 21433, 34924, 46825, 47157, 49318, 61712, 71099, 111446, 167364, 208137, 860467, 1379932, 2320329, 2382910, 2649984, 4787585, 11091426, 15218805, 78392228, 80073065},
				Alias:             "amazon",
				Name:              "Amazon",
				Followers:         "34M",
				Employees:         "10K+",
				AssociatedMembers: "727,302",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "aws", // "amzn",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Amazon",
				Employees:   "10,000+",
				Salary:      "$46K ~ $300K a year",
				Reviews:     "27,199",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "amazon",
				Employees: "865,406",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Amazon-EI_IE6036.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Amazon-Reviews-E6036.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Amazon-Jobs-E6036.htm",
				Jobs:        "19K",
				Reviews:     "221K",
				Salaries:    "657K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Amazon.com",
			},
			OttaProfileSlug:   "Amazon",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
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
			RustFoundationMember: true,
		},

		// BigTech | IBM
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IBM",
			Website: "https://www.ibm.com/",
			Careers: "https://www.ibm.com/careers",
			About:   "https://www.ibm.com/about",
			Blog:    "https://developer.ibm.com/blogs/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1009,
				Alias:             "ibm",
				Name:              "IBM",
				Followers:         "18M",
				Employees:         "10K+",
				AssociatedMembers: "317,108",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "IBM",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "IBM",
				Employees:   "10,000+",
				Salary:      "$50K ~ $280K a year",
				Reviews:     "1,941",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ibm",
				Employees: "524,558",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IBM-EI_IE354.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IBM-Reviews-E354.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IBM-Jobs-E354.htm",
				Jobs:        "297",
				Reviews:     "123K",
				Salaries:    "198K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "IBM",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@IBMTechnology",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 231,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Developer (Back-End) — Go/Python",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4036994106/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075807665/",
							Date:                 mustDate("2024-11-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Development Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115168875/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Developer (Go/Python/Scala)",
							ShortDescription:     "Be one of the technical leaders of watsonx.ai ecosystems, focusing on development in Go, Scala & Python languages",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4095286825/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Strong proficiency in backend development languages such as Go (Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149775881/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Junior Software Developer — Back-End",
							ShortDescription:     "Knowledge of Golang/Python programming",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161292030/",
							Date:                 mustDate("2025-02-21"),
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
							Title:            "Lead Scala Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3998480636/",
							Date:             mustDate("2025-01-13"), // mustDate("2024-11-08"),
						},
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063746170/",
							Date:             mustDate("2024-11-02"),
						},
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4124146753/",
							Date:             mustDate("2025-01-15"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational technology company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | SAP
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SAP",
			Website: "https://www.sap.com/",
			Careers: "https://jobs.sap.com/",
			About:   "https://www.sap.com/about/company.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1115,
				IDs:               []int{1115, 2573558, 2818, 166185, 5822},
				Alias:             "sap",
				Name:              "SAP",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "125,121",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "SAP",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "sap",
				Employees:   "10,000+",
				Salary:      "$31K ~ $269K a year",
				Reviews:     "1,394",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sap",
				Employees: "117,730",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SAP-EI_IE10471.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SAP-Reviews-E10471.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SAP-Jobs-E10471.htm",
				Jobs:        "2.6K",
				Reviews:     "33K",
				Salaries:    "46K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "SAP",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 54,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Software Engineer for Cloud Development",
							ShortDescription:     "Design, develop, and maintain high-quality software applications using Golang and cloud technologies",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4072909509/",
							Date:                 mustDate("2024-12-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4072911263/",
							Date:                 mustDate("2024-12-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go and Kubernetes Developer for Open-Source project Gardener",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4073625824/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Developer for Unified Customer Landscape",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4091947137/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer for Unified Metadata Service",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4072914533/",
							Date:                 mustDate("2025-02-17"), // mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Developer for Unified Customer Landscape",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137984859/",
							Date:                 mustDate("2025-02-19"), // mustDate("2025-01-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer C/C++, Rust, Go",
							ShortDescription:     "Open Source Cloud Infrastructure (IronCore)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149287476/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer C/C++, Rust, Go",
							ShortDescription:     "Open Source Cloud Infrastructure (IronCore)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151455339/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Junior) Cloud Engineer — Python/Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4059368134/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer C/C++, Rust, Go",
							ShortDescription:     "Open Source Cloud Infrastructure (IronCore)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149287476/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer C/C++, Rust, Go",
							ShortDescription:     "Open Source Cloud Infrastructure (IronCore)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151455339/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Developer with Java or Golan",
							ShortDescription:     "SAP iXP Intern",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117406944/",
							Date:                 mustDate("2025-02-26"),
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
			ShortDescription: "Enterprise application and business AI company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Oracle
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Oracle",
			Website: "https://www.oracle.com/",
			Careers: "https://www.oracle.com/careers/",
			About:   "https://www.oracle.com/corporate/",
			Blog:    "https://blogs.oracle.com/developers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1028,
				Alias:             "oracle",
				Name:              "Oracle",
				Followers:         "10M",
				Employees:         "10K+",
				AssociatedMembers: "199,448",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "oracle",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Oracle",
				Employees:   "10,000+",
				Salary:      "$30K ~ $317K a year",
				Reviews:     "2,598",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "oracle",
				Employees: "212,570",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oracle-EI_IE1737.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oracle-Reviews-E1737.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Oracle-Jobs-E1737.htm",
				Jobs:        "3.6K",
				Reviews:     "59K",
				Salaries:    "146K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Oracle",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Software Engineer — Java/Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034901984/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Principal Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4125824194/",
							Date:             mustDate("2025-01-16"),
						},
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4125825071/",
							Date:             mustDate("2025-01-16"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational technology company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social
		// Social | GitHub
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GitHub",
			Website: "https://github.com/",
			Careers: "https://www.github.careers/careers-home",
			About:   "https://github.com/about",
			Blog:    "https://github.blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1418841,
				Alias:             "github",
				Name:              "GitHub",
				Followers:         "5M",
				Employees:         "501-1K",
				AssociatedMembers: "6,253",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "github",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "github",
				Employees:   "501 to 1,000",
				Salary:      "$47K ~ $259K a year",
				Reviews:     "350",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "github",
				Employees: "3,630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitHub-EI_IE671945.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitHub-Reviews-E671945.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GitHub-Jobs-E671945.htm",
				Jobs:        "63",
				Reviews:     "450",
				Salaries:    "1.2K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Github",
			},
			OttaProfileSlug:   "Github",
			YouTubeChannelURL: "https://www.youtube.com/@GitHub",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 39,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Copilot",
							ShortDescription: "Write, review, and maintain code primarily in Go",
							URL:              "https://www.linkedin.com/jobs/view/3914880703/",
							Date:             mustDate("2024-05-24"),
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
			ShortDescription: "Developer platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | GitLab
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GitLab",
			Website: "https://gitlab.com/",
			Careers: "https://about.gitlab.com/jobs/all-jobs/",
			About:   "https://about.gitlab.com/company/",
			Blog:    "https://about.gitlab.com/blog/categories/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5101804,
				Alias:             "gitlab-com",
				Name:              "GitLab",
				Followers:         "1M",
				Employees:         "1K-5K",
				AssociatedMembers: "2,850",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gitlabhq",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "gitlab",
				Employees:   "1,001 to 5,000",
				Salary:      "$55K ~ $200K a year",
				Reviews:     "110",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gitlab",
				Employees: "1,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitLab-EI_IE1296544.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitLab-Reviews-E1296544.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GitLab-Jobs-E1296544.htm",
				Jobs:        "2",
				Reviews:     "633",
				Salaries:    "1K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Gitlab-Inc",
			},
			OttaProfileSlug:   "GitLab-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Intermediate Backend Engineer (Go)",
							ShortDescription: "Significant experience with Go",
							URL:              "https://www.linkedin.com/jobs/view/4018780627/",
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
			ShortDescription: "DevOps platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | LinkedIn
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LinkedIn",
			Website: "https://www.linkedin.com/",
			Careers: "https://www.linkedin.com/jobs",
			About:   "https://about.linkedin.com/",
			Blog:    "https://www.linkedin.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1337,
				Alias:             "linkedin",
				Name:              "LinkedIn",
				Followers:         "29M",
				Employees:         "10K+",
				AssociatedMembers: "25,395",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "linkedin",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "linkedin",
				Employees:   "10,000+",
				Salary:      "$40K ~ $356K a year",
				Reviews:     "2,578",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "linkedin",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LinkedIn-EI_IE34865.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LinkedIn-Reviews-E34865.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LinkedIn-Jobs-E34865.htm",
				Jobs:        "17",
				Reviews:     "8.1K",
				Salaries:    "23K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Linkedin",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Rust & Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4023583351/",
							Date:             mustDate("2024-09-16"),
						},
						{
							Title:            "Senior Software Engineer, Languages (Rust, Go, & Python) Tooling",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103336471/",
							Date:             mustDate("2025-02-10"), // mustDate("2024-02-18"),
						},
						{
							Title:                "Senior Software Engineer, Languages (Rust, Go, & Python) Dev Infra",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127847711/",
							Date:                 mustDate("2025-02-10"), // mustDate("2025-01-17"),
							WithSalary:           true,                   // The pay range for this role is $121,000 - $198,000.
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Rust Expertise",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4073651431/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Business and employment-focused social media platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Reddit
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Reddit",
			Website: "https://www.reddit.com/",
			Careers: "https://redditinc.com/careers",
			About:   "https://redditinc.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                150573,
				Alias:             "reddit-com",
				Name:              "Reddit, Inc.",
				Followers:         "400K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,697",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "reddit",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Reddit",
				Employees:   "201 to 500",
				Salary:      "$56K ~ $377K a year",
				Reviews:     "313",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "reddit",
				Employees: "1,820",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Reddit-EI_IE796358.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Reddit-Reviews-E796358.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Reddit-Jobs-E796358.htm",
				Jobs:        "",
				Reviews:     "370",
				Salaries:    "1.1K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Reddit",
			},
			OttaProfileSlug:   "Reddit-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 27,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer — Ads Billing",
							ShortDescription: "Languages: Python, Scala, Go",
							URL:              "https://www.linkedin.com/jobs/view/4075867913/",
							Date:             mustDate("2024-12-11"),
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer — Ads Billing",
							ShortDescription: "Languages: Python, Scala, Go",
							URL:              "https://www.linkedin.com/jobs/view/4075867913/",
							Date:             mustDate("2024-12-11"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Network of communities where people can dive into their interests, hobbies and passions",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Social | Medium
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Medium",
			Website: "https://medium.com/",
			Careers: "https://medium.pinpointhq.com/",
			About:   "https://medium.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3039001,
				Alias:             "medium-com",
				Name:              "Medium",
				Followers:         "191K",
				Employees:         "51-200",
				AssociatedMembers: "8,665",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Medium",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Medium",
				Employees:   "201 to 500",
				Salary:      "$120K ~ $365K a year",
				Reviews:     "24",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "medium",
				Employees: "4,680",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Medium-EI_IE784883.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Medium-Reviews-E784883.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Medium-Jobs-E784883.htm",
				Jobs:        "3",
				Reviews:     "120",
				Salaries:    "193",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Medium.com",
			},
			OttaProfileSlug:   "Medium",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Core Platform Engineer",
							ShortDescription: "3+ years experience in software engineering, specifically in Golang or similar backend language and interacting with API’s",
							URL:              "https://www.linkedin.com/jobs/view/3910516815/",
							Date:             mustDate("2024-04-25"),
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
			ShortDescription: "Home for human stories and ideas",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Pinterest
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pinterest",
			Website: "https://www.pinterest.com/",
			Careers: "https://www.pinterestcareers.com/",
			About:   "https://help.pinterest.com/en/guide/all-about-pinterest",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1124131,
				Alias:             "pinterest",
				Name:              "Pinterest",
				Followers:         "950K",
				Employees:         "1K-5K",
				AssociatedMembers: "10,702",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pinterest",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Pinterest",
				Employees:   "1,001 to 5,000",
				Salary:      "$60K ~ $332K a year",
				Reviews:     "798",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pinterest",
				Employees: "5,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinterest-EI_IE503467.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinterest-Reviews-E503467.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pinterest-Jobs-E503467.htm",
				Jobs:        "62",
				Reviews:     "1.1K",
				Salaries:    "3.3K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Pinterest",
			},
			OttaProfileSlug:   "Pinterest",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Software Engineer",
							ShortDescription: "Experience in Python, Java, C++, or Go or another language and a willingness to learn",
							URL:              "https://www.linkedin.com/jobs/view/3947425956/",
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
			ShortDescription: "Visual discovery engine for finding ideas",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Snap
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Snap",
			Website: "https://snap.com/",
			Careers: "https://careers.snap.com/",
			About:   "https://investor.snap.com/about-snap/",
			Blog:    "https://eng.snap.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15191764,
				Alias:             "snap-inc-co",
				Name:              "Snap Inc.",
				Followers:         "471K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,527",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Snapchat",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Snap",
				Employees:   "1,001 to 5,000",
				Salary:      "$76K ~ $339K a year",
				Reviews:     "960",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "snap",
				Employees: "6,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snap-EI_IE671946.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snap-Reviews-E671946.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Snap-Jobs-E671946.htm",
				Jobs:        "272",
				Reviews:     "1.3K",
				Salaries:    "4.2K",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Snap-Inc.",
			},
			OttaProfileSlug:   "Snap",
			YouTubeChannelURL: "https://www.youtube.com/@snapinc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Highly proficient in Java, Golang, NodeJs, and/or Python",
							URL:              "https://www.linkedin.com/jobs/view/4103915582/",
							Date:             mustDate("2024-12-19"),
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
			ShortDescription: "American multimedia instant messaging app",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | BeReal.
		{
			ID:      0,
			Type:    "",
			Name:    "BeReal.",
			Website: "https://bereal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34731272,
				Alias:             "bereal-app",
				Name:              "BeReal.",
				Followers:         "23K",
				Employees:         "51-200",
				AssociatedMembers: "71",
				Verified:          true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "BeReal-App",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BeReal-EI_IE7468524.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BeReal-Reviews-E7468524.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BeReal-Jobs-E7468524.htm",
				Jobs:        "",
				Reviews:     "15",
				Salaries:    "12",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:  "BeReal",
			ShortDescription: "Simplest photo sharing app",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{},
			Ignore:                    true, // Voodoo acquires BeReal to take authentic social network to new heights https://voodoo.io/news/voodoo-acquires-bereal
		},

		// Favorites
		// Favorites | VictoriaMetrics
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VictoriaMetrics",
			Website: "https://victoriametrics.com/",
			Careers: "",
			About:   "https://victoriametrics.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30169914,
				Alias:             "victoriametrics",
				Name:              "VictoriaMetrics",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "33",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "VictoriaMetrics",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				// NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Site Reliability Engineer",
							ShortDescription: "Ability to understand, write and review code in Go",
							URL:              "https://www.linkedin.com/jobs/view/4096907990/",
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
			ShortDescription: "Time series database and easy-to-use high performance monitoring solutions",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Favorites | Grammarly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Grammarly",
			Website: "https://www.grammarly.com/",
			Careers: "https://www.grammarly.com/careers",
			About:   "https://www.grammarly.com/about",
			Blog:    "https://www.grammarly.com/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1253671,
				Alias:             "grammarly",
				Name:              "Grammarly",
				Followers:         "223K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,728",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "grammarly",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Grammarly",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "147",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "grammarly",
				Employees: "720",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grammarly-EI_IE636873.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grammarly-Reviews-E636873.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Grammarly-Jobs-E636873.htm",
				Jobs:        "37",
				Reviews:     "347",
				Salaries:    "458",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Grammarly-5",
			},
			OttaProfileSlug:   "Grammarly",
			YouTubeChannelURL: "https://www.youtube.com/@grammarly",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "System Engineer, Enterprise Infrastructure",
							ShortDescription: "Writing code by using Golang, PowerShell, Bash",
							URL:              "https://www.linkedin.com/jobs/view/3983712751/",
							Date:             mustDate("2024-12-04"),
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
			ShortDescription: "Typing assistant",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Favorites | Influ2
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Influ2",
			Website: "https://www.influ2.com/",
			Careers: "https://jobs.lever.co/influ2",
			About:   "https://www.influ2.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18119989,
				Alias:             "influ2",
				Name:              "Influ2",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "120",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "influ2",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Influ2-EI_IE4066744.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Influ2-Reviews-E4066744.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Influ2-Jobs-E4066744.htm",
				Jobs:        "",
				Reviews:     "10",
				Salaries:    "3",
				ReviewsRate: "4.8",
				Verified:    true,
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
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058445939/",
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
			ShortDescription: "Person-based advertising platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech
		// Tech | Netlify
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netlify",
			Website: "https://www.netlify.com/",
			Careers: "https://www.netlify.com/careers/",
			About:   "https://www.netlify.com/about/",
			Blog:    "https://www.netlify.com/blog/tags/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6392431,
				Alias:             "netlify",
				Name:              "Netlify",
				Followers:         "29K",
				Employees:         "201-500",
				AssociatedMembers: "202",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "netlify",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Netlify",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "12",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "netlify",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netlify-EI_IE1426251.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netlify-Reviews-E1426251.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netlify-Jobs-E1426251.htm",
				Jobs:        "",
				Reviews:     "38",
				Salaries:    "100",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Netlify",
			},
			OttaProfileSlug:   "Netlify",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 35,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Backend Engineer",
							ShortDescription: "Experience developing production-level Golang",
							URL:              "https://app.welcometothejungle.com/jobs/qdIzdfOu",
							Date:             mustDate("2024-10-18"), // The date is approximate because it is missing on the job listing page
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
			ShortDescription: "Static hosting service",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Vercel
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vercel",
			Website: "https://vercel.com/",
			Careers: "https://vercel.com/careers",
			About:   "https://vercel.com/about",
			Blog:    "https://vercel.com/blog/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16181286,
				Alias:             "vercel",
				Name:              "Vercel",
				Followers:         "124K",
				Employees:         "201-500",
				AssociatedMembers: "601",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "vercel",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Vercel",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "23",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vercel",
				Employees: "45",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vercel-EI_IE6510369.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vercel-Reviews-E6510369.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vercel-Jobs-E6510369.htm",
				Jobs:        "63",
				Reviews:     "83",
				Salaries:    "161",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Vercel",
			},
			OttaProfileSlug:   "Vercel",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Compute",
							ShortDescription: "You will be writing Golang on a daily basis, while terraform will be your go-to tool when it comes to provisioning the needed infra",
							URL:              "https://www.linkedin.com/jobs/view/4048326475/",
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
			ShortDescription: "Static hosting service",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Tech | Fastly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fastly",
			Website: "https://www.fastly.com/",
			Careers: "https://www.fastly.com/about/careers/current-openings",
			About:   "https://www.fastly.com/company",
			Blog:    "https://www.fastly.com/blog/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2602522,
				Alias:             "fastly",
				Name:              "Fastly",
				Followers:         "57K",
				Employees:         "501-1K",
				AssociatedMembers: "1,292",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fastly",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Fastly",
				Employees:   "501 to 1,000",
				Salary:      "$20K ~ $424K a year",
				Reviews:     "65",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fastly",
				Employees: "990",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fastly-EI_IE814479.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fastly-Reviews-E814479.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fastly-Jobs-E814479.htm",
				Jobs:        "3",
				Reviews:     "312",
				Salaries:    "722",
				ReviewsRate: "2.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Fastly",
			},
			OttaProfileSlug:   "Fastly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 35,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer — API Services",
							ShortDescription: "Contribute, support and deploy applications and scripts written in Go to cloud environments",
							URL:              "https://www.linkedin.com/jobs/view/4077452954/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 29,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Edge Applications",
							ShortDescription: "Develop, test, and maintain new features in Rust for a variety of products Fastly offers",
							URL:              "https://www.linkedin.com/jobs/view/4097075408/",
							Date:             mustDate("2024-12-21"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Edge cloud platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Tech | Dropbox
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dropbox",
			Website: "https://www.dropbox.com/",
			Careers: "https://jobs.dropbox.com/all-jobs",
			About:   "https://www.dropbox.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167251,
				IDs:               nil,
				Alias:             "dropbox",
				Name:              "Dropbox",
				Followers:         "475K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,551",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "dropbox",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "dropbox",
				Employees:   "1,001 to 5,000",
				Salary:      "$67K ~ $350K a year",
				Reviews:     "887",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dropbox",
				Employees: "4,050",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dropbox-EI_IE415350.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dropbox-Reviews-E415350.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dropbox-Jobs-E415350.htm",
				Jobs:        "115",
				Reviews:     "1.7K",
				Salaries:    "3.5K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Dropbox-ab4b15a9",
			},
			OttaProfileSlug:   "Dropbox",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 22,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Product Software Engineer, Sync",
							ShortDescription: "Highly skilled at developing and debugging in Rust, C, C++, or Go",
							URL:              "https://www.linkedin.com/jobs/view/3983369109/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Product Software Engineer, Sync",
							ShortDescription: "Highly skilled at developing and debugging in Rust, C, C++, or Go",
							URL:              "https://www.linkedin.com/jobs/view/3983369109/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "File hosting service",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/DBX/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/DBX:NASDAQ",
			YCombinatorURL:   "https://www.ycombinator.com/companies/dropbox",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
			RustFoundationMember: true,
		},

		// Tech | Docker
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Docker",
			Website: "https://www.docker.com/",
			Careers: "https://www.docker.com/careers/",
			About:   "https://www.docker.com/company/",
			Blog:    "https://www.docker.com/blog/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1301808,
				Alias:             "docker",
				Name:              "Docker, Inc",
				Followers:         "707K",
				Employees:         "501-1K",
				AssociatedMembers: "941",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "docker",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Docker-EI_IE1089506.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Docker-Reviews-E1089506.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Docker-Jobs-E1089506.htm",
				Jobs:        "14",
				Reviews:     "137",
				Salaries:    "267",
				ReviewsRate: "3.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Docker-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 49,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Multi-platform application for developers",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Grafana Labs
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Grafana Labs",
			Website: "https://grafana.com/",
			Careers: "https://grafana.com/about/careers/",
			About:   "",
			Blog:    "https://grafana.com/categories/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11062162,
				Alias:             "grafana-labs",
				Name:              "Grafana Labs",
				Followers:         "213K",
				Employees:         "1K-5K",
				AssociatedMembers: "1420",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "grafana",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Grafana-Labs",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "21",
				ReviewsRate: "4.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "grafana",
				Employees: "630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grafana-Labs-EI_IE2300269.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grafana-Labs-Reviews-E2300269.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Grafana-Labs-Jobs-E2300269.htm",
				Jobs:        "2",
				Reviews:     "188",
				Salaries:    "447",
				ReviewsRate: "4.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Grafana-Labs",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCYCwgQAMm9sTJv0rgwQLCxw",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 332,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Adaptive Telemetry",
							ShortDescription: "We use Go, but if you have familiarity with Python, C, C++, Rust or similar then that translates well",
							URL:              "https://www.linkedin.com/jobs/view/4076482530/",
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
			ShortDescription: "Open source monitoring and observability platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | HashiCorp
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HashiCorp",
			Website: "https://www.hashicorp.com/",
			Careers: "https://www.hashicorp.com/careers",
			About:   "https://www.hashicorp.com/about",
			Blog:    "https://www.hashicorp.com/blog/tags/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2830763,
				Alias:             "hashicorp",
				Name:              "HashiCorp",
				Followers:         "287K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,476",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hashicorp",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "hashicorp",
				Employees:   "501 to 1,000",
				Salary:      "$81K ~ $248K a year",
				Reviews:     "209",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hashicorp",
				Employees: "1,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HashiCorp-EI_IE1359860.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HashiCorp-Reviews-E1359860.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HashiCorp-Jobs-E1359860.htm",
				Jobs:        "193",
				Reviews:     "425",
				Salaries:    "1K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Hashicorp",
			},
			OttaProfileSlug:   "HashiCorp",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 471,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer II",
							ShortDescription: "You have professional experience developing with modern programming languages and frameworks, and are interested in working in Golang and Ruby specifically",
							URL:              "https://www.linkedin.com/jobs/view/4102997754/",
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
			ShortDescription: "Software company that provides modular DevOps infrastructure provisioning and management products",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | CrowdStrike
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CrowdStrike",
			Website: "https://www.crowdstrike.com/",
			Careers: "https://www.crowdstrike.com/careers/",
			About:   "https://www.crowdstrike.com/about-us/",
			Blog:    "https://www.crowdstrike.com/en-us/blog/category.engineering-and-technology/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2497653,
				Alias:             "crowdstrike",
				Name:              "CrowdStrike",
				Followers:         "775K",
				Employees:         "5K-10K",
				AssociatedMembers: "9,641",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "CrowdStrike",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "crowdstrike",
				Employees:   "1,001 to 5,000",
				Salary:      "$34K ~ $285K a year",
				Reviews:     "215",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "crowdstrike",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CrowdStrike-EI_IE795976.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CrowdStrike-Reviews-E795976.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CrowdStrike-Jobs-E795976.htm",
				Jobs:        "407",
				Reviews:     "1.1K",
				Salaries:    "2.3K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Crowdstrike",
			},
			OttaProfileSlug:   "CrowdStrike-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer III, Identity Protection (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3996818515/",
							Date:             mustDate("2024-11-03"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Rust (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4040870810/",
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
			ShortDescription: "Cybersecurity technology company",
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

		// Tech | Cockroach Labs
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cockroach Labs",
			Website: "https://www.cockroachlabs.com/",
			Careers: "https://www.cockroachlabs.com/careers/",
			About:   "https://www.cockroachlabs.com/about/",
			Blog:    "https://www.cockroachlabs.com/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9309408,
				Alias:             "cockroach-labs",
				Name:              "Cockroach Labs",
				Followers:         "104K",
				Employees:         "501-1K",
				AssociatedMembers: "661",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cockroachdb",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cockroach-Lab",
				Employees:   "51 to 200",
				Salary:      "$100K ~ $190K a year",
				Reviews:     "33",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cockroach-labs",
				Employees: "360",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cockroach-Labs-EI_IE1168502.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cockroach-Labs-Reviews-E1168502.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cockroach-Labs-Jobs-E1168502.htm",
				Jobs:        "",
				Reviews:     "78",
				Salaries:    "205",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cockroach-Labs",
			},
			OttaProfileSlug:   "Cockroach-Labs",
			YouTubeChannelURL: "https://www.youtube.com/@cockroachdb",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 102,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Backend (Disaster Recovery)",
							ShortDescription: "Develop in Go, but if you don't know it you'll learn while you're here",
							URL:              "https://www.linkedin.com/jobs/view/4062743637/",
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
			ShortDescription: "Distributed database with standard SQL for cloud applications",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Timescale
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Timescale",
			Website: "https://www.timescale.com/",
			Careers: "https://www.timescale.com/careers",
			About:   "https://www.timescale.com/about",
			Blog:    "https://www.timescale.com/blog/tag/product",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11065434,
				Alias:             "timescaledb",
				Name:              "Timescale",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "164",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "timescale",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Timescale",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "timescale",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Timescale-EI_IE2214356.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Timescale-Reviews-E2214356.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Timescale-Jobs-E2214356.htm",
				Jobs:        "2",
				Reviews:     "48",
				Salaries:    "49",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Timescale",
			YouTubeChannelURL: "https://www.youtube.com/@TimescaleDB",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Backend",
							ShortDescription: "Design, develop, and maintain Golang-based services",
							URL:              "https://www.linkedin.com/jobs/view/4024039384/",
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
			ShortDescription: "Relational open-source database with PostgreSQL",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | ScyllaDB
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ScyllaDB",
			Website: "https://www.scylladb.com/",
			Careers: "https://www.scylladb.com/company/careers/",
			About:   "https://www.scylladb.com/company/",
			Blog:    "https://www.scylladb.com/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10201068,
				Alias:             "scylladb",
				Name:              "ScyllaDB",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "214",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "scylladb",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "scylladb",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ScyllaDB-EI_IE1622223.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ScyllaDB-Reviews-E1622223.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ScyllaDB-Jobs-E1622223.htm",
				Jobs:        "1",
				Reviews:     "42",
				Salaries:    "35",
				ReviewsRate: "4.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "ScyllaDB",
			YouTubeChannelURL: "https://www.youtube.com/@ScyllaDB",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 19,
					Vacancies: []domain.Vacancy{
						{
							Title:                "FullStack Software Engineer (Golang, React)",
							ShortDescription:     "Design & Build APIs with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4113720305/", // redirected from https://www.linkedin.com/jobs/view/3897167605/
							Date:                 mustDate("2024-12-31"),
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
			ShortDescription: "ScyllaDB is the distributed database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Percona
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Percona",
			Website: "https://www.percona.com/",
			Careers: "https://www.percona.com/about/careers",
			About:   "https://www.percona.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                421929,
				Alias:             "percona",
				Name:              "Percona",
				Followers:         "24K",
				Employees:         "201-500",
				AssociatedMembers: "345",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "percona",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Percona",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "5.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "percona",
				Employees: "280",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-EI_IE283779.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Percona-Reviews-E283779.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Percona-Jobs-E283779.htm",
				Jobs:        "28",
				Reviews:     "138",
				Salaries:    "121",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Percona",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 47,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Senior Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4110883947/",
							Date:             mustDate("2024-12-27"),
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
			ShortDescription: "Open source database software",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Elastic
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Elastic",
			Website: "https://www.elastic.co/",
			Careers: "https://www.elastic.co/careers",
			About:   "https://www.elastic.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                814025,
				Alias:             "elastic-co",
				Name:              "Elastic",
				Followers:         "447K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,151",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "elastic",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Elastic",
				Employees:   "1,001 to 5,000",
				Salary:      "$79K ~ $314K a year",
				Reviews:     "135",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "elastic",
				Employees: "2,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Elastic-EI_IE751551.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Elastic-Reviews-E751551.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Elastic-Jobs-E751551.htm",
				Jobs:        "36",
				Reviews:     "840",
				Salaries:    "1.5K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Elastic",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 131,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4001888204/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:                "Software Engineer II (Go, K8s)",
							ShortDescription:     "Our services run on multiple cloud service provider platforms and are built on Docker, Kubernetes, Go/Scala, and custom orchestration architectures",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4069010453/",
							Date:                 mustDate("2025-01-20"), // mustDate("2025-01-03"),
							WithSalary:           true,                   // $110,900—$175,500 USD
							Remote:               true,
						},
						{
							Title:                "Software Engineer II (Go)",
							ShortDescription:     "Platform — Ingest",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133600886/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-24"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Scala, JVM, IAM",
							ShortDescription:     "Control Plane",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4110414724/",
							Date:                 mustDate("2025-02-22"),
							WithSalary:           true, // $133.100 - $252.900 per year
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Search engine",
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

		// Tech | MongoDB
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MongoDB",
			Website: "https://www.mongodb.com/",
			Careers: "https://www.mongodb.com/company/careers",
			About:   "https://www.mongodb.com/company",
			Blog:    "https://www.mongodb.com/blog/channel/engineering-blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                783611,
				Alias:             "mongodbinc",
				Name:              "MongoDB",
				Followers:         "813K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,990",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mongodb",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "MongoDB",
				Employees:   "1,001 to 5,000",
				Salary:      "$91K ~ $287K a year",
				Reviews:     "360",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mongodb",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MongoDB-EI_IE433703.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MongoDB-Reviews-E433703.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/MongoDB-Jobs-E433703.htm",
				Jobs:        "489",
				Reviews:     "2.3K",
				Salaries:    "3.8K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Mongodb",
			},
			OttaProfileSlug:   "MongoDB",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 36,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Engineer, Cluster-to-Cluster Sync",
							ShortDescription: "Identify, design, implement, test, and support new features in a relatively new and quickly evolving Go codebase",
							URL:              "https://www.linkedin.com/jobs/view/4099667362/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | FerretDB
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FerretDB",
			Website: "https://www.ferretdb.com/",
			Careers: "https://www.ferretdb.com/about#careers",
			About:   "https://www.ferretdb.com/about",
			Blog:    "https://blog.ferretdb.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80672744,
				Alias:             "ferretdb",
				Name:              "FerretDB",
				Followers:         "1K",
				Employees:         "2-10",
				AssociatedMembers: "8",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "FerretDB",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "FerretDB",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "MongoDB compatible database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Redis
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Redis",
			Website: "https://redis.io/",
			Careers: "https://redis.io/careers/",
			About:   "https://redis.io/about/",
			Blog:    "https://redis.io/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2014725,
				Alias:             "redisinc",
				Name:              "Redis",
				Followers:         "251K",
				Employees:         "501-1K",
				AssociatedMembers: "1,135",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "redis",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Redis",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "16",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "redis",
				Employees: "580",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Redis-EI_IE928722.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Redis-Reviews-E928722.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Redis-Jobs-E928722.htm",
				Jobs:        "12",
				Reviews:     "219",
				Salaries:    "305",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Redis-2",
			},
			OttaProfileSlug:   "Redis",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4021697979/",
							Date:             mustDate("2024-10-28"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3963199017/",
							Date:                 mustDate("2024-11-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4082647102/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160394582/",
							Date:                 mustDate("2025-02-20"),
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
			ShortDescription: "Key-value database",
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
			},
		},

		// Tech | DigitalOcean
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DigitalOcean",
			Website: "https://www.digitalocean.com/",
			Careers: "https://www.digitalocean.com/careers",
			About:   "https://www.digitalocean.com/about",
			Blog:    "https://www.digitalocean.com/blog/tags/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2601253,
				Alias:             "digitalocean",
				Name:              "DigitalOcean",
				Followers:         "121K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,814",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "digitalocean",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "DigitalOcean",
				Employees:   "501 to 1,000",
				Salary:      "$58K ~ $298K a year",
				Reviews:     "142",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "digitalocean",
				Employees: "840",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DigitalOcean-EI_IE823482.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DigitalOcean-Reviews-E823482.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DigitalOcean-Jobs-E823482.htm",
				Jobs:        "251",
				Reviews:     "416",
				Salaries:    "673",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Digitalocean",
			},
			OttaProfileSlug:   "DigitalOcean",
			YouTubeChannelURL: "https://www.youtube.com/@DigitalOcean",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 122,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer II (Virtual Networking)",
							ShortDescription: "Go is a plus",
							URL:              "https://www.linkedin.com/jobs/view/4049491677/",
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
			ShortDescription: "Cloud service provider",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Canonical
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Canonical",
			Website: "https://canonical.com/",
			Careers: "https://canonical.com/careers",
			About:   "https://canonical.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                234280,
				Alias:             "canonical",
				Name:              "Canonical",
				Followers:         "480K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,631",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "canonical",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "canonical",
				Employees:   "501 to 1,000",
				Salary:      "$81K ~ $150K a year",
				Reviews:     "23",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "canonical",
				Employees: "930",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Canonical-EI_IE230560.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Canonical-Reviews-E230560.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Canonical-Jobs-E230560.htm",
				Jobs:        "215",
				Reviews:     "433",
				Salaries:    "726",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Canonical",
			},
			OttaProfileSlug:   "canonical",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 123,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "Golang is an essential language for our engineering teams, who build the systems that deliver Ubuntu to the world",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3736604544/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-25"), // mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "C, Golang Software Engineer",
							ShortDescription:     "Working on dqlite, a Raft extension for SQLite",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4140337335/",
							Date:                 mustDate("2025-02-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "Juju",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3783563353/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Observability Platform Developer — Python/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/2902876557/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Go (Golang) Software Engineer",
							ShortDescription:     "Identity Management",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3166882197/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Commercial Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3748409583/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Developer Tooling and Containers",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3696926569/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:            "System Software Engineer",
							ShortDescription: "We are building a new team to focus on the Rust programming language and its ecosystem on Ubuntu",
							URL:              "https://www.linkedin.com/jobs/view/3210061536/",
							Date:             mustDate("2024-12-04"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Ubuntu publisher",
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

		// Tech | SUSE
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SUSE",
			Website: "https://www.suse.com/",
			Careers: "https://jobs.suse.com/",
			About:   "https://www.suse.com/company/about/",
			Blog:    "https://www.suse.com/c/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1455,
				Alias:             "suse",
				Name:              "SUSE",
				Followers:         "169K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,589",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "suse",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SUSE",
				Employees:   "1,001 to 5,000",
				Salary:      "$58K ~ $271K a year",
				Reviews:     "19",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "suse",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SUSE-EI_IE466462.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SUSE-Reviews-E466462.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SUSE-Jobs-E466462.htm",
				Jobs:        "46",
				Reviews:     "639",
				Salaries:    "933",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Suse-2c6f8ada",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 77,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034618699/",
							Date:             mustDate("2024-11-28"),
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
			ShortDescription: "Secure enterprise open source solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Tech | Kong
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kong",
			Website: "https://konghq.com/",
			Careers: "https://konghq.com/company/careers",
			About:   "https://konghq.com/company/about-us",
			Blog:    "https://konghq.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                278819,
				Alias:             "konghq",
				Name:              "Kong Inc.",
				Followers:         "48K",
				Employees:         "501-1K",
				AssociatedMembers: "694",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Kong",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Kong",
				Employees:   "51 to 200",
				Salary:      "$120K ~ $225K a year",
				Reviews:     "43",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kong",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kong-EI_IE801963.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kong-Reviews-E801963.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kong-Jobs-E801963.htm",
				Jobs:        "66",
				Reviews:     "139",
				Salaries:    "244",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Kong",
			YouTubeChannelURL: "https://www.youtube.com/@KongInc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 66,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Expertise in designing and writing Restful and WebSocket APIs (in Golang)",
							URL:              "https://www.linkedin.com/jobs/view/4062919998/",
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
			ShortDescription: "Cloud-native, platform-agnostic, scalable API Gateway",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Exasol
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exasol",
			Website: "https://www.exasol.com/",
			Careers: "https://careers.exasol.com/",
			About:   "https://www.exasol.com/company/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1741694,
				Alias:             "exasol-ag",
				Name:              "Exasol",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "201",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "exasol",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "exasol",
				Employees: "220",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Exasol-EI_IE968677.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Exasol-Reviews-E968677.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Exasol-Jobs-E968677.htm",
				Jobs:        "19",
				Reviews:     "89",
				Salaries:    "127",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Exasol",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "3+ years of professional software development experience with a strong foundation in the Go",
							URL:              "https://app.welcometothejungle.com/jobs/yfqng_As",
							Date:             mustDate("2024-09-17"), // Appropriate date
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
			ShortDescription: "Analytics database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Palantir
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Palantir",
			Website: "https://www.palantir.com/",
			Careers: "https://www.palantir.com/careers/",
			About:   "https://www.palantir.com/about/",
			Blog:    "https://blog.palantir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20708,
				Alias:             "palantir-technologies",
				Name:              "Palantir Technologies",
				Followers:         "370K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,387",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "palantir",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "palantir",
				Employees:   "1,001 to 5,000",
				Salary:      "$95K ~ $300K a year",
				Reviews:     "225",
				ReviewsRate: "4.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "palantir",
				Employees: "3,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Palantir-Technologies-EI_IE236375.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Palantir-Technologies-Reviews-E236375.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Palantir-Technologies-Jobs-E236375.htm",
				Jobs:        "11",
				Reviews:     "861",
				Salaries:    "2.6K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Palantir-Technologies",
			},
			OttaProfileSlug:   "Palantir",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 64,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Substrate",
							ShortDescription: "Systems programming experience with strong proficiency in Golang, C/C++ or equivalent",
							URL:              "https://www.linkedin.com/jobs/view/3840764882/",
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
			ShortDescription: "Platform for big data analytics",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Buf
		// Template
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Buf",
			Website: "https://buf.build/",
			Careers: "https://buf.build/careers",
			About:   "",
			Blog:    "https://buf.build/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68684262,
				Alias:             "bufbuild",
				Name:              "Buf",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "49",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bufbuild",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Buf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 26,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4066312363/",
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
			ShortDescription: "Protobuf developer platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// FinTech
		// FinTech | Stripe
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stripe",
			Website: "https://stripe.com/",
			Careers: "https://stripe.com/jobs",
			About:   "",
			Blog:    "https://stripe.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2135371,
				Alias:             "stripe",
				Name:              "Stripe",
				Followers:         "946K",
				Employees:         "1K-5K",
				AssociatedMembers: "10,673",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stripe",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "stripe",
				Employees:   "1,001 to 5,000",
				Salary:      "$66K ~ $300K a year",
				Reviews:     "1,486",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "stripe",
				Employees: "4,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stripe-EI_IE671932.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stripe-Reviews-E671932.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stripe-Jobs-E671932.htm",
				Jobs:        "8",
				Reviews:     "1.1K",
				Salaries:    "3.7K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Stripe",
			},
			OttaProfileSlug:   "Stripe",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer, Developer SDKs, Golang expert",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4032599929/",
							Date:             mustDate("2024-11-28"),
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
			ShortDescription: "A financial infrastructure platform for businesses",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Wise
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wise",
			Website: "https://wise.com/",
			Careers: "https://wise.jobs/",
			About:   "https://wise.com/about/",
			Blog:    "https://medium.com/wise-engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1769571,
				Alias:             "wiseaccount",
				Name:              "Wise",
				Followers:         "413K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,511",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "transferwise",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "TransferWise",
				Employees:   "1,001 to 5,000",
				Salary:      "$90K ~ $215K a year",
				Reviews:     "75",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "wise",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wise-EI_IE637715.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wise-Reviews-E637715.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wise-Jobs-E637715.htm",
				Jobs:        "166",
				Reviews:     "2.1K",
				Salaries:    "3.8K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Wise-6",
			},
			OttaProfileSlug:   "Wise",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 37,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Platform Engineer",
							ShortDescription: "Experience coding in Golang and interacting/building API’s",
							URL:              "https://app.welcometothejungle.com/jobs/UVRcHQhe",
							Date:             mustDate("2024-09-17"), // Appropriate date
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
			ShortDescription: "FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | American Express
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "American Express",
			Website: "https://www.americanexpress.com/",
			Careers: "https://www.americanexpress.com/careers",
			About:   "https://www.americanexpress.com/en-us/company/",
			Blog:    "https://americanexpress.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1277,
				Alias:             "american-express",
				Name:              "American Express",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "81,656",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "americanexpress",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "american-express",
				Employees:   "10,000+",
				Salary:      "$31K ~ $224K a year",
				Reviews:     "1,011",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "american-express",
				Employees: "75,230",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-American-Express-EI_IE35.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/American-Express-Reviews-E35.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/American-Express-Jobs-E35.htm",
				Jobs:        "519",
				Reviews:     "20K",
				Salaries:    "34K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "American-Express",
			},
			OttaProfileSlug:   "American-Express",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3926321626/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer — Java, JavaScript, Node, Go & Python",
							ShortDescription:     "Global Tech",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4166983411/",
							Date:                 mustDate("2025-02-24"),
							WithSalary:           true, // $110.000 - $190.000 per year + bonus + benefits
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
			ShortDescription: "Bank holding",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Mastercard
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mastercard",
			Website: "https://www.mastercard.com/",
			Careers: "https://careers.mastercard.com/",
			About:   "https://www.mastercard.us/en-us/vision/who-we-are.html",
			Blog:    "https://developer.mastercard.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3015,
				Alias:             "mastercard",
				Name:              "Mastercard",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "39,600",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Mastercard",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "mastercard",
				Employees:   "10,000+",
				Salary:      "$51K ~ $281K a year",
				Reviews:     "460",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mastercard",
				Employees: "28,090",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mastercard-EI_IE3677.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mastercard-Reviews-E3677.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mastercard-Jobs-E3677.htm",
				Jobs:        "1K",
				Reviews:     "7.9K",
				Salaries:    "15K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Mastercard",
			YouTubeChannelURL: "https://www.youtube.com/@MasterCard",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3842471708/",
							Date:             mustDate("2024-07-26"),
						},
						{
							Title:                "Senior Software (Golang) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4113621612/",
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
			ShortDescription: "Payment card services corporation",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Morgan Stanley
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Morgan Stanley",
			Website: "https://www.morganstanley.com/",
			Careers: "https://www.morganstanley.com/careers/career-opportunities-search",
			About:   "https://www.morganstanley.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID: 0,
				IDs: []int{
					497017,  // Morgan Stanley
					1292145, // Graystone Consulting from Morgan Stanley
				},
				Alias:             "morgan-stanley",
				Name:              "Morgan Stanley",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "88,989",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MorganStanley",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Morgan-Stanley",
				Employees:   "10,000+",
				Salary:      "$30K ~ $275K a year",
				Reviews:     "669",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "morgan-stanley",
				Employees: "88,940",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Morgan-Stanley-EI_IE2282.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Morgan-Stanley-Reviews-E2282.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Morgan-Stanley-Jobs-E2282.htm",
				Jobs:        "17",
				Reviews:     "22K",
				Salaries:    "45K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Morgan-Stanley",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Python/Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3952881485/",
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
			},
		},

		// FinTech | Monzo
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Monzo",
			Website: "https://monzo.com/",
			Careers: "https://monzo.com/careers",
			About:   "https://monzo.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9471107,
				Alias:             "monzo-bank",
				Name:              "Monzo Bank",
				Followers:         "524K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,352",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "monzo",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Monzo-Bank-EI_IE1557148.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Monzo-Bank-Reviews-E1557148.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Monzo-Bank-Jobs-E1557148.htm",
				Jobs:        "27",
				Reviews:     "886",
				Salaries:    "2.3K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Monzo",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 81,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/ZwnXtENr",
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
			ShortDescription: "Mobile banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Cynergy Bank
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cynergy Bank",
			Website: "https://www.cynergybank.co.uk/",
			Careers: "https://www.cynergybank.co.uk/about-us/careers",
			About:   "https://www.cynergybank.co.uk/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18921842,
				Alias:             "cynergy-bank",
				Name:              "Cynergy Bank",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "363",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynergy-Bank-EI_IE769090.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cynergy-Bank-Reviews-E769090.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cynergy-Bank-Jobs-E769090.htm",
				Jobs:        "5",
				Reviews:     "112",
				Salaries:    "124",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "", // NOP
			YouTubeChannelURL: "https://www.youtube.com/@cynergybank",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		// FinTech | Atom bank
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Atom bank",
			Website: "https://www.atombank.co.uk/",
			Careers: "https://www.atombank.co.uk/careers/",
			About:   "https://www.atombank.co.uk/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5197064,
				Alias:             "atom-bank",
				Name:              "Atom bank",
				Followers:         "62K",
				Employees:         "501-1K",
				AssociatedMembers: "528",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "atombank",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Atom-Bank-EI_IE1354887.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Atom-Bank-Reviews-E1354887.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Atom-Bank-Jobs-E1354887.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "237",
				ReviewsRate: "4.5",
				Verified:    true,
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		// FinTech | Citi
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Citi",
			Website: "https://www.citi.com/",
			Careers: "https://jobs.citi.com/",
			About:   "https://www.citigroup.com/global/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11448,
				Alias:             "citi",
				Name:              "Citi",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "189,284",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Citi",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Citi",
				Employees:   "10,000+",
				Salary:      "$54K ~ $283K a year",
				Reviews:     "640",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "citi",
				Employees: "225,150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Citi-EI_IE8843.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Citi-Reviews-E8843.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Citi-Jobs-E8843.htm",
				Jobs:        "4.8K",
				Reviews:     "38K",
				Salaries:    "68K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Citibank-(citi)",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Citi",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang, Kubernetes)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912849197/",
							Date:             mustDate("2024-04-25"),
						},
						{
							Title:                "Java/Golang Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Preferably Go but will consider exceptional Java/Spring Boot engineers",
							URL:                  "https://www.linkedin.com/jobs/view/4089424578/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principle Generative AI Software Engineer",
							ShortDescription:     "Golang, Kubernetes",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132150228/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go/Java Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4088795722/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Generative AI Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134704785/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go/Java Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4089242986/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go/Java Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161097569/",
							Date:                 mustDate("2025-02-21"),
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
							Title:                "Senior Big Data Developer",
							ShortDescription:     "Hadoop, Spark, Scala, Python, Java",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4131052408/",
							Date:                 mustDate("2025-01-22"),
							WithSalary:           false,
							Remote:               false,
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

		// Internet
		// Internet Bitly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitly",
			Website: "https://bitly.com/",
			Careers: "https://bitly.com/pages/careers",
			About:   "https://bitly.com/pages/about",
			Blog:    "https://bitly.com/blog/why-we-write-everything-in-go/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                552285,
				Alias:             "bitly",
				Name:              "Bitly",
				Followers:         "37K",
				Employees:         "201-500",
				AssociatedMembers: "381",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bitly",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bitly-EI_IE800313.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bitly-Reviews-E800313.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bitly-Jobs-E800313.htm",
				Jobs:        "1",
				Reviews:     "94",
				Salaries:    "223",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Bitly",
			YouTubeChannelURL: "",
			GoMainLanguage:    true, // https://bitly.com/blog/why-we-write-everything-in-go/
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "URL shortening service and a link management platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},

		// Internet Cloudflare
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cloudflare",
			Website: "https://www.cloudflare.com/",
			Careers: "https://www.cloudflare.com/careers/",
			About:   "https://www.cloudflare.com/about-overview/",
			Blog:    "https://blog.cloudflare.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                407222,
				Alias:             "cloudflare",
				Name:              "Cloudflare",
				Followers:         "1M",
				Employees:         "1K-5K",
				AssociatedMembers: "4,969",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cloudflare",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cloudflare",
				Employees:   "1,001 to 5,000",
				Salary:      "$66K ~ $293K a year",
				Reviews:     "316",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudflare",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudflare-EI_IE430862.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudflare-Reviews-E430862.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cloudflare-Jobs-E430862.htm",
				Jobs:        "79",
				Reviews:     "849",
				Salaries:    "2.2K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cloudflare-6",
			},
			OttaProfileSlug:   "Cloudflare-1",
			YouTubeChannelURL: "https://www.youtube.com/@cloudflare",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 98,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Distributed Systems (Go and/or Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4038005635/",
							Date:             mustDate("2024-12-05"),
						},
						{
							Title:                "Software Engineer, Distributed Systems (Go and/or Rust)",
							ShortDescription:     "Confidence to work in multiple programming languages — bonus points for Go and/or Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077533479/",
							Date:                 mustDate("2025-01-30"), // mustDate("2025-01-03"),
							WithSalary:           true,                   // For Colorado-based hires: Estimated annual salary of $168,000 - $206,000
							Remote:               false,
						},
						{
							Title:            "Software Engineer, Distributed Systems (Go and/or Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4038010253/",
							Date:             mustDate("2025-02-15"), // mustDate("2025-01-20"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 58,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Data Center Networking (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4041774900/",
							Date:             mustDate("2024-11-25"),
						},
						{
							Title:                "Software Engineer, Distributed Systems (Go and/or Rust)",
							ShortDescription:     "Confidence to work in multiple programming languages — bonus points for Go and/or Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077533479/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           true, // For Colorado-based hires: Estimated annual salary of $168,000 - $206,000
							Remote:               false,
						},
						{
							Title:                "Software Engineer, Distributed Systems (Go and/or Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4038010253/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Rust)",
							ShortDescription:     "Data Loss Prevention",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4079142238/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           true, // $137.000 - $240.000 per year
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
			ShortDescription: "Web application security and performance",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Internet Namecheap
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Namecheap",
			Website: "https://www.namecheap.com/",
			Careers: "https://www.namecheap.com/careers/",
			About:   "https://www.namecheap.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                486932,
				Alias:             "namecheap-inc",
				Name:              "Namecheap, Inc",
				Followers:         "60K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,563",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "namecheap",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Namecheap",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "namecheap",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Namecheap-EI_IE994113.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Namecheap-Reviews-E994113.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Namecheap-Jobs-E994113.htm",
				Jobs:        "50",
				Reviews:     "209",
				Salaries:    "208",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Namecheap",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@namecheap",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Cloud Software Engineer",
							ShortDescription: "Hands-on experience with multiple programming languages (Python, Go are a plus)",
							URL:              "https://www.linkedin.com/jobs/view/3902607066/",
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
			ShortDescription: "Domain and web-hosting company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Internet GoDaddy
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GoDaddy",
			Website: "https://www.godaddy.com/",
			Careers: "https://careers.godaddy/",
			About:   "https://aboutus.godaddy.net/about-us/overview/default.aspx",
			Blog:    "https://www.godaddy.com/resources/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7846,
				Alias:             "godaddy",
				Name:              "GoDaddy",
				Followers:         "139K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,227",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "godaddy",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "godaddy",
				Employees:   "5,001 to 10,000",
				Salary:      "$28K ~ $230K a year",
				Reviews:     "305",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "godaddy",
				Employees: "7,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GoDaddy-EI_IE35337.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GoDaddy-Reviews-E35337.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GoDaddy-Jobs-E35337.htm",
				Jobs:        "61",
				Reviews:     "3.3K",
				Salaries:    "5.5K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Godaddy",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Backend (PHP/Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4011177772/",
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
			ShortDescription: "Domain and web-hosting company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Security
		// Security 1Password
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1Password",
			Website: "https://1password.com/",
			Careers: "https://1password.com/careers",
			About:   "https://1password.com/company",
			Blog:    "https://blog.1password.com/categories/developers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18648301,
				Alias:             "1password",
				Name:              "1Password",
				Followers:         "36K+",
				Employees:         "1K-5K",
				AssociatedMembers: "2,450",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "1Password",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "1Password",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "21",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "1password",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1Password-EI_IE2984143.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1Password-Reviews-E2984143.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/1Password-Jobs-E2984143.htm",
				Jobs:        "48",
				Reviews:     "316",
				Salaries:    "631",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "1password",
			},
			OttaProfileSlug:   "1Password",
			YouTubeChannelURL: "https://www.youtube.com/@1PasswordVideos",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer, Item Management (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3905310871/",
							Date:             mustDate("2024-04-25"),
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
			ShortDescription: "Service for users to store various passwords",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},

		// Security Okta
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Okta",
			Website: "https://www.okta.com/",
			Careers: "https://www.okta.com/company/careers/",
			About:   "https://www.okta.com/company/",
			Blog:    "https://developer.okta.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                926041,
				Alias:             "okta-inc-",
				Name:              "Okta",
				Followers:         "446K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,097",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "okta",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Okta",
				Employees:   "1,001 to 5,000",
				Salary:      "$48K ~ $300K a year",
				Reviews:     "733",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "okta",
				Employees: "4,620",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Okta-EI_IE444756.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Okta-Reviews-E444756.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Okta-Jobs-E444756.htm",
				Jobs:        "385",
				Reviews:     "1.5K",
				Salaries:    "3.5K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Okta",
			},
			OttaProfileSlug:   "Okta",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go/Terraform)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4090884682/",
							Date:                 mustDate("2024-12-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133471372/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           true, // $135,000—$203,000 CAD per year
							Remote:               true,
						},
						{
							Title:                "Senior Go Backend Engineer",
							ShortDescription:     "PAM (Agents/Clients)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4135551225/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true, // $118,000—$178,000 CAD per year
							Remote:               true,
						},
						{
							Title:                "Staff Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4164184734/",
							Date:                 mustDate("2025-02-26"),
							WithSalary:           true, // $139,000—$209,000 CAD per year
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
			ShortDescription: "Identity and access management company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security Nord Security
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nord Security",
			Website: "https://nordsecurity.com/",
			Careers: "https://nordsecurity.com/careers",
			About:   "https://nordsecurity.com/about-us",
			Blog:    "https://nordsecurity.com/blog/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64277258,
				Alias:             "nordsecurity",
				Name:              "Nord Security",
				Followers:         "53K+",
				Employees:         "1K-5K",
				AssociatedMembers: "1,573",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "NordSecurity",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nord-security",
				Employees: "990",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nord-Security-EI_IE4015819.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nord-Security-Reviews-E4015819.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nord-Security-Jobs-E4015819.htm",
				Jobs:        "",
				Reviews:     "229",
				Salaries:    "341",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Nord-Security",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@nordsecurity",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go) (B2B Checkout)",
							ShortDescription:     "Experience working with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4074033527/",
							Date:                 mustDate("2024-11-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Mid-Senior Backend Engineer (Golang, AI)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132260341/",
							Date:                 mustDate("2025-01-22"),
							WithSalary:           true, // 3500 — 7200 eur per month
							Remote:               false,
						},
						{
							Title:                "Middle Backend Engineer Go",
							ShortDescription:     "Pricing team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149651027/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // 3500 — 6000 EUR per month
							Remote:               false,
						},
						{
							Title:                "Mid-Senior Go Backend Engineer",
							ShortDescription:     "NordProtect",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152291366/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // Gross Salary 19800 - 33400 PLN/Month
							Remote:               false,
						},
						{
							Title:                "Mid-Senior Go Backend Engineer",
							ShortDescription:     "NordPayments team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152273092/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // Gross Salary 19800 - 33400 PLN/Month
							Remote:               true,
						},
						{
							Title:                "Senior Backend engineer PHP / Go",
							ShortDescription:     "Business enablement",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158335201/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           true, // Gross Salary 3,500 - 7,200 EUR per month
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
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security Proton
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Proton",
			Website: "https://proton.me/",
			Careers: "https://proton.me/careers",
			About:   "https://proton.me/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5241679,
				Alias: "protonprivacy",
				Name:  "Proton",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ProtonMail",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Proton",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "3",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "proton",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proton-privacy-by-default-EI_IE1405328.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proton-privacy-by-default-Reviews-E1405328.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Proton-privacy-by-default-Jobs-E1405328.htm",
				Jobs:        "16",
				Reviews:     "114",
				Salaries:    "187",
				ReviewsRate: "4.2",
				Verified:    true,
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
							URL:              "https://www.linkedin.com/jobs/view/4108568196/",
							Date:             mustDate("2025-02-11"), // mustDate("2024-01-16"), // mustDate("2024-12-23"),
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
		},

		// Security Fortinet
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fortinet",
			Website: "https://www.fortinet.com/",
			Careers: "https://www.fortinet.com/corporate/careers",
			About:   "https://www.fortinet.com/corporate/about-us/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6460,
				Alias:             "fortinet",
				Name:              "Fortinet",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "14,428",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fortinet",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Fortinet",
				Employees:   "5,001 to 10,000",
				Salary:      "$27K ~ $647K a year",
				Reviews:     "215",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fortinet",
				Employees: "8,240",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fortinet-EI_IE23128.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fortinet-Reviews-E23128.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fortinet-Jobs-E23128.htm",
				Jobs:        "968",
				Reviews:     "2.9K",
				Salaries:    "5.1K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Fortinet",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fortinet",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer (Golang and Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3907970778/",
							Date:             mustDate("2024-08-23"),
						},
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4104785473/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:                "Software Developer (Golang)",
							ShortDescription:     "Experience with Golang and Python and software development methodology",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122351024/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false, // Canada base salary range for this full-time position is expected to be between $81,000 - $110,000 annually
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
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security SentinelOne
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SentinelOne",
			Website: "https://www.sentinelone.com/",
			Careers: "https://www.sentinelone.com/careers/",
			About:   "https://sentinelone.com/company/",
			Blog:    "https://www.sentinelone.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2886771,
				Alias:             "sentinelone",
				Name:              "SentinelOne",
				Followers:         "253K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,799",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Sentinel-One",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SentinelOne",
				Employees:   "1,001 to 5,000",
				Salary:      "$100K ~ $189K a year",
				Reviews:     "27",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sentinelone",
				Employees: "1,120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SentinelOne-EI_IE1361978.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SentinelOne-Reviews-E1361978.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SentinelOne-Jobs-E1361978.htm",
				Jobs:        "115",
				Reviews:     "670",
				Salaries:    "658",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sentinelone",
			},
			OttaProfileSlug:   "SentinelOne",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4096029115/",
							Date:                 mustDate("2024-12-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "3+ years of software engineering experience in at least one high-level programming language (Go, Java or Python)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125792094/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "Cloud Scanner Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125792093/",
							Date:                 mustDate("2025-02-11"),
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

		// Famous
		// Famous Uber
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Uber",
			Website: "https://www.uber.com/",
			Careers: "https://www.uber.com//careers/",
			About:   "https://www.uber.com/about/",
			Blog:    "https://www.uber.com/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1815218,
				Alias:             "uber-com",
				Name:              "Uber",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "108,744",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "uber",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Uber",
				Employees:   "10,000+",
				Salary:      "$34K ~ $350K a year",
				Reviews:     "2,735",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "uber",
				Employees: "26,900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Uber-EI_IE575263.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Uber-Reviews-E575263.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Uber-Jobs-E575263.htm",
				Jobs:        "2.1K",
				Reviews:     "33K",
				Salaries:    "58K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Uber",
			},
			OttaProfileSlug:   "Uber",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 33,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — UberEats Order Platform",
							ShortDescription: "Highly efficient coding in Golang, Java or any similar languages",
							URL:              "https://www.linkedin.com/jobs/view/4093231440/",
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
			ShortDescription: "Transportation company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Siemens
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Siemens",
			Website: "https://www.siemens.com/",
			Careers: "https://www.siemens.com/global/en/company/jobs.html",
			About:   "https://www.siemens.com/global/en/company/about.html",
			Blog:    "https://blogs.sw.siemens.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				IDs:               []int{1043, 157241},
				Alias:             "siemens",
				Name:              "Siemens",
				Followers:         "7M",
				Employees:         "10K+",
				AssociatedMembers: "225,230",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "siemens",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Siemens",
				Employees:   "10,000+",
				Salary:      "$32K ~ $235K a year",
				Reviews:     "279",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "siemens",
				Employees: "253,720",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Siemens-EI_IE3510.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Siemens-Reviews-E3510.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Siemens-Jobs-E3510.htm",
				Jobs:        "589",
				Reviews:     "24K",
				Salaries:    "40K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Siemens",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@SiemensKnowledgeHub",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 19,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Developer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103717982/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4044911988/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4129459999/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4018413094/",
							Date:             mustDate("2024-09-25"),
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

		// Famous Ericsson
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ericsson",
			Website: "https://www.ericsson.com/",
			Careers: "https://www.ericsson.com/careers",
			About:   "https://www.ericsson.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1060,
				Alias:             "ericsson",
				Name:              "Ericsson",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "107,578",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Ericsson",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "ericsson",
				Employees:   "10,000+",
				Salary:      "$38K ~ $228K a year",
				Reviews:     "223",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ericsson",
				Employees: "131,940",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ericsson-Worldwide-EI_IE3472.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ericsson-Worldwide-Reviews-E3472.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ericsson-Worldwide-Jobs-E3472.htm",
				Jobs:        "2.1K",
				Reviews:     "20K",
				Salaries:    "27K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Ericsson--worldwide-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Developer",
							ShortDescription:     "Advanced development experience in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4038208151/",
							Date:                 mustDate("2024-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Python/Golang Development Engineer",
							ShortDescription:     "Design, develop, deploy and maintain a suite of Cloud components/microservices written in Golang, Java, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4098448765/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4030181056/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           false,
							Remote:               false,
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
							Title:                "Erlang/Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149110859/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Famous SoundCloud
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SoundCloud",
			Website: "https://soundcloud.com/",
			Careers: "https://careers.soundcloud.com/",
			About:   "https://press.soundcloud.com/about/",
			Blog:    "https://developers.soundcloud.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                200200,
				Alias:             "soundcloud",
				Name:              "SoundCloud",
				Followers:         "255K",
				Employees:         "201-500",
				AssociatedMembers: "582",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "soundcloud",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SoundCloud",
				Employees:   "201 to 500",
				Salary:      "$70K ~ $170K a year",
				Reviews:     "34",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "soundcloud",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SoundCloud-EI_IE407066.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SoundCloud-Reviews-E407066.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SoundCloud-Jobs-E407066.htm",
				Jobs:        "12",
				Reviews:     "224",
				Salaries:    "456",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Soundcloud",
			},
			OttaProfileSlug:   "SoundCloud",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Media Streaming Backend Engineer",
							ShortDescription: "Proficiency in Go is essential, and experience with additional programming languages such as Scala is highly advantageous",
							URL:              "https://app.welcometothejungle.com/jobs/Q8WY5QwP",
							Date:             mustDate("2024-12-27"), // Appropriate date
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Media Streaming Backend Engineer",
							ShortDescription: "Proficiency in Go is essential, and experience with additional programming languages such as Scala is highly advantageous",
							URL:              "https://app.welcometothejungle.com/jobs/Q8WY5QwP",
							Date:             mustDate("2024-12-27"), // Appropriate date
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Audio streaming service",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Spotify
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Spotify",
			Website: "https://www.lifeatspotify.com/",
			Careers: "https://www.lifeatspotify.com/jobs",
			About:   "https://newsroom.spotify.com/company-info/",
			Blog:    "https://engineering.atspotify.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                207470,
				Alias:             "spotify",
				Name:              "Spotify",
				Followers:         "4M",
				Employees:         "5K-10K",
				AssociatedMembers: "15,151",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "spotify",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "spotify",
				Employees:   "1,001 to 5,000",
				Salary:      "$56K ~ $500K a year",
				Reviews:     "825",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "spotify",
				Employees: "9,630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Spotify-EI_IE408251.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spotify-Reviews-E408251.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Spotify-Jobs-E408251.htm",
				Jobs:        "76",
				Reviews:     "1.9K",
				Salaries:    "5.4K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Spotify",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer — Advertising",
							ShortDescription: "Ideally you have worked with Java, and an additional language such as Go or C++ is a bonus",
							URL:              "https://www.linkedin.com/jobs/view/4035153021/", // "https://app.welcometothejungle.com/jobs/BAoL6B1_",
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
			ShortDescription: "Music streaming service",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous DoorDash
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DoorDash",
			Website: "https://doordash.com/",
			Careers: "https://careersatdoordash.com/",
			About:   "https://about.doordash.com/en-us/company",
			Blog:    "https://doordash.engineering/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3205573,
				Alias:             "doordash",
				Name:              "DoorDash",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "62,210",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "doordash",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "doordash",
				Employees:   "5,001 to 10,000",
				Salary:      "$50K ~ $283K a year",
				Reviews:     "1,040",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "doordash",
				Employees: "28,800",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DoorDash-EI_IE813073.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DoorDash-Reviews-E813073.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DoorDash-Jobs-E813073.htm",
				Jobs:        "499",
				Reviews:     "17K",
				Salaries:    "15K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Doordash",
			},
			OttaProfileSlug:   "DoorDash",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Observability",
							ShortDescription: "Development using Golang, Kotlin, and/or Python",
							URL:              "https://www.linkedin.com/jobs/view/4081729069/",
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
			ShortDescription: "Food delivery",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Just Eat Takeaway.com",
			Website: "https://www.justeattakeaway.com/",
			Careers: "https://careers.justeattakeaway.com/global/en",
			About:   "https://www.justeattakeaway.com/about/our-story/default.aspx",
			Blog:    "https://medium.com/justeattakeaway-tech",
			LinkedInProfile: domain.LinkedInProfile{
				ID: 0,
				IDs: []int{
					103531,  // Just Eat Takeaway.com
					2708300, //  SkipTheDishes
				},
				Alias:             "just-eat-takeaway-com",
				Name:              "Just Eat Takeaway.com",
				Followers:         "246K",
				Employees:         "10K+",
				AssociatedMembers: "12,411",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "justeattakeaway",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Just-Eat-Takeaway",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "15",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "just-eat",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Just-Eat-Takeaway-com-EI_IE490124.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Just-Eat-Takeaway-com-Reviews-E490124.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Just-Eat-Takeaway-com-Jobs-E490124.htm",
				Jobs:        "277",
				Reviews:     "2.8K",
				Salaries:    "4.9K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Just-Eat-Takeaway.com",
			},
			OttaProfileSlug:   "Just-Eat",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4039505901/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4117868597/",
							Date:             mustDate("2025-01-09"),
						},
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133479761/",
							Date:                 mustDate("2025-01-25"),
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
			ShortDescription: "Food delivery",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Famous Sixt
		{
			ID:      0,
			Type:    "",
			Name:    "Sixt",
			Website: "https://www.sixt.com/",
			Careers: "https://www.sixt.jobs/",
			About:   "https://about.sixt.com/",
			Blog:    "https://www.sixt.tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17120,
				Alias:             "sixt",
				PreviousAliases:   nil,
				Name:              "SIXT",
				Followers:         "136K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,413",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Sixt",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Sixt",
				Employees:   "5,001 to 10,000",
				Salary:      "$30K ~ $134K a year",
				Reviews:     "50",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sixt",
				Employees: "5,550",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sixt-EI_IE10875.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sixt-Reviews-E10875.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sixt-Jobs-E10875.htm",
				Jobs:        "856",
				Reviews:     "1.8K",
				Salaries:    "2.7K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sixt-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Development Engineer III (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085642905/",
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
			ShortDescription: "Car Rental service",
			Industries:       []domain.Industry{},
		},

		// Famous Motorola Solutions
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Motorola Solutions",
			Website: "https://www.motorolasolutions.com/",
			Careers: "https://www.motorolasolutions.com/en_us/about/careers.html",
			About:   "https://www.motorolasolutions.com/en_us/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1771432,
				Alias:             "motorolasolutions",
				Name:              "Motorola Solutions",
				Followers:         "592K",
				Employees:         "10K+",
				AssociatedMembers: "24,454",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Motorola-Mobility",
				Employees:   "1,001 to 5,000",
				Salary:      "$37K ~ $216K a year",
				Reviews:     "116",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "motorola-solutions",
				Employees: "19,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Motorola-Solutions-EI_IE427189.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Motorola-Solutions-Reviews-E427189.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Motorola-Solutions-Jobs-E427189.htm",
				Jobs:        "561",
				Reviews:     "5.1K",
				Salaries:    "9K",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Motorola-Solutions",
			},
			OttaProfileSlug:   "Motorola-Solutions",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Software Engineer",
							ShortDescription: "Knowledge of cloud platforms such as AWS, Azure, or Google Cloud",
							URL:              "https://www.linkedin.com/jobs/view/4082840850/",
							Date:             mustDate("2025-02-04"), // mustDate("2025-01-07"), // mustDate("2024-12-12"),
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
			ShortDescription: "American multinational telecommunications company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Samsung
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Samsung",
			Website: "https://www.samsung.com/",
			Careers: "https://www.samsung.com/us/careers/",
			About:   "https://www.samsung.com/us/about-us/our-business/",
			Blog:    "https://developer.samsung.com/blog/en",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				IDs:               []int{1753, 13749561, 3270132, 77752302, 3812597, 3220515, 10633911, 1447520, 3238801, 2711168, 10261221, 87464967, 5089912, 9500165, 10218505, 895705, 78467539, 10332849, 10660176, 79815984, 3290134, 76958044, 85881048, 27127559, 9278177, 37470202, 22316801, 81590082, 11229641, 15213487, 68478415, 5552815, 14472582},
				Alias:             "samsung-electronics",
				Name:              "Samsung Electronics",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "139,362",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Samsung",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Samsung",
				Employees:   "10,000+",
				Salary:      "$22K ~ $283K a year",
				Reviews:     "868",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "samsung",
				Employees: "145,470",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Samsung-Electronics-EI_IE3363.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Samsung-Electronics-Reviews-E3363.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Samsung-Electronics-Jobs-E3363.htm",
				Jobs:        "214",
				Reviews:     "17K",
				Salaries:    "23K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Samsung-Electronics-9",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Samsung",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Engineer, Ad Serving — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3919334242/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:                "Golang Engineer",
							ShortDescription:     "Samsung TV Plus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139319243/",
							Date:                 mustDate("2025-01-31"),
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
			ShortDescription: "Electronics corporation",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Salesforge
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Salesforge",
			Website: "https://www.salesforge.ai/",
			Careers: "",
			About:   "https://www.salesforge.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91706116,
				Alias:             "salesforgeai",
				Name:              "Salesforge",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "34",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
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
			YouTubeChannelURL: "https://www.youtube.com/@salesforge",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Backend Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912551615/",
							Date:             mustDate("2024-04-25"),
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
			ShortDescription: "All-inclusive sales app",
			Industries:       []domain.Industry{},
		},

		// Some
		// Some | Careem
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Careem",
			Website: "https://www.careem.com/",
			Careers: "https://jobs.careem.com/",
			About:   "https://www.careem.com/en-AE/about-us/",
			Blog:    "https://engineering.careem.com/tech/categories/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2852511,
				Alias:             "careem",
				Name:              "Careem",
				Followers:         "500K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,759",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "careem",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Careem-EI_IE1438731.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Careem-Reviews-E1438731.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Careem-Jobs-E1438731.htm",
				Jobs:        "58",
				Reviews:     "822",
				Salaries:    "915",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Careem",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
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
			ShortDescription: "SuperApp",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Dailymotion
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dailymotion",
			Website: "https://www.dailymotion.com/",
			Careers: "https://careers.dailymotion.com/jobs/",
			About:   "https://about.dailymotion.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24411,
				Alias:             "dailymotion",
				Name:              "Dailymotion",
				Followers:         "84K",
				Employees:         "201-500",
				AssociatedMembers: "574",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "dailymotion",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dailymotion-EI_IE372880.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dailymotion-Reviews-E372880.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dailymotion-Jobs-E372880.htm",
				Jobs:        "22",
				Reviews:     "236",
				Salaries:    "357",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Dailymotion",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
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
			ShortDescription: "Video hosting and sharing platforms",
			Industries:       []domain.Industry{},
		},

		// Some | Stream
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stream",
			Website: "https://getstream.io/",
			Careers: "https://getstream.io/team/#jobs",
			About:   "https://getstream.io/team/",
			Blog:    "https://getstream.io/blog/topic/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5338728,
				Alias:             "getstream",
				Name:              "Stream",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "275",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "GetStream",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Stream",
				Employees:   "51-200",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "getstreamio",
				Employees: "90",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stream-CO-EI_IE1703813.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stream-CO-Reviews-E1703813.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stream-CO-Jobs-E1703813.htm",
				Jobs:        "",
				Reviews:     "56",
				Salaries:    "96",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Stream",
			YouTubeChannelURL: "https://www.youtube.com/@streamdevelopers",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 40,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3873637399/",
							Date:                 mustDate("2024-07-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Software Engineer (Golang)",
							ShortDescription:     "In addition to Go we use CockroachDB/Postgres, RocksDB, Raft and Redis",
							SwitchingOpportunity: "We are willing to train you on Go if you’re experienced on a different tech stack (we have a 10-week internal onboarding program focused on Go, scalability etc.)",
							URL:                  "https://www.linkedin.com/jobs/view/4092131109/",
							Date:                 mustDate("2025-02-16"), // mustDate("2025-01-26"), // mustDate("2025-01-03"),
							WithSalary:           true,                   // €70,000 to €160,000 EUR
							Remote:               false,
						},
						{
							Title:                "Senior Full-stack Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092129283/",
							Date:                 mustDate("2025-02-18"),
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
			ShortDescription: "Scalable and fast APIs for building social networks and apps",
			Industries:       []domain.Industry{},
		},

		// Some | Workato
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Workato",
			Website: "https://www.workato.com/",
			Careers: "https://www.workato.com/careers",
			About:   "https://www.workato.com/about_us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3675685,
				Alias:             "workato",
				Name:              "Workato",
				Followers:         "68K",
				Employees:         "501-1K",
				AssociatedMembers: "1,051",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "workato",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Workato-EI_IE1333040.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Workato-Reviews-E1333040.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Workato-Jobs-E1333040.htm",
				Jobs:        "93",
				Reviews:     "366",
				Salaries:    "610",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Workato",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Form3
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Form3",
			Website: "https://www.form3.tech/",
			Careers: "https://www.form3.tech/careers",
			About:   "https://www.form3.tech/about/our-story",
			Blog:    "https://www.form3.tech/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15156804,
				Alias:             "form3-financial-cloud",
				Name:              "Form3",
				Followers:         "23K",
				Employees:         "501-1K",
				AssociatedMembers: "359",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "form3tech-oss",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Form3",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "form3",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Form3-EI_IE2008415.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Form3-Reviews-E2008415.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Form3-Jobs-E2008415.htm",
				Jobs:        "4",
				Reviews:     "153",
				Salaries:    "294",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Form3",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 81,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "Foundation API",
							URL:              "https://www.linkedin.com/jobs/view/4081153248/",
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
			ShortDescription: "Financial technology company focused on money transfers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Assertive Yield
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Assertive Yield",
			Website: "https://www.assertiveyield.com/",
			Careers: "https://www.assertiveyield.com/careers/",
			About:   "https://www.assertiveyield.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76806664,
				Alias:             "assertive-yield",
				Name:              "Assertive Yield B.V.",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "32",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Assertive-Yield",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
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
			ShortDescription: "Marketing services company",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
				domain.IndustryMarTech,
			},
		},

		// Some | Splunk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Splunk",
			Website: "https://www.splunk.com/",
			Careers: "https://www.splunk.com/en_us/careers.html",
			About:   "https://www.splunk.com/en_us/about-splunk.html",
			Blog:    "https://www.splunk.com/en_us/blog/learn.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20226,
				Alias:             "splunk",
				Name:              "Splunk",
				Followers:         "697K",
				Employees:         "5K-10K",
				AssociatedMembers: "10,195",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "splunk",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "splunk",
				Employees:   "5,001 to 10,000",
				Salary:      "$78K ~ $340K a year",
				Reviews:     "882",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "splunk",
				Employees: "8,140",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splunk-EI_IE117313.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splunk-Reviews-E117313.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Splunk-Jobs-E117313.htm",
				Jobs:        "1K",
				Reviews:     "2K",
				Salaries:    "6.2K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Splunk",
			},
			OttaProfileSlug:   "Splunk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 49,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Software Engineer — Backend (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056824974/",
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
			ShortDescription: "Big Data platform",
			Industries: []domain.Industry{
				domain.IndustryBigData,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | 90POE
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "90POE",
			Website: "https://www.90poe.io/",
			Careers: "",
			About:   "https://www.90poe.io/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18466590,
				Alias:             "90poe",
				Name:              "Ninety Percent of Everything (90POE)",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "187",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "90poe",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-90-POE-EI_IE3898368.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/90-POE-Reviews-E3898368.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/90-POE-Jobs-E3898368.htm",
				Jobs:        "",
				Reviews:     "21",
				Salaries:    "24",
				ReviewsRate: "4.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 31,
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
			ShortDescription: "Digital platform for maritime transporting",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | HelloFresh
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HelloFresh",
			Website: "https://www.hellofresh.com/",
			Careers: "https://careers.hellofresh.com/global/en",
			About:   "https://www.hellofresh.com/about/about-us",
			Blog:    "https://engineering.hellofresh.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2454643,
				Alias:             "hellofresh",
				Name:              "HelloFresh",
				Followers:         "380K",
				Employees:         "10K+",
				AssociatedMembers: "15,625",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hellofresh",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "hellofresh",
				Employees:   "5,001 to 10,000",
				Salary:      "$65K ~ $185K a year",
				Reviews:     "90",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hellofresh",
				Employees: "6,560",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HelloFresh-EI_IE998728.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HelloFresh-Reviews-E998728.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HelloFresh-Jobs-E998728.htm",
				Jobs:        "205",
				Reviews:     "3.5K",
				Salaries:    "6.4K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Hellofresh",
			},
			OttaProfileSlug:   "HelloFresh",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 32,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Engineer",
							ShortDescription: "Golang experience would be a plus",
							URL:              "https://www.linkedin.com/jobs/view/4096795323/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | AUTODOC
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AUTODOC",
			Website: "https://autodoc.group/",
			Careers: "https://autodoc.group/en/career",
			About:   "https://autodoc.group/en/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7703911,
				Alias:             "autodoc",
				Name:              "AUTODOC",
				Followers:         "61K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,148",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AUTODOC-EI_IE2179604.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AUTODOC-Reviews-E2179604.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AUTODOC-Jobs-E2179604.htm",
				Jobs:        "114",
				Reviews:     "143",
				Salaries:    "252",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Autodoc",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Developer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3757109938/",
							Date:                 mustDate("2025-01-07"), // mustDate("2024-12-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4131277355/",
							Date:                 mustDate("2025-02-18"), // mustDate("2025-01-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144895838/",
							Date:                 mustDate("2025-02-05"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Gymondo
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gymondo",
			Website: "https://www.gymondo.com/",
			Careers: "https://gymondo-gmbh.jobs.personio.com/",
			About:   "https://www.gymondo.com/en/press/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5233814,
				Alias:             "gymondo-gmbh",
				Name:              "Gymondo",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "102",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Gymondo-git",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gymondo-EI_IE1344198.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gymondo-Reviews-E1344198.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gymondo-Jobs-E1344198.htm",
				Jobs:        "2",
				Reviews:     "52",
				Salaries:    "95",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Gymondo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "",
								Date:             mustDate(""),
							},
						*/
						{
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4165424189/",
							Date:                 mustDate("2025-02-26"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Delivery Hero
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Delivery Hero",
			Website: "https://www.deliveryhero.com/",
			Careers: "https://careers.deliveryhero.com/",
			About:   "https://www.deliveryhero.com/about/",
			Blog:    "https://tech.deliveryhero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2393200,
				Alias:             "delivery-hero-se",
				Name:              "Delivery Hero",
				Followers:         "239K",
				Employees:         "10K+",
				AssociatedMembers: "35,000",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "deliveryhero",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Delivery-Hero",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "287",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "delivery-hero",
				Employees: "19,830",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Delivery-Hero-EI_IE504556.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Delivery-Hero-Reviews-E504556.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Delivery-Hero-Jobs-E504556.htm",
				Jobs:        "639",
				Reviews:     "1.7K",
				Salaries:    "3.6K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Delivery-Hero",
			},
			OttaProfileSlug:   "Delivery-Hero",
			YouTubeChannelURL: "https://www.youtube.com/@DeliveryHeroDH",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Junior) Software Engineer (Golang) — Demand Domain (AdTech)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4023882932/",
							Date:                 mustDate("2024-11-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "Quick Commerce",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148753946/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Software Engineer, Golang",
							ShortDescription:     "Quick Commerce",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148759219/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "Quick Commerce",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4136211696/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "Logistics, Global Service Tech",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139118740/",
							Date:                 mustDate("2025-02-21"),
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
			ShortDescription: "Delivery platform",
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

		// Some | Weaviate
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Weaviate",
			Website: "https://weaviate.io/",
			Careers: "https://weaviate.io/company/careers",
			About:   "https://weaviate.io/company/about-us",
			Blog:    "https://weaviate.io/blog/tags/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11702022,
				Alias:             "weaviate-io",
				Name:              "Weaviate",
				Followers:         "26K",
				Employees:         "51-200",
				AssociatedMembers: "103",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "weaviate",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Weaviate-EI_IE7479527.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Weaviate-Reviews-E7479527.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Weaviate-Jobs-E7479527.htm",
				Jobs:        "",
				Reviews:     "9",
				Salaries:    "4",
				ReviewsRate: "5.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Weaviate",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Fubo
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fubo",
			Website: "https://www.fubo.tv/",
			Careers: "https://careers.fubo.tv/",
			About:   "https://corporate.fubo.tv/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5316737,
				Alias:             "fubotv",
				Name:              "Fubo",
				Followers:         "19K",
				Employees:         "501-1K",
				AssociatedMembers: "682",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fubotv",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "fuboTV",
				Employees:   "201 to 500",
				Salary:      "$115K ~ $245K a year",
				Reviews:     "25",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fubotv",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FuboTV-EI_IE1006878.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FuboTV-Reviews-E1006878.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FuboTV-Jobs-E1006878.htm",
				Jobs:        "12",
				Reviews:     "72",
				Salaries:    "194",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Fubotv",
			},
			OttaProfileSlug:   "fubo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{

						{
							Title:            "Senior Software Engineer, Backend — Ad Engineering",
							ShortDescription: "Expertise in building microservices in Golang, Java, C++ or similar",
							URL:              "https://www.linkedin.com/jobs/view/3935382943/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Yassir
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yassir",
			Website: "https://yassir.com/",
			Careers: "https://yassir.com/career",
			About:   "https://yassir.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19069709,
				Alias:             "yassir",
				Name:              "Yassir",
				Followers:         "88K",
				Employees:         "501-1K",
				AssociatedMembers: "1,519",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "YAtechnologies",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "yassir",
				Employees: "900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-YASSIR-EI_IE2601333.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/YASSIR-Reviews-E2601333.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/YASSIR-Jobs-E2601333.htm",
				Jobs:        "7",
				Reviews:     "182",
				Salaries:    "146",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Yassir",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Backend Engineer",
							ShortDescription: "Solid backend engineering experience with Node.js and Go",
							URL:              "https://www.linkedin.com/jobs/view/4110428412/",
							Date:             mustDate("2024-12-27"),
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
		},

		// Some | Vio.com
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vio.com",
			Website: "https://www.vio.com/",
			Careers: "https://www.vio.com/careers",
			About:   "https://www.vio.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1192098,
				Alias:             "viodotcom",
				Name:              "Vio.com",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "viodotcom",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vio-com-EI_IE940798.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vio-com-Reviews-E940798.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vio-com-Jobs-E940798.htm",
				Jobs:        "5",
				Reviews:     "57",
				Salaries:    "83",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Vio-com",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3819771736/",
							Date:             mustDate("2024-04-25"),
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

		// Some | Vodeno
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vodeno",
			Website: "https://vodeno.com/",
			Careers: "https://vodeno.com/careers/",
			About:   "https://vodeno.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19016391,
				Alias:             "vodeno",
				Name:              "Vodeno",
				Followers:         "9K",
				Employees:         "201-500",
				AssociatedMembers: "221",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "vodeno",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vodeno-EI_IE2877672.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vodeno-Reviews-E2877672.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vodeno-Jobs-E2877672.htm",
				Jobs:        "1",
				Reviews:     "37",
				Salaries:    "60",
				ReviewsRate: "3.2",
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
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://justjoin.it/offers/vodeno-go-developer-kielce-358668",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Utility Warehouse
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Utility Warehouse",
			Website: "https://uw.co.uk/",
			Careers: "https://uw.co.uk/careers",
			About:   "https://uw.co.uk/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                457903,
				Alias:             "utilitywarehouse",
				Name:              "Utility Warehouse",
				Followers:         "31K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,349",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "utilitywarehouse",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Utility-Warehouse-EI_IE512935.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Utility-Warehouse-Reviews-E512935.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Utility-Warehouse-Jobs-E512935.htm",
				Jobs:        "2",
				Reviews:     "723",
				Salaries:    "1.2K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Utility-Warehouse",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 106,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Codenotary
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Codenotary",
			Website: "https://codenotary.com/",
			Careers: "https://codenotary.com/job",
			About:   "",
			Blog:    "https://codenotary.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35523736,
				Alias:             "codenotary",
				Name:              "Codenotary",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "16",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "codenotary",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CodeNotary-EI_IE3677292.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CodeNotary-Reviews-E3677292.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CodeNotary-Jobs-E3677292.htm",
				Jobs:        "2",
				Reviews:     "7",
				Salaries:    "4",
				ReviewsRate: "2.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Audigent
		{
			ID:      0,
			Type:    "",
			Name:    "Audigent",
			Website: "https://audigent.com/",
			Careers: "https://audigent.com/careers",
			About:   "https://audigent.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10642467,
				Alias: "audigent",
				Name:  "Audigent",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "AuDigent",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Audigent-EI_IE5815298.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Audigent-Reviews-E5815298.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Audigent-Jobs-E5815298.htm",
				Jobs:        "",
				Reviews:     "19",
				Salaries:    "32",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Audigent",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | runZero
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "runZero",
			Website: "https://www.runzero.com/",
			Careers: "https://www.runzero.com/about/careers/",
			About:   "https://www.runzero.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33274038,
				Alias:             "runzero",
				Name:              "runZero",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "78",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "runZeroInc",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "runzero",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-runZero-EI_IE7717209.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/runZero-Reviews-E7717209.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/runZero-Jobs-E7717209.htm",
				Jobs:        "",
				Reviews:     "31",
				Salaries:    "28",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "runZero",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085466064/",
							Date:             mustDate("2024-11-25"),
						},
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4131339080/",
							Date:             mustDate("2025-01-21"),
							Remote:           true,
							WithSalary:       true, // $160K — $190K per year
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
		},

		// Some | Tyk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tyk",
			Website: "https://tyk.io/",
			Careers: "https://tyk.io/current-vacancies/",
			About:   "https://tyk.io/company/",
			Blog:    "https://tyk.io/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10775050,
				Alias:             "tyk",
				Name:              "Tyk",
				Followers:         "31K",
				Employees:         "51-200",
				AssociatedMembers: "161",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "TykTechnologies",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tyk-EI_IE2321465.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tyk-Reviews-E2321465.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tyk-Jobs-E2321465.htm",
				Jobs:        "3",
				Reviews:     "64",
				Salaries:    "72",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Tyk",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 86,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/CwkAD14C",
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
			ShortDescription: "Open-source API gateway & API management platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | OpenTag
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenTag",
			Website: "https://theopentag.com/",
			Careers: "https://theopentag.com/careers/",
			About:   "https://theopentag.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20565935,
				Alias:             "theopentag",
				Name:              "OpenTag",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "88",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "OpenTagOS",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenTag-EI_IE3310869.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenTag-Reviews-E3310869.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OpenTag-Jobs-E3310869.htm",
				Jobs:        "",
				Reviews:     "41",
				Salaries:    "32",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Oxla
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Oxla",
			Website: "https://oxla.com/",
			Careers: "https://www.oxla.com/careers",
			About:   "",
			Blog:    "https://blog.oxla.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79378182,
				Alias:             "oxla",
				Name:              "Oxla",
				Followers:         "830",
				Employees:         "11-50",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oxla-EI_IE9210109.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oxla-Reviews-E9210109.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Oxla-Jobs-E9210109.htm",
				Jobs:        "1",
				Reviews:     "",
				Salaries:    "2",
				ReviewsRate: "",
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
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://justjoin.it/offers/oxla-golang-developer-gdansk-362959",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Lightspeed
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lightspeed",
			Website: "https://www.lightspeedhq.com/",
			Careers: "https://www.lightspeedhq.com/careers/",
			About:   "https://www.lightspeedhq.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1557218,
				Alias:             "lightspeedcommerce",
				Name:              "Lightspeed Commerce",
				Followers:         "77K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,974",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "lightspeed",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Lightspeed-POS",
				Employees:   "1,001 to 5,000",
				Salary:      "$62K ~ $250K a year",
				Reviews:     "43",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lightspeed-commerce",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lightspeed-EI_IE648762.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lightspeed-Reviews-E648762.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lightspeed-Jobs-E648762.htm",
				Jobs:        "325",
				Reviews:     "996",
				Salaries:    "2.2K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Lightspeed-5",
			},
			OttaProfileSlug:   "Lightspeed",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCqOEKwLpolZBcj4LfU3R0Fg",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Fullstack Software Developer (Golang & React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4040512029/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Fullstack Developer (Go & React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119051974/",
							Date:                 mustDate("2025-01-20"), // mustDate("2025-01-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior/Staff Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092327926/",
							Date:                 mustDate("2025-02-17"), // mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Fullstack Software Developer (Golang & React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4143662469/",
							Date:                 mustDate("2025-02-26"), // mustDate("2025-02-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160162842/",
							Date:                 mustDate("2025-02-20"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Squarespace
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Squarespace",
			Website: "https://www.squarespace.com/",
			Careers: "https://www.squarespace.com/about/careers",
			About:   "https://www.squarespace.com/about/company",
			Blog:    "https://engineering.squarespace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                265314,
				Alias:             "squarespace",
				Name:              "Squarespace",
				Followers:         "136K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,812",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "squarespace",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "squarespace",
				Employees:   "501 to 1,000",
				Salary:      "$49K ~ $262K a year",
				Reviews:     "195",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "squarespace",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Squarespace-EI_IE466343.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Squarespace-Reviews-E466343.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Squarespace-Jobs-E466343.htm",
				Jobs:        "27",
				Reviews:     "508",
				Salaries:    "1.5K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Squarespace",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/squarespace",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Go — Social",
							ShortDescription: "5+ years of Go, Java or other object-oriented language application development",
							URL:              "https://www.linkedin.com/jobs/view/4022342302/",
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Curve
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Curve",
			Website: "https://curve.com/",
			Careers: "https://www.curve.com/en/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10023464,
				Alias:             "curve-ltd",
				Name:              "Curve",
				Followers:         "59K",
				Employees:         "201-500",
				AssociatedMembers: "283",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "curve",
				Employees: "440",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Curve-EI_IE1739754.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Curve-Reviews-E1739754.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Curve-Jobs-E1739754.htm",
				Jobs:        "15",
				Reviews:     "224",
				Salaries:    "573",
				ReviewsRate: "3.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Curve-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Software Engineer (PHP)",
							ShortDescription: "The purpose of the Backend Software Engineer will be to support the build of elegant, performant, maintainable, operable and secure Golang microservices",
							URL:              "https://www.linkedin.com/jobs/view/3872933701/",
							Date:             mustDate("2024-04-25"),
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
			HasEmployeesFromCountries: []domain.Country{
				//
			},
		},

		// Some | Tradevest
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tradevest",
			Website: "https://www.tradevest.ai/",
			Careers: "https://tradevest.ai/careers",
			About:   "https://tradevest.ai/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92827682,
				Alias:             "tradevestgmbh",
				Name:              "Tradevest",
				Followers:         "501",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tradevest-EI_IE9595327.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tradevest-Reviews-E9595327.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tradevest-Jobs-E9595327.htm",
				Jobs:        "2",
				Reviews:     "",
				Salaries:    "2",
				ReviewsRate: "",
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
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://justjoin.it/offers/tv-development-gmbh-senior-backend-developer",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Woolsocks
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Woolsocks",
			Website: "https://woolsocks.eu/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79728837,
				Alias:             "woolsocks",
				Name:              "Woolsocks",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "37",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "amsterdam-platform-creation",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Woolsocks-EI_IE8302146.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Woolsocks-Reviews-E8302146.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Woolsocks-Jobs-E8302146.htm",
				Jobs:        "",
				Reviews:     "2",
				Salaries:    "5",
				ReviewsRate: "3.5",
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
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3869628047/",
							Date:             mustDate("2024-03-25"),
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
		},

		// Some | Applied Systems Canada
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Systems Canada",
			Website: "https://www.appliedsystems.com/",
			Careers: "https://www1.appliedsystems.com/en-uk/about-us/jobs/",
			About:   "https://www1.appliedsystems.com/en-uk/about-us/about-applied/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                908801,
				IDs:               nil,
				Alias:             "applied-systems-canada",
				Name:              "Applied Systems Canada",
				Followers:         "9K",
				Employees:         "1K-5K",
				AssociatedMembers: "76",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Applied-Systems",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "4.0",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Applied-Systems-EI_IE8534.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Applied-Systems-Reviews-E8534.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Applied-Systems-Jobs-E8534.htm",
				Jobs:        "76",
				Reviews:     "671",
				Salaries:    "1K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Applied-Systems-6",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer (Platform, Reliability, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4068633971/",
							Date:                 mustDate("2025-02-10"), // mustDate("2024-12-20"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4045130649/",
							Date:                 mustDate("2025-02-10"), // mustDate("2025-01-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Golang/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4045127878/",
							Date:                 mustDate("2025-01-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Platform, Reliability, Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4067693574/",
							Date:                 mustDate("2025-02-15"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Autodesk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Autodesk",
			Website: "https://www.autodesk.com/",
			Careers: "https://www.autodesk.com/careers/overview",
			About:   "https://www.autodesk.com/company",
			Blog:    "https://aps.autodesk.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1879,
				Alias:             "autodesk",
				Name:              "Autodesk",
				Followers:         "886K",
				Employees:         "10K+",
				AssociatedMembers: "15,691",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Autodesk",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "autodesk",
				Employees:   "5,001 to 10,000",
				Salary:      "$63K ~ $300K a year",
				Reviews:     "564",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "autodesk",
				Employees: "12,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Autodesk-EI_IE1155.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Autodesk-Reviews-E1155.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Autodesk-Jobs-E1155.htm",
				Jobs:        "473",
				Reviews:     "5.2K",
				Salaries:    "9.8K",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Autodesk",
			},
			OttaProfileSlug:   "Autodesk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3843045776/",
							Date:             mustDate("2024-03-25"),
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

		// Some | Vonage
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vonage",
			Website: "https://www.vonage.com/",
			Careers: "https://www.vonage.com/careers/",
			About:   "https://www.vonage.com/about-us/",
			Blog:    "https://developer.vonage.com/en/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				IDs:               []int{5028, 1345545, 66428, 76778, 2102},
				Alias:             "vonage",
				Name:              "Vonage",
				Followers:         "128K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,821",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Vonage",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Vonage",
				Employees:   "1,001 to 5,000",
				Salary:      "$55K ~ $280K a year",
				Reviews:     "45",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vonage",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vonage-EI_IE23019.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vonage-Reviews-E23019.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vonage-Jobs-E23019.htm",
				Jobs:        "105",
				Reviews:     "1.3K",
				Salaries:    "1.9K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Vonage",
			},
			OttaProfileSlug:   "Vonage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3818119706/",
							Date:             mustDate("2024-04-25"),
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

		// Some | OpenWeb
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenWeb",
			Website: "https://www.openweb.com/",
			Careers: "https://www.openweb.com/jobs/",
			About:   "https://www.openweb.com/who-we-are/",
			Blog:    "https://www.openweb.com/blog/category/tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2506872,
				Alias:             "openwebhq",
				Name:              "OpenWeb",
				Followers:         "27K",
				Employees:         "201-500",
				AssociatedMembers: "300",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenWeb-EI_IE1675932.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenWeb-Reviews-E1675932.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OpenWeb-Jobs-E1675932.htm",
				Jobs:        "2",
				Reviews:     "125",
				Salaries:    "96",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "OpenWeb",
			YouTubeChannelURL: "https://www.youtube.com/@openwebhq",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/x-Xm2wSF",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Arenko
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Arenko",
			Website: "https://arenko.group/",
			Careers: "https://arenko.group/careers/",
			About:   "https://arenko.group/vision/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10288973,
				Alias:             "arenko-cleantech",
				Name:              "Arenko",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "60",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "arenko-group",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Arenko-Group-EI_IE4554199.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Arenko-Group-Reviews-E4554199.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Arenko-Group-Jobs-E4554199.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "19",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Arenko-Group",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4037978333/",
							Date:                 mustDate("2024-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118955695/",
							Date:                 mustDate("2025-02-19"), // mustDate("2025-02-07"), // mustDate("2025-02-03"), // mustDate("2025-01-15"),
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
			ShortDescription:          "Technology provider enabling the clean energy transition",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Xata
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Xata",
			Website: "https://xata.io/",
			Careers: "https://xata.io/careers",
			About:   "https://xata.io/about",
			Blog:    "https://xata.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69560619,
				Alias:             "xataio",
				Name:              "Xata.io",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "26",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "xataio",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "xata",
				Employees: "15",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Xata-EI_IE5816263.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Xata-Reviews-E5816263.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Xata-Jobs-E5816263.htm",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "7",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Xata",
			YouTubeChannelURL: "https://www.youtube.com/@xataio",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Golang Backend Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3887215661/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Dojo
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dojo",
			Website: "https://dojo.tech/",
			Careers: "https://dojo.careers/",
			About:   "https://dojo.tech/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42391390,
				Alias:             "dojo-tech",
				Name:              "Dojo",
				Followers:         "50K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,508",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dojo",
				Employees: "1,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dojo-EI_IE687810.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dojo-Reviews-E687810.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dojo-Jobs-E687810.htm",
				Jobs:        "18",
				Reviews:     "649",
				Salaries:    "1.1K",
				ReviewsRate: "3.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Dojo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3926951138/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Unnax
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unnax",
			Website: "https://www.unnax.com/",
			Careers: "",
			About:   "https://www.unnax.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10295665,
				Alias:             "unnax-emi",
				Name:              "Unnax",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "65",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "unnax",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unnax-EI_IE2108310.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unnax-Reviews-E2108310.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Unnax-Jobs-E2108310.htm",
				Jobs:        "",
				Reviews:     "19",
				Salaries:    "38",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@unnax-emi",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3890431726/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | AB Tasty
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AB Tasty",
			Website: "https://www.abtasty.com/",
			Careers: "https://careers.abtasty.com/",
			About:   "https://www.abtasty.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1885711,
				Alias:             "ab-tasty",
				Name:              "AB Tasty",
				Followers:         "26K",
				Employees:         "201-500",
				AssociatedMembers: "356",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ab-tasty",
				Employees: "300",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AB-Tasty-EI_IE1309242.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AB-Tasty-Reviews-E1309242.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AB-Tasty-Jobs-E1309242.htm",
				Jobs:        "5",
				Reviews:     "127",
				Salaries:    "221",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Ab-Tasty",
			},
			OttaProfileSlug:   "AB-Tasty",
			YouTubeChannelURL: "https://www.youtube.com/@abtasty",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3872953963/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Firebolt
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Firebolt",
			Website: "https://www.firebolt.io/",
			Careers: "https://www.firebolt.io/careers",
			About:   "https://www.firebolt.io/about-us",
			Blog:    "https://www.firebolt.io/knowledge-center/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40719957,
				Alias:             "firebolt",
				Name:              "Firebolt",
				Followers:         "29K",
				Employees:         "51-200",
				AssociatedMembers: "173",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "firebolt-db",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "firebolt",
				Employees: "60",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firebolt-EI_IE5001853.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firebolt-Reviews-E5001853.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Firebolt-Jobs-E5001853.htm",
				Jobs:        "",
				Reviews:     "32",
				Salaries:    "46",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "firebolt",
			YouTubeChannelURL: "https://www.youtube.com/@FireboltHQ",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3855486951/",
							Date:             mustDate("2024-04-25"),
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

		// Some | Nine
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nine",
			Website: "https://www.nineforbrands.com.au/",
			Careers: "https://ninecareers.com.au/",
			About:   "https://www.nineforbrands.com.au/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14933,
				Alias:             "nine-entertainment-co.",
				Name:              "Nine",
				Followers:         "89K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,170",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nine-Entertainment-EI_IE229827.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nine-Entertainment-Reviews-E229827.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nine-Entertainment-Jobs-E229827.htm",
				Jobs:        "37",
				Reviews:     "342",
				Salaries:    "708",
				ReviewsRate: "3.4",
				Verified:    true,
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
							Title:            "Senior Software Engineer (Golang) — Core Publishing",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3890491488/",
							Date:             mustDate("2024-05-24"),
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
				domain.Czechia,
			},
		},

		// Some | Isovalent
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Isovalent",
			Website: "https://isovalent.com/",
			Careers: "https://isovalent.com/careers/",
			About:   "https://isovalent.com/about-us/",
			Blog:    "https://isovalent.com/resource-library/blogs/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34714411,
				Alias:             "isovalent",
				Name:              "Isovalent",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "152",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "isovalent",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Isovalent",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "isovalent",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Isovalent-EI_IE3180689.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Isovalent-Reviews-E3180689.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Isovalent-Jobs-E3180689.htm",
				Jobs:        "",
				Reviews:     "4",
				Salaries:    "5",
				ReviewsRate: "5.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Isovalent",
			YouTubeChannelURL: "https://www.youtube.com/@isovalent",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 19,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Software Engineer — Hubble",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3893426061/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | ABC Fitness
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ABC Fitness",
			Website: "https://abcfitness.com/",
			Careers: "https://abcfitness.com/careers/",
			About:   "https://abcfitness.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                110372,
				Alias:             "abc-fitness",
				Name:              "ABC Fitness",
				Followers:         "34K",
				Employees:         "501-1K",
				AssociatedMembers: "2,697",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "ABC-Fitness-Solutions",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "abc-fitness",
				Employees: "810",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ABC-Fitness-EI_IE28305.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ABC-Fitness-Reviews-E28305.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ABC-Fitness-Jobs-E28305.htm",
				Jobs:        "29",
				Reviews:     "662",
				Salaries:    "858",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "ABC-Fitness-Solutions-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ABC-Fitness",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3890939506/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Device42
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Device42",
			Website: "https://www.device42.com/",
			Careers: "https://jobs.lever.co/device42",
			About:   "https://www.device42.com/company/",
			Blog:    "https://www.device42.com/blog/category/engineering-2/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2345405,
				Alias:             "device42",
				Name:              "Device42",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "121",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "device42",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Device42-EI_IE1705087.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Device42-Reviews-E1705087.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Device42-Jobs-E1705087.htm",
				Jobs:        "",
				Reviews:     "19",
				Salaries:    "39",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Device42",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Acronis
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Acronis",
			Website: "https://www.acronis.com/",
			Careers: "https://www.acronis.com/careers/",
			About:   "https://www.acronis.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13179,
				Alias:             "acronis",
				Name:              "Acronis",
				Followers:         "131K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,974",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "acronis",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Acronis",
				Employees:   "1001-5000",
				Salary:      "",
				Reviews:     "10",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "acronis",
				Employees: "1,600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Acronis-EI_IE152824.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Acronis-Reviews-E152824.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Acronis-Jobs-E152824.htm",
				Jobs:        "1",
				Reviews:     "574",
				Salaries:    "786",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Acronis-International-Gmbh",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Acronis",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3888974596/",
							Date:             mustDate("2024-04-25"),
						},
						{
							Title:            "Senior Backend Golang Developer",
							ShortDescription: "CyberSecurity",
							URL:              "https://www.linkedin.com/jobs/view/4130349633/",
							Date:             mustDate("2025-01-20"),
						},
						{
							Title:            "C++/Golang System Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4133051684/",
							Date:             mustDate("2025-01-23"),
						},
						{
							Title:                "Senior Backend Golang Developer",
							ShortDescription:     "Cybersecurity",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139224092/",
							Date:                 mustDate("2025-02-21"), // mustDate("2025-01-31"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Go/PHP/Python Senior Software Developer",
							ShortDescription:     "Billing/CSA",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4143606852/",
							Date:                 mustDate("2025-02-26"), // mustDate("2025-02-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go/Python Software Developer",
							ShortDescription:     "Application Security skills",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139225077/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "C++/Golang System Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139224113/",
							Date:                 mustDate("2025-02-21"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Gcore
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gcore",
			Website: "https://gcore.com/",
			Careers: "https://gcore.com/careers",
			About:   "https://gcore.com/about",
			Blog:    "https://gcore.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10351246,
				Alias:             "g-core",
				Name:              "Gcore",
				Followers:         "17K",
				Employees:         "501-1K",
				AssociatedMembers: "474",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "G-Core",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gcore-EI_IE2156026.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gcore-Reviews-E2156026.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gcore-Jobs-E2156026.htm",
				Jobs:        "7",
				Reviews:     "53",
				Salaries:    "83",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Gcore",
			YouTubeChannelURL: "https://www.youtube.com/@GCoreOfficial",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer (CDN)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3888070191/",
							Date:             mustDate("2024-04-25"),
						},
						{
							Title:            "Software Engineer (Go / DNS)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4132269261/",
							Date:             mustDate("2025-01-22"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Zep AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zep AI",
			Website: "https://www.getzep.com/",
			Careers: "https://www.getzep.com/careers",
			About:   "",
			Blog:    "https://blog.getzep.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                97181599,
				Alias:             "zep-ai",
				Name:              "Zep AI",
				Followers:         "679",
				Employees:         "2-10",
				AssociatedMembers: "6",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
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
			OttaProfileSlug:   "",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Gelato
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gelato",
			Website: "https://www.gelato.com/",
			Careers: "https://www.gelato.com/careers/careers-home",
			About:   "https://www.gelato.com/about-gelato",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5037871,
				Alias:             "gelato",
				Name:              "Gelato",
				Followers:         "47K",
				Employees:         "201-500",
				AssociatedMembers: "496",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gelatoas",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Gelato",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "2.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gelato-network",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gelato-EI_IE1297272.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gelato-Reviews-E1297272.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gelato-Jobs-E1297272.htm",
				Jobs:        "19",
				Reviews:     "172",
				Salaries:    "161",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Gelato",
			YouTubeChannelURL: "https://www.youtube.com/@GelatoConnects",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3866739991/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:            "Senior Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4129880810/",
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
				domain.Czechia,
			},
		},

		// Some | SumUp
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SumUp",
			Website: "https://www.sumup.com/",
			Careers: "https://www.sumup.com/careers/",
			About:   "https://www.sumup.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2619512,
				Alias:             "sumup",
				Name:              "SumUp",
				Followers:         "144k",
				Employees:         "1K-5K",
				AssociatedMembers: "3,565",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sumup",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SumUp",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "31",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sumup",
				Employees: "3,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SumUp-EI_IE673829.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SumUp-Reviews-E673829.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SumUp-Jobs-E673829.htm",
				Jobs:        "454",
				Reviews:     "1.3K",
				Salaries:    "2.6K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sumup",
			},
			OttaProfileSlug:   "sumup",
			YouTubeChannelURL: "https://www.youtube.com/@SumUpGlobal",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3848461191/",
							Date:                 mustDate("2024-05-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Software Engineer (Golang) — Online Payments",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3903334014/",
							Date:                 mustDate("2025-02-16"), // mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3943738572/",
							Date:                 mustDate("2025-02-10"), // mustDate("2025-01-14"),
							WithSalary:           true,                   // Salary range: 5250 - 7500 Euro GROSS / month + VSOP
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "Merchant Protect",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092326175/",
							Date:                 mustDate("2025-02-17"), // mustDate("2025-02-15"), // mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "Merchant Protect",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092325280/",
							Date:                 mustDate("2025-02-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "Global Bank",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3641494844/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               false,
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
							Title:                "(Senior) Backend Engineer — Elixir",
							ShortDescription:     "Our tech stack includes: Elixir, OTP, Postgres, Redis, Docker",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4056886705/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-12"), // mustDate("2024-12-06"), // mustDate("2024-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "(Senior) Backend Engineer — Elixir",
							ShortDescription:     "Our tech stack includes: Elixir, OTP, Postgres, Redis, Docker",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4054449086/",
							Date:                 mustDate("2025-02-22"), // mustDate("2025-02-07"), // mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Payments company",
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

		// Some | Level Home
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Level Home",
			Website: "https://level.co/",
			Careers: "https://level.co/about-us/careers/",
			About:   "https://level.co/about-us/people/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28779237,
				Alias:             "levelhome",
				Name:              "Level Home Inc.",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "136",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "LevelHome",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Level",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "3",
				ReviewsRate: "2.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "level-home",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Level-Home-EI_IE3556695.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Level-Home-Reviews-E3556695.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Level-Home-Jobs-E3556695.htm",
				Jobs:        "6",
				Reviews:     "46",
				Salaries:    "72",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Level-Home",
			YouTubeChannelURL: "https://www.youtube.com/@LevelHome",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3897014183/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | SonicWall
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SonicWall",
			Website: "https://www.sonicwall.com/",
			Careers: "https://www.sonicwall.com/about-sonicwall/careers",
			About:   "https://www.sonicwall.com/about-sonicwall/30-year-anniversary",
			Blog:    "https://www.sonicwall.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4926,
				Alias:             "sonicwall",
				Name:              "SonicWall",
				Followers:         "102K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,005",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sonicwall",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sonicwall",
				Employees: "1,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SonicWall-EI_IE9777.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SonicWall-Reviews-E9777.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SonicWall-Jobs-E9777.htm",
				Jobs:        "30",
				Reviews:     "1.1K",
				Salaries:    "1.2K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sonicwall",
			},
			OttaProfileSlug:   "SonicWall",
			YouTubeChannelURL: "https://www.youtube.com/@SonicWallInc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal API Engineer",
							ShortDescription: "Go Programming Language is a must",
							URL:              "https://www.linkedin.com/jobs/view/3875719837/",
							Date:             mustDate("2024-04-25"),
						},
						{
							Title:            "Backend Software Engineer",
							ShortDescription: "Golang experience is a must",
							URL:              "https://www.linkedin.com/jobs/view/4034349867/",
							Date:             mustDate("2025-01-09"), // mustDate("2024-12-19"),
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
		},

		// Some | Pindrop
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pindrop",
			Website: "https://www.pindrop.com/",
			Careers: "https://www.pindrop.com/company/careers/",
			About:   "https://www.pindrop.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2326557,
				Alias:             "pindrop",
				Name:              "Pindrop",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "284",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Pindrop",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "10",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pindrop",
				Employees: "240",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pindrop-EI_IE709157.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pindrop-Reviews-E709157.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pindrop-Jobs-E709157.htm",
				Jobs:        "2",
				Reviews:     "169",
				Salaries:    "248",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Pindrop",
			},
			OttaProfileSlug:   "Pindrop",
			YouTubeChannelURL: "https://www.youtube.com/@Pindropsecurity",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Backend Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3901803458/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Seedtag
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Seedtag",
			Website: "https://www.seedtag.com/",
			Careers: "https://jobs.seedtag.com/",
			About:   "https://www.seedtag.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5171806,
				Alias:             "seedtag",
				Name:              "Seedtag",
				Followers:         "97K",
				Employees:         "501-1K",
				AssociatedMembers: "645",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "seedtag",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Seedtag-EI_IE1421858.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Seedtag-Reviews-E1421858.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Seedtag-Jobs-E1421858.htm",
				Jobs:        "",
				Reviews:     "109",
				Salaries:    "180",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Seedtag",
			},
			OttaProfileSlug:   "Seedtag",
			YouTubeChannelURL: "https://www.youtube.com/@seedtag.advertising",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Performance AdTech)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3853123918/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Flix
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flix",
			Website: "https://www.flixbus.com/",
			Careers: "https://flix.careers/",
			About:   "https://corporate.flix.com/about-flix/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2726149,
				Alias:             "flixbus",
				Name:              "Flix",
				Followers:         "97K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,246",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "flix-tech",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Flixbus",
				Employees:   "1,001-5,000",
				Salary:      "",
				Reviews:     "4",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flix",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flix-EI_IE976463.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flix-Reviews-E976463.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flix-Jobs-E976463.htm",
				Jobs:        "200",
				Reviews:     "589",
				Salaries:    "1.1K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Flix-Se-2",
			},
			OttaProfileSlug:   "FlixBus-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 17,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4101217107/",
							Date:                 mustDate("2024-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Back-End (Go) Engineer  — Content Platform team",
							ShortDescription:     "Backend: Go, ArangoDB, MySQL, Confluent Kafka, Snowflake, GraphQL + different AWS services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4091237831/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4093658841/",
							Date:                 mustDate("2025-01-20"), // mustDate("2025-01-16"), // mustDate("2025-01-11"),
							WithSalary:           false,
							Remote:               false, // Up to 60 days of working from (m)anywhere
						},
						{
							Title:                "Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4143253999/",
							Date:                 mustDate("2025-02-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Back-End (Go) Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4138356213/",
							Date:                 mustDate("2025-02-21"),
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
				domain.Czechia,
			},
		},

		// Some | Press Ganey
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Press Ganey",
			Website: "https://www.pressganey.com/",
			Careers: "https://www.pressganey.com/company/careers/",
			About:   "https://www.pressganey.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18007,
				Alias:             "press-ganey-associates",
				Name:              "Press Ganey",
				Followers:         "90K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,070",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Press-Ganey",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "10",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "press-ganey",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Press-Ganey-EI_IE35100.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Press-Ganey-Reviews-E35100.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Press-Ganey-Jobs-E35100.htm",
				Jobs:        "57",
				Reviews:     "784",
				Salaries:    "1.2K",
				ReviewsRate: "3.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Press-Ganey-4",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PressGaney",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3898197086/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Atmail
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Atmail",
			Website: "https://www.atmail.com/",
			Careers: "",
			About:   "https://www.atmail.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                918978,
				Alias:             "atmail",
				Name:              "Atmail",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "32",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
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
							URL:              "https://www.linkedin.com/jobs/view/3843643222/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Dusty Robotics
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dusty Robotics",
			Website: "https://www.dustyrobotics.com/",
			Careers: "https://www.dustyrobotics.com/careers",
			About:   "https://www.dustyrobotics.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33298433,
				Alias:             "dusty-robotics",
				Name:              "Dusty Robotics",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "88",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dusty-Robotics-EI_IE3518259.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dusty-Robotics-Reviews-E3518259.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dusty-Robotics-Jobs-E3518259.htm",
				Jobs:        "",
				Reviews:     "25",
				Salaries:    "34",
				ReviewsRate: "4.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dustyrobotics",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer fullstack/backend",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3804188613/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Cimri
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cimri",
			Website: "https://www.cimri.com/",
			Careers: "",
			About:   "https://www.cimri.com/cimri-hakkinda",
			Blog:    "https://engineering.cimri.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                411498,
				Alias:             "cimri",
				Name:              "Cimri",
				Followers:         "32K",
				Employees:         "51-200",
				AssociatedMembers: "178",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cimri-EI_IE2401296.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cimri-Reviews-E2401296.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cimri-Jobs-E2401296.htm",
				Jobs:        "",
				Reviews:     "2",
				Salaries:    "20",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cimri",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3863565726/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Quadcode
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quadcode",
			Website: "https://quadcode.com/",
			Careers: "https://jobs.quadcode.com/jobs",
			About:   "https://quadcode.com/about",
			Blog:    "https://quadcode.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42345997,
				Alias:             "quadcodecareer",
				Name:              "Quadcode",
				Followers:         "38K",
				Employees:         "501-1K",
				AssociatedMembers: "422",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "quadcode",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quadcode-EI_IE3293771.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quadcode-Reviews-E3293771.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Quadcode-Jobs-E3293771.htm",
				Jobs:        "60",
				Reviews:     "52",
				Salaries:    "107",
				ReviewsRate: "3.3",
				Verified:    true,
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
							URL:              "https://www.linkedin.com/jobs/view/3907613945/",
							Date:             mustDate("2024-04-25"),
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
			ShortDescription: "Fintech company developing software for the trading and investment industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | FinCompare
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FinCompare",
			Website: "https://fincompare.de/",
			Careers: "https://fincompare.de/jobs",
			About:   "https://fincompare.de/ueber-fincompare",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10976436,
				Alias:             "fincompare",
				Name:              "FinCompare",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fincompare",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FinCompare-EI_IE1644778.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FinCompare-Reviews-E1644778.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FinCompare-Jobs-E1644778.htm",
				Jobs:        "4",
				Reviews:     "45",
				Salaries:    "51",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3846776554/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Mellifera
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mellifera",
			Website: "https://mellifera.team/",
			Careers: "https://mellifera.team/careers/",
			About:   "https://mellifera.team/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98533069,
				Alias:             "mellifera-operations-limited",
				Name:              "Mellifera Operations Limited",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "85",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go (Golang) Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912533167/",
							Date:             mustDate("2024-04-25"),
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

		// Some | BNP Paribas — Securities Services
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BNP Paribas — Securities Services",
			Website: "https://securities.cib.bnpparibas/",
			Careers: "https://securities.cib.bnpparibas/who-we-are/your-career-with-us/",
			About:   "https://securities.cib.bnpparibas/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3625182,
				Alias:             "bnpparibassecuritiesservices",
				Name:              "BNP Paribas - Securities Services",
				Followers:         "137K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,620",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "BNP-Paribas",
				Employees:   "10,000+",
				Salary:      "$45K ~ $236K a year",
				Reviews:     "43",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bnp-paribas-securities",
				Employees: "17,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BNP-Paribas-EI_IE10342.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BNP-Paribas-Reviews-E10342.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BNP-Paribas-Jobs-E10342.htm",
				Jobs:        "3.1K",
				Reviews:     "18K",
				Salaries:    "28K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@labanquedunmondequichange",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Tech Lead Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3909846279/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Apifonica
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Apifonica",
			Website: "https://www.apifonica.com/",
			Careers: "https://www.apifonica.com/en/industries/hr-and-recruitment/",
			About:   "https://www.apifonica.com/en/company/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10943475,
				Alias:             "apifonica",
				Name:              "Apifonica",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "45",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "apifonica",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apifonica-EI_IE1805118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apifonica-Reviews-E1805118.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Apifonica-Jobs-E1805118.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "10",
				ReviewsRate: "4.3",
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
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3909698518/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "3+ years of engineering and development experience (PHP/Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122032404/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142622628/",
							Date:                 mustDate("2025-02-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142623623/",
							Date:                 mustDate("2025-02-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161464959/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161466631/",
							Date:                 mustDate("2025-02-21"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Cybus
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Cybus",
			Website: "https://www.cybus.io/",
			Careers: "https://www.cybus.io/en/jobs/career/",
			About:   "https://www.cybus.io/en/company/about-cybus/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7798667,
				Alias:             "cybus",
				Name:              "Cybus",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "62",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cybusio",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cybus-EI_IE2928520.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cybus-Reviews-E2928520.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cybus-Jobs-E2928520.htm",
				Jobs:        "3",
				Reviews:     "5",
				Salaries:    "11",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Cybus",
			YouTubeChannelURL: "https://www.youtube.com/@cybus_io",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend (Node.js) Engineer",
							ShortDescription: "Work with our application core technologies, which is Node.js and Go and modern broker technologies",
							URL:              "https://www.linkedin.com/jobs/view/4054430926/",
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
			ShortDescription:          "Data architecture for the industrial IoT world",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Flink
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flink",
			Website: "https://www.goflink.com/",
			Careers: "https://careers.smartrecruiters.com/Flink3/joinus",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71241902,
				Alias:             "goflink",
				Name:              "Flink",
				Followers:         "52K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,478",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flink",
				Employees: "12,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flink-EI_IE4921496.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flink-Reviews-E4921496.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flink-Jobs-E4921496.htm",
				Jobs:        "370",
				Reviews:     "724",
				Salaries:    "1.2K",
				ReviewsRate: "2.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Flink",
			YouTubeChannelURL: "https://www.youtube.com/@flink7309",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "2 years of commercial experience with Golang",
							URL:              "https://www.linkedin.com/jobs/view/3906306472/",
							Date:             mustDate("2024-04-25"),
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
				domain.Czechia,
			},
		},

		// Some | Greenbone
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Greenbone",
			Website: "https://www.greenbone.net/",
			Careers: "https://www.greenbone.net/careers/",
			About:   "https://www.greenbone.net/about-greenbone/",
			Blog:    "https://community.greenbone.net/blog/category/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                697428,
				Alias:             "greenbone-networks-gmbh",
				Name:              "Greenbone AG",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "104",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "greenbone",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Greenbone-Networks-EI_IE379229.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Greenbone-Networks-Reviews-E379229.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Greenbone-Networks-Jobs-E379229.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "12",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@greenbone",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3892843872/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | OLX
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OLX",
			Website: "https://www.olxgroup.com/",
			Careers: "https://careers.olxgroup.com/",
			About:   "https://www.olxgroup.com/about-us/",
			Blog:    "https://tech.olx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167557,
				Alias:             "olx-group",
				Name:              "OLX",
				Followers:         "213K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,306",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "OLX",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "11",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "olx-group",
				Employees: "87",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OLX-Group-EI_IE517166.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OLX-Group-Reviews-E517166.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OLX-Group-Jobs-E517166.htm",
				Jobs:        "",
				Reviews:     "1.2K",
				Salaries:    "2K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Olx-Group",
			},
			OttaProfileSlug:   "OLX-Group",
			YouTubeChannelURL: "https://www.youtube.com/@OLXGroup",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Middle/Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055838812/",
							Date:             mustDate("2024-12-22"),
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133389022/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "Real Estate",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142492503/",
							Date:                 mustDate("2025-02-25"), // mustDate("2025-02-03"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang, Java, PHP)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137532355/",
							Date:                 mustDate("2025-02-19"),
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
			ShortDescription: "Marketplace ecosystems",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Snyk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Snyk",
			Website: "https://snyk.io/",
			Careers: "https://snyk.io/careers/",
			About:   "https://snyk.io/about/",
			Blog:    "https://snyk.io/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10043614,
				Alias:             "snyk",
				Name:              "Snyk",
				Followers:         "92K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,290",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "snyk",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Snyk",
				Employees:   "1,001 to 5,000 Employees",
				Salary:      "",
				Reviews:     "83",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "snyk",
				Employees: "930",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snyk-EI_IE2094989.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snyk-Reviews-E2094989.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Snyk-Jobs-E2094989.htm",
				Jobs:        "40",
				Reviews:     "299",
				Salaries:    "655",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Snyk",
			},
			OttaProfileSlug:   "Snyk",
			YouTubeChannelURL: "https://www.youtube.com/@Snyksec",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 36,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4006143681/",
							Date:                 mustDate("2025-01-08"), // mustDate("2024-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Policies Team (TypeScript, Go)",
							ShortDescription:     "7+ years experience in software engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4024339260/",
							Date:                 mustDate("2025-01-14"),
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
			HasEmployeesFromCountries: nil,
		},

		// Some | Sinch
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sinch",
			Website: "https://www.sinch.com/",
			Careers: "https://sinch.com/careers/",
			About:   "https://sinch.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3726743,
				Alias:             "sinch",
				Name:              "Sinch",
				Followers:         "398K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,169",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sinch",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Sinch",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "9",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sinch",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sinch-EI_IE788805.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sinch-Reviews-E788805.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sinch-Jobs-E788805.htm",
				Jobs:        "118",
				Reviews:     "547",
				Salaries:    "1K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sinch",
			},
			OttaProfileSlug:   "sinch",
			YouTubeChannelURL: "https://www.youtube.com/@WeAreSinch",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3903797937/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | FOX Tech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FOX Tech",
			Website: "https://tech.fox.com/",
			Careers: "https://www.foxcareers.com/",
			About:   "https://www.foxcorporation.com/about-us/",
			Blog:    "https://medium.com/fox-tech",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14850572,
				Alias:             "foxtechteam",
				Name:              "FOX Tech",
				Followers:         "11K",
				Employees:         "5K-10K",
				AssociatedMembers: "111",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Fox",
				Employees:   "10,000+",
				Salary:      "$66K ~ $318K a year",
				Reviews:     "67",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fox",
				Employees: "120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FOX-Broadcasting-EI_IE13272.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FOX-Broadcasting-Reviews-E13272.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FOX-Broadcasting-Jobs-E13272.htm",
				Jobs:        "173",
				Reviews:     "313",
				Salaries:    "552",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Fox-Corporation",
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
							URL:              "https://www.linkedin.com/jobs/view/3903338524/",
							Date:             mustDate("2024-08-26"),
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
		},

		// Some | SailPoint
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SailPoint",
			Website: "https://www.sailpoint.com/",
			Careers: "https://www.sailpoint.com/company/careers",
			About:   "https://www.sailpoint.com/why-us/about-us",
			Blog:    "https://developer.sailpoint.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47456,
				Alias:             "sailpoint-technologies",
				Name:              "SailPoint",
				Followers:         "125K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,830",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sailpoint-oss",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SailPoint-Technologies",
				Employees:   "1,001 to 5,000",
				Salary:      "$60K ~ $250K a year",
				Reviews:     "98",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sailpoint",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SailPoint-Technologies-EI_IE449696.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SailPoint-Technologies-Reviews-E449696.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SailPoint-Technologies-Jobs-E449696.htm",
				Jobs:        "94",
				Reviews:     "939",
				Salaries:    "1.3K",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sailpoint-Technologies-2",
			},
			OttaProfileSlug:   "Sailpoint",
			YouTubeChannelURL: "https://www.youtube.com/@SailPointTechnologies",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4030943277/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4107495969/",
							Date:                 mustDate("2024-12-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077458993/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149693865/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // $110,250 - $157,500 - $204,750 (min-mid-max)
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148940672/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Java, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151653458/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // salary range $74,060 - $105,800 - $137,540
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137755120/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158963804/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           true, // salary range $69,860 - $99,800 - $129,740
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
			Industries:       []domain.Industry{},
		},

		// Some | Proofpoint
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Proofpoint",
			Website: "https://www.proofpoint.com/",
			Careers: "https://www.proofpoint.com/us/company/careers",
			About:   "https://www.proofpoint.com/us/company/about",
			Blog:    "https://www.proofpoint.com/us/blog/engineering-insights",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11681,
				Alias:             "proofpoint",
				Name:              "Proofpoint",
				Followers:         "154K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,698",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "proofpoint",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Proofpoint",
				Employees:   "1,001 to 5,000",
				Salary:      "$62K ~ $324K a year",
				Reviews:     "162",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "proofpoint",
				Employees: "3,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proofpoint-EI_IE39140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proofpoint-Reviews-E39140.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Proofpoint-Jobs-E39140.htm",
				Jobs:        "140",
				Reviews:     "1.1K",
				Salaries:    "2.1K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Proofpoint",
			YouTubeChannelURL: "https://www.youtube.com/@proofpoint",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3797062001/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Asset Reality
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Asset Reality",
			Website: "https://www.assetreality.com/",
			Careers: "",
			About:   "https://www.assetreality.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42805677,
				Alias:             "asset-reality",
				Name:              "Asset Reality",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "50",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
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
			OttaProfileSlug:   "Asset-Reality",
			YouTubeChannelURL: "https://www.youtube.com/@assetreality",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "",
								Date:             mustDate(""),
							},
						*/
						{
							Title:                "Senior Backend Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4165243684/",
							Date:                 mustDate("2025-02-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Limango
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Limango",
			Website: "https://www.limango.pl/",
			Careers: "https://praca.limango.pl/oferty-pracy/oferta-product-owner/",
			About:   "https://www.limango.pl/info/limango",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2966982,
				Alias:             "limango-sp-z-o-o-",
				Name:              "Limango",
				Followers:         "2K",
				Employees:         "501-1K",
				AssociatedMembers: "179",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Limango-Polska-EI_IE2884426.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Limango-Polska-Reviews-E2884426.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Limango-Polska-Jobs-E2884426.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "28",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@limangoPolska",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://justjoin.it/offers/limango-polska-mid-senior-backend-developer-golang",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | RxBenefits
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RxBenefits",
			Website: "https://www.rxbenefits.com/",
			Careers: "https://www.rxbenefits.com/about-us/culture/",
			About:   "https://www.rxbenefits.com/about-us/about-rxbenefits/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2873210,
				Alias:             "rxbenefits-inc-",
				Name:              "RxBenefits, Inc.",
				Followers:         "38K",
				Employees:         "1K-5K",
				AssociatedMembers: "879",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RxBenefits-EI_IE1175839.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RxBenefits-Reviews-E1175839.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RxBenefits-Jobs-E1175839.htm",
				Jobs:        "103",
				Reviews:     "147",
				Salaries:    "247",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Rxbenefits,-Inc.-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3910221910/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | SmithRx
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmithRx",
			Website: "https://www.smithrx.com/",
			Careers: "https://www.smithrx.com/careers",
			About:   "https://www.smithrx.com/our-story",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10886362,
				Alias:             "smithrx",
				Name:              "SmithRx",
				Followers:         "11K",
				Employees:         "501-1K",
				AssociatedMembers: "409",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SmithRx-EI_IE1901555.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SmithRx-Reviews-E1901555.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SmithRx-Jobs-E1901555.htm",
				Jobs:        "11",
				Reviews:     "66",
				Salaries:    "112",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "SmithRx",
			YouTubeChannelURL: "https://www.youtube.com/@SmithRxPBM",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | k-ID
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "k-ID",
			Website: "https://www.k-id.com/",
			Careers: "https://www.k-id.com/careers",
			About:   "https://www.k-id.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                45973117,
				Alias:             "k-id",
				Name:              "k-ID",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "41",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "k-id",
				Employees: "42",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@k-IDofficial",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912660665/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | CAFU
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CAFU",
			Website: "https://www.cafu.com/",
			Careers: "https://www.cafu.com/careers",
			About:   "https://www.cafu.com/about",
			Blog:    "https://cafu.engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13892019,
				Alias:             "mycafu",
				Name:              "CAFU",
				Followers:         "293K",
				Employees:         "51-200",
				AssociatedMembers: "1,090",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CAFU-EI_IE3713615.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CAFU-Reviews-E3713615.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CAFU-Jobs-E3713615.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "118",
				ReviewsRate: "3.0",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Rollee
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rollee",
			Website: "https://www.getrollee.com/",
			Careers: "https://www.getrollee.com/company/careers",
			About:   "https://www.getrollee.com/company/mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76353840,
				Alias:             "rollee",
				Name:              "Rollee",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "26",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
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
			OttaProfileSlug:   "Rollee",
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
								URL:              "https://app.welcometothejungle.com/jobs/4ajPfsAU",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Net2Phone
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Net2Phone",
			Website: "https://www.net2phone.com/",
			Careers: "https://www.net2phone.com/careers",
			About:   "https://www.net2phone.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2965,
				Alias:             "net2phone",
				Name:              "Net2Phone",
				Followers:         "10K",
				Employees:         "1K-5K",
				AssociatedMembers: "259",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "net2phone",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Net2Phone-EI_IE9342.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Net2Phone-Reviews-E9342.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Net2Phone-Jobs-E9342.htm",
				Jobs:        "5",
				Reviews:     "62",
				Salaries:    "104",
				ReviewsRate: "3.6",
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
							URL:              "https://www.linkedin.com/jobs/view/3914839595/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Ola Chat
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ola Chat",
			Website: "https://olachat.sg/",
			Careers: "",
			About:   "https://olachat.sg/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68332088,
				Alias:             "ola-chat",
				Name:              "Ola Chat",
				Followers:         "47K",
				Employees:         "1K-5K",
				AssociatedMembers: "243",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "olachat",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912130901/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:            "Back-End Developer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105518392/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:            "Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4132246194/",
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
		},

		// Some | Veracity Software Inc
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Veracity Software Inc",
			Website: "https://veracity-us.com/",
			Careers: "https://veracity-us.com//career",
			About:   "https://veracity-us.com//about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11137552,
				Alias:             "veracitysoftwareinc",
				Name:              "Veracity Software Inc",
				Followers:         "46K",
				Employees:         "51-200",
				AssociatedMembers: "60",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "veracity-software",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Veracity-Software-EI_IE1357198.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Veracity-Software-Reviews-E1357198.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Veracity-Software-Jobs-E1357198.htm",
				Jobs:        "11",
				Reviews:     "78",
				Salaries:    "115",
				ReviewsRate: "4.1",
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
							Title:            "Full Stack Java / Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3818383367/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Treecard
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Treecard",
			Website: "https://www.treecard.org/",
			Careers: "https://www.treecard.org/careers",
			About:   "https://www.treecard.org/the-mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68821773,
				Alias:             "treecardapp",
				Name:              "Treecard",
				Followers:         "16K",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "TreeCard",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Treecard-EI_IE5675051.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Treecard-Reviews-E5675051.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Treecard-Jobs-E5675051.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "15",
				ReviewsRate: "4.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Treecard",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Product Engineer (Backend — Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3909975571/",
							Date:             mustDate("2024-04-25"),
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
		},

		// Some | Openprovider
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Openprovider",
			Website: "https://www.openprovider.com/",
			Careers: "https://www.openprovider.com/company/careers",
			About:   "https://www.openprovider.com/company/about-us",
			Blog:    "https://www.openprovider.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                550698,
				Alias:             "openprovider",
				Name:              "Openprovider",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "73",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "openprovider",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Openprovider-EI_IE1267124.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Openprovider-Reviews-E1267124.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Openprovider-Jobs-E1267124.htm",
				Jobs:        "17",
				Reviews:     "30",
				Salaries:    "28",
				ReviewsRate: "4.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@openprovider453",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3914419419/",
							Date:             mustDate("2024-04-25"),
						},
						{
							Title:                "Full Stack Developer (ReactJS & Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4166998360/",
							Date:                 mustDate("2025-02-24"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | fiskaly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "fiskaly",
			Website: "https://www.fiskaly.com/",
			Careers: "https://www.fiskaly.com/jobs",
			About:   "https://www.fiskaly.com/about",
			Blog:    "https://developer.fiskaly.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18929063,
				Alias:             "fiskaly",
				Name:              "fiskaly",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "86",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fiskaly",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-fiskaly-EI_IE3059515.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/fiskaly-Reviews-E3059515.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/fiskaly-Jobs-E3059515.htm",
				Jobs:        "8",
				Reviews:     "4",
				Salaries:    "25",
				ReviewsRate: "4.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fiskaly",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914242491/", // 5 years working with Golang or other OOP languages
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4055768396/", // 5 years working with Golang or other OOP languages
							Date:                 mustDate("2024-12-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157742979/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true, // €56.000 - €70.000 per year
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
			ShortDescription: "fiskaly builds innovative tools to make receipts simple",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Vay
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vay",
			Website: "https://vay.io/",
			Careers: "https://vay.io/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12584218,
				Alias:             "vaytechnology",
				Name:              "Vay",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "162",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Reemote",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Vay",
			YouTubeChannelURL: "https://www.youtube.com/@vay_io",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Voltus
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Voltus",
			Website: "https://www.voltus.co/",
			Careers: "https://www.voltus.co/careers",
			About:   "https://www.voltus.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10866628,
				Alias:             "voltus-inc.",
				Name:              "Voltus",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "256",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "voltusdev",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voltus-EI_IE2090197.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voltus-Reviews-E2090197.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Voltus-Jobs-E2090197.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "178",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Voltus",
			YouTubeChannelURL: "https://www.youtube.com/@Voltusinc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/pB_hGf0W",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Stonebranch
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stonebranch",
			Website: "https://www.stonebranch.com/",
			Careers: "https://www.stonebranch.com/careers",
			About:   "https://www.stonebranch.com/about-us",
			Blog:    "https://www.stonebranch.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71261,
				Alias:             "stonebranch",
				Name:              "Stonebranch",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "161",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stonebranch-marketplace",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "stonebranch",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stonebranch-EI_IE408996.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stonebranch-Reviews-E408996.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stonebranch-Jobs-E408996.htm",
				Jobs:        "17",
				Reviews:     "30",
				Salaries:    "46",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Stonebranch",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Stonebranch",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3914103960/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:            "Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106867221/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Rapid7
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rapid7",
			Website: "https://www.rapid7.com/",
			Careers: "https://careers.rapid7.com/",
			About:   "https://www.rapid7.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                39624,
				Alias:             "rapid7",
				Name:              "Rapid7",
				Followers:         "183K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,070",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "rapid7",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Rapid8",
				Employees:   "1,001 to 5,000",
				Salary:      "$75K ~ $350K a year",
				Reviews:     "69",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rapid7",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rapid7-EI_IE243542.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rapid7-Reviews-E243542.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rapid7-Jobs-E243542.htm",
				Jobs:        "193",
				Reviews:     "893",
				Salaries:    "2.1K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Rapid7",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@GoRapid7",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "Open-source platform — Velociraptor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3915341496/",
							Date:                 mustDate("2024-09-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4159762093/",
							Date:                 mustDate("2025-02-20"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Toggle AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Toggle AI",
			Website: "https://toggle.ai/",
			Careers: "https://toggle.ai/careers",
			About:   "https://toggle.ai/#aboutUs",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28508827,
				Alias:             "toggle-ai",
				Name:              "Toggle AI",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Toggle-EI_IE3924898.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Toggle-Reviews-E3924898.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Toggle-Jobs-E3924898.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "11",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ToggleAI-Investing",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3915194291/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | hearX Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "hearX Group",
			Website: "https://www.hearxgroup.com/",
			Careers: "https://hearxgroup.simplify.hr/",
			About:   "https://hearxgroup.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18002825,
				Alias:             "hearx-group",
				Name:              "hearX Group",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "160",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hearxgroup",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-hearX-Group-EI_IE5800566.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/hearX-Group-Reviews-E5800566.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/hearX-Group-Jobs-E5800566.htm",
				Jobs:        "",
				Reviews:     "4",
				Salaries:    "11",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hearxgroupptyltd8061",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | MarkiTech.AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MarkiTech.AI",
			Website: "https://markitech.ca/",
			Careers: "https://markitech.ca/jobs/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9294422,
				Alias:             "markitech-ai",
				Name:              "MarkiTech.AI",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "32",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Markitech-EI_IE4190913.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Markitech-Reviews-E4190913.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Markitech-Jobs-E4190913.htm",
				Jobs:        "1",
				Reviews:     "33",
				Salaries:    "33",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@markitech-digitaltransform6173",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Lantronix
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lantronix",
			Website: "https://www.lantronix.com/",
			Careers: "https://www.lantronix.com/about-us/careers/",
			About:   "https://www.lantronix.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12612,
				Alias:             "lantronix",
				Name:              "Lantronix",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "443",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Lantronix",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lantronix-EI_IE5498.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lantronix-Reviews-E5498.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@LantronixInc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | INFOLOB
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "INFOLOB",
			Website: "https://www.infolob.com/",
			Careers: "https://www.infolob.com/careers/",
			About:   "https://www.infolob.com/about/",
			Blog:    "https://www.infolob.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                404211,
				Alias:             "infolob-global",
				Name:              "INFOLOB",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "468",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infolob-Solutions-EI_IE423113.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infolob-Solutions-Reviews-E423113.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Infolob-Solutions-Jobs-E423113.htm",
				Jobs:        "6",
				Reviews:     "120",
				Salaries:    "524",
				ReviewsRate: "4.3",
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
							Title:            "Senior Python AWS Developer / Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3914415181/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Argela Technologies
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Argela Technologies",
			Website: "https://www.argela.com.tr/",
			Careers: "https://www.argela.com.tr/en/career",
			About:   "https://www.argela.com.tr/en/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26805,
				Alias:             "argela-technologies",
				Name:              "Argela Technologies",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "190",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Argela-EI_IE389528.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Argela-Reviews-E389528.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Argela-Jobs-E389528.htm",
				Jobs:        "",
				Reviews:     "57",
				Salaries:    "57",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@argelatechnologies",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer — C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914290466/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151349412/",
							Date:                 mustDate("2025-02-13"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Top Doctors
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Top Doctors",
			Website: "https://topdoctors.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3035481,
				Alias:             "top-doctors-europe",
				Name:              "Top Doctors",
				Followers:         "35K",
				Employees:         "201-500",
				AssociatedMembers: "452",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Top-Doctors-EI_IE1712721.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Top-Doctors-Reviews-E1712721.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Top-Doctors-Jobs-E1712721.htm",
				Jobs:        "27",
				Reviews:     "95",
				Salaries:    "145",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Top-Doctors",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TopDoctorsUK",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Developer Backend Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077191404/",
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

		// Some | Recurly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Recurly",
			Website: "https://recurly.com/",
			Careers: "https://recurly.com/careers/",
			About:   "https://recurly.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                810383,
				Alias:             "recurly",
				Name:              "Recurly",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "300",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "recurly",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "recurly",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Recurly-EI_IE692611.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Recurly-Reviews-E692611.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Recurly-Jobs-E692611.htm",
				Jobs:        "25",
				Reviews:     "178",
				Salaries:    "319",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Recurly",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Recurly",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Features (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912991626/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Cynet Systems
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cynet Systems",
			Website: "https://www.cynetsystems.com/",
			Careers: "https://www.cynetsystems.com/jobs/",
			About:   "https://www.cynetsystems.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5003556,
				Alias:             "cynet-systems",
				Name:              "Cynet Systems",
				Followers:         "72K",
				Employees:         "1K-5K",
				AssociatedMembers: "682",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynet-Systems-EI_IE654628.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cynet-Systems-Reviews-E654628.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cynet-Systems-Jobs-E654628.htm",
				Jobs:        "5",
				Reviews:     "436",
				Salaries:    "343",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cynet-Systems",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cynetsystems2026",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3912977786/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Odyssey Information Services
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Odyssey Information Services",
			Website: "https://www.odysseyis.com/",
			Careers: "",
			About:   "https://www.odysseyis.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                56445,
				Alias:             "odyssey-information-services",
				Name:              "Odyssey Information Services",
				Followers:         "117K",
				Employees:         "201-500",
				AssociatedMembers: "164",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "odyssey-information-services",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Odyssey-Information-Services-EI_IE558201.11,39.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Odyssey-Information-Services-Reviews-E558201.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Odyssey-Information-Services-Jobs-E558201.htm",
				Jobs:        "18",
				Reviews:     "47",
				Salaries:    "43",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Odyssey-Information-Services",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@odysseyinformationservices1641",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Cloud Security Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3905563294/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152455199/",
							Date:                 mustDate("2025-02-26"), // mustDate("2025-02-19"), // mustDate("2025-02-18"), // mustDate("2025-02-17"), // mustDate("2025-02-13"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Infomatics Corp
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Infomatics Corp",
			Website: "https://infomaticscorp.com/",
			Careers: "https://infomaticscorp.com/careers/",
			About:   "https://infomaticscorp.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                821065,
				Alias:             "infomatics-corp",
				Name:              "Infomatics Corp",
				Followers:         "36K",
				Employees:         "51-200",
				AssociatedMembers: "99",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infomatics-EI_IE925978.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infomatics-Reviews-E925978.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Infomatics-Jobs-E925978.htm",
				Jobs:        "",
				Reviews:     "54",
				Salaries:    "253",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Infomatics-Corp",
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
							URL:              "https://www.linkedin.com/jobs/view/3912956713/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4124982450/",
							Date:             mustDate("2025-01-15"),
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
		},

		// Some | Mindera
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mindera",
			Website: "https://mindera.com/",
			Careers: "https://mindera.com/careers",
			About:   "https://mindera.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80044268,
				Alias:             "mindera-world",
				Name:              "Mindera",
				Followers:         "96K",
				Employees:         "1K-5K",
				AssociatedMembers: "952",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Mindera",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Mindera",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "6",
				ReviewsRate: "4.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mindera",
				Employees: "760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mindera-EI_IE1139926.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mindera-Reviews-E1139926.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mindera-Jobs-E1139926.htm",
				Jobs:        "53",
				Reviews:     "265",
				Salaries:    "516",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Mindera",
			},
			OttaProfileSlug:   "Mindera",
			YouTubeChannelURL: "https://www.youtube.com/@MinderaSoftware",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Back-End Golang Senior Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3915092213/",
							Date:             mustDate("2024-05-24"),
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
				domain.Czechia,
			},
		},

		// Some | Sytac
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sytac",
			Website: "https://sytac.io/",
			Careers: "https://sytac.io/careers/",
			About:   "https://sytac.io/aboutus/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                626751,
				Alias:             "sytac",
				Name:              "Sytac",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "102",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sytac",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sytac-EI_IE1255983.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sytac-Reviews-E1255983.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@sytac",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Qumulus Cloud Platform
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qumulus Cloud Platform",
			Website: "https://www.qumulus.io/",
			Careers: "",
			About:   "https://www.qumulus.io/about",
			Blog:    "https://www.qumulus.io/blog-pages/blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81905634,
				Alias:             "qumuluscloudplatform",
				Name:              "Qumulus Cloud Platform",
				Followers:         "1K",
				Employees:         "2-10",
				AssociatedMembers: "13",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "QumulusTechnology",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Golang Engineer — Kubernetes",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3914358799/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Saxon AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Saxon AI",
			Website: "https://saxon.ai/",
			Careers: "https://saxon.ai/career/",
			About:   "https://saxon.ai/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                224935,
				Alias:             "saxonai",
				Name:              "Saxon AI",
				Followers:         "135K",
				Employees:         "201-500",
				AssociatedMembers: "374",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@saxonai",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Dyninno Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dyninno Group",
			Website: "https://dyninno.com/",
			Careers: "https://dyninno.com/en/careers/",
			About:   "https://dyninno.com/en/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9456141,
				Alias:             "dyninno-group",
				Name:              "Dyninno Group",
				Followers:         "60K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,713",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dyninno",
				Employees: "5,400",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DYNINNO-Group-EI_IE2562842.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DYNINNO-Group-Reviews-E2562842.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DYNINNO-Group-Jobs-E2562842.htm",
				Jobs:        "11",
				Reviews:     "285",
				Salaries:    "316",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Dyninno-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dyninnogroup2702",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3914135515/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Holland & Barrett
		{
			ID:      0,
			Name:    "Holland & Barrett",
			Website: "https://www.hollandandbarrett.com/",
			Careers: "https://www.hollandandbarrettjobs.com/",
			About:   "https://www.hollandandbarrett.com/info/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    815488,
				Alias: "815488",
				Name:  "Holland & Barrett",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "holland-and-barrett",
				Employees: "7,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Holland-and-Barrett-EI_IE36174.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Holland-and-Barrett-Reviews-E36174.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Holland-and-Barrett-Jobs-E36174.htm",
				Jobs:        "182",
				Reviews:     "2.1K",
				Salaries:    "2.7K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hollandandbarrett",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Engineer (Golang Engineer)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3911059005/",
							Date:             mustDate("2024-05-24"),
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

		// Some | Group Avows
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Group Avows",
			Website: "https://avowstech.com/",
			Careers: "https://avowstech.com/career/",
			About:   "https://avowstech.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3651016,
				Alias:             "group-avows",
				Name:              "Group Avows",
				Followers:         "31K",
				Employees:         "501-1K",
				AssociatedMembers: "267",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AVOWS-Technologies-EI_IE870406.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AVOWS-Technologies-Reviews-E870406.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AVOWS-Technologies-Jobs-E870406.htm",
				Jobs:        "18",
				Reviews:     "77",
				Salaries:    "71",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@avowsgroupofficial9907",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full Stack Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3911007016/",
							Date:             mustDate("2024-05-24"),
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "3 years experience as Golang Developer",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123654018/",
							Date:                 mustDate("2025-01-14"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Cognizant
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cognizant",
			Website: "https://www.cognizant.com/",
			Careers: "https://careers.cognizant.com/global-en/",
			About:   "https://www.cognizant.com/au/en/about-cognizant",
			Blog:    "https://www.cognizant.com/se/en/insights/blog/home-page-se",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1680,
				Alias:             "cognizant",
				Name:              "Cognizant",
				Followers:         "8M",
				Employees:         "10K+",
				AssociatedMembers: "326,260",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cognizant",
				Employees:   "10,000+",
				Salary:      "$50K ~ $220K a year",
				Reviews:     "612",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cognizant",
				Employees: "309,110",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognizant-Technology-Solutions-EI_IE8014.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognizant-Technology-Solutions-Reviews-E8014.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cognizant-Technology-Solutions-Jobs-E8014.htm",
				Jobs:        "2.8K",
				Reviews:     "117K",
				Salaries:    "226K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cognizant-Technology-Solutions",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cognizant",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914839625/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120737001/",
							Date:                 mustDate("2025-01-24"), // mustDate("2025-01-11"),
							WithSalary:           true,                   // The annual salary for this position is between $95K to 115K depending on experience and other qualifications of the successful candidate
							Remote:               false,
						},
						{
							Title:                "Senior AWS developer with Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149258542/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           true, // $55.000 — $114.000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Engineer (Golang/AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144963593/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // salary for this position is between $68,000 – $114,000
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Developer (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4066607632/",
							Date:             mustDate("2024-11-07"),
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Data Engineer — AWS/Hive/Scala/Spark",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124669157/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // $68K — $108K per year
							Remote:               false,
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
				domain.Czechia,
			},
		},

		// Some | Nuro
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nuro",
			Website: "https://www.nuro.ai/",
			Careers: "https://www.nuro.ai/careers",
			About:   "https://www.nuro.ai/company",
			Blog:    "https://medium.com/nuro",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12957486,
				Alias:             "nuro-inc.",
				Name:              "Nuro",
				Followers:         "86K",
				Employees:         "501-1K",
				AssociatedMembers: "930",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nuro-EI_IE1550693.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nuro-Reviews-E1550693.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Nuro",
			YouTubeChannelURL: "https://www.youtube.com/@NuroTeam",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/P5a_50Xb",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | CloudWalk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudWalk",
			Website: "https://cloudwalk.io/",
			Careers: "https://www.cloudwalk.io/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3523168,
				Alias:             "cloudwalk-inc",
				Name:              "CloudWalk, Inc.",
				Followers:         "40K",
				Employees:         "501-1K",
				AssociatedMembers: "601",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cloudwalk",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudwalk",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudWalk-Inc-EI_IE2722308.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudWalk-Inc-Reviews-E2722308.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CloudWalk-Inc-Jobs-E2722308.htm",
				Jobs:        "",
				Reviews:     "282",
				Salaries:    "490",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cloudwalk_shorts",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3917078597/",
							Date:             mustDate("2024-08-26"),
						},
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088770341/",
							Date:             mustDate("2024-12-25"),
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121841670/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior AI-Driven Engineer — Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4129248449/",
							Date:                 mustDate("2025-02-12"), // mustDate("2025-01-20"),
							WithSalary:           false,
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
			Industries:       []domain.Industry{},
		},

		// Some | Transition Technologies PSC
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Transition Technologies PSC",
			Website: "https://ttpsc.com/",
			Careers: "https://kariera.ttpsc.com/en/job-offers/",
			About:   "https://ttpsc.com/en/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17880075,
				Alias:             "transition-technologies-psc",
				Name:              "Transition Technologies PSC",
				Followers:         "6K",
				Employees:         "501-1K",
				AssociatedMembers: "622",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Transition-Technologies-PSC-EI_IE1875542.11,38.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Transition-Technologies-PSC-Reviews-E1875542.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Transition-Technologies-PSC-Jobs-E1875542.htm",
				Jobs:        "",
				Reviews:     "47",
				Salaries:    "83",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TransitionTechnologiesPSC",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3886949034/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4167039797/",
							Date:                 mustDate("2025-02-24"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Kroger
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kroger",
			Website: "https://www.kroger.com/",
			Careers: "https://www.krogerfamilycareers.com/en/sites/CX_2001",
			About:   "https://www.thekrogerco.com/about-kroger/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4914,
				Alias:             "kroger",
				Name:              "Kroger",
				Followers:         "454K",
				Employees:         "10K+",
				AssociatedMembers: "135,606",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Kroger-Technology",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Kroger",
				Employees:   "10,000+",
				Salary:      "$23K ~ $200K a year",
				Reviews:     "72",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kroger",
				Employees: "465,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kroger-EI_IE386.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kroger-Reviews-E386.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kroger-Jobs-E386.htm",
				Jobs:        "9.9K",
				Reviews:     "24K",
				Salaries:    "34K",
				ReviewsRate: "3.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Kroger-54356a75",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@KrogerCo",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Advanced Software Engineer (Java/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3917025739/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Precisely
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Precisely",
			Website: "https://www.precisely.com/",
			Careers: "https://www.precisely.com/careers-and-culture",
			About:   "https://www.precisely.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64863146,
				Alias:             "preciselydata",
				Name:              "Precisely",
				Followers:         "149K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,958",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "PreciselyData",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "precisely",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Precisely-EI_IE3372755.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Precisely-Reviews-E3372755.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Precisely-Jobs-E3372755.htm",
				Jobs:        "24",
				Reviews:     "434",
				Salaries:    "789",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Precisely-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PreciselyData",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3917000992/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | R Systems
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "R Systems",
			Website: "https://www.rsystems.com/",
			Careers: "https://www.rsystems.com/apac/career/",
			About:   "https://www.rsystems.com/about-us/",
			Blog:    "https://eu.rsystems.com/category/blog/tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165636,
				Alias:             "r-systems",
				Name:              "R Systems",
				Followers:         "146K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,830",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "R-Systems",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "r-systems",
				Employees: "6,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-R-Systems-EI_IE32444.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/R-Systems-Reviews-E32444.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/R-Systems-Jobs-E32444.htm",
				Jobs:        "12",
				Reviews:     "1.5K",
				Salaries:    "2K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "R-Systems-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@RSystems_inc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3916555239/",
							Date:             mustDate("2024-05-24"),
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

		// Some | Hays
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hays",
			Website: "https://www.haysplc.com/",
			Careers: "https://www.haysplc.com/joinhays",
			About:   "https://www.haysplc.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3486,
				Alias:             "hays",
				Name:              "Hays",
				Followers:         "8M",
				Employees:         "5K-10K",
				AssociatedMembers: "26,240",
				Verified:          true,
			},
			Ignore: true, // The deleted recruiting agency was added by mistake
		},

		// Some | Consort Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Consort Group",
			Website: "https://consort-group.com/",
			Careers: "https://consort-group.com/en/join-us/our-jobs/",
			About:   "https://consort-group.com/en/consort-group-2/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                46088,
				Alias:             "consortgroup",
				Name:              "Consort Group",
				Followers:         "65K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,666",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Groupe-Consort-EI_IE915503.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Groupe-Consort-Reviews-E915503.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Groupe-Consort-Jobs-E915503.htm",
				Jobs:        "",
				Reviews:     "237",
				Salaries:    "469",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Consort-Group-5",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@consortgroup",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3913831485/",
							Date:             mustDate("2024-05-24"),
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
		},

		// Some | Ascendion
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ascendion",
			Website: "https://ascendion.com/",
			Careers: "https://ascendion.com/careers/",
			About:   "https://ascendion.com/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86694680,
				Alias:             "ascendion",
				Name:              "Ascendion",
				Followers:         "1M",
				Employees:         "5K-10K",
				AssociatedMembers: "3,039",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Ascendion",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ascendion",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ascendion-EI_IE7774544.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ascendion-Reviews-E7774544.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ascendion-Jobs-E7774544.htm",
				Jobs:        "67",
				Reviews:     "549",
				Salaries:    "682",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Ascendion",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ascendioninc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117260547/",
							Date:                 mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Checkout.com
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Checkout.com",
			Website: "https://www.checkout.com/",
			Careers: "https://www.checkout.com/careers",
			About:   "https://www.checkout.com/mission",
			Blog:    "https://medium.com/checkout-com-techblog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3110635,
				Alias:             "checkout",
				Name:              "Checkout.com",
				Followers:         "205K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,956",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "checkout",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Checkoutcom",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "27",
				ReviewsRate: "2.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "checkoutcom",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Checkout-com-EI_IE837487.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Checkout-com-Reviews-E837487.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Checkout-com-Jobs-E837487.htm",
				Jobs:        "10",
				Reviews:     "837",
				Salaries:    "1.8K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Checkout.com",
			},
			OttaProfileSlug:   "Checkout-com",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang) — Payment Performance",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3920739722/",
							Date:             mustDate("2024-06-25"),
						},
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "ProcessOut — Payments",
							URL:              "https://www.linkedin.com/jobs/view/4125751681/",
							Date:             mustDate("2025-02-04"), // mustDate("2025-01-14"),
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

		// Some | Unlimit
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unlimit",
			Website: "https://www.unlimit.com/",
			Careers: "https://www.unlimit.com/careers/",
			About:   "https://www.unlimit.com/about-company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                56459436,
				Alias:             "unlimit-com",
				Name:              "Unlimit",
				Followers:         "138K",
				Employees:         "201-500",
				AssociatedMembers: "559",
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
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Unlimit",
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
								URL:              "https://app.welcometothejungle.com/jobs/P4ept_aQ",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Chime
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chime",
			Website: "https://www.chime.com/",
			Careers: "https://careers.chime.com/",
			About:   "https://www.chime.com/about-us/",
			Blog:    "https://careers.chime.com/en/life-at-chime/engineering-at-chime/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3738695,
				Alias:             "chime-card",
				Name:              "Chime",
				Followers:         "138K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,971",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Chime-EI_IE1493686.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Chime-Reviews-E1493686.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Chime",
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
								URL:              "https://app.welcometothejungle.com/jobs/nKhnfPzD",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Collective Minds Radiology
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Collective Minds Radiology",
			Website: "https://cmrad.com/",
			Careers: "https://careers.cmrad.com/jobs",
			About:   "https://about.cmrad.com/",
			Blog:    "https://about.cmrad.com/articles/tag/technology",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11564593,
				Alias:             "cmrad",
				Name:              "Collective Minds Radiology",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "75",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Collective-Minds-Radiology-EI_IE4686955.11,37.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Collective-Minds-Radiology-Reviews-E4686955.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Collective-Minds-Radiology",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cruise
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cruise",
			Website: "https://www.getcruise.com/",
			Careers: "https://www.getcruise.com/careers/",
			About:   "https://www.getcruise.com/about/",
			Blog:    "https://medium.com/cruise/engineering/home",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6577635,
				Alias:             "getcruise",
				Name:              "Cruise",
				Followers:         "164K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,211",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cruise-automation",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cruise-EI_IE977351.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cruise-Reviews-E977351.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/xpGrXDnd",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Compass
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Compass",
			Website: "https://www.compass.com/",
			Careers: "https://www.compass.com/careers/",
			About:   "https://www.compass.com/about/",
			Blog:    "https://medium.com/compass-true-north/tagged/software-development",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2866215,
				Alias:             "compassinc",
				Name:              "Compass",
				Followers:         "328K",
				Employees:         "1K-5K",
				AssociatedMembers: "29,272",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "UrbanCompass",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Compass-EI_IE719025.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Compass-Reviews-E719025.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Compass-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/EK9Av13p",
								Date:             mustDate(""),
							},
						*/
						{
							Title:                "Golang Backend Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4162630665/",
							Date:                 mustDate("2025-02-25"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Mercury
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mercury",
			Website: "https://mercury.com/",
			Careers: "https://mercury.com/jobs",
			About:   "https://mercury.com/about",
			Blog:    "https://mercury.com/blog/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19107985,
				Alias:             "mercuryhq",
				Name:              "Mercury",
				Followers:         "57K",
				Employees:         "501-1K",
				AssociatedMembers: "1,033",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Mercury",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "12",
				ReviewsRate: "4.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mercury",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mercury-EI_IE3583070.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mercury-Reviews-E3583070.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mercury-Jobs-E3583070.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "200",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Mercury-Technologies",
			},
			OttaProfileSlug:   "Mercury",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Risk Intelligence",
							ShortDescription: "Join a compassionate team of experienced Haskell engineers",
							URL:              "https://www.linkedin.com/jobs/view/4017721524/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
			},
			ShortDescription: "FinTech company",
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
			Name:    "Tailscale",
			Website: "https://tailscale.com/",
			Careers: "https://tailscale.com/careers",
			About:   "https://tailscale.com/company",
			Blog:    "https://tailscale.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35653234,
				Alias:             "tailscale",
				Name:              "Tailscale",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "150",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tailscale",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Tailscale",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tailscale",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tailscale-EI_IE6841860.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tailscale-Reviews-E6841860.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tailscale-Jobs-E6841860.htm",
				Jobs:        "",
				Reviews:     "1",
				Salaries:    "6",
				ReviewsRate: "5.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 73,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer",
							ShortDescription: "Experience with Go is a plus",
							URL:              "https://www.linkedin.com/jobs/view/4039096196/",
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
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Customer.io",
			Website: "https://customer.io/",
			Careers: "https://customer.io/careers",
			About:   "https://customer.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2492471,
				Alias:             "customer-io",
				Name:              "Customer.io",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "373",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "customerio",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Customer-io-EI_IE1308885.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Customer-io-Reviews-E1308885.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Customer-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 22,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/QaTOxLFC",
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
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tabby",
			Website: "https://tabby.ai/",
			Careers: "https://tabby.ai/en-AE/careers",
			About:   "",
			Blog:    "https://insights.tabby.ai/tagged/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26615638,
				Alias:             "tabbypay",
				Name:              "Tabby",
				Followers:         "113K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,048",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tabby",
				Employees: "360",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-tabby-EI_IE6075206.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/tabby-Reviews-E6075206.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/tabby-Jobs-E6075206.htm",
				Jobs:        "104",
				Reviews:     "137",
				Salaries:    "157",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Tabby",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Merchant Activation)",
							ShortDescription: "Strong coding ability in Golang",
							URL:              "https://www.linkedin.com/jobs/view/4012619449/",
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
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sonatus",
			Website: "https://www.sonatus.com/",
			Careers: "https://www.sonatus.com/company/careers/",
			About:   "https://www.sonatus.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18780890,
				Alias:             "sonatus",
				Name:              "Sonatus",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "182",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sonatus-EI_IE3258616.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sonatus-Reviews-E3258616.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Sonatus",
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
								URL:              "https://app.welcometothejungle.com/jobs/0g6xbMY_",
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
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "KOHO",
			Website: "https://www.koho.ca/",
			Careers: "https://www.koho.ca/careers/",
			About:   "https://www.koho.ca/about/",
			Blog:    "https://koho.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9443598,
				Alias:             "koho",
				Name:              "KOHO",
				Followers:         "40K",
				Employees:         "201-500",
				AssociatedMembers: "411",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "kohofinancial",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "koho",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-KOHO-EI_IE2155372.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/KOHO-Reviews-E2155372.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/KOHO-Jobs-E2155372.htm",
				Jobs:        "",
				Reviews:     "166",
				Salaries:    "254",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "KOHO-Financial",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer II (Backend)",
							ShortDescription: "You have a strong understanding of Go, RESTful APIs, Docker and AWS",
							URL:              "https://www.linkedin.com/jobs/view/4032999298/",
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
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nicolab",
			Website: "https://www.nicolab.com/",
			Careers: "https://nicolab1.recruitee.com/",
			About:   "https://www.nicolab.com/our-story/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10873366,
				Alias:             "nico-lab",
				Name:              "Nicolab",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "53",
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
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
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
								URL:              "https://app.welcometothejungle.com/jobs/h0EkrVdy",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
				domain.IndustryMedTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Operant AI",
			Website: "https://www.operant.ai/",
			Careers: "https://www.operant.ai/careers",
			About:   "https://www.operant.ai/company/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67522377,
				Alias:             "operantai",
				Name:              "Operant AI",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "34",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "OperantAI",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "operant-ai",
				Employees: "21",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Operant-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Backend",
							ShortDescription: "1-2+ years experience with Golang / GraphQL / REST / API data modeling",
							URL:              "https://www.linkedin.com/jobs/view/4000951216/",
							Date:             mustDate("2024-08-26"),
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
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CrowdRiff",
			Website: "https://crowdriff.com/",
			Careers: "https://crowdriff.com/careers/",
			About:   "https://crowdriff.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2781710,
				Alias:             "crowd-riff",
				Name:              "CrowdRiff",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "113",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "crowdriff",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "crowdriff",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CrowdRiff-EI_IE1643945.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CrowdRiff-Reviews-E1643945.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CrowdRiff-Jobs-E1643945.htm",
				Jobs:        "2",
				Reviews:     "73",
				Salaries:    "126",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Crowdriff",
			},
			OttaProfileSlug:   "CrowdRiff",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Developer",
							ShortDescription: "You'll primarily work with Go, TypeScript and Node.js",
							URL:              "https://www.linkedin.com/jobs/view/4022030075/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Swiss Post",
			Website: "https://www.post.ch/",
			Careers: "https://www.post.ch/en/jobs/jobs-and-careers",
			About:   "https://www.post.ch/en/jobs/who-we-are",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                338556,
				Alias:             "swiss-post",
				Name:              "Swiss Post",
				Followers:         "111K",
				Employees:         "10K+",
				AssociatedMembers: "9,482",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "swisspost",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Die-Schweizerische-Post-EI_IE12870.11,34.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Die-Schweizerische-Post-Reviews-E12870.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Die-Schweizerische-Post-Jobs-E12870.htm",
				Jobs:        "251",
				Reviews:     "216",
				Salaries:    "550",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077388502/",
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
			ShortDescription: "Postal service",
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
			Name:    "Rialtic",
			Website: "https://www.rialtic.io/",
			Careers: "https://www.rialtic.io/careers",
			About:   "https://www.rialtic.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68821761,
				Alias:             "rialtic-io",
				Name:              "Rialtic",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Rialtic",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rialtic",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rialtic-EI_IE4497416.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rialtic-Reviews-E4497416.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rialtic-Jobs-E4497416.htm",
				Jobs:        "2",
				Reviews:     "16",
				Salaries:    "27",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Rialtic",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4062048996/",
							Date:                 mustDate("2024-11-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144741921/",
							Date:                 mustDate("2025-02-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4146386667/",
							Date:                 mustDate("2025-02-10"),
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
			ShortDescription: "Administrative and financial technology for healthcare",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Prisma",
			Website: "https://www.prisma.io/",
			Careers: "https://www.prisma.io/careers",
			About:   "https://www.prisma.io/about",
			Blog:    "https://www.prisma.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10614939,
				Alias:             "prisma-io",
				Name:              "Prisma",
				Followers:         "15K",
				Employees:         "11-50",
				AssociatedMembers: "138",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "prisma",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Prisma-Data-EI_IE2431237.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Prisma-Data-Reviews-E2431237.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Prisma",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/k6ASPhV7",
								Date:             mustDate(""),
							},
						*/
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/k6ASPhV7",
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
			ShortDescription: "",
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
			Name:    "Zipline",
			Website: "https://www.flyzipline.com/",
			Careers: "https://www.flyzipline.com/careers/",
			About:   "https://www.flyzipline.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7602863,
				Alias:             "flyzipline",
				Name:              "Zipline",
				Followers:         "109K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,315",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zipline-EI_IE1394276.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zipline-Reviews-E1394276.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Zipline",
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
								URL:              "https://app.welcometothejungle.com/jobs/DwReS2t8",
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
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "VisionAI",
			Website: "https://www.visionai.co/",
			Careers: "https://visionai.jobs.personio.de/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18858131,
				Alias:             "wearevisionai",
				Name:              "VisionAI",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "56",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-VisionAI-EI_IE9666263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/VisionAI-Reviews-E9666263.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/VisionAI-Jobs-E9666263.htm",
				Jobs:        "",
				Reviews:     "3",
				Salaries:    "9",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "VisionAI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "(Senior) Rust Backend-Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4020879940/",
							Date:             mustDate("2024-09-27"),
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
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vector Atomic",
			Website: "https://www.vectoratomic.com/",
			Careers: "https://vectoratomic.com/#careers_section",
			About:   "https://vectoratomic.com/#about_section",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18630145,
				Alias:             "vectoratomic",
				Name:              "Vector Atomic",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "49",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vector-Atomic-EI_IE3378320.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vector-Atomic-Reviews-E3378320.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vector-Atomic-Jobs-E3378320.htm",
				Jobs:        "11",
				Reviews:     "2",
				Salaries:    "3",
				ReviewsRate: "5.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Vector-Atomic",
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
							URL:              "https://www.linkedin.com/jobs/view/4023783229/",
							Date:             mustDate("2024-09-25"),
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
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Freeform",
			Website: "https://freeform.co/",
			Careers: "https://freeform.co/careers",
			About:   "https://freeform.co/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19207744,
				Alias:             "freeformfuture",
				Name:              "Freeform",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "freeform",
				Employees: "45",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Freeform-Future-EI_IE8374898.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Freeform-Future-Reviews-E8374898.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Freeform-Future-Jobs-E8374898.htm",
				Jobs:        "27",
				Reviews:     "10",
				Salaries:    "6",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Freeform",
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
							URL:              "https://www.linkedin.com/jobs/view/4023717146/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Metal 3D printing factories",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Daedalean AI",
			Website: "https://daedalean.ai/",
			Careers: "https://daedalean.ai/company/careers",
			About:   "https://www.daedalean.ai/company/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10999325,
				Alias:             "daedalean",
				Name:              "Daedalean AI",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "129",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "daedaleanai",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Daedalean-EI_IE3150803.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Daedalean-Reviews-E3150803.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Daedalean-Jobs-E3150803.htm",
				Jobs:        "9",
				Reviews:     "8",
				Salaries:    "20",
				ReviewsRate: "4.4",
				Verified:    true,
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
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Software Engineer (Data&Tools)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4029289382/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Autonomous flight control for the aircraft",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Statista",
			Website: "https://www.statista.com/",
			Careers: "https://www.statista.com/working-at-statista/",
			About:   "https://www.statista.com/aboutus/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				IDs:               []int{1042846, 42875927},
				Alias:             "statista",
				Name:              "Statista",
				Followers:         "258K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,455",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Statista",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "4",
				ReviewsRate: "2.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "statista",
				Employees: "750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Statista-EI_IE800158.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Statista-Reviews-E800158.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Statista-Jobs-E800158.htm",
				Jobs:        "62",
				Reviews:     "439",
				Salaries:    "738",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Statista-Ltd.",
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
							Title:            "(Senior) Software Engineer (Rust, Backend)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4031949216/",
							Date:             mustDate("2024-09-25"),
						},
						{
							Title:                "Senior Search Engineer (OpenSearch, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134905562/",
							Date:                 mustDate("2025-01-26"),
							WithSalary:           false,
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
			ShortDescription: "Business-data platform",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DroneSense",
			Website: "https://www.dronesense.com/",
			Careers: "https://www.dronesense.com/join-us",
			About:   "https://www.dronesense.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10128253,
				Alias:             "dronesense",
				Name:              "DroneSense, Inc.",
				Followers:         "10K",
				Employees:         "11-50",
				AssociatedMembers: "53",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DroneSense-EI_IE1830147.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DroneSense-Reviews-E1830147.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DroneSense-Jobs-E1830147.htm",
				Jobs:        "4",
				Reviews:     "5",
				Salaries:    "14",
				ReviewsRate: "4.7",
				Verified:    true,
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
							Title:            "Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4035121961/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Drone software platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zurich Instruments",
			Website: "https://www.zhinst.com/",
			Careers: "https://www.zhinst.com/company/careers",
			About:   "https://www.zhinst.com/company/our-story",
			Blog:    "https://www.zhinst.com/blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2289023,
				Alias:             "zurich-instruments-ag",
				Name:              "Zurich Instruments",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "158",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "zhinst",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zurich-Instruments-AG-EI_IE3109985.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zurich-Instruments-AG-Reviews-E3109985.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zurich-Instruments-AG-Jobs-E3109985.htm",
				Jobs:        "10",
				Reviews:     "20",
				Salaries:    "24",
				ReviewsRate: "4.5",
				Verified:    true,
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
							Title:            "Senior Software Engineer Quantum Computing (Python/Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4033619128/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Instrumentation for Quantum Computing and Periodic Signal Measurement",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Palo Alto Networks",
			Website: "https://www.paloaltonetworks.com/",
			Careers: "https://jobs.paloaltonetworks.com/",
			About:   "https://www.paloaltonetworks.com/about-us",
			Blog:    "https://www.paloaltonetworks.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30086,
				Alias:             "palo-alto-networks",
				Name:              "Palo Alto Networks",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "16,701",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "paloaltonetworks",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "palo-alto-networks",
				Employees:   "5,001 to 10,000",
				Salary:      "$39K ~ $278K a year",
				Reviews:     "612",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "palo-alto-networks",
				Employees: "11,410",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Palo-Alto-Networks-EI_IE115142.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Palo-Alto-Networks-Reviews-E115142.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Palo-Alto-Networks-Jobs-E115142.htm",
				Jobs:        "761",
				Reviews:     "2.7K",
				Salaries:    "8.8K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Palo-Alto-Networks",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 46,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Backend (Shared Services)",
							ShortDescription:     "Design and develop robust backend services using 1 or 2 years of expertise Golang, Node.js",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4033692599/",
							Date:                 mustDate("2024-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go/Python Software Engineer",
							ShortDescription:     "Principal (Cortex)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081871661/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer Golang, C",
							ShortDescription:     "Cloud Platform Management",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4159521277/",
							Date:                 mustDate("2025-02-20"),
							WithSalary:           true, // $147.000 - $203.500 per year
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Rust Engineer — Data Classification (Prisma Cloud)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4035556879/",
							Date:             mustDate("2024-11-26"),
						},
						{
							Title:            "Principal Rust Engineer — Data Classification (Prisma Cloud)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4104811687/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:            "Senior Engineer, Rust — Data Classification (Prisma Cloud)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4104810631/",
							Date:             mustDate("2025-02-11"), // mustDate("2025-01-17"),
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
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SciTec",
			Website: "https://scitec.com/",
			Careers: "https://scitec.com/join/",
			About:   "https://scitec.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                22302976,
				Alias:             "scitecinc",
				Name:              "SciTec, Inc.",
				Followers:         "16K",
				Employees:         "201-500",
				AssociatedMembers: "304",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SciTec-EI_IE1000832.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SciTec-Reviews-E1000832.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SciTec-Jobs-E1000832.htm",
				Jobs:        "31",
				Reviews:     "32",
				Salaries:    "86",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Scitec-1",
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
							Title:            "Rust Staff /Senior Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4018720987/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stackx.me",
			Website: "https://stackx.me/",
			Careers: "",
			About:   "https://stackx.me/en/om-oss",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71717406,
				Alias:             "stack-x-me",
				Name:              "Stackx.me",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "12",
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
							Title:            "Senior Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034993832/",
							Date:             mustDate("2024-09-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Social investment app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Walt Disney",
			Website: "https://thewaltdisneycompany.com/",
			Careers: "https://www.disneycareers.com/",
			About:   "https://thewaltdisneycompany.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1292,
				IDs:               []int{1291, 1292, 1294, 1295, 1296, 1301, 1303, 1304, 1305, 2080, 3595, 3598, 3815, 3907, 3908, 3910, 6696, 7827, 7829, 7956, 15817, 51941, 165481, 166865, 255569, 418687, 444310, 513007, 1075902, 1080391, 1097772, 1500141, 2486040, 2525658, 2569065, 2576937, 2739821, 2777915, 3146632, 10288516, 11826960, 18380517, 35634467, 68995044, 73443968, 74544528, 82749060, 103560838, 105582255},
				Alias:             "the-walt-disney-company",
				Name:              "The Walt Disney Company",
				Followers:         "6M",
				Employees:         "10K+",
				AssociatedMembers: "177,388",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "disney",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "the-walt-disney-company",
				Employees:   "10,000+",
				Salary:      "$36K ~ $275K a year",
				Reviews:     "451",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "disney",
				Employees: "223,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Walt-Disney-Company-EI_IE717.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Walt-Disney-Company-Reviews-E717.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Walt-Disney-Company-Jobs-E717.htm",
				Jobs:        "1.4K",
				Reviews:     "19K",
				Salaries:    "33K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "The-Walt-Disney-Company",
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
							Title:                "Senior Software Engineer (Rust Developer)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4071554607/",
							Date:                 mustDate("2024-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer II (Rust Developer)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4088472435/",
							Date:                 mustDate("2024-12-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Software Engineer, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118591892/",
							Date:                 mustDate("2025-01-07"),
							WithSalary:           true, // The hiring range for this position in Seattle, Washington and New York NY is $159,500 to $213,900 per year and in Burbank, CA and Bristol CT is $152,200 to $204,100 per year.
							Remote:               false,
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4063482617/",
							Date:                 mustDate("2024-11-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151374279/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           true, // for this position $145,400 to $195,000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Front-End — Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4156260740/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           true, // $145.500 - $195.000 per year
							Remote:               false,
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
			Name:    "Materialise",
			Website: "https://www.materialise.com/",
			Careers: "https://www.materialise.com/careers/vacancies",
			About:   "https://www.materialise.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4463,
				Alias:             "materialise",
				Name:              "Materialise",
				Followers:         "66K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,800",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "materialise",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Materialise-EI_IE223927.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Materialise-Reviews-E223927.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Materialise-Jobs-E223927.htm",
				Jobs:        "70",
				Reviews:     "377",
				Salaries:    "537",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Materialise",
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
							Title:            "Rust Software Development Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3948655227/",
							Date:             mustDate("2024-11-14"),
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
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "InfluxData",
			Website: "https://www.influxdata.com/",
			Careers: "https://www.influxdata.com/careers/",
			About:   "https://www.influxdata.com/about/",
			Blog:    "https://www.influxdata.com/blog/category/tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5159145,
				Alias:             "influxdb",
				Name:              "InfluxData",
				Followers:         "20K",
				Employees:         "201-500",
				AssociatedMembers: "197",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "influxdata",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "InfluxData-Inc",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "21",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "influxdata",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-InfluxData-EI_IE1402855.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/InfluxData-Reviews-E1402855.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/InfluxData-Jobs-E1402855.htm",
				Jobs:        "8",
				Reviews:     "72",
				Salaries:    "146",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Influxdata",
			},
			OttaProfileSlug:   "InfluxData",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Rust",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3978294129/",
							Date:             mustDate("2024-10-18"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Time series database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "amo",
			Website: "https://amo.co/",
			Careers: "https://amo.co/jobs/",
			About:   "https://amo.co/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92897504,
				Alias:             "amoamoamo",
				Name:              "amo",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "149",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "amo",
				Employees: "60",
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
							Title:            "Software Engineer, Backend (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4038779132/",
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
			ShortDescription:          "Social Networking Platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sony Interactive Entertainment",
			Website: "http://www.sonyinteractive.com/",
			Careers: "https://sonyinteractive.com/careers/",
			About:   "https://sonyinteractive.com/our-company/",
			Blog:    "https://sonyinteractive.com/news/blog/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                57993393,
				Alias:             "sony-interactive-entertainment-llc",
				Name:              "Sony Interactive Entertainment",
				Followers:         "111K",
				Employees:         "10K+",
				AssociatedMembers: "10,714",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sony-Interactive-Entertainment-EI_IE5580180.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sony-Interactive-Entertainment-Reviews-E5580180.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sony-Interactive-Entertainment-Jobs-E5580180.htm",
				Jobs:        "",
				Reviews:     "170",
				Salaries:    "866",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "PlayStation",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust, C++)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4022259692/",
							Date:                 mustDate("2025-02-19"), // mustDate("2024-12-05"),
							WithSalary:           true,                   // $172.100 - $258.100 per year
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
			ShortDescription: "PlayStation",
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
			Name:    "Datadog",
			Website: "https://www.datadoghq.com/",
			Careers: "https://careers.datadoghq.com/",
			About:   "https://www.datadoghq.com/about/",
			Blog:    "https://www.datadoghq.com/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1066442,
				Alias:             "datadog",
				Name:              "Datadog",
				Followers:         "372K",
				Employees:         "1K-5K",
				AssociatedMembers: "7,942",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "datadog",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "datadog",
				Employees:   "5,001 to 10,000 Employees",
				Salary:      "$43K ~ $391K a year",
				Reviews:     "431",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "datadog",
				Employees: "3,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Datadog-EI_IE762009.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Datadog-Reviews-E762009.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Datadog-Jobs-E762009.htm",
				Jobs:        "337",
				Reviews:     "1.4K",
				Salaries:    "3.7K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Datadog",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 259,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Library Software Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4030885124/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Static Analysis Engine (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3967138677/",
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
			ShortDescription: "Datadog provides cloud-scale monitoring and security for metrics, traces and logs in one unified platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,               // system
			Type:    "",              // system
			Name:    "Genius Sports", // "Second Spectrum" somehow related to "Genius Sports"
			Website: "https://www.geniussports.com/",
			Careers: "https://www.geniussports.com/careers/",
			About:   "https://www.geniussports.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                22208562,
				Alias:             "geniussports",
				Name:              "Genius Sports",
				Followers:         "89K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,816",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Genius-Sports",
				Employees:   "1001-5000",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "genius-sports",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Genius-Sports-EI_IE769838.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Genius-Sports-Reviews-E769838.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Genius-Sports-Jobs-E769838.htm",
				Jobs:        "18",
				Reviews:     "434",
				Salaries:    "672",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Genius-Sports",
			},
			OttaProfileSlug:   "Genius-Sports",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4039692138/",
							Date:                 mustDate("2024-10-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4068518142/",
							Date:                 mustDate("2024-12-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075838872/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // salary range for this role is $145,000 – $185,000
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Rust)",
							ShortDescription:     "Realtime Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4091331711/",
							Date:                 mustDate("2025-02-15"),
							WithSalary:           true, // $145.000 - $185.000 per year
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Windmill",
			Website: "https://www.windmill.dev/",
			Careers: "https://www.ycombinator.com/companies/windmill/jobs",
			About:   "https://www.windmill.dev/docs/intro",
			Blog:    "https://www.windmill.dev/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82268299,
				Alias:             "windmill-dev",
				Name:              "Windmill",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "18",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "windmill-labs",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "windmill",
				Employees: "75",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Windmill",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Software Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3992496369/",
							Date:             mustDate("2024-08-26"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "OSS self-hostable developer platform for APIs, background jobs, workflows and UIs",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Fetcherr",
			Website: "https://fetcherr.io",
			Careers: "https://www.fetcherr.io/careers",
			About:   "https://www.fetcherr.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31454791,
				Alias:             "fetcherr-ltd",
				Name:              "Fetcherr",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "116",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fetcherr-EI_IE7854340.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fetcherr-Reviews-E7854340.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fetcherr-Jobs-E7854340.htm",
				Jobs:        "16",
				Reviews:     "5",
				Salaries:    "9",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Fetcherr",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Cloud Service Developer (US)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4043996376/",
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
			ShortDescription:          "AI-driven solutions for the airline industry",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flok Health",
			Website: "https://flok.health/",
			Careers: "https://careers.flok.health/",
			About:   "https://flok.health/#about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84475363,
				Alias:             "flok-health",
				Name:              "Flok Health",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "10",
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
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Flok-Health",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior/Staff Systems Programmer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047041896/",
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
			ShortDescription: "Digital-native healthcare provider",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rhebo",
			Website: "https://rhebo.com/",
			Careers: "https://rhebo.com/company/career/",
			About:   "https://rhebo.com/company/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5312885,
				Alias:             "rhebo",
				Name:              "Rhebo - a Landis+Gyr company",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "32",
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
							Title:            "Senior Software Developer Rust — Focus Network",
							ShortDescription: "You have 3+ years of experience in Rust",
							URL:              "https://www.linkedin.com/jobs/view/4050077229/",
							Date:             mustDate("2024-10-28"),
						},
						{
							Title:            "Senior Software Developer Rust — Focus Network",
							ShortDescription: "You have 3+ years of experience in Rust",
							URL:              "https://www.linkedin.com/jobs/view/4130761747/",
							Date:             mustDate("2025-01-21"),
							Remote:           true,
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
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bjak",
			Website: "https://bjak.my/",
			Careers: "https://travel.cdn.bjak.my/careers",
			About:   "https://bjak.my/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14607755,
				Alias:             "bjak",
				Name:              "Bjak",
				Followers:         "101K",
				Employees:         "501-1K",
				AssociatedMembers: "245",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bjak-EI_IE3055055.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bjak-Reviews-E3055055.htm",
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
							Title:            "Rust Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4046735639/",
							Date:             mustDate("2024-10-28"),
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
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
			Ignore:                    true, // Smart contracts and decentralized applications
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "EverCharge",
			Website: "https://evercharge.com/",
			Careers: "https://evercharge.com/careers",
			About:   "https://evercharge.com/aboutus",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4817778,
				Alias:             "evercharge",
				Name:              "EverCharge",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "111",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "evercharge",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-EverCharge-EI_IE1421791.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/EverCharge-Reviews-E1421791.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/EverCharge-Jobs-E1421791.htm",
				Jobs:        "3",
				Reviews:     "19",
				Salaries:    "42",
				ReviewsRate: "2.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Evercharge-1",
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4040225807/",
							Date:                 mustDate("2024-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081026868/",
							Date:                 mustDate("2024-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114584098/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // The base salary range for this position is between $180,000 - $230,000
							Remote:               false,
						},
						{
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158798582/",
							Date:                 mustDate("2025-02-25"), // mustDate("2025-02-19"),
							WithSalary:           true,                   // $180.000 - $230.000 per year
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
			ShortDescription: "Electric vehicle (EV) charging",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ebury",
			Website: "https://ebury.com/",
			Careers: "https://careers.ebury.com/jobs/",
			About:   "https://ebury.com/about/about-ebury/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                963919,
				Alias:             "eburyfintech",
				Name:              "Ebury",
				Followers:         "75K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,614",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ebury",
				Employees: "1,404",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ebury-Partners-EI_IE823195.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ebury-Partners-Reviews-E823195.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ebury-Partners-Jobs-E823195.htm",
				Jobs:        "208",
				Reviews:     "901",
				Salaries:    "1.4K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Ebury",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3990418832/",
							Date:             mustDate("2025-01-16"), // mustDate("2024-11-16"),
						},
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4132372469/",
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
			ShortDescription: "Ebury is a hyper-growth FinTech firm",
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
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Talon.One",
			Website: "https://talon.one",
			Careers: "https://www.talon.one/jobs",
			About:   "https://www.talon.one/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10393857,
				Alias:             "talon.one",
				Name:              "Talon.One",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "217",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "talon-one",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "talonone",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Talon-One-EI_IE2176357.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Talon-One-Reviews-E2176357.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Talon-One-Jobs-E2176357.htm",
				Jobs:        "10",
				Reviews:     "76",
				Salaries:    "130",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Talon.one",
			},
			OttaProfileSlug:   "Talon-One",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:            "(Senior) Backend Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3945724263/",
							Date:             mustDate("2024-11-25"),
						},
						{
							Title:            "Senior Backend Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4092377724/",
							Date:             mustDate("2025-02-15"), // mustDate("2024-12-18"),
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
			ShortDescription: "Loyalty and promotion engine for enterprises",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Jome",
			Website: "https://jome.com/",
			Careers: "",
			About:   "https://jome.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86332598,
				Alias:             "jomehq",
				PreviousAliases:   []string{"newhomesmate"},
				Name:              "Jome",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "104",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NewHomesMate-EI_IE8164563.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NewHomesMate-Reviews-E8164563.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NewHomesMate-Jobs-E8164563.htm",
				Jobs:        "3",
				Reviews:     "17",
				Salaries:    "14",
				ReviewsRate: "4.1",
				Verified:    true,
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
							Title:            "Senior Backend Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3887627093/",
							Date:             mustDate("2024-04-25"),
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
			ShortDescription: "Platform to search, compare, and purchase newly constructed homes",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Density",
			Website: "https://www.density.io/",
			Careers: "https://www.density.io/careers",
			About:   "https://www.density.io/about-density",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7589398,
				Alias:             "density-inc-",
				Name:              "Density",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "106",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Density",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "density",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Density-EI_IE1627818.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Density-Reviews-E1627818.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Density-Jobs-E1627818.htm",
				Jobs:        "",
				Reviews:     "75",
				Salaries:    "109",
				ReviewsRate: "2.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Density",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Fullstack Engineer",
							ShortDescription: "Lead the development of fullstack web applications, writing clean, efficient, and scalable code in Go, React, and Python, with a focus on both frontend and API development",
							URL:              "https://www.linkedin.com/jobs/view/4047223146/",
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
			ShortDescription:          "Analytics platform for measuring and improving workplaces",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Surfly",
			Website: "https://www.surfly.com",
			Careers: "",
			About:   "https://www.surfly.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1943587,
				Alias:             "surfly",
				Name:              "Surfly",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "30",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "surfly",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Surfly-EI_IE1265203.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Surfly-Reviews-E1265203.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Surfly-Jobs-E1265203.htm",
				Jobs:        "2",
				Reviews:     "17",
				Salaries:    "19",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Surfly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "We're looking for a Senior Software Engineer with expertise in Python, Go (server-side) and vanilla JavaScript (client-side) to join our dynamic team",
							URL:              "https://www.linkedin.com/jobs/view/3931538747/",
							Date:             mustDate("2024-05-24"),
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
			ShortDescription:          "Co-browsing middleware provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cross River",
			Website: "https://www.crossriver.com/",
			Careers: "https://www.crossriver.com/careers",
			About:   "https://www.crossriver.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                967395,
				Alias:             "cross-river-bank",
				Name:              "Cross River",
				Followers:         "28K",
				Employees:         "501-1K",
				AssociatedMembers: "1,411",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cross-River",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "9",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cross-river",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cross-River-Bank-EI_IE1177112.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cross-River-Bank-Reviews-E1177112.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cross-River-Bank-Jobs-E1177112.htm",
				Jobs:        "45",
				Reviews:     "205",
				Salaries:    "284",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cross-River-Bank",
			},
			OttaProfileSlug:   "Cross-River",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4015069840/",
							Date:             mustDate("2024-12-12"),
						},
						{
							Title:            "Senior Software Engineer DevOps Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3847871749/",
							Date:             mustDate("2025-02-13"), // mustDate("2025-01-22"),
							WithSalary:       true,                   // $150k — $180k per year
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
			ShortDescription: "Bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Abnormal Security",
			Website: "https://abnormalsecurity.com/",
			Careers: "https://careers.abnormalsecurity.com/",
			About:   "https://abnormalsecurity.com/about",
			Blog:    "https://abnormalsecurity.com/blog/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18586257,
				Alias:             "abnormalsecurity",
				Name:              "Abnormal Security",
				Followers:         "64K",
				Employees:         "501-1K",
				AssociatedMembers: "945",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "abnormal-security",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Abnormal-Security",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "81",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "abnormal-security",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Abnormal-Security-EI_IE3146005.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Abnormal-Security-Reviews-E3146005.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Abnormal-Security-Jobs-E3146005.htm",
				Jobs:        "14",
				Reviews:     "216",
				Salaries:    "340",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Abnormal-Security",
			},
			OttaProfileSlug:   "Abnormal-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Platform Security Team",
							ShortDescription: "Experience with Golang and Python programming languages",
							URL:              "https://www.linkedin.com/jobs/view/4092851428/",
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
			ShortDescription: "Keeps your email protected",
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
			Name:    "Scope3",
			Website: "https://scope3.com/",
			Careers: "https://scope3.com/careers",
			About:   "https://scope3.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80943372,
				Alias:             "scope3data",
				Name:              "Scope3",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "106",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "scope3data",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "scope3",
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
			OttaProfileSlug:   "Scope3",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Engineer, API & Real Time Systems",
							ShortDescription: "The current tech stack focuses on running high throughput + low latency Golang services deployed to a Kubernetes engine running in Google Cloud Platform",
							URL:              "https://www.linkedin.com/jobs/view/4099234574/",
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
			ShortDescription:          "Sustainability platform for media & advertising",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Apollo GraphQL",
			Website: "https://www.apollographql.com/",
			Careers: "https://www.apollographql.com/careers",
			About:   "https://www.apollographql.com/leadership",
			Blog:    "https://www.apollographql.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28602935,
				Alias:             "apollo-graphql",
				Name:              "Apollo GraphQL",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "243",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "apollographql",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Apollo-GraphQL",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "6",
				ReviewsRate: "2.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "apollo-graphql",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apollo-GraphQL-EI_IE893757.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apollo-GraphQL-Reviews-E893757.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Apollo-GraphQL-Jobs-E893757.htm",
				Jobs:        "18",
				Reviews:     "65",
				Salaries:    "146",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Apollo-Graphql",
			},
			OttaProfileSlug:   "Apollo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4043213870/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           true, // $144.000 - $182.000 per year
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145959230/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           true, // $144.000 - $182.000 per year
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
			ShortDescription: "Deliver GraphQL at any scale",
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
			Name:    "Paessler",
			Website: "https://www.paessler.com/",
			Careers: "https://www.paessler.com/company/career",
			About:   "https://www.paessler.com/company/about-us",
			Blog:    "https://blog.paessler.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                409588,
				Alias:             "paessler-gmbh",
				Name:              "Paessler GmbH",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "367",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "PaesslerAG",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Paessler-GmbH-EI_IE1430882.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Paessler-GmbH-Reviews-E1430882.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Paessler-GmbH-Jobs-E1430882.htm",
				Jobs:        "19",
				Reviews:     "43",
				Salaries:    "94",
				ReviewsRate: "3.1",
				Verified:    true,
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
							Title:                "Back-End Developer Delphi / Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154723829/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Developer Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4039210719/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Developer Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154733084/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
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
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hasura",
			Website: "https://hasura.io/",
			Careers: "https://hasura.io/careers/",
			About:   "https://hasura.io/about/",
			Blog:    "https://hasura.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13211008,
				Alias:             "hasura",
				Name:              "Hasura",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "170",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hasura",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Hasura",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "5.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hasura-EI_IE1451757.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hasura-Reviews-E1451757.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hasura-Jobs-E1451757.htm",
				Jobs:        "",
				Reviews:     "52",
				Salaries:    "55",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Hasura",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 34,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior/Staff Software Engineer — Go Backend",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4026175934/",
							Date:             mustDate("2024-12-20"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior/Staff Software Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4036409356/",
							Date:             mustDate("2024-12-12"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "API platform, built for data delivery",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Aira",
			Website: "https://company.airahome.com/",
			Careers: "https://company.airahome.com/careers",
			About:   "https://company.airahome.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93205887,
				Alias:             "airahome",
				Name:              "Aira",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "662",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aira-EI_IE8976549.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aira-Reviews-E8976549.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aira-Jobs-E8976549.htm",
				Jobs:        "59",
				Reviews:     "30",
				Salaries:    "38",
				ReviewsRate: "4.2",
				Verified:    true,
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
							Title:            "Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4039345477/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Tech Lead / Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075293644/",
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
			ShortDescription:          "Clean energy-tech accessible and affordable for millions of homes",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DDN",
			Website: "https://www.ddn.com/",
			Careers: "https://www.ddn.com/careers/",
			About:   "https://www.ddn.com/about-us/",
			Blog:    "https://www.ddn.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                21665,
				Alias:             "ddn", // before ddn-storage
				Name:              "DDN",
				Followers:         "51K",
				Employees:         "1K-5K",
				AssociatedMembers: "914",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "DDNStorage",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "DataDirect-Networks",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ddn",
				Employees: "600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DDN-EI_IE262933.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DDN-Reviews-E262933.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DDN-Jobs-E262933.htm",
				Jobs:        "82",
				Reviews:     "312",
				Salaries:    "302",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Ddn-7",
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
							Title:            "Senior Development Engineer — RUST",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4064738212/",
							Date:             mustDate("2024-11-02"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Data Intelligence Platform",
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
			Name:    "Starling Bank",
			Website: "https://www.starlingbank.com/",
			Careers: "https://www.starlingbank.com/careers/",
			About:   "https://www.starlingbank.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9180896,
				Alias:             "starlingbank",
				Name:              "Starling Bank",
				Followers:         "266K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,642",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "starlingbank",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Starling-Bank",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "7",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "starling-bank",
				Employees: "1,800",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Starling-Bank-EI_IE1337967.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Starling-Bank-Reviews-E1337967.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Starling-Bank-Jobs-E1337967.htm",
				Jobs:        "52",
				Reviews:     "558",
				Salaries:    "1.4K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Starling-Bank",
			},
			OttaProfileSlug:   "Starling-Bank",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer — Go / IaC (Security)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4013536315/",
							Date:                 mustDate("2024-08-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Cloud Security Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4159614481/",
							Date:                 mustDate("2025-02-21"),
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
			ShortDescription: "Mobile-first bank",
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
			Name:    "Spice AI",
			Website: "https://spice.ai/",
			Careers: "https://spice.ai/careers",
			About:   "https://spice.ai/about-us",
			Blog:    "https://spice.ai/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74148478,
				Alias:             "spice-ai",
				Name:              "Spice AI",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "17",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "spiceai",
				Verified: true,
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
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "Design and implement big data, distributed, and machine learning systems using modern tech stacks in Rust, Golang, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4108994660/",
							Date:                 mustDate("2024-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 29,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "Design and implement big data, distributed, and machine learning systems using modern tech stacks in Rust, Golang, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4108994660/",
							Date:                 mustDate("2024-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4108998246/",
							Date:                 mustDate("2025-02-11"), // mustDate("2025-01-17"),
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
			ShortDescription:          "Blocks for data and time-series AI applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Algolia",
			Website: "https://www.algolia.com/",
			Careers: "https://www.algolia.com/careers",
			About:   "https://www.algolia.com/about",
			Blog:    "https://www.algolia.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2728700,
				Alias:             "algolia",
				Name:              "Algolia",
				Followers:         "56K",
				Employees:         "501-1K",
				AssociatedMembers: "817",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "algolia",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Algolia",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "30",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "algolia",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Algolia-EI_IE998983.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Algolia-Reviews-E998983.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Algolia-Jobs-E998983.htm",
				Jobs:        "18",
				Reviews:     "251",
				Salaries:    "513",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Algolia",
			},
			OttaProfileSlug:   "Algolia",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Backend (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3984452580/",
							Date:             mustDate("2025-02-15"), // mustDate("2024-12-05"),
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
			ShortDescription: "End-to-end AI Search solution",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Observe",
			Website: "https://www.observeinc.com/",
			Careers: "https://www.observeinc.com/careers/",
			About:   "https://www.observeinc.com/about-us/",
			Blog:    "https://www.observeinc.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18796376,
				Alias:             "observe-inc",
				Name:              "Observe, Inc.",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "227",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "observeinc",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Observe",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "observe",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Observe-CA-EI_IE4567528.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Observe-CA-Reviews-E4567528.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Observe-CA-Jobs-E4567528.htm",
				Jobs:        "",
				Reviews:     "51",
				Salaries:    "56",
				ReviewsRate: "4.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff/Principal Software Engineer: Data Backend",
							ShortDescription: "Experience with Go is a plus",
							URL:              "https://www.linkedin.com/jobs/view/3840172575/",
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
			ShortDescription: "SaaS Observability means fewer incidents",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chalk",
			Website: "https://chalk.ai/",
			Careers: "https://chalk.ai/careers",
			About:   "https://chalk.ai/about",
			Blog:    "https://chalk.ai/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                78829224,
				Alias:             "chalkai",
				Name:              "Chalk",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "53",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "chalk",
				Employees: "10",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Chalk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Write Python and Rust",
							URL:              "https://www.linkedin.com/jobs/view/4091287030/",
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
			ShortDescription:          "Data platform for machine learning",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Enova",
			Website: "https://www.enova.com/",
			Careers: "https://www.enova.com/careers/",
			About:   "https://www.enova.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                670584,
				Alias:             "enova-international",
				Name:              "Enova International",
				Followers:         "62K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,401",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "enova",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Enova-International",
				Employees:   "1,001 to 5,000",
				Salary:      "$84K ~ $185K a year",
				Reviews:     "26",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "enova-international",
				Employees: "960",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Enova-EI_IE298072.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Enova-Reviews-E298072.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Enova-Jobs-E298072.htm",
				Jobs:        "35",
				Reviews:     "809",
				Salaries:    "1.5K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Enova-International",
			},
			OttaProfileSlug:   "Enova",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Platform Accounting",
							ShortDescription: "Write clean, maintainable, and efficient code using the Go and Ruby programming languages",
							URL:              "https://www.linkedin.com/jobs/view/3977106723/",
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
			ShortDescription: "Financial credit provider",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JustWatch",
			Website: "https://media.justwatch.com/",
			Careers: "https://media.justwatch.com/careers",
			About:   "https://media.justwatch.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9216347,
				Alias:             "justwatch",
				Name:              "JustWatch",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "225",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "justwatch",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JustWatch-EI_IE1348307.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JustWatch-Reviews-E1348307.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/JustWatch-Jobs-E1348307.htm",
				Jobs:        "1",
				Reviews:     "51",
				Salaries:    "106",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "JustWatch-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Backend Developer",
							ShortDescription: "As part of our team, you'll contribute by coding in Go",
							URL:              "https://www.linkedin.com/jobs/view/4004164547/",
							Date:             mustDate("2024-08-26"),
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
			ShortDescription: "Cross platform streaming guide",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Sourcegraph",
			Website: "https://sourcegraph.com/",
			Careers: "https://sourcegraph.com/jobs",
			About:   "https://sourcegraph.com/about",
			Blog:    "https://sourcegraph.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4803356,
				Alias:             "sourcegraph",
				Name:              "Sourcegraph",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "191",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sourcegraph",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Sourcegraph",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "34",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sourcegraph",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sourcegraph-EI_IE1356770.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sourcegraph-Reviews-E1356770.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sourcegraph-Jobs-E1356770.htm",
				Jobs:        "",
				Reviews:     "110",
				Salaries:    "208",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sourcegraph",
			},
			OttaProfileSlug:   "Sourcegraph",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 190,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Code Search [IC2]",
							ShortDescription: "You are proficient in Go, with exposure to scaling and concurrency",
							URL:              "https://www.linkedin.com/jobs/view/3987576655/",
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
			ShortDescription:          "Universal code search and code intelligence for developers",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Plaid",
			Website: "https://plaid.com/",
			Careers: "https://plaid.com/careers/",
			About:   "https://plaid.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2684737,
				Alias:             "plaid-",
				Name:              "Plaid",
				Followers:         "188K",
				Employees:         "501-1K",
				AssociatedMembers: "1,273",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "plaid",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "plaid",
				Employees:   "201 to 500",
				Salary:      "$66K ~ $264K a year",
				Reviews:     "297",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "plaid",
				Employees: "890",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Plaid-EI_IE1156368.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Plaid-Reviews-E1156368.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Plaid-Jobs-E1156368.htm",
				Jobs:        "69",
				Reviews:     "173",
				Salaries:    "546",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Plaid",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Scaled Access",
							ShortDescription: "Experience with Golang",
							URL:              "https://www.linkedin.com/jobs/view/3972591536/",
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
			ShortDescription: "Infrastructure and APIs for FinTech",
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
			Name:    "Coalition",
			Website: "https://www.coalitioninc.com/",
			Careers: "https://www.coalitioninc.com/jobs",
			About:   "https://www.coalitioninc.com/about",
			Blog:    "https://www.coalitioninc.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11193112,
				Alias:             "coalitioninc",
				Name:              "Coalition, Inc.",
				Followers:         "82K",
				Employees:         "501-1K",
				AssociatedMembers: "718",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Coalition-Inc",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "67",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "coalition",
				Employees: "360",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Coalition-EI_IE1717118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Coalition-Reviews-E1717118.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Coalition-Jobs-E1717118.htm",
				Jobs:        "34",
				Reviews:     "157",
				Salaries:    "247",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Coalition",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Scanning Engine",
							ShortDescription: "Experience with at least one of Python or Golang",
							URL:              "https://www.linkedin.com/jobs/view/4079546853/",
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
			ShortDescription: "Cyber insurance and security for businesses",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mechanical Orchard",
			Website: "https://www.mechanical-orchard.com/",
			Careers: "https://jobs.lever.co/mechanicalorchard",
			About:   "https://www.mechanical-orchard.com/#approach",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89798793,
				Alias:             "mechanical-orchard",
				Name:              "Mechanical Orchard",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "103",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mechanical-orchard",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Mechanical-Orchard",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Mechanical Orchard using Elixir",
							URL:              "https://app.welcometothejungle.com/jobs/h2lbAL5a",
							Date:             mustDate("2024-12-27"), // Approximate date
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "AI testing to help modernize critical systems",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Moveworks",
			Website: "https://www.moveworks.com/",
			Careers: "https://www.moveworks.com/us/en/company/careers",
			About:   "https://www.moveworks.com/us/en/company/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18311685,
				Alias:             "moveworksai",
				Name:              "Moveworks",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "606",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "moveworks",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "moveworks",
				Employees:   "201 to 500 Employees",
				Salary:      "$105K ~ $300K a year",
				Reviews:     "121",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "moveworks",
				Employees: "300",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Moveworks-EI_IE1730936.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Moveworks-Reviews-E1730936.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Moveworks-Jobs-E1730936.htm",
				Jobs:        "64",
				Reviews:     "202",
				Salaries:    "439",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Moveworks",
			},
			OttaProfileSlug:   "Moveworks",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer l, Core Platform",
							ShortDescription: "Experience building and maintaining micro services in Python/Go",
							URL:              "https://www.linkedin.com/jobs/view/3923093175/",
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
			ShortDescription:          "Enterprise-built conversational AI platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Emitwise",
			Website: "https://emitwise.com/",
			Careers: "https://emitwise.com/careers/",
			About:   "https://emitwise.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20511470,
				Alias:             "emitwise",
				Name:              "Emitwise",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "53",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "emitwise",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Emitwise-EI_IE3915102.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Emitwise-Reviews-E3915102.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Emitwise-Jobs-E3915102.htm",
				Jobs:        "3",
				Reviews:     "19",
				Salaries:    "38",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Emitwise",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Full-Stack Engineer",
							ShortDescription: "The majority of our backend is written in Go",
							URL:              "https://www.linkedin.com/jobs/view/4091187613/",
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
			ShortDescription:          "Carbon accounting platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Upvest",
			Website: "https://upvest.co/",
			Careers: "https://careers.upvest.co/",
			About:   "https://upvest.co/about",
			Blog:    "https://engineering.upvest.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18388245,
				Alias:             "upvest",
				Name:              "Upvest",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "180",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "upvestco",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Upvest-EI_IE2948356.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Upvest-Reviews-E2948356.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Upvest-Jobs-E2948356.htm",
				Jobs:        "24",
				Reviews:     "21",
				Salaries:    "38",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Upvest",
			},
			OttaProfileSlug:   "Upvest",
			YouTubeChannelURL: "",
			GoMainLanguage:    true, // https://app.welcometothejungle.com/jobs/JTNWvpsj Work with cutting-edge technologies (Go is the primary language) without a legacy codebase
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Product Engineer — Trading",
							ShortDescription: "Go is the primary language",
							URL:              "https://www.linkedin.com/jobs/view/4064898198/",
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
			ShortDescription: "The Investment API",
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
			Name:    "CoreWeave",
			Website: "https://www.coreweave.com/",
			Careers: "https://www.coreweave.com/careers",
			About:   "https://www.coreweave.com/about-us",
			Blog:    "https://www.coreweave.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                36121341,
				Alias:             "coreweave",
				Name:              "CoreWeave",
				Followers:         "37K",
				Employees:         "501-1K",
				AssociatedMembers: "713",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "coreweave",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "CoreWeave",
				Employees:   "1 to 50",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "4.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "coreweave",
				Employees: "61",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CoreWeave-EI_IE5711823.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CoreWeave-Reviews-E5711823.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CoreWeave-Jobs-E5711823.htm",
				Jobs:        "70",
				Reviews:     "43",
				Salaries:    "56",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 51,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, HPC Network",
							ShortDescription: "Python & Shell scripting are required, Go is a big plus",
							URL:              "https://www.linkedin.com/jobs/view/4029510008/",
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
			ShortDescription: "Specialized cloud provider",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Neo4j",
			Website: "https://www.neo4j.com/",
			Careers: "https://neo4j.com/careers/",
			About:   "https://neo4j.com/company/",
			Blog:    "https://neo4j.com/developer-blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                828370,
				Alias:             "neo4j",
				Name:              "Neo4j",
				Followers:         "85K",
				Employees:         "501-1K",
				AssociatedMembers: "919",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "neo4j",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Neo4J",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "9",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "neo4j",
				Employees: "730",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Neo4j-EI_IE665767.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Neo4j-Reviews-E665767.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Neo4j-Jobs-E665767.htm",
				Jobs:        "46",
				Reviews:     "215",
				Salaries:    "417",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Neo4j",
			},
			OttaProfileSlug:   "Neo4j",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer — Infrastructure",
							ShortDescription:     "Experience using Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4102014529/",
							Date:                 mustDate("2024-12-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (K8s & Go) — GraphQL",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4100001937/",
							Date:                 mustDate("2025-02-15"),
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
			ShortDescription: "Graph-based data analysis",
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
			Name:    "Jack Henry",
			Website: "https://www.jackhenry.com/",
			Careers: "https://careers.jackhenry.com/",
			About:   "https://www.jackhenry.com/who-we-are",
			Blog:    "https://jackhenry.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6970,
				Alias:             "jack-henry",
				Name:              "Jack Henry",
				Followers:         "81K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,818",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "jkhy",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Jack-Henry--Associates",
				Employees:   "5,001 to 10,000",
				Salary:      "$28K ~ $185K a year",
				Reviews:     "21",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "jack-henry-and-associates",
				Employees: "6,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Jack-Henry-and-Associates-EI_IE1543.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Jack-Henry-and-Associates-Reviews-E1543.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Jack-Henry-and-Associates-Jobs-E1543.htm",
				Jobs:        "29",
				Reviews:     "1.3K",
				Salaries:    "2.3K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Jack-Henry-&-Associates-7",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer: Golang/Back-End",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4057290510/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Staff Engineer: Golang/GCP Infrastructure",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075815840//",
							Date:             mustDate("2024-11-16"),
						},
						{
							Title:                "Staff Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142268247/",
							Date:                 mustDate("2025-02-03"),
							WithSalary:           true, // $110.000 — $165.000 per year
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
			ShortDescription: "Financial technology company that strengthens connections between people and their financial institutions",
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
			Name:    "GoCardless",
			Website: "https://gocardless.com/",
			Careers: "https://gocardless.com/about/careers/",
			About:   "https://gocardless.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2808012,
				Alias:             "gocardless",
				Name:              "GoCardless",
				Followers:         "56K",
				Employees:         "501-1K",
				AssociatedMembers: "826",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gocardless",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "GoCardless",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "36",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gocardless",
				Employees: "710",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GoCardless-EI_IE1001543.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GoCardless-Reviews-E1001543.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GoCardless-Jobs-E1001543.htm",
				Jobs:        "14",
				Reviews:     "297",
				Salaries:    "874",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Gocardless",
			},
			OttaProfileSlug:   "GoCardless",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 24,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Development Engineer — Billing & Revenue",
							ShortDescription: "You'll work with Golang, Ruby on Rails, Postgres and Google Cloud Platform",
							URL:              "https://www.linkedin.com/jobs/view/4066058946/",
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
			Name:    "Zalando",
			Website: "https://www.zalando.com/",
			Careers: "https://jobs.zalando.com/",
			About:   "https://corporate.zalando.com/en",
			Blog:    "https://engineering.zalando.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                698916,
				Alias:             "zalando",
				Name:              "Zalando",
				Followers:         "347K",
				Employees:         "10K+",
				AssociatedMembers: "11,355",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "zalando",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Zalando",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "24",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "zalando",
				Employees: "8,340",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zalando-EI_IE613421.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zalando-Reviews-E613421.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zalando-Jobs-E613421.htm",
				Jobs:        "400",
				Reviews:     "3.1K",
				Salaries:    "6.6K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Zalando",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3978516002/",
							Date:                 mustDate("2024-11-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Scala / Kotlin)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4055507332/",
							Date:                 mustDate("2024-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer (Scala/Java)",
							ShortDescription:     "Inspiration & Entertainment",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127911217/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Data Engineer — Scala/Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4126005258/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Data Engineer — Scala/Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130152966/",
							Date:                 mustDate("2025-02-15"), // mustDate("2025-01-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — Java/Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148673191/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Online fashion platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FreeWheel",
			Website: "https://www.freewheel.com/",
			Careers: "https://www.freewheel.com/careers",
			About:   "https://www.freewheel.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                458871,
				Alias:             "freewheel",
				Name:              "FreeWheel",
				Followers:         "42K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,434",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "freewheel",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "freewheel",
				Employees:   "1,001 to 5,000",
				Salary:      "$75K ~ $255K a year",
				Reviews:     "49",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "freewheel",
				Employees: "1,130",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FreeWheel-EI_IE313920.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FreeWheel-Reviews-E313920.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FreeWheel-Jobs-E313920.htm",
				Jobs:        "1.2K",
				Reviews:     "296",
				Salaries:    "592",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Freewheel-4",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "Build with our tech stack and tooling which includes, but is not limited to: Go, gRPC, Redis cache, MySQL Server, AWS services like lambda, EKS, S3, SNS and others, Tilt, Kubernetes, Docker, Cypress, GitHub, CI and TDD practices",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4008715409/",
							Date:                 mustDate("2024-11-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang Cloud Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4004983611/",
							Date:                 mustDate("2025-02-22"), // mustDate("2025-02-17"), // mustDate("2025-01-27"),
							WithSalary:           true,                   // $96.600 — $226.408 per year
							Remote:               false,
						},
						{
							Title:                "Senior Software Development Engineer",
							ShortDescription:     "Golang, React, JavaScript, MySQL",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122185632/",
							Date:                 mustDate("2025-02-22"), // mustDate("2025-02-15"),
							WithSalary:           true,                   // $96.600 - $226.400 per year
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
			ShortDescription: "Platform makes TV and video advertising work",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GeoComply",
			Website: "https://www.geocomply.com/",
			Careers: "https://www.geocomply.com/careers/",
			About:   "https://www.geocomply.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2457172,
				Alias:             "geocomply",
				Name:              "GeoComply",
				Followers:         "20K",
				Employees:         "201-500",
				AssociatedMembers: "534",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "GeoComply",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "GeoComply",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "6",
				ReviewsRate: "2.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "geocomply",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GeoComply-EI_IE2333177.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GeoComply-Reviews-E2333177.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GeoComply-Jobs-E2333177.htm",
				Jobs:        "30",
				Reviews:     "204",
				Salaries:    "321",
				ReviewsRate: "2.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Geocomply",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4026401818/",
							Date:                 mustDate("2024-12-24"), // 2024-11-01
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152516554/",
							Date:                 mustDate("2025-02-14"),
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
			ShortDescription: "Fraud prevention and cybersecurity solutions",
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
			Name:    "Electronic Arts",
			Website: "https://www.ea.com/",
			Careers: "https://www.ea.com/careers",
			About:   "https://www.ea.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1449,
				Alias:             "electronic-arts",
				Name:              "Electronic Arts",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "33,233",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "electronicarts",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "electronic-arts",
				Employees:   "5,001 to 10,000",
				Salary:      "$60K ~ $250K a year",
				Reviews:     "271",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "electronic-arts",
				Employees: "25,900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Electronic-Arts-EI_IE1628.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Electronic-Arts-Reviews-E1628.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Electronic-Arts-Jobs-E1628.htm",
				Jobs:        "369",
				Reviews:     "4.8K",
				Salaries:    "11K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Electronic-Arts",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full Stack Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4092326507/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
	}
}
