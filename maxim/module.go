package maxim

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
	if err := m.server.Templates.AddTemplate("maxims"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/maxims", web.GET, m.indexPage)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/maxims",
		HyperText: "Legal Maxims",
		Info:      "Legal Mxims",
		Icon:      "feather icon-package",
	})

	return nil
}
