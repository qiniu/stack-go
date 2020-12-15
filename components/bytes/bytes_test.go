package bytes

import (
	"io"
	"testing"
)

// ---------------------------------------------------

func TestBuffer(t *testing.T) {

	b := NewBuffer()
	n, err := b.WriteStringAt("Hello", 4)
	if n != 5 || err != nil {
		t.Fatal("WriteStringAt failed:", n, err)
	}
	if b.Len() != 9 {
		t.Fatal("Buffer.Len invalid (9 is required):", b.Len())
	}

	buf := make([]byte, 10)
	n, err = b.ReadAt(buf, 50)
	if n != 0 || err != io.EOF {
		t.Fatal("ReadAt failed:", n, err)
	}

	n, err = b.ReadAt(buf, 6)
	if n != 3 || err != io.EOF || string(buf[:n]) != "llo" {
		t.Fatal("ReadAt failed:", n, err, string(buf[:n]))
	}

	n, err = b.WriteAt([]byte("Hi h"), 1)
	if n != 4 || err != nil {
		t.Fatal("WriteAt failed:", n, err)
	}
	if b.Len() != 9 {
		t.Fatal("Buffer.Len invalid (9 is required):", b.Len())
	}

	n, err = b.ReadAt(buf, 0)
	if n != 9 || err != io.EOF || string(buf[:n]) != "\x00Hi hello" {
		t.Fatal("ReadAt failed:", n, err)
	}

	n, err = b.WriteStringAt("LO world!", 7)
	if n != 9 || err != nil {
		t.Fatal("WriteStringAt failed:", n, err)
	}
	if b.Len() != 16 {
		t.Fatal("Buffer.Len invalid (16 is required):", b.Len())
	}

	buf = make([]byte, 17)
	n, err = b.ReadAt(buf, 0)
	if n != 16 || err != io.EOF || string(buf[:n]) != "\x00Hi helLO world!" {
		t.Fatal("ReadAt failed:", n, err, string(buf[:n]))
	}
}

// ---------------------------------------------------

func TestReader(t *testing.T) {
	r := NewReader([]byte("hello world"))
	if r.Len() != 11 {
		t.Fatal("len invalid (11 is required)", r.Len())
	}

	if string(r.Bytes()) != "hello world" {
		t.Fatal("bytes invalid ('hello world' is required)", string(r.Bytes()))
	}

	_, err := r.Seek(-10, 1)
	if err == nil {
		t.Fatal("err invalid (err is required)", err)
	}

	ret, err := r.Seek(9, 0)
	if err != nil {
		t.Fatal("err invalid", err)
	}
	if ret != 9 {
		t.Fatal("ret invalid, (9 is required)", ret)
	}

	r.SeekToBegin()

	b := make([]byte, 100)
	cnt, err := r.Read(b)
	if err != nil {
		t.Fatal("err invalid", err)
	}
	if cnt != 11 {
		t.Fatal("read cnt invalid (11 is required)", cnt)
	}

	_, err = r.Read(b)
	if err == nil {
		t.Fatal("err invalid (EOF is required)", err)
	}

	r.Close()
}

func TestWriter(t *testing.T) {
	w := NewWriter(make([]byte, 100))
	w.Write([]byte("hello world"))
	if w.Len() != 11 {
		t.Fatal("len invalid (11 is required)", w.Len())
	}

	if string(w.Bytes()) != "hello world" {
		t.Fatal("bytes invalid ('hello world' is required)", string(w.Bytes()))
	}

	w.Reset()
}
