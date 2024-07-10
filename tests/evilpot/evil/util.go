package evil

import (
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

var SepFunc = func(b byte) bool {
	if 48 <= b && b <= 57 {
		return false
	}
	if b >= 128 {
		return false
	}
	if b >= 65 && b <= 90 {
		return false
	}
	if b >= 97 && b <= 122 {
		return false
	}
	return true
}

func Split(data []byte, sep func(b byte) bool, handler func([]byte) bool) {
	if sep == nil {
		sep = SepFunc
	}
	visible := 0
	invisible := 0
	for i := 0; i < len(data); i++ {
		if !sep(data[i]) {
			continue
		}
		invisible = i
		if invisible == visible {
			visible++
			continue
		}
		b := data[visible:invisible]
		if !handler(b) {
			return
		}
		visible = invisible + 1
	}
	if visible < len(data) {
		handler(data[visible:])
	}
	return
}

func calculate(expression string) (interface{}, error) {
	// 去除多余空格
	expression = strings.TrimSpace(expression)
	expression = strings.ReplaceAll(expression, " ", "")

	// 定义运算符优先级
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	// 定义两个栈，一个存操作数，一个存操作符
	var numStack []interface{}
	var opStack []rune

	// 辅助函数，用于执行整数计算
	applyOpInt := func(op rune, b, a *big.Int) *big.Int {
		result := new(big.Int)
		switch op {
		case '+':
			return result.Add(a, b)
		case '-':
			return result.Sub(a, b)
		case '*':
			return result.Mul(a, b)
		case '/':
			return result.Div(a, b)
		default:
			return big.NewInt(0)
		}
	}

	// 辅助函数，用于执行浮点数计算
	applyOpFloat := func(op rune, b, a *big.Float) *big.Float {
		result := new(big.Float)
		switch op {
		case '+':
			return result.Add(a, b)
		case '-':
			return result.Sub(a, b)
		case '*':
			return result.Mul(a, b)
		case '/':
			return result.Quo(a, b)
		default:
			return big.NewFloat(0)
		}
	}

	// 判断字符串是否为整数
	isInteger := func(s string) bool {
		_, err := strconv.ParseInt(s, 10, 64)
		return err == nil
	}

	// 将操作数转换为浮点数
	toBigFloat := func(num interface{}) *big.Float {
		switch v := num.(type) {
		case *big.Int:
			return new(big.Float).SetInt(v)
		case *big.Float:
			return v
		default:
			return big.NewFloat(0)
		}
	}

	// 解析表达式并进行计算
	i := 0
	for i < len(expression) {
		char := rune(expression[i])

		if unicode.IsDigit(char) || char == '.' {
			// 处理数字和小数点
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || rune(expression[j]) == '.') {
				j++
			}
			numStr := expression[i:j]
			if isInteger(numStr) {
				num := new(big.Int)
				num, ok := num.SetString(numStr, 10)
				if !ok {
					return big.NewFloat(0), nil
				}
				numStack = append(numStack, num)
			} else {
				num, ok := new(big.Float).SetString(numStr)
				if !ok {
					return big.NewFloat(0), nil
				}
				numStack = append(numStack, num)
			}
			i = j
		} else if char == '(' {
			// 处理左括号
			opStack = append(opStack, char)
			i++
		} else if char == ')' {
			// 处理右括号
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				b := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				a := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]

				if aInt, okA := a.(*big.Int); okA {
					if bInt, okB := b.(*big.Int); okB {
						result := applyOpInt(op, bInt, aInt)
						numStack = append(numStack, result)
					} else {
						result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
						numStack = append(numStack, result)
					}
				} else {
					result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
					numStack = append(numStack, result)
				}
			}
			// 移除左括号
			if len(opStack) > 0 {
				opStack = opStack[:len(opStack)-1]
			}
			i++
		} else if strings.ContainsRune("+-*/", char) {
			// 处理运算符
			for len(opStack) > 0 && precedence[char] <= precedence[opStack[len(opStack)-1]] {
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				b := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				a := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]

				if aInt, okA := a.(*big.Int); okA {
					if bInt, okB := b.(*big.Int); okB {
						result := applyOpInt(op, bInt, aInt)
						numStack = append(numStack, result)
					} else {
						result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
						numStack = append(numStack, result)
					}
				} else {
					result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
					numStack = append(numStack, result)
				}
			}
			opStack = append(opStack, char)
			i++
		} else {
			return big.NewFloat(0), nil
		}
	}

	// 处理剩余的运算符
	for len(opStack) > 0 {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		b := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		a := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]

		if aInt, okA := a.(*big.Int); okA {
			if bInt, okB := b.(*big.Int); okB {
				result := applyOpInt(op, bInt, aInt)
				numStack = append(numStack, result)
			} else {
				result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
				numStack = append(numStack, result)
			}
		} else {
			result := applyOpFloat(op, toBigFloat(b), toBigFloat(a))
			numStack = append(numStack, result)
		}
	}

	if len(numStack) != 1 {
		return big.NewFloat(0), nil
	}

	return numStack[0], nil
}
