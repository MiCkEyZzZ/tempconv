package tempconv

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

// almostEqual проверяет, что два числа почти равны с заданной погрешностью.
func almostEqual(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// TestNewCelsius проверяет создание объектов Celsius с корректными значениями и
// ошибкой для значений ниже абсолютного нуля.
func TestNewCelsius(t *testing.T) {
	tests := []struct {
		input    float64
		expected Celsius
		err      error
	}{
		{-273.15, -273.15, nil},
		{0, 0, nil},
		{-274, 0, ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Celsius %v", tt.input), func(t *testing.T) {
			c, err := NewCelsius(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && c != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, c)
			}
		})
	}
}

// TestNewFahrenheit проверяет создание объектов Fahrenheit с корректными значениями и
// ошибкой для значений ниже абсолютного нуля.
func TestNewFahrenheit(t *testing.T) {
	tests := []struct {
		input    float64
		expected Fahrenheit
		err      error
	}{
		{-459.67, -459.67, nil},
		{32, 32, nil},
		{-460, 0, ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Fahrenheit %v", tt.input), func(t *testing.T) {
			f, err := NewFahrenheit(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && f != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, f)
			}
		})
	}
}

// TestNewKelvin проверяет создание объектов Kelvin с корректными значениями и ошибкой для
// значений ниже абсолютного нуля.
func TestNewKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected Kelvin
		err      error
	}{
		{0, 0, nil},
		{273.15, 273.15, nil},
		{-1, 0, ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Kelvin %v", tt.input), func(t *testing.T) {
			k, err := NewKelvin(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && k != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, k)
			}
		})
	}
}

// TestNewRankine проверяет создание объектов Rankine с корректными значениями и ошибкой
// для значений ниже абсолютного нуля.
func TestNewRankine(t *testing.T) {
	tests := []struct {
		input    float64
		expected Rankine
		err      error
	}{
		{0, 0, nil},
		{459.67, 459.67, nil},
		{-1, 0, ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rankine %v", tt.input), func(t *testing.T) {
			r, err := NewRankine(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && r != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, r)
			}
		})
	}
}

// TestNewReaumur проверяет создание объектов Reaumur с корректными значениями и ошибкой
// для значений ниже абсолютного нуля.
func TestNewReaumur(t *testing.T) {
	tests := []struct {
		input    float64
		expected Reaumur
		err      error
	}{
		{-218.52, -218.52, nil},
		{0, 0, nil},
		{-219, 0, ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Reaumur %v", tt.input), func(t *testing.T) {
			re, err := NewReaumur(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && re != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, re)
			}
		})
	}
}

// TestNewDelisle проверяет создание объектов Delisle с корректными значениями и ошибкой
// для значений ниже абсолютного нуля.
func TestNewDelisle(t *testing.T) {
	tests := []struct {
		input    float64
		expected Delisle
		err      error
	}{
		{559.725, 559.725, nil},        // Абсолютный ноль в шкале Делисля
		{0, 0, nil},                    // Точка кипения воды
		{560, 0, ErrBelowAbsoluteZero}, // Температура ниже абсолютного нуля
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Delisle %v", tt.input), func(t *testing.T) {
			d, err := NewDelisle(tt.input)
			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && d != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, d)
			}
		})
	}
}

