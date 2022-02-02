### Spinner

```go 
func main() {
	var done bool
	
	go Spinner(&done, 50)

	time.Sleep(time.Second * 5) // long-running function

	done = true

	time.Sleep(time.Second * 3) // program operation
}
```
### Indicator
```go 
func main() {
	var done bool
	
	go PrBar(&done, "*", ".", 200, 50)

	time.Sleep(time.Second * 10) // long-running function

	done = true

	time.Sleep(time.Second * 3) // program operation
}
```

### Imitation indicator

```go 
func main() {
    var done bool
    
    go PrBarImitation(&done, "*", ".", 50)

	time.Sleep(time.Second * 15) // long-running function

	done = true

	time.Sleep(time.Second * 3) // program operation
}
```

### Self-writing text

```go 
func main() {
	var done bool

	go SWrText(&done, "Please wait. Loading...", 150, 3, 500)

	time.Sleep(time.Second * 15) // long-running function

	done = true

	time.Sleep(time.Second * 3) // program operation
}
```