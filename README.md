# arsenio

Gentle stale bot for GitHub Issues and PullRequests.

# Description

History: when you work on middle-big project people create issues and PRs but then after a while this are in a stale forgot status.

Arsenio will be a gentle golang bot, providing utilty to maintainers to opensource projects, for respecting roadmaps, milestones and timing.

For the first version, the goal is to ping authors with a message on the issue/PR they forgot since XX months. (default could be 2 months).

Further versions can include remind on issues that were due to a previous milestone. ( this can be usefull for Release Manager or Scrum-master, or any project working with GitHub milestone/scrums).

# Contribute:

For any suggestion idea, feel free to open an issue or fix an issue or spend a PR ( at moment  project is under dev).

# Roadmap:

## Version 0.1

Implement:

`arsenio ping -c repos.yml` 

Description:
- [ ] given a list of github Repos, find out issues/pr older then X Months, when true ping authors with a  template comment on the PR/ISSUEs


# Development

This project use go modules, you will need latest golang version.
