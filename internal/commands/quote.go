package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/rand"
)

var quotes = []string{
	"I cannot tell a lie.\n\t- George Washington",
	"Early to bed and early to rise makes a man healthy, wealthy, and wise.\n\t- Benjamin Franklin",
	"Life, liberty, and the pursuit of happiness.\n\t- Declaration of Independence",
	"Talk is cheap. Show me the code.\n\t- Linus Torvalds",
	"Stay hungry, stay foolish.\n\t- Steve Jobs",
	"The best way to predict the future is to invent it.\n\t- Alan Kay",
	"In the middle of difficulty lies opportunity.\n\t- Albert Einstein",
	"Knowledge speaks, but wisdom listens.\n\t- Jimi Hendrix",
	"Do or do not. There is no try.\n\t- Yoda",
	"Innovation distinguishes between a leader and a follower.\n\t- Steve Jobs",
	"Be yourself; everyone else is already taken.\n\t- Oscar Wilde",
	"Simplicity is the ultimate sophistication.\n\t- Leonardo da Vinci",
	"The journey of a thousand miles begins with one step.\n\t- Lao Tzu",
	"The only true wisdom is in knowing you know nothing.\n\t- Socrates",
	"Genius is one percent inspiration and ninety-nine percent perspiration.\n\t- Thomas Edison",
	"To err is human; to forgive, divine.\n\t- Alexander Pope",
	"Imagination is more important than knowledge.\n\t- Albert Einstein",
	"The secret of getting ahead is getting started.\n\t- Mark Twain",
	"I think, therefore I am.\n\t- René Descartes",
	"Fortune favors the bold.\n\t- Latin Proverb",
	"You miss 100%% of the shots you don't take.\n\t- Wayne Gretzky",
	"The only thing we have to fear is fear itself.\n\t- Franklin D. Roosevelt",
	"That's one small step for man, one giant leap for mankind.\n\t- Neil Armstrong",
	"Stay foolish to stay sane.\n\t- Maxime Lagacé",
	"Whatever you are, be a good one.\n\t- Abraham Lincoln",
	"The purpose of our lives is to be happy.\n\t- Dalai Lama",
	"Turn your wounds into wisdom.\n\t- Oprah Winfrey",
	"Embrace the glorious mess that you are.\n\t- Elizabeth Gilbert",
	"Impossible is for the unwilling.\n\t- John Keats",
	"No pressure, no diamonds.\n\t- Thomas Carlyle",
	"Dream big and dare to fail.\n\t- Norman Vaughan",
	"You can if you think you can.\n\t- George Reeves",
	"Action is the foundational key to all success.\n\t- Pablo Picasso",
	"What you do speaks so loudly that I cannot hear what you say.\n\t- Ralph Waldo Emerson",
	"Believe you can and you're halfway there.\n\t- Theodore Roosevelt",
	"I sexually Identify as an Attack Helicopter. Ever since I was a boy I dreamed of soaring over the oilfields dropping hot sticky loads on disgusting foreigners. People say to me that a person being a helicopter is Impossible and I'm fucking retarded but I don't care, I'm beautiful. I'm having a plastic surgeon install rotary blades, 30 mm cannons and AMG-114 Hellfire missiles on my body. From now on I want you guys to call me \"Apache\" and respect my right to kill from above and kill needlessly. If you can't accept me you're a heliphobe and need to check your vehicle privilege. Thank you for being so understanding.",
	"Did you ever hear the tragedy of Darth Plagueis The Wise? I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis was a Dark Lord of the Sith, so powerful and so wise he could use the Force to influence the midichlorians to create life… He had such a knowledge of the dark side that he could even keep the ones he cared about from dying. The dark side of the Force is a pathway to many abilities some consider to be unnatural. He became so powerful… the only thing he was afraid of was losing his power, which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew, then his apprentice killed him in his sleep. Ironic. He could save others from death, but not himself.",
}

func cmdQuote(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(uint64(time.Now().UnixNano()))
	quote := quotes[rand.Intn(len(quotes))]
	s.ChannelMessageSend(m.ChannelID, quote)
}
