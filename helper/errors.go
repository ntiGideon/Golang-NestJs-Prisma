package helper

func PanicAllErrors(err error) {
	if err != nil {
		panic(err.Error())
	}
}
