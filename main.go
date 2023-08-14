package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	// Define a string flag with name "lang", default value "en", and a usage description
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse() // Parse the flags from the command line
	greeting := greet(language(lang))
	fmt.Println(greeting) // Print the greeting
}

// language represents the language's code
type language string

// phrasebook holds greetings for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",       // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in various languages
// It returns an error message if the language is unsupported
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		// Return an error message if the language is not found in the phrasebook
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting // Return the greeting for the specified language
}
