package tempconv

import (
	"errors"
	"fmt"
)

// Ошибки для недопустимых температур
var (
	ErrBelowAbsoluteZero = errors.New("температура ниже абсолютного нуля")
)

// Константы для температурных точек
const (
	// absoluteZeroC - абсолютный ноль по Цельсию (-273.15°C)
	absoluteZeroC Celsius = -273.15
	// absoluteZeroF - абсолютный ноль по Фаренгейту (-459.67°F)
	absoluteZeroF Fahrenheit = -459.67
	// absoluteZeroK - абсолютный ноль по Кельвину (0K)
	absoluteZeroK Kelvin = 0
	// absoluteZeroR - абсолютный ноль по Ранкину (0°R)
	absoluteZeroR Rankine = 0
	// absoluteZeroRe - абсолютный ноль по Реомюру (-218.52°Re)
	absoluteZeroRe Reaumur = -218.52
)

// Temperature - интерфейс для работы с температурой.
type Temperature interface {
	ToCelsius() Celsius
	ToFahrenheit() Fahrenheit
	ToKelvin() Kelvin
	ToRankine() Rankine
	ToReaumur() Reaumur
	String() string
	ScaleName() string
}

// Типы для температурных шкал
type (
	// Celsius - тип для шкалы Цельсия
	Celsius float64
	// Fahrenheit - тип для шкалы Фаренгейта
	Fahrenheit float64
	// Kelvin - тип для шкалы Кельвина
	Kelvin float64
	// Rankine - тип для шкалы Ранкина
	Rankine float64
	// Reaumur - тип для шкалы Реомюра
	Reaumur float64
)

// Константы для преобразования температур
const (
	// cToFMultiplier - коэффициент для преобразования из Цельсия в Фаренгейт
	cToFMultiplier = 9.0 / 5.0
	// cToFOffset - смещение для преобразования из Цельсия в Фаренгейт
	cToFOffset = 32.0
	// cToKOffset - смещение для преобразования из Цельсия в Кельвин
	cToKOffset = 273.15
	// cToRMultiplier - коэффициент для преобразования из Цельсия в Ранкин
	cToRMultiplier = cToFMultiplier
	// cToROffset - смещение для преобразования из Цельсия в Ранкин
	cToROffset = cToKOffset
	// cToReMultiplier - коэффициент для преобразования из Цельсия в Реомюр
	cToReMultiplier = 4.0 / 5.0
	// fToCOffset - смещение для преобразования из Фаренгейта в Цельсий
	fToCOffset = cToFOffset
	// fToCMultiplier - коэффициент для преобразования из Фаренгейта в Цельсий
	fToCMultiplier = 1.0 / cToFMultiplier
	// fToROffset - смещение для преобразования из Фаренгейта в Ранкин
	fToROffset = 459.67
	// fToKMultiplier - коэффициент для преобразования из Фаренгейта в Кельвин
	fToKMultiplier = fToCMultiplier
	// fToRMultiplier - коэффициент для преобразования из Фаренгейта в Ранкин
	fToRMultiplier = 1.0
	// kToCOffset - смещение для преобразования из Кельвина в Цельсий
	kToCOffset = cToKOffset
	// kToRMultiplier - коэффициент для преобразования из Кельвина в Ранкин
	kToRMultiplier = cToFMultiplier
	// rToFOffset - смещение для преобразования из Ранкина в Фаренгейт
	rToFOffset = fToROffset
	// rToCMultiplier - коэффициент для преобразования из Ранкина в Цельсий
	rToCMultiplier = fToCMultiplier
	// rToKMultiplier - коэффициент для преобразования из Ранкина в Кельвин
	rToKMultiplier = 5.0 / 9.0
	// reToCMultiplier - коэффициент для преобразования из Реомюра в Цельсий
	reToCMultiplier = 5.0 / 4.0
)

// NewCelsius создает объект Цельсий и проверяет, что значение температуры
// не ниже абсолютного нуля по Цельсию (-273.15°C). Если значение корректно,
// возвращается объект Цельсий, иначе - ошибка.
func NewCelsius(c float64) (Celsius, error) {
	if err := validateTemperature(c, float64(absoluteZeroC), "°C"); err != nil {
		return 0, err
	}
	return Celsius(c), nil
}

// NewFahrenheit создает объект Фаренгейт и проверяет, что значение температуры
// не ниже абсолютного нуля по Фаренгейту (-459.67°F). Если значение корректно,
// возвращается объект Фаренгейт, иначе - ошибка.
func NewFahrenheit(f float64) (Fahrenheit, error) {
	if err := validateTemperature(f, float64(absoluteZeroF), "°F"); err != nil {
		return 0, err
	}
	return Fahrenheit(f), nil
}

