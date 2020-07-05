# Mordred

Mordred is a Git repository scanning tool.

Git is a widely adopted VCS (Version Control Software). You can learn more about it here: https://git-scm.com.

One thing amazing with it, is that you cannot delete previously pushed contents (to be fair you can but it is really difficult).
It is pretty handy since it means you can't lost a file contents by accidentally deleting it in your working directory. You can always set the code state to a previous version and work from there.

However, from time to time, people accidentally push sensitive data to a Git repository. This can be an hardcoded API token in a script, secrets written in an environment variable file that wasn't properly git-ignored, or even a binary that contains an uncrypted secret string variable that can be recovered by decompiling it.

Even if developers delete the file afterwards and push a fix commit, Git won't remove the previously written data from its internal objects store. Moreover, deleting the commit itself is not enough.

Most of the time, people rotate the secrets to avoid getting hacked. Sometimes they forget about it, or they simply don't realize such a file was pushed.

This is what Mordred is for.

Mordred helps you to find sensitive informations that could be stored in your repository.
It is a CLI tool that extracts information from different sources:
- commits that contain personal informations (names, e-mails, GPG signing keys) about the repository authors and can be useful for an OSINT research
- blobs that can contain secrets committed by mistake.

With such a tool you can get a better idea of the attack surface exposed by a repository.

Mordred currently indexes two types of data:
- identities: names and e-mails of authors and committers of a repository
- strings: any string catched in Golang source code files, sorted by type (IP, DNS record, base64, etc.)

## Build

```shell
make build
```

## Install

```shell
make install
```

## Usage

```shell
mordred --help
```
should get you started!

## Roadmap
Mordred is a WIP.

Planned features:
- add tests for the parsing and indexing logic
- support blob parsing of other files types
- add a default parser for blobs that don't match any supported pattern
- support new strings types
