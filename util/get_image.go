package util

import (
    "strconv"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "github.com/Sirupsen/logrus"
)

/*
 * Function: GetResponse
 * Get a random image or the response associated with the given command
 *
 * Params:
 * c: The command to use
 *
 * Return:
 * The response to send
 */
func GetResponse(c string) string {
    s := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-1")}))
    p := &dynamodb.QueryInput{
        TableName: aws.String("ToukaBot"),
        AttributesToGet: []*string{
            aws.String("images"),
            aws.String("response"),
        },
        KeyConditions: map[string]*dynamodb.Condition{
            "command": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(c),
                    },
                },
            },
        },
    }

    req, resp := s.QueryRequest(p)
    err := req.Send()

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Error("DynamoDB query failed")

        return "I forgot how to respond to that :P"
    } else {
        if r, ok := resp.Items[0]["response"]; ok {
            // Check if response is a URL
            if (*r.S)[0:7] == "http://" {
                if TestURLValidity(*r.S) {
                    return *r.S
                } else {
                    return "I couldn't find the image I wanted to send ;-;"
                }
            } else {
                return *r.S
            }
        } else {
            ci := make(map[int]bool)

            for true {
                var i int
                n := false

                for !n {
                    i = RandomRange(0, len(resp.Items[0]["images"].L))

                    if !ci[i] {
                        ci[i] = true
                        n = true
                    }
                }

                u := resp.Items[0]["images"].L[i]

                if TestURLValidity(*u.S) {
                    logrus.Info("Got " + c + "[" + strconv.Itoa(i) + "]")
                    return *(u.S)
                } else {
                    logrus.WithFields(logrus.Fields{
                        "url": *(u.S),
                    }).Warn(c + "[" + strconv.Itoa(i) + "] potentially broken")

                    if len(ci) == len(resp.Items[0]["images"].L) {
                        logrus.Warn("All links for " + c + " potentially broken")
                        return "I couldn't find an image to send ;-;"
                    }
                }
            }
        }
    }

    return ""
}