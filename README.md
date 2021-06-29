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

# #1.6 If with a Twist (03:59)

## Variable expression in if: would excluesively be used only inside if

```go
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
```

# #1.7 Switch (02:44)

## Variable expression is available as well

```go
func canMyAgeDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

```

## expression at case is available

```go
func canMyAgeDrink(age int) bool {
	switch {
	case age < 10:
		return false
	case age == 18:
		return true
	}
	return false
}
```

# #1.8 Pointers! (07:24)

## `&`: address of var

## `*`: look through address

```go
func main() {
	a := 2  // assign
	b := a  // value of a
	c := &a // address of a
	a = 5   // c is affected
	fmt.Println(&a, &b, c, b, *c)
	// 0xc00001a0a8 0xc00001a0c0 0xc00001a0a8 2 5
	*c = 20 // change value of a
	fmt.Println(a)
	// 20
}
```

# #1.9 Arrays and Slices (04:01)

## `array`: should specify length

## `slice`: no need explicit length but is decide when define

## `append(slice, elem)`: keep same address (doesn't modify and return new array like js)

## `slice := [...]string{"nico", "dal", "flynn"}`: can't even append

```go
array := [4]string{"nico", "dal", "flynn"}
array[3] = "henry"
array[4] = "liha" // index 4 (constant of type int) is out of bounds
slice := []string{"nico", "dal", "flynn"}
slice[4] = "liha" // index out of range [4] with length 3
slice = append(slice, "a")
fmt.Println(array, slice)
```

# #1.10 Maps (03:37)

## map: object with same type

## `map[keyType]valueType{"key1": "val1",}`

```go
nico := map[string]string{"name": "nico", "age": "12"}
	for key, index := range nico {
		fmt.Println(key, index)
	}
```

# #1.11 Structs (06:38)

## struct: object + class

## no constructor => have to make manually

```go
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi, ramen"}
	nico1 := person{"nico", 18, favFood}
	nico2 := person{name: "nico", age: 18, favFood} // mixture of field:value and value elements in struct literal
	nico3 := person{name: "nico", favFood: favFood, age: 18}
	fmt.Println(nico1, nico3.name)
}
```
