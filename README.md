# Mistake
My personal simple error handler built in Go.

## Why Mistake?
Our approach begins with designing an unified error handler just for being more clear due to `error` package is extremely flexible and then __Mistake__ born to be it.  
Also we have a preference on `go-playground/validator` package and was integrated from the beginning.

This package is inspired on API error handler and we recommend used it for this.

## How to start?
We have 2 ways to create a mistake object:  

Create a new single error with id and message.
``` go
func New(statusCode int, messageID string, message string) *M
```

But if your project uses the package `github.com/go-playground/validator`, you can use this constructor and pass the error.  
Internally the errors are converted to `validator.ValidationErrors` and retrieve __Field__ and __Tag__ attributes only.
``` go
func NewStructValidation(err error, statusCode int, messageID string, message string) *M
```

## Dynamic messages
Mistake provides two functions for those messages where use verbs (%v).  
Append values add all those variables will be replaced on the verbs.  
Formatter takes the message with verbs and apply `fmt.Sprintf` on it, at the end append values will be dropped.
``` go
func AppendValue(v any)
func Formatter(message string) string
```

Here an example:  
``` go
mistake.AppendValue(request.Email)
mt := mistake.New(
    resources.DUPLICATED,
    resources.CREATE_USER_EXIST,
    mistake.Formatter("here is my email: '%v'"), // here is my email: 'mymy@emailme.com'
)
```
