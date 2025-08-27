package main

import "fmt"

// PaymentMethod 接口定义了支付方法的基本操作
type PayMethod interface {
	Account
	Pay(amount int) bool //Pay(int) bool amount可省略
}

type Account interface {
	GetBalance() int
}

type Account1 interface {
	GetBalance() int
}
type TestValue struct {
	balance int
}

func (t TestValue) GetBalance() int {
	return t.balance
}

type TestValue1 struct {
	balance int
}

func (t *TestValue1) GetBalance() int {
	return t.balance
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}
func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买:")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	purchaseItem(debitCard, 300)

	//----------------------------
	creditCard1 := CreditCard{balance: 1, limit: 1000}
	fmt.Println(creditCard1.GetBalance())
	purchaseItem(&creditCard1, 800)
	p := PayMethod(&creditCard1)
	fmt.Println(p)
	p1 := Account(&creditCard1) // 调用pay失败，因为Account接口没有pay方法
	fmt.Println(p1.GetBalance())
	p2 := p1.(*CreditCard) //Account接口，转换回实现类 creditcard，然后再调用pay试试
	p2.Pay(1)              //可以调用

	//---------------对于一个结构体而言， 如果A接口的所有方法都被值接收者实现，那么结构体的值或者指针都可以赋值给A接口，go会自动解引用。
	//---------------如果结构体存在方法需要指针接收者，则必须使用指针-----------
	aa := TestValue{balance: 1}
	var a11 Account = aa
	var a12 Account = &aa
	fmt.Println(a11.GetBalance())
	fmt.Println(a12.GetBalance())

	// 由于结构体  TestValue 实现了Account的方法，采用的是值接收者，所以赋值给Account时，传递指针或者值都可以，Go会自动解引用

	aa1 := TestValue1{balance: 1}
	//var a13 Account = aa1 直接传值报错，因为值是没办法转成地址的
	var a14 Account = &aa1
	fmt.Println(a14.GetBalance())
	//var r1 Account = aa  //报错，因为Account接口的方法都是指针
	//var r2 Account1 = aa

}

/*
在 Go 中接口是一种抽象类型，是一组方法的集合，里面只声明方法，而没有任何数据成员。

而在 Go 中实现一个接口  也
不需要显式的声明 不需要显式的声明 不需要显式的声明，
只需要其他类型实现了接口中
所有 所有 所有所有所有所有所有所有所有所有所有所有所有 所有
的方法，就是实现了这个接口。



使用过程中需要注意以下几点：

1. Go 中接口声明的方法并
不要求 全部公开
不要求 全部公开
不要求 全部公开 不要求 全部公开。
2. 直接用接口类型作为变量时，
如果方法都是值接收者，可以用值或指针。
如果任何方法是指针接收者，则必须用指针。

- 接口可以嵌套。
- 接口中声明的方法，参数可以没有名称。
- 如果函数参数使用 interface{}可以接受任何类型的实参。
同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量

*/
