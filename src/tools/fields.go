package tools

func CopyField[T string | float32 | int](src T, des T, default_value T) T {
	if src != default_value {
		des = src
	}
	return des
}
