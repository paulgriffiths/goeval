simplecalc
==========

**simplecalc** is a basic front end to the *goeval* simple expression
evaluation package.

Sample session:

	paul@horus:~$ ./simplecalc
	Enter simple mathematical expressions, 'q' to quit
	> 1 + 2
	  = 3
	> 3 * 4
	  = 12
	> (5 + 6) * 7
	  = 77
	> 2 ^ 10
	  = 1024
	> ((ln(e^2) * log(1000)) * sin(90)) ^ 2 + sqrt(36)
	  = 42
	> 2 * pi
	  = 6.283185307179586
	> 4 / 0
	error: divide by zero
	> 3 +
	error: missing factor
	> (8 * 7
	error: unbalanced parentheses
	> 4 ^ 3 with some extraneous text
	error: trailing tokens
	> q
	paul@horus:~$ 
