package main

import (
    "github.com/jbillote/toukabot/modules"

    "encoding/json"
    "flag"
    "os"
    "os/signal"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

type ToukaBotConfig struct {
    OwnerID    string   `json:"ownerId"`
    BotToken   string   `json:"botToken"`
    Statuses   []string `json:"statuses"`
    StatusTime int64    `json:"statusTime"`
}

var (
    // discordgo session
    d *discordgo.Session

    // Config
    cn ToukaBotConfig

    // Bot user ID
    uid string

    // Command list
    c map[string] string
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
    logrus.Info("ToukaBot started")
    s.UpdateStatus(0, "#JustBotThings")

    // Get initial command list
    c = modules.GetCommands()

    // Save bot user id
    uid = e.User.ID
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Parse messages in a new thread
    go modules.MessageParse(s, m, uid, c)
}

func main() {
    // Check to see if a path to a new config file was given
    cl := flag.String("config", "./config.json", "Configuration file")
    flag.Parse()

    // Try to open the config file
    cf, err := os.Open(*cl)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
            "path": *cl,
        }).Fatal("Unable to open config")
        os.Exit(-1)
    }

    // Parse the config file into a ToukaBotConfig struct
    j := json.NewDecoder(cf)
    err = j.Decode(&cn)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Unable to parse config file")
        os.Exit(-1)
    }

    // Create discord session
    logrus.Info("Starting Discord session...")
    d, err = discordgo.New(cn.BotToken)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Failed to create Discord session")
    }

    // Add ready and message create handlers
    d.AddHandler(onReady)
    d.AddHandler(onMessageCreate)

    err = d.Open()
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Failed to create Discord websocket connection")
        os.Exit(-1)
    }

    // Start thread to change the status
    go modules.RotateStatuses(d, cn.Statuses, cn.StatusTime)

    // Start thread to update command list
    go modules.UpdateCommands(&c)

    // Wait for a signal to quit
    s := make(chan os.Signal, 1)
    signal.Notify(s, os.Interrupt, os.Kill)
    <-s

    return
}
