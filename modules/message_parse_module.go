package modules

import (
    "github.com/jbillote/toukabot/util"

    "github.com/bwmarrin/discordgo"
)

func MessageParse(s *discordgo.Session, m *discordgo.MessageCreate, uid string, c map[string]string) {
    // Make sure the bot doesn't respond to itself
    if m.Author.ID == uid {
        return
    }

    var d bool

    // Check response functions
    if d, r := Response(s, m); d {
        util.SafeSendMessage(s, m.ChannelID, r)
    }

    // Check commands if a response wasn't generated
    if !d {
        if d, r := Command(s, m, c); d {
            util.SafeSendMessage(s, m.ChannelID, r)
        }
    }

    return
}
