package api

import (
	"testing"
)

// Test the origin validation
func TestOriginValidator(t *testing.T) {
	testcases := []struct {
		origin string
		allow  bool
	}{
		// `null` should be denied
		{"null", false},
		// HTTPS for detahard.io should be allowed
		{"https://detahard.io", true},
		{"https://foo.detahard.io", true},
		{"https://bar.foo.detahard.io", true},
		// but HTTP for detahard.io should be denied
		{"http://detahard.io", false},
		{"http://foo.detahard.io", false},
		{"http://bar.foo.detahard.io", false},
		// Fakes should be denied
		{"https://fakedetahard.io", false},
		{"https://foo.fakedetahard.io", false},
		{"https://foo.detahard.ioo", false},
		{"http://fakedetahard.io", false},
		{"http://foo.fakedetahard.io", false},
		{"http://foo.detahard.ioo", false},
		// detahard onion should be allowed
		{"http://detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"http://foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"http://bar.foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://bar.foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		// Fake detahard onions should be denied
		{"http://detahardiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"https://detahardiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"http://detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqbd.onion", false},
		{"https://detahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqbd.onion", false},
		{"http://detahardiowpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://detahardiowpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"https://bar.foo.detahardiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"http://fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://foo.fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://foo.fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://bar.foo.fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://bar.foo.fakedetahardiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		// Localhost 8xxx and 5xxx should be allowed for local development
		{"https://localhost:8000", true},
		{"http://localhost:8000", true},
		{"http://localhost:8999", true},
		{"https://localhost:5000", true},
		{"http://localhost:5000", true},
		{"http://localhost:5999", true},
		// SatoshiLabs dev servers should be allowed
		{"https://sldev.cz", true},
		{"https://foo.sldev.cz", true},
		{"https://bar.foo.sldev.cz", true},
		// Fake SatoshiLabs dev servers should be denied
		{"https://fakesldev.cz", false},
		{"https://foo.fakesldev.cz", false},
		{"https://foo.sldev.czz", false},
		{"http://foo.detahard.sldev.cz", false},
		// Other ports should be denied
		{"http://localhost", false},
		{"http://localhost:1234", false},
	}
	validator := corsValidator()
	for _, tc := range testcases {
		allow := validator(tc.origin)
		if allow != tc.allow {
			t.Errorf("Origin %q: expected %v, got %v", tc.origin, tc.allow, allow)
		}
	}
}
