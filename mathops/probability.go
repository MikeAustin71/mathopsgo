package mathops

/*
	Permutations
  ============

	nPr	 =     n!
            ----
            (n-r)!

	n things taken r at a time

  Combinations
  ============

	nCr		=            n!
                -----------
								(n-r)! (r!)

 */

type Probability struct {
		PercentCertainty 	BigIntNum
		NoOfPossibilitiesThatMeetConditions BigIntNum
		NoOfPossibilities   BigIntNum
}
