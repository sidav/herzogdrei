package spritesatlas

import (
	"herzog/lib/geometry"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteAtlas struct {
	// first index is color mask, second is sprite number (rotation is there), third is frame number (animation)
	atlas [][][]rl.Texture2D
	// spriteSize int // width of square sprite
}

func (sa *SpriteAtlas) TotalFrames() int {
	if sa != nil && len(sa.atlas) > 0 && len(sa.atlas[0]) > 0 {
		return len(sa.atlas[0][0])
	}
	return 0
}

func (sa *SpriteAtlas) GetSpriteByFrame(frameNum int) rl.Texture2D {
	return sa.atlas[0][0][frameNum]
}

func (sa *SpriteAtlas) GetSpriteByColorAndFrame(color, frameNum int) rl.Texture2D {
	if frameNum >= sa.TotalFrames() {
		frameNum = frameNum % sa.TotalFrames()
	}
	return sa.atlas[color][0][frameNum]
}

//func (sa *spriteAtlas) getSpriteByFrame(frameNum int) rl.Texture2D {
//	return sa.atlas[0][0][frameNum]
//}

func (sa *SpriteAtlas) GetSpriteByDirectionAndFrameNumber(dx, dy, num int) rl.Texture2D {
	var spriteGroup uint8 = 0
	if dx == 1 {
		spriteGroup = 3
	}
	if dx == -1 {
		spriteGroup = 1
	}
	if dy == 1 {
		spriteGroup = 2
	}
	num = num % len(sa.atlas[spriteGroup])
	return sa.atlas[0][spriteGroup][num]
}

func (sa *SpriteAtlas) GetSpriteByColorDegreeAndFrameNumber(color, degree, num int) rl.Texture2D {
	totalRotationSprites := len(sa.atlas[color])
	rotFrame := geometry.DegreeToSectorNumber(degree, totalRotationSprites)
	// Zero degree looks to the right, but the first sprite in atlas looks to the up.
	// So, add totalRotationSprites/4 to rotFrame to compensate
	rotFrame = (totalRotationSprites/4 + rotFrame) % totalRotationSprites
	num = num % len(sa.atlas[color][rotFrame])
	return sa.atlas[color][rotFrame][num]
}
