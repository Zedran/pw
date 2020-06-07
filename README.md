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

<h3>Example use</h3>

    pw a 10    ->    mO6I|_r0u(
    pw n 4     ->    0885
</p>

<br><br>

<h2>License</h2>

<p>Password Generator is available under MIT license.</p>
