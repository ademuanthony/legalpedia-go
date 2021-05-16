package legalresource

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

type store interface {
	FindLegalResources(ctx context.Context, req web.PagedResultRequest) (*PagedResponse, error)
}

type module struct {
	db     store
	server *web.Server
}

type LegalResource struct {
	ID          int    `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	URL         string `boil:"Url" json:"Url,omitempty" toml:"Url" yaml:"Url,omitempty"`
	Title       string `boil:"Title" json:"Title,omitempty" toml:"Title" yaml:"Title,omitempty"`
	Description string `boil:"Description" json:"Description,omitempty" toml:"Description" yaml:"Description,omitempty"`
}

type PagedResponse struct {
	Data       []LegalResource
	TotalCount int64
}
