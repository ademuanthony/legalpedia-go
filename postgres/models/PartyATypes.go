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

// PartyAType is an object representing the database table.
type PartyAType struct {
	ID         int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	PartyAType null.String `boil:"PartyAType" json:"PartyAType,omitempty" toml:"PartyAType" yaml:"PartyAType,omitempty"`

	R *partyATypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L partyATypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PartyATypeColumns = struct {
	ID         string
	PartyAType string
}{
	ID:         "Id",
	PartyAType: "PartyAType",
}

// Generated where

var PartyATypeWhere = struct {
	ID         whereHelperint
	PartyAType whereHelpernull_String
}{
	ID:         whereHelperint{field: "\"PartyATypes\".\"Id\""},
	PartyAType: whereHelpernull_String{field: "\"PartyATypes\".\"PartyAType\""},
}

// PartyATypeRels is where relationship names are stored.
var PartyATypeRels = struct {
}{}

// partyATypeR is where relationships are stored.
type partyATypeR struct {
}

// NewStruct creates a new relationship struct
func (*partyATypeR) NewStruct() *partyATypeR {
	return &partyATypeR{}
}

// partyATypeL is where Load methods for each relationship are stored.
type partyATypeL struct{}

var (
	partyATypeAllColumns            = []string{"Id", "PartyAType"}
	partyATypeColumnsWithoutDefault = []string{"Id", "PartyAType"}
	partyATypeColumnsWithDefault    = []string{}
	partyATypePrimaryKeyColumns     = []string{"Id"}
)

type (
	// PartyATypeSlice is an alias for a slice of pointers to PartyAType.
	// This should generally be used opposed to []PartyAType.
	PartyATypeSlice []*PartyAType

	partyATypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	partyATypeType                 = reflect.TypeOf(&PartyAType{})
	partyATypeMapping              = queries.MakeStructMapping(partyATypeType)
	partyATypePrimaryKeyMapping, _ = queries.BindMapping(partyATypeType, partyATypeMapping, partyATypePrimaryKeyColumns)
	partyATypeInsertCacheMut       sync.RWMutex
	partyATypeInsertCache          = make(map[string]insertCache)
	partyATypeUpdateCacheMut       sync.RWMutex
	partyATypeUpdateCache          = make(map[string]updateCache)
	partyATypeUpsertCacheMut       sync.RWMutex
	partyATypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single partyAType record from the query.
func (q partyATypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PartyAType, error) {
	o := &PartyAType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for PartyATypes")
	}

	return o, nil
}

// All returns all PartyAType records from the query.
func (q partyATypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (PartyATypeSlice, error) {
	var o []*PartyAType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PartyAType slice")
	}

	return o, nil
}

// Count returns the count of all PartyAType records in the query.
func (q partyATypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count PartyATypes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q partyATypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if PartyATypes exists")
	}

	return count > 0, nil
}

// PartyATypes retrieves all the records using an executor.
func PartyATypes(mods ...qm.QueryMod) partyATypeQuery {
	mods = append(mods, qm.From("\"PartyATypes\""))
	return partyATypeQuery{NewQuery(mods...)}
}

// FindPartyAType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPartyAType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*PartyAType, error) {
	partyATypeObj := &PartyAType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"PartyATypes\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, partyATypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from PartyATypes")
	}

	return partyATypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PartyAType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PartyATypes provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(partyATypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	partyATypeInsertCacheMut.RLock()
	cache, cached := partyATypeInsertCache[key]
	partyATypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			partyATypeAllColumns,
			partyATypeColumnsWithDefault,
			partyATypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(partyATypeType, partyATypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(partyATypeType, partyATypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"PartyATypes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"PartyATypes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into PartyATypes")
	}

	if !cached {
		partyATypeInsertCacheMut.Lock()
		partyATypeInsertCache[key] = cache
		partyATypeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PartyAType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PartyAType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	partyATypeUpdateCacheMut.RLock()
	cache, cached := partyATypeUpdateCache[key]
	partyATypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			partyATypeAllColumns,
			partyATypePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update PartyATypes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"PartyATypes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, partyATypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(partyATypeType, partyATypeMapping, append(wl, partyATypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update PartyATypes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for PartyATypes")
	}

	if !cached {
		partyATypeUpdateCacheMut.Lock()
		partyATypeUpdateCache[key] = cache
		partyATypeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q partyATypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for PartyATypes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for PartyATypes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PartyATypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partyATypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"PartyATypes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, partyATypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in partyAType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all partyAType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PartyAType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PartyATypes provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(partyATypeColumnsWithDefault, o)

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

	partyATypeUpsertCacheMut.RLock()
	cache, cached := partyATypeUpsertCache[key]
	partyATypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			partyATypeAllColumns,
			partyATypeColumnsWithDefault,
			partyATypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			partyATypeAllColumns,
			partyATypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert PartyATypes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(partyATypePrimaryKeyColumns))
			copy(conflict, partyATypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"PartyATypes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(partyATypeType, partyATypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(partyATypeType, partyATypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert PartyATypes")
	}

	if !cached {
		partyATypeUpsertCacheMut.Lock()
		partyATypeUpsertCache[key] = cache
		partyATypeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PartyAType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PartyAType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PartyAType provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), partyATypePrimaryKeyMapping)
	sql := "DELETE FROM \"PartyATypes\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from PartyATypes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for PartyATypes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q partyATypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no partyATypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from PartyATypes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PartyATypes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PartyATypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partyATypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"PartyATypes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, partyATypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from partyAType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PartyATypes")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PartyAType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPartyAType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PartyATypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PartyATypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partyATypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"PartyATypes\".* FROM \"PartyATypes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, partyATypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PartyATypeSlice")
	}

	*o = slice

	return nil
}

// PartyATypeExists checks if the PartyAType row exists.
func PartyATypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"PartyATypes\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if PartyATypes exists")
	}

	return exists, nil
}
