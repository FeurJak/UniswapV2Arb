func ArbCalc(Ra []*big.Float, Rb []*big.Float, Fees []*big.Float, _reverseFlags []bool) (*big.Int, *big.Int) {

	_OptimalInput := new(big.Float)
	_Lambda := new(big.Float)
	_OptimalOut := new(big.Float)

	_OptimalInput_int := new(big.Int)
	_OptimalOut_int := new(big.Int)

	_lp_len := len(Ra)

	// Now we calculate Input Amount and Output Amount
	FeeProduct := make([]*big.Float, _lp_len)
	FeeProduct[0] = Fees[0]

	Ea := make([]*big.Float, _lp_len)
	Eb := make([]*big.Float, _lp_len)

	for i := 1; i < _lp_len; i++ {
		FeeProduct[i] = big.NewFloat(0).Mul(Fees[i], FeeProduct[i-1])
	}

	for i := 0; i < _lp_len; i++ {
		if i == 0 {
			if !_reverseFlags[0] {
				Ea[0] = Ra[0]
				Eb[0] = Rb[0]
			} else {
				Ea[0] = Rb[0]
				Eb[0] = Ra[0]
			}

		} else {
			if !_reverseFlags[i] { // not Reversed
				Ea[i] = big.NewFloat(0).Quo(big.NewFloat(0).Mul(Ra[i], Ea[i-1]), big.NewFloat(0).Add(Ra[i], big.NewFloat(0).Mul(Eb[i-1], FeeProduct[i-1])))
				Eb[i] = big.NewFloat(0).Quo(big.NewFloat(0).Mul(Rb[i], Eb[i-1]), big.NewFloat(0).Add(Ra[i], big.NewFloat(0).Mul(Eb[i-1], FeeProduct[i-1])))
			} else {

				Ea[i] = big.NewFloat(0).Quo(big.NewFloat(0).Mul(Rb[i], Ea[i-1]), big.NewFloat(0).Add(Rb[i], big.NewFloat(0).Mul(Eb[i-1], FeeProduct[i-1])))
				Eb[i] = big.NewFloat(0).Quo(big.NewFloat(0).Mul(Ra[i], Eb[i-1]), big.NewFloat(0).Add(Rb[i], big.NewFloat(0).Mul(Eb[i-1], FeeProduct[i-1])))

			}
		}
	}
	// Then We calculate Input Amount
	_Numerator := big.NewFloat(0).Sub(big.NewFloat(0).Sqrt(big.NewFloat(0).Mul(big.NewFloat(0).Mul(Ea[_lp_len-1], Eb[_lp_len-1]), FeeProduct[_lp_len-1])), Ea[_lp_len-1])
	_OptimalInput.Quo(_Numerator, FeeProduct[_lp_len-2])
	if _OptimalInput.Cmp(big.NewFloat(0)) > 0 {
		_OptimalInput.Int(_OptimalInput_int)
		_Lambda.Sqrt(big.NewFloat(0).Quo(big.NewFloat(0).Mul(Eb[_lp_len-1], FeeProduct[_lp_len-1]), Ea[_lp_len-1]))
		_OptimalOut.Mul(_OptimalInput, _Lambda)
		_OptimalOut.Int(_OptimalOut_int)
		return _OptimalInput_int, _OptimalOut_int
	}
	return nil, nil
}