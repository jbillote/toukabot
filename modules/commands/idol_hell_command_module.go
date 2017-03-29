package commands

import (
    "strings"

    "github.com/jbillote/toukabot/util"
)

/*
 * Function: IdolHellCommand
 * Select a random image from the lovelive and idolmaster commands
 *
 * Params:
 * p: Split message to parse
 *
 * Return:
 * True if the message contained the idolhell command, else false
 * String containing response if command is present
 */
func IdolHellCommand(p []string) (bool, string) {
    if strings.Contains(strings.ToLower(p[0]), "idolhell") {
        pc := util.RandomRange(1, 101)

        if pc < 50 {
            return true, util.GetResponse("lovelive")
        } else {
            return true, util.GetResponse("idolmaster")
        }
    } else {
        return false, ""
    }
}