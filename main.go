package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math"
	"math/rand"
	"strconv"
	"github.com/hajimehoshi/bitmapfont"
	"golang.org/x/image/colornames"
	"github.com/hajimehoshi/ebiten/text"
	"log"
	"github.com/hajimehoshi/ebiten"
)

var buttons []*sprite
var myHand int
var yourHand int
var waiting bool
var turn int
var myPt int
var yourPt int
var ready bool
var getPt int
var gameSet bool
var myName string

func init() {
	myName = js.Global.Call("prompt", "名前を入力してください(二文字まで)").String()

	buttonImage, _ := ebiten.NewImage(20, 20, ebiten.FilterDefault)
	buttonImage.Fill(colornames.Silver)
	buttons = append(buttons, newSprite(buttonImage, 15, 100), newSprite(buttonImage, 30, 130), newSprite(buttonImage, 60, 130), newSprite(buttonImage, 90, 130), newSprite(buttonImage, 108, 100))

	turn = rand.Intn(2)
}

type sprite struct {
	image *ebiten.Image
	x     float64
	y     float64
	scaleX float64
	scaleY float64
}

func newSprite(image *ebiten.Image, x, y float64) *sprite {
	return &sprite{
		image: image,
		x: x,
		y: y,
		scaleX: 1.0,
		scaleY: 1.0,
		// alpha: 1.0,
		// monochrome: false,
	}
}

func (s *sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scaleX, s.scaleY)
	op.GeoM.Translate(s.x, s.y)
	/*
	op.ColorM.Scale(1.0, 1.0, 1.0, s.alpha)
	if s.monochrome {
		op.ColorM.ChangeHSV(0, 0, 1)
	}
	*/
	
	screen.DrawImage(s.image, op)
}

func update(screen *ebiten.Image) error {

	for i, v := range buttons {
		if v.IsJustPressed() && !waiting && !gameSet {
			println(i + 1)
			myHand = i + 1
			waiting = true
			break
		}
	}

	text.Draw(screen, "2円が...", bitmapfont.Gothic12r, 20, 20, colornames.White)
	

	for _, v := range buttons {
		v.Draw(screen)
	}

	text.Draw(screen, "1", bitmapfont.Gothic12r, 25, 110, colornames.Black)
	text.Draw(screen, "2", bitmapfont.Gothic12r, 40, 140, colornames.Black)
	text.Draw(screen, "3", bitmapfont.Gothic12r, 70, 140, colornames.Black)
	text.Draw(screen, "4", bitmapfont.Gothic12r, 100, 140, colornames.Black)
	text.Draw(screen, "5", bitmapfont.Gothic12r, 118, 110, colornames.Black)


	text.Draw(screen, myName, bitmapfont.Gothic12r, 40, 58, colornames.White)
	text.Draw(screen, "相手", bitmapfont.Gothic12r, 86, 58, colornames.White)
	if myHand > 0 {
		text.Draw(screen, strconv.Itoa(myHand), bitmapfont.Gothic12r, 50, 75, colornames.White)
	}
	if yourHand > 0 {
		text.Draw(screen, strconv.Itoa(yourHand), bitmapfont.Gothic12r, 96, 75, colornames.White)
	}

	if turn % 2 == 0 {
		text.Draw(screen, "攻め", bitmapfont.Gothic12r, 40, 40, colornames.White)
		text.Draw(screen, "守り", bitmapfont.Gothic12r, 86, 40, colornames.White)
	}else {
		text.Draw(screen, "守り", bitmapfont.Gothic12r, 40, 40, colornames.White)
		text.Draw(screen, "攻め", bitmapfont.Gothic12r, 86, 40, colornames.White)
	}

	if waiting {
		yourHand = rand.Intn(5) + 1
		ready = true
		waiting = false
	}

	if ready {
		getPt = int(math.Abs(float64(myHand - yourHand)))
		if turn % 2 == 0 {//my turn
			myPt += getPt
		}else {//your turn
			yourPt += getPt
		}
		ready = false
		turn += 3
	} 

	text.Draw(screen, myName, bitmapfont.Gothic12r, 40, 180, colornames.White)
	text.Draw(screen, "相手", bitmapfont.Gothic12r, 86, 180, colornames.White)
	text.Draw(screen, strconv.Itoa(myPt), bitmapfont.Gothic12r, 50, 197, colornames.White)
	text.Draw(screen, strconv.Itoa(yourPt), bitmapfont.Gothic12r, 96, 197, colornames.White)
	
	if turn > 2 {
		if turn % 2 == 0 {
			text.Draw(screen, "+" + strconv.Itoa(getPt), bitmapfont.Gothic12r, 93, 214, colornames.Red)
		}else {
			text.Draw(screen, "+" + strconv.Itoa(getPt), bitmapfont.Gothic12r, 47, 214, colornames.Red)
		}
	}

	if myPt >= 10 {
		text.Draw(screen, "勝ち！", bitmapfont.Gothic12r, 62, 232, colornames.Red)
		gameSet = true
	}else if yourPt >= 10 {
		text.Draw(screen, "負け...", bitmapfont.Gothic12r, 62, 232, colornames.Blue)
		gameSet = true
	}
	

	//https://data-14938.firebaseio.com/

	return nil
}

func main() {
	if err := ebiten.Run(update, 144, 256, 3, "2円"); err != nil {
		log.Fatal(err)
	}
}