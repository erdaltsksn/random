# Contributing

This document contains a set of guidelines to help you during the contribution
process.

We are happy to welcome all contributions from anyone willing to improve this
project. Thank you for helping out and remember, no contribution is too small.

## Step by step contribution guide

01. **Install** the [prerequisites](#prerequisites)
02. **Fork** the repository and **Clone** the fork to your own machine
03. **Open** an issue before working on your changes
04. **Create** a git branch and **Start working** on your fix, feature and etc
    on **its own branch**
05. **Create** and **Run** the **tests** (if there are any)
    - **Run** `go run docs/gen.go` to automatically generate command docs
06. **Commit** your changes according to [**Commit Message Guidelines**](#commit-message-guidelines)
07. **Push** the changes to your fork and **Create** a Pull Request for review
08. **Communicate** with the maintainer about the revisions
09. **Be responsive** if someone request changes for your contributions
10. Your contribution gets accepted. 🎉🎉🎉

**NOTE:** Be sure to get the latest from "**upstream**" before making a pull
request!

### Prerequisites

- [Git](https://git-scm.com)
- [Golang](https://golang.org)

## Commit Message Guidelines

We have strict rules over how our **git commit messages** can be formatted. This
leads to **more readable messages** that are easy to follow when looking through
the **project history**.

### Commit Message Format

```sh
<type>: (<scope>) <subject>
<BLANK LINE>
<type>: <body>
<BLANK LINE>
<type>: <footer>
```

### Template

We use the commit message template from [dotfiles]. You may find any other
additional *types*, *scopes*, and more information below.

[dotfiles]: https://github.com/erdaltsksn/dotfiles/blob/master/git/.gittemplate

### Types

As well as those specified in the [dotfiles] file, The following is the list of
supported types:

### Scopes

The following is the list of supported scopes:

- **(*)** Multiple scopes or Unknown.
- **(devops)** e.g. git, travis, pre-commit, netlify, heroku, npm, gulp and etc.
- **(readme)** e.g. readme, contributing, changelog, license and etc.
- **(cmd)** Root command and shared functions.
- **(card)**
- **(digit)**
- **(letter)**
- **(lorem)**
- **(name)**
- **(password)**
- **(state)**
- **(user-agent)**

#### Old / Deprecated Scopes

Some of the scopes have become old/deprecated. The following is the list and
their newer equivalents.

| OLD                                  | NEW                                   |
|--------------------------------------|---------------------------------------|
| (pkg)                                | `REMOVED`                             |

### Revert and Merge

Use default GIT templates for `revert` and `merge`.

## Issues

Before creating a new issue, please search for similar issues, open or closed,
to see if someone else has already noticed the same problem and possible
solutions.

Do not comment on open issues unless you can provide more information to resolve
it. Use the subscribe function to keep up-to-date with the issue or the voting
system to support it.

## Bug reports

When you can't find a previous bug, open an issue keeping in mind the following
considerations:

- Try to reproduce the bug using the code found on the master branch
- Copy and paste the full error message, including the `backtrace`
- Be as detailed as possible and include any additional information

## Feature requests

If you want to request or implement a new feature please submit an issue
describing the details and possible use cases.

## Thanks

Thank you to all who have contributed in this great project.
