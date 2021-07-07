# Go For Beginners

```
VSC => install all => lint tool: golint
```

# #0.2 Introduction (01:02)

- Go: multicore, concurrncy

# #0.5 Software and Installation (09:15)

- domain.username based download

# #1.0 Main Package (04:02)

- main is only for compile but mandatory
- function = func
- format = fmt

# #1.1 Packages and Imports (04:37)

## Exporting function should start with `uppercase`

```
go mod init github.com/devgony/learngo
go run main.go
```

# #1.2 Variables and Constants (03:46)

## `const`: like const of JS

```go
const name string = "henry"
name = "Lynn" // cannot assign to name (declared const)
```

## `var`: like let of JS

```go
var vName string = "henry"
vName = "liha" // mutable
```

## Short-hand for `var`

- It guesses type with first assigned value

```go
vName := "henry"
```

# #1.3 Functions part One (08:40)

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

## multiple return

```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
    totalLenght, upperName := lenAndUpper("henry")
    fmt.Println(totalLenght, upperName)
}
```

## Ignore return value with `_`

```go
totalLenght, _ := lenAndUpper("henry")
```

## Repeat params

```go
func repeatMe(words ...string) {
	fmt.Println(words)
}
repeatMe("a", "b", "c")
```

# #1.4 Functions part Two (05:23)

## Naked return

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

## defer

- do sth after func finished

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
...
}
```

# #1.5 for, range, ...args (05:59)

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

## `array := [index]valType{v1, v2...}`: should specify length

## `slice := []valType{v1, v2, ...}`: no need explicit length but is decide when define

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

## map: object with unified type

## need to init first with `v := map[string]string{}`

## `map[keyType]valueType{"key1": "val1",}`

```go
nico := map[string]string{"name": "nico", "age": "12"}
	for key, index := range nico {
		fmt.Println(key, index)
	}
```

# #1.11 Structs (06:38)

## struct: object + class

1. Define type
2. assign value with values way or key:value way (should be unified)

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
	nico2 := person{name: "nico", age: 18, favFood} // Err: mixture of field:value and value elements in struct literal
	nico3 := person{name: "nico", favFood: favFood, age: 18} // key: value way
	fmt.Println(nico1, nico3.name)
}
```

# #2.0 Account + NewAccount (09:45)

## To export struct, or property, name should start with UpperCase

### If lower, cannot access from other package

## Use lowerCase and manipulate with func (starting with Upper)

```go
// mkdir accounts
// touch accounts/accounts.go
package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// main.go
func main() {
	account := accounts.NewAccount("nico")
	account.balance = 10 // Err: account.balance undefined (type *accounts.Account has no field or method balance)
	// if lower, cannot access
}
```

# #2.1 Methods part One (05:05)

## Receiver: similar to method

- naming convention: starts with lower first char of struct

```go
// Deposit n amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println(amount)
	a.balance += amount // copy of struct
}
```

# #2.2 Methods part Two (09:31)

## every time, a is copy , to protect origin method

## Pointer recevier

```go
func (a *Account) Deposit(amount int) {
	a.balance += amount // original struct
```

## error handling

## no trycatch => handle error manually

```go
log.Fatalln(err)
```

## error var should start with `err`

```go
var errNoMoney = errors.New("Can't withdraw you are poor")

// Withdraw n amount
func (a *Account) Withdraw(amonut int) error {
	if a.balance < amonut {
		return errNoMoney
	}
	a.balance -= amonut
	return nil
}
```

```go
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(50)
	if err != nil {
		fmt.Println(err) // or log.Fatalln(err)
	}
	fmt.Println(account)
}
```

## Overwrite string: when printing struct, automatically call String()

```go
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

func main() {
	fmt.Println(account) // > nico's account.\nHas: 10
}
```

# #2.4 Dictionary part One (07:57)

## Dictionary: custom type with map

- `map[key]` returns value, exists

```go
// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func main() {
	dictionary := mydict.Dictionary{"key1": "1stWord"}
	dictionary["key2"] = "value2"
	fmt.Println(dictionary["key1"])
	definition, err := dictionary.Search("key2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
```

# #2.5 Add Method (05:58)

## map receiver automatically use \* (not copied)

```go
var errWordExists = errors.New("That word already exists.")

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition: ", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
```

# #2.6 Update Delete (08:00)

```go
var errCantUpdate = errors.New("Can't update non-existing word")

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary.Search(word))
	dictionary.Update(word, "bye")
	fmt.Println(dictionary.Search(word))
	dictionary.Delete(word)
	fmt.Println(dictionary.Search(word))
}
```

# #3.0 hitURL (06:07)

## http.Get: std lib retuning `response, err`

```go
func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errReqeustFailed
	}
	return nil
}
```

# #3.1 Slow URLChecker (07:33)

## map without initialization get error

```go
var results map[string]string
results[url] = result // Err: panic: assignment to entry in nil map
```

## init map with make

```go
var results = make(map[string]string)
```

## init map with {}

```go
var results = make(map[string]string){}
```

# #3.2 Goroutines (08:47)

## run at background but main doesn't wait goRoutines

```go
func main() {
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
```

# #3.3 Channels (11:05)

## Using `<-c`, wait for routine

## Blocking operation: If wait for just one channel, it finishes ASA receive first resp

### If over than the number => deadlock

```go
func main() {
	c := make(chan string) // init channel
	people := [2]string{"nico", "henry"}
	for _, person := range people {
		fmt.Println("executed: ", person)
		go isSexy(person, c) // send channel n times
	}
	fmt.Println(<-c) // if wait for just one channel, it finishes ASA receive first resp
	fmt.Println(<-c)
}

func isSexy(person string, c chan string) {
	if person == "henry" {
		time.Sleep(time.Second * 3)
	} else {
		time.Sleep(time.Second * 5)
	}
	c <- person + " is returned" // return resp when routine finished
}
```

# #3.4 Channels Recap (07:37)

## iterate blocking operation

```go
for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
```

# #3.5 One more Recap (04:30)

## explicit send only channel

```go
func hitURL(url string, c chan<- requestResult) { // send only
...
	c <- requestResult{url: url, status: status}
...
```
