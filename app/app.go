package app

import (
	"fmt"

	"github.com/urfave/cli"
)

func ConverterGraus() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação Para Converter Temperatura"
	app.Usage = "Converte entre Celsius e Fahrenheit"

	flags := []cli.Flag{
		cli.Float64Flag{
			Name:  "c",
			Usage: "Temperatura em graus Celsius",
		},
		cli.Float64Flag{
			Name:  "f",
			Usage: "Temperatura em graus Fahrenheit",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "celsius",
			Aliases: []string{"ctof"},
			Usage: "Converte de Celcius em Fahrenheit",
			Flags: flags,
			Action: func(c *cli.Context) error {
				celsius := c.Float64("c")
				fahrenheit := celsiusToFahrenheit(celsius)
				fmt.Printf("%.2f graus Celsius é igual a %.2f graus Fahrenheit\n", celsius, fahrenheit)
				return nil
			},
		},
		{
			Name:  "fahrenheit",
			Aliases: []string{"ftoc"},
			Usage: "Converte de Fahrenheit em Celcius",
			Flags: flags,
			Action: func(c *cli.Context) error {
				fahrenheit := c.Float64("f")
				celsius := fahrenheitToCelsius(fahrenheit)
				fmt.Printf("%.2f graus Fahrenheit é igual a %.2f graus Celsius\n", fahrenheit, celsius)
				return nil
			},
		},
	}

	return app
}

// FahrenheitToCelsius converte graus Fahrenheit para Celsius.
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// CelsiusToFahrenheit converte graus Celsius para Fahrenheit.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
