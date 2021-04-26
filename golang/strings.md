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
