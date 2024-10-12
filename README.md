# CBS-02 | Team 09

Team members:

* Vadim Iarullin - v.iarullin@innopolis.university
* Nikita Sannikov - n.sannikov@innopolis.university
* Aleksandr Efremov - a.efremov@innopolis.university
* Artur Lukianov - a.lukianov@innopolis.university
* Andrew Boronin - a.boronin@innopolis.university


# Running instructions

```bash
git clone https://github.com/Poxidq/sa-assignment-5
cd sa-assignment-5
docker compose up -d --build

# CLI commands
docker-compose run cli-service -h
Usage:
  cli-service [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  feed        Get the latest 10 messages
  help        Help about any command
  like        Like a message by ID
  post        Post a new message
  register    Register a new user

Flags:
  -h, --help   help for cli-service

Use "cli-service [command] --help" for more information about a command.


# OR 
docker exec -it <cli-service-container-id> /bin/sh
```