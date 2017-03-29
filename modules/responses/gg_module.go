package responses

/*
 * Function: GG
 * Respond with "ez" when someone says "gg"
 *
 * Params:
 * m: Message to parse
 * a: Message's author's ID
 *
 * Return:
 * True if the message is valid, else false
 * String containing response if message is valid
 */
func GG(m string, a string) (bool, string) {
    if (m == "gg") {
        return true, "ez"
    }

    return false, ""
}
