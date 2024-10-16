package packages

import (
	"flag"
	"fmt"
	"os"
)

func FlagExample() {
	var (
		port   int
		host   string
		secure bool
	)
	
	flag.IntVar(&port, "port", 8080, "Порт для сервера")
	flag.StringVar(&host, "host", "localhost", "Хост для подключения")
	flag.BoolVar(&secure, "secure", false, "Использовать HTTPS")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование программы:\n")
		fmt.Fprintf(os.Stderr, "  [опции] <аргументы>\n")
		fmt.Fprintf(os.Stderr, "Опции:\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()
	
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: не указаны аргументы.")
		flag.Usage()
		os.Exit(1)
	}
	
	fmt.Printf("Запуск сервера на %s:%d (HTTPS: %t)\n", host, port, secure)
	fmt.Printf("Позиционные аргументы: %v\n", args)
}
