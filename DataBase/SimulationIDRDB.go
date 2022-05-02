package DataBase

var IDRdb map[string]string //hash:id

func init() {
	IDRdb = make(map[string]string)
	IDRdb["af8baa3110d3226f27f92b6c625806a14213c9641a1c6aec3ce0122739c2e596"] = "nh679sg5x6wl"
}

func CalibrateIDR(hash string) (string, bool) {
	order, res := IDRdb[hash]
	return order, res
}
