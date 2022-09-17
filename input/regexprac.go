package main

func regexprac() {
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})_(TEST)`)

	// -1 = 取得するマッチ数:全部のマッチ箇所を取得する-1 
	res := re.FindAllStringSubmatch(str, -1)
}