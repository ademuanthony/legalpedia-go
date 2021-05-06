package homepage

import (
	"io"
	"net/http"

	"github.com/ademuanthony/legalpedia/web"
)

type Home struct {
	server *web.Server
}

func New(server *web.Server) (*Home, error) {
	hm := &Home{
		server: server,
	}
	err := server.Templates.AddTemplate("home")

	if err != nil {
		return nil, err
	}

	server.AddRoute("/", web.GET, hm.homepage)
	return hm, nil
}

func (hm *Home) homepage(w http.ResponseWriter, r *http.Request) {
	str, err := hm.server.Templates.ExecTemplateToString("home", struct {
		*web.CommonPageData
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: hm.server.CommonData(r),
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Home",
				Active:    true,
			},
		},
	})

	if err != nil {
		log.Errorf("Template execute failure: %v", err)
		hm.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "", web.ExpStatusError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	if _, err = io.WriteString(w, str); err != nil {
		log.Error(err)
	}
}
