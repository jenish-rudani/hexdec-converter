# Hexadecimal-Decimal Converter

This Go module provides command-line utilities for converting between hexadecimal and decimal numbers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Examples](#examples)
- [Building](#building)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Development](#development)
  - [Makefile Commands](#makefile-commands)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/jenish-rudani/hexdec-converter.git
cd hexdec-converter
```

2. Build the binaries:

```bash
make
```

3. Install the utilities to your system PATH:

```bash
make install
```

Note: The `make install` command may require sudo privileges on Linux and macOS. On Windows, you may need to run your terminal as Administrator.

## Usage

After installation, you can use the utilities from anywhere in your terminal:

Convert hexadecimal to decimal:

```bash
h2d <hexadecimal number>
```

Convert decimal to hexadecimal:

```bash
d2h <decimal number>
```

### Examples

Here are some examples of how to use the utilities:

1. Hexadecimal to Decimal:

   ```bash
   $ h2d A
   10

   $ h2d FF
   255

   $ h2d 100
   256
   ```

2. Decimal to Hexadecimal:

   ```bash
   $ d2h 10
   A

   $ d2h 255
   FF

   $ d2h 256
   100
   ```

## Building

To build binaries for multiple platforms:

```bash
make
```

This will create binaries in the `binary` directory, organized by platform.

## Testing

To run the unit and integration tests:

```bash
make test
```

This will run tests for the converter package and integration tests for the command-line tools.

## Project Structure

```
hexdec-converter/
├── cmd/
│   ├── h2d/
│   │   └── main.go
│   └── d2h/
│       └── main.go
├── pkg/
│   └── converter/
│       ├── converter.go
│       └── converter_test.go
├── test/
│   └── integration_test.go
├── Makefile
├── go.mod
└── README.md
```

- `cmd/`: Contains the main applications for h2d and d2h.
- `pkg/`: Contains the core logic for conversion in the converter package.
- `test/`: Contains integration tests for the command-line tools.

## Development

### Makefile Commands

The project includes a Makefile with the following commands:

- `make`: Build the binaries for all supported platforms
- `make test`: Run unit and integration tests
- `make clean`: Remove all generated binaries
- `make info`: Display usage information and examples
- `make install`: Install the utilities to your system PATH

To view the usage information and examples in the terminal, run:

```bash
make info
```

This will display:

```
Hexadecimal-Decimal Converter
Usage:
  h2d <hexadecimal number>  - Convert hex to decimal
  d2h <decimal number>      - Convert decimal to hex

Examples:
  h2d FF  -> 255
  d2h 255 -> FF

Build:   make
Test:    make test
Clean:   make clean
Install: make install
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
```