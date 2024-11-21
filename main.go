package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/tempconv/tempconv"
)

func main() {
	// Создаем значение температуры в градусах Цельсия
	tC, err := tempconv.NewCelsius(100)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Преобразуем в градусы Фаренгейта
	tF := tC.ToFahrenheit()
	fmt.Printf("%.2f°C в Фаренгейтах: %.2f°F\n", tC, tF)

	// Преобразуем в Кельвины
	tK := tC.ToKelvin()
	fmt.Printf("%.2f°C в Кельвинах: %.2fK\n", tC, tK)

	// Преобразуем в шкалу Делисля
	tDe := tC.ToDelisle()
	fmt.Printf("%.2f°C в шкале Делисля: %.2f°De\n", tC, tDe)

	// Используем интерфейс Temperature для универсального доступа
	var temp tempconv.Temperature = tC
	fmt.Printf("Температура в Ранкинах: %.2f°R\n", temp.ToRankine())
}
