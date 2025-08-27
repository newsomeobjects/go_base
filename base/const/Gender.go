package main

type Gender byte

const (
	Male Gender = iota
	Female
)

//如果 iota 定义在 const 定义组中的第 n 行，那么 iota 的值为 n - 1。
//所以一定要注意 iota 出现在定义组中的第几行，而不是当前代码中它第几次出现。
