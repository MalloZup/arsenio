# arsenio

[![Go Report Card](https://goreportcard.com/badge/github.com/MalloZup/arsenio)](https://goreportcard.com/report/github.com/MalloZup/arsenio)  [![GoDoc](https://godoc.org/github.com/MalloZup/arsenio?status.svg)](https://godoc.org/github.com/MalloZup/arsenio)
___
Arsenio: always in search of the Beauty and minimalistic design. Golang bot for GitHub  opensource projects.
___

# Description:

History: when you work on middle-big project people create issues and PRs but then after a while people do something else or they simply forget about the issue or PR for some reason.

Arsenio will be a gentle golang bot, providing utilty to maintainers to opensource projects, for respecting roadmaps, milestones and timing.


# Roadmap:

Arsenio is under active dev. see https://github.com/MalloZup/arsenio/milestones

# Howto:

## Installation

```golang 
go get -u github.com/MalloZup/arsenio
```
## Quickstart:

### Authentication

GitHub GraphQL API v4 [requires authentication](https://developer.github.com/v4/guides/forming-calls/#authenticating-with-graphql).

0) create a token via GitHub.
1) Put the token in `.arsenio.yml` configuration fil

```yml
GITHUB_TOKEN: mytokenGithub
```

### Run some commands
2) test it with running commands ( this might change in future)

```golang
./arsenio ping
```
 

## Goals: 

For the first version, the goal is to ping the authors of PRs/ISSUe with a message on the issue/PR they forgot since XX months. (default could be 2 months).

Further versions can include remind on issues that were due to a previous milestone. ( this can be usefull for Release Manager or Scrum-master, or any project working with GitHub milestone/scrums).


## Version 0.1

Implement:

`arsenio ping -c repos.yml` 

Description:
- [ ] given a list of github Repos, find out issues/pr older then X Months, when true ping authors with a  template comment on the PR/ISSUEs


# Development


## Technical choices, architecture:

Arsenio use the GitHub GraphQL API v4 (https://developer.github.com/v4/),  https://github.com/shurcooL/githubv4.
For the command line https://github.com/spf13/cobra and viper.


Project layout: https://github.com/golang-standards/project-layout

This project use go modules, you will need latest golang version.

# Contribute:

For any suggestion idea, feel free to open an issue or fix an issue or spend a PR ( at moment  project is under dev).


