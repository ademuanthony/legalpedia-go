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

// LawOfFedSection is an object representing the database table.
type LawOfFedSection struct {
	ID            int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	SectionHeader null.String `boil:"SectionHeader" json:"SectionHeader,omitempty" toml:"SectionHeader" yaml:"SectionHeader,omitempty"`
	SectionBody   null.String `boil:"SectionBody" json:"SectionBody,omitempty" toml:"SectionBody" yaml:"SectionBody,omitempty"`
	LawId         null.Int    `boil:"LawId" json:"LawId,omitempty" toml:"LawId" yaml:"LawId,omitempty"`
	PartId        null.Int    `boil:"PartId" json:"PartId,omitempty" toml:"PartId" yaml:"PartId,omitempty"`

	R *lawOfFedSectionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lawOfFedSectionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LawOfFedSectionColumns = struct {
	ID            string
	SectionHeader string
	SectionBody   string
	LawId         string
	PartId        string
}{
	ID:            "Id",
	SectionHeader: "SectionHeader",
	SectionBody:   "SectionBody",
	LawId:         "LawId",
	PartId:        "PartId",
}

// Generated where

var LawOfFedSectionWhere = struct {
	ID            whereHelperint
	SectionHeader whereHelpernull_String
	SectionBody   whereHelpernull_String
	LawId         whereHelpernull_Int
	PartId        whereHelpernull_Int
}{
	ID:            whereHelperint{field: "\"LawOfFedSections\".\"Id\""},
	SectionHeader: whereHelpernull_String{field: "\"LawOfFedSections\".\"SectionHeader\""},
	SectionBody:   whereHelpernull_String{field: "\"LawOfFedSections\".\"SectionBody\""},
	LawId:         whereHelpernull_Int{field: "\"LawOfFedSections\".\"LawId\""},
	PartId:        whereHelpernull_Int{field: "\"LawOfFedSections\".\"PartId\""},
}

// LawOfFedSectionRels is where relationship names are stored.
var LawOfFedSectionRels = struct {
}{}

// lawOfFedSectionR is where relationships are stored.
type lawOfFedSectionR struct {
}

// NewStruct creates a new relationship struct
func (*lawOfFedSectionR) NewStruct() *lawOfFedSectionR {
	return &lawOfFedSectionR{}
}

// lawOfFedSectionL is where Load methods for each relationship are stored.
type lawOfFedSectionL struct{}

var (
	lawOfFedSectionAllColumns            = []string{"Id", "SectionHeader", "SectionBody", "LawId", "PartId"}
	lawOfFedSectionColumnsWithoutDefault = []string{"Id", "SectionHeader", "SectionBody", "LawId", "PartId"}
	lawOfFedSectionColumnsWithDefault    = []string{}
	lawOfFedSectionPrimaryKeyColumns     = []string{"Id"}
)

type (
	// LawOfFedSectionSlice is an alias for a slice of pointers to LawOfFedSection.
	// This should generally be used opposed to []LawOfFedSection.
	LawOfFedSectionSlice []*LawOfFedSection

	lawOfFedSectionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lawOfFedSectionType                 = reflect.TypeOf(&LawOfFedSection{})
	lawOfFedSectionMapping              = queries.MakeStructMapping(lawOfFedSectionType)
	lawOfFedSectionPrimaryKeyMapping, _ = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, lawOfFedSectionPrimaryKeyColumns)
	lawOfFedSectionInsertCacheMut       sync.RWMutex
	lawOfFedSectionInsertCache          = make(map[string]insertCache)
	lawOfFedSectionUpdateCacheMut       sync.RWMutex
	lawOfFedSectionUpdateCache          = make(map[string]updateCache)
	lawOfFedSectionUpsertCacheMut       sync.RWMutex
	lawOfFedSectionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single lawOfFedSection record from the query.
func (q lawOfFedSectionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LawOfFedSection, error) {
	o := &LawOfFedSection{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for LawOfFedSections")
	}

	return o, nil
}

// All returns all LawOfFedSection records from the query.
func (q lawOfFedSectionQuery) All(ctx context.Context, exec boil.ContextExecutor) (LawOfFedSectionSlice, error) {
	var o []*LawOfFedSection

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LawOfFedSection slice")
	}

	return o, nil
}

