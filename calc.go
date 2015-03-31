package main

import "fmt"
import "strconv"
import "errors"

func toNumber(eq []byte) (int, error) {
	return strconv.Atoi(string(eq))
}

func splitOnOp(eq []byte, op byte) ([]byte, []byte, error) {
	for i, c := range eq {
		if c == op {
			return eq[0:i], eq[i+1 : len(eq)], nil
		}
	}
	return nil, nil, errors.New("No " + string(op))
}

func calc(eq []byte) (int, error) {

	if num, err := toNumber(eq); err == nil {
		return num, nil
	}

	for _, op := range []byte("+*") {

		if l, r, err := splitOnOp(eq, op); err == nil {
			l, err := calc(l)
			r, err := calc(r)
			fmt.Println(l, string(op), r)
			if err != nil {
				return 0, err
			}
			switch op {
			case '+':
				return l + r, nil
			case '*':
				return l * r, nil
			}
		}
	}

	return 0, errors.New("Oops.")

}

func main() {

	eq := []byte("1+2*3+4*5+6")
	fmt.Println(calc(eq))

}
