package fractional

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tmp := Model(1, 2)
	tmp2 := Model(1, 2)
	result := Model(1)            // or Model(1,1)
	actualResult := tmp.Add(tmp2) // (1/2) + (1/2)
	if *actualResult != result {
		t.Errorf("Expected %s, got %s", result, *actualResult)
	}
}

func TestSub(t *testing.T) {
	tmp := Model(1, 3)
	tmp1 := Model(1, 2)
	result := Model(-1, 6)
	actualResult := tmp.Sub(tmp1) // (1/3) - (1/2)
	if *actualResult != result {
		t.Errorf("Expected %s, got %s", result, *actualResult)
	}
}

func TestMul(t *testing.T) {
	tmp := Model(1, 3)
	tmp1 := Model(1, 2)
	result := Model(1, 6)
	actualResult := tmp.Mul(tmp1) // (1/3) * (1/2)
	if *actualResult != result {
		t.Errorf("Expected %s, got %s", result, *actualResult)
	}
}

func TestDiv(t *testing.T) {
	tmp := Model(1, 3)
	tmp1 := Model(1, 2)
	result := Model(2, 3)
	actualResult := tmp.Div(tmp1) // (1/3) / (1/2)
	if *actualResult != result {
		t.Errorf("Expected %s, got %s", result, *actualResult)
	}
}

// func TestError(t *testing.T) {
// 	tmp := Model(3, 0)
// 	fmt.Println("tmp:", tmp.Verdict())
// }

func TestChainRule(t *testing.T) {
	tmp := Model(1)
	tmp.Add(Model(1, 3)).Mul(Model(1, 2)) // (1 + (1/3))*(1/2)
	result := Model(2, 3)
	if tmp != result {
		t.Errorf("Expected %s, got %s", result, tmp)
	}
}