// NewKelvin создает объект Кельвин и проверяет, что значение температуры
// не ниже абсолютного нуля по Кельвину (0K). Если значение корректно,
// возвращается объект Кельвин, иначе - ошибка.
func NewKelvin(k float64) (Kelvin, error) {
	if err := validateTemperature(k, float64(absoluteZeroK), "K"); err != nil {
		return 0, err
	}
	return Kelvin(k), nil
}

// NewRankine создает объект Ранкин и проверяет, что значение температуры
// не ниже абсолютного нуля по Ранкину (0°R). Если значение корректно,
// возвращается объект Ранкин, иначе - ошибка.
func NewRankine(r float64) (Rankine, error) {
	if err := validateTemperature(r, float64(absoluteZeroR), "°R"); err != nil {
		return 0, err
	}
	return Rankine(r), nil
}

// NewReaumur создает объект Реомюр и проверяет, что значение температуры
// не ниже абсолютного нуля по Реомюру (-218.52°Re). Если значение корректно,
// возвращается объект Реомюр, иначе - ошибка.
func NewReaumur(re float64) (Reaumur, error) {
	if err := validateTemperature(re, float64(absoluteZeroRe), "°Re"); err != nil {
		return 0, err
	}
	return Reaumur(re), nil
}

// Реализация методов для типа Celsius

// ToCelsius возвращает температуру в шкале Цельсия (сама по себе).
// Метод возвращает объект типа Celsius, который уже представляет температуру в шкале Цельсия.
func (c Celsius) ToCelsius() Celsius { return c }

// ToFahrenheit преобразует температуру из Цельсия в Фаренгейт.
// Метод возвращает объект типа Fahrenheit, который представляет температуру в шкале Фаренгейта.
func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*cToFMultiplier + cToFOffset) }

// ToKelvin преобразует температуру из Цельсия в Кельвин.
// Метод возвращает объект типа Kelvin, который представляет температуру в шкале Кельвина.
func (c Celsius) ToKelvin() Kelvin { return Kelvin(c + cToKOffset) }

// ToRankine преобразует температуру из Цельсия в Ранкин.
// Метод возвращает объект типа Rankine, который представляет температуру в шкале Ранкина.
func (c Celsius) ToRankine() Rankine { return Rankine((c + cToROffset) * cToRMultiplier) }

// ToReaumur преобразует температуру из Цельсия в Реомюр.
// Метод возвращает объект типа Reaumur, который представляет температуру в шкале Реомюра.
func (c Celsius) ToReaumur() Reaumur { return Reaumur(c * cToReMultiplier) }

// String возвращает строковое представление температуры в шкале Цельсия.
func (c Celsius) String() string { return fmt.Sprintf("%.2f°C", c) }

// ScaleName возвращает строковое название шкалы температуры (Цельсий).
func (c Celsius) ScaleName() string { return "Celsius" }

// Реализация методов для типа Fahrenheit

// ToFahrenheit возвращает температуру в шкале Фаренгейта (сама по себе).
// Метод возвращает объект типа Fahrenheit, который уже представляет температуру в шкале Фаренгейта.
func (f Fahrenheit) ToFahrenheit() Fahrenheit { return f }

// ToCelsius преобразует температуру из Фаренгейта в Цельсий.
// Метод возвращает объект типа Celsius, который представляет температуру в шкале Цельсия.
func (f Fahrenheit) ToCelsius() Celsius { return Celsius((f - fToCOffset) * fToCMultiplier) }

// ToKelvin преобразует температуру из Фаренгейта в Кельвин.
// Метод возвращает объект типа Kelvin, который представляет температуру в шкале Кельвина.
func (f Fahrenheit) ToKelvin() Kelvin { return Kelvin((f-fToCOffset)*fToKMultiplier + kToCOffset) }

// ToRankine преобразует температуру из Фаренгейта в Ранкин.
// Метод возвращает объект типа Rankine, который представляет температуру в шкале Ранкина.
func (f Fahrenheit) ToRankine() Rankine { return Rankine(f + fToROffset) }

// ToReaumur преобразует температуру из Фаренгейта в Реомюр.
// Метод возвращает объект типа Reaumur, который представляет температуру в шкале Реомюра.
func (f Fahrenheit) ToReaumur() Reaumur { return f.ToCelsius().ToReaumur() }

// String возвращает строковое представление температуры в шкале Фаренгейта.
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }

// ScaleName возвращает строковое название шкалы температуры (Фаренгейт).
func (f Fahrenheit) ScaleName() string { return "Fahrenheit" }

// Реализация методов для типа Kelvin

// ToKelvin возвращает температуру в шкале Кельвина (сама по себе).
// Метод возвращает объект типа Kelvin, который уже представляет температуру в шкале Кельвина.
func (k Kelvin) ToKelvin() Kelvin { return k }

// ToCelsius преобразует температуру из Кельвина в Цельсий.
// Метод возвращает объект типа Celsius, который представляет температуру в шкале Цельсия.
func (k Kelvin) ToCelsius() Celsius { return Celsius(k - kToCOffset) }

