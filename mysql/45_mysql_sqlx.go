package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

// ざっくり構造体
// 「お外に持ち出せる」用に先頭大文字にしてみる
type GoTest struct {
	Go_test_id string
	String     string
	B_string   string
	Num        int
	Created_at string
}

func main() {
	db, err := sqlx.Open("mysql", "go_test_user:go_Test_Pass1@tcp(127.0.0.1:3306)/go_test")
	if err != nil {
		panic(err)
	}
	defer db.Close() // XXX これ、やらなくてもよい???

	err = db.Ping() // 接続エラーの時、ここで始めて「エラー」って分かるんだ。OKならnilが返るぽ
	fmt.Println(err)

	// 名前付きプレースホルダでinsert
	{
		// 別行にしなくてもいいんだけど、とりあえずわかりやすく
		// 順番は「あえて」バラバラに
		/*
			m := map[string]interface{}{
				"num":        777,
				"string":     "ふげ",
				"created_at": time.Now(),
				"b_string":   "fuge",
			}
		*/
		// 実際はこっちのほうがケース的におおかろ
		m := map[string]any{}
		m["num"] = 777
		m["created_at"] = time.Now()
		m["string"] = "ふげ"
		m["b_string"] = "fuge"
		//
		_, err := db.NamedExec(`
			INSERT INTO go_tests(string, b_string, num, created_at)
			VALUES(:string, :b_string, :num, :created_at)
			`, m)
		if err != nil {
			panic(err)
		}
	}
	// selectを構造体で
	{
		m := map[string]any{}
		m["num"] = 777
		rows, err := db.NamedQuery(`SELECT * FROM go_tests WHERE num=:num`, m)
		if err != nil {
			panic(err)
		}
		defer rows.Close() // deferで閉じるの忘れずに!!

		for rows.Next() {
			var datum GoTest
			// StructScanで「構造体にScanする」が出来る
			err := rows.StructScan(&datum)
			if err != nil {
				panic(err)
			}
			fmt.Println(datum)
		}
	}
	// selectをmapで
	{
		m := map[string]any{}
		m["num"] = 777
		rows, err := db.NamedQuery(`SELECT * FROM go_tests WHERE num=:num`, m)
		if err != nil {
			panic(err)
		}
		defer rows.Close() // deferで閉じるの忘れずに!!

		for rows.Next() {
			datum := make(map[string]interface{})
			// StructScanで「構造体にScanする」が出来る
			err := rows.MapScan(datum)
			if err != nil {
				panic(err)
			}
			// fmt.Println(datum) // 出力が多分 []byte になってる
			for k, v := range datum {
				switch v.(type) {
				case int, int64, float64, float32:
					fmt.Println(k, v)
				case []uint8:
					fmt.Println(k, string(v.([]uint8)))
				default:
					fmt.Printf("%s %T \n", k, v)
				}
			}
		}
	}

}
