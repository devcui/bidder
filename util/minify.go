/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-06 15:15:02
 * @LastEditTime: 2020-11-11 09:51:31
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/util/minify.go
 * @LICENSE: Apache-2.0
 */
package util

import (
	"io"
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

// Minify ..
func Minify(w io.Writer, r io.Reader) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.Minify("text/html", w, r)
}
