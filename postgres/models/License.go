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

// License is an object representing the database table.
type License struct {
	ID                   int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	CustomerId           null.Int    `boil:"CustomerId" json:"CustomerId,omitempty" toml:"CustomerId" yaml:"CustomerId,omitempty"`
	SystemId             null.String `boil:"SystemId" json:"SystemId,omitempty" toml:"SystemId" yaml:"SystemId,omitempty"`
	LicensedDays         null.Int    `boil:"LicensedDays" json:"LicensedDays,omitempty" toml:"LicensedDays" yaml:"LicensedDays,omitempty"`
	UnlockDate           null.Time   `boil:"UnlockDate" json:"UnlockDate,omitempty" toml:"UnlockDate" yaml:"UnlockDate,omitempty"`
	Package              null.String `boil:"Package" json:"Package,omitempty" toml:"Package" yaml:"Package,omitempty"`
	LastModificationTime null.Time   `boil:"LastModificationTime" json:"LastModificationTime,omitempty" toml:"LastModificationTime" yaml:"LastModificationTime,omitempty"`
	LastModifierUserId   null.Int64  `boil:"LastModifierUserId" json:"LastModifierUserId,omitempty" toml:"LastModifierUserId" yaml:"LastModifierUserId,omitempty"`
	CreationTime         null.Time   `boil:"CreationTime" json:"CreationTime,omitempty" toml:"CreationTime" yaml:"CreationTime,omitempty"`
	CreatorUserId        null.Int64  `boil:"CreatorUserId" json:"CreatorUserId,omitempty" toml:"CreatorUserId" yaml:"CreatorUserId,omitempty"`
	TenantId             null.Int    `boil:"TenantId" json:"TenantId,omitempty" toml:"TenantId" yaml:"TenantId,omitempty"`
	UpdateLicensedDays   null.Int    `boil:"UpdateLicensedDays" json:"UpdateLicensedDays,omitempty" toml:"UpdateLicensedDays" yaml:"UpdateLicensedDays,omitempty"`
	MobileSystemId       null.String `boil:"MobileSystemId" json:"MobileSystemId,omitempty" toml:"MobileSystemId" yaml:"MobileSystemId,omitempty"`
	ExpDate              null.Time   `boil:"ExpDate" json:"ExpDate,omitempty" toml:"ExpDate" yaml:"ExpDate,omitempty"`

	R *licenseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L licenseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LicenseColumns = struct {
	ID                   string
	CustomerId           string
	SystemId             string
	LicensedDays         string
	UnlockDate           string
	Package              string
	LastModificationTime string
	LastModifierUserId   string
	CreationTime         string
	CreatorUserId        string
	TenantId             string
	UpdateLicensedDays   string
	MobileSystemId       string
	ExpDate              string
}{
	ID:                   "Id",
	CustomerId:           "CustomerId",
	SystemId:             "SystemId",
	LicensedDays:         "LicensedDays",
	UnlockDate:           "UnlockDate",
	Package:              "Package",
	LastModificationTime: "LastModificationTime",
	LastModifierUserId:   "LastModifierUserId",
	CreationTime:         "CreationTime",
	CreatorUserId:        "CreatorUserId",
	TenantId:             "TenantId",
	UpdateLicensedDays:   "UpdateLicensedDays",
	MobileSystemId:       "MobileSystemId",
	ExpDate:              "ExpDate",
}

// Generated where

var LicenseWhere = struct {
	ID                   whereHelperint
	CustomerId           whereHelpernull_Int
	SystemId             whereHelpernull_String
	LicensedDays         whereHelpernull_Int
	UnlockDate           whereHelpernull_Time
	Package              whereHelpernull_String
	LastModificationTime whereHelpernull_Time
	LastModifierUserId   whereHelpernull_Int64
	CreationTime         whereHelpernull_Time
	CreatorUserId        whereHelpernull_Int64
	TenantId             whereHelpernull_Int
	UpdateLicensedDays   whereHelpernull_Int
	MobileSystemId       whereHelpernull_String
	ExpDate              whereHelpernull_Time
}{
	ID:                   whereHelperint{field: "\"License\".\"Id\""},
	CustomerId:           whereHelpernull_Int{field: "\"License\".\"CustomerId\""},
	SystemId:             whereHelpernull_String{field: "\"License\".\"SystemId\""},
	LicensedDays:         whereHelpernull_Int{field: "\"License\".\"LicensedDays\""},
	UnlockDate:           whereHelpernull_Time{field: "\"License\".\"UnlockDate\""},
	Package:              whereHelpernull_String{field: "\"License\".\"Package\""},
	LastModificationTime: whereHelpernull_Time{field: "\"License\".\"LastModificationTime\""},
	LastModifierUserId:   whereHelpernull_Int64{field: "\"License\".\"LastModifierUserId\""},
	CreationTime:         whereHelpernull_Time{field: "\"License\".\"CreationTime\""},
	CreatorUserId:        whereHelpernull_Int64{field: "\"License\".\"CreatorUserId\""},
	TenantId:             whereHelpernull_Int{field: "\"License\".\"TenantId\""},
	UpdateLicensedDays:   whereHelpernull_Int{field: "\"License\".\"UpdateLicensedDays\""},
	MobileSystemId:       whereHelpernull_String{field: "\"License\".\"MobileSystemId\""},
	ExpDate:              whereHelpernull_Time{field: "\"License\".\"ExpDate\""},
}

// LicenseRels is where relationship names are stored.
var LicenseRels = struct {
}{}

// licenseR is where relationships are stored.
type licenseR struct {
}

// NewStruct creates a new relationship struct
func (*licenseR) NewStruct() *licenseR {
	return &licenseR{}
}

// licenseL is where Load methods for each relationship are stored.
type licenseL struct{}

var (
	licenseAllColumns            = []string{"Id", "CustomerId", "SystemId", "LicensedDays", "UnlockDate", "Package", "LastModificationTime", "LastModifierUserId", "CreationTime", "CreatorUserId", "TenantId", "UpdateLicensedDays", "MobileSystemId", "ExpDate"}
	licenseColumnsWithoutDefault = []string{"Id", "CustomerId", "SystemId", "LicensedDays", "UnlockDate", "Package", "LastModificationTime", "LastModifierUserId", "CreationTime", "CreatorUserId", "TenantId", "UpdateLicensedDays", "MobileSystemId", "ExpDate"}
	licenseColumnsWithDefault    = []string{}
	licensePrimaryKeyColumns     = []string{"Id"}
)

type (
	// LicenseSlice is an alias for a slice of pointers to License.
	// This should generally be used opposed to []License.
	LicenseSlice []*License

	licenseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	licenseType                 = reflect.TypeOf(&License{})
	licenseMapping              = queries.MakeStructMapping(licenseType)
	licensePrimaryKeyMapping, _ = queries.BindMapping(licenseType, licenseMapping, licensePrimaryKeyColumns)
	licenseInsertCacheMut       sync.RWMutex
	licenseInsertCache          = make(map[string]insertCache)
	licenseUpdateCacheMut       sync.RWMutex
	licenseUpdateCache          = make(map[string]updateCache)
	licenseUpsertCacheMut       sync.RWMutex
	licenseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single license record from the query.
func (q licenseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*License, error) {
	o := &License{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for License")
	}

	return o, nil
}

// All returns all License records from the query.
func (q licenseQuery) All(ctx context.Context, exec boil.ContextExecutor) (LicenseSlice, error) {
	var o []*License

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to License slice")
	}

	return o, nil
}

// Count returns the count of all License records in the query.
func (q licenseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count License rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q licenseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if License exists")
	}

	return count > 0, nil
}

