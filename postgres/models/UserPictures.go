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

// UserPicture is an object representing the database table.
type UserPicture struct {
	ID      string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	UserId  int64       `boil:"UserId" json:"UserId" toml:"UserId" yaml:"UserId"`
	Content null.String `boil:"Content" json:"Content,omitempty" toml:"Content" yaml:"Content,omitempty"`

	R *userPictureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userPictureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserPictureColumns = struct {
	ID      string
	UserId  string
	Content string
}{
	ID:      "Id",
	UserId:  "UserId",
	Content: "Content",
}

// Generated where

var UserPictureWhere = struct {
	ID      whereHelperstring
	UserId  whereHelperint64
	Content whereHelpernull_String
}{
	ID:      whereHelperstring{field: "\"UserPictures\".\"Id\""},
	UserId:  whereHelperint64{field: "\"UserPictures\".\"UserId\""},
	Content: whereHelpernull_String{field: "\"UserPictures\".\"Content\""},
}

// UserPictureRels is where relationship names are stored.
var UserPictureRels = struct {
}{}

// userPictureR is where relationships are stored.
type userPictureR struct {
}

// NewStruct creates a new relationship struct
func (*userPictureR) NewStruct() *userPictureR {
	return &userPictureR{}
}

// userPictureL is where Load methods for each relationship are stored.
type userPictureL struct{}

var (
	userPictureAllColumns            = []string{"Id", "UserId", "Content"}
	userPictureColumnsWithoutDefault = []string{"Id", "UserId", "Content"}
	userPictureColumnsWithDefault    = []string{}
	userPicturePrimaryKeyColumns     = []string{"Id"}
)

type (
	// UserPictureSlice is an alias for a slice of pointers to UserPicture.
	// This should generally be used opposed to []UserPicture.
	UserPictureSlice []*UserPicture

	userPictureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userPictureType                 = reflect.TypeOf(&UserPicture{})
	userPictureMapping              = queries.MakeStructMapping(userPictureType)
	userPicturePrimaryKeyMapping, _ = queries.BindMapping(userPictureType, userPictureMapping, userPicturePrimaryKeyColumns)
	userPictureInsertCacheMut       sync.RWMutex
	userPictureInsertCache          = make(map[string]insertCache)
	userPictureUpdateCacheMut       sync.RWMutex
	userPictureUpdateCache          = make(map[string]updateCache)
	userPictureUpsertCacheMut       sync.RWMutex
	userPictureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single userPicture record from the query.
func (q userPictureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserPicture, error) {
	o := &UserPicture{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for UserPictures")
	}

	return o, nil
}

// All returns all UserPicture records from the query.
func (q userPictureQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserPictureSlice, error) {
	var o []*UserPicture

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserPicture slice")
	}

	return o, nil
}

// Count returns the count of all UserPicture records in the query.
func (q userPictureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count UserPictures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userPictureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if UserPictures exists")
	}

	return count > 0, nil
}

// UserPictures retrieves all the records using an executor.
func UserPictures(mods ...qm.QueryMod) userPictureQuery {
	mods = append(mods, qm.From("\"UserPictures\""))
	return userPictureQuery{NewQuery(mods...)}
}

// FindUserPicture retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserPicture(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*UserPicture, error) {
	userPictureObj := &UserPicture{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"UserPictures\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userPictureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from UserPictures")
	}

	return userPictureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserPicture) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no UserPictures provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(userPictureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userPictureInsertCacheMut.RLock()
	cache, cached := userPictureInsertCache[key]
	userPictureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userPictureAllColumns,
			userPictureColumnsWithDefault,
			userPictureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userPictureType, userPictureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userPictureType, userPictureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"UserPictures\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"UserPictures\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into UserPictures")
	}

	if !cached {
		userPictureInsertCacheMut.Lock()
		userPictureInsertCache[key] = cache
		userPictureInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the UserPicture.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserPicture) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	userPictureUpdateCacheMut.RLock()
	cache, cached := userPictureUpdateCache[key]
	userPictureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userPictureAllColumns,
			userPicturePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update UserPictures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"UserPictures\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userPicturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userPictureType, userPictureMapping, append(wl, userPicturePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update UserPictures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for UserPictures")
	}

	if !cached {
		userPictureUpdateCacheMut.Lock()
		userPictureUpdateCache[key] = cache
		userPictureUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q userPictureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for UserPictures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for UserPictures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserPictureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"UserPictures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userPicturePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userPicture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userPicture")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserPicture) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no UserPictures provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(userPictureColumnsWithDefault, o)

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

	userPictureUpsertCacheMut.RLock()
	cache, cached := userPictureUpsertCache[key]
	userPictureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userPictureAllColumns,
			userPictureColumnsWithDefault,
			userPictureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userPictureAllColumns,
			userPicturePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert UserPictures, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userPicturePrimaryKeyColumns))
			copy(conflict, userPicturePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"UserPictures\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userPictureType, userPictureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userPictureType, userPictureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert UserPictures")
	}

	if !cached {
		userPictureUpsertCacheMut.Lock()
		userPictureUpsertCache[key] = cache
		userPictureUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single UserPicture record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserPicture) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserPicture provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPicturePrimaryKeyMapping)
	sql := "DELETE FROM \"UserPictures\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from UserPictures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for UserPictures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userPictureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userPictureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from UserPictures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for UserPictures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserPictureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"UserPictures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPicturePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userPicture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for UserPictures")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserPicture) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserPicture(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserPictureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserPictureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"UserPictures\".* FROM \"UserPictures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPicturePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserPictureSlice")
	}

	*o = slice

	return nil
}

// UserPictureExists checks if the UserPicture row exists.
func UserPictureExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"UserPictures\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if UserPictures exists")
	}

	return exists, nil
}
