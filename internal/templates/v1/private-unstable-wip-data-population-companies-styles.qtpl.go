// Code generated by qtc from "private-unstable-wip-data-population-companies-styles.qtpl". DO NOT EDIT.
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

func streamdataPopulationCompaniesStyles(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<style>
    body {
        font-family: Arial, sans-serif;
        background-color: #f9f9f9;
        padding: 20px;
        max-width: 800px;
        margin: auto;
    }

    h1 {
        text-align: center;
    }

    ul {
        list-style: none;
        padding: 0;
    }

    li {
        background: #ffffff;
        margin: 10px 0;
        padding: 15px;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .company-name {
        font-size: 1.2em;
        font-weight: bold;
    }

    .links {
        margin-top: 5px;
    }

    .links a {
        color: #0073e6;
        text-decoration: none;
        margin-right: 10px;
    }

    .links a:hover {
        text-decoration: underline;
    }
</style>
`)
}

func writedataPopulationCompaniesStyles(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamdataPopulationCompaniesStyles(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func dataPopulationCompaniesStyles() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writedataPopulationCompaniesStyles(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
