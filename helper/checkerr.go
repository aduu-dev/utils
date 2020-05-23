package helper

// CheckErr panics with the given error if it is not nil.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
