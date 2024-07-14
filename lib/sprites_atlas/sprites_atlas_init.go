package spritesatlas

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sa *SpriteAtlas) GetDebugDataLine() string {
	return fmt.Sprintf("Colors %d, directions %d, frames in first %d", len(sa.atlas), len(sa.atlas[0]), len(sa.atlas[0][0]))
}

func CreateAtlasFromFile(filename string, desiredSpriteW, desiredSpriteH, totalFrames int, createAllColors bool) *SpriteAtlas {

	newAtlas := SpriteAtlas{
		// spriteSize: desiredSpriteSize * int(SPRITE_SCALE_FACTOR),
	}

	if createAllColors {
		newAtlas.atlas = make([][][]rl.Texture2D, len(FactionColors))
	} else {
		newAtlas.atlas = make([][][]rl.Texture2D, 1)
	}

	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	file.Close()

	originalSpriteW := img.Bounds().Dx() / totalFrames
	originalSpriteH := img.Bounds().Dy()

	for i := range newAtlas.atlas {
		newAtlas.atlas[i] = make([][]rl.Texture2D, 1)

		for currFrame := 0; currFrame < totalFrames; currFrame++ {
			currPic := extractSubimageFromImage(img, currFrame*originalSpriteW, 0, originalSpriteW, originalSpriteH)
			rlImg := rl.NewImageFromImage(currPic)
			if createAllColors {
				replaceImageColorsToFactionImages(rlImg, i)
			}
			rl.ImageResizeNN(rlImg, int32(desiredSpriteW)*int32(SpriteScaleFactor), int32(desiredSpriteH)*int32(SpriteScaleFactor))
			newAtlas.atlas[i][0] = append(newAtlas.atlas[i][0], rl.LoadTextureFromImage(rlImg))
		}
	}
	DebugWritef("LOADING %s: created atlas {%s}\n", filename, newAtlas.GetDebugDataLine())
	return &newAtlas
}

func CreateDirectionalAtlasFromFile(filename string, originalSpriteSize, desiredSpriteSize, totalFrames, directionsInFile int, createAllColors bool) *SpriteAtlas {
	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	file.Close()

	newAtlas := SpriteAtlas{
		// spriteSize: desiredSpriteSize * int(SPRITE_SCALE_FACTOR),
	}

	if createAllColors {
		newAtlas.atlas = make([][][]rl.Texture2D, len(FactionColors))
	} else {
		newAtlas.atlas = make([][][]rl.Texture2D, 1)
	}

	for i := range newAtlas.atlas {

		newAtlas.atlas[i] = make([][]rl.Texture2D, 4*directionsInFile)

		for currFrame := 0; currFrame < totalFrames; currFrame++ {
			for currDirectionFromFile := 0; currDirectionFromFile < directionsInFile; currDirectionFromFile++ {
				currPic := extractSubimageFromImage(img, currFrame*originalSpriteSize, currDirectionFromFile*originalSpriteSize, originalSpriteSize, originalSpriteSize)
				rlImg := rl.NewImageFromImage(currPic)
				if createAllColors {
					replaceImageColorsToFactionImages(rlImg, i)
				}
				rl.ImageResizeNN(rlImg, int32(desiredSpriteSize)*int32(SpriteScaleFactor), int32(desiredSpriteSize)*int32(SpriteScaleFactor))
				newAtlas.atlas[i][currDirectionFromFile] = append(newAtlas.atlas[i][currDirectionFromFile], rl.LoadTextureFromImage(rlImg))
				for currDir := 1; currDir < 4; currDir++ {
					rl.ImageRotateCW(rlImg)
					newAtlas.atlas[i][currDir*directionsInFile+currDirectionFromFile] =
						append(newAtlas.atlas[i][currDir*directionsInFile+currDirectionFromFile], rl.LoadTextureFromImage(rlImg))
				}
			}
		}
	}
	DebugWritef("LOADING %s: created dir-atlas {%s}", filename, newAtlas.GetDebugDataLine())
	return &newAtlas
}

func replaceImageColorsToFactionImages(img *rl.Image, factionColorNumber int) {
	// rl.ImageColorReplace(img, color.RGBA{192, 192, 192, 255}, factionColors[factionColorNumber])
	rl.ImageColorReplace(img, color.RGBA{255, 0, 255, 255}, FactionColors[factionColorNumber])
	darkerFactionTint := FactionColors[factionColorNumber]
	darkerFactionTint.R /= 2
	darkerFactionTint.G /= 2
	darkerFactionTint.B /= 2
	// rl.ImageColorReplace(img, color.RGBA{128, 128, 128, 255}, darkerFactionTint)
	rl.ImageColorReplace(img, color.RGBA{128, 0, 128, 255}, darkerFactionTint)
	darkestFactionTint := FactionColors[factionColorNumber]
	darkestFactionTint.R /= 3
	darkestFactionTint.G /= 3
	darkestFactionTint.B /= 3
	// rl.ImageColorReplace(img, color.RGBA{64, 64, 64, 255}, darkestFactionTint)
	rl.ImageColorReplace(img, color.RGBA{64, 0, 64, 255}, darkestFactionTint)
}

func extractSubimageFromImage(img image.Image, fromx, fromy, w, h int) image.Image {
	minx, miny := img.Bounds().Min.X, img.Bounds().Min.Y
	//maxx, maxy := img.Bounds().Min.X, img.Bounds().Max.Y
	switch img.(type) {
	case *image.RGBA:
		subImg := img.(*image.RGBA).SubImage(
			image.Rect(minx+fromx, miny+fromy, minx+fromx+w, miny+fromy+h),
		)
		// reset img bounds, because RayLib goes nuts about it otherwise
		subImg.(*image.RGBA).Rect = image.Rect(0, 0, w, h)
		return subImg
	case *image.NRGBA:
		subImg := img.(*image.NRGBA).SubImage(
			image.Rect(minx+fromx, miny+fromy, minx+fromx+w, miny+fromy+h),
		)
		// reset img bounds, because RayLib goes nuts about it otherwise
		subImg.(*image.NRGBA).Rect = image.Rect(0, 0, w, h)
		return subImg
	case *image.Paletted:
		subImg := img.(*image.Paletted).SubImage(
			image.Rect(minx+fromx, miny+fromy, minx+fromx+w, miny+fromy+h),
		)
		// reset img bounds, because RayLib goes nuts about it otherwise
		subImg.(*image.Paletted).Rect = image.Rect(0, 0, w, h)
		return subImg
	default:
	}
	panic(fmt.Sprintf("\nUnknown image type %T", img))
}
