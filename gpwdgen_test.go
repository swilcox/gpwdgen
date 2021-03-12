package main

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected: %v (type %v)  Got: %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	// back to normal state
	w.Close()
	os.Stdout = orig
	out := <-outC
	return out
}

func TestGeneratePassPhrase(t *testing.T) {
	// random seed
	rand.Seed(42)

	passphrase := GeneratePassPhrase("", 4, "-", 1, 9)
	expect(t, "Nullah-gatekeeper-glint-hansom4", passphrase)

	passphrase = GeneratePassPhrase("", 1, "", 0, 0)
	expect(t, "memory", passphrase)

	passphrase = GeneratePassPhrase("", 2, ".", 2, 99)
	expect(t, "holder.Partial38", passphrase)

	passphrase = GeneratePassPhrase("test_words.txt", 2, "_", 0, 0)
	expect(t, "does_fine", passphrase)
}

func TestMain(t *testing.T) {
	passphrase := captureOutput(main)
	if len(passphrase) > 100 || len(passphrase) < 3 {
		t.Error("main didn't appear to run correctly!")
	}
}
