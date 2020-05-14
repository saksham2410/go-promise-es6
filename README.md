# go-promise-es6
Implementation of Promises in GoLang

A Promise is a proxy for a value not necessarily known when the promise is created. It allows you to associate handlers with an asynchronous action's eventual success value or failure reason. This lets asynchronous methods return values like synchronous methods: instead of immediately returning the final value, the asynchronous method returns a promise to supply the value at some point in the future.

A Promise is in one of these states:

pending: initial state, neither fulfilled nor rejected.
fulfilled: meaning that the operation completed successfully.
rejected: meaning that the operation failed.

The following files contains implementation of Promises in GoLang.
Go doesn't have prototypes but it has interfaces. Go interface matches the Promise prototype methods:

catch(onRejected)
then(onFulfilled[, onRejected])
finally(onFinally)

The main.go in test directory contains basic tests.

You can edit TestingFunction in promise.go package to create customs http calls os Database calls to check the working.

@saksham2410
