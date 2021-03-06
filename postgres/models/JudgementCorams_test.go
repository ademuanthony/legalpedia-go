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

func testJudgementCorams(t *testing.T) {
	t.Parallel()

	query := JudgementCorams()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJudgementCoramsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
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

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJudgementCoramsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := JudgementCorams().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJudgementCoramsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JudgementCoramSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJudgementCoramsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JudgementCoramExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if JudgementCoram exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JudgementCoramExists to return true, but got false.")
	}
}

func testJudgementCoramsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	judgementCoramFound, err := FindJudgementCoram(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if judgementCoramFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJudgementCoramsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = JudgementCorams().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJudgementCoramsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := JudgementCorams().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJudgementCoramsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	judgementCoramOne := &JudgementCoram{}
	judgementCoramTwo := &JudgementCoram{}
	if err = randomize.Struct(seed, judgementCoramOne, judgementCoramDBTypes, false, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}
	if err = randomize.Struct(seed, judgementCoramTwo, judgementCoramDBTypes, false, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = judgementCoramOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = judgementCoramTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := JudgementCorams().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJudgementCoramsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	judgementCoramOne := &JudgementCoram{}
	judgementCoramTwo := &JudgementCoram{}
	if err = randomize.Struct(seed, judgementCoramOne, judgementCoramDBTypes, false, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}
	if err = randomize.Struct(seed, judgementCoramTwo, judgementCoramDBTypes, false, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = judgementCoramOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = judgementCoramTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testJudgementCoramsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJudgementCoramsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(judgementCoramColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJudgementCoramsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
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

func testJudgementCoramsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JudgementCoramSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJudgementCoramsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := JudgementCorams().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	judgementCoramDBTypes = map[string]string{`ID`: `integer`, `CoramID`: `integer`, `SuitNo`: `character varying`}
	_                     = bytes.MinRead
)

func testJudgementCoramsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(judgementCoramPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(judgementCoramAllColumns) == len(judgementCoramPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJudgementCoramsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(judgementCoramAllColumns) == len(judgementCoramPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &JudgementCoram{}
	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, judgementCoramDBTypes, true, judgementCoramPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(judgementCoramAllColumns, judgementCoramPrimaryKeyColumns) {
		fields = judgementCoramAllColumns
	} else {
		fields = strmangle.SetComplement(
			judgementCoramAllColumns,
			judgementCoramPrimaryKeyColumns,
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

	slice := JudgementCoramSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJudgementCoramsUpsert(t *testing.T) {
	t.Parallel()

	if len(judgementCoramAllColumns) == len(judgementCoramPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := JudgementCoram{}
	if err = randomize.Struct(seed, &o, judgementCoramDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert JudgementCoram: %s", err)
	}

	count, err := JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, judgementCoramDBTypes, false, judgementCoramPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JudgementCoram struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert JudgementCoram: %s", err)
	}

	count, err = JudgementCorams().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
