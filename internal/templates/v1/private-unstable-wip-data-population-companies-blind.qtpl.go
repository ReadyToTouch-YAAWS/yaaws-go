// Code generated by qtc from "private-unstable-wip-data-population-companies-blind.qtpl". DO NOT EDIT.
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

func StreamDataPopulationCompaniesBlind(qw422016 *qt422016.Writer, companies []Company, title string) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
    <title>`)
	qw422016.E().S(title)
	qw422016.N().S(`</title>
    <meta name="description" content="`)
	qw422016.E().S(title)
	qw422016.N().S(`">

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="author" type="text/plain" href="https://readytotouch.com/humans.txt"/>
    <meta property="og:image" content="/assets/images/og/organizers-light.jpg">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
	qw422016.N().S(`
    `)
	streamdataPopulationCompaniesStyles(qw422016)
	qw422016.N().S(`
</head>
<body>
<h1>`)
	qw422016.E().S(title)
	qw422016.N().S(` (`)
	qw422016.N().D(len(companies))
	qw422016.N().S(`)</h1>

<ul>
    `)
	for _, company := range companies {
		qw422016.N().S(`
    <li>
        <div class="company-name">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</div>
        <div class="links">
            <a href='`)
		qw422016.E().S(googleSearchBlind(hostname(company.Website)))
		qw422016.N().S(`' target="_blank">
                <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> site:teamblind.com `)
		qw422016.E().S(hostname(company.Website))
		qw422016.N().S(`
            </a>
            <a href='`)
		qw422016.E().S(googleSearchBlind(company.Name))
		qw422016.N().S(`' target="_blank">
                <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> site:teamblind.com `)
		qw422016.E().S(company.Name)
		qw422016.N().S(`
            </a>
        </div>
    </li>
    `)
	}
	qw422016.N().S(`
</ul>

</body>
</html>
`)
}

func WriteDataPopulationCompaniesBlind(qq422016 qtio422016.Writer, companies []Company, title string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamDataPopulationCompaniesBlind(qw422016, companies, title)
	qt422016.ReleaseWriter(qw422016)
}

func DataPopulationCompaniesBlind(companies []Company, title string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteDataPopulationCompaniesBlind(qb422016, companies, title)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
