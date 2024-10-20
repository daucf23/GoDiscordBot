package commands

import "github.com/bwmarrin/discordgo"

func cmdHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Commands:\n" +
		"!ping - Responds with 'pong' (wow, amazing)\n" +
		"!roll - Rolls a die\n" +
		"!serverinfo - Gives server info\n" +
		"!quote - Random quotes\n" +
		"!ask - Chat with low-effort-bot\n" +
		//"!weather [location] - Provides the current weather\n" +
		"!joke - Tells a random joke\n" +
		"!fact - Shares an interesting fact\n" +
		"!8ball [question] - Magic 8-Ball answers your question\n" +
		"!reverse [text] - Reverses the provided text\n" +
		"!mock [text] - Mocks the provided text\n"
		//"... and more!"

	s.ChannelMessageSend(m.ChannelID, helpMessage)
}
