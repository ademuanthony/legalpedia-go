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

// JudgementPrinciple is an object representing the database table.
type JudgementPrinciple struct {
	ID          int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	PrincipleId null.Int    `boil:"PrincipleId" json:"PrincipleId,omitempty" toml:"PrincipleId" yaml:"PrincipleId,omitempty"`
	SuitNo      null.String `boil:"SuitNo" json:"SuitNo,omitempty" toml:"SuitNo" yaml:"SuitNo,omitempty"`

	R *judgementPrincipleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L judgementPrincipleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JudgementPrincipleColumns = struct {
	ID          string
	PrincipleId string
	SuitNo      string
}{
	ID:          "Id",
	PrincipleId: "PrincipleId",
	SuitNo:      "SuitNo",
}

// Generated where

var JudgementPrincipleWhere = struct {
	ID          whereHelperint
	PrincipleId whereHelpernull_Int
	SuitNo      whereHelpernull_String
}{
	ID:          whereHelperint{field: "\"JudgementPrinciples\".\"Id\""},
	PrincipleId: whereHelpernull_Int{field: "\"JudgementPrinciples\".\"PrincipleId\""},
	SuitNo:      whereHelpernull_String{field: "\"JudgementPrinciples\".\"SuitNo\""},
}

// JudgementPrincipleRels is where relationship names are stored.
var JudgementPrincipleRels = struct {
}{}

// judgementPrincipleR is where relationships are stored.
type judgementPrincipleR struct {
}

// NewStruct creates a new relationship struct
func (*judgementPrincipleR) NewStruct() *judgementPrincipleR {
	return &judgementPrincipleR{}
}

// judgementPrincipleL is where Load methods for each relationship are stored.
type judgementPrincipleL struct{}

var (
	judgementPrincipleAllColumns            = []string{"Id", "PrincipleId", "SuitNo"}
	judgementPrincipleColumnsWithoutDefault = []string{"Id", "PrincipleId", "SuitNo"}
	judgementPrincipleColumnsWithDefault    = []string{}
	judgementPrinciplePrimaryKeyColumns     = []string{"Id"}
)

type (
	// JudgementPrincipleSlice is an alias for a slice of pointers to JudgementPrinciple.
	// This should generally be used opposed to []JudgementPrinciple.
	JudgementPrincipleSlice []*JudgementPrinciple

	judgementPrincipleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	judgementPrincipleType                 = reflect.TypeOf(&JudgementPrinciple{})
	judgementPrincipleMapping              = queries.MakeStructMapping(judgementPrincipleType)
	judgementPrinciplePrimaryKeyMapping, _ = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, judgementPrinciplePrimaryKeyColumns)
	judgementPrincipleInsertCacheMut       sync.RWMutex
	judgementPrincipleInsertCache          = make(map[string]insertCache)
	judgementPrincipleUpdateCacheMut       sync.RWMutex
	judgementPrincipleUpdateCache          = make(map[string]updateCache)
	judgementPrincipleUpsertCacheMut       sync.RWMutex
	judgementPrincipleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single judgementPrinciple record from the query.
func (q judgementPrincipleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*JudgementPrinciple, error) {
	o := &JudgementPrinciple{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for JudgementPrinciples")
	}

	return o, nil
}

// All returns all JudgementPrinciple records from the query.
func (q judgementPrincipleQuery) All(ctx context.Context, exec boil.ContextExecutor) (JudgementPrincipleSlice, error) {
	var o []*JudgementPrinciple

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JudgementPrinciple slice")
	}

	return o, nil
}

// Count returns the count of all JudgementPrinciple records in the query.
func (q judgementPrincipleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count JudgementPrinciples rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q judgementPrincipleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if JudgementPrinciples exists")
	}

	return count > 0, nil
}

// JudgementPrinciples retrieves all the records using an executor.
func JudgementPrinciples(mods ...qm.QueryMod) judgementPrincipleQuery {
	mods = append(mods, qm.From("\"JudgementPrinciples\""))
	return judgementPrincipleQuery{NewQuery(mods...)}
}

// FindJudgementPrinciple retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJudgementPrinciple(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*JudgementPrinciple, error) {
	judgementPrincipleObj := &JudgementPrinciple{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"JudgementPrinciples\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, judgementPrincipleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from JudgementPrinciples")
	}

	return judgementPrincipleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *JudgementPrinciple) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no JudgementPrinciples provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(judgementPrincipleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	judgementPrincipleInsertCacheMut.RLock()
	cache, cached := judgementPrincipleInsertCache[key]
	judgementPrincipleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			judgementPrincipleAllColumns,
			judgementPrincipleColumnsWithDefault,
			judgementPrincipleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"JudgementPrinciples\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"JudgementPrinciples\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into JudgementPrinciples")
	}

	if !cached {
		judgementPrincipleInsertCacheMut.Lock()
		judgementPrincipleInsertCache[key] = cache
		judgementPrincipleInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the JudgementPrinciple.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *JudgementPrinciple) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	judgementPrincipleUpdateCacheMut.RLock()
	cache, cached := judgementPrincipleUpdateCache[key]
	judgementPrincipleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			judgementPrincipleAllColumns,
			judgementPrinciplePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update JudgementPrinciples, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"JudgementPrinciples\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, judgementPrinciplePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, append(wl, judgementPrinciplePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update JudgementPrinciples row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for JudgementPrinciples")
	}

	if !cached {
		judgementPrincipleUpdateCacheMut.Lock()
		judgementPrincipleUpdateCache[key] = cache
		judgementPrincipleUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q judgementPrincipleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for JudgementPrinciples")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for JudgementPrinciples")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JudgementPrincipleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementPrinciplePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"JudgementPrinciples\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, judgementPrinciplePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in judgementPrinciple slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all judgementPrinciple")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *JudgementPrinciple) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no JudgementPrinciples provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(judgementPrincipleColumnsWithDefault, o)

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

	judgementPrincipleUpsertCacheMut.RLock()
	cache, cached := judgementPrincipleUpsertCache[key]
	judgementPrincipleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			judgementPrincipleAllColumns,
			judgementPrincipleColumnsWithDefault,
			judgementPrincipleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			judgementPrincipleAllColumns,
			judgementPrinciplePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert JudgementPrinciples, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(judgementPrinciplePrimaryKeyColumns))
			copy(conflict, judgementPrinciplePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"JudgementPrinciples\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(judgementPrincipleType, judgementPrincipleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert JudgementPrinciples")
	}

	if !cached {
		judgementPrincipleUpsertCacheMut.Lock()
		judgementPrincipleUpsertCache[key] = cache
		judgementPrincipleUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single JudgementPrinciple record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JudgementPrinciple) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no JudgementPrinciple provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), judgementPrinciplePrimaryKeyMapping)
	sql := "DELETE FROM \"JudgementPrinciples\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from JudgementPrinciples")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for JudgementPrinciples")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q judgementPrincipleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no judgementPrincipleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from JudgementPrinciples")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for JudgementPrinciples")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JudgementPrincipleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementPrinciplePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"JudgementPrinciples\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgementPrinciplePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from judgementPrinciple slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for JudgementPrinciples")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JudgementPrinciple) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJudgementPrinciple(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JudgementPrincipleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JudgementPrincipleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgementPrinciplePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"JudgementPrinciples\".* FROM \"JudgementPrinciples\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgementPrinciplePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JudgementPrincipleSlice")
	}

	*o = slice

	return nil
}

// JudgementPrincipleExists checks if the JudgementPrinciple row exists.
func JudgementPrincipleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"JudgementPrinciples\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if JudgementPrinciples exists")
	}

	return exists, nil
}
