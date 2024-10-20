package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/rand"
)

var facts = []string{
	"Honey never spoils. Archaeologists have found edible honey in ancient Egyptian tombs.",
	"Bananas are berries, but strawberries aren't.",
	"The Eiffel Tower can be 15 cm taller during hot days due to thermal expansion of the metal.",
	"Octopuses have three hearts.",
	"A day on Venus is longer than a year on Venus.",
	"There are more possible iterations of a game of chess than atoms in the observable universe.",
	"Wombat poop is cube-shaped.",
	"The world's smallest reptile was discovered in 2021 in Madagascar—a tiny chameleon.",
	"An adult human has fewer bones than a baby does.",
	"Some cats are allergic to humans.",
	"Scotland's national animal is the unicorn.",
	"The majority of your brain is fat.",
	"Kangaroos can't walk backwards.",
	"There are more stars in the universe than grains of sand on Earth.",
	"A group of flamingos is called a 'flamboyance'.",
	"The hottest temperature ever recorded on Earth was 134°F (56.7°C) in Death Valley, California.",
	"Polar bears have black skin under their white fur.",
	"Sloths can hold their breath longer than dolphins can.",
	"An octopus can fit through any hole larger than its beak.",
	"The first oranges weren't orange; they were green.",
	"There's a basketball court on the top floor of the U.S. Supreme Court Building known as the 'Highest Court in the Land.'",
	"Humans share approximately 50% of their DNA with bananas.",
	"The shortest war in history lasted just 38 minutes.",
	"Bees can make colored honey.",
	"There's a species of jellyfish that is considered biologically immortal.",
	"Cows have best friends and can become stressed when separated.",
	"Your nose and ears continue to grow throughout your life.",
	"The inventor of the Frisbee was turned into a Frisbee after he died.",
	"Dolphins have unique names for one another.",
	"A bolt of lightning contains enough energy to toast 100,000 slices of bread.",
	"Apples float in water because they are 25% air.",
	"The longest English word without a true vowel is 'rhythms.'",
	"In Switzerland, it's illegal to own just one guinea pig.",
	"Octopuses lay 56,000 eggs at a time.",
	"Humans are the only animals that blush.",
}

func cmdFact(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(uint64(time.Now().UnixNano()))
	fact := facts[rand.Intn(len(facts))]
	s.ChannelMessageSend(m.ChannelID, fact)
}
