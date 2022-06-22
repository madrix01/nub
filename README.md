# Nub

> Command line markdown editor which uses Github gists as a database.

## Installation

- Clone repository
```shell
git clone https://github.com/madrix01/nub.git
cd nub
```
- Download the dependencies
```shell
go mod download
```

- build the application
```shell
make build
```
- Copy output to /bin 
```
cp ./bin/nub /bin/nub
```

## Setting the config

As you run `nub` it creates `config.json` file in `~/.config/nub`
```json
{
    "Editor": "vim",
    "GistId": "<you gist id>",
    "Token": "<your github token>",
    "TempFolder": "/tmp",
    "Username": "madrix01"
}
```

## Usage

- `edit` - Edit note in a gist
```shell
nub edit
```
- `create` - Create note in a gist
```shell
nub create <filename>.md
```