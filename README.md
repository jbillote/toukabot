# ToukaBot

ToukaBot is a basic chat bot written in Go for the voice-over-IP program,
Discord. Its main function is to respond to various commands and certain
messages. These responses are obtained from a DynamoDB database or, for more
complex commands, functions.

Commands at the moment are expected to be preceded by ``+``. Command parsing
is also non-exact, meaning that ``+lol`` and ``+3i4jt43ajololr3q3r2q`` would be
considered equivalent since both contain ``lol``.

## Compiling

Follow [these instructions](https://golang.org/doc/install) to install Go for
your operating system.

The following dependencies must be installed in order to compile and run
ToukaBot:

1. ``go get github.com/bwmarrin/discordgo``
2. ``go get -u github.com/aws/aws-sdk-go``
3. ``go get github.com/sirupsen/logrus``

You can pull the newest release of ToukaBot by running
``go get github.com/jbillote/toukabot``.

ToukaBot can then be compiled by running ``make`` or ``go build``.

If you get any errors during compilation, make sure that all dependencies are
installed.

## Configuration

ToukaBot uses a ``.json`` file for configuration. By default, it uses the 
``config.json`` file found in the same directory as the executable. However,
you can specify which configuration file to use using the ``-config`` flag when
running ToukaBot.

The fields inside the configuration ``.json`` file are as follows:

* **ownerId** is the user ID of the "owner" of the bot. The owner is the user
who should be in charge of maintaining and configuring ToukaBot for the
server. At the current moment, the only thing special the owner can do is use
the ``+changename`` command to change the bot's name.
* **botToken** is the bot's user token provided from the [Discord developers
console](https://discordapp.com/developers/docs/intro). Make sure to include
the ``Bot `` prefix.
* **statuses** is an array of strings that are to be used as the bot's status
messages. These statuses will be randomly rotated at set intervals.
* **statusTime** is how often statuses should be rotated, in seconds.

## Database Configuration

ToukaBot uses Amazon DynamoDB to store image URLs and responses. It is up to
you to have your own AWS account and DynamoDB table as well as any user
authentication information on your computer.

The DynamoDB table is expected to be named ``ToukaBot``.

The DynamoDB table is expected to have the following attributes:

* **command** is the primary key and is what the bot expects when parsing commands.
* **helpString** is the text that is displayed when ``+help <command name>`` is
used. If this attribute is not present, the command will not be shown by
``+help`` nor will users be able to directly use it.
* **images** is a list of strings containing all images associated with a
command. ToukaBot will retreive a random image from this list, if this
attribute is present.
* **response** is a predefined, single response associated with a command. It
may or may not be an image.

*Note*: It is expected that both ``images`` and ``response`` are not present at
the same time. If both attributes are present, ``response`` is used. If neither
are present and there is no function that can handle the command, the bot will
crash.

## Modules

ToukaBot was designed to be modular. That is, new features, such as games or 
more complex commands, can be added by making new modules and registering it
with as little effort as possible.

At the current moment, only adding new command and response modules is
streamlined. This can be done by creating a new module file with the
appropriate function signature and registering it in ``command_module.go`` or
``response_module.go``. For examples of this, see the files under the
``modules/commands`` and ``modules/responses`` folders.

*Note*: The provided ``idol_hell_command.go`` module will only work if your
ToukaBot table has the commands ``idolmaster`` and ``lovelive`` in it.
Similarly, the ``today_command.go`` module will only work if the commands
listed inside it are present in the table.

## Other Features

### URL Checking

ToukaBot features URL checking to make sure that dead URLs are not posted.
Currently, URLs are checked when fetched from the database, and they are only
checked once to see if they respond with an HTTP status code of 200. If a link
is dead at the time of access, it is logged so those running the bot can go 
back and double check to make sure it is not just a one-time access problem.

### Rotating Statuses

ToukaBot changes its status at predefined intervals. This interval and the list
of statuses to choose from is defined in the configuration file. See above for
more details.

### Automated Command List Updates

ToukaBot stores the list of available commands on launch, and updates this list
at midnight. This allows for the quick addition of simple commands without
having to recompile. Currently, there is no way to force a command list update
without restarting the bot.

## Known Issues

* If a command is present in the database and there is no ``images`` or
``response`` attribute present or a function that can handle the command, the
bot will crash. This will be addressed in a later update.