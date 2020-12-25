# Command Line Password Generator

<br>

## Introduction

<p>This is a very basic password generator written in Go. It is intended to be integrated into PATH environmental variable and called from command line.

<br><br>

## Usage

### Issuing command
    
    pw -l <length> [-m <mode>] 

Length can be anywhere between `1` and `4096`.

<br>

### Available modes

    a    -    alphanumeric password
    n    -    numeric password

<br>

Mode is an optional flag. If it is not provided, alphanumeric password is generated.

The alphanumeric mode generates a random string containing characters from ASCII `33` to `125` (`!` to `}`). I decided not to include space (`32`) due to visibility issues. To include it change the following value:

```Go
const MIN_ALPHANUM_CODE = 32    // from 33 to 32
```

<br>

Also tilde (`126`) was not included, due to specific way of typing it (at least on my machine). To include it change the following value:

```Go
const MAX_ALPHANUM_CODE = 126   // from 125 to 126
```

<br>

### Example use

    pw -l 15 -m a     ->    FbUgA5=/>?
    pw -l 4 -m -n     ->    9229
    pw -l 8           ->    a"G;L-yo
</p>

<br><br>

## License

<p>Password Generator is available under MIT license.</p>
