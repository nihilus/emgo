package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"unicode"
	"unicode/utf8"
)

func warn(info ...interface{}) {
	fmt.Fprintln(os.Stderr, info...)
}

func die(info ...interface{}) {
	fmt.Fprintln(os.Stderr, info...)
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		die(err)
	}
}

func doxy(s, tag string) string {
	n := strings.Index(s, tag)
	if n < 0 {
		return ""
	}
	s = s[n+len(tag):]
	if len(s) == 0 || s[0] != ' ' && s[0] != '\t' {
		return ""
	}
	return strings.TrimSpace(s)
}

func split(s string) (first, rest string) {
	s = strings.TrimSpace(s)
	i := strings.IndexFunc(s, unicode.IsSpace)
	if i < 0 {
		return s, ""
	}
	return s[:i], strings.TrimSpace(s[i+1:])
}

// cwc wraps io.WriteCloser. It terminates program on any error.
type cwc struct {
	c io.WriteCloser
}

func create(path string) cwc {
	f, err := os.Create(path)
	checkErr(err)
	return cwc{f}
}

func (w cwc) Write(b []byte) (int, error) {
	n, err := w.c.Write(b)
	checkErr(err)
	return n, nil
}

func (w cwc) WriteString(s string) (int, error) {
	n, err := io.WriteString(w.c, s)
	checkErr(err)
	return n, nil
}

func (w cwc) Close() error {
	var name string
	if f, ok := w.c.(*os.File); ok {
		name = f.Name()
	}
	checkErr(w.c.Close())
	if name != "" {
		name, err := filepath.Abs(name)
		checkErr(err)
		gofmt := exec.Command("gofmt", "-w", name)
		gofmt.Stdout = os.Stdout
		gofmt.Stderr = os.Stderr
		checkErr(gofmt.Run())
	}
	return nil
}

func (w cwc) donotedit() {
	io.WriteString(
		w,
		"\n// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.\n\n",
	)
}

func mkdir(path string) {
	err := os.Mkdir(path, 0755)
	if e, ok := err.(*os.PathError); !ok ||
		e.Err != os.ErrExist && e.Err != syscall.EEXIST {
		checkErr(err)
	}
}

func chdir(path string) {
	checkErr(os.Chdir(path))
}

type scanner struct {
	*bufio.Scanner
	Name string
	n    int
}

func newScanner(r io.Reader, name string) *scanner {
	return &scanner{
		Scanner: bufio.NewScanner(r),
		Name:    name,
	}
}

func (s *scanner) Scan() bool {
	s.n++
	return s.Scanner.Scan()
}

func (s *scanner) Die(v ...interface{}) {
	v = append(
		[]interface{}{fmt.Sprintf("%s:%d", s.Name, s.n)},
		v...,
	)
	die(v)
}

func upperFirst(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError || unicode.IsUpper(r) {
		return s
	}
	s = s[n:]
	buf := make([]byte, 4+len(s))
	n = utf8.EncodeRune(buf, unicode.ToUpper(r))
	n += copy(buf[n:], s)
	return string(buf[:n])
}

func ident(s string) string {
	r, _ := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError || unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return s
	}
	return "V" + s
}

func trailingZeros32(u uint32) uint {
	n := uint(32)
	x := u << 16
	if x != 0 {
		n -= 16
		u = x
	}
	x = u << 8
	if x != 0 {
		n -= 8
		u = x
	}
	x = u << 4
	if x != 0 {
		n -= 4
		u = x
	}
	x = u << 2
	if x != 0 {
		n -= 2
		u = x
	}
	x = u << 1
	if x != 0 {
		n -= 1
		u = x
	}
	return n - uint(u>>31)
}
