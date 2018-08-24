simplecalc2
===========

**simplecalc2** is a basic front end to the *goeval* simple expression
evaluation package with variables.

Sample session:

	paul@horus:~$ ./simplecalc2
	Enter simple mathematical expressions, 'q' to quit
	> 3+4
	  = 7
	> print 3+4
	7
	> 5 * foo
	error: unknown identifier
	> let foo = 4
	> 5 * foo
	  = 20
	> print 5 * foo
	20
	> let foo = pi
	> let r = 1    
	> print foo * r^2
	3.141593
	> let foo = true
	> true == foo
	  = true
	> false != true
	  = true
	> print r < 3
	true
	> print r > 0
	true
	> print r < 0
	false
	> let bar = r > 3
	> print bar
	false
	> q
    > let s = "foo"
    > print s + "bar"
    "foobar"
	paul@horus:~$ 
