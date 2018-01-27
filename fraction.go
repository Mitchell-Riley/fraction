package fraction

import (
	"fmt"
	"math"
	"strconv"
)

// taken from here because I am too lazy to figure out how to do this on my own
// https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
// returns a slice containing the prime factorization of n
func primeFactors(n int) []int {
	factors := []int{}

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

// returns how many times element n is repeated in array arr
func elementCount(arr []int, n int) int {
	count := 0
	for _, v := range arr {
		if v == n {
			count++
		}
	}
	return count
}

// returns the length of the repetend with base-b and a given prime
// factor not given in the prime factors of b p
func repetendLength(b, p int) int {
	for i := 1; ; i++ {
		if (int(math.Pow(float64(b), float64(i)))-1)%p == 0 {
			return i
		}
	}
}

// returns true if a belongs to s
func intInSlice(a int, s []int) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}
	return false
}

// BaseRExpansion returns a string representation of the base b expansion of n/d along
// with the termination/delay length and the repetend length
// The resulting string contains the quotient. "[]" signifies the repetend
func BaseRExpansion(b, n, d int) (string, int, int, error) {
	if n > d {
		return "", 0, 0, fmt.Errorf("Improper fraction (numerator is larger than denominator)")
	}

	baseFactors := primeFactors(b)
	denominatorFactors := primeFactors(int(d))

	// delay is the maximum power of the base's prime factors
	// if the fraction is terminating, than the delay is the length of the termination
	delay := 0
	for _, v := range baseFactors {
		if e := elementCount(denominatorFactors, v); e > delay {
			delay = e
		}
	}

	potentialC := []int{}
	for _, v := range denominatorFactors {
		if !intInSlice(v, baseFactors) {
			potentialC = append(potentialC, v)
		}
	}

	// if c:=0, then we can't multiply by the rest of the factors
	// if c:=1, then its impossible to have an accurate repetend of length 0
	// start at the first element to avoid multiplying by 0
	c := potentialC[0]
	for i := 1; i < len(potentialC); i++ {
		c *= potentialC[i]
	}

	repetend := repetendLength(b, c)
	fraction := "0."

	// not a great name, but it's the best I can think of rn
	j := 0
	significand := float64(b) * (float64(n) / float64(d))
	for j < delay+repetend {
		// Modf splits a mixed number into its whole number component and its decimal component
		i, m := math.Modf(significand)

		if j == delay {
			fraction += "["
		}

		fraction += strconv.Itoa(int(i))

		if j == delay+repetend-1 {
			fraction += "]"
		}

		// I didn't know this syntax existed/worked
		j++
		significand = float64(b) * m
	}

	return fraction, delay, repetend, nil
}
