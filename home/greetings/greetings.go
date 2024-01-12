package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
    // if no name was given, return an error with a message
    if name == "" {
        return "", errors.New("empty name")
    }


    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}


// In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).
// long way:
// var message string
// message = fmt.Sprintf("Hi, %v. Welcome!", name)
// the first argument is a format string, sprintf substitutes the name parameter's value for the %v format verb


// Hellos returns a map that associates each of the named people
// with a greeting message

func Hellos(names []string) (map[string]string, error) {
    // a map to associate names with messages
    messages := make(map[string]string)
    // loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // in the map, associate the retrieved message with the name
        messages[name] = message
    }
    return messages,nil
}
// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
    // a slice of message formats
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}

