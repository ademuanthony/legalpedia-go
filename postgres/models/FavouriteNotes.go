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

// FavouriteNote is an object representing the database table.
type FavouriteNote struct {
	ID        string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	NoteId    null.String `boil:"NoteId" json:"NoteId,omitempty" toml:"NoteId" yaml:"NoteId,omitempty"`
	UserId    int64       `boil:"UserId" json:"UserId" toml:"UserId" yaml:"UserId"`
	CreatedAt time.Time   `boil:"CreatedAt" json:"CreatedAt" toml:"CreatedAt" yaml:"CreatedAt"`

	R *favouriteNoteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L favouriteNoteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FavouriteNoteColumns = struct {
	ID        string
	NoteId    string
	UserId    string
	CreatedAt string
}{
	ID:        "Id",
	NoteId:    "NoteId",
	UserId:    "UserId",
	CreatedAt: "CreatedAt",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var FavouriteNoteWhere = struct {
	ID        whereHelperstring
	NoteId    whereHelpernull_String
	UserId    whereHelperint64
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"FavouriteNotes\".\"Id\""},
	NoteId:    whereHelpernull_String{field: "\"FavouriteNotes\".\"NoteId\""},
	UserId:    whereHelperint64{field: "\"FavouriteNotes\".\"UserId\""},
	CreatedAt: whereHelpertime_Time{field: "\"FavouriteNotes\".\"CreatedAt\""},
}

// FavouriteNoteRels is where relationship names are stored.
var FavouriteNoteRels = struct {
}{}

// favouriteNoteR is where relationships are stored.
type favouriteNoteR struct {
}

// NewStruct creates a new relationship struct
func (*favouriteNoteR) NewStruct() *favouriteNoteR {
	return &favouriteNoteR{}
}

// favouriteNoteL is where Load methods for each relationship are stored.
type favouriteNoteL struct{}

var (
	favouriteNoteAllColumns            = []string{"Id", "NoteId", "UserId", "CreatedAt"}
	favouriteNoteColumnsWithoutDefault = []string{"Id", "NoteId", "UserId", "CreatedAt"}
	favouriteNoteColumnsWithDefault    = []string{}
	favouriteNotePrimaryKeyColumns     = []string{"Id"}
)

type (
	// FavouriteNoteSlice is an alias for a slice of pointers to FavouriteNote.
	// This should generally be used opposed to []FavouriteNote.
	FavouriteNoteSlice []*FavouriteNote

	favouriteNoteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	favouriteNoteType                 = reflect.TypeOf(&FavouriteNote{})
	favouriteNoteMapping              = queries.MakeStructMapping(favouriteNoteType)
	favouriteNotePrimaryKeyMapping, _ = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, favouriteNotePrimaryKeyColumns)
	favouriteNoteInsertCacheMut       sync.RWMutex
	favouriteNoteInsertCache          = make(map[string]insertCache)
	favouriteNoteUpdateCacheMut       sync.RWMutex
	favouriteNoteUpdateCache          = make(map[string]updateCache)
	favouriteNoteUpsertCacheMut       sync.RWMutex
	favouriteNoteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single favouriteNote record from the query.
func (q favouriteNoteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FavouriteNote, error) {
	o := &FavouriteNote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for FavouriteNotes")
	}

	return o, nil
}

// All returns all FavouriteNote records from the query.
func (q favouriteNoteQuery) All(ctx context.Context, exec boil.ContextExecutor) (FavouriteNoteSlice, error) {
	var o []*FavouriteNote

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FavouriteNote slice")
	}

	return o, nil
}

// Count returns the count of all FavouriteNote records in the query.
func (q favouriteNoteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count FavouriteNotes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q favouriteNoteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if FavouriteNotes exists")
	}

	return count > 0, nil
}

// FavouriteNotes retrieves all the records using an executor.
func FavouriteNotes(mods ...qm.QueryMod) favouriteNoteQuery {
	mods = append(mods, qm.From("\"FavouriteNotes\""))
	return favouriteNoteQuery{NewQuery(mods...)}
}

// FindFavouriteNote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFavouriteNote(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*FavouriteNote, error) {
	favouriteNoteObj := &FavouriteNote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"FavouriteNotes\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, favouriteNoteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from FavouriteNotes")
	}

	return favouriteNoteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FavouriteNote) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no FavouriteNotes provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(favouriteNoteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	favouriteNoteInsertCacheMut.RLock()
	cache, cached := favouriteNoteInsertCache[key]
	favouriteNoteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			favouriteNoteAllColumns,
			favouriteNoteColumnsWithDefault,
			favouriteNoteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"FavouriteNotes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"FavouriteNotes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into FavouriteNotes")
	}

	if !cached {
		favouriteNoteInsertCacheMut.Lock()
		favouriteNoteInsertCache[key] = cache
		favouriteNoteInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the FavouriteNote.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FavouriteNote) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	favouriteNoteUpdateCacheMut.RLock()
	cache, cached := favouriteNoteUpdateCache[key]
	favouriteNoteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			favouriteNoteAllColumns,
			favouriteNotePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update FavouriteNotes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"FavouriteNotes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, favouriteNotePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, append(wl, favouriteNotePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update FavouriteNotes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for FavouriteNotes")
	}

	if !cached {
		favouriteNoteUpdateCacheMut.Lock()
		favouriteNoteUpdateCache[key] = cache
		favouriteNoteUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q favouriteNoteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for FavouriteNotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for FavouriteNotes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FavouriteNoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favouriteNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"FavouriteNotes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, favouriteNotePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in favouriteNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all favouriteNote")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FavouriteNote) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no FavouriteNotes provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(favouriteNoteColumnsWithDefault, o)

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

	favouriteNoteUpsertCacheMut.RLock()
	cache, cached := favouriteNoteUpsertCache[key]
	favouriteNoteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			favouriteNoteAllColumns,
			favouriteNoteColumnsWithDefault,
			favouriteNoteColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			favouriteNoteAllColumns,
			favouriteNotePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert FavouriteNotes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(favouriteNotePrimaryKeyColumns))
			copy(conflict, favouriteNotePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"FavouriteNotes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(favouriteNoteType, favouriteNoteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert FavouriteNotes")
	}

	if !cached {
		favouriteNoteUpsertCacheMut.Lock()
		favouriteNoteUpsertCache[key] = cache
		favouriteNoteUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single FavouriteNote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FavouriteNote) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FavouriteNote provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), favouriteNotePrimaryKeyMapping)
	sql := "DELETE FROM \"FavouriteNotes\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from FavouriteNotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for FavouriteNotes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q favouriteNoteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no favouriteNoteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from FavouriteNotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for FavouriteNotes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FavouriteNoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favouriteNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"FavouriteNotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, favouriteNotePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from favouriteNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for FavouriteNotes")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FavouriteNote) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFavouriteNote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FavouriteNoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FavouriteNoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favouriteNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"FavouriteNotes\".* FROM \"FavouriteNotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, favouriteNotePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FavouriteNoteSlice")
	}

	*o = slice

	return nil
}

// FavouriteNoteExists checks if the FavouriteNote row exists.
func FavouriteNoteExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"FavouriteNotes\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if FavouriteNotes exists")
	}

	return exists, nil
}
