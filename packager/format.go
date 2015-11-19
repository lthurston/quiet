package packager


func columnize(string string) out string {
	out := ""
	for i := 1.00 ; i <= math.Ceil(float64(len(string)) / 80.0); i++ {
		if (i * 80) <= float64(len(base64d)) {
			out = out + base64d[int((i - 1) * 80):int(i * 80)]
		} else {
			out = out + base64d[int((i - 1) * 80):]
		}
	}
	return out
}

func decolumnize(column string) out string {
	// do the opposite, son!
	return "a string"
}