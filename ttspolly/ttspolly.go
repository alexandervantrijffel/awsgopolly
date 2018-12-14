package ttspolly

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
)

func SaySSML(svc *polly.Polly, text, fileName, voice string) {
	text = "<speak>" + text + "</speak>"
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(text),
		VoiceId:      aws.String(voice),
		TextType:     aws.String("ssml")}

	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		panic(err)
	}

}

func SetSpeed(text string, speed int) string {
	if speed != 100 {
		return fmt.Sprintf(`<prosody rate="%d%%">%s</prosody>`, speed, text)
	}
	return text
}

const (
	Geraint   = "Geraint"
	Gwyneth   = "Gwyneth"
	Mads      = "Mads"
	Naja      = "Naja"
	Hans      = "Hans"
	Marlene   = "Marlene"
	Nicole    = "Nicole"
	Russell   = "Russell"
	Amy       = "Amy"
	Brian     = "Brian"
	Emma      = "Emma"
	Raveena   = "Raveena"
	Ivy       = "Ivy"
	Joanna    = "Joanna"
	Joey      = "Joey"
	Justin    = "Justin"
	Kendra    = "Kendra"
	Kimberly  = "Kimberly"
	Salli     = "Salli"
	Conchita  = "Conchita"
	Enrique   = "Enrique"
	Miguel    = "Miguel"
	Penelope  = "Penelope"
	Chantal   = "Chantal"
	Celine    = "Celine"
	Mathieu   = "Mathieu"
	Dora      = "Dora"
	Karl      = "Karl"
	Carla     = "Carla"
	Giorgio   = "Giorgio"
	Mizuki    = "Mizuki"
	Liv       = "Liv"
	Lotte     = "Lotte"
	Ruben     = "Ruben"
	Ewa       = "Ewa"
	Jacek     = "Jacek"
	Jan       = "Jan"
	Maja      = "Maja"
	Ricardo   = "Ricardo"
	Vitoria   = "Vitoria"
	Cristiano = "Cristiano"
	Ines      = "Ines"
	Carmen    = "Carmen"
	Maxim     = "Maxim"
	Tatyana   = "Tatyana"
	Astrid    = "Astrid"
	Filiz     = "Filiz"
	Aditi     = "Aditi"
	Matthew   = "Matthew"
)
