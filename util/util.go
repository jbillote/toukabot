package util

import (
    "math/rand"
    "net/http"
    "time"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

/*
 * Function: RandomRange
 * Generate a random number greater than or equal to min and less than max
 *
 * Params:
 * min: Minimum number in the range
 * max: Maximum number in the range + 1
 *
 * Return:
 * Random number greater than or equal to min and less than max
 */
func RandomRange(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(max - min) + min
}

/*
 * Function: SafeSendMessage
 * Send a message to a Discord channel, including error checking
 *
 * Params:
 * s: Pointer to the Discord session to use
 * c: ID of the channel to send the message to
 * m: Message to send
 */
func SafeSendMessage(s *discordgo.Session, c string, m string) *discordgo.Message {
    v, err := s.ChannelMessageSend(c, m)

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Error("Unable to send message")
    }

    return v
}

/*
 * Function: TestURLValidity
 * Test whether or not a URL is accessible
 *
 * Params:
 * s: URL to test
 *
 * Return:
 * True if URL is accessible, else false
 */
func TestURLValidity(s string) bool {
    r, err := http.Get(s)
    return err == nil && r.StatusCode == 200
}