// TestConversions проверяет конверсии температур между различными шкалами.
func TestConversions(t *testing.T) {
	tests := []struct {
		fromCelsius        Celsius
		expectedFahrenheit Fahrenheit
		expectedKelvin     Kelvin
		expectedRankine    Rankine
		expectedReaumur    Reaumur
		expectedDelisle    Delisle
	}{
		{0, 32, 273.15, 491.67, 0, 150},
		{-273.15, -459.67, 0, 0, -218.52, 559.725},
		{100, 212, 373.15, 671.67, 80, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Convert %v°C", tt.fromCelsius), func(t *testing.T) {
			// Проверка конверсии в Фаренгейт
			if got := float64(tt.fromCelsius.ToFahrenheit()); !almostEqual(got, float64(tt.expectedFahrenheit), 0.01) {
				t.Errorf("ToFahrenheit() = %v, want %v", got, tt.expectedFahrenheit)
			}
			// Проверка конверсии в Кельвин
			if got := float64(tt.fromCelsius.ToKelvin()); !almostEqual(got, float64(tt.expectedKelvin), 0.01) {
				t.Errorf("ToKelvin() = %v, want %v", got, tt.expectedKelvin)
			}
			// Проверка конверсии в Ранкин
			if got := float64(tt.fromCelsius.ToRankine()); !almostEqual(got, float64(tt.expectedRankine), 0.01) {
				t.Errorf("ToRankine() = %v, want %v", got, tt.expectedRankine)
			}
			// Проверка конверсии в Реомюр
			if got := float64(tt.fromCelsius.ToReaumur()); !almostEqual(got, float64(tt.expectedReaumur), 0.01) {
				t.Errorf("ToReaumur() = %v, want %v", got, tt.expectedReaumur)
			}
			// Проверка конверсии в Делисль
			if got := float64(tt.fromCelsius.ToDelisle()); !almostEqual(got, float64(tt.expectedDelisle), 0.01) {
				t.Logf("From Celsius: %v, Calculated Delisle: %v", tt.fromCelsius, tt.fromCelsius.ToDelisle())
				t.Errorf("ToDelisle() = %v, want %v", got, tt.expectedDelisle)
			}
		})
	}
}

// TestInvalidTemperatures проверяет обработку ошибок для температур ниже
// абсолютного нуля.
func TestInvalidTemperatures(t *testing.T) {
	tests := []struct {
		input       float64
		scale       string
		expectedErr error
	}{
		{-274, "Celsius", ErrBelowAbsoluteZero},
		{-460, "Fahrenheit", ErrBelowAbsoluteZero},
		{-1, "Kelvin", ErrBelowAbsoluteZero},
		{-1, "Rankine", ErrBelowAbsoluteZero},
		{-219, "Reaumur", ErrBelowAbsoluteZero},
		{-560, "Delisle", ErrBelowAbsoluteZero},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Invalid %s %v", tt.scale, tt.input), func(t *testing.T) {
			var err error
			switch tt.scale {
			case "Celsius":
				_, err = NewCelsius(tt.input)
			case "Fahrenheit":
				_, err = NewFahrenheit(tt.input)
			case "Kelvin":
				_, err = NewKelvin(tt.input)
			case "Rankine":
				_, err = NewRankine(tt.input)
			case "Reaumur":
				_, err = NewReaumur(tt.input)
			case "Delisle":
				_, err = NewDelisle(tt.input)
			}

			if err != nil && !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

// TestStringAndScaleName проверяет корректность строкового представления и названия
// шкалы для различных температур.
func TestStringAndScaleName(t *testing.T) {
	tests := []struct {
		input             Temperature
		expectedString    string
		expectedScaleName string
	}{
		{Celsius(0), "0.00°C", "Celsius"},
		{Fahrenheit(32), "32.00°F", "Fahrenheit"},
		{Kelvin(273.15), "273.15K", "Kelvin"},
		{Rankine(491.67), "491.67°R", "Rankine"},
		{Reaumur(0), "0.00°Re", "Reaumur"},
		{Delisle(559.725), "559.725°De", "Delisle"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("String and ScaleName %v", tt.input), func(t *testing.T) {
			if got := tt.input.String(); got != tt.expectedString {
				t.Errorf("String() = %v, want %v", got, tt.expectedString)
			}
			if got := tt.input.ScaleName(); got != tt.expectedScaleName {
				t.Errorf("ScaleName() = %v, want %v", got, tt.expectedScaleName)
			}
		})
	}
}
