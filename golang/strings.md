generate random string

```go
func GetRandomString(length int) string {
  var symbolRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
  b := make([]rune, length)
  for i := range b {
    b[i] = symbolRunes[rand.Intn(len(symbolRunes))]
  }
  return string(b)
}
```

welcome message. used: 
* https://patorjk.com/software/taag/#p=display&f=Standard&t=new%20solution
* https://github.com/fatih/color

```go
func printHello() {
	str := `
                                  _       _   _             
  _ __   _____      __  ___  ___ | |_   _| |_(_) ___  _ __  
 | '_ \ / _ \ \ /\ / / / __|/ _ \| | | | | __| |/ _ \| '_ \ 
 | | | |  __/\ V  V /  \__ \ (_) | | |_| | |_| | (_) | | | |
 |_| |_|\___| \_/\_/   |___/\___/|_|\__,_|\__|_|\___/|_| |_|
                                                            `
	color.Blue(str)
	color.White("by Sagleft (github.com/sagleft). Automate and prosper.")
	fmt.Println()
}
```
