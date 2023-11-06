package main

import (
	"context"
	"database/sql"
	"fmt"
	models "testSQLBoiler/models"

	// Import this so we don't have to use qm.Limit etc.
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go_test_user:go_Test_Pass1@tcp(127.0.0.1:3306)/go_test?parseTime=true&loc=Asia%2FTokyo") // ?parseTime=true がないと、日付型でエラーが出るっぽい & locがないとUTCとして取得するぽ(これは設定にもよるんだろうなぁ)
	if err != nil {
		panic(err)
	}
	defer db.Close()      // XXX これ、やらなくてもよい???
	boil.DebugMode = true // これあるとSQL文とか把握できる
	boil.SetDB(db)

	{
		ctx := context.Background()
		//gt, _ := models.GoTests().All(context.Background(), db)
		//gt, err := models.FindGoTest(ctx, db, 1)
		gt, err := models.FindGoTest(ctx, db, 999)
		//gt, _ := models.GoTests(qm.Where("go_test_id=?", 1)).One(ctx, db)
		// gt, err := models.GoTests().All(ctx, db)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("row not find")
			} else {
				fmt.Println("in panic")
				fmt.Println(err)
				panic(err)

			}
		} else {
			fmt.Println("GoTest:", gt)
			fmt.Printf("%T \n", gt)
			//fmt.Println(string(gt.BString))
		}
	}

	// トラン切る
	{
		ctx := context.Background()
		tx, err := boil.BeginTx(ctx, nil)
		if err != nil {
			panic(err)
		}

		gt, err := models.FindGoTest(ctx, tx, 1)
		//gt, _ := models.GoTests(qm.Where("go_test_id=?", 1)).One(ctx, db)
		// gt, err := models.GoTests().All(ctx, db)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("row not find")
			} else {
				fmt.Println("in panic")
				fmt.Println(err)
				panic(err)
			}
			tx.Rollback()
			return
		}
		fmt.Println("GoTest:", gt)
		fmt.Printf("%T \n", gt)
		//fmt.Println(string(gt.BString))

		tx.Commit()
	}
}
