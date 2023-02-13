package core

// Coordinate right-handed coordinate system
type Coordinate struct {
	X int
	Y int
}

func TransformMatrixCoordinatesToArrayIndex(x, y, width, height int) int {
	return y*width + x
}

func TransformArrayIndexToMatrixCoordinates(index, width, height uint) (uint, uint) {
	return index % width, index / width
}
