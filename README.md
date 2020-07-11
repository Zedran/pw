<h1>Command Line Password Generator</h1>

<p>This is a very basic password generator written in Go. It is intended to be integrated into PATH environmental variable and called from command line.

<br><br>

<h3>Issuing command</h3>
    
    pw <mode> <length> 

Length can be anywhere between 1 and 4096.

<br>

<h3>Available modes</h3>

    a    -    alphanumeric password
    n    -    numeric password

<br>

The alphanumeric mode generates a random string containing characters from ASCII 33 to 125 ('!' to '}'). I decided not to include space (32) due to visibility issues. To include it change the following value:

```Go
const MIN_ALPHANUM_CODE = 32    // from 33 to 32
```

<br>

Also tilde (126) was not included, due to specific way of typing it (at least on my machine). To include it change the following value:

```Go
const MAX_ALPHANUM_CODE = 126   // from 125 to 126
```

<br>

<h3>Example use</h3>

    pw a 10    ->    FbUgA5=/>?
    pw n 4     ->    9229
</p>

<br><br>

<h2>License</h2>

<p>Password Generator is available under MIT license.</p>
