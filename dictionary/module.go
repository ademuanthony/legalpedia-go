package dictionary

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
	if err := m.server.Templates.AddTemplate("dictionary_index"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/lawdictionary", web.GET, m.indexPage)

	// api routes
	// m.server.AddAPIRoute("articles/Filter", web.GET, m.articlesEndpoint)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/lawdictionary",
		HyperText: "Law Dictionary",
		Info:      "Law Dictionary",
		Icon:      "feather icon-package",
	})

	return nil
}
