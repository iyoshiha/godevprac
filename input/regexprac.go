package main

func regexprac() {
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})_(TEST)`)

	res := re.FindAllStringSubmatch(str, -1)
}