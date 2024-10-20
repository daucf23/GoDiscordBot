package commands

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jokes = []string{
	"Why don't scientists trust atoms? Because they make up everything.",
	"Why did the math book look sad? Because it had too many problems.",
	"Why don't programmers like nature? It has too many bugs.",
	"What do you call fake spaghetti? An impasta.",
	"Why did the scarecrow win an award? Because he was outstanding in his field.",
	"Why can't you hear a pterodactyl go to the bathroom? Because the 'P' is silent.",
	"Why did the bicycle fall over? It was two-tired.",
	"What's orange and sounds like a parrot? A carrot.",
	"Why did the coffee file a police report? It got mugged.",
	"What do you call cheese that isn't yours? Nacho cheese.",
	"Why did the tomato blush? Because it saw the salad dressing.",
	"Why did the golfer bring two pairs of pants? In case he got a hole in one.",
	"What do you call a belt made of watches? A waist of time.",
	"Why don't skeletons fight each other? They don't have the guts.",
	"Why was the math teacher late to work? He took the rhombus.",
	"Why did the cookie go to the doctor? Because he felt crummy.",
	"Why did the chicken cross the playground? To get to the other slide.",
	"Why do bees have sticky hair? Because they use honeycombs.",
	"Why can't you trust stairs? They're always up to something.",
	"Why did the picture go to jail? Because it was framed.",
	"Why did the computer go to the doctor? Because it had a virus.",
	"Why was the broom late? It overswept.",
	"What do you call an alligator in a vest? An investigator.",
	"Why don't elephants use computers? Because they're afraid of the mouse.",
	"Why did the barber win the race? He took a shortcut.",
	"Why was the stadium so cool? It was filled with fans.",
	"Why did the music teacher need a ladder? To reach the high notes.",
	"Why was the belt arrested? For holding up the pants.",
	"What do you call a bear with no teeth? A gummy bear.",
	"Why did the man run around his bed? Because he was trying to catch up on sleep.",
	"Why did the astronaut break up with his girlfriend? He needed space.",
	"Why don't oysters share their pearls? Because they're shellfish.",
	"Why did the cowboy get a wiener dog? He wanted to get a long little doggy.",
	"Why did the melon jump into the lake? It wanted to be a watermelon.",
	"Why did the student eat his homework? Because the teacher said it was a piece of cake.",
	"Why did the banana go to the doctor? Because it wasn't peeling well.",
}

func cmdJoke(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]
	s.ChannelMessageSend(m.ChannelID, joke)
}
