package helper

/**
Export function
***********************
To export a function capitilize the first letter of the function
**/

func ComputeFact(val int) int {
	if val == 1 {
		return 1
	}
	return val * ComputeFact(val-1)
}
