package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func f1() error {
	err := f2();
	if err != nil {
		return err
	}
	return nil
}

func f2() error {
	err := f3();
	if err != nil {
		return err
	}
	return nil
}

func f3() error {
	return errors.New("my error")
}


func main() {
	fmt.Println("hello, world")

	err := f1();
	if err != nil {
		// 一番外側で一度WithStackすれば、そこまでのerrのスタックトレースが取れる？
		err = errors.WithStack(err)
		fmt.Printf("%+v\n", err)
	}
}
