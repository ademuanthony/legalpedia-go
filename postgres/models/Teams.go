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

// Team is an object representing the database table.
type Team struct {
	ID            string      `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	Name          null.String `boil:"Name" json:"Name,omitempty" toml:"Name" yaml:"Name,omitempty"`
	Description   null.String `boil:"Description" json:"Description,omitempty" toml:"Description" yaml:"Description,omitempty"`
	Logo          null.String `boil:"Logo" json:"Logo,omitempty" toml:"Logo" yaml:"Logo,omitempty"`
	CreationTime  time.Time   `boil:"CreationTime" json:"CreationTime" toml:"CreationTime" yaml:"CreationTime"`
	CreatorUserId null.Int64  `boil:"CreatorUserId" json:"CreatorUserId,omitempty" toml:"CreatorUserId" yaml:"CreatorUserId,omitempty"`

	R *teamR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teamL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeamColumns = struct {
	ID            string
	Name          string
	Description   string
	Logo          string
	CreationTime  string
	CreatorUserId string
}{
	ID:            "Id",
	Name:          "Name",
	Description:   "Description",
	Logo:          "Logo",
	CreationTime:  "CreationTime",
	CreatorUserId: "CreatorUserId",
}

// Generated where

var TeamWhere = struct {
	ID            whereHelperstring
	Name          whereHelpernull_String
	Description   whereHelpernull_String
	Logo          whereHelpernull_String
	CreationTime  whereHelpertime_Time
	CreatorUserId whereHelpernull_Int64
}{
	ID:            whereHelperstring{field: "\"Teams\".\"Id\""},
	Name:          whereHelpernull_String{field: "\"Teams\".\"Name\""},
	Description:   whereHelpernull_String{field: "\"Teams\".\"Description\""},
	Logo:          whereHelpernull_String{field: "\"Teams\".\"Logo\""},
	CreationTime:  whereHelpertime_Time{field: "\"Teams\".\"CreationTime\""},
	CreatorUserId: whereHelpernull_Int64{field: "\"Teams\".\"CreatorUserId\""},
}

// TeamRels is where relationship names are stored.
var TeamRels = struct {
	TeamIdTeamMembers string
}{
	TeamIdTeamMembers: "TeamIdTeamMembers",
}

// teamR is where relationships are stored.
type teamR struct {
	TeamIdTeamMembers TeamMemberSlice `boil:"TeamIdTeamMembers" json:"TeamIdTeamMembers" toml:"TeamIdTeamMembers" yaml:"TeamIdTeamMembers"`
}

// NewStruct creates a new relationship struct
func (*teamR) NewStruct() *teamR {
	return &teamR{}
}

// teamL is where Load methods for each relationship are stored.
type teamL struct{}

var (
	teamAllColumns            = []string{"Id", "Name", "Description", "Logo", "CreationTime", "CreatorUserId"}
	teamColumnsWithoutDefault = []string{"Id", "Name", "Description", "Logo", "CreatorUserId"}
	teamColumnsWithDefault    = []string{"CreationTime"}
	teamPrimaryKeyColumns     = []string{"Id"}
)

type (
	// TeamSlice is an alias for a slice of pointers to Team.
	// This should generally be used opposed to []Team.
	TeamSlice []*Team

	teamQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teamType                 = reflect.TypeOf(&Team{})
	teamMapping              = queries.MakeStructMapping(teamType)
	teamPrimaryKeyMapping, _ = queries.BindMapping(teamType, teamMapping, teamPrimaryKeyColumns)
	teamInsertCacheMut       sync.RWMutex
	teamInsertCache          = make(map[string]insertCache)
	teamUpdateCacheMut       sync.RWMutex
	teamUpdateCache          = make(map[string]updateCache)
	teamUpsertCacheMut       sync.RWMutex
	teamUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single team record from the query.
func (q teamQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Team, error) {
	o := &Team{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Teams")
	}

	return o, nil
}

// All returns all Team records from the query.
func (q teamQuery) All(ctx context.Context, exec boil.ContextExecutor) (TeamSlice, error) {
	var o []*Team

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Team slice")
	}

	return o, nil
}

// Count returns the count of all Team records in the query.
func (q teamQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Teams rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teamQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Teams exists")
	}

	return count > 0, nil
}

// TeamIdTeamMembers retrieves all the TeamMember's TeamMembers with an executor via TeamId column.
func (o *Team) TeamIdTeamMembers(mods ...qm.QueryMod) teamMemberQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"TeamMembers\".\"TeamId\"=?", o.ID),
	)

	query := TeamMembers(queryMods...)
	queries.SetFrom(query.Query, "\"TeamMembers\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"TeamMembers\".*"})
	}

	return query
}

// LoadTeamIdTeamMembers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (teamL) LoadTeamIdTeamMembers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTeam interface{}, mods queries.Applicator) error {
	var slice []*Team
	var object *Team

	if singular {
		object = maybeTeam.(*Team)
	} else {
		slice = *maybeTeam.(*[]*Team)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &teamR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &teamR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`TeamMembers`),
		qm.WhereIn(`TeamMembers.TeamId in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TeamMembers")
	}

	var resultSlice []*TeamMember
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TeamMembers")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on TeamMembers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for TeamMembers")
	}

	if singular {
		object.R.TeamIdTeamMembers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &teamMemberR{}
			}
			foreign.R.TeamIdTeam = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.TeamId) {
				local.R.TeamIdTeamMembers = append(local.R.TeamIdTeamMembers, foreign)
				if foreign.R == nil {
					foreign.R = &teamMemberR{}
				}
				foreign.R.TeamIdTeam = local
				break
			}
		}
	}

	return nil
}

