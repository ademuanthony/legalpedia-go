// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// JudgementsSummary is an object representing the database table.
type JudgementsSummary struct {
	ID             string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	Title          null.String `boil:"Title" json:"Title,omitempty" toml:"Title" yaml:"Title,omitempty"`
	SummaryOfFacts null.String `boil:"SummaryOfFacts" json:"SummaryOfFacts,omitempty" toml:"SummaryOfFacts" yaml:"SummaryOfFacts,omitempty"`
	Held           null.String `boil:"Held" json:"Held,omitempty" toml:"Held" yaml:"Held,omitempty"`
	Issues         null.String `boil:"Issues" json:"Issues,omitempty" toml:"Issues" yaml:"Issues,omitempty"`
	CasesCited     null.String `boil:"CasesCited" json:"CasesCited,omitempty" toml:"CasesCited" yaml:"CasesCited,omitempty"`
	StatutesCited  null.String `boil:"StatutesCited" json:"StatutesCited,omitempty" toml:"StatutesCited" yaml:"StatutesCited,omitempty"`
	LPCitation     null.String `boil:"LPCitation" json:"LPCitation,omitempty" toml:"LPCitation" yaml:"LPCitation,omitempty"`
	JudgementDate  null.Time   `boil:"JudgementDate" json:"JudgementDate,omitempty" toml:"JudgementDate" yaml:"JudgementDate,omitempty"`
	OtherCitations null.String `boil:"OtherCitations" json:"OtherCitations,omitempty" toml:"OtherCitations" yaml:"OtherCitations,omitempty"`
	HoldenAtID     null.Int    `boil:"HoldenAtID" json:"HoldenAtID,omitempty" toml:"HoldenAtID" yaml:"HoldenAtID,omitempty"`
	CourtID        null.Int    `boil:"CourtID" json:"CourtID,omitempty" toml:"CourtID" yaml:"CourtID,omitempty"`
	PartyATypeID   null.Int    `boil:"PartyATypeID" json:"PartyATypeID,omitempty" toml:"PartyATypeID" yaml:"PartyATypeID,omitempty"`
	PartyBTypeID   null.Int    `boil:"PartyBTypeID" json:"PartyBTypeID,omitempty" toml:"PartyBTypeID" yaml:"PartyBTypeID,omitempty"`
	CategoryId     null.Int    `boil:"CategoryId" json:"CategoryId,omitempty" toml:"CategoryId" yaml:"CategoryId,omitempty"`
	Tags           null.String `boil:"Tags" json:"Tags,omitempty" toml:"Tags" yaml:"Tags,omitempty"`
	CreatedAt      null.Time   `boil:"CreatedAt" json:"CreatedAt,omitempty" toml:"CreatedAt" yaml:"CreatedAt,omitempty"`
	UpdatedAt      null.Time   `boil:"UpdatedAt" json:"UpdatedAt,omitempty" toml:"UpdatedAt" yaml:"UpdatedAt,omitempty"`

	R *judgementsSummaryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L judgementsSummaryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JudgementsSummaryColumns = struct {
	ID             string
	Title          string
	SummaryOfFacts string
	Held           string
	Issues         string
	CasesCited     string
	StatutesCited  string
	LPCitation     string
	JudgementDate  string
	OtherCitations string
	HoldenAtID     string
	CourtID        string
	PartyATypeID   string
	PartyBTypeID   string
	CategoryId     string
	Tags           string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "Id",
	Title:          "Title",
	SummaryOfFacts: "SummaryOfFacts",
	Held:           "Held",
	Issues:         "Issues",
	CasesCited:     "CasesCited",
	StatutesCited:  "StatutesCited",
	LPCitation:     "LPCitation",
	JudgementDate:  "JudgementDate",
	OtherCitations: "OtherCitations",
	HoldenAtID:     "HoldenAtID",
	CourtID:        "CourtID",
	PartyATypeID:   "PartyATypeID",
	PartyBTypeID:   "PartyBTypeID",
	CategoryId:     "CategoryId",
	Tags:           "Tags",
	CreatedAt:      "CreatedAt",
	UpdatedAt:      "UpdatedAt",
}

// Generated where

var JudgementsSummaryWhere = struct {
	ID             whereHelperstring
	Title          whereHelpernull_String
	SummaryOfFacts whereHelpernull_String
	Held           whereHelpernull_String
	Issues         whereHelpernull_String
	CasesCited     whereHelpernull_String
	StatutesCited  whereHelpernull_String
	LPCitation     whereHelpernull_String
	JudgementDate  whereHelpernull_Time
	OtherCitations whereHelpernull_String
	HoldenAtID     whereHelpernull_Int
	CourtID        whereHelpernull_Int
	PartyATypeID   whereHelpernull_Int
	PartyBTypeID   whereHelpernull_Int
	CategoryId     whereHelpernull_Int
	Tags           whereHelpernull_String
	CreatedAt      whereHelpernull_Time
	UpdatedAt      whereHelpernull_Time
}{
	ID:             whereHelperstring{field: "\"JudgementsSummaries\".\"Id\""},
	Title:          whereHelpernull_String{field: "\"JudgementsSummaries\".\"Title\""},
	SummaryOfFacts: whereHelpernull_String{field: "\"JudgementsSummaries\".\"SummaryOfFacts\""},
	Held:           whereHelpernull_String{field: "\"JudgementsSummaries\".\"Held\""},
	Issues:         whereHelpernull_String{field: "\"JudgementsSummaries\".\"Issues\""},
	CasesCited:     whereHelpernull_String{field: "\"JudgementsSummaries\".\"CasesCited\""},
	StatutesCited:  whereHelpernull_String{field: "\"JudgementsSummaries\".\"StatutesCited\""},
	LPCitation:     whereHelpernull_String{field: "\"JudgementsSummaries\".\"LPCitation\""},
	JudgementDate:  whereHelpernull_Time{field: "\"JudgementsSummaries\".\"JudgementDate\""},
	OtherCitations: whereHelpernull_String{field: "\"JudgementsSummaries\".\"OtherCitations\""},
	HoldenAtID:     whereHelpernull_Int{field: "\"JudgementsSummaries\".\"HoldenAtID\""},
	CourtID:        whereHelpernull_Int{field: "\"JudgementsSummaries\".\"CourtID\""},
	PartyATypeID:   whereHelpernull_Int{field: "\"JudgementsSummaries\".\"PartyATypeID\""},
	PartyBTypeID:   whereHelpernull_Int{field: "\"JudgementsSummaries\".\"PartyBTypeID\""},
	CategoryId:     whereHelpernull_Int{field: "\"JudgementsSummaries\".\"CategoryId\""},
	Tags:           whereHelpernull_String{field: "\"JudgementsSummaries\".\"Tags\""},
	CreatedAt:      whereHelpernull_Time{field: "\"JudgementsSummaries\".\"CreatedAt\""},
	UpdatedAt:      whereHelpernull_Time{field: "\"JudgementsSummaries\".\"UpdatedAt\""},
}

// JudgementsSummaryRels is where relationship names are stored.
var JudgementsSummaryRels = struct {
}{}

// judgementsSummaryR is where relationships are stored.
type judgementsSummaryR struct {
}

// NewStruct creates a new relationship struct
func (*judgementsSummaryR) NewStruct() *judgementsSummaryR {
	return &judgementsSummaryR{}
}

// judgementsSummaryL is where Load methods for each relationship are stored.
type judgementsSummaryL struct{}

var (
	judgementsSummaryAllColumns            = []string{"Id", "Title", "SummaryOfFacts", "Held", "Issues", "CasesCited", "StatutesCited", "LPCitation", "JudgementDate", "OtherCitations", "HoldenAtID", "CourtID", "PartyATypeID", "PartyBTypeID", "CategoryId", "Tags", "CreatedAt", "UpdatedAt"}
	judgementsSummaryColumnsWithoutDefault = []string{"Id", "Title", "SummaryOfFacts", "Held", "Issues", "CasesCited", "StatutesCited", "LPCitation", "JudgementDate", "OtherCitations", "HoldenAtID", "CourtID", "PartyATypeID", "PartyBTypeID", "CategoryId", "Tags", "CreatedAt", "UpdatedAt"}
	judgementsSummaryColumnsWithDefault    = []string{}
	judgementsSummaryPrimaryKeyColumns     = []string{"Id"}
)

type (
	// JudgementsSummarySlice is an alias for a slice of pointers to JudgementsSummary.
	// This should generally be used opposed to []JudgementsSummary.
	JudgementsSummarySlice []*JudgementsSummary

	judgementsSummaryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	judgementsSummaryType                 = reflect.TypeOf(&JudgementsSummary{})
	judgementsSummaryMapping              = queries.MakeStructMapping(judgementsSummaryType)
	judgementsSummaryPrimaryKeyMapping, _ = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, judgementsSummaryPrimaryKeyColumns)
	judgementsSummaryInsertCacheMut       sync.RWMutex
	judgementsSummaryInsertCache          = make(map[string]insertCache)
	judgementsSummaryUpdateCacheMut       sync.RWMutex
	judgementsSummaryUpdateCache          = make(map[string]updateCache)
	judgementsSummaryUpsertCacheMut       sync.RWMutex
	judgementsSummaryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single judgementsSummary record from the query.
func (q judgementsSummaryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*JudgementsSummary, error) {
	o := &JudgementsSummary{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for JudgementsSummaries")
	}

	return o, nil
}

// All returns all JudgementsSummary records from the query.
func (q judgementsSummaryQuery) All(ctx context.Context, exec boil.ContextExecutor) (JudgementsSummarySlice, error) {
	var o []*JudgementsSummary

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JudgementsSummary slice")
	}

	return o, nil
}

