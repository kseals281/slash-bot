package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type villager struct {
	description string
	birthday    string
	loves       []string
	likes       []string
	dislikes    []string
	hates       []string
}

var Villagers = map[string]villager{
	"abigail": {
		description: "Abigail is a villager who lives at Pierre's General Store in Pelican Town. She is one of the" +
			" twelve characters available to marry.",
		birthday: "Fall 13",
		loves: []string{"amethyst", "banana pudding", "blackberry cobbler", "chocolate cake", "pufferfish",
			"pumpkin", "spicy eel"},
		likes:    []string{"quartz"},
		dislikes: []string{"sugar", "wild horseradish"},
		hates:    []string{"clay", "holly"},
	},
	"alex": {
		description: "Alex is a villager who lives in the house southeast of Pierre's General Store. Alex is one of " +
			"the twelve characters available to marry.",
		birthday: "Summer 13",
		loves:    []string{"complete breakfast", "salmon dinner"},
		likes:    []string{"all eggs except void and dino"},
		dislikes: []string{"salmonberry", "wild horseradish"},
		hates:    []string{"quartz", "holly"},
	},
	"caroline": {
		description: "Caroline is married to Pierre and they live with their daughter Abigail in Pierre's General" +
			" Store.",
		birthday: "Summer 13",
		loves:    []string{"fish taco", "green tea", "summer spangle", "tropical curry"},
		likes:    []string{"daffodil", "tea leaves"},
		dislikes: []string{"amaranth", "chanterelle", "common mushroom", "dandelion", "duck mayonnaise", "ginger",
			"hazelnut", "holly", "leek", "magma cap", "mayonnaise", "morel", "purple mushroom", "snow yam",
			"wild horseradish", "wild root"},
		hates: []string{"quartz", "salmonberry"},
	},
	"clint": {
		description: "Clint is unavailable for shopping, tool upgrades, or geode processing on Fridays after the " +
			"Community Center is restored, unless it is raining.",
		birthday: "winter 26",
		loves: []string{"amethyst", "aquamarine", "artichoke dip", "emerald", "fiddlehead risotto", "gold bar",
			"iridium bar", "jade", "omni geode", "ruby", "topaz"},
		likes:    []string{"copper bar", "iron bar"},
		dislikes: []string{"quartz", "salmonberry", "wild horseradish"},
		hates:    []string{"holly"},
	},
	"demetrius": {
		description: "Demetrius is a scientist who studies the valley's local wildlife. He can often be found working" +
			" in his laboratory or outdoors taking notes.",
		birthday: "summer 19",
		loves:    []string{"bean hotpot", "ice cream", "rice pudding", "strawberry"},
		likes:    []string{"purple mushroom"},
		dislikes: []string{"quartz"},
		hates:    []string{"holly"},
	},
	"dwarf": {
		description: "Initially the way is blocked off by an unbreakable rock, just inside the entrance. After " +
			"upgrading to a steel pickaxe the stone can be broken. A cherry bomb also works..",
		birthday: "summer 22",
		loves:    []string{"amethyst", "aquamarine", "emerald", "jade", "lemon stone", "omni geode", "ruby", "topaz"},
		likes:    []string{"cave carrot", "quartz"},
		dislikes: []string{"chanterelle", "common mushroom", "daffodil", "dandelion", "ginger",
			"hazelnut", "holly", "leek", "magma cap", "morel", "purple mushroom", "snow yam",
			"wild horseradish", "winter root"},
		hates: []string{"All universal hates"},
	},
	"elliot": {
		description: "Elliott lives alone in a cabin on the beach near Willy's Fish Shop. He is friends with Willy " +
			"and Leah.",
		birthday: "summer 22",
		loves:    []string{"crab cakes", "duck feather", "lobster", "pomegranate", "squid ink", "tom kha soup"},
		likes:    []string{"octopus", "squid"},
		dislikes: []string{"chanterelle", "common mushroom", "daffodil", "dandelion", "ginger",
			"hazelnut", "holly", "leek", "magma cap", "morel", "pizza", "purple mushroom", "snow yam",
			"wild horseradish", "winter root"},
		hates: []string{"amaramth", "quartz", "salmonberry", "sea cucumber"},
	},
	"emily": {
		description: "Emily loves to make her own clothing, but fabric can be difficult to come by in town. Among her" +
			" favorite gifts are cloth and wool.",
		birthday: "spring 27",
		loves: []string{"amethyst", "aquamarine", "cloth", "emerald", "jade", "ruby", "survival burger",
			"topaz", "wool"},
		likes:    []string{"daffodil", "quartz"},
		dislikes: []string{"fried eel", "ice cream cone", "rice pudding", "salmonberry", "spicy eel"},
		hates:    []string{"fish taco", "holly", "maki roll", "salmon dinner", "sashimi"},
	},
	"evelyn": {
		description: "On the 2nd of every season, she has an appointment at the clinic, and on the 23rd of every " +
			"season she accompanies George to his appointment.",
		birthday: "winter 20",
		loves:    []string{"beet", "chocolate cake", "diamond", "fairy rose", "stuffing", "tulip"},
		likes:    []string{"daffodil"},
		dislikes: []string{"quartz", "wild horseradish"},
		hates: []string{"clam", "clay", "coral", "fried eel", "garlic", "holly", "maki roll", "salmonberry",
			"sashimi", "spice berry", "spicy eel", "trout soup"},
	},
	"george": {
		description: "George is married to Evelyn and lives with his grandson, Alex. Alex's deceased mother Clara was" +
			" his daughter.",
		birthday: "fall 24",
		loves:    []string{"fried mushrooms", "leek"},
		likes:    []string{"daffodil"},
		dislikes: []string{"salmonberry", "wild horseradish"},
		hates:    []string{"clay", "dandelion", "holly", "quartz"},
	},
	"gus": {
		description: "If the player has unlocked the beach resort on Ginger Island, Gus will leave on random days to" +
			" tend the resort's bar. A variety of alcohol items and the recipe for Tropical Curry can be purchased.",
		birthday: "summer 8",
		loves:    []string{"diamond", "escargot", "fish taco", "orange", "tropical curry"},
		likes:    []string{"daffodil"},
		dislikes: []string{"salmonberry", "wild horseradish"},
		hates:    []string{"coleslaw", "holly", "quartz"},
	},
	"haley": {
		description: "Her behavior changes if it's raining or snowing outside. She will not go to the fountain " +
			"Tuesday-Sunday if it is raining.",
		birthday: "spring 14",
		loves:    []string{"coconut", "fruit salad", "pink cake", "sunflower"},
		likes:    []string{"daffodil"},
		dislikes: []string{"chanterelle", "common mushroom", "dandelion", "ginger", "hazelnut", "holly", "leek",
			"magma cap", "morel", "purple mushroom", "quartz", "snow yam", "wild horseradish", "winter root"},
		hates: []string{"clay", "prismatic shard", "wild horseradish"},
	},
	"harvey": {
		description: "He runs the town's medical clinic and is passionate about the health of the townsfolk. " +
			"He's one of the twelve characters available to marry.",
		birthday: "winter 14",
		loves:    []string{"coffee", "pickles", "super meal", "truffle oil", "wine"},
		likes: []string{"chanterelle", "common mushroom", "daffodil", "dandelion", "duck egg", "duck feather",
			"ginger", "goat milk", "hazelnut", "holly", "large goat milk", "leek", "magma cap", "morel",
			"purple mushroom", "quartz", "snow yam", "spring onion", "wild horseradish", "winter root"},
		dislikes: []string{"blueberry tart", "bread", "cheese", "chocolate cake", "cookie", "cranberry sauce",
			"fried mushroom", "glazed yams", "goat cheese", "hashbrowns", "ice cream", "pancakes", "pink cake", "pizza",
			"rhubarb pie", "rice pudding"},
		hates: []string{"coral", "nautilus shell", "rainbow shell", "salmonberry", "spice berry"},
	},
	"jas": {
		description: "She's a young girl, and can often be found with her best friend Vincent.",
		birthday:    "summer 4",
		loves:       []string{"fairy rose", "pink cake", "plum pudding"},
		likes:       []string{"coconut", "daffodil"},
		dislikes: []string{"chanterelle", "common mushroom", "dandelion", "ginger", "hazelnut", "holly", "leek",
			"magma cap", "morel", "purple mushroom", "quartz", "snow yam", "winter root"},
		hates: []string{"clay", "pina colada", "triple shot espresso", "wild horseradish"},
	},
	"jodi": {
		description: "She lives at 1 Willow Lane with her husband Kent (who is away serving in the military until the " +
			"first of Spring Year 2) and two sons, Sam and Vincent.",
		birthday: "fall 11",
		loves: []string{"chocolate cake", "crispy bass", "diamond", "eggplant parmesan", "fried eel", "pancakes",
			"rhubarb pie", "vegetable medley"},
		likes: []string{"all eggs except void", "all milk"},
		dislikes: []string{"chanterelle", "common mushroom", "garlic", "ginger", "hazelnut", "holly", "leek",
			"magma cap", "morel", "purple mushroom", "quartz", "snow yam", "wild horseradish", "winter root"},
		hates: []string{"daffodil", "dandelion", "spice berry"},
	},
	"robin": {
		description: "Robin is a villager who resides at 24 Mountain Road, on The Mountain, with her husband " +
			"Demetrius, daughter Maru, and son Sebastian.",
		birthday: "Fall 21",
		loves:    []string{"goat cheese", "peach", "spaghetti"},
		likes:    []string{"hardwood", "quartz"},
		dislikes: []string{"wild horseradish"},
		hates:    []string{"holly"},
	},
}

