simplecalc
==========

**simplecalc** is a basic front end to the *goeval* simple expression
evaluation package.

Sample session:

	paul@horus:$ ./simplecalc
	Enter simple mathematical expressions, 'q' to quit
	> 1+2
	  = 3
	> 5-3
	  = 2
	> 2 * 7
	  = 14
	> 4^5
	  = 1024
	> (4+5) * 6
	  = 54
	> 4/0
	error: divide by zero
	> 3+
	error: missing factor
	> -1 ^ 0.5
	error: domain error
	> (8 * 7
	error: unbalanced parentheses
	> q
	paul@horus:$
