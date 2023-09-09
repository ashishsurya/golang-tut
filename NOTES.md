# Golang Tutorial
by Frontend Masters


## Functions

Some important things about go funcitons

- Last argument can variable length
- Functions can return more than one value at once
- Functions can return labeled variables
- Functions recieve arguments **by value**


```go
// some examples of go functions

func getFullname(firstName string, lastName string) string {
  return firstName + lastName
}

```
> if all the arguments are having same type, then the type 

```go
func getFullname(firstName , lastName string) string {
  // return
}
```
> functions can return more than one value
```go
func getUserDetails(id string) (string , uint) {
  // do some api call based on id

  return fullname , age
}

// use the above function as
fullName, age := getUserDetails(id)

// if you dont want fullname
_, age := getUseDetails(id)
//or
fullName, _ getUserDetails(id)
```
> variables returned can be labeld
```go
func getUserDetails(id string) (fullName string, age uint) {
  // do some api call based on id
  fullName := "Surya"
  age := 22
  return
}

// this function works same as the above function
```

## Pointers

```go
func birthday(age int) {
	age = age + 1
}

func main()  {
	age := 1
	fmt.Println(age)
	birthday(age)
	fmt.Println(age)
}

// OUTPUT
// 1
// 1
```
here the age variable is passed by value.

>To solve this issue pointers were included
```go
func birthday(age *int) {
	*age = *age + 1
}

func main()  {
	age := 1
	fmt.Println(age)
	birthday(&age)
	fmt.Println(age)
}

// OUTPUT
// 1
// 2
```

## Panic, Defer and Error Design Patern

- Panic is Go's style of throwing error messages, similar to `throw new Error` in Javascript.


- Defer runs the code which is specified at the end of the function call in which the defer is initiated
- `defer` is used for clean up functions
- Error design pattern is notifying the person who calls the function about the reason of any failure by providing the reason in return viaraible

```go
func getUser(id string) (user, err) {
  // do some db call

  if everythingFine {
    return user, nil
  } else {
    return nil , errorDetails
  }
}


func main() {
  user, err := getUser(2)
}
```


