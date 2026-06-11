type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""

	for _, str := range strs {
		encoded += strconv.Itoa(len(str)) + "#" + str
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		start := j + 1
		end := start + length

		res = append(res, encoded[start:end])

		i = end
	}

	return res
}