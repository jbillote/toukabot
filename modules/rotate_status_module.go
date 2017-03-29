package modules

import (
    "github.com/jbillote/toukabot/util"

    "time"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

/*
 * Function: RotateStatuses
 * Goroutine for rotating status messages at a regular interval
 *
 * Params:
 * s: Pointer to Discord session to update
 * l: Array of status messages
 * t: Interval at which to rotate statuses, in seconds
 */
func RotateStatuses(s *discordgo.Session, l []string, t int64) {
    for true {
        time.Sleep(time.Duration(t) * time.Second)

        n := l[util.RandomRange(0, len(l))]

        err := s.UpdateStatus(0, n)
        if err != nil {
            logrus.WithFields(logrus.Fields{
                "err": err,
            }).Error("Unable to change status")
        }

        logrus.Info("Changed status to " + n)
    }
}