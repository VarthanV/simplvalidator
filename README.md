# simplvalidator

Golang  package provides helper functions for commonly use validations that we do in our projects

## Features
- Simple ,as the name says, It has simple yet robust logic and intiutive function names.
- You can fork this project and extend according to your needs and use it in your projects/organisations.
- Functions specific for Indian use 🇮🇳🇮🇳 (More will be added PR's are welcome).
- Validated against test cases.

## Usage

```go
    import "github.com/VarthanV/simplvalidator"

    func main() {
	email := "foo@gmail.com"
	if ok := IsEmail(email); ok {
		// Business logic
	}
}

```

## TODO 
[] Add test cases for all the functions <br/>
[] Comprehensive documentation <br/>
[] Add more validation functions 

PRS are welcome !

