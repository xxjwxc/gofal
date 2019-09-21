# gofal

[English](README.md)

- golang分数运算相关函数
- 支持精准运算
- 支持加减乘除
- 支持链式表达式
- 支持结果输出(flat64)

### 调用示例


1、 加法 

```go
	tmp := fractional.Model(7, 12)
	tmp1 := fractional.Model(1, 12)
	fmt.Println(tmp.Add(tmp1))
```

- out
```sh
2/3
```

2、 减法 

```go
	tmp = fractional.Model(1, 4)
	tmp1 = fractional.Model(1, 3)
	fmt.Println(tmp.Sub(tmp1))
```

- out
```sh
-1/12
```

3、乘法 

```go
	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Mul(tmp1))
```

- out
```sh
1/2
```

4、 除法

```go
	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Div(*tmp1))
```

- out
```sh
9/8
```

5、 结果输出(flat64) 

```go
    tmp = fractional.Model(1, 3)
	fmt.Println(tmp.Verdict())
```

- out
```sh
0.3333333333333333
```

6、链式表达式

```go
	tmp := fractional.Model(1, 3)
	tmp.Add(fractional.Model(1)).Mul(tmp)
	fmt.Println(tmp)
```

- out
```sh
16/9
```
- 所有函数支持链式表达式
