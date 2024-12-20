# tempconv — конвертер температур

`tempconv` — это пакет на Go, который предоставляет функциональность для конверсии значений
температуры между различными шкалами, включая:

- Цельсия (°C)
- Фаренгейта (°F)
- Кельвина (K)
- Ранкина (°R)
- Реомюра (°Re)

Пакет удобен для работы с температурными вычислениями и поддерживает строгую проверку
допустимости значений температуры (например, невозможность задать значения ниже абсолютного нуля).

## Особенности

- Поддержка 5 температурных шкал.
- Проверка значений температуры: предотвращение создания объектов с некорректными значениями
(ниже абсолютного нуля).
- Удобный интерфейс `Temperature`: позволяет работать с различными температурными шкалами через
единый интерфейс.
- Методы преобразования: каждый тип температуры имеет методы для конверсии в другие шкалы.
- Интуитивное строковое представление объектов температуры.

## Установка

Чтобы использовать этот пакет в своем проекте, выполните команду:

```zsh
go get github.com/MiCkEyZzZ/tempconv
```

## Пример использования

Ниже представлен пример, демонстрирующий работу пакета:

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tempconv/tempconv"
)

func main() {
	// Создание объекта температуры в Цельсиях
	tC, err := tempconv.NewCelsius(25)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Преобразование в Фаренгейт
	tF := tC.ToFahrenheit()
	fmt.Printf("25°C в Фаренгейтах: %s\n", tF)

	// Преобразование в Кельвины
	tK := tC.ToKelvin()
	fmt.Printf("25°C в Кельвинах: %s\n", tK)

	// Использование интерфейса Temperature
	var temp tempconv.Temperature = tC
	fmt.Printf("Температура через интерфейс (в Ранкинах): %s\n", temp.ToRankine())
}
```

## Вывод программы:

```zsh
25°C в Фаренгейтах: 77.00°F
25°C в Кельвинах: 298.15K
Температура через интерфейс (в Ранкинах): 536.67°R
```

## Константы
В пакете определены ключевые константы, включая абсолютные значения нуля для каждой шкалы:

- Абсолютный ноль по Цельсию: -273.15°C
- Абсолютный ноль по Фаренгейту: -459.67°F
- Абсолютный ноль по Кельвину: 0 K
- Абсолютный ноль по Ранкину: 0°R
- Абсолютный ноль по Реомюру: -218.52°Re

## API

Пакет предоставляет следующие основные функции:

### Конструкторы температур

- NewCelsius(c float64) (Celsius, error)
- NewFahrenheit(f float64) (Fahrenheit, error)
- NewKelvin(k float64) (Kelvin, error)
- NewRankine(r float64) (Rankine, error)
- NewReaumur(re float64) (Reaumur, error)

### Методы типов температуры

Все типы (Celsius, Fahrenheit, Kelvin, Rankine, Reaumur) реализуют:

- Конвертацию в другие шкалы: `ToCelsius`, `ToFahrenheit`, `ToKelvin`, `ToRankine`, `ToReaumur`.
- Строковое представление: `String`.
- Название шкалы: `ScaleName`.

## Проверка значений

Пакет автоматически проверяет, чтобы значения температур не были ниже
абсолютного нуля. Например:

```go
_, err := tempconv.NewCelsius(-300)
if err != nil {
    fmt.Println(err)
}

```

### Вывод:

```zsh
температура ниже абсолютного нуля: -300.00 °C ниже абсолютного нуля
```

## Лицензия

Этот пакет распространяется без лицензии и предоставляется "как есть". Вы можете использовать
его в своих проектах.
