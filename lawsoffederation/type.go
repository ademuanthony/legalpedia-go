package lawsoffederation

import (
	"context"
	"time"

	"github.com/ademuanthony/legalpedia/web"
)

type store interface {
	GetAllLawsOfFederation(ctx context.Context, req FindRequest) (*PagedResponse, error)
	GetLawDetails(ctx context.Context, id int) (*LawDetail, error)
}

type module struct {
	db     store
	server *web.Server
}

type LawsOfFederation struct {
	ID                    int       `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	CatId                 int       `boil:"CatId" json:"CatId,omitempty" toml:"CatId" yaml:"CatId,omitempty"`
	LawNo                 string    `boil:"LawNo" json:"LawNo,omitempty" toml:"LawNo" yaml:"LawNo,omitempty"`
	Title                 string    `boil:"Title" json:"Title,omitempty" toml:"Title" yaml:"Title,omitempty"`
	LawDate               time.Time `boil:"LawDate" json:"LawDate,omitempty" toml:"LawDate" yaml:"LawDate,omitempty"`
	Descr                 string    `boil:"Descr" json:"Descr,omitempty" toml:"Descr" yaml:"Descr,omitempty"`
	SubsidiaryLegislation string    `boil:"SubsidiaryLegislation" json:"SubsidiaryLegislation,omitempty" toml:"SubsidiaryLegislation" yaml:"SubsidiaryLegislation,omitempty"`
	Tags                  string    `boil:"Tags" json:"Tags,omitempty" toml:"Tags" yaml:"Tags,omitempty"`
	CreatedAt             time.Time `boil:"CreatedAt" json:"CreatedAt,omitempty" toml:"CreatedAt" yaml:"CreatedAt,omitempty"`
	UpdatedAt             time.Time `boil:"UpdatedAt" json:"UpdatedAt,omitempty" toml:"UpdatedAt" yaml:"UpdatedAt,omitempty"`
}

type Part struct {
	ID         int    `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	PartHeader string `boil:"PartHeader" json:"PartHeader,omitempty" toml:"PartHeader" yaml:"PartHeader,omitempty"`
	LawId      int    `boil:"LawId" json:"LawId,omitempty" toml:"LawId" yaml:"LawId,omitempty"`

	Sections []Section `json:"sctions"`
}

type Section struct {
	ID            int    `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	SectionHeader string `boil:"SectionHeader" json:"SectionHeader,omitempty" toml:"SectionHeader" yaml:"SectionHeader,omitempty"`
	SectionBody   string `boil:"SectionBody" json:"SectionBody,omitempty" toml:"SectionBody" yaml:"SectionBody,omitempty"`
	LawId         int    `boil:"LawId" json:"LawId,omitempty" toml:"LawId" yaml:"LawId,omitempty"`
	PartId        int    `boil:"PartId" json:"PartId,omitempty" toml:"PartId" yaml:"PartId,omitempty"`
}

type Schedule struct {
	ID          int    `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	SchedHeader string `boil:"SchedHeader" json:"SchedHeader,omitempty" toml:"SchedHeader" yaml:"SchedHeader,omitempty"`
	SchedBody   string `boil:"SchedBody" json:"SchedBody,omitempty" toml:"SchedBody" yaml:"SchedBody,omitempty"`
	LawId       int    `boil:"LawId" json:"LawId,omitempty" toml:"LawId" yaml:"LawId,omitempty"`
}

type LawDetail struct {
	LawsOfFederation

	Parts     []Part     `json:"parts"`
	Schedules []Schedule `json:"schedules"`
}

type FindRequest struct {
	web.PagedResultRequest
	Title string
	Year  int
}

type PagedResponse struct {
	Data       []LawsOfFederation
	TotalCount int64
}
