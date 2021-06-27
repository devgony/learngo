# Go For Beginners

## #0.2 Introduction (01:02)

- Go: multicore, concurrncy

## #0.5 Software and Installation (09:15)

- domain.username based download

## #1.0 Main Package (04:02)

- main is only for compile but mandatory
- function = func
- format = fmt

## #1.1 Packages and Imports (04:37)

- exporting function should start with uppercase

```
go mod init github.com/devgony/learngo
go run main.go
```

## #1.2 Variables and Constants (03:46)

### `const`: like const of JS

```go
const name string = "henry"
name = "Lynn" // cannot assign to name (declared const)
```

### `var`: like let of JS

```go
var vName string = "henry"
vName = "liha" // mutable
```

### Short-hand for `var`

- It guesses type with first assigned value

```go
vName := "henry"
```

## #1.3 Functions part One (08:40)

- go has types

```go
string
bool
int, int8, float32, float64 ...
```

- function

```go
func multiply(a int, b int) int {
	return a * b
}
```

- short-hand

```go
func multiply(a, b int) int {
```

- return void is empty?

### multiple return

```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
    totalLenght, upperName := lenAndUpper("henry")
    fmt.Println(totalLenght, upperName)
}
```

### Ignore return value

```go
totalLenght, _ := lenAndUpper("henry")
```

### Repeat params

```go
func repeatMe(words ...string) {
	fmt.Println(words)
}
repeatMe("a", "b", "c")
```

## #1.4 Functions part Two (05:23)

### Naked return

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

### defer

- do sth after func finished

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
...
}
```

## #1.5 for, range, ...args (05:59)

- for the only loop

```go
func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println((numbers[i]))
	}
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}
```

- accomulate

```go
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
```
