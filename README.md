# pokedexcli

A lightweight and fast Pokédex CLI built entirely in Go. Fill up your Pokédex by exploring cities and catching Pokémon. Inspect your caught Pokémon for additional details.

* External API calls are cached to drastically decrease response times

## Installation
### Install via Go
If you have Go installed:
```sh
go install github.com/jzaager/pokedexcli@latest
```
Otherwise, to build from source:
```sh
git clone https://github.com/jzaager/pokedexcli.git  
cd pokedexcli  
go build -o pokedexcli  
```

## Usage
Run the executable to start the Pokédex REPL.
### Available Commands
| Command                 | Description                                    |
|-------------------------|------------------------------------------------|
| help                    | Displays help message with list of commands    |
| exit                    | Exit the Pokedex                               |
| map                     | Displays next page of locations                |
| mapb                    | Displays previous page of locations            |
| explore <location_name> | Displays a list of pokemon at a given location |
| catch <pokemon_name>    | Attempt to catch a pokemon                     |
| inspect <pokemon_name>  | View details about a caught pokemon            |
| pokedex                 | Displays a list of all your caught pokemon     |

The *up* arrow can be used to scroll through your previoius commands.

The chance to catch a given Pokémon is based on its base experience.

## Planned Features
* Changes to location exploration to limit choices
* Storing Pokémon in a "party"
* Battles and experience gain for Pokémon in a party
* Random encounters with wild Pokémon
* Persistent Pokédex storage to disk
