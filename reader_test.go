package rangereader

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNormalBehaviour(t *testing.T) {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	rr, err := New(reader, 16, 19)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(rr)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "fox" {
		t.Fatalf("Did not read back what we expected. Expected 'fox', got '%s'", string(b))
	}
}

func TestNegativeStart(t *testing.T) {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	_, err := New(reader, -1, 10)
	if err == nil {
		t.Fatalf("Expected an error when creating the Reader")
	}
}

func TestNegativeEnd(t *testing.T) {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	_, err := New(reader, 3, -6)
	if err == nil {
		t.Fatalf("Expected an error when creating the Reader")
	}
}

func TestStartBiggerThanEnd(t *testing.T) {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	_, err := New(reader, 10, 3)
	if err == nil {
		t.Fatalf("Expected an error when creating the Reader")
	}
}

func TestNilReader(t *testing.T) {
	_, err := New(nil, 1, 2)
	if err == nil {
		t.Fatalf("Expected an error when creating the Reader")
	}
}

func TestReadPastBounds(t *testing.T) {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	rr, err := New(reader, 16, 19)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ioutil.ReadAll(rr)
	if err != nil {
		t.Fatal(err)
	}

	p := make([]byte, 10)
	n, err := rr.Read(p)

	if err != io.EOF {
		t.Fatalf("Did not get 'EOF' after reading past the bounds")
	}

	if n != 0 {
		t.Fatalf("Read some bytes although we're out of bounds")
	}
}
