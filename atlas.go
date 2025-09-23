package grender

import (
	"grender/util"
	"image"
	"image/draw"
	"log"
	"math"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Atlas ref impl: https://github.com/juj/RectangleBinPack/blob/83e7e1132d93777e3732dfaae26b0f3703be2036/MaxRectsBinPack.cpp
type Atlas struct {
	Size                 uint32
	FreeRects            []RectP
	usedRectangles       []RectP
	newFreeRectsLastSize int
	newFreeRects         []RectP
	Img                  *image.NRGBA
}

type AtlasOption int32

func NewAtlas(opts ...AtlasOption) *Atlas {
	if len(opts) > 1 {
		log.Fatalln("more than one opts int to new atlas")
	}
	var maxTextureSize int32
	if len(opts) == 0 {
		maxTextureSize = util.GetGlIntP(gl.MAX_TEXTURE_SIZE)
	} else {
		maxTextureSize = int32(opts[0])
	}

	return &Atlas{
		Size:      uint32(maxTextureSize),
		FreeRects: []RectP{{0, 0, maxTextureSize, maxTextureSize}},
		Img:       image.NewNRGBA(image.Rect(0, 0, int(maxTextureSize), int(maxTextureSize))),
	}
}

func (a *Atlas) AddTexture(t *Texture) {
	// TODO: commented as unused in ref impl but will investigate later
	score1 := math.MaxInt32
	score2 := math.MaxInt32

	newNode := a.findPositionForNewNode(t.Size.X, t.Size.Y, int32(score1), int32(score2))
	a.placeRect(newNode)

	dstRect := image.Rect(int(newNode.X), int(newNode.Y), int(newNode.X+newNode.W), int(newNode.Y+newNode.H))

	t.AtlasPos = RectP{
		X: newNode.X,
		Y: newNode.Y,
		W: newNode.X + newNode.W,
		H: newNode.Y + newNode.H,
	}

	draw.Draw(a.Img, dstRect, t.Image, image.Point{0, 0}, draw.Over)
}

func (a *Atlas) placeRect(node RectP) {
	for i := 0; i < len(a.FreeRects); {
		if a.splitFreeNode(a.FreeRects[i], node) {
			a.FreeRects[i] = a.FreeRects[len(a.FreeRects)-1]
			a.FreeRects = a.FreeRects[:len(a.FreeRects)-1]
		} else {
			i++
		}
	}

	a.pruneFreeList()

	a.usedRectangles = append(a.usedRectangles, node)
}

func (a *Atlas) splitFreeNode(freeNode, usedNode RectP) bool {
	if usedNode.X >= freeNode.X+freeNode.W || usedNode.X+usedNode.W <= freeNode.X || usedNode.Y >= freeNode.Y+freeNode.H || usedNode.Y+usedNode.H <= freeNode.Y {
		return false
	}

	a.newFreeRectsLastSize = len(a.newFreeRects)

	if usedNode.X < freeNode.X+freeNode.W && usedNode.X+usedNode.W > freeNode.X {
		if usedNode.Y > freeNode.Y && usedNode.Y < freeNode.Y+freeNode.H {
			newNode := freeNode
			newNode.H = usedNode.Y - newNode.Y
			a.insertNewFreeRect(newNode)
		}

		if usedNode.Y+usedNode.H < freeNode.Y+freeNode.H {
			newNode := freeNode
			newNode.Y = usedNode.Y + usedNode.H
			newNode.H = freeNode.Y + freeNode.H - (usedNode.Y + usedNode.H)
			a.insertNewFreeRect(newNode)
		}
	}

	if usedNode.Y < freeNode.Y+freeNode.H && usedNode.Y+usedNode.H > freeNode.Y {
		if usedNode.X > freeNode.X && usedNode.X < freeNode.X+freeNode.W {
			newNode := freeNode
			newNode.W = usedNode.X - newNode.X
			a.insertNewFreeRect(newNode)
		}

		if usedNode.X+usedNode.W < freeNode.X+freeNode.W {
			newNode := freeNode
			newNode.X = usedNode.X + usedNode.W
			newNode.W = freeNode.X + freeNode.W - (usedNode.X + usedNode.W)
			a.insertNewFreeRect(newNode)
		}
	}

	return true
}

func (a *Atlas) insertNewFreeRect(newFreeRect RectP) {
	// equiv to assert(newFreeRect.width > 0); from ref impl
	if newFreeRect.W < 0 {
		log.Fatalln("newFreeRect.W must be > 0")
	}
	// equiv to assert(newFreeRect.height > 0); from ref impl
	if newFreeRect.H < 0 {
		log.Fatalln("newFreeRect.H must be > 0")
	}

	for i := 0; i < a.newFreeRectsLastSize; {
		if newFreeRect.IsContainedIn(a.newFreeRects[i]) {
			return
		}

		if a.newFreeRects[i].IsContainedIn(newFreeRect) {
			a.newFreeRectsLastSize--
			a.newFreeRects[i] = a.newFreeRects[a.newFreeRectsLastSize]
			a.newFreeRects[a.newFreeRectsLastSize] = a.newFreeRects[len(a.newFreeRects)-1]
			a.newFreeRects = a.newFreeRects[:len(a.newFreeRects)-1]
		} else {
			i++
		}
	}

	a.newFreeRects = append(a.newFreeRects, newFreeRect)
}

func (a *Atlas) findPositionForNewNode(width, height, bestShortSideFit, bestLongSideFit int32) RectP {
	bestNode := RectP{}

	bestShortSideFit = math.MaxInt32
	bestLongSideFit = math.MaxInt32

	for i := 0; i < len(a.FreeRects); i++ {
		if a.FreeRects[i].W >= width && a.FreeRects[i].H >= height {
			leftoverHoriz := util.AbsI32(a.FreeRects[i].W - width)
			leftoverVert := util.AbsI32(a.FreeRects[i].H - height)
			shortSideFit := max(leftoverHoriz, leftoverVert)
			longSideFit := max(leftoverHoriz, leftoverVert)

			if shortSideFit < bestShortSideFit || (shortSideFit == bestShortSideFit && longSideFit < bestLongSideFit) {
				bestNode.X = a.FreeRects[i].X
				bestNode.Y = a.FreeRects[i].Y
				bestNode.W = width
				bestNode.H = height
				bestShortSideFit = shortSideFit
				bestLongSideFit = longSideFit
			}
		}
	}

	return bestNode
}

func (a *Atlas) pruneFreeList() {
	for i := 0; i < len(a.FreeRects); i++ {
		for j := 0; j < len(a.newFreeRects); {
			if a.newFreeRects[j].IsContainedIn(a.FreeRects[i]) {
				a.newFreeRects[j] = a.newFreeRects[len(a.newFreeRects)-1]
				a.newFreeRects = a.newFreeRects[:len(a.newFreeRects)-1]
			} else {
				// equiv to assert(!IsContainedIn(freeRectangles[i], newFreeRectangles[j])); in ref impl
				if a.FreeRects[i].IsContainedIn(a.newFreeRects[j]) {
					panic("a.FreeRects[i].IsContainedIn(a.newFreeRects[j] = true")
				}
				j++
			}
		}
	}

	// equiv to freeRectangles.insert(freeRectangles.end(), newFreeRectangles.begin(), newFreeRectangles.end()); in ref impl
	a.FreeRects = append(a.FreeRects, a.newFreeRects...)
	a.newFreeRects = []RectP{}
}
