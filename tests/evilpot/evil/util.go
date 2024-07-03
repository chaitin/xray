package evil

var SepFunc = func(b byte) bool {
	if 48 <= b && b <= 57 {
		return false
	}
	if b >= 128 {
		return false
	}
	if b >= 65 && b <= 90 {
		return false
	}
	if b >= 97 && b <= 122 {
		return false
	}
	return true
}

func Split(data []byte, sep func(b byte) bool, handler func([]byte) bool) {
	if sep == nil {
		sep = SepFunc
	}
	visible := 0
	invisible := 0
	for i := 0; i < len(data); i++ {
		if !sep(data[i]) {
			continue
		}
		invisible = i
		if invisible == visible {
			visible++
			continue
		}
		b := data[visible:invisible]
		if !handler(b) {
			return
		}
		visible = invisible + 1
	}
	if visible < len(data) {
		handler(data[visible:])
	}
	return
}
