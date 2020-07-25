# `rangereader` -- a Reader to read from a range of bytes

`RangeReader` wraps another `io.Reader` and will only read data in the range you specify.

Here's the first test, which illustrates the concept pretty well:

```golang
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
```

It was initially developed to ease implementation of HTTP range requests.