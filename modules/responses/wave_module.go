package responses

/*
 * Function: Wave
 * Respond to o/ and \o by reversing what was said and pinging the user
 *
 * Params:
 * m: Message to parse
 * a: Message's author's ID
 *
 * Return:
 * True if the message was valid, else false
 * String containing response if message is valid
 */
func Wave(m string, a string) (bool, string) {
    if m == "o/" {
        return true, "<@" + a + "> \\o"
    } else if m == "\\o" {
        return true, "<@" + a + "> o/"
    }

    return false, ""
}
