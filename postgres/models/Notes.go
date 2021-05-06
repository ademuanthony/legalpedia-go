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

// Note is an object representing the database table.
type Note struct {
	ID         string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	TeamId     null.String `boil:"TeamId" json:"TeamId,omitempty" toml:"TeamId" yaml:"TeamId,omitempty"`
	CreatorId  int64       `boil:"CreatorId" json:"CreatorId" toml:"CreatorId" yaml:"CreatorId"`
	Title      null.String `boil:"Title" json:"Title,omitempty" toml:"Title" yaml:"Title,omitempty"`
	Summary    null.String `boil:"Summary" json:"Summary,omitempty" toml:"Summary" yaml:"Summary,omitempty"`
	Body       null.String `boil:"Body" json:"Body,omitempty" toml:"Body" yaml:"Body,omitempty"`
	Image      null.String `boil:"Image" json:"Image,omitempty" toml:"Image" yaml:"Image,omitempty"`
	AccessType int         `boil:"AccessType" json:"AccessType" toml:"AccessType" yaml:"AccessType"`
	Labels     null.String `boil:"Labels" json:"Labels,omitempty" toml:"Labels" yaml:"Labels,omitempty"`
	CreatedAt  time.Time   `boil:"CreatedAt" json:"CreatedAt" toml:"CreatedAt" yaml:"CreatedAt"`

	R *noteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L noteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NoteColumns = struct {
	ID         string
	TeamId     string
	CreatorId  string
	Title      string
	Summary    string
	Body       string
	Image      string
	AccessType string
	Labels     string
	CreatedAt  string
}{
	ID:         "Id",
	TeamId:     "TeamId",
	CreatorId:  "CreatorId",
	Title:      "Title",
	Summary:    "Summary",
	Body:       "Body",
	Image:      "Image",
	AccessType: "AccessType",
	Labels:     "Labels",
	CreatedAt:  "CreatedAt",
}

// Generated where

var NoteWhere = struct {
	ID         whereHelperstring
	TeamId     whereHelpernull_String
	CreatorId  whereHelperint64
	Title      whereHelpernull_String
	Summary    whereHelpernull_String
	Body       whereHelpernull_String
	Image      whereHelpernull_String
	AccessType whereHelperint
	Labels     whereHelpernull_String
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "\"Notes\".\"Id\""},
	TeamId:     whereHelpernull_String{field: "\"Notes\".\"TeamId\""},
	CreatorId:  whereHelperint64{field: "\"Notes\".\"CreatorId\""},
	Title:      whereHelpernull_String{field: "\"Notes\".\"Title\""},
	Summary:    whereHelpernull_String{field: "\"Notes\".\"Summary\""},
	Body:       whereHelpernull_String{field: "\"Notes\".\"Body\""},
	Image:      whereHelpernull_String{field: "\"Notes\".\"Image\""},
	AccessType: whereHelperint{field: "\"Notes\".\"AccessType\""},
	Labels:     whereHelpernull_String{field: "\"Notes\".\"Labels\""},
	CreatedAt:  whereHelpertime_Time{field: "\"Notes\".\"CreatedAt\""},
}

// NoteRels is where relationship names are stored.
var NoteRels = struct {
}{}

// noteR is where relationships are stored.
type noteR struct {
}

// NewStruct creates a new relationship struct
func (*noteR) NewStruct() *noteR {
	return &noteR{}
}

// noteL is where Load methods for each relationship are stored.
type noteL struct{}

var (
	noteAllColumns            = []string{"Id", "TeamId", "CreatorId", "Title", "Summary", "Body", "Image", "AccessType", "Labels", "CreatedAt"}
	noteColumnsWithoutDefault = []string{"Id", "TeamId", "CreatorId", "Title", "Summary", "Body", "Image", "AccessType", "Labels", "CreatedAt"}
	noteColumnsWithDefault    = []string{}
	notePrimaryKeyColumns     = []string{"Id"}
)

type (
	// NoteSlice is an alias for a slice of pointers to Note.
	// This should generally be used opposed to []Note.
	NoteSlice []*Note

	noteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	noteType                 = reflect.TypeOf(&Note{})
	noteMapping              = queries.MakeStructMapping(noteType)
	notePrimaryKeyMapping, _ = queries.BindMapping(noteType, noteMapping, notePrimaryKeyColumns)
	noteInsertCacheMut       sync.RWMutex
	noteInsertCache          = make(map[string]insertCache)
	noteUpdateCacheMut       sync.RWMutex
	noteUpdateCache          = make(map[string]updateCache)
	noteUpsertCacheMut       sync.RWMutex
	noteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single note record from the query.
func (q noteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Note, error) {
	o := &Note{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Notes")
	}

	return o, nil
}

// All returns all Note records from the query.
func (q noteQuery) All(ctx context.Context, exec boil.ContextExecutor) (NoteSlice, error) {
	var o []*Note

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Note slice")
	}

	return o, nil
}

// Count returns the count of all Note records in the query.
func (q noteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Notes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q noteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Notes exists")
	}

	return count > 0, nil
}

// Notes retrieves all the records using an executor.
func Notes(mods ...qm.QueryMod) noteQuery {
	mods = append(mods, qm.From("\"Notes\""))
	return noteQuery{NewQuery(mods...)}
}

// FindNote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNote(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Note, error) {
	noteObj := &Note{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Notes\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, noteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Notes")
	}

	return noteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Note) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Notes provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(noteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	noteInsertCacheMut.RLock()
	cache, cached := noteInsertCache[key]
	noteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			noteAllColumns,
			noteColumnsWithDefault,
			noteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(noteType, noteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(noteType, noteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Notes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Notes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into Notes")
	}

	if !cached {
		noteInsertCacheMut.Lock()
		noteInsertCache[key] = cache
		noteInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Note.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Note) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	noteUpdateCacheMut.RLock()
	cache, cached := noteUpdateCache[key]
	noteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			noteAllColumns,
			notePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Notes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Notes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, notePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(noteType, noteMapping, append(wl, notePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Notes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Notes")
	}

	if !cached {
		noteUpdateCacheMut.Lock()
		noteUpdateCache[key] = cache
		noteUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q noteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Notes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Notes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Notes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, notePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in note slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all note")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Note) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Notes provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(noteColumnsWithDefault, o)

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

	noteUpsertCacheMut.RLock()
	cache, cached := noteUpsertCache[key]
	noteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			noteAllColumns,
			noteColumnsWithDefault,
			noteColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			noteAllColumns,
			notePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert Notes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(notePrimaryKeyColumns))
			copy(conflict, notePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"Notes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(noteType, noteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(noteType, noteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert Notes")
	}

	if !cached {
		noteUpsertCacheMut.Lock()
		noteUpsertCache[key] = cache
		noteUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Note record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Note) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Note provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), notePrimaryKeyMapping)
	sql := "DELETE FROM \"Notes\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Notes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Notes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q noteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no noteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Notes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Notes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Notes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, notePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from note slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Notes")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Note) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Notes\".* FROM \"Notes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, notePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NoteSlice")
	}

	*o = slice

	return nil
}

// NoteExists checks if the Note row exists.
func NoteExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Notes\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Notes exists")
	}

	return exists, nil
}
