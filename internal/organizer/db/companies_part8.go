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
							Date:                 mustDate("2025-01-10"), // mustDate("2025-01-10"),
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
							Date:                 mustDate("2025-01-16"), // mustDate("2025-01-10"),
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mondu",
			Website: "https://www.mondu.ai/",
			Careers: "https://www.mondu.ai/careers/",
			About:   "https://www.mondu.ai/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76198623,
				Alias:             "mondu-ai",
				Name:              "Mondu",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "111",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mondu-ai",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior GoLang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121324436/",
							Date:                 mustDate("2025-01-13"),
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
			ShortDescription: "European Fintech offering innovative payment solutions for B2B commerce",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "株式会社マネーフォワード — Money Forward",
			Website: "https://corp.moneyforward.com/en/",
			Careers: "https://recruit.moneyforward.com/en/",
			About:   "https://corp.moneyforward.com/en/aboutus/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5775746,
				Alias:             "money-forward",
				Name:              "株式会社マネーフォワード — Money Forward",
				Followers:         "9K",
				Employees:         "1K-5K",
				AssociatedMembers: "893",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "moneyforward",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4126118648/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // 659,000 JPY — 〜834,000 JPY / month
						},
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PrizePicks",
			Website: "https://prizepicks.com/",
			Careers: "https://www.prizepicks.com/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11477578,
				Alias:             "prize-picks",
				Name:              "PrizePicks",
				Followers:         "21K",
				Employees:         "501-1K",
				AssociatedMembers: "543",
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
							Title:                "Software Engineer III (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124190818/",
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
			ShortDescription: "A daily fantasy sports operator",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CHEP",
			Website: "https://www.chep.com/",
			Careers: "https://www.chep.com/careers/your-career-chep",
			About:   "https://www.chep.com/about-us/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5085,
				Alias:             "chep",
				Name:              "CHEP",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "8,831",
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
							Title:                "Full Stack Developer (Golang and Svelte)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123557217/",
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
			ShortDescription: "Company that specializes in pooling and managing pallets, crates, and containers for efficient and sustainable supply chain operations",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Optimizely",
			Website: "https://www.optimizely.com/",
			Careers: "https://careers.optimizely.com/",
			About:   "https://www.optimizely.com/company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1189697,
				Alias:             "optimizely",
				Name:              "Optimizely",
				Followers:         "140K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,528",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "optimizely",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120974608/",
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
			ShortDescription: "American Digital Experience Platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OPSWAT",
			Website: "https://www.opswat.com/",
			Careers: "https://www.opswat.com/careers",
			About:   "https://www.opswat.com/company/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105936,
				Alias:             "opswat",
				Name:              "OPSWAT",
				Followers:         "106K",
				Employees:         "501-1K",
				AssociatedMembers: "952",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "OPSWAT",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124852341/",
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
			ShortDescription: "American cybersecurity company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VideoAmp",
			Website: "https://videoamp.com/",
			Careers: "https://videoamp.com/jobs/",
			About:   "https://videoamp.com/our-company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3704165,
				Alias:             "videoamp",
				Name:              "VideoAmp",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "611",
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
							Title:                "Senior Full Stack Engineer — HTMX / Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122141025/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           true, // This position has a minimum salary of $150,000- $170,000
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
			ShortDescription: "Media measurement company revolutionizing advertising",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Broadcom",
			Website: "https://www.broadcom.com/",
			Careers: "https://www.broadcom.com/company/careers",
			About:   "https://www.broadcom.com/company/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3072,
				Alias:             "broadcom",
				Name:              "Broadcom",
				Followers:         "498K",
				Employees:         "10K+",
				AssociatedMembers: "60,347",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "broadcom",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127239750/",
							Date:                 mustDate("2025-01-17"),
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
			ShortDescription: "American global technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Toyota North America",
			Website: "toyota.com",
			Careers: "http://www.toyota.com/careers",
			About:   "https://www.toyota.com/brand/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2107,
				Alias:             "toyota-north-america",
				Name:              "Toyota North America",
				Followers:         "435K",
				Employees:         "10K+",
				AssociatedMembers: "19,741",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125175696/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           true, // $60.00 - $77.00/ hr
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
			ShortDescription: "Manufacturing companies",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ABOUT YOU",
			Website: "aboutyou.de",
			Careers: "https://corporate.aboutyou.de/en/career",
			About:   "https://corporate.aboutyou.de/en/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5214302,
				Alias:             "about-you",
				Name:              "ABOUT YOU",
				Followers:         "102K",
				Employees:         "1K-5K",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124350189/",
							Date:                 mustDate("2025-01-16"),
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
			ShortDescription: "German fashion online retailer",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Magicmotorsport",
			Website: "https://www.magicmotorsport.com/",
			Careers: "https://www.magicmotorsport.com/about-us/careers/",
			About:   "https://www.magicmotorsport.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10541609,
				Alias:             "magicmotorsport",
				Name:              "Magicmotorsport",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "128",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Ruby/Elixir Back-End Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119784307/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Development and production of equipment for chip tuning and diagnostics of electronic control units (ECU) of cars",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "STATION F",
			Website: "https://stationf.co/",
			Careers: "https://jobs.stationf.co/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10776313,
				Alias:             "stationf",
				Name:              "STATION F",
				Followers:         "171K",
				Employees:         "11-50",
				AssociatedMembers: "266",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stationf",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Engineer (Elixir/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124020794/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           true, // 55000€ — 69000€
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "French startup campus",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nordea",
			Website: "https://www.nordea.com/",
			Careers: "https://www.nordea.com/en/careers",
			About:   "https://www.nordea.com/en/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2007,
				Alias:             "nordea",
				Name:              "Nordea",
				Followers:         "283K",
				Employees:         "10K+",
				AssociatedMembers: "29,892",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior IT Developer — Scala/Spark/Hadoop/Big Data",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4056511122/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Nordic financial services group",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deutsche Börse",
			Website: "https://www.deutsche-boerse.com/dbg-en/",
			Careers: "https://careers.deutsche-boerse.com/",
			About:   "https://www.deutsche-boerse.com/dbg-en/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163430,
				Alias:             "deutscheboerse",
				Name:              "Deutsche Börse",
				Followers:         "97K",
				Employees:         "10K+",
				AssociatedMembers: "3,934",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Big Data Developer / Engineer (Scala/Java)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124323501/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "German stock exchange",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hootsuite",
			Website: "https://www.hootsuite.com/",
			Careers: "https://careers.hootsuite.com/",
			About:   "https://www.hootsuite.com/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                288540,
				Alias:             "hootsuite",
				Name:              "Hootsuite",
				Followers:         "511K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,668",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Developer, Backend (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123118327/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // $98,400—$137,800 CAD
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Social media management platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HSBC",
			Website: "https://www.hsbc.com/",
			Careers: "https://www.hsbc.com/careers",
			About:   "https://www.hsbc.com/who-we-are",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1241,
				Alias:             "hsbc",
				Name:              "HSBC",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "180,941",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hsbc",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior React and Scala/Java Developer (Investment Banking Tech)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122037400/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "British FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AffiniPay",
			Website: "https://www.affinipay.com/",
			Careers: "https://www.affinipay.com/careers/",
			About:   "https://www.affinipay.com/company/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1276962,
				Alias:             "affinipay",
				Name:              "AffiniPay",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "651",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "affinipay",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Contractor — Senior Back End Developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120259237/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
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
