package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"

	"github.com/alexandervantrijffel/awsgopolly/ttspolly"
)

func main() {
	// Initialize a session that the SDK will use to load credentials from the shared credentials file ~/.aws/credentials, load your configuration from the shared configuration file ~/.aws/config, and create an Amazon Polly client.

	// Format of ~/.aws/credentials https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-file-format

	// Format of ~/.aws/config https://docs.aws.amazon.com/cli/latest/topic/config-vars.html<Paste>
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := polly.New(sess)
	text := "Cessna 9130 Delta, Oakland Center, radio check"
	generateMp3(svc, text)
	generateHigherSpeed(svc, text)
}

func generateMp3(svc *polly.Polly, text string) {
	ttspolly.SaySSML(svc, text, "en-gb-brian-radiocheck.mp3", ttspolly.Brian)
	ttspolly.SaySSML(svc, text, "en-gb-amy-radiocheck.mp3", ttspolly.Amy)
	ttspolly.SaySSML(svc, text, "en-gb-emma-radiocheck.mp3", ttspolly.Emma)
	ttspolly.SaySSML(svc, text, "en-au-russell-radiocheck.mp3", ttspolly.Russell)
	ttspolly.SaySSML(svc, text, "en-au-nicole-radiocheck.mp3", ttspolly.Nicole)
	ttspolly.SaySSML(svc, text, "en-in-aditi-radiocheck.mp3", ttspolly.Aditi)
	ttspolly.SaySSML(svc, text, "en-in-raveena-radiocheck.mp3", ttspolly.Raveena)
	ttspolly.SaySSML(svc, text, "en-us-joey-radiocheck.mp3", ttspolly.Joey)
	ttspolly.SaySSML(svc, text, "en-us-justin-radiocheck.mp3", ttspolly.Justin)
	ttspolly.SaySSML(svc, text, "en-us-matthew-radiocheck.mp3", ttspolly.Matthew)
	ttspolly.SaySSML(svc, text, "en-us-ivy-radiocheck.mp3", ttspolly.Ivy)
	ttspolly.SaySSML(svc, text, "en-us-joanna-radiocheck.mp3", ttspolly.Joanna)
	ttspolly.SaySSML(svc, text, "en-us-kendra-radiocheck.mp3", ttspolly.Kendra)
	ttspolly.SaySSML(svc, text, "en-us-kimberly-radiocheck.mp3", ttspolly.Kimberly)
	ttspolly.SaySSML(svc, text, "en-us-salli-radiocheck.mp3", ttspolly.Salli)
	ttspolly.SaySSML(svc, text, "en-gb-wls-geraint-radiocheck.mp3", ttspolly.Geraint)
	ttspolly.SaySSML(svc, text, "es-us-miguel-radiocheck.mp3", ttspolly.Miguel)
	ttspolly.SaySSML(svc, text, "es-us-penelope-radiocheck.mp3", ttspolly.Penelope)
	ttspolly.SaySSML(svc, text, "fr-fr-mathieu-radiocheck.mp3", ttspolly.Mathieu)
}

func generateHigherSpeed(svc *polly.Polly, text string) {
	ttspolly.SaySSML(svc, ttspolly.SetSpeed(text, 110), "en-us-matthew-110-radiocheck.mp3", ttspolly.Matthew)
	ttspolly.SaySSML(svc, ttspolly.SetSpeed(text, 120), "en-us-matthew-120-radiocheck.mp3", ttspolly.Matthew)
	ttspolly.SaySSML(svc, ttspolly.SetSpeed(text, 130), "en-us-matthew-130-radiocheck.mp3", ttspolly.Matthew)
	ttspolly.SaySSML(svc, ttspolly.SetSpeed(text, 140), "en-us-matthew-140-radiocheck.mp3", ttspolly.Matthew)

}
