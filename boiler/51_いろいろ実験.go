package main

import (
	"context"
	"database/sql"
	"fmt"
	"testSQLBoiler/models"

	// Import this so we don't have to use qm.Limit etc.
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

func main() {
	db, err := sql.Open("mysql", "go_test_user:go_Test_Pass1@tcp(127.0.0.1:3306)/go_test?parseTime=true&loc=Asia%2FTokyo") // ?parseTime=true がないと、日付型でエラーが出るっぽい & locがないとUTCとして取得するぽ(これは設定にもよるんだろうなぁ)
	if err != nil {
		panic(err)
	}
	defer db.Close() // XXX これ、やらなくてもよい???
	//boil.DebugMode = true // これあるとSQL文とか把握できる
	boil.SetDB(db)

	// 事前に: 一気にお片付け
	{
		ctx := context.Background()
		models.GoTest2NDS().DeleteAll(ctx, db)
	}

	// insert
	m := &models.GoTest2ND{
		UserID: 1,
		Num:    100,
	}
	ctx := context.Background()
	if err := m.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}
	fmt.Println("insert", m)

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
		m.Num += 100
		m.Update(ctx, db, boil.Infer())
		fmt.Println("update", m)
	}

	// upsertの確認
	{
		sql := `
		INSERT INTO go_test_2nd(user_id, num, created_at, updated_at)
		       	VALUES(?, ?, NOW(), NOW())
			ON DUPLICATE KEY
				UPDATE num = num + ?, updated_at=NOW();
		`
		r, err := queries.Raw(sql, 2, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ := r.RowsAffected()                 // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(2),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}

		// 二度目
		r, err = queries.Raw(sql, 2, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ = r.RowsAffected()                 // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(2),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}

		// 三度目
		r, err = queries.Raw(sql, 2, 100, 150).Exec(db) // 第一戻り値 sql.Result って型らしい
		rowAffect, _ = r.RowsAffected()                 // これで「影響行」が取れるんだ
		fmt.Println("upsert", rowAffect, err)
		{
			ctx := context.Background()
			// XXX 最後のsが大文字なのか小文字なのか、はたぶんなんかルールがありそう(別テーブルだと小文字だった): あと、これfunc (調べるときの手がかり)
			m, err := models.GoTest2NDS(
				models.GoTest2NDWhere.UserID.EQ(2),
			).One(ctx, db)
			if err != nil {
				panic(err) // データあるはずだから、 err == sql.ErrNoRows とか見ない
			}
			fmt.Println(m)
		}
	}

}
