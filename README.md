# Command Line Password Generator
<br>

## Introduction
<br>

This is a very basic password generator written in Go. It is intended to be integrated into PATH environmental variable and called from command line. It utilizes Go's [cryptographically secure pseudorandom number generator](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator) for good quality randomness. The software generates a random sequence of characters (numeric only as well as alphanumeric with symbols) or uses a [diceware](https://en.wikipedia.org/wiki/Diceware) mode to generate word sequences. The default source of words is [EFF's word list](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases), but users may introduce their own resources.

<br>

## Installation

You can download and unpack the zipped release into the directory of your choice. I would, however, encourage you to clone this repo and compile the code yourself. I have provided build scripts for Windows and Linux. 

Once you have your program directory set up, the next step is to add it to PATH. This will allow you to get a password by calling `pw` in the command line.

## Usage
<br>

### Command syntax
<br>
    
```
pw -l <length> -m <mode> -f <wordlist_file>
```

<br>

### Flags
<br>

* `l` - password length, can be anywhere between `1` and `4096`. The default is `6` - a convenient compromise between modes. You probably do not want to lock your smartphone with 16-digit code or have to remember 16 words. However, when it comes to alphanumeric password used for serious applications, 16 characters may be quite reasonable.
* `m` - generation mode [(see the next section)](#modes)
* `f` - use custom list for diceware mod

<br>

### Modes
<br>

* `a` - default, alphanumeric mode, which generates a password containing letters, numbers and symbols.
More specifically, the characters from ASCII `33` to `125` (`!` to `}`) are used. I have decided not to include space (`32`) due to visibility issues. Also, some websites I have encountered do not allow space in passwords. Along with space, tilde (`126`) was not included, due to specific way of typing it which may lead to mistakes in forms with hidden characters.

* `d` - diceware mode. Generates a password by randomly selecting words from the provided list. The users may then combine and modify those words as they see fit. An additional flag is accepted for this mode: `-f`, which lets the user pass the name of their own word list file. It must be a `.txt` file located within the `res` directory.

* `n` - numeric mode, used for generating codes containing only numbers.

<br>

### Example
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

[Detailed information](./doc/RES_INFO.md)

<br>

## License
<br>

This software is available under MIT license.
