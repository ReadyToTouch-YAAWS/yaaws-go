package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart6() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hasaki.vn",
			Website: "https://hasaki.vn/",
			Careers: "https://tuyendung.hasaki.vn/viec-lam",
			About:   "https://hotro.hasaki.vn/gioi-thieu-hasaki.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31033180,
				Alias:             "hasaki-vn",
				Name:              "Hasaki.vn",
				Followers:         "9K",
				Employees:         "1K-5K",
				AssociatedMembers: "243",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Back-End Developer (PHP/ Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4109417330/",
							Date:                 mustDate("2024-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160627315/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Back-End Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4164581568/",
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
			ShortDescription: "Retail Health and Personal Care Products company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nutanix",
			Website: "https://www.nutanix.com/",
			Careers: "https://www.nutanix.com/company/careers",
			About:   "https://www.nutanix.com/company",
			Blog:    "https://www.nutanix.dev/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                735085,
				Alias:             "nutanix",
				Name:              "Nutanix",
				Followers:         "521K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,633",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "nutanix",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "nutanix",
				Employees:   "5,001 to 10,000",
				Salary:      "$22K ~ $330K a year",
				Reviews:     "672",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nutanix",
				Employees: "6,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nutanix-EI_IE429159.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nutanix-Reviews-E429159.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nutanix-Jobs-E429159.htm",
				Jobs:        "239",
				Reviews:     "2.5K",
				Salaries:    "5K",
				ReviewsRate: "4.0",
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
							Title:                "Software Engineer (C++/Golang/Systems Development)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4110235695/",
							Date:                 mustDate("2025-02-26"), // mustDate("2025-02-15"), // mustDate("2024-12-26"),
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
			ShortDescription: "Private and hybrid cloud software provider",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Umbra",
			Website: "https://umbra.space/",
			Careers: "https://umbra.space/careers/",
			About:   "https://umbra.space/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18111757,
				Alias:             "umbraspace",
				Name:              "Umbra",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "156",
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
				Alias:     "umbra",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Umbra-Space-EI_IE4502080.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Umbra-Space-Reviews-E4502080.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Umbra-Space-Jobs-E4502080.htm",
				Jobs:        "18",
				Reviews:     "25",
				Salaries:    "14",
				ReviewsRate: "4.8",
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
							Title:                "Senior Spacecraft Flight/Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4105560497/",
							Date:                 mustDate("2024-12-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer — Rust",
							ShortDescription:     "Spacecraft Flight Software",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145412132/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           true, // $150.000 - $200.000 per year
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
			ShortDescription: "Creator of space technology",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cognite",
			Website: "https://www.cognite.com",
			Careers: "https://www.cognite.com/company/careers",
			About:   "https://www.cognite.com/en/company/about-us-cognite",
			Blog:    "https://www.cognite.com/en/blog/main",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18045548,
				Alias:             "cognitedata",
				Name:              "Cognite",
				Followers:         "58K",
				Employees:         "501-1K",
				AssociatedMembers: "657",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cognitedata",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cognite",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cognite",
				Employees: "660",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognite-EI_IE2374321.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognite-Reviews-E2374321.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cognite-Jobs-E2374321.htm",
				Jobs:        "86",
				Reviews:     "155",
				Salaries:    "384",
				ReviewsRate: "3.6",
				Verified:    true,
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
							Title:            "Senior Software Engineer (Rust/Kotlin)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106227833/",
							Date:             mustDate("2024-12-20"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Automates and scales industrial data contextualization",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Europa-Park Resort",
			Website: "https://jobs.europapark.de/en",
			Careers: "https://jobs.europapark.de/en/careers-and-jobs-europa-park-resort",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2325165,
				Alias:             "europa-park-rust-germany-",
				Name:              "Europa-Park Resort",
				Followers:         "24K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,108",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Germany's largest theme park and holiday resort",
			Ignore:           true, // Rust is a municipality in the district of Ortenau in Baden-Württemberg in Germany
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ubisoft",
			Website: "https://www.ubisoft.com/",
			Careers: "https://jobs.ubisoft.com/",
			About:   "https://www.ubisoft.com/en-us/company/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2528,
				Alias:             "ubisoft",
				Name:              "Ubisoft",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "22,599",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ubisoft",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Ubisoft",
				Employees:   "10,000+",
				Salary:      "$50K ~ $206K a year",
				Reviews:     "90",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ubisoft",
				Employees: "21,620",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ubisoft-EI_IE12717.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ubisoft-Reviews-E12717.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ubisoft-Jobs-E12717.htm",
				Jobs:        "337",
				Reviews:     "5.7K",
				Salaries:    "11K",
				ReviewsRate: "3.6",
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
							Title:            "Senior Software Engineer — Rust Backend",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4094489606/",
							Date:             mustDate("2024-12-20"),
						},
						{
							Title:                "Back-End Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127761267/",
							Date:                 mustDate("2025-01-16"),
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
			ShortDescription: "French video game publisher",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Two Six Technologies",
			Website: "https://twosixtech.com/",
			Careers: "https://twosixtech.com/careers/",
			About:   "https://twosixtech.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75453814,
				Alias:             "twosixtechnologies",
				Name:              "Two Six Technologies",
				Followers:         "8K",
				Employees:         "501-1K",
				AssociatedMembers: "659",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "twosixlabs",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Two-Six-Technologies",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "13",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "two-six-technologies",
				Employees: "310",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Two-Six-Technologies-EI_IE4806671.11,31.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Two-Six-Technologies-Reviews-E4806671.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Two-Six-Technologies-Jobs-E4806671.htm",
				Jobs:        "48",
				Reviews:     "55",
				Salaries:    "175",
				ReviewsRate: "4.2",
				Verified:    true,
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
							Title:                "Senior Research Engineer – Rust Development and Formal Methods",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151624125/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           true, // salary range $103,200 - $209,000
							Remote:               false,
						},
						{
							Title:                "Lead Rust Software Engineer",
							ShortDescription:     "Design and implement software solutions using Rust, emphasizing security, performance, and maintainability",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081194563/",
							Date:                 mustDate("2025-02-15"), // mustDate("2024-12-20"),
							WithSalary:           true,                   // $114.800 - $232.600 per year
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
			ShortDescription: "Provides cybersecurity and technology solutions",
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