// Count returns the count of all JudgementsSummary records in the query.
func (q judgementsSummaryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count JudgementsSummaries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q judgementsSummaryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if JudgementsSummaries exists")
	}

	return count > 0, nil
}

// JudgementsSummaries retrieves all the records using an executor.
func JudgementsSummaries(mods ...qm.QueryMod) judgementsSummaryQuery {
	mods = append(mods, qm.From("\"JudgementsSummaries\""))
	return judgementsSummaryQuery{NewQuery(mods...)}
}

// FindJudgementsSummary retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJudgementsSummary(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*JudgementsSummary, error) {
	judgementsSummaryObj := &JudgementsSummary{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"JudgementsSummaries\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, judgementsSummaryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from JudgementsSummaries")
	}

	return judgementsSummaryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *JudgementsSummary) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no JudgementsSummaries provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(judgementsSummaryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	judgementsSummaryInsertCacheMut.RLock()
	cache, cached := judgementsSummaryInsertCache[key]
	judgementsSummaryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			judgementsSummaryAllColumns,
			judgementsSummaryColumnsWithDefault,
			judgementsSummaryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"JudgementsSummaries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"JudgementsSummaries\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into JudgementsSummaries")
	}

	if !cached {
		judgementsSummaryInsertCacheMut.Lock()
		judgementsSummaryInsertCache[key] = cache
		judgementsSummaryInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the JudgementsSummary.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *JudgementsSummary) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	judgementsSummaryUpdateCacheMut.RLock()
	cache, cached := judgementsSummaryUpdateCache[key]
	judgementsSummaryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			judgementsSummaryAllColumns,
			judgementsSummaryPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update JudgementsSummaries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"JudgementsSummaries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, judgementsSummaryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, append(wl, judgementsSummaryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update JudgementsSummaries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for JudgementsSummaries")
	}

	if !cached {
		judgementsSummaryUpdateCacheMut.Lock()
		judgementsSummaryUpdateCache[key] = cache
		judgementsSummaryUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q judgementsSummaryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for JudgementsSummaries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for JudgementsSummaries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JudgementsSummarySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementsSummaryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"JudgementsSummaries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, judgementsSummaryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in judgementsSummary slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all judgementsSummary")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *JudgementsSummary) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no JudgementsSummaries provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(judgementsSummaryColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	judgementsSummaryUpsertCacheMut.RLock()
	cache, cached := judgementsSummaryUpsertCache[key]
	judgementsSummaryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			judgementsSummaryAllColumns,
			judgementsSummaryColumnsWithDefault,
			judgementsSummaryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			judgementsSummaryAllColumns,
			judgementsSummaryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert JudgementsSummaries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(judgementsSummaryPrimaryKeyColumns))
			copy(conflict, judgementsSummaryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"JudgementsSummaries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(judgementsSummaryType, judgementsSummaryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert JudgementsSummaries")
	}

	if !cached {
		judgementsSummaryUpsertCacheMut.Lock()
		judgementsSummaryUpsertCache[key] = cache
		judgementsSummaryUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single JudgementsSummary record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JudgementsSummary) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no JudgementsSummary provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), judgementsSummaryPrimaryKeyMapping)
	sql := "DELETE FROM \"JudgementsSummaries\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from JudgementsSummaries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for JudgementsSummaries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q judgementsSummaryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no judgementsSummaryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from JudgementsSummaries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for JudgementsSummaries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JudgementsSummarySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementsSummaryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"JudgementsSummaries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgementsSummaryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from judgementsSummary slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for JudgementsSummaries")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JudgementsSummary) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJudgementsSummary(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JudgementsSummarySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JudgementsSummarySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementsSummaryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"JudgementsSummaries\".* FROM \"JudgementsSummaries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgementsSummaryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JudgementsSummarySlice")
	}

	*o = slice

	return nil
}

// JudgementsSummaryExists checks if the JudgementsSummary row exists.
func JudgementsSummaryExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"JudgementsSummaries\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if JudgementsSummaries exists")
	}

	return exists, nil
}
