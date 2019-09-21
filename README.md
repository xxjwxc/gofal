# gofal

[中文文档](README_zh_cn.md)

- Fractional Operational Correlation API base on golang
- Supporting Precision Operations
- Supporting addition, subtraction, multiplication and division
- Support chain expression
- Support result output(flat64)

### Example


1、 Add 

```go
	tmp := fractional.Model(7, 12)
	tmp1 := fractional.Model(1, 12)
	fmt.Println(tmp.Add(tmp1))
```

- out
```sh
2/3
```

2、 Subtraction 

```go
	tmp = fractional.Model(1, 4)
	tmp1 = fractional.Model(1, 3)
	fmt.Println(tmp.Sub(tmp1))
```

- out
```sh
-1/12
```

3、Multiplication

```go
	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Mul(tmp1))
```

- out
```sh
1/2
```

4、 Division

```go
	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Div(tmp1))
```

- out
```sh
9/8
```

5、 out put (flat64) 

```go
    tmp = fractional.Model(1, 3)
	fmt.Println(tmp.Verdict())
```

- out
```sh
0.3333333333333333
```

6、Chain expression

```go
	tmp := fractional.Model(1, 3)
	tmp.Add(fractional.Model(1)).Mul(tmp)
	fmt.Println(tmp)
```

- out
```sh
16/9
```
- All functions can support chain expressions
