# Dependency Injection

In software engineering, [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) is a design pattern in which an object receives other objects that it depends on. 
A form of inversion of control, dependency injection aims to separate the concerns of constructing objects and using them.

## Examples

### Example with direct dependency

In this case Pay depends directly on PayMethod and you are not able to 
easily change PayMethod. 
Notice that if you want to write Unit Test for `Pay`
you will face issue with switching PayMethod.

```go
func Pay() {
    // initialize a pay method inside Pay
    payMethod := PayMethod{
        name: "CB"
    }
	
    // use PayMethod to do things
    if payMethod === "CB" {
        println("pay with CB success")
    } else if payMethod === "VISA" {
        println("pay with VISA success")
    }
}
```

### Example using dependency injection

In this case Pay gets PayMethod front its arguments.
Notice that if you want to write Unit Test for `Pay`
you pass in any PayMethod you want in the argument.

```go
// get pay method from argument
func Pay(payMethod PayMethod) {
    // use PayMethod to do things
    if payMethod === "CB" {
        println("pay with CB success")
    } else if payMethod === "VISA" {
        println("pay with VISA success")
    }
}
```

## Exercice

Fetch the project and run the `unit test`
```shell
> git clone git@github.com:adeo/dependency-injection.git
> go test ./...
--- FAIL: TestComputeWeeksLeftBeforeNextYear (0.04s)
    compute-weeks-left_test.go:10:
        expected: [~ 46 weeks left this year]
        received: [~ 36 weeks left this year]
FAIL
FAIL	main/cmd	0.158s
FAIL
```

Now you should try to use dependency injection to make the unit test pass 😉🤞🏽