// AddTeamIdTeamMembers adds the given related objects to the existing relationships
// of the Team, optionally inserting them as new records.
// Appends related to o.R.TeamIdTeamMembers.
// Sets related.R.TeamIdTeam appropriately.
func (o *Team) AddTeamIdTeamMembers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TeamMember) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.TeamId, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"TeamMembers\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"TeamId"}),
				strmangle.WhereClause("\"", "\"", 2, teamMemberPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.TeamId, o.ID)
		}
	}

	if o.R == nil {
		o.R = &teamR{
			TeamIdTeamMembers: related,
		}
	} else {
		o.R.TeamIdTeamMembers = append(o.R.TeamIdTeamMembers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &teamMemberR{
				TeamIdTeam: o,
			}
		} else {
			rel.R.TeamIdTeam = o
		}
	}
	return nil
}

// SetTeamIdTeamMembers removes all previously related items of the
// Team replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.TeamIdTeam's TeamIdTeamMembers accordingly.
// Replaces o.R.TeamIdTeamMembers with related.
// Sets related.R.TeamIdTeam's TeamIdTeamMembers accordingly.
func (o *Team) SetTeamIdTeamMembers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TeamMember) error {
	query := "update \"TeamMembers\" set \"TeamId\" = null where \"TeamId\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TeamIdTeamMembers {
			queries.SetScanner(&rel.TeamId, nil)
			if rel.R == nil {
				continue
			}

			rel.R.TeamIdTeam = nil
		}

		o.R.TeamIdTeamMembers = nil
	}
	return o.AddTeamIdTeamMembers(ctx, exec, insert, related...)
}

// RemoveTeamIdTeamMembers relationships from objects passed in.
// Removes related items from R.TeamIdTeamMembers (uses pointer comparison, removal does not keep order)
// Sets related.R.TeamIdTeam.
func (o *Team) RemoveTeamIdTeamMembers(ctx context.Context, exec boil.ContextExecutor, related ...*TeamMember) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.TeamId, nil)
		if rel.R != nil {
			rel.R.TeamIdTeam = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("TeamId")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TeamIdTeamMembers {
			if rel != ri {
				continue
			}

			ln := len(o.R.TeamIdTeamMembers)
			if ln > 1 && i < ln-1 {
				o.R.TeamIdTeamMembers[i] = o.R.TeamIdTeamMembers[ln-1]
			}
			o.R.TeamIdTeamMembers = o.R.TeamIdTeamMembers[:ln-1]
			break
		}
	}

	return nil
}

// Teams retrieves all the records using an executor.
func Teams(mods ...qm.QueryMod) teamQuery {
	mods = append(mods, qm.From("\"Teams\""))
	return teamQuery{NewQuery(mods...)}
}

// FindTeam retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeam(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Team, error) {
	teamObj := &Team{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Teams\" where \"Id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, teamObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Teams")
	}

	return teamObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Team) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Teams provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(teamColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teamInsertCacheMut.RLock()
	cache, cached := teamInsertCache[key]
	teamInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teamAllColumns,
			teamColumnsWithDefault,
			teamColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teamType, teamMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Teams\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Teams\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into Teams")
	}

	if !cached {
		teamInsertCacheMut.Lock()
		teamInsertCache[key] = cache
		teamInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Team.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Team) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	teamUpdateCacheMut.RLock()
	cache, cached := teamUpdateCache[key]
	teamUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teamAllColumns,
			teamPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Teams, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Teams\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, teamPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, append(wl, teamPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Teams row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Teams")
	}

	if !cached {
		teamUpdateCacheMut.Lock()
		teamUpdateCache[key] = cache
		teamUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q teamQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Teams")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Teams")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeamSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Teams\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, teamPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in team slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all team")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Team) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Teams provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(teamColumnsWithDefault, o)

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

	teamUpsertCacheMut.RLock()
	cache, cached := teamUpsertCache[key]
	teamUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teamAllColumns,
			teamColumnsWithDefault,
			teamColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			teamAllColumns,
			teamPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert Teams, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(teamPrimaryKeyColumns))
			copy(conflict, teamPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"Teams\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teamType, teamMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert Teams")
	}

	if !cached {
		teamUpsertCacheMut.Lock()
		teamUpsertCache[key] = cache
		teamUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Team record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Team) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Team provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teamPrimaryKeyMapping)
	sql := "DELETE FROM \"Teams\" WHERE \"Id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Teams")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Teams")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teamQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no teamQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Teams")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Teams")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeamSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Teams\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from team slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Teams")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Team) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTeam(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeamSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeamSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Teams\".* FROM \"Teams\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TeamSlice")
	}

	*o = slice

	return nil
}

// TeamExists checks if the Team row exists.
func TeamExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Teams\" where \"Id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Teams exists")
	}

	return exists, nil
}
