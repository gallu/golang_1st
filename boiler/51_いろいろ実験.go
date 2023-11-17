package main

import (
	"context"
	"database/sql"
	"fmt"
	"testSQLBoiler/models"

	// Import this so we don't have to use qm.Limit etc.
	_ "github.com/go-sql-driver/mysql"
	// "github.com/guregu/null" // null許可のカラムがある時は必要
	"github.com/volatiletech/null/v8" // nullは「こっち」ぽい
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	db, err := sql.Open("mysql", "go_test_user:go_Test_Pass1@tcp(127.0.0.1:3306)/go_test?parseTime=true&loc=Asia%2FTokyo") // ?parseTime=true がないと、日付型でエラーが出るっぽい & locがないとUTCとして取得するぽ(これは設定にもよるんだろうなぁ)
	if err != nil {
		panic(err)
	}
	defer db.Close()      // XXX これ、やらなくてもよい???
	boil.DebugMode = true // これあるとSQL文とか把握できる
	boil.SetDB(db)
	ctx := context.Background()

	// 事前に: 一気にお片付け
	{
		ctx := context.Background()
		models.GoTest2NDS().DeleteAll(ctx, db)
	}

	// insert
	/*
		i := null.NewInt(100, true)
		fmt.Printf("%T, %+v\n", i, i)
	*/
	{
		m := &models.GoTest2ND{
			UserID: 1,
			Num:    null.Int64From(100), // 単純に値を作りたいならこっちが楽(NewInt64()もあるけど)
		}
		ctx := context.Background()
		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			panic(err)
		}
		fmt.Println("insert", m)
	}
	{
		m := &models.GoTest2ND{
			UserID: 10,
			Num:    null.NewInt64(0, false), // 第二引数をfalseにするとnullが入る
		}
		ctx := context.Background()
		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			panic(err)
		}
		fmt.Println("insert", m)
	}

	// (selectしてからの)update
	{
		ctx := context.Background()
		// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
		m, err := models.GoTest2NDS(
			models.GoTest2NDWhere.UserID.EQ(1),
		).One(ctx, db)
		if err != nil {
			panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
		}
		/*
			num, _ := m.Num.Value()
			fmt.Printf("*************** %T, %#v \n", num, num)
			num2 := num.(int64) + 100 // numはValue()で取り出しても driver.Value なので、型を変換してから
		*/
		// num := m.Num.Int64 + 100
		m.Num = null.Int64From(m.Num.Int64 + 123) // これでいけるのか…… https://github.com/volatiletech/null/blob/6ee82df11e0c5503159b918642256e2d8f1beb39/int64.go#L14C2-L14C8
		m.Update(ctx, db, boil.Infer())
		fmt.Println("update", m)
	}

	// upsertの確認
	{
		user_id := uint64(2)
		sql := `
			INSERT INTO go_test_2nd(user_id, num, created_at, updated_at)
			       	VALUES(?, ?, NOW(), NOW())
				ON DUPLICATE KEY
					UPDATE num = num + ?, updated_at=NOW();
			`
		r, err := queries.Raw(sql, user_id, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ := r.RowsAffected()                       // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(user_id),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}

		// 二度目
		r, err = queries.Raw(sql, user_id, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ = r.RowsAffected()                       // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(user_id),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}

		// 三度目
		r, err = queries.Raw(sql, user_id, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ = r.RowsAffected()                       // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(user_id),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}
	}

	// 生SQLを使う(スライス)
	{
		// o := &GoTest{}
		var o []models.GoTest
		sql := `SELECT * FROM go_tests;`
		err := queries.Raw(sql).Bind(ctx, db, &o)
		if err != nil {
			panic(err)
		}
		for _, v := range o {
			fmt.Println(v)
		}
	}
	// 生SQLを使う(1レコード)
	{
		o := models.GoTest{}
		id := 1
		sql := `SELECT * FROM go_tests WHERE go_test_id = ? FOR UPDATE;`
		err := queries.Raw(sql, id).Bind(ctx, db, &o)
		if err != nil {
			panic(err)
		}
		fmt.Println(o)
	}
	// 生SQLを使う(1レコード) レコードがないとき
	{
		o := models.GoTest{}
		id := 1234
		err := queries.Raw(`SELECT * FROM go_tests WHERE go_test_id = ? FOR UPDATE;`, id).Bind(ctx, db, &o)
		if err == sql.ErrNoRows {
			fmt.Println("not found")
		} else if err != nil {
			panic(err)
		}
		fmt.Println(o)
	}

	// ネストした条件(ベタ)
	{
		ctx := context.Background()
		// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
		m, err := models.GoTest2NDS(
			qm.Where("user_id >= ? AND (num IS NULL OR num > ?)", 1, 50),
		).All(ctx, db)
		if err != nil {
			panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
		}
		for _, v := range m {
			fmt.Println(v)
		}
	}
	// ネストした条件(ちょい綺麗に)
	{
		ctx := context.Background()
		// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
		m, err := models.GoTest2NDS(
			qm.Where("user_id = ?", 1), qm.Or("user_id = ?", 2),
			qm.Expr(qm.And("num > ?", 50), qm.Or("num IS NULL")),
		).All(ctx, db)
		if err != nil {
			panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
		}
		for _, v := range m {
			fmt.Println(v)
		}
	}
	// ネストした条件(綺麗に)
	{
		ctx := context.Background()
		// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
		m, err := models.GoTest2NDS(
			models.GoTest2NDWhere.UserID.EQ(1),
			qm.Or2(models.GoTest2NDWhere.UserID.EQ(2)),
			qm.Expr(
				models.GoTest2NDWhere.Num.GTE(null.Int64From(55)),
				models.GoTest2NDWhere.Num.IsNull(),
			),
		).All(ctx, db)
		if err != nil {
			panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
		}
		for _, v := range m {
			fmt.Println(v)
		}
	}
}
