// Code generated by qtc from "footer.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func streamfooter(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<footer class="footer">
    <div class="footer__wrapper wrapper-u8">
        <div class="footer__info">
            <a href="/" class="footer__title">u8views</a>
            <p class="footer__subtitle">Registration statistics on u8views</p>
            <div class="footer__map">
                <img class="footer__map-Ukraine" src="/assets/images/map-of-ukraine.png" alt="Map of Ukraine">
            </div>
            <div class="footer__copyrights">
                <span>© 2023 Yaroslav Podorvanov</span>
                <img class="footer__flag-of-ukraine" src="/assets/images/flag-of-ukraine.svg" alt="Flag Of Ukraine">
            </div>
        </div>
        <div class="footer__support">
            <a href="" class="footer__link">
                <figure class="footer__figure">
                    <img class="footer__img" width="93" height="24" src="/assets/images/national-bank-of-ukraine.png"
                         alt="support">
                    <figcaption class="footer__caption">Support</figcaption>
                    <img src="/assets/images/arrow.svg" alt="arrow">

                </figure>
            </a>
            <a href="" class="footer__link">
                <figure class="footer__figure">
                    <img class="" width="60" height="24" src="/assets/images/come-back-alive-ukraine.png" alt="support">
                    <figcaption class="footer__caption">Support</figcaption>
                    <img src="/assets/images/arrow.svg" alt="arrow">

                </figure>
            </a>
            <a href="" class="footer__link">
                <figure class="footer__figure">
                    <figcaption class="footer__caption">war.ukraine.ua</figcaption>
                    <img src="/assets/images/arrow.svg" alt="arrow">
                </figure>
            </a>
        </div>
    </div>
</footer>
`)
}

func writefooter(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamfooter(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func footer() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writefooter(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
