package buftermio

import (
	"fmt"
    "strings"
)

func (b *Buffer) backspace() {
	if b.len > 1 {
		b.removeSlice(2)
		left(b.cursor)
		fmt.Print(string(b.buf), " ")
		left(b.len - b.cursor + 1)
	} else {
		b.removeSlice(1)
		bell()
	}
}

func (b *Buffer) upIndex() {
	b.removeSlice(3)
	if b.index > 0 {
		b.addIndex(b.index)
		b.clear()
		b.index--
		if b.index > 0 && b.isIndexEmpty(b.index) {
			b.index--
		}
		b.newLine(b.cache[b.index]...)
		fmt.Print(string(b.buf))
	} else {
		bell()
	}
}

func (b *Buffer) downIndex() {
	b.removeSlice(3)
	if b.index < len(b.cache)-1 {
		b.addIndex(b.index)
		b.clear()
		b.index++
		b.newLine(b.cache[b.index]...)
		fmt.Print(string(b.buf))
	} else {
		bell()
	}
}

func (b *Buffer) cursorRight() {
	b.removeSlice(3)
	if b.cursor < b.len {
		fmt.Print(string(rightArrow))
		b.cursor++
	} else {
		bell()
	}
}

func (b *Buffer) cursorLeft() {
	b.removeSlice(3)
	if b.cursor > 0 {
		left(1)
		b.cursor--
	} else {
		bell()
	}
}

func (b *Buffer) cursorSOL() {
    b.removeSlice(2)
    if b.cursor > 0 {
        left(b.cursor)
        b.cursor = 0
    } else {
        bell()
    }
}

func (b *Buffer) cursorEOL() {
    b.removeSlice(2)
    if b.cursor < b.len {
        fmt.Print(strings.Repeat(string(rightArrow), b.len - b.cursor))
        b.cursor = b.len
    } else {
        bell()
    }
}

func (b *Buffer) fourSpaces() {
	b.removeSlice(1)
	b.insert(fourSpaces)
	left(b.cursor - 4)
	fmt.Print(string(b.buf))
	left(b.len - b.cursor)
}

func (b *Buffer) enter() string {
	b.removeSlice(1)
	if b.index < len(b.cache) {
		newCache := make([][]byte, 0)
		newCache = append(newCache, b.cache[:b.index]...)
		newCache = append(newCache, b.cache[b.index+1:]...)
		b.cache = newCache
	}
	if b.len > 0 {
		b.addIndex(len(b.cache))
	}
	b.index = len(b.cache)
	copy := make([]byte, 0)
	copy = append(copy, b.buf...)
	b.newLine()
	return string(copy)
}