// ToFahrenheit преобразует температуру из Кельвина в Фаренгейт.
// Метод возвращает объект типа Fahrenheit, который представляет температуру в шкале Фаренгейта.
func (k Kelvin) ToFahrenheit() Fahrenheit { return k.ToCelsius().ToFahrenheit() }

// ToRankine преобразует температуру из Кельвина в Ранкин.
// Метод возвращает объект типа Rankine, который представляет температуру в шкале Ранкина.
func (k Kelvin) ToRankine() Rankine { return Rankine(k * kToRMultiplier) }

// ToReaumur преобразует температуру из Кельвина в Реомюр.
// Метод возвращает объект типа Reaumur, который представляет температуру в шкале Реомюра.
func (k Kelvin) ToReaumur() Reaumur { return k.ToCelsius().ToReaumur() }

// String возвращает строковое представление температуры в шкале Кельвина.
func (k Kelvin) String() string { return fmt.Sprintf("%.2fK", k) }

// ScaleName возвращает строковое название шкалы температуры (Кельвин).
func (k Kelvin) ScaleName() string { return "Kelvin" }

// Реализация методов для типа Rankine

// ToRankine возвращает температуру в шкале Ранкина (сама по себе).
// Метод возвращает объект типа Rankine, который уже представляет температуру в шкале Ранкина.
func (r Rankine) ToRankine() Rankine { return r }

// ToCelsius преобразует температуру из Ранкина в Цельсий.
// Метод возвращает объект типа Celsius, который представляет температуру в шкале Цельсия.
func (r Rankine) ToCelsius() Celsius { return Celsius((r - rToFOffset) * rToCMultiplier) }

// ToFahrenheit преобразует температуру из Ранкина в Фаренгейт.
// Метод возвращает объект типа Fahrenheit, который представляет температуру в шкале Фаренгейта.
func (r Rankine) ToFahrenheit() Fahrenheit { return Fahrenheit(r - rToFOffset) }

// ToKelvin преобразует температуру из Ранкина в Кельвин.
// Метод возвращает объект типа Kelvin, который представляет температуру в шкале Кельвина.
func (r Rankine) ToKelvin() Kelvin { return Kelvin(r * rToKMultiplier) }

// ToReaumur преобразует температуру из Ранкина в Реомюр.
// Метод возвращает объект типа Reaumur, который представляет температуру в шкале Реомюра.
func (r Rankine) ToReaumur() Reaumur { return r.ToCelsius().ToReaumur() }

// String возвращает строковое представление температуры в шкале Ранкина.
func (r Rankine) String() string { return fmt.Sprintf("%.2f°R", r) }

// ScaleName возвращает строковое название шкалы температуры (Ранкин).
func (r Rankine) ScaleName() string { return "Rankine" }

// Реализация методов для типа Reaumur

// ToReaumur возвращает температуру в шкале Реомюра (сама по себе).
// Метод возвращает объект типа Reaumur, который уже представляет температуру в шкале Реомюра.
func (re Reaumur) ToReaumur() Reaumur { return re }

// ToCelsius преобразует температуру из Реомюра в Цельсий.
// Метод возвращает объект типа Celsius, который представляет температуру в шкале Цельсия.
func (re Reaumur) ToCelsius() Celsius { return Celsius(re * reToCMultiplier) }

// ToFahrenheit преобразует температуру из Реомюра в Фаренгейт.
// Метод возвращает объект типа Fahrenheit, который представляет температуру в шкале Фаренгейта.
func (re Reaumur) ToFahrenheit() Fahrenheit { return re.ToCelsius().ToFahrenheit() }

// ToKelvin преобразует температуру из Реомюра в Кельвин.
// Метод возвращает объект типа Kelvin, который представляет температуру в шкале Кельвина.
func (re Reaumur) ToKelvin() Kelvin { return re.ToCelsius().ToKelvin() }

// ToRankine преобразует температуру из Реомюра в Ранкин.
// Метод возвращает объект типа Rankine, который представляет температуру в шкале Ранкина.
func (re Reaumur) ToRankine() Rankine { return re.ToCelsius().ToRankine() }

// String возвращает строковое представление температуры в шкале Реомюра.
func (re Reaumur) String() string { return fmt.Sprintf("%.2f°Re", re) }

// ScaleName возвращает строковое название шкалы температуры (Реомюр).
func (re Reaumur) ScaleName() string { return "Reaumur" }

// ValidateTemperature проверяет, что температура не ниже абсолютного нуля для
// соответствующей шкалы и возвращает ошибку, если температура некорректна.
func validateTemperature(value, absoluteZero float64, scale string) error {
	if value < absoluteZero {
		return fmt.Errorf("%w: %.2f %s ниже абсолютного нуля", ErrBelowAbsoluteZero, value, scale)
	}
	return nil
}
