package formsandprecedents

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

func Activate(ctx context.Context, db store, server *web.Server) error {
	m := module{
		db:     db,
		server: server,
	}

	return m.setupHTTP()
}

func (m *module) setupHTTP() error {
	if err := m.server.Templates.AddTemplate("formspresedents/index"); err != nil {
		return err
	}
	if err := m.server.Templates.AddTemplate("formspresedents/view"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/formspresedents", web.GET, m.listPage)
	m.server.AddRoute("/formspresedents/view/{id}", web.GET, m.viewPage, web.IDParamCtx)

	// api routes
	// m.server.AddAPIRoute("articles/Filter", web.GET, m.articlesEndpoint)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/formspresedents",
		HyperText: "Forms and Presedents",
		Info:      "Forms and Presedents",
		Icon:      "feather icon-package",
	})

	return nil
}
