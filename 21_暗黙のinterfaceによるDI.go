package main

import (
	"errors"
	"fmt"
)

// ログ記録関数(だとおもいねぇ)
func LogOutput(message string) {
	fmt.Println(message)
}

// データの保存場所(だとおもいねぇ)
type userDataType map[string]string
type SimpleDataStore struct {
	userData userDataType
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// いわゆるコンストラクタ的なの。初期データ込み
// XXX コンストラクタだし外からデータ受け取りたいよねぇ、ってことで書籍からざっくり変更
func NewSimpleDataStore(data userDataType) SimpleDataStore {
	return SimpleDataStore{
		userData: data,
	}
}

/* ------------------------------- */

/* ビジネスロジックが必要としているinterfaceの定義 */
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
type Logger interface {
	Log(message string)
}

// 関数LogOutputがinterface Logger に適合するように、以下を追加
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

/* ------------------------------- */

/* ビジネスロジック的なナニカ */
type SimpleLogic struct {
	l  Logger    // interfaceだ
	ds DataStore // こっちもinterfaceだ
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello: (" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("謎いユーザ")
	}
	return name + "さんこんにちは! こんにちは!", nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("SayHello" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("謎いユーザ")
	}
	return name + "さん、おばんでやんす", nil
}

/* ------------------------------- */
// 呼び出すためのinterface
type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

/* ------------------------------- */
// これ、httpのrequest受け取ってresponse返す系だよねぇ?
/*
func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHallo内: ")
	userID := r.URL.Query().Get("uesr_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}
*/
func (c Controller) HandleDummy(userID string) {
	c.l.Log("SayHallo内: ")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(message)
}

/* ------------------------------- */
func main() {
	/* この辺でこねこね作ってる感じだねぇ */
	l := LoggerAdapter(LogOutput) // あ、ここで「注入」するんだ: 自動化できると便利、かもだけど、どうなんだろ???: wire ってのがあるぽ。後で出てくる?
	ds := NewSimpleDataStore(userDataType{
		"1": "gallu",
		"2": "furu",
		"3": "melmay",
	})
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	/* ここで「使う」んだ */
	/*
		http.HandleFunc("/hello", c.SayHello) // c.SayHello って、解決できないよねぇ???
		http.ListenAndServe(":8080", nil)
	*/
	// とりあえず「動かしてみたい」ので小細工
	c.HandleDummy("1")
	c.HandleDummy("2")
	c.HandleDummy("5")
}
