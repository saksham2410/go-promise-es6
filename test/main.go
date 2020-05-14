package main

import (
	"fmt"

	"github.com/saksham2410/01/promise"
)

func main() {
	//Creating Promise
	var newProm = promise.NewPromise(promise.TestingFunction)

	//Creating Promise Interface
	var p promise.PromInterface = newProm

	p.Then(
		func(resolve string) {
			fmt.Println("Then method is called")
			fmt.Println(resolve)
		},
		func(err error) {
			fmt.Println("Then method is called")
			fmt.Println("Some error is there")
			fmt.Println(err)
		})

	p.Catch(func(err error) {
		fmt.Println("Catch Method is called")
		fmt.Println("Some error is there")
		fmt.Println(err)
	})

	p.Finally(func(result int) {
		if result == 1 {
			fmt.Println("Promise is Fulfilled")
		} else if result == 2 {
			fmt.Println("Promise is Rejected")
		} else {
			fmt.Println("Error has Occured")
		}
	})

}