func createEmbedGift(title, url, thumbnail string) *discordgo.MessageEmbed {
	embed := new(discordgo.MessageEmbed)
	lovesField, hatesField := new(discordgo.MessageEmbedField), new(discordgo.MessageEmbedField)
	likesField, dislikesField := new(discordgo.MessageEmbedField), new(discordgo.MessageEmbedField)
	lovesField.Name, hatesField.Name = "Loves", "Hates"
	likesField.Name, dislikesField.Name = "Likes", "Dislikes"

	embed.Title = capitalize(title)
	embed.URL = url
	embed.Description = Villagers[title].description
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: thumbnail}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("%s's birthday is %s", capitalize(title), Villagers[title].birthday),
	}

	for _, love := range Villagers[title].loves {
		lovesField.Value += capitalize(love) + "\n"
	}
	for _, like := range Villagers[title].likes {
		likesField.Value += capitalize(like) + "\n"
	}
	for _, dislike := range Villagers[title].dislikes {
		dislikesField.Value += capitalize(dislike) + "\n"
	}
	for _, hate := range Villagers[title].hates {
		hatesField.Value += capitalize(hate) + "\n"
	}

	embed.Fields = []*discordgo.MessageEmbedField{lovesField, likesField, dislikesField, hatesField}

	return embed
}
