package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type shop struct {
	hours string
	stock []item
}

type item struct {
	name      string
	cost      int
	materials map[string]int
}

var Shops = map[string]shop{
	"robin": {
		hours: "9am - 5pm except Tuesdays",
		stock: []item{
			{name: "wood", cost: 50},
			{name: "stone", cost: 100},
			{name: "workbench", cost: 2000},
			{name: "telephone", cost: 2000},
			{name: "mini-fridge", cost: 3000},
			{name: "furniture catalogue", cost: 200000},
			{name: "barn", cost: 6000, materials: map[string]int{
				"wood":  350,
				"stone": 150,
			}},
			{name: "big barn", cost: 12000, materials: map[string]int{
				"wood":  450,
				"stone": 200,
			}},
			{name: "deluxe barn", cost: 25000, materials: map[string]int{
				"wood":  550,
				"stone": 300,
			}},
			{name: "coop", cost: 4000, materials: map[string]int{
				"wood":  300,
				"stone": 100,
			}},
			{name: "big Coop", cost: 10000, materials: map[string]int{
				"wood":  400,
				"stone": 150,
			}},
			{name: "deluxe Coop", cost: 20000, materials: map[string]int{
				"wood":  500,
				"stone": 200,
			}},
			{name: "fish Pond", cost: 5000, materials: map[string]int{
				"stone":       200,
				"seaweed":     5,
				"green algae": 5,
			}},
			{name: "shed", cost: 15000, materials: map[string]int{
				"stone": 300,
			}},
			{name: "big Shed", cost: 20000, materials: map[string]int{
				"wood":  550,
				"stone": 300,
			}},
			{name: "silo", cost: 100, materials: map[string]int{
				"stone":      100,
				"clay":       10,
				"copper bar": 5,
			}},
			{name: "community Upgrade 1", cost: 500000, materials: map[string]int{
				"wood": 950,
			}},
			{name: "community Upgrade 2", cost: 300000},
		},
	},

	"marnie": {
		hours: "9am - 4pm, Wednesday - Sunday",
		stock: []item{
			{name: "hay", cost: 50},
			{name: "heater", cost: 2000},
			{name: "auto-grabber", cost: 25000},
			{name: "chicken", cost: 800},
			{name: "cow", cost: 1500},
			{name: "goat", cost: 4000},
			{name: "duck", cost: 1200},
			{name: "sheep", cost: 8000},
			{name: "rabbit", cost: 8000},
			{name: "pig", cost: 16000},
		},
	},

	"pierre-general": {
		hours: "9am - 5pm Everyday",
		stock: []item{
			{name: "large pack", cost: 2000},
			{name: "deluxe pack", cost: 10000},
			{name: "grass starter", cost: 100},
			{name: "sugar", cost: 100},
			{name: "wheat flour", cost: 100},
			{name: "rice", cost: 200},
			{name: "oil", cost: 200},
			{name: "vinegar", cost: 200},
			{name: "quality fertilizer", cost: 150},
			{name: "deluxe speed-gro", cost: 150},
			{name: "cherry sapling", cost: 3400},
			{name: "apricot sapling", cost: 2000},
			{name: "orange sapling", cost: 4000},
			{name: "peach sapling", cost: 6000},
			{name: "pomegranate sapling", cost: 6000},
			{name: "apple sapling", cost: 4000},
		},
	},
	"pierre-spring": {
		hours: "9am - 5pm Everyday",
		stock: []item{
			{name: "parsnip seeds", cost: 20},
			{name: "bean starter", cost: 60},
			{name: "cauliflower seeds", cost: 80},
			{name: "potato seeds", cost: 50},
			{name: "tulip bulb", cost: 20},
			{name: "kale seeds", cost: 70},
			{name: "jazz seeds", cost: 30},
			{name: "garlic seeds", cost: 40},
			{name: "rice shoot", cost: 40},
		},
	},
	"pierre-summer": {
		hours: "9am - 5pm Everyday",
		stock: []item{
			{name: "melon seeds", cost: 80},
			{name: "blueberry seeds", cost: 80},
			{name: "pepper seeds", cost: 40},
			{name: "wheat seeds", cost: 10},
			{name: "radish seeds", cost: 40},
			{name: "poppy seeds", cost: 100},
			{name: "spangle seeds", cost: 50},
			{name: "hops starter", cost: 60},
			{name: "corn seeds", cost: 150},
			{name: "sunflower seeds", cost: 200},
			{name: "red cabbage seeds", cost: 100},
		},
	},
	"pierre-fall": {
		hours: "9am - 5pm Everyday",
		stock: []item{
			{name: "eggplant seeds", cost: 20},
			{name: "corn seeds", cost: 150},
			{name: "pumpkin seeds", cost: 100},
			{name: "bok choy seeds", cost: 50},
			{name: "yam seeds", cost: 60},
			{name: "cranberry seeds", cost: 240},
			{name: "sunflower seeds", cost: 200},
			{name: "fairy seeds", cost: 200},
			{name: "amaranth seeds", cost: 70},
			{name: "grape starter", cost: 60},
			{name: "wheat seeds", cost: 10},
			{name: "artichoke seeds", cost: 30},
		},
	},

	"clint": {
		hours: "9am - 4pm except Friday",
		stock: []item{
			{name: "copper ore", cost: 150},
			{name: "iron ore", cost: 250},
			{name: "coal", cost: 250},
			{name: "gold ore", cost: 750},
			{name: "copper tool upgrade", cost: 2000, materials: map[string]int{
				"copper bar": 5,
			}},
			{name: "iron tool upgrade", cost: 5000, materials: map[string]int{
				"iron bar": 5,
			}},
			{name: "gold tool upgrade", cost: 10000, materials: map[string]int{
				"gold bar": 5,
			}},
			{name: "iridium tool upgrade", cost: 25000, materials: map[string]int{
				"iridium bar": 5,
			}},
			{name: "copper trash can", cost: 1000, materials: map[string]int{
				"copper bar": 5,
			}},
			{name: "iron trash can", cost: 2500, materials: map[string]int{
				"iron bar": 5,
			}},
			{name: "gold trash can", cost: 5000, materials: map[string]int{
				"gold bar": 5,
			}},
			{name: "iridium trash can", cost: 12500, materials: map[string]int{
				"iridium bar": 5,
			}},
		},
	},

	"wizard": {
		hours: "6am to 11pm Everyday",
		stock: []item{
			{name: "junimo hut", cost: 20000, materials: map[string]int{
				"stone":     200,
				"starfruit": 9,
				"fiber":     100,
			}},
			{name: "earth obelisk", cost: 500000, materials: map[string]int{
				"iridium bar":   10,
				"earth crystal": 10,
			}},
			{name: "water obelisk", cost: 500000, materials: map[string]int{
				"iridium bar": 5,
				"clam":        10,
				"coral":       10,
			}},
			{name: "desert obelisk", cost: 1000000, materials: map[string]int{
				"iridium bar":  20,
				"coconut":      10,
				"cactus fruit": 10,
			}},
			{name: "island obelisk", cost: 1000000, materials: map[string]int{
				"iridium bar":  10,
				"banana":       10,
				"dragon tooth": 10,
			}},
			{name: "gold clock", cost: 10000000},
		},
	},

	"saloon": {
		hours: "12pm to 12am Everyday",
		stock: []item{
			{name: "beer", cost: 400},
			{name: "salad", cost: 220},
			{name: "bread", cost: 120},
			{name: "spaghetti", cost: 240},
			{name: "pizza", cost: 600},
			{name: "coffee", cost: 300},
			{name: "crab cakes", cost: 550},
		},
	},

	"fish": {
		hours: "9am to 5pm Everyday",
		stock: []item{
			{name: "trout soup", cost: 250},
			{name: "bait", cost: 5},
			{name: "crab pot", cost: 1500},
			{name: "spinner", cost: 500},
			{name: "trap bobber", cost: 500},
			{name: "lead bobber", cost: 200},
			{name: "treasure hunter", cost: 750},
			{name: "cork bobber", cost: 750},
			{name: "barbed hook", cost: 1000},
			{name: "dressed spinner", cost: 1000},
			{name: "magnet", cost: 1000},
			{name: "iridium rod", cost: 7500},
		},
	},

	"joja-general": {
		hours: "9am - 11pm",
		stock: []item{
			{name: "joja cola", cost: 75},
			{name: "grass starter", cost: 125},
			{name: "sugar", cost: 125},
			{name: "wheat flour", cost: 125},
			{name: "rice", cost: 250},
			{name: "auto-Petter", cost: 50000},
		},
	},
	"joja-spring": {
		hours: "9am - 11pm",
		stock: []item{
			{name: "parsnip seeds", cost: 25},
			{name: "bean starter", cost: 75},
			{name: "cauliflower seeds", cost: 100},
			{name: "potato seeds", cost: 62},
			{name: "tulip bulb", cost: 25},
			{name: "kale seeds", cost: 87},
			{name: "jazz seeds", cost: 37},
			{name: "garlic seeds", cost: 40},
			{name: "rice shoot", cost: 40},
		},
	},
	"joja-summer": {
		hours: "9am - 11pm",
		stock: []item{
			{name: "tomato seeds", cost: 62},
			{name: "pepper seeds", cost: 50},
			{name: "wheat seeds", cost: 12},
			{name: "melon seeds", cost: 100},
			{name: "hops starter", cost: 75},
			{name: "poppy seeds", cost: 125},
			{name: "spangle seeds", cost: 62},
			{name: "sunflower seeds", cost: 125},
		},
	},
	"joja-fall": {
		hours: "9am - 11pm",
		stock: []item{
			{name: "corn seeds", cost: 187},
			{name: "eggplant seeds", cost: 25},
			{name: "pumpkin seeds", cost: 125},
			{name: "amaranth seeds", cost: 87},
			{name: "grape starter", cost: 75},
			{name: "yam seeds", cost: 75},
			{name: "bok choy seeds", cost: 62},
			{name: "cranberry seeds", cost: 300},
			{name: "sunflower seeds", cost: 125},
			{name: "fairy seeds", cost: 250},
			{name: "wheat seeds", cost: 12},
		},
	},

	"oasis": {
		hours: "9am - 11pm",
		stock: []item{
			{name: "cactus seeds", cost: 150},
			{name: "rhubarb seeds", cost: 100},
			{name: "starfruit seeds", cost: 400},
			{name: "beet seeds", cost: 20},
			{name: "seasonal plant", cost: 100},
		},
	},

	"adventurers-guild": {
		hours: "2pm - 10pm",
		stock: []item{
			{name: "bone sword", cost: 6000, materials: map[string]int{"the mines": 75}},
			{name: "lava katana", cost: 25000, materials: map[string]int{"the mines": 120}},
			{name: "galaxy sword", cost: 50000, materials: map[string]int{"desert prismatic shard": 1}},
			{name: "galaxy dagger", cost: 35000, materials: map[string]int{"desert prismatic shard": 1}},
			{name: "galaxy hammer", cost: 75000, materials: map[string]int{"desert prismatic shard": 1}},
			{name: "slime charmer ring", cost: 25000, materials: map[string]int{"slimes": 1000}},
			{name: "savage ring", cost: 25000, materials: map[string]int{"void spirits": 150}},
			{name: "burglar's ring", cost: 20000, materials: map[string]int{"dust sprites": 500}},
			{name: "vampire ring", cost: 15000, materials: map[string]int{"bat": 200}},
			{name: "crabshell ring", cost: 15000, materials: map[string]int{"rock crabs": 60}},
			{name: "napalm ring", cost: 30000, materials: map[string]int{"serpents": 250}},
		},
	},
}

