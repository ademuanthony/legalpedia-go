package legalresource

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
	if err := m.server.Templates.AddTemplate("legalresource/index"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/legalresource", web.GET, m.indexPage)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/legalresource",
		HyperText: "Foreign Resources",
		Info:      "Foreign Legal Resources",
		Icon:      "feather icon-package",
	})

	return nil
}
