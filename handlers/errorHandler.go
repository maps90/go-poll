package handlers

func Error(err error) {
	if err != nil {
		panic(err)
		return
	}
}
