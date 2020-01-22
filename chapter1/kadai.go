package chapter1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/apbgo/go-study-group/chapter1/lib"
)

// Calc opには+,-,×,÷の4つが渡ってくることを想定してxとyについて計算して返却(正常時はerrorはnilでよい)
// 想定していないopが渡って来た時には0とerrorを返却
func Calc(op string, x, y int) (int, error) {

	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	// TODO Q1
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "×":
		return x * y, nil
	case "÷":
		if y == 0 {
			return 0, fmt.Errorf("/ by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("invalid op=%s", op)
	}
}

// StringEncode 引数strの長さが5以下の時キャメルケースにして返却、それ以外であればスネークケースにして返却
func StringEncode(str string) string {
	// ヒント：長さ(バイト長)はlen(str)で取得できる
	// chapter1/libのToCamelとToSnakeを使うこと

	// TODO Q2
	if len(str) <= 5 {
		return lib.ToCamel(str)
	}
	return lib.ToSnake(str)
}

// Sqrt 数値xが与えられたときにz²が最もxに近い数値zを返却
func Sqrt(x float64) float64 {

	// TODO Q3
	return math.Sqrt(x)
}

// Pyramid x段のピラミッドを文字列にして返却
// 期待する戻り値の例：x=5のとき "1\n12\n123\n1234\n12345"
// （x<=0の時は"error"を返却）
func Pyramid(x int) string {
	// ヒント：string <-> intにはstrconvを使う
	// int -> stringはstrconv.Ioa() https://golang.org/pkg/strconv/#Itoa

	// TODO Q4
	if x <= 0 {
		return "error"
	}

	var sb strings.Builder
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			sb.WriteString(strconv.Itoa(j))
		}
		if i < x {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// StringSum x,yをintにキャストし合計値を返却 (正常終了時、errorはnilでよい)
// キャスト時にエラーがあれば0とエラーを返却
func StringSum(x, y string) (int, error) {

	// ヒント：string <-> intにはstrconvを使う
	// string -> intはstrconv.Atoi() https://golang.org/pkg/strconv/#Atoi

	// TODO Q5
	xNumber, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}

	yNumber, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}

	return xNumber + yNumber, nil
}

// SumFromFileNumber ファイルを開いてそこに記載のある数字の和を返却
func SumFromFileNumber(filePath string) (int, error) {
	// ヒント：ファイルの扱い：os.Open()/os.Close()
	// bufio.Scannerなどで１行ずつ読み込むと良い

	// TODO Q6 オプション
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		sum += number
	}

	scErr := scanner.Err()
	if scErr != nil {
		return 0, scErr
	}

	return sum, nil
}
