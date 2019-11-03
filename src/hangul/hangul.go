package hangul

var (
	start = rune(44032)  // "가" 의 unicode
	end = rune(55204)    // "힣" 의 unicode
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false

	for _, r := range s {
		if start <= r && r < end {
			index :=int(r-start)
			result = index % numEnds != 0
		}
	}
	return result
}

