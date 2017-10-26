// Package secret implements the secret-handshake exercise
package secret

const testVersion = 2

// Handshake generates nerdy Handshake
func Handshake(n uint) (out []string) {
	var reverse bool

	if n&1 == 1 {
		out = append(out, "wink")
	}
	if n&2 == 2 {
		out = append(out, "double blink")
	}
	if n&4 == 4 {
		out = append(out, "close your eyes")
	}
	if n&8 == 8 {
		out = append(out, "jump")
	}
	if n&16 == 16 {
		reverse = true
	}
	if reverse {
		tmp := make([]string, len(out))
		pos := len(out) - 1
		for i, v := range out {
			tmp[pos-i] = v
		}
		out = tmp
	}
	return out
}
