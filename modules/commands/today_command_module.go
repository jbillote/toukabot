package commands

import (
    "strings"
    "time"

    "github.com/jbillote/toukabot/util"
)

/*
 * Function: TodayCommand
 * Select a random image dependent on the current day of the week
 *
 * Params:
 * p: Split message to parse
 *
 * Return:
 * True if the message contained the today command, else false
 * String containing the response, if the command was present
 */
func TodayCommand(p[] string) (bool, string) {
    if strings.Contains(strings.ToLower(p[0]), "today") {
        t := time.Now().Weekday()

        switch t {
        case time.Monday:
            return true, util.GetResponse("monday")
        case time.Tuesday:
            return true, util.GetResponse("tuesday")
        case time.Friday:
            return true, util.GetResponse("friday")
        case time.Saturday:
            return true, util.GetResponse("saturday")
        case time.Sunday:
            return true, util.GetResponse("sunday")
        default:
            return true, util.GetResponse("idklol")
        }
    } else {
        return false, ""
    }
}
