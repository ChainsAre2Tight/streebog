# Функция шеширования Стрибог (ГОСТ 34.11 - 2018)

Проект реализует функции хеширования Стрибог с длинами хеш-кода 256 и 512 бит.

## Установка

```bash
go get github.com/ChainsAre2Tight/streebog
```

## Пример использования 

```go
package main

import (
	"fmt"
	"log"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	message := []byte("any-message")
	hash, err := streebog.Streebog512(message)
	if err != nil {
		log.Fatalf("Error during hash computation: %s", err)
	}
	fmt.Printf("Hash: %x\n", hash)
}

```

## Производительность
~60 kH/s
# <img alt="benchmark-results" src="https://raw.githubusercontent.com/ChainsAre2Tight/streebog/refs/heads/master/examples/benchmark.png">