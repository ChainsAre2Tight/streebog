# Функция шеширования Стрибог (ГОСТ 34.11 - 2018)

Проект реализует функции хеширования Стрибог с длинами хеш-кода 256 и 512 бит, реализующий интерфейс hash.Hash.

## Установка

```bash
go get github.com/ChainsAre2Tight/streebog
```

## Пример использования 


#### HASH
```go
package main

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	message := []byte("any-message")
	hash := streebog.New(64)
	hash.Write(message)
	fmt.Printf("Hash: %x\n", hash.Sum(nil))
}


```
#### HMAC
```Go
package main

import (
	"crypto/hmac"
	"fmt"
	"hash"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	key := []byte("any-key")
	message := []byte("any-message")
	h := hmac.New(func() hash.Hash { return streebog.New(32) }, key)
	h.Write(message)
	fmt.Printf("Hash: %x\n", h.Sum(nil))
}
```
## Производительность
~60 kH/s
# <img alt="benchmark-results" src="https://raw.githubusercontent.com/ChainsAre2Tight/streebog/refs/heads/master/examples/benchmark.png">