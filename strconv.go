package cursed

import "strconv"

// ınt turns your string into an int!!!
func ınt(s string) int {
	// Ignore errors YOLO
	i, _ := strconv.Atoi(s)
	return i
}
