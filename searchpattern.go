package godtype

type PatternFinder struct {
	pattern	[]rune
	lps 	[]int
}

// use KMP algorithm
func NewPatternFinder(p string) *PatternFinder {
	return &PatternFinder{
		pattern:	[]rune(p),
		lps:		calculateLPS(p),
	}
}

func calculateLPS(p string) []int {
	// treat string as unicode
	r := []rune(p)

	n := len(r)
	lps := make([]int, n)

	// proper prefix of a single rune is ""
	lps[0] = 0

	/* 
	  pointer: 
		lg: prefix current
		i:  suffix current

	  lg: represent also the len of previous longest proper prefix (which is also suffix),
	      so the value need to be incremented if match with the suffix current, before write to the table
	*/
	lg, i := 0, 1

	for i < n {
		// the rune after the previous longest proper prefix equals to the rune at the current end (suffix)
		if r[i] == r[lg] {
			// in this case increment and register the longest proper prefix at i
			// lg in the same time point to the next rune, for next round's comparison
			lg++
			lps[i] = lg
			i++
		}else{  // r[i] != r[lg]: don't match
			// longest proper prefix is zero, fill the table and check the next rune
			if lg == 0 {
				lps[i] = 0
				i++
			}else{
			// longest proper prefix is not zero, pointer roll back in prefix 
			// use the table to find the previous compare point
				lg = lps[lg-1]
			}
		}
	}

	return lps
}

func (pf *PatternFinder) FindIn(target string) []int {
	r := []rune(target)
	rslt := []int{}

	i, m := 0, len(r)
	j, n := 0, len(pf.pattern)

	// pointer i in the text doesn't go back
	for i < m {
		// current rune in text doesn't match current rune in pattern
		if r[i] != pf.pattern[j] {
			// start of the pattern, skip the current rune in text
			if j == 0 {
				i++
			// go back to the previous longest proper prefix position
			}else{
				j = pf.lps[j-1]
			}		
		}else{
			// match
			i++
			j++

			// find a result
			if j == n {
				rslt = append(rslt, i - n)
				// don't need to compare from the begin of the pattern
				j = pf.lps[j-1]
			}
		}
	}

	return rslt
}

func (pf *PatternFinder) Pattern() []rune{
	return pf.pattern
}

func (pf *PatternFinder) GetLPS() []int {
	return pf.lps
}
