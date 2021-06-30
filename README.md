# GitHub Simple Downloader

It's often required to fetch part of the codebase from a big mono-repo. Unfortunately github doesn't permit for such operation directly. It provides an API endpoints to achieve that objective though. This repo contains the codebase of the CLI tool to download part of the codebase, either a file/blob, or directory/tree from the git repo.

## Requirements

- GITHUB_TOKEN

## How to use

1. Download the binary and copy it to `/usr/bin/` or `/usr/local/bin`
2. Run for the info:

```bash
github-download -h
```
