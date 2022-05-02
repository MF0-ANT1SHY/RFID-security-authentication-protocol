package DataBase

var IDTdb map[string]string // hash:id

func init() {
	IDTdb = make(map[string]string)
	IDTdb["1737eed9489ea3b54a5550e858e7a6d2d41054c3d95d594c06af395bc463bc0f"] = "9p9x1ydp7uue"
}

func CalibrateIDT(hash string) (string, bool) {
	order, res := IDTdb[hash]
	return order, res
}
