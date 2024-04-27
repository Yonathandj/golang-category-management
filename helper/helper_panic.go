package helper

func HelperPanic(err error) {
	if err != nil {
		panic(err)
	}
}
