package l33tc0d3

import "strconv"

func groupAnagrams(strs []string) [][]string {
	// Store the groups in hash map
	groups := map[string][]string{}
	for _, s := range strs {
		// count frequency of characters in string.
		// to main order of lettering across different words, we use an array instead of map
		frequency := [26]int{}
		for _, f := range s {
			frequency[int(f)-97]++
		}

		// create hash from the frequency map.
		hashIdx := ""
		for _, f := range frequency {
			hashIdx = hashIdx + strconv.Itoa(f) + "," // comma separation necessary for #s with more than one digit
		}
		hashIdx = hashIdx[:len(hashIdx)-1]

		// finally check if hash exists, if it does
		if _, ok := groups[hashIdx]; ok {
			groups[hashIdx] = append(groups[hashIdx], s)
		} else {
			groups[hashIdx] = []string{s}
		}
	}
	result := [][]string{}
	for _, s := range groups {
		result = append(result, s)
	}
	return result
}

// Below is an optimal solution I found on the forums, it is much nicer
/*
func groupAnagrams(words []string) [][]string {

	// This is basically a map that groups the anagrams based on the letter counts
	// The interesting thing is the key is an array of bytes, there is no need to convert this into a string key
	// The operations on strings are oftentimes a lot more expensive, and underneath they operate on byte slices/arrays
	// So keys are a slice of bytes, and values are the index where you would append the word in the result array
	// Again secret sauce here is [26]byte the slice of bytes that we store the count for reach letter
	// Many languages do not support arbitrary data structures as keys for Maps, hence the bias in multiple solutions
	// to convert the list to a string, and/or to work with strings. No need. A slice of bytes for a map key
	// is perfectly fine in golang
	cache := make(map[[26]byte]int)

	// No magic here, just a result slice to accumulate all the anagrams and return at the end
	result := make([][]string, 0)

	// For every word in the words input
	for i := range words {
		// We initialize an empty slice of bytes
		// The 26 is the number of letter in the alphabet, i.e. "abcdefghijklmnopqrstuvwxyz"
		list := [26]byte{}

		// For each letter in the words[i] word
		for j := range words[i] {

			// This is very nice here, 'a' is the rune (please note single quotes) of the character "a"
			// which will resolve to the ASCII code of lowercase a in this case 97
			// So for example if the character is a, then words[i][j]-'a' == 0, e.g. 97 - 97 == 0
			// That means that letter a will be stored at index 0 in the list
			// If we would have letter "b" we will get words[i][j]-'a' == 1, and b will be stored at index 1 in the list
			// Because the list is initialized with 0 at every index, we just need to increment the value
			list[words[i][j]-'a']++
		}

		// We look into the cache, if the list "signature" is the same for two different words then they should be
		// bucketed in the same result slice
		// If we find it in the cache, the cache value will be the index of slice we need to append in the result slice
		if idx, ok := cache[list]; ok {
			// we append the word to the right slice in the result
			result[idx] = append(result[idx], words[i])
		} else {
			// This means that we have not found before a word with the same list signature
			// so we need to push an entry into the cache
			// We initialize a slice of strings with one entry which is the current word
			result = append(result, []string{words[i]})

			// We make an entry into the cache, with the key of list, and the value of len(result) - 1
			// The value is len(result) - 1, because if we do not find an entry in the cache, that means we need to
			// create a new slice in the result.
			// In this case we have already appended a new slice into the result, so if this is the second time around
			// this will be 1, and then next one will be 2, and so on
			cache[list] = len(result) - 1
		}
	}

	// We return the result
	return result
}```
*/
