package main

import (
	"errors"
	"fmt"
)

var (
	err_age_too_low  = errors.New("age too low")
	err_age_too_high = errors.New("age too high")
)

func validate_age(age int) error {
	if age < 18 {
		return fmt.Errorf("age too low")
	} else if age > 60 {
		return fmt.Errorf("age too high")
	}
	return nil
}

func main() {
	err := validate_age(50)
	if err == err_age_too_low {
		fmt.Println("CAN NOT ENTER")
		return
	} else if err == err_age_too_high {
		fmt.Println("LOL")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("welcome")
	}
}
