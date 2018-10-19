package stringutils

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (fullname string, length int) {
	fullname = f + " " + l
	length = len(fullname)
	//return fullname, length
	//naked return since specified the return from parameter
	return
}
