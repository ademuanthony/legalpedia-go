package rulesofcourt

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

type module struct {
	db     store
	server *web.Server
}

type store interface {
	GetAll(context.Context, GetAllInput) (*PagedRuleResponse, error)
	GetNames(ctx context.Context, ruleType string) ([]string, error)
	State(context.Context, GetAllInput) (*PagedRuleResponse, error)
	Other(context.Context, GetAllInput) (*PagedRuleResponse, error)
	Detail(ctx context.Context, id int) (*Rule, error)
}

type Rule struct {
	ID        int    `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	UUID      string `boil:"uuid" json:"uuid,omitempty" toml:"uuid" yaml:"uuid,omitempty"`
	Name      string `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Title     string `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Section   string `boil:"section" json:"section,omitempty" toml:"section" yaml:"section,omitempty"`
	Type      string `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Content   string `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	VersionNo int    `boil:"version_no" json:"version_no,omitempty" toml:"version_no" yaml:"version_no,omitempty"`
}

type GetAllInput struct {
	web.PagedResultRequest
	Title   string
	State   string
	Section string
}

type PagedRuleResponse struct {
	Rules      []Rule
	TotalCount int64
}
