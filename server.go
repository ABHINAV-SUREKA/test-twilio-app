// Replies to whatsapp message
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/twilio/twilio-go/twiml"
)

func echoString(w http.ResponseWriter, r *http.Request) {

	jokes := []string{"Two guys walked into a bar. The third guy ducked.",
		"I had a date last night and it was perfect. Tomorrow, I’ll have a fig.",
		"How do you get a country girl’s attention? A tractor.",
		"What do you call it when a group of apes starts a company? Monkey business.",
		"I wanted to buy a pair of camouflage pants, but I couldn't find any.",
		"Why are elevator jokes so classic and good? They work on many levels.",
		"A ghost walks into a bar. The bartender says, 'Sorry we don't serve spirits'",
		"I asked my dog what's two minus two. He said nothing.",
		"A guy lost left leg and left hand in an accident. He is all right now."}

	message := &twiml.MessagingMessage{
		Body: string(jokes[rand.Intn(len(jokes))]),
	}
	twimlResult, _ := twiml.Messages([]twiml.Element{message})

	fmt.Fprintf(w, twimlResult)
}

func main() {
	http.HandleFunc("/", echoString)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
