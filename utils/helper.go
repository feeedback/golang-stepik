package utils

type MapFnType[I any, O any] func(I) O

func MapFn[I any, O any](input []I, fn MapFnType[I, O]) []O {
	output := make([]O, len(input))

	for i, val := range input {
		output[i] = fn(val)
	}

	return output
}
