package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart8() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Coherent Solutions",
			Website: "https://www.coherentsolutions.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                38745,
				Alias:             "coherent-solutions",
				Name:              "Coherent Solutions",
				Followers:         "13K",
				Employees:         "1K-5K",
				AssociatedMembers: "680",
				Verified:          true,
			},
			ShortDescription: "Software development company",
			Ignore:           true, // The deleted outsource company was added by mistake
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ATG (Auction Technology Group)",
			Website: "https://www.auctiontechnologygroup.com/",
			Careers: "https://www.auctiontechnologygroup.com/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                565927,
				Alias:             "auctiontechnologygroup",
				Name:              "ATG (Auction Technology Group)",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "509",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "3+ years experience with Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117172166/",
							Date:                 mustDate("2025-01-08"),
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
			ShortDescription: "Online platform for auctions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "gridX",
			Website: "https://www.gridx.ai/",
			Careers: "https://www.gridx.ai/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10599040,
				Alias:             "gridX",
				Name:              "gridX",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "199",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "grid-x",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "Golang for at least 2+ years and preferably have an understanding of IoT that you can apply",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118507935/",
							Date:                 mustDate("2025-01-07"),
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
			ShortDescription: "Developments in the energy industry",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Evolve",
			Website: "https://evolve.com/",
			Careers: "https://evolve.com/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2114845,
				Alias:             "evolve-vacation-rental",
				Name:              "Evolve",
				Followers:         "24K",
				Employees:         "501-1K",
				AssociatedMembers: "1,529",
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
							Title:                "Senior Software Engineer Golang",
							ShortDescription:     "5+ years of backend development experience building applications in Go, NodeJS, and integrations with 3rd Party tools, APIs, and other services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119005072/",
							Date:                 mustDate("2025-01-07"),
							WithSalary:           true, // salary range is $120,000 to $185,000
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
			ShortDescription: "Technology-driven vacation rental company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "StrangeBee",
			Website: "https://strangebee.com/",
			Careers: "https://strangebee.com/careers-at-strangebee/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18993016,
				Alias:             "strangebee",
				Name:              "StrangeBee",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "43",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "StrangeBeeCorp",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121276014/",
							Date:                 mustDate("2025-01-09"),
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
			ShortDescription: "CyberSecurity software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "thatgamecompany",
			Website: "https://thatgamecompany.com/",
			Careers: "https://thatgamecompany.com/careers/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1637083,
				Alias:             "thatgamecompany",
				Name:              "thatgamecompany",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "197",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "thatgameco",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120264764/",
							Date:                 mustDate("2025-01-09"),
							WithSalary:           true, // Salary range is $95,000 USD to $177,000 USD
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
			ShortDescription: "American video game development company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BetterMe",
			Website: "https://betterme.world/",
			Careers: "https://betterme.world/career",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18550246,
				Alias:             "betterme-company",
				Name:              "BetterMe",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "482",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "betterme-dev",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Go + PHP)",
							ShortDescription:     "Design and implement backend features from conception through to deployment, ensuring they meet strategic business objectives",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121587727/",
							Date:                 mustDate("2025-01-10"),
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
			ShortDescription: "Digital fitness company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SPE Solution",
			Website: "https://www.spesolution.com/",
			Careers: "https://www.spesolution.com/career",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30205768,
				Alias:             "spe-solution",
				Name:              "SPE Solution",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "190",
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
							Title:                "Backend Developer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120414743/",
							Date:                 mustDate("2025-01-10"),
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
			ShortDescription: "Indonesian FinTech solution company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quantum Systems",
			Website: "https://quantum-systems.com/",
			Careers: "https://career.quantum-systems.com/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25007323,
				Alias:             "quantum-systems-gmbh",
				Name:              "Quantum Systems",
				Followers:         "37K",
				Employees:         "201-500",
				AssociatedMembers: "284",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "5+ years of professional experience in backend development with a strong focus on real-time systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119548956/",
							Date:                 mustDate("2025-01-10"),
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
			ShortDescription: "German technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "General Dynamics Mission Systems",
			Website: "https://gdmissionsystems.com/",
			Careers: "https://gdmissionsystems.com/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1908,
				Alias:             "gdms",
				Name:              "General Dynamics Mission Systems",
				Followers:         "165K",
				Employees:         "10K",
				AssociatedMembers: "8,822",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Embedded Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118966168/",
							Date:                 mustDate("2025-01-09"),
							WithSalary:           true, // Starting at $108K/yr
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
			ShortDescription: "Integrates secure communication and information systems and technology",
		},

		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Antmicro",
			Website: "https://antmicro.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3267482,
				Alias:             "antmicro-ltd",
				Name:              "Antmicro",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "83",
				Verified:          true,
			},
			ShortDescription: "Sweden software-driven tech company",
			Ignore:           true, // The deleted outsource company was added by mistake
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
		//},
	}
}