func createEmbedShop(title, owner, url, note string, author *discordgo.MessageEmbedAuthor) *discordgo.MessageEmbed {
	embed := new(discordgo.MessageEmbed)
	materials := false
	itemName, itemCost, itemMaterials := new(discordgo.MessageEmbedField), new(discordgo.MessageEmbedField), new(discordgo.MessageEmbedField)
	itemName.Name, itemCost.Name, itemMaterials.Name = "Name", "Cost", "Materials"
	itemName.Inline, itemCost.Inline, itemMaterials.Inline = true, true, true

	embed.Title = title
	embed.URL = url
	embed.Description = Shops[owner].hours
	embed.Author = author
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: note,
	}
	for _, itm := range Shops[owner].stock {
		itemName.Value += capitalize(itm.name) + "\n"
		itemCost.Value += strconv.Itoa(itm.cost) + "g\n"
		if len(itm.materials) == 0 {
			itemMaterials.Value += "None"
		}
		count := 0
		for mat, num := range itm.materials {
			materials = true
			itemMaterials.Value += capitalize(mat) + " (" + strconv.Itoa(num) + ")"
			if count == len(itm.materials)-1 {
				break
			}
			itemMaterials.Value += "  -  "
			count++
		}
		itemMaterials.Value += "\n"
	}
	if materials {
		embed.Fields = []*discordgo.MessageEmbedField{itemName, itemCost, itemMaterials}
	} else {
		embed.Fields = []*discordgo.MessageEmbedField{itemName, itemCost}
	}
	return embed
}
