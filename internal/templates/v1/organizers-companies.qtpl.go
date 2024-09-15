// Code generated by qtc from "organizers-companies.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOrganizersCompanies(qw422016 *qt422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	authQueryParams string,
) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>Yet another anonymous work search</title>
	<meta name="description" content="Yet another anonymous work search">

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersCompaniesStyles(qw422016)
	qw422016.N().S(`
</head>

<body class="organizer-companies-inner">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/organizers" class="header__logo">
      <img
        width="129"
        height="32"
        class="header__logo-img"
        src="/assets/images/pages/organizer/`)
	qw422016.E().S(organizerFeature.Organizer.Logo)
	qw422016.N().S(`"
        alt="organizer logo"
      />
    </a>
    `)
	var navigation = toFeatureNavigation(organizerFeature.Path)

	qw422016.N().S(`
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.companiesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.companiesActive)
	qw422016.N().S(`">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.vacanciesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.vacanciesActive)
	qw422016.N().S(`">Vacancies</a>
      </li>
    </ul>
    `)
	streamorganizersHeaderStars(qw422016)
	qw422016.N().S(`
    `)
	if len(headerProfiles) > 0 {
		qw422016.N().S(`
    `)
		streamorganizersHeaderProfile(qw422016, headerProfiles)
		qw422016.N().S(`
    `)
	} else {
		qw422016.N().S(`
    <a href="/organizers/`)
		qw422016.E().S(organizerFeature.Organizer.Alias)
		qw422016.N().S(`/welcome`)
		qw422016.E().S(authQueryParams)
		qw422016.N().S(`" class="button button--small-padding button--black header__login-button">Log in</a>
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<div class="container">
  <nav aria-label="breadcrumb" aria-labelledby="navigation through the bread crumbs" class="breadcrumb">
    <ul class="breadcrumb__list">
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/">Main</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/organizers">Organizers</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="javascript:void(0);">`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <span class="breadcrumb__page" aria-current="page">`)
	qw422016.E().S(organizerFeature.Title)
	qw422016.N().S(`</span>
      </li>
    </ul>
  </nav>
</div>

<section class="search-container container">
  <div class="search search--projects search--organizer">
    <div class="search__input-group">
      <input class="search__input" id="company" type="text" placeholder="Search" />
      <img class="search__icon" alt="Search icon" width="20" height="20" src="/assets/images/pages/common/search.svg" />
    </div>
  </div>
</section>

<div class="search-result mt-32">
  <div class="container">
    <div class="search-result__wrapper">
      <!-- filters -->
      <aside class="search-result__filters">
  <div class="search-result__filter-group search-result__filter-group--wide">
    <div class="search-result__filter-header">
      <h2 class="search-result__filter-headline">Filters:</h2>
      <button class="button button--light-link search-result__filter-headline-reset">Reset all</button>
    </div>

    <div class="filters">
      <!-- English level -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Project type</h4>
        </header>
        <div class="filters__elements">
          <label class="checkbox filters__element">
            <input checked class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Product
          </label>
          <label class="checkbox filters__element">
            <input class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Startup
          </label>
        </div>
      </div>
      <!-- /English level -->

      <!-- Direction -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Direction</h4>
        </header>

        <div class="filters__elements">
          <div class="filters__elements-inner">
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              CyberSecurity
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              Embedded
            </label>
            <label class="checkbox filters__element">
              <input checked class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              EdTech
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              eCommerce
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              HealthTech
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              CyberSecurity
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              Embedded
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              EdTech
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              eCommerce
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox" />
              <span class="checkbox__element"></span>
              HealthTech
            </label>
          </div>
        </div>
      </div>
      <!-- /Direction -->

      <!-- Company type -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Other</h4>
        </header>
        <div class="filters__elements">
          <label class="checkbox filters__element">
            <input checked class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Has Ukrainian employees
            <img
              class="checkbox__content-image"
              alt="Flag of Ukraine with waves"
              width="24"
              height="24"
              src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
            />
          </label>
          <label class="checkbox filters__element">
            <input class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Favorites
          </label>
          <div class="switch-wrapper filters__element">
            <label class="switch">
              <input type="checkbox" name="active_search" id="active_search" class="switch__input" />
              <span class="switch__slider switch__slider--round"></span>
            </label>
            <span class="switch-wrapper__text">Active search</span>
            <img
              class="ml-16"
              alt="info icon"
              width="20"
              height="20"
              title="info"
              src="/assets/images/pages/common/info.svg"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</aside>

      <!-- /filters -->
      <div class="search-result--group">
        <!-- selected filters -->
<div id="search_result_filter_used" class="filter-used filter-used--small">
  <div class="filter-used__title">Applied filters:</div>
  <ul class="filter-used__list">
    <li class="filter-used__item">
      Product
      <button class="filter-used__link" type="button">
        <img
          class="filter-used__cross-icon"
          alt="cross icon"
          width="10"
          height="10"
          title="info"
          src="/assets/images/pages/common/cross.svg"
        />
      </button>
    </li>
    <li class="filter-used__item">
      EdTech
      <button class="filter-used__link" type="button">
        <img
          class="filter-used__cross-icon"
          alt="cross icon"
          width="10"
          height="10"
          title="info"
          src="/assets/images/pages/common/cross.svg"
        />
      </button>
    </li>
    <li class="filter-used__item">
      Has Ukrainian employees
      <img
        class="filter-used__content-image"
        alt="Flag of Ukraine with waves"
        width="20"
        height="20"
        src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
      />
      <button class="filter-used__link" type="button">
        <img
          class="filter-used__cross-icon"
          alt="cross icon"
          width="10"
          height="10"
          title="info"
          src="/assets/images/pages/common/cross.svg"
        />
      </button>
    </li>
  </ul>
</div>
<!-- /selected filters -->

        <div id="search_result_list" class="search-result__list">
          <p class="search-result-found"><span class="search-result-found__amount">`)
	qw422016.N().D(len(companies))
	qw422016.N().S(`</span> results</p>
          <!-- card list -->
          <div class="search-result__cards row-gap-8 mt-24">
            `)
	for _, company := range companies {
		qw422016.N().S(`
            <div class="card">
              <aside class="card__action">
                `)
		qw422016.N().S(`

                <button class="favorite in-favorite card__action-button button-group__item" title="Remove from favorites">
                  <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
                    <path
                      d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
                    />
                  </svg>
                </button>
                <button class="button-group__item" title="View statistics">
                  <img width="20" height="20" alt="icon stats" src="/assets/images/pages/common/stats.svg" />
                </button>
              </aside>
              <figure class="card__header">
                <div class="card__image-overlay card__image-overlay--small">
                  <img
                    width="18"
                    height="18"
                    class="card__image card__image--preview"
                    alt="card image preview icon"
                    src="/assets/images/pages/common/image-preview.svg"
                  />
                </div>
                <figcaption class="card__header-caption">
                  `)
		qw422016.N().S(`
                  <a href="`)
		qw422016.E().S(company.URL)
		qw422016.N().S(`" class="card__headline vacancy__link">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</a>
                </figcaption>
              </figure>
              <div class="card__info">
                <figure class="card__figure">
                  <img
                    class="card__icon"
                    alt="card type icon"
                    width="16"
                    height="16"
                    src="/assets/images/pages/vacancy/building.svg"
                  />
                  <figcaption class="card__figcaption">Product|Startup</figcaption>
                </figure>
                <figure class="card__figure">
                  <img
                    class="card__icon"
                    alt="card type icon"
                    width="16"
                    height="16"
                    src="/assets/images/pages/vacancy/company-type.svg"
                  />
                  <figcaption class="card__figcaption">TodoTech</figcaption>
                </figure>
              </div>
              <p class="card__text">
                `)
		qw422016.N().S(`
              </p>
              <div class="card__links">
                <ul class="card__links-group">
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="linkedin icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/linkedin.svg"
                    />
                    <a href="https://www.linkedin.com/company/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" class="button-link card__links-link">LinkedIn</a>
                  </li>
                  <li class="card__links-item">
                    <a href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, nil))
		qw422016.N().S(`" class="button-link card__links-link">Connections</a>
                  </li>
                  <li class="card__links-item">
                    `)
		qw422016.N().S(`
                    <a href="`)
		qw422016.E().S(linkedinJobsURL([]Company{company}, golangKeywordsTitles))
		qw422016.N().S(`" class="button-link card__links-link">Jobs</a>
                  </li>
                </ul>
`)
		if company.GitHubProfile.Login == "" {
			qw422016.N().S(`                  <ul class="card__links-group">
                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="github icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/github.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchGitHub(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">GitHub</span>
                    </li>
                  </ul>
                `)
		} else {
			qw422016.N().S(`
                  <ul class="card__links-group">
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="github icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/github.svg"
                      />
                      <a href="https://github.com/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`" class="button-link card__links-link">GitHub</a>
                    </li>
                    <li class="card__links-item">
                      `)
			qw422016.N().S(`
                      <a href="https://github.com/orgs/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`/repositories?q=lang:go" class="button-link card__links-link" title="`)
			qw422016.N().D(company.GitHubProfile.GoRepositoryCount)
			qw422016.N().S(` repositories">Repositories</a>
                    </li>
                  </ul>
`)
		}
		if company.GlassdoorProfile.OverviewURL == "" {
			qw422016.N().S(`                  <ul class="card__links-group">
                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="glassdoor icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/glassdoor.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchGlassdoor(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">Glassdoor</span>
                    </li>
                  </ul>
                `)
		} else {
			qw422016.N().S(`
                  <ul class="card__links-group">
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="glassdoor icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/glassdoor.svg"
                      />
                      <a href="`)
			qw422016.E().S(company.GlassdoorProfile.OverviewURL)
			qw422016.N().S(`" class="button-link card__links-link">Glassdoor</a>
                    </li>
                    <li class="card__links-item">
                      <a href="`)
			qw422016.E().S(company.GlassdoorProfile.ReviewsURL)
			qw422016.N().S(`" class="button-link card__links-link">Reviews</a>
                    </li>
                  </ul>
`)
		}
		qw422016.N().S(`                <ul class="card__links-group card__links-group--unbordered">
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="similarweb icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/similarweb.svg"
                    />
                    <a href="`)
		qw422016.E().S(similarwebURL(company.URL))
		qw422016.N().S(`" class="button-link card__links-link">SimilarWeb</a>
                  </li>
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="whois icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/whois.svg"
                    />
                    <a href="`)
		qw422016.E().S(whoisURL(company.URL))
		qw422016.N().S(`" class="button-link card__links-link">Whois</a>
                  </li>

                  `)
		qw422016.N().S(`
                  <li class="card__links-item card__links-item--disabled">
                    <img
                      class="card__links-icon"
                      alt="xing icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/xing.svg"
                    />
                    <a href="`)
		qw422016.E().S(googleSearchXing(company.Name))
		qw422016.N().S(`" class="card__links-link card__links-link--google">
                      <img
                        class="card__links-icon card__links-icon--google"
                        alt="google icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/google.svg"
                      />
                    </a>
                    <span class="button-link card__links-link">XING</span>
                  </li>

`)
		if company.OttaProfileSlug == "" {
			qw422016.N().S(`                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="otta icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/otta.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchOtta(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">Otta</span>
                    </li>
                  `)
		} else {
			qw422016.N().S(`
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="otta icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/otta.svg"
                      />
                      <a href="https://app.otta.com/companies/`)
			qw422016.E().S(company.OttaProfileSlug)
			qw422016.N().S(`" class="button-link card__links-link">Otta</a>
                    </li>
`)
		}
		qw422016.N().S(`                </ul>
              </div>
            </div>
            `)
	}
	qw422016.N().S(`
          </div>

          <!-- /card list -->
          <!-- search footer -->
          <footer class="search-result__footer">
            <button class="button button--bordered-black-transparent" type="button">More</button>

            <!-- pagination -->
            <nav class="pagination search-result__pagination">
  <a class="pagination__button pagination__button--prev" href="javascript:" aria-label="arrow left">
    <img
      class="pagination__icon pagination__icon--prev"
      width="12"
      height="12"
      src="/assets/images/pages/common/double-left.svg"
      alt="arrow left"
    />
  </a>
  <a class="pagination__button pagination__button--prev" href="javascript:" aria-label="arrow left">
    <img
      class="pagination__icon pagination__icon--prev"
      width="14"
      height="14"
      src="/assets/images/pages/common/left.svg"
      alt="arrow left"
    />
  </a>
  <a class="pagination__item pagination__item--active" href="javascript:">1</a>
  <a class="pagination__item" href="javascript:">2</a>
  <a class="pagination__item" href="javascript:">3</a>
  <a class="pagination__item" href="javascript:">4</a>
  <a class="pagination__item" href="javascript:">5</a>
  <span class="pagination__item">...</span>
  <a class="pagination__item" href="javascript:">9</a>
  <a class="pagination__button pagination__button--prev" href="javascript:" aria-label="arrow left">
    <img
      class="pagination__icon pagination__icon--prev"
      width="14"
      height="14"
      src="/assets/images/pages/common/right.svg"
      alt="arrow left"
    />
  </a>
  <a class="pagination__button pagination__button--next" href="javascript:" aria-label="arrow right">
    <img
      class="pagination__icon pagination__icon--next"
      width="12"
      height="12"
      src="/assets/images/pages/common/double-left.svg"
      alt="arrow right"
    />
  </a>
</nav>

            <!-- /pagination -->
          </footer>
          <!-- /search footer -->
        </div>
      </div>
    </div>
  </div>
</div>


</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`

</body>
</html>
`)
}

func WriteOrganizersCompanies(qq422016 qtio422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	authQueryParams string,
) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersCompanies(qw422016, organizerFeature, headerProfiles, companies, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersCompanies(
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	authQueryParams string,
) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersCompanies(qb422016, organizerFeature, headerProfiles, companies, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
