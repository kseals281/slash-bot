package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"time"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
		{
			Name:        "calculate",
			Description: "Run simple calculation on two numbers",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "first",
					Description: "First Number",
					Required:    true,
				}, {
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "operator",
					Description: "+ OR - OR * OR /",
					Required:    true,
				}, {
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "second",
					Description: "Second Number",
					Required:    true,
				},
			},
		},
		{
			Name:        "options",
			Description: "Command for demonstrating options",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "String option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "integer-option",
					Description: "Integer option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "bool-option",
					Description: "Boolean option",
					Required:    true,
				},

				// Required options must be listed first since optional parameters
				// always come after when they're used.
				// The same concept applies to Discord's Slash-commands API

				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role-option",
					Description: "Role option",
					Required:    false,
				},
			},
		},
		{
			Name:        "subcommands",
			Description: "Subcommands and command groups example",
			Options: []*discordgo.ApplicationCommandOption{
				// When a command has subcommands/subcommand groups
				// It must not have top-level options, they aren't accesible in the UI
				// in this case (at least not yet), so if a command has
				// subcommands/subcommand any groups registering top-level options
				// will cause the registration of the command to fail

				{
					Name:        "scmd-grp",
					Description: "Subcommands group",
					Options: []*discordgo.ApplicationCommandOption{
						// Also, subcommand groups aren't capable of
						// containing options, by the name of them, you can see
						// they can only contain subcommands
						{
							Name:        "nst-subcmd",
							Description: "Nested subcommand",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommandGroup,
				},
				// Also, you can create both subcommand groups and subcommands
				// in the command at the same time. But, there's some limits to
				// nesting, count of subcommands (top level and nested) and options.
				// Read the intro of slash-commands docs on Discord dev portal
				// to get more information
				{
					Name:        "subcmd",
					Description: "Top-level subcommand",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "responses",
			Description: "Interaction responses testing initiative",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "resp-type",
					Description: "Response type",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Channel message with source",
							Value: 4,
						},
						{
							Name:  "Deferred response With Source",
							Value: 5,
						},
					},
					Required: true,
				},
			},
		},
		{
			Name:        "followups",
			Description: "Followup messages",
		},
		{
			Name:        "stardew",
			Description: "All stardew related commands",
			Options: []*discordgo.ApplicationCommandOption{
				// When a command has subcommands/subcommand groups
				// It must not have top-level options, they aren't accesible in the UI
				// in this case (at least not yet), so if a command has
				// subcommands/subcommand any groups registering top-level options
				// will cause the registration of the command to fail

				{
					Name:        "shops",
					Description: "Shops around the valley.",
					Options: []*discordgo.ApplicationCommandOption{
						// Also, subcommand groups aren't capable of
						// containing options, by the name of them, you can see
						// they can only contain subcommands
						{
							Name:        "robin",
							Description: "Robin's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "marnie",
							Description: "Marnie's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "clint",
							Description: "Clint's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "wizard",
							Description: "Wizard's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "saloon",
							Description: "The Stardrop Saloon's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "fish",
							Description: "Fish Shop's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "pierre-general",
							Description: "Pierre's general shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "pierre-spring",
							Description: "Pierre's spring shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "pierre-summer",
							Description: "Pierre's summer shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "pierre-fall",
							Description: "Pierre's fall shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "joja-general",
							Description: "Joja Mart's general shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "joja-spring",
							Description: "Joja Mart's spring shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "joja-summer",
							Description: "Joja Mart's summer shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "joja-fall",
							Description: "Joja Mart's fall shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "oasis",
							Description: "Oasis's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "adventurers-guild",
							Description: "Oasis's shop info and schedule",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommandGroup,
				}, {
					Name:        "gifts",
					Description: "Gift info for any giftable villager.",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "abigail",
							Description: "Abigail is a villager who lives at Pierre's General Store in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name: "alex",
							Description: "Alex is a villager who lives in the house southeast of Pierre's General" +
								" Store.",
							Type: discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "caroline",
							Description: "Caroline is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "clint",
							Description: "Clint is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "demetrius",
							Description: "Demetrius is a villager who resides at 24 Mountain Road, on The Mountain.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "dwarf",
							Description: "The Dwarf is a valley resident who lives in The Mines.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "elliott",
							Description: "Elliott is a villager who lives on the beach south of Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "emily",
							Description: "Emily is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "evelyn",
							Description: "Evelyn is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "george",
							Description: "George is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "gus",
							Description: "Gus is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "haley",
							Description: "Haley is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "harvey",
							Description: "Harvey is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "jas",
							Description: "Jas is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "jodi",
							Description: "Jodi is a villager who lives in Pelican Town.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						}, {
							Name:        "robin",
							Description: "Robin is the town carpenter.",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommandGroup,
				},
				// Also, you can create both subcommand groups and subcommands
				// in the command at the same time. But, there's some limits to
				// nesting, count of subcommands (top level and nested) and options.
				// Read the intro of slash-commands docs on Discord dev portal
				// to get more information
				{
					Name:        "farming",
					Description: "Top-level subcommand",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				}, {
					Name:        "fishing",
					Description: "Top-level subcommand",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
		"calculate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			num1 := i.Data.Options[0].IntValue()
			op := i.Data.Options[1].StringValue()
			num2 := i.Data.Options[2].IntValue()

			var result string

			switch op {
			case "+":
				result = strconv.FormatInt(num1+num2, 10)
			case "-":
				result = strconv.FormatInt(num1-num2, 10)
			case "*":
				result = strconv.FormatInt(num1*num2, 10)
			case "/":
				result = strconv.FormatInt(num1/num2, 10)
			default:
				result = "Unknown operator given"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: string(result),
				},
			})
		},
		"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			margs := []interface{}{
				// Here we need to convert raw interface{} value to wanted type.
				// Also, as you can see, here is used utility functions to convert the value
				// to particular type. Yeah, you can use just switch type,
				// but this is much simpler
				i.Data.Options[0].StringValue(),
				i.Data.Options[1].IntValue(),
				i.Data.Options[2].BoolValue(),
			}
			msgformat :=
				` Now you just learned how to use command options. Take a look to the value of which you've just entered:
				> string_option: %s
				> integer_option: %d
				> bool_option: %v
`
			if len(i.Data.Options) >= 4 {
				margs = append(margs, i.Data.Options[3].ChannelValue(nil).ID)
				msgformat += "> channel-option: <#%s>\n"
			}
			if len(i.Data.Options) >= 5 {
				margs = append(margs, i.Data.Options[4].UserValue(nil).ID)
				msgformat += "> user-option: <@%s>\n"
			}
			if len(i.Data.Options) >= 6 {
				margs = append(margs, i.Data.Options[5].RoleValue(nil, "").ID)
				msgformat += "> role-option: <@&%s>\n"
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, we'll discuss them in "responses" part
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
		"subcommands": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			content := ""

			// As you can see, the name of subcommand (nested, top-level) or subcommand group
			// is provided through arguments.
			switch i.Data.Options[0].Name {
			case "subcmd":
				content =
					"The top-level subcommand is executed. Now try to execute the nested one."
			default:
				if i.Data.Options[0].Name != "scmd-grp" {
					return
				}
				switch i.Data.Options[0].Options[0].Name {
				case "nst-subcmd":
					content = "Nice, now you know how to execute nested commands too"
				default:
					// I added this in the case something might go wrong
					content = "Oops, something gone wrong.\n" +
						"Hol' up, you aren't supposed to see this message."
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: content,
				},
			})
		},
		"responses": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Responses to a command are very important.
			// First of all, because you need to react to the interaction
			// by sending the response in 3 seconds after receiving, otherwise
			// interaction will be considered invalid and you can no longer
			// use the interaction token and ID for responding to the user's request

			content := ""
			// As you can see, the response type names used here are pretty self-explanatory,
			// but for those who want more information see the official documentation
			switch i.Data.Options[0].IntValue() {
			case int64(discordgo.InteractionResponseChannelMessageWithSource):
				content =
					"You just responded to an interaction, sent a message and showed the original one. " +
						"Congratulations!"
				content +=
					"\nAlso... you can edit your response, wait 5 seconds and this message will be changed"
			default:
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseType(i.Data.Options[0].IntValue()),
				})
				if err != nil {
					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
						Content: "Something went wrong",
					})
				}
				return
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseType(i.Data.Options[0].IntValue()),
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: content,
				},
			})
			if err != nil {
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
				return
			}
			time.AfterFunc(time.Second*5, func() {
				err = s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
					Content: content + "\n\nWell, now you know how to create and edit responses. " +
						"But you still don't know how to delete them... so... wait 10 seconds and this " +
						"message will be deleted.",
				})
				if err != nil {
					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
						Content: "Something went wrong",
					})
					return
				}
				time.Sleep(time.Second * 10)
				s.InteractionResponseDelete(s.State.User.ID, i.Interaction)
			})
		},
		"followups": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Followup messages are basically regular messages (you can create as many of them as you wish)
			// but work as they are created by webhooks and their functionality
			// is for handling additional messages after sending a response.

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					// Note: this isn't documented, but you can use that if you want to.
					// This flag just allows you to create messages visible only for the caller of the command
					// (user who triggered the command)
					Flags:   1 << 6,
					Content: "Surprise!",
				},
			})
			msg, err := s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: "Followup message has been created, after 5 seconds it will be edited",
			})
			if err != nil {
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
				return
			}
			time.Sleep(time.Second * 5)

			s.FollowupMessageEdit(s.State.User.ID, i.Interaction, msg.ID, &discordgo.WebhookEdit{
				Content: "Now the original message is gone and after 10 seconds this message will ~~self-destruct~~ be deleted.",
			})

			time.Sleep(time.Second * 10)

			s.FollowupMessageDelete(s.State.User.ID, i.Interaction, msg.ID)

			s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
				Content: "For those, who didn't skip anything and followed tutorial along fairly, " +
					"take a unicorn :unicorn: as reward!\n" +
					"Also, as bonus... look at the original interaction response :D",
			})
		},
		"stardew": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			embed := new(discordgo.MessageEmbed)

			// As you can see, the name of subcommand (nested, top-level) or subcommand group
			// is provided through arguments.
			switch i.Data.Options[0].Name {
			case "farming":
				embed.Description =
					"The top-level subcommand is executed. Now try to execute the nested one."
			case "fishing":
				embed.Description =
					"Stardew fishing info."
			case "shops":
				switch i.Data.Options[0].Options[0].Name {

				case "robin":
					title := "Carpenter's shop"
					url := "https://stardewvalleywiki.com/Carpenter%27s_Shop"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Robin",
						Name:    "Robin",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/1/1b/Robin.png",
					}
					embed = createEmbedShop(title, "robin", url, "Most construction projects are finished "+
						"in 2 days.", author)

				case "marnie":
					title := "Marnie's Ranch"
					url := "https://stardewvalleywiki.com/Marnie%27s_Ranch"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Marnie",
						Name:    "Marnie",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/5/52/Marnie.png",
					}
					embed = createEmbedShop(title, "marnie", url, "Also closed on Fall 18 and Winter 18.",
						author)

				case "clint":
					title := "Blacksmith"
					url := "https://stardewvalleywiki.com/Blacksmith"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Clint",
						Name:    "Clint",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/3/31/Clint.png",
					}
					embed = createEmbedShop(title, "clint", url, "Clint is available on Friday's before the "+
						"communitty center is completed.", author)

				case "wizard":
					title := "Wizard's Tower"
					url := "https://stardewvalleywiki.com/Wizard%27s_Tower"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Wizard",
						Name:    "Wizard",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/c/c7/Wizard.png",
					}
					embed = createEmbedShop(title, "wizard", url, "The wizard's notebook is only available "+
						"after completing the community center.", author)

				case "saloon":
					title := "The Stardrop Saloon"
					url := "https://stardewvalleywiki.com/The_Stardrop_Saloon"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Gus",
						Name:    "Gus",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/5/52/Gus.png",
					}
					embed = createEmbedShop(title, "saloon", url, "The Saloon will sell an unlimited number "+
						"of crab cakes for approximately 4 days after Willy's 6 heart cutscene.", author)

				case "fish":
					title := "Fish Shop"
					url := "https://stardewvalleywiki.com/Fish_Shop"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Willy",
						Name:    "Willy",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/8/82/Willy.png",
					}
					embed = createEmbedShop(title, "fish", url, "More goods become available as the "+
						"player increases their fishing skill.", author)

				case "pierre-general":
					title := "Pierre's General Store"
					url := "https://stardewvalleywiki.com/Pierre%27s_General_Store"
					embed = createEmbedShop(title, "pierre-general", url, "Pierre is closed on Wednesday until the community "+
						"center is completed", pierreAuthor())

				case "pierre-spring":
					title := "Pierre's General Store"
					url := "https://stardewvalleywiki.com/Pierre%27s_General_Store"
					embed = createEmbedShop(title, "pierre-spring", url, "Pierre is closed on Wednesday until the community "+
						"center is completed", pierreAuthor())

				case "pierre-summer":
					title := "Pierre's General Store"
					url := "https://stardewvalleywiki.com/Pierre%27s_General_Store"
					embed = createEmbedShop(title, "pierre-summer", url, "Pierre is closed on Wednesday until the community "+
						"center is completed", pierreAuthor())

				case "pierre-fall":
					title := "Pierre's General Store"
					url := "https://stardewvalleywiki.com/Pierre%27s_General_Store"
					embed = createEmbedShop(title, "pierre-fall", url, "Pierre is closed on Wednesday until the community "+
						"center is completed", pierreAuthor())

				case "joja-general":
					title := "JojaMart"
					url := "https://stardewvalleywiki.com/JojaMart"
					embed = createEmbedShop(title, "joja-general", url, "After the player purchases a Joja"+
						" membership, prices are the same as at Pierre's.", morrisAuthor())

				case "joja-spring":
					title := "JojaMart"
					url := "https://stardewvalleywiki.com/JojaMart"
					embed = createEmbedShop(title, "joja-spring", url, "After the player purchases a Joja"+
						" membership, prices are the same as at Pierre's.", morrisAuthor())

				case "joja-summer":
					title := "JojaMart"
					url := "https://stardewvalleywiki.com/JojaMart"
					embed = createEmbedShop(title, "joja-summer", url, "After the player purchases a Joja"+
						" membership, prices are the same as at Pierre's.", morrisAuthor())

				case "joja-fall":
					title := "JojaMart"
					url := "https://stardewvalleywiki.com/JojaMart"
					embed = createEmbedShop(title, "joja-fall", url, "After the player purchases a Joja"+
						" membership, prices are the same as at Pierre's.", morrisAuthor())

				case "oasis":
					title := "Oasis"
					url := "https://stardewvalleywiki.com/Oasis"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Sandy",
						Name:    "Sandy",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/4/4e/Sandy.png",
					}
					embed = createEmbedShop(title, "oasis", url, "The Oasis sells at least two items that"+
						" change daily.", author)

				case "adventurers-guild":
					title := "Adventurer's Guild"
					url := "https://stardewvalleywiki.com/Adventurer%27s_Guild"
					author := &discordgo.MessageEmbedAuthor{
						URL:     "https://stardewvalleywiki.com/Marlon",
						Name:    "Marlon",
						IconURL: "https://stardewvalleywiki.com/mediawiki/images/3/37/Marlon.png",
					}
					embed = createEmbedShop(title, "adventurers-guild", url, "Neither Marlon nor Gil accept gifts, gain"+
						" friendship points, or have bedrooms for the player to enter.", author)

				default:
					//I added this in the case something might go wrong
					embed.Description = "Oops, something gone wrong.\n" +
						"There's no villager with that name."
				}

			case "gifts":
				switch i.Data.Options[0].Options[0].Name {

				case "abigail":
					url := "https://stardewvalleywiki.com/Abigail"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/8/88/Abigail.png"
					embed = createEmbedGift("abigail", url, thumbnail)

				case "alex":
					url := "https://stardewvalleywiki.com/Alex"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/0/04/Alex.png"
					embed = createEmbedGift("alex", url, thumbnail)

				case "caroline":
					url := "https://stardewvalleywiki.com/Caroline"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/8/87/Caroline.png"
					embed = createEmbedGift("caroline", url, thumbnail)

				case "clint":
					url := "https://stardewvalleywiki.com/Clint"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/3/31/Clint.png"
					embed = createEmbedGift("clint", url, thumbnail)

				case "demetrius":
					url := "https://stardewvalleywiki.com/Demetrius"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/f/f9/Demetrius.png"
					embed = createEmbedGift("demetrius", url, thumbnail)

				case "dwarf":
					url := "https://stardewvalleywiki.com/Dwarf"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/e/ed/Dwarf.png"
					embed = createEmbedGift("dwarf", url, thumbnail)

				case "elliott":
					url := "https://stardewvalleywiki.com/Elliott"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/b/bd/Elliott.png"
					embed = createEmbedGift("elliott", url, thumbnail)

				case "emily":
					url := "https://stardewvalleywiki.com/Emily"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/2/28/Emily.png"
					embed = createEmbedGift("emily", url, thumbnail)

				case "evelyn":
					url := "https://stardewvalleywiki.com/Evelyn"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/8/8e/Evelyn.png"
					embed = createEmbedGift("evelyn", url, thumbnail)

				case "george":
					url := "https://stardewvalleywiki.com/George"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/7/78/George.png"
					embed = createEmbedGift("george", url, thumbnail)

				case "gus":
					url := "https://stardewvalleywiki.com/Gus"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/5/52/Gus.png"
					embed = createEmbedGift("gus", url, thumbnail)

				case "haley":
					url := "https://stardewvalleywiki.com/Haley"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/1/1b/Haley.png"
					embed = createEmbedGift("haley", url, thumbnail)

				case "harvey":
					url := "https://stardewvalleywiki.com/Harvey"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/9/95/Harvey.png"
					embed = createEmbedGift("harvey", url, thumbnail)

				case "jas":
					url := "https://stardewvalleywiki.com/Jas"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/5/55/Jas.png"
					embed = createEmbedGift("jas", url, thumbnail)

				case "jodi":
					url := "https://stardewvalleywiki.com/Jodi"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/4/41/Jodi.png"
					embed = createEmbedGift("jodi", url, thumbnail)

				case "robin":
					url := "https://stardewvalleywiki.com/Robin"
					thumbnail := "https://stardewvalleywiki.com/mediawiki/images/1/1b/Robin.png"
					embed = createEmbedGift("robin", url, thumbnail)
				}

			default:
				panic("no valid option chosen")
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
			if err != nil {
				log.Println(err)
			}
		},
	}
)

func morrisAuthor() *discordgo.MessageEmbedAuthor {
	return &discordgo.MessageEmbedAuthor{
		URL:     "https://stardewvalleywiki.com/Morris",
		Name:    "Morris",
		IconURL: "https://stardewvalleywiki.com/mediawiki/images/9/90/Morris.png",
	}
}

func pierreAuthor() *discordgo.MessageEmbedAuthor {
	return &discordgo.MessageEmbedAuthor{
		URL:     "https://stardewvalleywiki.com/Pierre",
		Name:    "Pierre",
		IconURL: "https://stardewvalleywiki.com/mediawiki/images/7/7e/Pierre.png",
	}
}
