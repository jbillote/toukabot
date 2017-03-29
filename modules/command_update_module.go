package modules

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "github.com/Sirupsen/logrus"
    "time"
)

/*
 * Function: GetCommands
 * Get a list of commands "publicly" available to users
 *
 * Return:
 * Map of command name to help string
 */
func GetCommands() map[string]string {
    c := make(map[string]string)

    s := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-1")}))
    p := &dynamodb.ScanInput{
        TableName: aws.String("ToukaBot"),
        ProjectionExpression: aws.String("command, helpString"),
    }

    req, resp := s.ScanRequest(p)
    err := req.Send()

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Error("Unable to get list of commands from DynamoDB")
    }

    for _, e := range resp.Items {
        if _, ok := e["helpString"]; ok {
            c[*e["command"].S] = *e["helpString"].S
        }
    }

    logrus.Info("Updated command list")

    return c
}

/*
 * Function: UpdateCommands
 * Goroutine to update list of commands "publicly" available to users at midnight
 *
 * Params:
 * c: Pointer to map of commands to update. Map is of command name to help string
 */
func UpdateCommands(c *map[string]string) {
    f := true
    n := time.Now()
    t := n.AddDate(0, 0, 1)
    m := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

    for true {
        if f {
            time.Sleep(m.Sub(n) * time.Minute)
            f = false
        } else {
            time.Sleep(time.Duration(24) * time.Hour)
        }

        *c = GetCommands()
    }
}