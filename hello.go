
package main

import (
	"os"
	"fmt"
	"log"			// https://pkg.go.dev/log
	// "rsc.io/quote"		// https://pkg.go.dev/rsc.io/quote/v4
	"github.com/Naegolus/greetings-go"
)

func main () {

	log.SetPrefix("greetings: ")
	log.SetFlags(
		log.LstdFlags |
		log.Lshortfile |
		log.Lmsgprefix)

	//fmt.Println(quote.Go())

	var msg string
	var err error

	err, msg = greetings.Hello("Joe")
	if err != nil {
		// log.Fatal(err)
		log.Print(err)
		os.Exit(1)
	}

	fmt.Println(msg)

	var names []string

	names = []string{"Gladys", "Samantha", "Darrin"}

	err, msgs := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(msgs)

	for name, msg := range msgs {
		fmt.Print(name + ": ")
		fmt.Println(msg)
	}
}

