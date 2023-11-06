// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testGoTest2NDS(t *testing.T) {
	t.Parallel()

	query := GoTest2NDS()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGoTest2NDSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
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

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoTest2NDSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GoTest2NDS().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoTest2NDSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoTest2NDSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoTest2NDSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GoTest2NDExists(ctx, tx, o.GoTest2NDID)
	if err != nil {
		t.Errorf("Unable to check if GoTest2ND exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GoTest2NDExists to return true, but got false.")
	}
}

func testGoTest2NDSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	goTest2NDFound, err := FindGoTest2ND(ctx, tx, o.GoTest2NDID)
	if err != nil {
		t.Error(err)
	}

	if goTest2NDFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGoTest2NDSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GoTest2NDS().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGoTest2NDSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GoTest2NDS().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGoTest2NDSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	goTest2NDOne := &GoTest2ND{}
	goTest2NDTwo := &GoTest2ND{}
	if err = randomize.Struct(seed, goTest2NDOne, goTest2NDDBTypes, false, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}
	if err = randomize.Struct(seed, goTest2NDTwo, goTest2NDDBTypes, false, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = goTest2NDOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = goTest2NDTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GoTest2NDS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGoTest2NDSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	goTest2NDOne := &GoTest2ND{}
	goTest2NDTwo := &GoTest2ND{}
	if err = randomize.Struct(seed, goTest2NDOne, goTest2NDDBTypes, false, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}
	if err = randomize.Struct(seed, goTest2NDTwo, goTest2NDDBTypes, false, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = goTest2NDOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = goTest2NDTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func goTest2NDBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func goTest2NDAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GoTest2ND) error {
	*o = GoTest2ND{}
	return nil
}

func testGoTest2NDSHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &GoTest2ND{}
	o := &GoTest2ND{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GoTest2ND object: %s", err)
	}

	AddGoTest2NDHook(boil.BeforeInsertHook, goTest2NDBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	goTest2NDBeforeInsertHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.AfterInsertHook, goTest2NDAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	goTest2NDAfterInsertHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.AfterSelectHook, goTest2NDAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	goTest2NDAfterSelectHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.BeforeUpdateHook, goTest2NDBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	goTest2NDBeforeUpdateHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.AfterUpdateHook, goTest2NDAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	goTest2NDAfterUpdateHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.BeforeDeleteHook, goTest2NDBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	goTest2NDBeforeDeleteHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.AfterDeleteHook, goTest2NDAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	goTest2NDAfterDeleteHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.BeforeUpsertHook, goTest2NDBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	goTest2NDBeforeUpsertHooks = []GoTest2NDHook{}

	AddGoTest2NDHook(boil.AfterUpsertHook, goTest2NDAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	goTest2NDAfterUpsertHooks = []GoTest2NDHook{}
}

func testGoTest2NDSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoTest2NDSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(goTest2NDColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoTest2NDSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
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

func testGoTest2NDSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoTest2NDSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGoTest2NDSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GoTest2NDS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	goTest2NDDBTypes = map[string]string{`GoTest2NDID`: `bigint`, `UserID`: `bigint`, `Num`: `bigint`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                = bytes.MinRead
)

func testGoTest2NDSUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(goTest2NDPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(goTest2NDAllColumns) == len(goTest2NDPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGoTest2NDSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(goTest2NDAllColumns) == len(goTest2NDPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GoTest2ND{}
	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, goTest2NDDBTypes, true, goTest2NDPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(goTest2NDAllColumns, goTest2NDPrimaryKeyColumns) {
		fields = goTest2NDAllColumns
	} else {
		fields = strmangle.SetComplement(
			goTest2NDAllColumns,
			goTest2NDPrimaryKeyColumns,
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

	slice := GoTest2NDSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGoTest2NDSUpsert(t *testing.T) {
	t.Parallel()

	if len(goTest2NDAllColumns) == len(goTest2NDPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLGoTest2NDUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := GoTest2ND{}
	if err = randomize.Struct(seed, &o, goTest2NDDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GoTest2ND: %s", err)
	}

	count, err := GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, goTest2NDDBTypes, false, goTest2NDPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GoTest2ND struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GoTest2ND: %s", err)
	}

	count, err = GoTest2NDS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}