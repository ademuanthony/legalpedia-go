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

func testPackages(t *testing.T) {
	t.Parallel()

	query := Packages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPackagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
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

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Packages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PackageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PackageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Package exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PackageExists to return true, but got false.")
	}
}

func testPackagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	packageFound, err := FindPackage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if packageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPackagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Packages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPackagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Packages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPackagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	packageOne := &Package{}
	packageTwo := &Package{}
	if err = randomize.Struct(seed, packageOne, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}
	if err = randomize.Struct(seed, packageTwo, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = packageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = packageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Packages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPackagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	packageOne := &Package{}
	packageTwo := &Package{}
	if err = randomize.Struct(seed, packageOne, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}
	if err = randomize.Struct(seed, packageTwo, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = packageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = packageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPackagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPackagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(packageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPackagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
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

func testPackagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PackageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPackagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Packages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	packageDBTypes = map[string]string{`ID`: `integer`, `Name`: `character`, `Description`: `text`, `Price`: `real`, `LastModificationTime`: `timestamp without time zone`, `LastModifierUserId`: `bigint`, `CreationTime`: `timestamp without time zone`, `CreatorUserId`: `bigint`, `Days`: `integer`, `Features`: `text`, `Permalink`: `character`, `Platforms`: `text`}
	_              = bytes.MinRead
)

func testPackagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(packagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(packageAllColumns) == len(packagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, packageDBTypes, true, packagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPackagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(packageAllColumns) == len(packagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Package{}
	if err = randomize.Struct(seed, o, packageDBTypes, true, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, packageDBTypes, true, packagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(packageAllColumns, packagePrimaryKeyColumns) {
		fields = packageAllColumns
	} else {
		fields = strmangle.SetComplement(
			packageAllColumns,
			packagePrimaryKeyColumns,
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

	slice := PackageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPackagesUpsert(t *testing.T) {
	t.Parallel()

	if len(packageAllColumns) == len(packagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Package{}
	if err = randomize.Struct(seed, &o, packageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Package: %s", err)
	}

	count, err := Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, packageDBTypes, false, packagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Package: %s", err)
	}

	count, err = Packages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
