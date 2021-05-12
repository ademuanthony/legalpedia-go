package postgres

import (
	"context"
	"fmt"

	. "github.com/ademuanthony/legalpedia/dictionary"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) FindDictionaries(ctx context.Context, req FindRequest) (*PagedResponse, error) {
	var queries []qm.QueryMod
	if req.Title != "" {
		statement := fmt.Sprintf("(%s LIKE '%%%s%%')",
			models.DictionaryColumns.Title, req.Title,
		)
		queries = append(queries, qm.Where(statement))
	}

	totalCount, err := models.Dictionaries(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Limit(req.Limit), qm.Offset(req.Offset))

	data, err := models.Dictionaries(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var dictionaries = make([]Dictionay, len(data))
	for i, d := range data {
		dictionaries[i] = Dictionay{
			ID:        d.ID.Int,
			UUID:      d.UUID,
			Title:     d.Title.String,
			Content:   d.Content.String,
			VersionNo: d.VersionNo.Int,
		}
	}

	return &PagedResponse{
		Data:       dictionaries,
		TotalCount: totalCount,
	}, nil
}

func (pg *PgDb) FindDictionaryByID(ctx context.Context, id string) (*Dictionay, error) {
	d, err := models.Dictionaries(models.DictionaryWhere.UUID.EQ(id)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}
	return &Dictionay{
		ID:        d.ID.Int,
		UUID:      d.UUID,
		Title:     d.Title.String,
		Content:   d.Content.String,
		VersionNo: d.VersionNo.Int,
	}, nil
}
