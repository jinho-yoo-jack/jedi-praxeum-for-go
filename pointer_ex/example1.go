package pointer_ex

type Data struct {
	value int
	data  [200]int
}

func ChangeByVale(arg Data) {
	arg.value = 999
	arg.data[0] = 100
}

func ChangeByReference(arg *Data) {
	arg.value = 999
	arg.data[0] = 100
}