// Count returns the count of all LawOfFedSection records in the query.
func (q lawOfFedSectionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count LawOfFedSections rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lawOfFedSectionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if LawOfFedSections exists")
	}

	return count > 0, nil
}

// LawOfFedSections retrieves all the records using an executor.
func LawOfFedSections(mods ...qm.QueryMod) lawOfFedSectionQuery {
	mods = append(mods, qm.From("\"LawOfFedSections\""))
	return lawOfFedSectionQuery{NewQuery(mods...)}
}

// FindLawOfFedSection retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLawOfFedSection(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LawOfFedSection, error) {
	lawOfFedSectionObj := &LawOfFedSection{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"LawOfFedSections\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lawOfFedSectionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from LawOfFedSections")
	}

	return lawOfFedSectionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LawOfFedSection) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no LawOfFedSections provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lawOfFedSectionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lawOfFedSectionInsertCacheMut.RLock()
	cache, cached := lawOfFedSectionInsertCache[key]
	lawOfFedSectionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lawOfFedSectionAllColumns,
			lawOfFedSectionColumnsWithDefault,
			lawOfFedSectionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"LawOfFedSections\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"LawOfFedSections\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into LawOfFedSections")
	}

	if !cached {
		lawOfFedSectionInsertCacheMut.Lock()
		lawOfFedSectionInsertCache[key] = cache
		lawOfFedSectionInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the LawOfFedSection.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LawOfFedSection) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lawOfFedSectionUpdateCacheMut.RLock()
	cache, cached := lawOfFedSectionUpdateCache[key]
	lawOfFedSectionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lawOfFedSectionAllColumns,
			lawOfFedSectionPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update LawOfFedSections, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"LawOfFedSections\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lawOfFedSectionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, append(wl, lawOfFedSectionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update LawOfFedSections row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for LawOfFedSections")
	}

	if !cached {
		lawOfFedSectionUpdateCacheMut.Lock()
		lawOfFedSectionUpdateCache[key] = cache
		lawOfFedSectionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lawOfFedSectionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for LawOfFedSections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for LawOfFedSections")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LawOfFedSectionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lawOfFedSectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"LawOfFedSections\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lawOfFedSectionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in lawOfFedSection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all lawOfFedSection")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LawOfFedSection) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no LawOfFedSections provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lawOfFedSectionColumnsWithDefault, o)

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

	lawOfFedSectionUpsertCacheMut.RLock()
	cache, cached := lawOfFedSectionUpsertCache[key]
	lawOfFedSectionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lawOfFedSectionAllColumns,
			lawOfFedSectionColumnsWithDefault,
			lawOfFedSectionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lawOfFedSectionAllColumns,
			lawOfFedSectionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert LawOfFedSections, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lawOfFedSectionPrimaryKeyColumns))
			copy(conflict, lawOfFedSectionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"LawOfFedSections\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lawOfFedSectionType, lawOfFedSectionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert LawOfFedSections")
	}

	if !cached {
		lawOfFedSectionUpsertCacheMut.Lock()
		lawOfFedSectionUpsertCache[key] = cache
		lawOfFedSectionUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single LawOfFedSection record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LawOfFedSection) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LawOfFedSection provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lawOfFedSectionPrimaryKeyMapping)
	sql := "DELETE FROM \"LawOfFedSections\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from LawOfFedSections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for LawOfFedSections")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lawOfFedSectionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no lawOfFedSectionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from LawOfFedSections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for LawOfFedSections")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LawOfFedSectionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lawOfFedSectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"LawOfFedSections\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lawOfFedSectionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lawOfFedSection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for LawOfFedSections")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LawOfFedSection) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLawOfFedSection(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LawOfFedSectionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LawOfFedSectionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lawOfFedSectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"LawOfFedSections\".* FROM \"LawOfFedSections\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lawOfFedSectionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LawOfFedSectionSlice")
	}

	*o = slice

	return nil
}

// LawOfFedSectionExists checks if the LawOfFedSection row exists.
func LawOfFedSectionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"LawOfFedSections\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if LawOfFedSections exists")
	}

	return exists, nil
}
