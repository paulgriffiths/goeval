

# lar
`import "github.com/paulgriffiths/goeval/lar"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package lar provides a single character lookahead reader.

When performing lexical analysis, it often greatly simplifies matters to
have available one character of lookahead, i.e. to be able to peek at what
the next character to be read would be without actually reading it. With
this ability you can construct a series of tests to check if the next
character matches the start of a pattern, but only read and consume that
character from the input if it does, in fact, match.

The single character lookahead reader implemented by this package uses that
functionality to provide a set of matching functions which will extract
a character or set of characters from the input if and only if they match
a specified category.




## <a name="pkg-index">Index</a>
* [type LookaheadReader](#LookaheadReader)
  * [func NewLookaheadReader(reader io.Reader) (LookaheadReader, error)](#NewLookaheadReader)
  * [func (r LookaheadReader) EndOfInput() bool](#LookaheadReader.EndOfInput)
  * [func (r *LookaheadReader) MatchDigit() bool](#LookaheadReader.MatchDigit)
  * [func (r *LookaheadReader) MatchDigits() bool](#LookaheadReader.MatchDigits)
  * [func (r *LookaheadReader) MatchLetter() bool](#LookaheadReader.MatchLetter)
  * [func (r *LookaheadReader) MatchLetters() bool](#LookaheadReader.MatchLetters)
  * [func (r *LookaheadReader) MatchOneOf(vals ...byte) bool](#LookaheadReader.MatchOneOf)
  * [func (r *LookaheadReader) MatchSpace() bool](#LookaheadReader.MatchSpace)
  * [func (r *LookaheadReader) MatchSpaces() bool](#LookaheadReader.MatchSpaces)
  * [func (r *LookaheadReader) Next() (byte, error)](#LookaheadReader.Next)
  * [func (r LookaheadReader) Result() []byte](#LookaheadReader.Result)

#### <a name="pkg-examples">Examples</a>
* [Package](#example_)

#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/paulgriffiths/goeval/lar/doc.go) [lookahead_reader.go](/src/github.com/paulgriffiths/goeval/lar/lookahead_reader.go) 






## <a name="LookaheadReader">type</a> [LookaheadReader](/src/target/lookahead_reader.go?s=116:195#L10)
``` go
type LookaheadReader struct {
    // contains filtered or unexported fields
}

```
LookaheadReader implements a single character lookahead reader.







### <a name="NewLookaheadReader">func</a> [NewLookaheadReader](/src/target/lookahead_reader.go?s=284:350#L18)
``` go
func NewLookaheadReader(reader io.Reader) (LookaheadReader, error)
```
NewLookaheadReader returns a single character lookahead reader from
an io.Reader





### <a name="LookaheadReader.EndOfInput">func</a> (LookaheadReader) [EndOfInput](/src/target/lookahead_reader.go?s=3351:3393#L108)
``` go
func (r LookaheadReader) EndOfInput() bool
```
EndOfInput returns true if end of input has been reached, otherwise false.




### <a name="LookaheadReader.MatchDigit">func</a> (\*LookaheadReader) [MatchDigit](/src/target/lookahead_reader.go?s=2131:2174#L77)
``` go
func (r *LookaheadReader) MatchDigit() bool
```
MatchDigit returns true if the next character to be read is a digit
and stores that character in the result, otherwise it returns false
and clears the result.




### <a name="LookaheadReader.MatchDigits">func</a> (\*LookaheadReader) [MatchDigits](/src/target/lookahead_reader.go?s=3050:3094#L98)
``` go
func (r *LookaheadReader) MatchDigits() bool
```
MatchDigits returns true if the next character to be read is a digit
and stores that and all immediately following digit characters in
the result, otherwise it returns false and clears the result.




### <a name="LookaheadReader.MatchLetter">func</a> (\*LookaheadReader) [MatchLetter](/src/target/lookahead_reader.go?s=1602:1646#L63)
``` go
func (r *LookaheadReader) MatchLetter() bool
```
MatchLetter returns true if the next character to be read is a letter
and stores that character in the result, otherwise it returns false
and clears the result.




### <a name="LookaheadReader.MatchLetters">func</a> (\*LookaheadReader) [MatchLetters](/src/target/lookahead_reader.go?s=2434:2479#L84)
``` go
func (r *LookaheadReader) MatchLetters() bool
```
MatchLetters returns true if the next character to be read is a letter
and stores that and all immediately following letter characters in
the result, otherwise it returns false and clears the result.




### <a name="LookaheadReader.MatchOneOf">func</a> (\*LookaheadReader) [MatchOneOf](/src/target/lookahead_reader.go?s=1218:1273#L48)
``` go
func (r *LookaheadReader) MatchOneOf(vals ...byte) bool
```
MatchOneOf returns true if the next character to be read is among
the characters passed to the function and stores that character in
the result, otherwise it returns false and clears the result.




### <a name="LookaheadReader.MatchSpace">func</a> (\*LookaheadReader) [MatchSpace](/src/target/lookahead_reader.go?s=1869:1912#L70)
``` go
func (r *LookaheadReader) MatchSpace() bool
```
MatchSpace returns true if the next character to be read is whitespace
and stores that character in the result, otherwise it returns false
and clears the result.




### <a name="LookaheadReader.MatchSpaces">func</a> (\*LookaheadReader) [MatchSpaces](/src/target/lookahead_reader.go?s=2747:2791#L91)
``` go
func (r *LookaheadReader) MatchSpaces() bool
```
MatchSpaces returns true if the next character to be read is whitespace
and stores that and all immediately following whitespace characters in
the result, otherwise it returns false and clears the result.




### <a name="LookaheadReader.Next">func</a> (\*LookaheadReader) [Next](/src/target/lookahead_reader.go?s=755:801#L29)
``` go
func (r *LookaheadReader) Next() (byte, error)
```
Next returns the next character from a lookahead reader.
If there are no more characters, the function returns 0 and io.EOF.
On any other error, the function returns 0 and that error.




### <a name="LookaheadReader.Result">func</a> (LookaheadReader) [Result](/src/target/lookahead_reader.go?s=3210:3250#L103)
``` go
func (r LookaheadReader) Result() []byte
```
Result returns the result of the most recent matching test.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
