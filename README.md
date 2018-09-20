goeval
======

*goeval* is, or will become, a collection of Golang packages and applications
concerning lexing, parsing and expression evaluation.

The principle purpose of this project is for me to practise writing in Go.

Among other things, the packages implement the following:

* Automata

    * Data structures to simulate arbitrary deterministic and nondeterministic
    finite automata
    * Conversion of nondeterministic to deterministic finite automata
    * Construction of deterministic and nondeterministic finite automata from
    regular expressions

* Context-free grammars

    * A parser to build an arbitrary context-free grammar from a representation in a text file
    * Functionality to examine a context-free grammar and report:

        * If the grammar is left-recursive
        * Lists of cyclic, nullable, unreachable and unproductive nonterminals
        * First and follow sets for all nonterminals

    * Parsers for context-free grammars

        * Recursive descent parser

* A library and interactive interface for the evaluation of arithmetical,
boolean and string expressions using dynamically-typed variables
