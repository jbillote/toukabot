package modules

import (
    "github.com/jbillote/toukabot/modules/commands"
    "github.com/jbillote/toukabot/util"

    "strings"

    "github.com/bwmarrin/discordgo"
)

type CommandFunc func(p []string) (bool, string)

var cfs = []CommandFunc{
    commands.IdolHellCommand,
    commands.TodayCommand,
}

/*
 * Function: Command
 * Parse the input for any possible commands
 *
 * Params:
 * s: Pointer to the Discord session to use
 * m: Pointer to the MessageCreate event that activated
 * c: Map of command names to help string
 */
func Command(s *discordgo.Session, m *discordgo.MessageCreate, c map[string]string) (bool, string) {
    if len(m.Content) > 0 && m.Content[0] == '+' {
        p := strings.Split(m.Content, " ")

        if strings.Contains(strings.ToLower(p[0]), "help") {
            var sc *string

            if len(p) > 1 {
                *sc = strings.ToLower(p[1])
            } else {
                sc = nil
            }

            HelpCommand(s, m, sc, c)
        } else {
            var v bool

            // Check function-driven commands first
            for _, f := range cfs {
                if v, r := f(p); v {
                    return true, r
                }
            }

            // Database-driven commands
            if !v {
                for k := range(c) {
                    if strings.Contains(strings.ToLower(p[0]), k) {
                        return true, util.GetResponse(k)
                    }
                }
            }
        }
    }

    return false, ""
}