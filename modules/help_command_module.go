package modules

import (
    "github.com/jbillote/toukabot/util"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

/*
 * Function: HelpCommand
 * Process a help command and message the requester the appropriate information
 *
 * Params:
 * s: Pointer to Discord session to use
 * m: Pointer to MessageCreate event to process
 * sc: Pointer to string containing specific command to look up, or nil if none given
 * c: Map of commands
 */
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate, sc *string, c map[string]string) {
    dm, err := s.UserChannelCreate(m.Author.ID)

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Error("Unable to create direct message channel")
    }

    if sc != nil {
        if h, ok := c[*sc]; ok {
            util.SafeSendMessage(s, dm.ID, "``" + *sc + " - " + h + "``")
        } else {
            util.SafeSendMessage(s, dm.ID, "``" + *sc + "`` is not a valid command. For a list of valid commands, use ``+help``.")
        }
    } else {
        h := "Available commands:\n```"

        for k, _ := range c {
            h += k + "\n"
        }

        h += "```For more detailed information, use +help <Command name>"

        util.SafeSendMessage(s, dm.ID, h)
    }
}
