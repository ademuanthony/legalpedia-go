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

func testPartyATypes(t *testing.T) {
	t.Parallel()

	query := PartyATypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPartyATypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
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

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPartyATypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PartyATypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPartyATypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PartyATypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPartyATypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PartyATypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PartyAType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PartyATypeExists to return true, but got false.")
	}
}

func testPartyATypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	partyATypeFound, err := FindPartyAType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if partyATypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPartyATypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PartyATypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPartyATypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PartyATypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPartyATypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	partyATypeOne := &PartyAType{}
	partyATypeTwo := &PartyAType{}
	if err = randomize.Struct(seed, partyATypeOne, partyATypeDBTypes, false, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}
	if err = randomize.Struct(seed, partyATypeTwo, partyATypeDBTypes, false, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = partyATypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = partyATypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PartyATypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPartyATypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	partyATypeOne := &PartyAType{}
	partyATypeTwo := &PartyAType{}
	if err = randomize.Struct(seed, partyATypeOne, partyATypeDBTypes, false, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}
	if err = randomize.Struct(seed, partyATypeTwo, partyATypeDBTypes, false, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = partyATypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = partyATypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPartyATypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPartyATypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(partyATypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPartyATypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
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

func testPartyATypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PartyATypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPartyATypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PartyATypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	partyATypeDBTypes = map[string]string{`ID`: `integer`, `PartyAType`: `character varying`}
	_                 = bytes.MinRead
)

func testPartyATypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(partyATypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(partyATypeAllColumns) == len(partyATypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPartyATypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(partyATypeAllColumns) == len(partyATypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PartyAType{}
	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, partyATypeDBTypes, true, partyATypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(partyATypeAllColumns, partyATypePrimaryKeyColumns) {
		fields = partyATypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			partyATypeAllColumns,
			partyATypePrimaryKeyColumns,
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

	slice := PartyATypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPartyATypesUpsert(t *testing.T) {
	t.Parallel()

	if len(partyATypeAllColumns) == len(partyATypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PartyAType{}
	if err = randomize.Struct(seed, &o, partyATypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PartyAType: %s", err)
	}

	count, err := PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, partyATypeDBTypes, false, partyATypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PartyAType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PartyAType: %s", err)
	}

	count, err = PartyATypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
