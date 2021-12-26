package main

import (
	"fmt"
	"os"
)

func main() {

	salary := 130000

	if salary < 150000 {

		errorSalary := myError{status: 400}
		fmt.Println(myErrorTest(errorSalary.status))
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

type myError struct {
	status int
	msg    string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myError{
			status: status,
			msg:    "error: el salario ingresado no alcanza el minimo imponible",
		}
	}
	return 200, nil
}
