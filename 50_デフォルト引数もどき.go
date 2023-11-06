package main

import "fmt"

/*
// これはダメぽ
func test(i int) {
	test(i, "dummy")
}
func test(i int, s string) {
	fmt.Println(i, s)
}
*/

/* 真っ当なFunctional Option Pattern */
// デフォルトで引数になりうる値の群れを構造体で宣言
type testOptions struct {
	str string
	boo bool
}

// 「デフォルトの引数の値の解決」をするための処理
type testOption func(*testOptions)

// XXX これ、外部から使うから、publicにしておかないとなのか
func ParamStr(s string) testOption {
	return func(args *testOptions) {
		args.str = s
	}
}
func ParamBool(b bool) testOption {
	return func(args *testOptions) {
		args.boo = b
	}
}

// 関数本体
func Test(i int, additionArgs ...testOption) int {
	// デフォルト値をここで設定
	args := &testOptions{
		str: "default string",
		boo: false,
	}

	// デフォルト引数で「明示的に指定された値」を取得
	for _, setter := range additionArgs {
		setter(args)
	}

	// 値を使う
	fmt.Println(i, args.str, args.boo)
	return 1
}

/* 「テスト用にDI差し込みたい」程度の、ちょっと雑な Functional Option Pattern */
// DIすることをイメージしている雑構造体
// XXX 本当はこれに関数が紐付いてたりするんだと思う
type InjDI1 struct{}
type InjDI2 struct{}

type InjDI1Test struct{}
type InjDI2Test struct{}

// interface切りたいので関数実装
func (ijd InjDI1) hoge()     {}
func (ijd InjDI2) foo()      {}
func (ijd InjDI1Test) hoge() {}
func (ijd InjDI2Test) foo()  {}

// interface設定しておく
type InjDI1Interface interface {
	hoge()
}
type InjDI2Interface interface {
	foo()
}

// XXX 中味は「ポインタ型」を欲しいのだとしても、interface自体はポイント型にはしない(これで普通にポインタ型の構造体も受け入れられる)
type testDIOptions struct {
	req1 InjDI1Interface
	req2 InjDI2Interface
}

// 「デフォルトの引数の値の解決」をするための処理
type TestDIOption func(*testDIOptions)

func InjectionForTest(req1 InjDI1Interface, rep2 InjDI2Interface) TestDIOption {
	return func(args *testDIOptions) {
		args.req1 = req1
		args.req2 = rep2
	}
}

// 関数本体
func Test2(i int, additionArgs ...TestDIOption) int {
	// いったん空で作成
	args := new(testDIOptions)

	// デフォルト引数で「明示的に指定された値」を取得
	for _, setter := range additionArgs {
		setter(args)
	}
	// 値が「空」ならデフォルトの値を埋め込む
	// XXX 多分こっちのほうが「無駄が少ない」ような気がしてる(テストの時だけだから気にしなくても、とも思うんだけど……)
	// XXX 一応多段にして「一部だけ差し替えたい」とかも出来るようにはする(けど、いるのか?)
	if args.req1 == nil {
		args.req1 = new(InjDI1)
	}
	if args.req2 == nil {
		args.req2 = new(InjDI2)
	}

	// 値を使う
	fmt.Println(i, args.req1, args.req2)
	fmt.Printf("%T, %T \n", args.req1, args.req2)
	return 1
}

//
func main() {
	// Functional Optionsで解決
	{
		r := Test(1)
		fmt.Println(r)
	}
	{
		// なるほど。「オプション引数」は、「作った関数をcallする」んだ!!
		r := Test(2, ParamStr("string"), ParamBool(true))
		fmt.Println(r)
	}

	// テストを想定したやつ
	// XXX デフォルト引数、「全部ある」か「全部ない」かだけなので、関数1つにまとめてある
	{
		Test2(10)
	}
	{
		// オプション部分を1関数にまとめた
		Test2(22, InjectionForTest(new(InjDI1Test), new(InjDI2Test)))
	}
	{
		// やらんだろうけど「一部だけ使いたい」時
		Test2(22, InjectionForTest(new(InjDI1Test), nil))
	}
}
