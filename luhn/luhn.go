package luhn

// Validate check according to the Luhn algorithm
func Validate(number string) bool {
	var sum int
	var alternate bool

	for i := len(number) - 1; i >= 0; i-- {
		// convert character to integer digit
		digit := number[i] - '0'
		if digit < 0 || digit > 9 {
			return false
		}

		if alternate {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}

		sum += int(digit)
		alternate = !alternate
	}

	return sum%10 == 0
}
