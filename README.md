# Command Line Password Generator
<br>

## Introduction
<br>

This is a very basic password generator written in Go. It is intended to be integrated into PATH environmental variable and called from command line. It utilizes Go's [cryptographically secure pseudorandom number generator](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator) for added safety. The software allows to generate an entirely random sequence of characters (numeric only as well as alphanumeric with symbols) or use a [diceware](https://en.wikipedia.org/wiki/Diceware) mode to generate word sequences. [EFF's word list](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) is used as the default source of words, but users may introduce their own resources.

<br>

## Usage
<br>

### Issuing command
<br>
    
```
pw -l <length> -m <mode> -f <wordlist_file>
```

<br>

### Flags
<br>

* `l` - password length can be anywhere between `1` and `4096`. Default is `6` - a convenient compromise between modes. You probably do not want lock your smartphone with 16-digit code or having to remember 16 words. However, when it comes to alphanumeric password used for serious applications, 16 characters may be quite reasonable.
* `m` - generation mode (see the next section)
* `f` - use custom list for diceware mod

<br>

### Available modes
<br>

* `a` - default, alphanumeric mode, which generates a password containing letters, numbers and symbols.
More specifically, the characters from ASCII `33` to `125` (`!` to `}`) are used. I have decided not to include space (`32`) due to visibility issues. Also, some websites I have encountered do not allow space in passwords. Along with space, tilde (`126`) was not included, due to specific way of typing it which may lead to mistakes in forms with obscured characters.

* `d` - diceware mode. Generates a password by randomly selecting words from the provided list. The users may then combine and modify those words as they see fit. An additional flag is accepted for this mode: `-f`, which lets the user pass the name of their own word list file. It must be a `.txt` file located within the `res` directory.

* `n` - numeric mode, used for generating codes containing only numbers.

<br>

### Example use
<br>

```
pw -l 15 -m a     ->    FbUgA5=/>?
pw -l 4 -m -n     ->    9229
pw -l 8           ->    a"G;L-yo
pw -m d -l 6      ->    active never disarm heroics pretzel antacid
pw -m d -f eff    ->    maimed humid bony awaken tracing
```

<br>

## Resources
<br>

* [CSPRNG](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator)
* [Diceware](https://en.wikipedia.org/wiki/Diceware)
* [Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases)

<br>

## License
<br>

This software is available under MIT license.
