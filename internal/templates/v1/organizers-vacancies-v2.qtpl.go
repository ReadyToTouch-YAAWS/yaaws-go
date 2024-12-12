// Code generated by qtc from "organizers-vacancies-v2.qtpl". DO NOT EDIT.
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

func StreamOrganizersVacanciesV2(qw422016 *qt422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	vacancies []PreparedVacancy,
	userVacancyFavoriteMap map[int64]bool,
	vacancyMonthlyViewsMap map[int64]int64,
	authQueryParams string,
) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` vacancies | ReadyToTouch</title>
	<meta name="title" content="`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` vacancies | ReadyToTouch">
    `)
	qw422016.N().S(`

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="author" type="text/plain" href="https://readytotouch.com/humans.txt"/>
	<meta property="og:image" content="/assets/images/og/organizers-light.jpg">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersVacanciesV2Styles(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
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
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/companies" class="header__nav-link">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/vacancies" class="header__nav-link active">Vacancies</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/communities" class="header__nav-link">Communities</a>
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
  <div class="search search--org-vacancies search--organizer">
    <div class="search__input-group">
      <input class="search__input" id="js-vacancy-query" type="search" name="search" placeholder="Search" list="js-vacancy-query-datalist" />
      <datalist id="js-vacancy-query-datalist">
        `)
	for _, company := range companies {
		qw422016.N().S(`
          <option value="`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`"></option>
        `)
	}
	qw422016.N().S(`
      </datalist>
      <img class="search__icon" alt="Search icon" width="20" height="20" src="/assets/images/pages/common/search.svg" />
    </div>
    `)
	qw422016.N().S(`
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
      <button id="js-criteria-reset" class="button button--light-link search-result__filter-headline-reset" style="visibility: hidden;">Reset all</button>
    </div>

    <div class="filters">
      <!-- Company type -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Company type</h4>
        </header>
		<div class="filters__elements">
          <div class="filters__elements-inner">
            <label class="checkbox filters__element">
              <input class="js-criteria-company-type checkbox__input" type="checkbox" data-alias="product" />
              <span class="checkbox__element"></span>
              Product
            </label>
            <label class="checkbox filters__element">
              <input class="js-criteria-company-type checkbox__input" type="checkbox" data-alias="startup" />
              <span class="checkbox__element"></span>
              Startup
            </label>
          </div>
		</div>
      </div>
      <!-- /Company type  -->

      <!-- Direction -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Industry</h4>
        </header>
		<div class="filters__elements">
          <div class="filters__elements-inner">
            `)
	for _, industry := range industries {
		qw422016.N().S(`
            <label class="checkbox filters__element">
              <input class="js-criteria-company-industry checkbox__input" type="checkbox" data-alias="`)
		qw422016.E().S(industry.Alias)
		qw422016.N().S(`" />
              <span class="checkbox__element"></span>
              `)
		qw422016.E().S(industry.Name)
		qw422016.N().S(`
            </label>
            `)
	}
	qw422016.N().S(`
          </div>
        </div>
      </div>
      <!-- /Direction -->

      <!-- Source -->
      <div class="filters__group">
      	<header class="filters__header">
      		<h4 class="filters__headline">Source</h4>
      	</header>

      	<div class="filters__elements">
      		<div class="filters__elements-inner">
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox"/>
              <span class="checkbox__element"></span>
              LinkedIn
            </label>
            <label class="checkbox filters__element">
              <input class="checkbox__input" type="checkbox"/>
              <span class="checkbox__element"></span>
              Otto
            </label>
      		</div>
      	</div>
      </div>
      <!-- /Source  -->

      <!-- Other -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Other</h4>
        </header>
        <div class="filters__elements">
          <label class="checkbox filters__element">
            <input class="js-criteria-has-employees-from-country checkbox__input" type="checkbox" data-alias="ukraine" />
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
            <input class="js-criteria-has-employees-from-country checkbox__input" type="checkbox" data-alias="czechia" />
            <span class="checkbox__element"></span>
            Has Czechs employees
            <img
              class="checkbox__content-image"
              alt="Flag of Czechia"
              width="24"
              height="24"
              src="/assets/images/pages/online-new/cz_flag.svg"
            />
          </label>
          <label class="checkbox filters__element">
            <input id="js-criteria-in-favorites" class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Favorites
          </label>
          `)
	qw422016.N().S(`
        </div>
      </div>
      <!-- /Other -->

    </div>
  </div>
</aside>
<!-- /filters -->

      <div class="search-result--group">
        <!-- selected filters -->
        <div class="filter-used filter-used--small" style="visibility: hidden;">
          <div class="filter-used__title">Applied filters:</div>
          <ul id="js-vacancy-selected-criteria" class="filter-used__list"></ul>
        </div>
        <!-- /selected filters -->

        <div id="search_result_list" class="search-result__list">
          <p class="search-result-found"><span id="js-result-count" class="search-result-found__amount">`)
	qw422016.N().D(len(vacancies))
	qw422016.N().S(`</span> results</p>
          <!-- card list -->
          <div class="search-result__cards row-gap-8 mt-24">
			`)
	for _, vacancy := range vacancies {
		qw422016.N().S(`
				<div class="js-vacancy card"
					 data-vacancy-id="`)
		qw422016.N().DL(vacancy.ID)
		qw422016.N().S(`"
				>
					<aside class="card__action">
						`)
		if userVacancyFavoriteMap[vacancy.ID] {
			qw422016.N().S(`
						  <button class="js-vacancy-favorite favorite card__action-button button-group__item in-favorite" title="Remove from favorites">
							<svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
							  <path
								d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
							  />
							</svg>
						  </button>
						`)
		} else {
			qw422016.N().S(`
						  <button class="js-vacancy-favorite favorite card__action-button button-group__item" title="Add to favorite">
							<svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
							  <path
								d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
							  />
							</svg>
						  </button>
						`)
		}
		qw422016.N().S(`

						<button title="Hide vacancy" class="button-group__item button-group__item-sloth card__action-button-sloth"></button>
					</aside>
					<figure class="card__header card__header--organizer">
						<div class="card__image-overlay">
							<img class="card__image"
								alt="card image preview icon"
								src="/assets/images/pages/common-images/unknown.svg"
							/>
						</div>
						<figcaption class="card__header-caption">
							<a href="/organizers/v/`)
		qw422016.N().DL(vacancy.ID)
		qw422016.N().S(`" class="card__headline vacancy__link">`)
		qw422016.E().S(vacancy.Title)
		qw422016.N().S(`</a>
							<a href="/organizers/`)
		qw422016.E().S(organizerFeature.Organizer.Alias)
		qw422016.N().S(`/companies/`)
		qw422016.E().S(vacancy.Company.Alias)
		qw422016.N().S(`" class="card__sub-headline vacancy__link">`)
		qw422016.E().S(vacancy.Company.Name)
		qw422016.N().S(`</a>
						</figcaption>
					</figure>
					<p class="card__text card__text--organizer">`)
		qw422016.E().S(vacancy.ShortDescription)
		qw422016.N().S(`</p>
					<div class="card__footer">
						<div class="card__details">
							<figure class="card__figure" title="`)
		qw422016.E().S(formatVacancyDate(vacancy.Date))
		qw422016.N().S(`">
								<img class="card__icon"
									alt="calendar icon"
									width="16"
									height="16"
									src="/assets/images/pages/online/calendar.svg"
								/>
								<figcaption class="card__figcaption">`)
		qw422016.E().S(formatVacancyDiffDate(vacancy.Date))
		qw422016.N().S(`</figcaption>
							</figure>
							<figure class="card__figure">
								<img class="card__icon"
									alt="eye icon"
									width="16"
									height="16"
									src="/assets/images/pages/common/eye.svg"
								/>
								<figcaption class="card__figcaption">Monthly views: `)
		qw422016.N().DL(vacancyMonthlyViewsMap[vacancy.ID])
		qw422016.N().S(`</figcaption>
							</figure>
						</div>
						<a href="/organizers/v/`)
		qw422016.N().DL(vacancy.ID)
		qw422016.N().S(`" class="button button--bordered-gray button--gap-images">
							`)
		if isLinkedInVacancyURL(vacancy.URL) {
			qw422016.N().S(`
								<img
									width="20"
									height="20"
									src="/assets/images/pages/organizer/linkedin.svg"
									alt="linkedin logo"
									class="hero__button-icon"
								/>
							`)
		} else if isOttaVacancyURL(vacancy.URL) {
			qw422016.N().S(`
								<img
									width="20"
									height="20"
									src="/assets/images/pages/organizer/otta.svg"
									alt="otta logo"
									class="hero__button-icon"
								/>
							`)
		}
		qw422016.N().S(`
							View source
							<img
								width="18"
								height="18"
								src="/assets/images/pages/common/external-link.svg"
								alt="arrow black icon"
								class="hero__button-icon"
							/>
						</a>
					</div>
				</div>
			`)
	}
	qw422016.N().S(`
		  </div>
          <!-- /card list -->
        </div>
      </div>
    </div>
  </div>
</div>


</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`
<script src="/assets/js/organizers-vacancies-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersVacanciesV2(qq422016 qtio422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	vacancies []PreparedVacancy,
	userVacancyFavoriteMap map[int64]bool,
	vacancyMonthlyViewsMap map[int64]int64,
	authQueryParams string,
) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersVacanciesV2(qw422016, organizerFeature, headerProfiles, companies, vacancies, userVacancyFavoriteMap, vacancyMonthlyViewsMap, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersVacanciesV2(
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	vacancies []PreparedVacancy,
	userVacancyFavoriteMap map[int64]bool,
	vacancyMonthlyViewsMap map[int64]int64,
	authQueryParams string,
) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersVacanciesV2(qb422016, organizerFeature, headerProfiles, companies, vacancies, userVacancyFavoriteMap, vacancyMonthlyViewsMap, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
