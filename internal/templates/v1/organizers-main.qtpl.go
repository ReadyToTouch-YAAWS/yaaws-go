// Code generated by qtc from "organizers-main.qtpl". DO NOT EDIT.
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

func StreamOrganizersMain(qw422016 *qt422016.Writer, headerProfiles []SocialProviderUser, authQueryParams string) {
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
	streamorganizersMainStyles(qw422016)
	qw422016.N().S(`
</head>

<body class="authorized">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/" class="header__logo">
      <img
        width="129"
        height="32"
        class="header__logo-img"
        src="/assets/images/pages/organizer/golang-organizer.svg"
        alt="organizer logo"
      />
    </a>
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
    <a href="/organizers/golang/welcome`)
		qw422016.E().S(authQueryParams)
		qw422016.N().S(`" class="button button--small-padding button--black header__login-button">Log in</a>
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<section class="hero">
  <div class="container hero__container">
    <div class="hero__overlay">
      <div class="hero__info-wrapper">
        <div class="hero__status">
          <p class="hero__status-text">Service status:</p>
          <span class="badge badge--new">New</span>
          <span class="badge badge--work-in-progress">Work in progress</span>
        </div>
        <h1 class="hero__headline">Direct Contract Job Search Organizer</h1>
        <p class="hero__description">
          A list of useful links to simplify your job search. The basis is a list of companies searching for Golang
          developers.
        </p>
      </div>
      <div class="hero__buttons-group">
        <a href="/organizers/golang" class="button button--bordered-gray hero__button">
          <img
            width="20"
            height="20"
            src="/assets/images/pages/organizer/go-original.svg"
            alt="go original logo"
            class="hero__button-icon"
          />
          Golang Organizer
          <div class="hero__buttons-icon-box">
            <img
              width="12"
              height="12"
              src="/assets/images/pages/common/arrow-small-top-black.svg"
              alt="arrow black icon"
              class="hero__button-icon"
            />
          </div>
        </a>
        <a href="/organizers/rust" class="button button--bordered-gray hero__button">
          <img
            width="20"
            height="20"
            src="/assets/images/pages/organizer/rust.svg"
            alt="go original logo"
            class="hero__button-icon"
          />
          Rust Organizer
          <div class="hero__buttons-icon-box">
            <img
              width="12"
              height="12"
              src="/assets/images/pages/common/arrow-small-top-black.svg"
              alt="arrow black icon"
              class="hero__button-icon"
            />
          </div>
        </a>
        <a href="/organizers/scala" class="button button--bordered-gray hero__button">
          <img
            width="20"
            height="20"
            src="/assets/images/pages/organizer/scala.svg"
            alt="go original logo"
            class="hero__button-icon"
          />
          Scala Organizer
          <div class="hero__buttons-icon-box">
            <img
              width="12"
              height="12"
              src="/assets/images/pages/common/arrow-small-top-black.svg"
              alt="arrow black icon"
              class="hero__button-icon"
            />
          </div>
        </a>
        <a href="/organizers/elixir" class="button button--bordered-gray hero__button">
          <img
            width="20"
            height="20"
            src="/assets/images/pages/organizer/elixir.svg"
            alt="go original logo"
            class="hero__button-icon"
          />
          Elixir Organizer
          <div class="hero__buttons-icon-box">
            <img
              width="12"
              height="12"
              src="/assets/images/pages/common/arrow-small-top-black.svg"
              alt="arrow black icon"
              class="hero__button-icon"
            />
          </div>
        </a>
        <a href="/organizers/clojure" class="button button--bordered-gray hero__button">
          <img
            width="20"
            height="20"
            src="/assets/images/pages/organizer/clojure.svg"
            alt="go original logo"
            class="hero__button-icon"
          />
          Clojure Organizer
          <div class="hero__buttons-icon-box">
            <img
              width="12"
              height="12"
              src="/assets/images/pages/common/arrow-small-top-black.svg"
              alt="arrow black icon"
              class="hero__button-icon"
            />
          </div>
        </a>
      </div>
    </div>
  </div>
</section>

<section class="demo">
  <div class="container demo__container">
    <picture class="demo__picture">
      <source
        width="1216"
        height="560"
        type="image/avif"
        srcset="
          /assets/images/pages/organizer/direct-contract-job-search-organizer.avif,
          /assets/images/pages/organizer/direct-contract-job-search-organizer@2x.avif 2x
        "
      />
      <source
        width="1216"
        height="560"
        type="image/webp"
        srcset="
          /assets/images/pages/organizer/direct-contract-job-search-organizer.webp,
          /assets/images/pages/organizer/direct-contract-job-search-organizer@2x.webp 2x
        "
      />
      <source
        width="1216"
        height="560"
        media="@@minWidth"
        srcset="
          /assets/images/pages/organizer/direct-contract-job-search-organizer.jpg,
          /assets/images/pages/organizer/direct-contract-job-search-organizer@2x.jpg 2x
        "
      />
      <img
        class="demo__image"
        width="1216"
        height="560"
        loading="lazy"
        srcset="/assets/images/pages/organizer/direct-contract-job-search-organizer@2x.jpg 2x"
        src="/assets/images/pages/organizer/direct-contract-job-search-organizer.jpg"
        alt="organizer page screen"
      />
    </picture>
  </div>
</section>

<section class="projects">
  <div class="container projects__container">
    <div class="projects__overlay">
      <div class="projects__about">
        <figure class="projects__readytotouch-logo">
          <img
            width="40"
            height="40"
            class="projects__readytotouch-logo-img"
            src="/assets/images/pages/online/logo.svg"
            alt="logo"
          />
          <h3 class="projects__readytotouch-logo-title">ReadyToTouch</h3>
        </figure>
        <h2 class="projects__headline">ReadyToTouch is a space to help you find jobs</h2>
        <p class="projects__description">Our projects are ReadyToTouch, u8views, Organizer, and others.</p>
        <a
          href="https://readytotouch.com/"
          target="_blank"
          class="button button--small-padding button--black projects__button"
        >
          Check them out
          <div class="organizer-card__try-btn-icon-box">
            <img width="12" height="12" alt="arrow top icon" src="/assets/images/pages/common/arrow-small-top.svg" />
          </div>
        </a>
      </div>
      <div class="projects__group">
        <a class="projects__link" target="_blank" href="https://golang-companies-organizer.readytotouch.com/">
          <img
            class="projects__image"
            width="129"
            height="32"
            alt="organizer logo"
            src="/assets/images/pages/organizer/golang-organizer.svg"
          />
        </a>
        <a class="projects__link" href="https://u8views.com/" target="_blank">
          <img
            class="projects__image"
            width="109"
            height="35"
            alt="u8views logo"
            src="/assets/images/pages/organizer/u8views.svg"
          />
        </a>
        <a class="projects__link" href="https://doutivity.github.io/" target="_blank">
          <img
            class="projects__image"
            width="111"
            height="41"
            alt="doutivity logo"
            src="/assets/images/pages/organizer/doutivity.svg"
          />
        </a>
        <a class="projects__link" href="https://json-to-proto.github.io/" target="_blank">
          <img
            class="projects__image"
            width="157"
            height="18"
            alt="JSON-to-PROTO logo"
            src="/assets/images/pages/organizer/json_to_proto.svg"
          />
        </a>
        <a class="projects__link" href="https://xml-to-go.github.io/" target="_blank">
          <img
            class="projects__image"
            width="117"
            height="18"
            alt="XML-to-GO logo"
            src="/assets/images/pages/organizer/xml_to_go.svg"
          />
        </a>
      </div>
    </div>
  </div>
</section>


</main>
<footer class="footer">
  <div class="footer__overlay">
    <div class="footer__wrapper">
      <div class="footer__group">
        <div class="footer__info">
          <a href="/" class="footer__title">
            <h3 class="footer__title-link">ReadyToTouch</h3>
          </a>
          <p class="footer__subtitle">Space to help you find jobs</p>
          <img width="98"
               height="64"
               loading="lazy"
               class="footer__ukraine-map"
               src="/assets/images/pages/online-new/ukraine-map.svg"
               alt="Map of Ukraine"
          >
        </div>
        <div class="footer__right-side">
          <iframe class="footer__stars"
                  src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch&type=star&count=true&size=large"
                  width="168"
                  height="33"
                  title="GitHub"
          ></iframe>
          <div class="footer__support">
            <a href="https://savelife.in.ua/en/" target="_blank" class="button button--bordered-gray footer__link">
              <figure class="footer__support-figure">
                <img width="49"
                     height="23"
                     src="/assets/images/pages/common-images/come_back_alive.svg"
                     alt="come back alive logo"
                >
                <figcaption class="footer__support-caption">Support</figcaption>
              </figure>
              <div class="footer__support-icon-box">
                <img width="18"
                     height="18"
                     src="/assets/images/pages/common/external-link.svg"
                     alt="external link icon"
                >
              </div>
            </a>
            <a href="https://war.ukraine.ua/" target="_blank" class="button button--bordered-gray footer__link">
              war.ukraine.ua
              <div class="footer__support-icon-box">
                <img width="18"
                     height="18"
                     src="/assets/images/pages/common/external-link.svg"
                     alt="external link icon"
                >
              </div>
            </a>
          </div>
        </div>
      </div>
      <div class="footer__copyrights">
        &copy; 2024 Yaroslav Podorvanov
        <img width="24"
             height="24"
             class="footer__flag-ua"
             src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
             alt="Flag of Ukraine"
        >
      </div>
    </div>
  </div>
</footer>

</body>
</html>
`)
}

func WriteOrganizersMain(qq422016 qtio422016.Writer, headerProfiles []SocialProviderUser, authQueryParams string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersMain(qw422016, headerProfiles, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersMain(headerProfiles []SocialProviderUser, authQueryParams string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersMain(qb422016, headerProfiles, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
