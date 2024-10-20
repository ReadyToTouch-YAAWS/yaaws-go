package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart2() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Trusting Social",
			URL:  "https://www.trustingsocial.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3506386,
				Alias: "trustingsocial",
				Name:  "Trusting Social",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrustingSocial-EI_IE741535.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrustingSocial-Reviews-E741535.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4042510486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI fintech company revolutionizing credit scoring using big data technology",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "McAfee",
			URL:  "https://www.mcafee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2336,
				Alias: "mcafee",
				Name:  "McAfee",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mcafee",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McAfee-EI_IE2244.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McAfee-Reviews-E2244.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4041686516/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hadrian",
			URL:  "https://www.hadrian.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71668100,
				Alias: "hadrianspace",
				Name:  "Hadrian",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hadrian-EI_IE5364733.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hadrian-Reviews-E5364733.htm",
			},
			OttaProfileSlug:   "Hadrian-2",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/63inUEkF",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Building autonomous factories for aerospace components",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Scaleway",
			URL:  "https://www.scaleway.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9427185,
				Alias: "scaleway",
				Name:  "Scaleway",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scaleway",
				GoRepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scaleway-EI_IE1861876.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scaleway-Reviews-E1861876.htm",
			},
			OttaProfileSlug:   "Scaleway",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/dFlIR1ow",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cloud provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Demandbase",
			URL:  "https://www.demandbase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    89759,
				Alias: "demandbase",
				Name:  "Demandbase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Demandbase",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Demandbase-EI_IE271272.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Demandbase-Reviews-E271272.htm",
			},
			OttaProfileSlug:   "Demandbase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/sKrHcCx0",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Go-To-Market engagement & marketing software",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Overwolf",
			URL:  "https://www.overwolf.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    519584,
				Alias: "overwolf.com",
				Name:  "Overwolf",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "overwolf",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Overwolf-EI_IE1963582.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Overwolf-Reviews-E1963582.htm",
			},
			OttaProfileSlug:   "Overwolf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/duQ9FWwK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Gaming app & mod development framework",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bitfount",
			URL:  "https://www.bitfount.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69183486,
				Alias: "bitfount",
				Name:  "Bitfount",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bitfount",
				GoRepositoryCount: 0, // Rust 1
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Bitfount",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4044732400/",
					"https://app.welcometothejungle.com/jobs/AGxChIIS",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Federated AI & distributed data science platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hinge",
			URL:  "https://hinge.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2603651,
				Alias: "hinge-app",
				Name:  "Hinge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Hinge",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hinge-EI_IE871901.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hinge-Reviews-E871901.htm",
			},
			OttaProfileSlug:   "Hinge",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4043430809/",
					"https://app.welcometothejungle.com/jobs/V73cck97",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Dating app for meaningful relationships",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Firework",
			URL:  "https://firework.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18700473,
				Alias: "fireworkhq",
				Name:  "Firework",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firework-EI_IE2300706.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firework-Reviews-E2300706.htm",
			},
			OttaProfileSlug:   "Firework",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/srNX_4cO",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Video-based eCommerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Breakroom",
			URL:  "https://www.breakroom.cc/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18458137,
				Alias: "breakroom",
				Name:  "Breakroom",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "breakroom",
				GoRepositoryCount: 0, // Elixir 13
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Breakroom-EI_IE8002334.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Breakroom-Reviews-E8002334.htm",
			},
			OttaProfileSlug:   "Breakroom",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4052987616/",
					"https://app.welcometothejungle.com/jobs/p_OpKzJ8",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "The people-powered job comparison site",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Ansa",
			URL:  "https://www.ansa.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    91158109,
				Alias: "getansa",
				Name:  "Ansa",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Ansa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/EjVpa6gl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Branded customer wallet infrastructure",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "1GLOBAL",
			URL:  "https://www.1global.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39711,
				Alias: "1global",
				Name:  "1GLOBAL",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1GLOBAL-EI_IE9805631.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1GLOBAL-Reviews-E9805631.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048972208/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "eSIM",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Eco-Movement",
			URL:  "https://eco-movement.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15223273,
				Alias: "eco-movement-charge-point-data",
				Name:  "Eco-Movement",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eco-Movement-EI_IE5353422.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eco-Movement-Reviews-E5353422.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4046511967/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Charging station data platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Network Optix",
			URL:  "https://www.networkoptix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1896041,
				Alias: "network-optix",
				Name:  "Network Optix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "networkoptix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Network-Optix-EI_IE1103945.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Network-Optix-Reviews-E1103945.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4028418131/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Convert video data into practical insights for businesses and industries of all kinds",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cinemo",
			URL:  "https://www.cinemo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    564661,
				Alias: "cinemo",
				Name:  "Cinemo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cinemo-EI_IE2189711.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cinemo-Reviews-E2189711.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050001754/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Make every screen an opportunity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Readdle",
			URL:  "https://readdle.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    629551,
				Alias: "readdle",
				Name:  "Readdle",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "readdle",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Readdle-EI_IE613771.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Readdle-Reviews-E613771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3971228819/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Creating apps",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ipsos",
			URL:  "https://www.ipsos.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4318,
				Alias: "ipsos",
				Name:  "Ipsos",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ipsos-EI_IE13063.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ipsos-Reviews-E13063.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053349624/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Market research company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bedrock Streaming",
			URL:  "https://bedrockstreaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64853943,
				Alias: "bedrock-streaming",
				Name:  "Bedrock Streaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bedrockstreaming",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bedrock-Streaming-EI_IE5955934.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bedrock-Streaming-Reviews-E5955934.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4049210489/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Bedrock creates and operates full-scope streaming platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Whalebone",
			URL:  "https://www.whalebone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10961729,
				Alias: "whaleboneio",
				Name:  "Whalebone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "whalebone",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whalebone-EI_IE4889910.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whalebone-Reviews-E4889910.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4007694502/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,  // system
			Type: "", // system
			Name: "OneSignal",
			URL:  "https://onesignal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6424376,
				Alias: "onesignal",
				Name:  "OneSignal",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OneSignal",
				GoRepositoryCount: 5, // Rust 32
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSignal-EI_IE1952551.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSignal-Reviews-E1952551.htm",
			},
			OttaProfileSlug:   "OneSignal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Customer messaging and engagement platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Metabase",
			URL:  "https://www.metabase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6460313,
				Alias: "metabase",
				Name:  "Metabase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "metabase",
				GoRepositoryCount: 0, // Clojure 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Metabase-CA-EI_IE6095700.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Metabase-CA-Reviews-E6095700.htm",
			},
			OttaProfileSlug:   "Metabase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:     nil,
				domain.Rust:   nil,
				domain.Zig:    nil,
				domain.Scala:  nil,
				domain.Elixir: nil,
				domain.Clojure: []string{
					"https://app.welcometothejungle.com/jobs/K21IaTJ2",
				},
				domain.Haskell: nil,
			},
			ShortDescription: "Open-source Business Intelligence tool",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Anjuna Security",
			URL:  "https://www.anjuna.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18438300,
				Alias: "anjuna-security",
				Name:  "Anjuna Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "anjuna-security",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anjuna-Security-EI_IE3150191.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anjuna-Security-Reviews-E3150191.htm",
			},
			OttaProfileSlug:   "Anjuna",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Public-cloud network security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Aqua Security",
			URL:  "https://www.aquasec.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10034420,
				Alias: "aquasecteam",
				Name:  "Aqua Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "aquasecurity",
				GoRepositoryCount: 101,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aqua-Security-Software-EI_IE1785939.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aqua-Security-Software-Reviews-E1785939.htm",
			},
			OttaProfileSlug:   "Aqua-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4045761623/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud native security platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Verve",
			URL:  "https://verve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    108605,
				Alias: "verve-ad-solutions",
				Name:  "Verve",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verve-EI_IE4432646.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verve-Reviews-E4432646.htm",
			},
			OttaProfileSlug:   "Verve",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4015809270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,  // system
			Type: "", // system
			Name: "Wargaming",
			URL:  "https://wargaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    127309,
				Alias: "wargaming-net",
				Name:  "Wargaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "wgnet",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wargaming-EI_IE381713.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wargaming-Reviews-E381713.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048943270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Game developer and publisher",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Napier AI",
			URL:  "https://www.napier.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15197985,
				Alias: "napier",
				Name:  "Napier AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "napier-ai",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Napier-Technologies-EI_IE4370088.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Napier-Technologies-Reviews-E4370088.htm",
			},
			OttaProfileSlug:   "Napier",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4044948110/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI-driven AML solutions transform financial crime compliance from legal obligation to competitive edge",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Kaseya",
			URL:  "https://www.kaseya.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    21377,
				Alias: "kaseya",
				Name:  "Kaseya",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "kaseya",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kaseya-EI_IE262966.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kaseya-Reviews-E262966.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908613207/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Omio",
			URL:  "https://www.omio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2754440,
				Alias: "omio",
				Name:  "Omio",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "omio-labs",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omio-EI_IE2855962.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omio-Reviews-E2855962.htm",
			},
			OttaProfileSlug:   "Omio",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4025460474/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Comparison & booking website for public transport",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Liebherr",
			URL:  "https://www.liebherr.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11057,
				Alias: "liebherr",
				Name:  "Liebherr Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Liebherr-Group-EI_IE787369.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Liebherr-Group-Reviews-E787369.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3821037210/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Construction equipment manufacturer",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Fiserv",
			URL:  "https://www.fiserv.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3364,
				Alias: "fiserv",
				Name:  "Fiserv",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fiserv",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fiserv-EI_IE1384.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fiserv-Reviews-E1384.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4026685486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Gett",
			URL:  "https://www.gett.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1514929,
				Alias: "gettaxi",
				Name:  "Gett",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "GettEngineering",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gett-EI_IE810731.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gett-Reviews-E810731.htm",
			},
			OttaProfileSlug:   "Gett",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3819936157/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "SaaS platform for mobility providers",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stairwell",
			URL:  "https://stairwell.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68606343,
				Alias: "stairwell-inc",
				Name:  "Stairwell",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "stairwell-inc",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Stairwell",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/hnl2J9ws",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Threat detection & cybersecurity platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Solo.io",
			URL:  "https://www.solo.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11074869,
				Alias: "solo.io",
				Name:  "Solo.io",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "solo-io",
				GoRepositoryCount: 82,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-solo-io-EI_IE5382785.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/solo-io-Reviews-E5382785.htm",
			},
			OttaProfileSlug:   "Solo-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/vuvd0QAu",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,  // system
			Type: "", // system
			Name: "Splash Damage",
			URL:  "https://www.splashdamage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    47756,
				Alias: "splash-damage",
				Name:  "Splash Damage",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "splash-damage",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splash-Damage-EI_IE470252.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splash-Damage-Reviews-E470252.htm",
			},
			OttaProfileSlug:   "SplashDamage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/A7daqq7I",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Ditto",
			URL:  "https://ditto.live/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18702497,
				Alias: "dittolive",
				Name:  "Ditto",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "getditto",
				GoRepositoryCount: 6, // Rust 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "ditto",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4050904655/",
					"https://app.welcometothejungle.com/jobs/U3QtZINa",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Data sync platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Pinecone",
			URL:  "https://www.pinecone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20299330,
				Alias: "pinecone-io",
				Name:  "Pinecone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "pinecone-io",
				GoRepositoryCount: 6, // Rust 9
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinecone-Systems-EI_IE6085804.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinecone-Systems-Reviews-E6085804.htm",
			},
			OttaProfileSlug:   "Pinecone",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Beyond Identity",
			URL:  "https://www.beyondidentity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64665737,
				Alias: "beyond-identity-inc",
				Name:  "Beyond Identity",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gobeyondidentity",
				GoRepositoryCount: 2, // Rust 2
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Beyond-Identity-EI_IE3403008.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Beyond-Identity-Reviews-E3403008.htm",
			},
			OttaProfileSlug:   "Beyond-Identity",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/akm3fSOi",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Passwordless identity management solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Nightfall AI",
			URL:  "https://www.nightfall.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19031754,
				Alias: "nightfall-ai",
				Name:  "Nightfall AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "nightfallai",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nightfall-AI-EI_IE3097406.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nightfall-AI-Reviews-E3097406.htm",
			},
			OttaProfileSlug:   "Nightfall-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912841050/",
					"https://app.welcometothejungle.com/jobs/VFnWchKY",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud-native data protection platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Cyberhaven",
			URL:  "https://www.cyberhaven.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10639445,
				Alias: "cyberhaven",
				Name:  "Cyberhaven",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "CyberhavenInc",
				GoRepositoryCount: 13,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cyberhaven-EI_IE2985068.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cyberhaven-Reviews-E2985068.htm",
			},
			OttaProfileSlug:   "Cyberhaven",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/w0d6uKVc",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Fonoa",
			URL:  "https://www.fonoa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    30129363,
				Alias: "fonoa",
				Name:  "Fonoa",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Fonoa-Tech",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fonoa-EI_IE4125781.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fonoa-Reviews-E4125781.htm",
			},
			OttaProfileSlug:   "Fonoa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3960607071/",
					"https://app.welcometothejungle.com/jobs/5Q9cmWeg",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Tax automation for the internet economy",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Payrails",
			URL:  "https://www.payrails.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    75580859,
				Alias: "payrails",
				Name:  "Payrails",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Payrails",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/HCYGHzSQ",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,  // system
			Type: "", // system
			Name: "Teleport",
			URL:  "https://goteleport.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7941233,
				Alias: "go-teleport",
				Name:  "Teleport",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gravitational",
				GoRepositoryCount: 116,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teleport-EI_IE1967263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teleport-Reviews-E1967263.htm",
			},
			OttaProfileSlug:   "Teleport",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4000249330/",
					"https://app.welcometothejungle.com/jobs/vhRdlAui",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Bigblue",
			URL:  "https://www.bigblue.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18776044,
				Alias: "bigblue-co",
				Name:  "Bigblue",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bigblue-EI_IE3379658.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bigblue-Reviews-E3379658.htm",
			},
			OttaProfileSlug:   "Bigblue",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/UU1oL2xl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Developer of eCommerce logistics tools",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Taxfix",
			URL:  "https://taxfix.de/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10791106,
				Alias: "taxfix",
				Name:  "Taxfix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "taxfix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Taxfix-EI_IE1301207.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Taxfix-Reviews-E1301207.htm",
			},
			OttaProfileSlug:   "Taxfix",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/whJEvleb",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Tax filing app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Clerk",
			URL:  "https://clerk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69336355,
				Alias: "clerkinc",
				Name:  "Clerk.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "clerk",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Clerk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Identity tool for React applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ClearScore",
			URL:  "https://clearscore.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9320086,
				Alias: "clearscore",
				Name:  "ClearScore",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ClearScore",
				GoRepositoryCount: 2, // Scala 2
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ClearScore-EI_IE1046600.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ClearScore-Reviews-E1046600.htm",
			},
			OttaProfileSlug:   "ClearScore",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/RmSslgyY",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "OneSchema",
			URL:  "https://www.oneschema.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    74704820,
				Alias: "oneschema",
				Name:  "OneSchema",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "oneschema",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSchema-EI_IE7681431.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSchema-Reviews-E7681431.htm",
			},
			OttaProfileSlug:   "OneSchema",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/YvpdZmyz",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cloud-based data importing and validation platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "StrongDM",
			URL:  "https://www.strongdm.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9446266,
				Alias: "strongdm",
				Name:  "StrongDM",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "strongdm",
				GoRepositoryCount: 15,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-strongDM-EI_IE4514684.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/strongDM-Reviews-E4514684.htm",
			},
			OttaProfileSlug:   "strongDM",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3978025966/",
					"https://app.welcometothejungle.com/jobs/lys1kK5P",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Manages auditable database access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Template
		//{
		//	ID:   0,  // system
		//	Type: "", // system
		//	Name: "",
		//	URL:  "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "",
		//		GoRepositoryCount: 0,
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	GoMainLanguage:    false,
		//	Vacancies: domain.Vacancies{
		//		domain.Go:      []string{},
		//		domain.Rust:    nil,
		//		domain.Zig:     nil,
		//		domain.Scala:   nil,
		//		domain.Elixir:  nil,
		//		domain.Clojure: nil,
		//		domain.Haskell: nil,
		//	},
		//	ShortDescription:          "",
		//	Industries:                []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{},
		//},
	}
}
