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

func testUpdateMetas(t *testing.T) {
	t.Parallel()

	query := UpdateMetas()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUpdateMetasDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
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

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpdateMetasQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UpdateMetas().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpdateMetasSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UpdateMetasSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpdateMetasExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UpdateMetasExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UpdateMetas exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UpdateMetasExists to return true, but got false.")
	}
}

func testUpdateMetasFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	updateMetasFound, err := FindUpdateMetas(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if updateMetasFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUpdateMetasBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UpdateMetas().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUpdateMetasOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UpdateMetas().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUpdateMetasAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	updateMetasOne := &UpdateMetas{}
	updateMetasTwo := &UpdateMetas{}
	if err = randomize.Struct(seed, updateMetasOne, updateMetasDBTypes, false, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}
	if err = randomize.Struct(seed, updateMetasTwo, updateMetasDBTypes, false, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = updateMetasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = updateMetasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UpdateMetas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUpdateMetasCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	updateMetasOne := &UpdateMetas{}
	updateMetasTwo := &UpdateMetas{}
	if err = randomize.Struct(seed, updateMetasOne, updateMetasDBTypes, false, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}
	if err = randomize.Struct(seed, updateMetasTwo, updateMetasDBTypes, false, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = updateMetasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = updateMetasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testUpdateMetasInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpdateMetasInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(updateMetasColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpdateMetasReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
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

func testUpdateMetasReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UpdateMetasSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUpdateMetasSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UpdateMetas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	updateMetasDBTypes = map[string]string{`ID`: `integer`, `Date`: `timestamp without time zone`, `Size`: `integer`, `Description`: `character`, `DownloadLink`: `character`}
	_                  = bytes.MinRead
)

func testUpdateMetasUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(updateMetasPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(updateMetasAllColumns) == len(updateMetasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUpdateMetasSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(updateMetasAllColumns) == len(updateMetasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UpdateMetas{}
	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, updateMetasDBTypes, true, updateMetasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(updateMetasAllColumns, updateMetasPrimaryKeyColumns) {
		fields = updateMetasAllColumns
	} else {
		fields = strmangle.SetComplement(
			updateMetasAllColumns,
			updateMetasPrimaryKeyColumns,
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

	slice := UpdateMetasSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUpdateMetasUpsert(t *testing.T) {
	t.Parallel()

	if len(updateMetasAllColumns) == len(updateMetasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UpdateMetas{}
	if err = randomize.Struct(seed, &o, updateMetasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UpdateMetas: %s", err)
	}

	count, err := UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, updateMetasDBTypes, false, updateMetasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpdateMetas struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UpdateMetas: %s", err)
	}

	count, err = UpdateMetas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
