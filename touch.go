package main

import (
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"

)

func (s *sprite) IsJustPressed() bool {
	rect := s.image.Bounds()
	rx := int(float64(rect.Dx()) * s.scaleX)
	ry := int(float64(rect.Dy()) * s.scaleY)

	cursorX, cursorY := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if cursorX >= int(s.x) && cursorX <= int(s.x)+rx && cursorY >= int(s.y) && cursorY <= int(s.y)+ry {
			return true
		}
	}

	touchX, touchY := getJustTouchPosition()
	if touchX >= int(s.x) && touchX <= int(s.x)+rx && touchY >= int(s.y) && touchY <= int(s.y)+ry {
		return true
	}

	return false
}

func getJustTouchPosition() (int, int) {
	var x, y int
	ts := inpututil.JustPressedTouchIDs()
	// ts := gGame.justTouchIDs
	if len(ts) == 1 {
		x, y = ebiten.TouchPosition(ts[0])
	}

	return x, y
}