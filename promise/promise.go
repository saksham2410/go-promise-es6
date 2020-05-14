package promise

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//Global Wait Group to handle go corountines

//Constants pending, fulfilled and rejected which represents state of promise
//The states of promise can be pending, fulfilled and rejected, that's why we defined three constants here.
//Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers.
//Because it can be used in expressions, it provides a generality beyond that of simple enumerations
const (
	pending = iota
	fulfilled
	rejected
)

//Golang has the ability to declare and create own data types by combining one or more types,
//including both built-in and user-defined types.
//The struct of Promise can have value,error and state. The state is int because it can be pending, fulfiled or rejected (0,1,2)
type Promise struct {
	value     string
	err       error
	state     int
	catchWG   sync.WaitGroup
	finallyWG sync.WaitGroup
	globalWG  sync.WaitGroup
}

func testForErrorReturnValue() {
	//Test the promises
	// errorFunc := func() (string, error) {
	// 	fmt.Printf("Executing Async Task  \n");
	// 	time.Sleep(5 * time.Second);
	// 	fmt.Printf("Task Done , will throw error \n");
	// 	return "error", errors.New("Custom Error");
	// }
	// p2 := newPromise(errorFunc)
	// onFulfilled := func(r string) {
	// 	fmt.Printf("[onFulfilled]  Got this from promise %d\n , now sleeping \n", r);
	// 	time.Sleep(5 * time.Second);
	// 	fmt.Printf("[onFulfilled] done \n");
	// }
	// onRejected := func(e error) {
	// 	fmt.Println("[onRejected] Handling error ", e);
	// 	time.Sleep(2 * time.Second);
	// 	fmt.Println("[onRejected] Error handling done ");
	// }
	// p2.then(onFulfilled func(string), onRejected func(error)).Catch(func(e error) {
	// 	fmt.Println("[onRejected-ThenWithErrorHandler] Error in catch block2 done", e);
	// }).Finally(func() {
	// 	fmt.Println("[onRejected-ThenWithErrorHandler] Finally called2 done");
	// })
}

//Go doesn't have prototypes but it has interfaces
//In Go language, the interface is a custom type that is used to specify a set of one or more method signatures
//and the interface is abstract, so you are not allowed to create an instance of the interface.
//But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type
//value that has the methods the interface requires. Or in other words, the interface is a collection of
//methods as well as it is a custom type.
//In this case, the PromiseInterface can have Then, Catch, Finally methods
type PromInterface interface {
	Then(func(string), func(error))
	Catch(func(error))
	Finally(func(int))
}

//A new promise pointer is created.
// A wait group of 1 is added so that the corountine can run and value of string and error can be assigned to p.value and p.err
func NewPromise(f func() (string, error)) *Promise {
	p := &Promise{}
	p.globalWG.Add(1)
	go func() {
		p.value, p.err = f()
		p.state = pending
		p.globalWG.Done()
	}()
	return p
}

//The Then method is applied on struct promise pointer and this takes OnFulfilled and OnRejected as arguments
// wgGlobal.wait() checks if the wgGlobal.Done() is called or not (decremented the pointer by 1).
//We need to create a new promise before we can use Then method on it
// If promise struct has err value of not null, then OnRejected handler (function) is invoked with value p.err
func (p *Promise) Then(OnFulfilled func(string), OnRejected func(error)) {
	func() {
		p.globalWG.Wait()
		if p.err != nil {
			OnRejected(p.err)
			p.state = rejected
			return
		}
		OnFulfilled(p.value)
		p.state = fulfilled
	}()
}

//The Catch method is applied on struct promise pointer and this takes OnRejected as argument
// wgGlobal.wait() checks if the wgGlobal.Done() is called or not (decremented the pointer by 1).
// We need to create a new promise before we can use Then method on it
// If promise struct has err value of not null, then OnRejected handler (function)
func (p *Promise) Catch(OnRejected func(error)) {
	func() {
		p.globalWG.Wait()
		if p.err != nil {
			OnRejected(p.err)
			p.state = rejected
			return
		}
	}()
}

//Same step as above functions with OnFinally function as argument
func (p *Promise) Finally(OnFinally func(int)) {
	func() {
		p.globalWG.Wait()
		OnFinally(p.state)
	}()
}

//This is a test function to return (string,error) pair. You can generate your own.
func TestingFunction() (string, error) {
	<-time.Tick(time.Second * 1)

	//Return response with nil error
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular1", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(req)
		return "Error", errors.New("Some error")
	}

	fmt.Println(err)
	return "req", nil
	//To return a response with error, uncomment below line.
	//return "", errors.New("This returns error")

	// return "req", err

}
