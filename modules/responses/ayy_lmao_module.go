package responses

/*
 * Function: AyyLmao
 * Determine if the message is "ayy" with any amount of trailing y's
 *
 * Params:
 * m: Message to parse
 * a: Message's author's ID
 *
 * Return:
 * True if the message is valid, else false
 * String containing response if message is valid
 */
func AyyLmao(m string, a string) (bool, string) {
    var v bool
    i := 2

    if len(m) > 2 && m[0:3] == "ayy" {
        v = true

        for v && i < len(m) {
            v = (m[i] == 'y')
            i++
        }
    }

    // Don't bother with dummy value since value is used only if v is true
    return v, "lmao"
}