// Licenses retrieves all the records using an executor.
func Licenses(mods ...qm.QueryMod) licenseQuery {
	mods = append(mods, qm.From("\"License\""))
	return licenseQuery{NewQuery(mods...)}
}

// FindLicense retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLicense(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*License, error) {
	licenseObj := &License{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"License\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, licenseObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from License")
	}

	return licenseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *License) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no License provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(licenseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	licenseInsertCacheMut.RLock()
	cache, cached := licenseInsertCache[key]
	licenseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			licenseAllColumns,
			licenseColumnsWithDefault,
			licenseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(licenseType, licenseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(licenseType, licenseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"License\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"License\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into License")
	}

	if !cached {
		licenseInsertCacheMut.Lock()
		licenseInsertCache[key] = cache
		licenseInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the License.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *License) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	licenseUpdateCacheMut.RLock()
	cache, cached := licenseUpdateCache[key]
	licenseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			licenseAllColumns,
			licensePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update License, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"License\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, licensePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(licenseType, licenseMapping, append(wl, licensePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update License row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for License")
	}

	if !cached {
		licenseUpdateCacheMut.Lock()
		licenseUpdateCache[key] = cache
		licenseUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q licenseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for License")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for License")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LicenseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), licensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"License\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, licensePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in license slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all license")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *License) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no License provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(licenseColumnsWithDefault, o)

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

	licenseUpsertCacheMut.RLock()
	cache, cached := licenseUpsertCache[key]
	licenseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			licenseAllColumns,
			licenseColumnsWithDefault,
			licenseColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			licenseAllColumns,
			licensePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert License, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(licensePrimaryKeyColumns))
			copy(conflict, licensePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"License\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(licenseType, licenseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(licenseType, licenseMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert License")
	}

	if !cached {
		licenseUpsertCacheMut.Lock()
		licenseUpsertCache[key] = cache
		licenseUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single License record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *License) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no License provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), licensePrimaryKeyMapping)
	sql := "DELETE FROM \"License\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from License")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for License")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q licenseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no licenseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from License")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for License")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LicenseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), licensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"License\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, licensePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from license slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for License")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *License) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLicense(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LicenseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LicenseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), licensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"License\".* FROM \"License\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, licensePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LicenseSlice")
	}

	*o = slice

	return nil
}

// LicenseExists checks if the License row exists.
func LicenseExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"License\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if License exists")
	}

	return exists, nil
}
