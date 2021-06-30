# GitHub Simple Downloader

It is often required to fetch few modules of the codebase from a bulky mono-repo. Unfortunately github doesn't permit for such operation directly. It provides an API endpoints to achieve that objective though.

*The repo contains the codebase of the CLI tool* to download part of the codebase, either a file/blob, or directory/tree from a github repository.

## How to use

### Note

[cURL](https://curl.se/) is required to be install in your OS.

### Steps

1. Download the binary and copy it to `/usr/bin/` or `/usr/local/bin`:

```bash
todo
```
2.

- Run to see the man/help:

```bash
github-downloader -h
```

- Run to fetch the current version of the app:

```bash
github-downloader -b https://github.com/kislerdm/github-downloader/blob/master/cli/VERSION
```

You shall see the following in stdout:

```bash
2021/06/30 17:42:11 Download [4 bytes] cli/VERSION
```

Now if you were to run the command:

```bash
cat /tmp/cli/VERSION
```

You shall expect to see the current version of the app.

## How to use github token

It is advised to authenticate with github to overcome the API calls [rate limit](https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting). Please follow [the instructions](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to setup an access token.

The required authorization scope is `repo`.

There are two options to allow the CLI to access the token:

- Export the token as an environmental variable `GITHUB_TOKEN`

- Provide the token as a cmd option:

```bash
github-downloader -token=xxxxxx https://github.com/kislerdm/github-downloader/blob/master/cli/VERSION
```

### Note

The *cmd option token* is prioritized above the environmental variable. In the other words, if you export token as an envvar and provide another token as the cmd option `token`, the later value will be used to authenticate with the github API.
