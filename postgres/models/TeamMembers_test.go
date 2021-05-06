// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTeamMembers(t *testing.T) {
	t.Parallel()

	query := TeamMembers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTeamMembersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMembersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TeamMembers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMembersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamMemberSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMembersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TeamMemberExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TeamMember exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TeamMemberExists to return true, but got false.")
	}
}

func testTeamMembersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	teamMemberFound, err := FindTeamMember(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if teamMemberFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTeamMembersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TeamMembers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTeamMembersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TeamMembers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTeamMembersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	teamMemberOne := &TeamMember{}
	teamMemberTwo := &TeamMember{}
	if err = randomize.Struct(seed, teamMemberOne, teamMemberDBTypes, false, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}
	if err = randomize.Struct(seed, teamMemberTwo, teamMemberDBTypes, false, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = teamMemberOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamMemberTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamMembers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTeamMembersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	teamMemberOne := &TeamMember{}
	teamMemberTwo := &TeamMember{}
	if err = randomize.Struct(seed, teamMemberOne, teamMemberDBTypes, false, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}
	if err = randomize.Struct(seed, teamMemberTwo, teamMemberDBTypes, false, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = teamMemberOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamMemberTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTeamMembersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamMembersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(teamMemberColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamMembersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTeamMembersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamMemberSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTeamMembersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamMembers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	teamMemberDBTypes = map[string]string{`ID`: `text`, `TeamId`: `text`, `UserId`: `bigint`, `Role`: `integer`}
	_                 = bytes.MinRead
)

func testTeamMembersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(teamMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(teamMemberAllColumns) == len(teamMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTeamMembersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(teamMemberAllColumns) == len(teamMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamMember{}
	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamMemberDBTypes, true, teamMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(teamMemberAllColumns, teamMemberPrimaryKeyColumns) {
		fields = teamMemberAllColumns
	} else {
		fields = strmangle.SetComplement(
			teamMemberAllColumns,
			teamMemberPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TeamMemberSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTeamMembersUpsert(t *testing.T) {
	t.Parallel()

	if len(teamMemberAllColumns) == len(teamMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TeamMember{}
	if err = randomize.Struct(seed, &o, teamMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamMember: %s", err)
	}

	count, err := TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, teamMemberDBTypes, false, teamMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMember struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamMember: %s", err)
	}

	count, err = TeamMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
