package modules

import (
    "github.com/jbillote/toukabot/modules/responses"

    "github.com/bwmarrin/discordgo"
)

type ResponseFunc func(m string, a string) (bool, string)

var rfs = []ResponseFunc{
    responses.AyyLmao,
    responses.GG,
    responses.Wave,
}

/*
 * Function: Response
 * Parse the input to see if there are any pre-defined responses available
 *
 * Params:
 * s: Pointer to the Discord session to use
 * m: Pointer to the MessageCreate event that activated
 *
 * Return:
 * True if there was a pre-defined response, else false
 * String containing the pre-defined response, if there was one
 */
func Response(s *discordgo.Session, m *discordgo.MessageCreate) (bool, string) {
    d := false

    for i := 0; i < len(rfs) && !d; i++ {
        if d, r := rfs[i](m.Content, m.Author.ID); d {
            return true, r
        }
    }

    return false, ""
}