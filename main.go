package main

import "C"

import (
	"log"
	"math/rand"
	"time"
	"unsafe"

	gg "github.com/nanitefactory/amidakuji/glossary"
	"github.com/nanitefactory/amidakuji/glossary/jukebox"
	"github.com/nanitefactory/winmb"
)

//export Test
func Test() {
	winmb.MessageBoxPlain("export Test", "export Test")
}

// OnProcessAttach is an async callback (hook).
//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	defer func() {
		err := jukebox.Finalize()
		if err != nil {
			log.Fatal(err)
		}
	}()
	rand.Seed(time.Now().UnixNano())

	conf := askConf()
	if conf == nil {
		conf = map[string]interface{}{
			"window_width":  800.0,
			"window_height": 800.0,
			"max_player":    10.0,
			"max_level":     100.0,
			"width":         1500.0,
			"height":        1500.0,
			"zoom":          -4.0,
			"rotate_degree": 270.0,
			"margin_top":    50.0,
			"margin_right":  100.0,
			"margin_bottom": 50.0,
			"margin_left":   200.0,
			"font_size":     28.0,
			"picks":         []interface{}{"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon", "Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie", "Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill", "Pidgey", "Pidgeotto", "Pidgeot", "Rattata"},
			"prizes":        []interface{}{"TM88", "TM89", "TM90", "TM91", "TM92", "HM01", "HM02", "HM03", "HM04", "HM05", "HM06"},
		}
	}

	newGame(gameConfig{
		winWidth:            conf["window_width"].(float64),
		winHeight:           conf["window_height"].(float64),
		nParticipants:       int(conf["max_player"].(float64)),
		nLevel:              int(conf["max_level"].(float64)),
		width:               conf["width"].(float64),
		height:              conf["height"].(float64),
		initialZoomLevel:    conf["zoom"].(float64),
		initialRotateDegree: conf["rotate_degree"].(float64),
		paddingTop:          conf["margin_top"].(float64),
		paddingRight:        conf["margin_right"].(float64),
		paddingBottom:       conf["margin_bottom"].(float64),
		paddingLeft:         conf["margin_left"].(float64),
		fontSize:            conf["font_size"].(float64),
		nametagPicks:        gg.ItfsToStrs(conf["picks"].([]interface{})),
		nametagPrizes:       gg.ItfsToStrs(conf["prizes"].([]interface{})),
	}).Run()
}

func main() {
	// nothing really. xD
}
