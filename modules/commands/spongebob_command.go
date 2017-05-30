package commands

import "strings"

/*
 * Function: SpongebobCommand
 * Posts a Spongebob image and alternates message capitalization.
 *
 * Params:
 * p: Split message to parse
 *
 * Return:
 * True if nothing went wrong, else false
 * String containing response
 */
func SpongebobCommand(p []string) (bool, string) {
    capital := true
    message := ""

    if strings.Contains(strings.ToLower(p[0]), "spongebob") {
        for ndx, section := range p {
            if ndx > 0 {
                lower := strings.ToLower(section)

                for _, char := range lower {
                    if capital {
                        message += strings.ToUpper(string(char))
                    } else {
                        message += string(char)
                    }

                    capital = !capital
                }

                message += " "
            }
        }

        message += "\n\nhttps://cdn.discordapp.com/attachments/293841525148745748/319258428046311425/mocking-spongebob.png"
    } else {
        return false, ""
    }

    return true, message
}