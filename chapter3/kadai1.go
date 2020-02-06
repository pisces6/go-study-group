package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (a Kadai1) ID() int {
	return a.id
}

func (a Kadai1) Name() string {
	return a.name
}

func (a *Kadai1) SetID(id int) {
	a.id = id
}

func (a *Kadai1) SetName(name string) {
	a.name = name
}
