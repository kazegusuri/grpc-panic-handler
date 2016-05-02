# grpc-panic-handler

grpc-panic-handler is an interceptor to protect a process from aborting by panic and return Internal error as status code.

## Usage

```
import (
	panichandler "github.com/kazegusuri/grpc-panic-handler"
)

func main() {
	uIntOpt := grpc.UnaryInterceptor(panichandler.UnaryPanicHandler)
	sIntOpt := grpc.StreamInterceptor(panichandler.StreamPanicHandler)
	grpc.NewServer(uIntOpt, sIntOpt)
}
```

## Custom Panic Handler

You can write custom panic handler in case of panic. Use `InstallPanicHandler`.

```
func main() {
	panichandler.InstallPanicHandler(func(r interface{}) {
		fmt.Printf("panic happened: %v", r)
	}
}
```


### Built-in custom panic handler

- LogPanicDump
 - `debug.Stack()` to stderr
- LogPanicStackMultiLine
- show stack trace in multi line by glog

## Copyright

<table>
  <tr>
    <td>Author</td><td>Masahiro Sano <sabottenda@gmail.com></td>
  </tr>
  <tr>
    <td>Copyright</td><td>Copyright (c) 2016- Masahiro Sano</td>
  </tr>
  <tr>
    <td>License</td><td>MIT License</td>
  </tr>
</table>
