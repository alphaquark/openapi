package upbit

func Serve() (price float32, err error) {
	upbit := initClient()
	price, err = upbit.VAPrice()

	return
}
