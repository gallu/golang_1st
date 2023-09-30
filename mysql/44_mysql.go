package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE go_tests (
    go_test_id SERIAL ,
    string  VARCHAR(128) NOT NULL COMMENT '文字列',
    b_string  VARBINARY(128) NOT NULL COMMENT 'バイナリ',
    num INT NOT NULL COMMENT '整数',
    created_at DATETIME NOT NULL COMMENT '日付',
    PRIMARY KEY(go_test_id)
)CHARACTER SET 'utf8mb4', COMMENT='golang用テストテーブル';
*/

func main() {
	// ??? charset(か、よりベターなのはcollationだとか)の指定はどうすんだろ? デフォがutf8mb4っぽいんだよね……
	db, err := sql.Open("mysql", "go_test_user:go_Test_Pass1@tcp(127.0.0.1:3306)/go_test")
	// db, err := sql.Open("mysql", "go_test_user:go_Test_Ps1@tcp(127.0.0.1:3306)/go_test")
	if err != nil {
		panic(err)
	}
	defer db.Close() // XXX これ、やらなくてもよい???

	err = db.Ping() // 接続エラーの時、ここで始めて「エラー」って分かるんだ。OKならnilが返るぽ
	fmt.Println(err)

	// 簡単なクエリを流してみる
	{
		rows, err := db.Query("show variables like '%char%';")
		if err != nil {
			panic(err)
		}
		defer rows.Close() // めっちゃ大事らしい
		// 読み込み元は作っておかんといかんぽい
		var vname, v string
		// これで「カラム名の一覧」が取れる?
		cols, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		fmt.Println(cols)

		//
		for rows.Next() {
			err := rows.Scan(&vname, &v)
			if err != nil {
				panic(err)
			}
			fmt.Printf("data: %s\t%s\n", vname, v)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
	}
	// 動的な取得の実験
	{
		fmt.Println("動的な取得の実験")
		rows, err := db.Query("show variables like '%char%';")
		if err != nil {
			panic(err)
		}
		defer rows.Close() // めっちゃ大事らしい
		// これで「カラム名の一覧」が取れる?
		cols, _ := rows.Columns()
		// types, _ := rows.ColumnTypes() // 型情報が入ってるらしい、んだけど上手く文字にできない…… ColumnTypeDatabaseTypeName が使えるといいんだが……
		fmt.Println(cols)

		// カラム数分の「ポインタ」を用意する
		data := make([]any, len(cols))
		for i, _ := range cols {
			//data[i] = new(sql.RawBytes) // sql.RawBytes ネットで見かけたけど使いにくいから一度オミットしておく……
			data[i] = new([]byte)
			//data[i] = new(any) // うまくいってない orz
		}
		for rows.Next() {
			err = rows.Scan(data...) // 3点ドットで「スライスを展開して引数に渡す」
			for i := range data {
				/*
					b := data[i].(*[]byte)
					fmt.Println(string(*b))
				*/
				// 上の2行を縮めた
				fmt.Println(string(*data[i].(*[]byte)))
			}
		}
	}

	// insertしてselectしてみやう
	{
		// Exec()でプレースホルダ
		// https://www.wakuwakubank.com/posts/869-go-database-sql/#index_id6
		res, err := db.Exec(`
			INSERT INTO go_tests(string, b_string, num, created_at)
			VALUES(?, ?, ?, ?)
		`, "ほげ", "hoge", 123, time.Now())
		if err != nil {
			panic(err)
		}
		fmt.Println(res) // 謎い値。なのでこれ直接出力するもんじゃねぇな
		// あ、LastIDあるのか
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
	{
		/*
			// https://mattn.kaoriya.net/software/lang/go/20161106232834.htm にあったんだけど、駄目かねぇ?
			// sql.Param が見当たらない……
			row, err := db.Exec(`
				INSERT INTO go_tests(string, b_string, num, created_at)
				VALUES(:string, :b_string, :num, :created_at)
			`, sql.Param(":string", "ふげ"), sql.Param(":b_string", "binbin"), sql.Param(":num", 777), sql.Param(":created_at", time.Now()))
			if err != nil {
				panic(err)
			}
			id, err := res.LastInsertId()
			if err != nil {
				panic(err)
			}
			fmt.Println(id)
		*/
	}
	// とりあえず「数値のある」テーブルを読み込んでみよう
	// いけるぽい。ってことは全体的に「文字」で入ってくるねぇ多分
	{
		fmt.Println("動的な取得の実験")
		rows, err := db.Query("SELECT * FROM go_tests;")
		if err != nil {
			panic(err)
		}
		defer rows.Close() // めっちゃ大事らしい
		// これで「カラム名の一覧」が取れる?
		cols, _ := rows.Columns()
		// types, _ := rows.ColumnTypes() // 型情報が入ってるらしい、んだけど上手く文字にできない…… ColumnTypeDatabaseTypeName が使えるといいんだが……
		fmt.Println(cols)

		// カラム数分の「ポインタ」を用意する
		data := make([]any, len(cols))
		for i, _ := range cols {
			//data[i] = new(sql.RawBytes) // sql.RawBytes ネットで見かけたけど使いにくいから一度オミットしておく……
			data[i] = new([]byte)
			//data[i] = new(any) // うまくいってない orz
		}
		for rows.Next() {
			err = rows.Scan(data...) // 3点ドットで「スライスを展開して引数に渡す」
			for i := range data {
				/*
					b := data[i].(*[]byte)
					fmt.Println(string(*b))
				*/
				// 上の2行を縮めた
				fmt.Println(string(*data[i].(*[]byte)))
			}
		}

	}

}
