package postgres

import (
	"context"
	"fmt"
	"strings"

	lfn "github.com/ademuanthony/legalpedia/lawsoffederation"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) GetAllLawsOfFederation(ctx context.Context, req lfn.FindRequest) (*lfn.PagedResponse, error) {
	var queries []qm.QueryMod
	if req.Title != "" {
		statement := fmt.Sprintf("(%s LIKE '%%%s%%')",
			"\"LawsOfFederation\".\"Title\"", strings.ToUpper(req.Title),
		)
		queries = append(queries, qm.Where(statement))
	}

	if req.Year != 0 {
		// get year from date and compare
		statement := fmt.Sprintf("(TO_CHAR(%s, 'YYYY') = '%d')",
		"\"LawsOfFederation\".\"LawDate\"", req.Year,
		)
		queries = append(queries, qm.Where(statement))
	}

	totalCount, err := models.LawsOfFederations(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(
		queries, qm.Limit(req.Limit), qm.Offset(req.Offset), 
		qm.OrderBy("\"LawsOfFederation\".\"Title\""),
	)

	data, err := models.LawsOfFederations(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var dictionaries = make([]lfn.LawsOfFederation, len(data))
	for i, d := range data {
		dictionaries[i] = lfn.LawsOfFederation{
			ID:                    d.ID,
			Title:                 d.Title.String,
			LawNo:                 d.LawNo.String,
			LawDate:               d.LawDate.Time,
			CatId:                 d.CatId.Int,
			Descr:                 d.Descr.String,
			SubsidiaryLegislation: d.SubsidiaryLegislation.String,
			Tags:                  d.Tags.String,
			CreatedAt:             d.CreatedAt.Time,
			UpdatedAt:             d.UpdatedAt.Time,
		}
	}

	return &lfn.PagedResponse{
		Data:       dictionaries,
		TotalCount: totalCount,
	}, nil
}

func (pg *PgDb) GetLawDetails(ctx context.Context, id int) (*lfn.LawDetail, error) {
	var lawDetail = &lfn.LawDetail{}

	data, err := models.LawsOfFederations(models.LawsOfFederationWhere.ID.EQ(id)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	lawDetail.CatId = data.CatId.Int
	lawDetail.CreatedAt = data.CreatedAt.Time
	lawDetail.Descr = data.Descr.String
	lawDetail.ID = data.ID
	lawDetail.LawDate = data.LawDate.Time
	lawDetail.LawNo = data.LawNo.String
	lawDetail.Title = data.Title.String
	lawDetail.SubsidiaryLegislation = data.SubsidiaryLegislation.String
	lawDetail.Tags = data.Tags.String
	lawDetail.UpdatedAt = data.UpdatedAt.Time

	parts, err := models.LawOfFedParts(models.LawOfFedPartWhere.LawId.EQ(null.IntFrom(id))).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	for _, p := range parts {
		var part = lfn.Part{
			ID: p.ID,
			PartHeader: p.PartHeader.String,
			LawId: p.LawId.Int,
		}

		sections, err := models.LawOfFedSections(models.LawOfFedSectionWhere.LawId.EQ(null.IntFrom(id))).All(ctx, pg.db)
		if err != nil {
			return nil, err
		}

		for _, s := range sections {
			part.Sections = append(part.Sections, lfn.Section{
				ID: s.ID,
				LawId: s.LawId.Int,
				SectionHeader: s.SectionBody.String,
				SectionBody: s.SectionBody.String,
			})
		}

		lawDetail.Parts = append(lawDetail.Parts, part)
	}

	schedules, err := models.LawOfFedScheds(models.LawOfFedSchedWhere.LawId.EQ(null.IntFrom(id))).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	for _, s := range schedules {
		lawDetail.Schedules = append(lawDetail.Schedules, lfn.Schedule{
			ID: s.ID,
			LawId: s.LawId.Int,
			SchedHeader: s.SchedHeader.String,
			SchedBody: s.SchedBody.String,
		})
	}

	return lawDetail, nil
}
