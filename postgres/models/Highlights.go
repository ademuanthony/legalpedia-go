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

// Highlight is an object representing the database table.
type Highlight struct {
	ID           string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	UserId       int64       `boil:"UserId" json:"UserId" toml:"UserId" yaml:"UserId"`
	CollectionId null.String `boil:"CollectionId" json:"CollectionId,omitempty" toml:"CollectionId" yaml:"CollectionId,omitempty"`
	Text         null.String `boil:"Text" json:"Text,omitempty" toml:"Text" yaml:"Text,omitempty"`
	CaseId       null.String `boil:"CaseId" json:"CaseId,omitempty" toml:"CaseId" yaml:"CaseId,omitempty"`
	PageNumber   int         `boil:"PageNumber" json:"PageNumber" toml:"PageNumber" yaml:"PageNumber"`
	StartIndex   int         `boil:"StartIndex" json:"StartIndex" toml:"StartIndex" yaml:"StartIndex"`
	Length       int         `boil:"Length" json:"Length" toml:"Length" yaml:"Length"`

	R *highlightR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L highlightL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HighlightColumns = struct {
	ID           string
	UserId       string
	CollectionId string
	Text         string
	CaseId       string
	PageNumber   string
	StartIndex   string
	Length       string
}{
	ID:           "Id",
	UserId:       "UserId",
	CollectionId: "CollectionId",
	Text:         "Text",
	CaseId:       "CaseId",
	PageNumber:   "PageNumber",
	StartIndex:   "StartIndex",
	Length:       "Length",
}

// Generated where

var HighlightWhere = struct {
	ID           whereHelperstring
	UserId       whereHelperint64
	CollectionId whereHelpernull_String
	Text         whereHelpernull_String
	CaseId       whereHelpernull_String
	PageNumber   whereHelperint
	StartIndex   whereHelperint
	Length       whereHelperint
}{
	ID:           whereHelperstring{field: "\"Highlights\".\"Id\""},
	UserId:       whereHelperint64{field: "\"Highlights\".\"UserId\""},
	CollectionId: whereHelpernull_String{field: "\"Highlights\".\"CollectionId\""},
	Text:         whereHelpernull_String{field: "\"Highlights\".\"Text\""},
	CaseId:       whereHelpernull_String{field: "\"Highlights\".\"CaseId\""},
	PageNumber:   whereHelperint{field: "\"Highlights\".\"PageNumber\""},
	StartIndex:   whereHelperint{field: "\"Highlights\".\"StartIndex\""},
	Length:       whereHelperint{field: "\"Highlights\".\"Length\""},
}

// HighlightRels is where relationship names are stored.
var HighlightRels = struct {
}{}

// highlightR is where relationships are stored.
type highlightR struct {
}

// NewStruct creates a new relationship struct
func (*highlightR) NewStruct() *highlightR {
	return &highlightR{}
}

// highlightL is where Load methods for each relationship are stored.
type highlightL struct{}

var (
	highlightAllColumns            = []string{"Id", "UserId", "CollectionId", "Text", "CaseId", "PageNumber", "StartIndex", "Length"}
	highlightColumnsWithoutDefault = []string{"Id", "UserId", "CollectionId", "Text", "CaseId", "PageNumber", "StartIndex", "Length"}
	highlightColumnsWithDefault    = []string{}
	highlightPrimaryKeyColumns     = []string{"Id"}
)

type (
	// HighlightSlice is an alias for a slice of pointers to Highlight.
	// This should generally be used opposed to []Highlight.
	HighlightSlice []*Highlight

	highlightQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	highlightType                 = reflect.TypeOf(&Highlight{})
	highlightMapping              = queries.MakeStructMapping(highlightType)
	highlightPrimaryKeyMapping, _ = queries.BindMapping(highlightType, highlightMapping, highlightPrimaryKeyColumns)
	highlightInsertCacheMut       sync.RWMutex
	highlightInsertCache          = make(map[string]insertCache)
	highlightUpdateCacheMut       sync.RWMutex
	highlightUpdateCache          = make(map[string]updateCache)
	highlightUpsertCacheMut       sync.RWMutex
	highlightUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single highlight record from the query.
func (q highlightQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Highlight, error) {
	o := &Highlight{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Highlights")
	}

	return o, nil
}

// All returns all Highlight records from the query.
func (q highlightQuery) All(ctx context.Context, exec boil.ContextExecutor) (HighlightSlice, error) {
	var o []*Highlight

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Highlight slice")
	}

	return o, nil
}

// Count returns the count of all Highlight records in the query.
func (q highlightQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Highlights rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q highlightQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Highlights exists")
	}

	return count > 0, nil
}

// Highlights retrieves all the records using an executor.
func Highlights(mods ...qm.QueryMod) highlightQuery {
	mods = append(mods, qm.From("\"Highlights\""))
	return highlightQuery{NewQuery(mods...)}
}

// FindHighlight retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHighlight(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Highlight, error) {
	highlightObj := &Highlight{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Highlights\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, highlightObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Highlights")
	}

	return highlightObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Highlight) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Highlights provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(highlightColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	highlightInsertCacheMut.RLock()
	cache, cached := highlightInsertCache[key]
	highlightInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			highlightAllColumns,
			highlightColumnsWithDefault,
			highlightColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(highlightType, highlightMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Highlights\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Highlights\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into Highlights")
	}

	if !cached {
		highlightInsertCacheMut.Lock()
		highlightInsertCache[key] = cache
		highlightInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Highlight.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Highlight) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	highlightUpdateCacheMut.RLock()
	cache, cached := highlightUpdateCache[key]
	highlightUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			highlightAllColumns,
			highlightPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Highlights, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Highlights\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, highlightPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, append(wl, highlightPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Highlights row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Highlights")
	}

	if !cached {
		highlightUpdateCacheMut.Lock()
		highlightUpdateCache[key] = cache
		highlightUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q highlightQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Highlights")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Highlights")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HighlightSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Highlights\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, highlightPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in highlight slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all highlight")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Highlight) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Highlights provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(highlightColumnsWithDefault, o)

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

	highlightUpsertCacheMut.RLock()
	cache, cached := highlightUpsertCache[key]
	highlightUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			highlightAllColumns,
			highlightColumnsWithDefault,
			highlightColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			highlightAllColumns,
			highlightPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert Highlights, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(highlightPrimaryKeyColumns))
			copy(conflict, highlightPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"Highlights\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(highlightType, highlightMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(highlightType, highlightMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert Highlights")
	}

	if !cached {
		highlightUpsertCacheMut.Lock()
		highlightUpsertCache[key] = cache
		highlightUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Highlight record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Highlight) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Highlight provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), highlightPrimaryKeyMapping)
	sql := "DELETE FROM \"Highlights\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Highlights")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Highlights")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q highlightQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no highlightQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Highlights")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Highlights")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HighlightSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Highlights\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, highlightPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from highlight slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Highlights")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Highlight) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHighlight(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HighlightSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HighlightSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), highlightPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Highlights\".* FROM \"Highlights\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, highlightPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HighlightSlice")
	}

	*o = slice

	return nil
}

// HighlightExists checks if the Highlight row exists.
func HighlightExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Highlights\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Highlights exists")
	}

	return exists, nil
}
