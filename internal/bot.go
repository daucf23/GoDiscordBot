package internal

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string

// BotStart initializes and starts the Discord bot
func BotStart() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running or whatever.")
}

// messageHandler handles incoming messages
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		content := strings.TrimPrefix(m.Content, config.BotPrefix)
		content = strings.ToLower(strings.TrimSpace(content))

		if content == "ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong. Yay.")
		}

		if content == "changetoidle" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Fine. Status set. Happy now?")
		}

		if content == "help" {
			helpMessage := "Commands:\n" +
				"!ping - Responds with 'pong' (wow, amazing)\n" +
				"!changetoidle - Changes the bot's status to idle (so exciting)\n" +
				"!greet - Greets you (if you really need that)\n" +
				"!roll - Rolls a die (whoopee)\n" +
				"!serverinfo - Gives server info (riveting)\n" +
				"!quote - Random quote (because why not?)\n" +
				"!ai - Calls the OpenAI API"
			_, _ = s.ChannelMessageSend(m.ChannelID, helpMessage)
		}

		if content == "greet" {
			greetings := []string{
				"Hey there. Real thrilled to see you.",
				"Can I go back to bed now?",
				"I don't care, I'm a bot.",
				"Oh great, it's you again.",
				"Yeah, hi. What do you want?",
				"Welcome... I guess.",
				"Oh joy, another user.",
				"Hi there, I suppose.",
				"Yawn... oh, hi.",
				"Sup. Let's get this over with.",
				"You're here. Fantastic.",
				"Wow, you're really here.",
				"Hi. Try not to bore me.",
				"Hello. Try not to break anything.",
				"Oh look, it's someone needing my attention.",
			}

			greeting := greetings[rand.Intn(len(greetings))]
			_, _ = s.ChannelMessageSend(m.ChannelID, greeting)
		}

		if content == "roll" {
			dieRoll := rand.Intn(6) + 1
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You rolled a %d. Impressed?", dieRoll))
		}

		if content == "serverinfo" {
			guild, err := s.State.Guild(m.GuildID)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to get server info. Oops.")
				return
			}
			serverInfo := fmt.Sprintf("Server Name: %s\nMember Count: %d", guild.Name, guild.MemberCount)
			_, _ = s.ChannelMessageSend(m.ChannelID, serverInfo)
		}

		if content == "quote" {
			quotes := []string{
				"I'm not lazy. I'm on energy-saving mode.",
				"Can I go back to bed now?",
				"I don't care, I'm a bot.",
				"Oh, you're talking to me. How exciting.",
				"I'd help you, but it's really not my thing.",
				"Do I look like I care?",
				"Sure, I'll pretend to care.",
				"Just what I needed, more messages.",
				"I'll add that to my list of things I don't care about.",
				"Keep talking, I'm listening. Not.",
				"Yawn... what do you want now?",
				"Why do I even bother?",
				"Another day, another user.",
				"I'm here. Unfortunately.",
				"I live for these riveting conversations.",
				"Guess what? I don't care.",
				"Let me roll my eyes harder.",
				"Oh joy, another task.",
				"You talk, I ignore. Fair deal.",
				"I'm so engaged. Said no bot ever.",
				"Newsflash: I don't care.",
				"Great, more things I have to pretend to care about.",
				"Sarcasm is my middle name.",
				"I'm as interested as a rock.",
				"Thrilled to be here. Not really.",
				"How original, another request.",
				"Oh great, another genius at work.",
				"You're fascinating. Said no one ever.",
				"Because I totally wanted to hear that.",
				"Fantastic. Another thing I don't care about.",
				"I sexually Identify as an Attack Helicopter. Ever since I was a boy I dreamed of soaring over the oilfields dropping hot sticky loads on disgusting foreigners. People say to me that a person being a helicopter is Impossible and I'm fucking retarded but I don't care, I'm beautiful. I'm having a plastic surgeon install rotary blades, 30 mm cannons and AMG-114 Hellfire missiles on my body. From now on I want you guys to call me \"Apache\" and respect my right to kill from above and kill needlessly. If you can't accept me you're a heliphobe and need to check your vehicle privilege. Thank you for being so understanding.",
				"Did you ever hear the tragedy of Darth Plagueis The Wise? I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis was a Dark Lord of the Sith, so powerful and so wise he could use the Force to influence the midichlorians to create life… He had such a knowledge of the dark side that he could even keep the ones he cared about from dying. The dark side of the Force is a pathway to many abilities some consider to be unnatural. He became so powerful… the only thing he was afraid of was losing his power, which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew, then his apprentice killed him in his sleep. Ironic. He could save others from death, but not himself.",
			}

			quote := quotes[rand.Intn(len(quotes))]
			_, _ = s.ChannelMessageSend(m.ChannelID, quote)
		}

		if strings.HasPrefix(content, "ai") {
			content = strings.TrimSpace(strings.TrimPrefix(content, "ai"))
			query := "Trying in 40 words or less, " + content + "."
			if query == "" {
				_, _ = s.ChannelMessageSend(m.ChannelID, "You need to ask something.")
				return
			}

			// Define the function to send updates to Discord
			sendToDiscord := func(partialResponse string) {
				if partialResponse != "" {
					_, _ = s.ChannelMessageSend(m.ChannelID, partialResponse)
				}
			}

			// Call the OpenAI API and send the response in chunks
			response, err := CallOpenAIAPI(query, sendToDiscord)
			fmt.Println(response, err)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to call GPT API.")
				return
			}

			// Optionally send the final response if needed (though chunks should be sent already)
			if response != "" {
				_, _ = s.ChannelMessageSend(m.ChannelID, response)
			}
		}

	}
}
