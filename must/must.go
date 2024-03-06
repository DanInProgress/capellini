package must

// Succeed checks if err != nil, otherwise panic.
// for use in scenarios where errors cannot occur,
// but where traditional error handling is not possilbe.
func Succeed(err error) {
	if err != nil {
		panic(err)
	}
}

// Return if err == nil, return V, otherwise panic.
// for use in scenarios where errors cannot occur,
// but where traditional error handling is not possilbe.
func Return[V any](v V, err error) V {
	if err != nil {
		panic(err)
	}
	return v
}
