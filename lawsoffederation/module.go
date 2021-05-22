package lawsoffederation

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
	if err := m.server.Templates.AddTemplate("lawsoffederation/index"); err != nil {
		return err
	}
	
	if err := m.server.Templates.AddTemplate("lawsoffederation/detail"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/lfn", web.GET, m.indexPage)
	m.server.AddRoute("/lfn/{id}", web.GET, m.detailPage, web.IDParamCtx)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/lfn",
		HyperText: "Laws of Federation",
		Info:      "Laws of Federation",
		Icon:      "feather icon-package",
	})

	return nil
}
