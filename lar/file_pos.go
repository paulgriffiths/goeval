package lar

// FilePos represents a line number and character position in a file.
type FilePos struct {
    ch, line int
}

// clear clears all the fields in a FilePos
func (f *FilePos) clear() {
    f.ch, f.line = 0, 0
}

// setPos sets the FilePos's character position and line
func (f *FilePos) setPos(pos FilePos) {
	f.ch, f.line = pos.ch, pos.line
}

// inc increments the FilePos's character position
func (f *FilePos) inc() {
    f.ch++
}

// inc increments the FilePos's character position
func (f *FilePos) incLine() {
    f.ch = 0
    f.line++
}
