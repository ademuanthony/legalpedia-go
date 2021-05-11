package postgres

import (
	"context"

	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/ademuanthony/legalpedia/rulesofcourt"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) GetAll(ctx context.Context, input rulesofcourt.GetAllInput) (*rulesofcourt.PagedRuleResponse, error) {
	var queries []qm.QueryMod
	if input.Title != "" {
		queries = append(queries, qm.Where(models.RuleColumns.Title+" LIKE '%?%'", input.Title))
	}

	if input.Section != "" {
		queries = append(queries, models.RuleWhere.Section.EQ(null.StringFrom(input.Section)))
	}

	if input.State != "" {
		queries = append(queries, models.RuleWhere.Name.EQ(null.StringFrom(input.State)))
	}

	totalCount, err := models.Rules(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Offset(input.Offset), qm.Limit(input.Limit))

	records, err := models.Rules(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var rules = make([]rulesofcourt.Rule, len(records))
	for i, r := range records {
		rules[i] = ruleFromModel(r)
	}

	return &rulesofcourt.PagedRuleResponse{TotalCount: totalCount, Rules: rules}, nil

}

func (pg *PgDb) State(ctx context.Context, input rulesofcourt.GetAllInput) (*rulesofcourt.PagedRuleResponse, error) {
	var queries = []qm.QueryMod{models.RuleWhere.Type.EQ(null.StringFrom("State"))}
	if input.Title != "" {
		queries = append(queries, qm.Where(models.RuleColumns.Title+" LIKE '%?%'", input.Title))
	}

	if input.Section != "" {
		queries = append(queries, models.RuleWhere.Section.EQ(null.StringFrom(input.Section)))
	}

	if input.State != "" {
		queries = append(queries, models.RuleWhere.Name.EQ(null.StringFrom(input.State)))
	}

	totalCount, err := models.Rules(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Offset(input.Offset), qm.Limit(input.Limit))

	records, err := models.Rules(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var rules = make([]rulesofcourt.Rule, len(records))
	for i, r := range records {
		rules[i] = ruleFromModel(r)
	}

	return &rulesofcourt.PagedRuleResponse{TotalCount: totalCount, Rules: rules}, nil

}

func (pg *PgDb) Other(ctx context.Context, input rulesofcourt.GetAllInput) (*rulesofcourt.PagedRuleResponse, error) {
	var queries = []qm.QueryMod{models.RuleWhere.Type.EQ(null.StringFrom("Other"))}
	if input.Title != "" {
		queries = append(queries, qm.Where(models.RuleColumns.Title+" LIKE '%?%'", input.Title))
	}

	if input.Section != "" {
		queries = append(queries, models.RuleWhere.Section.EQ(null.StringFrom(input.Section)))
	}

	if input.State != "" {
		queries = append(queries, models.RuleWhere.Name.EQ(null.StringFrom(input.State)))
	}

	totalCount, err := models.Rules(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}
	
	queries = append(queries, qm.Offset(input.Offset), qm.Limit(input.Limit))

	records, err := models.Rules(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var rules = make([]rulesofcourt.Rule, len(records))
	for i, r := range records {
		rules[i] = ruleFromModel(r)
	}

	return &rulesofcourt.PagedRuleResponse{TotalCount: totalCount, Rules: rules}, nil

}

func ruleFromModel(r *models.Rule) rulesofcourt.Rule {
	return rulesofcourt.Rule{
		ID:        r.ID,
		UUID:      r.UUID.String,
		VersionNo: r.VersionNo.Int,
		Name:      r.Name.String,
		Section:   r.Section.String,
		Title:     r.Title.String,
		Type:      r.Type.String,
		Content:   r.Content.String,
	}
}

func (pg *PgDb)	GetNames(ctx context.Context, ruleType string) ([]string, error) {
	records, err := models.Rules(qm.Select("DISTINCT name"),
	 models.RuleWhere.Type.EQ(null.StringFrom(ruleType))).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var names = make([]string, len(records))
	for i, r := range records {
		names[i] = r.Name.String
	}
	return names, nil
}

func (pg *PgDb) Detail(ctx context.Context, id int) (*rulesofcourt.Rule, error) {
	var record, err = models.FindRule(ctx, pg.db, id)
	if err != nil {
		return nil, err
	}
	var rule = ruleFromModel(record)
	return &rule, nil
}
