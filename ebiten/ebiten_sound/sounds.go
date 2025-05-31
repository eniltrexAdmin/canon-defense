package ebiten_sound

import (
	"bytes"
	"canon-tower-defense/ebiten/constants"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
)

func MustNewPlayer(data []byte) *audio.Player {
	stream, err := wav.DecodeWithSampleRate(constants.AudioContext.SampleRate(), bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	player, err := constants.AudioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}
	return player
}
