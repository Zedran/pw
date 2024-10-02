# pw

A command line password generator written in Go. It can generate a password or passphrase and allows basic output customization.

## Quickstart

```
go install github.com/Zedran/pw
```

### Flags

| Flag   | Description                                                              |
|:-------|:-------------------------------------------------------------------------|
| `-h`   | Display help                                                             |
| `-l`   | Specify the desired number of elements (default 6)                       |
| `-m`   | Specify generation mode (default `c`)                                    |
| `-n`   | Do not print `LF` at the end of the generated string                     |

### Mode

An argument for the `-m` flag.

| Mode   | Description                                                              |
|:------:|:-------------------------------------------------------------------------|
| `c`    | Generates password (characters)                                          |
| `w`    | Generates passphrase (words)                                             |

#### Password options

Password consists of individual, random characters.

| Option | Description                                                              |
|:------:|:-------------------------------------------------------------------------|
| `-e`   | Exclude the specified characters from charset (default none)             |
| `-i`   | Include the following character groups into the charset (default `Aans`) |

##### Character groups

* `A` - upper case letters
* `a` - lower case letters
* `n` - numbers
* `s` - symbols

#### Passphrase options

Passphrase consists of a number of random words.

| Option | Description                                                              |
|:-------|:-------------------------------------------------------------------------|
| `-f`   | Select word list file                                                    |
| `-s`   | Separate the generated words with the specified string (default space)   |

##### Custom word list file format requirements

* one, lower case word per line
* `LF` line endings
* no `LF` at the end of the file

### Examples

```bash
# Generate password of length 15
pw -m c -l 15

# Include upper case letters, lower case letters and numbers
pw -m c -i Aan

# Exclude space and '$' sign from the charset
pw -m c -i Aans -e \ \$

# Generate a passphrase of default length using the default word list
pw -m w

# Separate words with '_'
pw -m w -s _

# Separate words with '--'
pw -m w -s --
```

## Attributions

Refer to [NOTICE.md](NOTICE.md).

## License

This software is available under MIT license.
