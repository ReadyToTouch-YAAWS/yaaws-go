// Code generated by qtc from "organizers-footer.qtpl". DO NOT EDIT.
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

func streamorganizersFooter(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
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
        &copy; 2025 Yaroslav Podorvanov
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
`)
}

func writeorganizersFooter(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamorganizersFooter(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func organizersFooter() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeorganizersFooter(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
