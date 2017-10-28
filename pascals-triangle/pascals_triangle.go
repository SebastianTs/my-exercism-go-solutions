package pascal

const testVersion = 1

// Triangle returns a pascals triangle
func Triangle(n int) (out [][]int) {

	for i := 0; i < n; i++ {
		out = append(out, []int{})
		out[i] = append(out[i], 1)
		if i > 0 {
			for j := 1; j < i; j++ {
				out[i] = append(out[i], out[i-1][j-1]+out[i-1][j])
			}
			out[i] = append(out[i], 1)
		}
	}
	return out
}
