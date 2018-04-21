package cipher

import (
	"testing"
)

// type for testing cipher encoding alone, without requiring any text prep.
type prepped struct {
	pt string // prepped text == decoded plain text
	ct string // cipher text
}

// +3? -3?  Positions vary.  Your code just has to pass the tests.
var caesarPrepped = []prepped{
	{"iamapandabear", "ldpdsdqgdehdu"},
	{"programmingisawesome", "surjudpplqjlvdzhvrph"},
	{"todayisholiday", "wrgdblvkrolgdb"},
	{"venividivici", "yhqlylglylfl"},
}

func TestCaesarPrepped(t *testing.T) {
	c := NewCaesar()
	for _, test := range caesarPrepped {
		if enc := c.Encode(test.pt); enc != test.ct {
			t.Fatalf("Caesar Encode(%q) = %q, want %q.", test.pt, enc, test.ct)
		}
	}
}

// type for testing implementations of the Cipher interface
type cipherTest struct {
	source string // source text, any UTF-8
	cipher string // cipher text, result of Encode(st)
	plain  string // decoded plain text, result of Decode(ct)
}

var caesarTests = []cipherTest{
	{"Go, go, gophers", "jrjrjrskhuv", "gogogophers"},
	{"I am a panda bear.", "ldpdsdqgdehdu", "iamapandabear"},
	{"Programming is AWESOME!", "surjudpplqjlvdzhvrph", "programmingisawesome"},
	{"today is holiday", "wrgdblvkrolgdb", "todayisholiday"},
	{"Twas the night before Christmas",
		"wzdvwkhqljkwehiruhfkulvwpdv",
		"twasthenightbeforechristmas"},
	{"venividivici", "yhqlylglylfl", "venividivici"},
	{" -- @#!", "", ""},
	{"", "", ""},
}

func TestCaesar(t *testing.T) {
	testCipher("Caesar", NewCaesar(), caesarTests, t)
}

func testCipher(name string, c Cipher, tests []cipherTest, t *testing.T) {
	for _, test := range tests {
		if enc := c.Encode(test.source); enc != test.cipher {
			t.Fatalf("%s Encode(%q) = %q, want %q.",
				name, test.plain, enc, test.cipher)
		}
		if dec := c.Decode(test.cipher); dec != test.plain {
			t.Fatalf("%s Decode(%q) = %q, want %q.",
				name, test.cipher, dec, test.plain)
		}
	}
}

// Benchmark combined time to run all tests.
// Note other ciphers test different data; times cannot be compared.
func BenchmarkEncodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Encode(test.source)
		}
	}
}

func BenchmarkDecodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Decode(test.cipher)
		}
	}
}
