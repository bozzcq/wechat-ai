package enum

// Virtual key codes
const (
	VK_A = 65 + iota
	VK_B
	VK_C
	VK_D
	VK_E
	VK_F
	VK_G
	VK_H
	VK_I
	VK_J
	VK_K
	VK_L
	VK_M
	VK_N
	VK_O
	VK_P
	VK_Q
	VK_R
	VK_S
	VK_T
	VK_U
	VK_V
	VK_W
	VK_X
	VK_Y
	VK_Z
)

type Offset struct {
	Left float32 `json:"left"`
	Top  float32 `json:"top"`
}
type Size struct {
	Width, Height int32
}
