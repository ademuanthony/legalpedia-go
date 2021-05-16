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

func testMaxims(t *testing.T) {
	t.Parallel()

	query := Maxims()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMaximsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
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

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaximsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Maxims().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaximsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MaximSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaximsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MaximExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Maxim exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MaximExists to return true, but got false.")
	}
}

func testMaximsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	maximFound, err := FindMaxim(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if maximFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMaximsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Maxims().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMaximsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Maxims().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMaximsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	maximOne := &Maxim{}
	maximTwo := &Maxim{}
	if err = randomize.Struct(seed, maximOne, maximDBTypes, false, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}
	if err = randomize.Struct(seed, maximTwo, maximDBTypes, false, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = maximOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = maximTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Maxims().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMaximsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	maximOne := &Maxim{}
	maximTwo := &Maxim{}
	if err = randomize.Struct(seed, maximOne, maximDBTypes, false, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}
	if err = randomize.Struct(seed, maximTwo, maximDBTypes, false, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = maximOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = maximTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testMaximsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMaximsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(maximColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMaximsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
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

func testMaximsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MaximSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMaximsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Maxims().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	maximDBTypes = map[string]string{`ID`: `integer`, `UUID`: `text`, `Title`: `text`, `Content`: `text`, `VersionNo`: `integer`}
	_            = bytes.MinRead
)

func testMaximsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(maximPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(maximAllColumns) == len(maximPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, maximDBTypes, true, maximPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMaximsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(maximAllColumns) == len(maximPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Maxim{}
	if err = randomize.Struct(seed, o, maximDBTypes, true, maximColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, maximDBTypes, true, maximPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(maximAllColumns, maximPrimaryKeyColumns) {
		fields = maximAllColumns
	} else {
		fields = strmangle.SetComplement(
			maximAllColumns,
			maximPrimaryKeyColumns,
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

	slice := MaximSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMaximsUpsert(t *testing.T) {
	t.Parallel()

	if len(maximAllColumns) == len(maximPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Maxim{}
	if err = randomize.Struct(seed, &o, maximDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Maxim: %s", err)
	}

	count, err := Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, maximDBTypes, false, maximPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maxim struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Maxim: %s", err)
	}

	count, err = Maxims().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
