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

func testCourts(t *testing.T) {
	t.Parallel()

	query := Courts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCourtsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
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

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCourtsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Courts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCourtsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CourtSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCourtsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CourtExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Court exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CourtExists to return true, but got false.")
	}
}

func testCourtsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	courtFound, err := FindCourt(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if courtFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCourtsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Courts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCourtsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Courts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCourtsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	courtOne := &Court{}
	courtTwo := &Court{}
	if err = randomize.Struct(seed, courtOne, courtDBTypes, false, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}
	if err = randomize.Struct(seed, courtTwo, courtDBTypes, false, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = courtOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = courtTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Courts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCourtsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	courtOne := &Court{}
	courtTwo := &Court{}
	if err = randomize.Struct(seed, courtOne, courtDBTypes, false, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}
	if err = randomize.Struct(seed, courtTwo, courtDBTypes, false, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = courtOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = courtTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testCourtsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCourtsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(courtColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCourtsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
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

func testCourtsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CourtSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCourtsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Courts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	courtDBTypes = map[string]string{`ID`: `integer`, `Court`: `character varying`}
	_            = bytes.MinRead
)

func testCourtsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(courtPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(courtAllColumns) == len(courtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, courtDBTypes, true, courtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCourtsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(courtAllColumns) == len(courtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Court{}
	if err = randomize.Struct(seed, o, courtDBTypes, true, courtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, courtDBTypes, true, courtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(courtAllColumns, courtPrimaryKeyColumns) {
		fields = courtAllColumns
	} else {
		fields = strmangle.SetComplement(
			courtAllColumns,
			courtPrimaryKeyColumns,
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

	slice := CourtSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCourtsUpsert(t *testing.T) {
	t.Parallel()

	if len(courtAllColumns) == len(courtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Court{}
	if err = randomize.Struct(seed, &o, courtDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Court: %s", err)
	}

	count, err := Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, courtDBTypes, false, courtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Court struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Court: %s", err)
	}

	count, err = Courts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
