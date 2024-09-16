// Code generated by qtc from "organizers-header-profile.qtpl". DO NOT EDIT.
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

func streamorganizersHeaderProfile(qw422016 *qt422016.Writer, profiles []SocialProviderUser) {
	qw422016.N().S(`
<div class="header__profile">
  <button class="header__profile-button">
    <img src="`)
	qw422016.E().S(SocialProviderUserHeaderPhoto(profiles))
	qw422016.N().S(`" alt="photo-profile" />
  </button>
  <div class="header__modal modal-profile">
    <div class="modal-profile__profile">
      <figure class="modal-profile__profile-info">
        <img
          src="`)
	qw422016.E().S(SocialProviderUserHeaderPhoto(profiles))
	qw422016.N().S(`"
          class="modal-profile__user-photo"
          width="48"
          height="48"
          alt="Profile photo"
        />
        <figcaption class="modal-profile__profile-name">`)
	qw422016.E().S(SocialProviderUserName(profiles[0]))
	qw422016.N().S(`</figcaption>
      </figure>
      <div class="modal-profile__profile-list">
        `)
	for _, profile := range profiles {
		qw422016.N().S(`
        `)
		switch profile.SocialProvider {
		case github:
			qw422016.N().S(`
        <a href="https://github.com/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`" target="_blank" class="modal-profile__link">
          <img src="/assets/images/pages/online-new/github-black.svg" width="20" height="20" alt="GitHub logo" />
          <span class="modal-profile__profile-github">github.com/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`</span>
        </a>
        `)
		case gitlab:
			qw422016.N().S(`
        <a href="https://gitlab.com/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`" target="_blank" class="modal-profile__link">
          <img src="/assets/images/pages/online-new/gitlab.svg" width="20" height="20" alt="GitLab logo" />
          <span class="modal-profile__profile-gitlab">gitlab.com/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`</span>
        </a>
        `)
		case bitbucket:
			qw422016.N().S(`
        <a href="https://bitbucket.org/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`" target="_blank" class="modal-profile__link">
          <img src="/assets/images/pages/online-new/bitbucket.svg" width="20" height="20" alt="GitLab logo" />
          <span class="modal-profile__profile-gitlab">bitbucket.org/`)
			qw422016.E().S(profile.Username)
			qw422016.N().S(`</span>
        </a>
        `)
		}
		qw422016.N().S(`
        `)
	}
	qw422016.N().S(`
      </div>
    </div>
    <div class="modal-profile__log-out">
      <a href="/logout" class="modal-profile__button">Log out</a>
    </div>
  </div>
</div>
<script>
  document.querySelector('.header__profile-button').addEventListener('click', () => {
    document.querySelector('.header__modal').classList.toggle('active');
  });
</script>
`)
}

func writeorganizersHeaderProfile(qq422016 qtio422016.Writer, profiles []SocialProviderUser) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamorganizersHeaderProfile(qw422016, profiles)
	qt422016.ReleaseWriter(qw422016)
}

func organizersHeaderProfile(profiles []SocialProviderUser) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeorganizersHeaderProfile(qb422016, profiles)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
