# Programmatic Hello

This repository contains a simple program (adapted from "Learn Go with Pocket-Sized Projects" by Aliénor Latour, Donia Chaiehloudj and Pascal Bertrand) that demonstrates greeting the world in various languages. The language can be selected using command-line flags. Additionally, a table-driven test approach is employed to ensure that the correct greetings are returned.

## Main Features

1. **Multi-Language Support**: Greetings in English, Greek, French, Hebrew, Urdu, and Vietnamese are supported.
2. **Command-Line Flags**: The language can be specified using a command-line flag.
3. **Table-Driven Tests**: Utilizes table-driven tests for comprehensive and maintainable testing.
4. **Hash Maps**: Efficient language lookup using Go's hash maps.

## Usage

Compile and run the program with the following command:

```bash
go run main.go -lang=en

Replace `en` with the desired language code to get the greeting in that language.

## Supported Languages

Here's a list of currently supported languages and their respective codes:

- English (`en`): "Hello world"
- Greek (`el`): "Χαίρετε Κόσμε"
- French (`fr`): "Bonjour le monde"
- Hebrew (`he`): "שלום עולם"
- Urdu (`ur`): "ہیلو دنیا"
- Vietnamese (`vi`): "Xin chào Thế Giới"

If an unsupported language is specified, the program returns an error message.

