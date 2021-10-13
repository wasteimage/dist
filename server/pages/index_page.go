package pages

import (
	"fmt"
	"net/http"
)

const indexPage = "index"

func init() {
	// Index page
	initPages = append(initPages, func(p *Pages) Page {
		var pg page
		pg.name = indexPage
		pg.get = func(rw http.ResponseWriter, r *http.Request) {
			userId := readSession(r)
			var params = map[string]interface{}{
				"loggedIn": userId > 0,
				"pages":    p.GetPagesInfo(),
			}
			err := p.tmpl.Lookup(pg.name).Execute(rw, params)
			if err != nil {
				fmt.Println(err)
			}
		}
		return &pg
	})
}
