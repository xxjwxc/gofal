package fractional

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tmp := Model(7, 12)
	tmp1 := Model(1, 12)
	result := tmp
	result.Add(tmp1)
	fmt.Println(fmt.Sprintf("(%v) + (%v) = (%v)", tmp, tmp1, result))
	fmt.Println("result:", result.Verdict())
}

func TestSub(t *testing.T) {
	tmp := Model(1, 4)
	tmp1 := Model(1, 3)
	result := tmp
	result.Sub(tmp1)
	fmt.Println(fmt.Sprintf("(%v) - (%v) = (%v)", tmp, tmp1, result))
	fmt.Println("result:", result.Verdict())
}
func TestMul(t *testing.T) {
	tmp := Model(3, 4)
	tmp1 := Model(2, 3)
	result := tmp
	result.Mul(tmp1)
	fmt.Println(fmt.Sprintf("(%v) * (%v) = (%v)", tmp, tmp1, result))
	fmt.Println("result:", result.Verdict())
}

func TestDiv(t *testing.T) {
	tmp := Model(3, 4)
	tmp1 := Model(2, 3)
	result := tmp
	result.Div(tmp1)
	fmt.Println(fmt.Sprintf("(%v) * (%v) = (%v)", tmp, tmp1, result))
	fmt.Println("result:", result.Verdict())
}

func TestError(t *testing.T) {
	tmp := Model(3, 0)
	fmt.Println("tmp:", tmp.Verdict())
}

func TestChainRule(t *testing.T) {
	tmp := Model(1, 2)
	tmp.Add(Model(1, 3)).Mul(Model(1, 4))
	fmt.Println(tmp)
}
