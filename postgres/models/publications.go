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

// Publication is an object representing the database table.
type Publication struct {
	ID        int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	UUID      null.String `boil:"uuid" json:"uuid,omitempty" toml:"uuid" yaml:"uuid,omitempty"`
	Title     null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Content   null.String `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	VersionNo null.Int    `boil:"version_no" json:"version_no,omitempty" toml:"version_no" yaml:"version_no,omitempty"`

	R *publicationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publicationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublicationColumns = struct {
	ID        string
	UUID      string
	Title     string
	Content   string
	VersionNo string
}{
	ID:        "Id",
	UUID:      "uuid",
	Title:     "title",
	Content:   "content",
	VersionNo: "version_no",
}

// Generated where

var PublicationWhere = struct {
	ID        whereHelperint
	UUID      whereHelpernull_String
	Title     whereHelpernull_String
	Content   whereHelpernull_String
	VersionNo whereHelpernull_Int
}{
	ID:        whereHelperint{field: "\"publications\".\"Id\""},
	UUID:      whereHelpernull_String{field: "\"publications\".\"uuid\""},
	Title:     whereHelpernull_String{field: "\"publications\".\"title\""},
	Content:   whereHelpernull_String{field: "\"publications\".\"content\""},
	VersionNo: whereHelpernull_Int{field: "\"publications\".\"version_no\""},
}

// PublicationRels is where relationship names are stored.
var PublicationRels = struct {
}{}

// publicationR is where relationships are stored.
type publicationR struct {
}

// NewStruct creates a new relationship struct
func (*publicationR) NewStruct() *publicationR {
	return &publicationR{}
}

// publicationL is where Load methods for each relationship are stored.
type publicationL struct{}

var (
	publicationAllColumns            = []string{"Id", "uuid", "title", "content", "version_no"}
	publicationColumnsWithoutDefault = []string{"Id", "uuid", "title", "content", "version_no"}
	publicationColumnsWithDefault    = []string{}
	publicationPrimaryKeyColumns     = []string{"Id"}
)

type (
	// PublicationSlice is an alias for a slice of pointers to Publication.
	// This should generally be used opposed to []Publication.
	PublicationSlice []*Publication

	publicationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publicationType                 = reflect.TypeOf(&Publication{})
	publicationMapping              = queries.MakeStructMapping(publicationType)
	publicationPrimaryKeyMapping, _ = queries.BindMapping(publicationType, publicationMapping, publicationPrimaryKeyColumns)
	publicationInsertCacheMut       sync.RWMutex
	publicationInsertCache          = make(map[string]insertCache)
	publicationUpdateCacheMut       sync.RWMutex
	publicationUpdateCache          = make(map[string]updateCache)
	publicationUpsertCacheMut       sync.RWMutex
	publicationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single publication record from the query.
func (q publicationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Publication, error) {
	o := &Publication{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publications")
	}

	return o, nil
}

// All returns all Publication records from the query.
func (q publicationQuery) All(ctx context.Context, exec boil.ContextExecutor) (PublicationSlice, error) {
	var o []*Publication

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Publication slice")
	}

	return o, nil
}

// Count returns the count of all Publication records in the query.
func (q publicationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publications rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publicationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publications exists")
	}

	return count > 0, nil
}

// Publications retrieves all the records using an executor.
func Publications(mods ...qm.QueryMod) publicationQuery {
	mods = append(mods, qm.From("\"publications\""))
	return publicationQuery{NewQuery(mods...)}
}

// FindPublication retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublication(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Publication, error) {
	publicationObj := &Publication{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publications\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, publicationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publications")
	}

	return publicationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Publication) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publications provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(publicationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publicationInsertCacheMut.RLock()
	cache, cached := publicationInsertCache[key]
	publicationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publicationAllColumns,
			publicationColumnsWithDefault,
			publicationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publicationType, publicationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publicationType, publicationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publications\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publications\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publications")
	}

	if !cached {
		publicationInsertCacheMut.Lock()
		publicationInsertCache[key] = cache
		publicationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Publication.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Publication) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	publicationUpdateCacheMut.RLock()
	cache, cached := publicationUpdateCache[key]
	publicationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publicationAllColumns,
			publicationPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publications, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publications\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, publicationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publicationType, publicationMapping, append(wl, publicationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publications row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publications")
	}

	if !cached {
		publicationUpdateCacheMut.Lock()
		publicationUpdateCache[key] = cache
		publicationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q publicationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publications")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublicationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publications\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, publicationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publication slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publication")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Publication) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publications provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(publicationColumnsWithDefault, o)

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

	publicationUpsertCacheMut.RLock()
	cache, cached := publicationUpsertCache[key]
	publicationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			publicationAllColumns,
			publicationColumnsWithDefault,
			publicationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			publicationAllColumns,
			publicationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publications, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(publicationPrimaryKeyColumns))
			copy(conflict, publicationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"publications\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(publicationType, publicationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publicationType, publicationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publications")
	}

	if !cached {
		publicationUpsertCacheMut.Lock()
		publicationUpsertCache[key] = cache
		publicationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Publication record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Publication) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Publication provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publicationPrimaryKeyMapping)
	sql := "DELETE FROM \"publications\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publications")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publicationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publicationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publications")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublicationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publications\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publicationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publication slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publications")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Publication) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPublication(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublicationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublicationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publications\".* FROM \"publications\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publicationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublicationSlice")
	}

	*o = slice

	return nil
}

// PublicationExists checks if the Publication row exists.
func PublicationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publications\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publications exists")
	}

	return exists, nil
}
