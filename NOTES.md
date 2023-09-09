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




