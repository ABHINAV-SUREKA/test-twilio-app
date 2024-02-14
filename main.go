// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"fmt"
	"math/rand"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure

	jokes := []string{"Two guys walked into a bar. The third guy ducked.",
		"I had a date last night and it was perfect. Tomorrow, I’ll have a fig.",
		"How do you get a country girl’s attention? A tractor.",
		"What do you call it when a group of apes starts a company? Monkey business.",
		"I wanted to buy a pair of camouflage pants, but I couldn't find any.",
		"Why are elevator jokes so classic and good? They work on many levels.",
		"A ghost walks into a bar. The bartender says, 'Sorry we don't serve spirits'",
		"I asked my dog what's two minus two. He said nothing."}

	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	joke := string(jokes[rand.Intn(len(jokes))])
	fmt.Println(joke)
	params.SetBody(joke)
	params.SetFrom("+14847256474")
	// params.SetMediaUrl([]string{"https://c1.staticflickr.com/3/2899/14341091933_1e92e62d12_b.jpg"})
	params.SetTo("+919740446065")

	resp2, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp2.Sid != nil {
			fmt.Println(*resp2.Sid)
		} else {
			fmt.Println(resp2.Sid)
		}
	}
}
