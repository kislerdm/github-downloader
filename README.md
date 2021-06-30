# GitHub Simple Downloader

**Motivation**: It is often required to fetch a couple of codebase modules from bulky monorepo. Unfortunately standard [git client](https://git-scm.com/) doesn't permit for such operation directly. Github provides a set of API endpoints to achieve that objective though.

*The repo contains the codebase of the CLI tool* to download part of the codebase: either a file/blob, or directory/tree from a github repository.

## Installation

### Requirements

- [cURL](https://curl.se/)

- [tar](https://man7.org/linux/man-pages/man1/tar.1.html)

### Instructions

Download and execute [the installation script](./install.sh).

You can run the commands to install *the latest app version*:

```bash
curl -S https://github.com/kislerdm/github-downloader/install.sh
chmod +x install.sh
./install.sh
```

**Note** you can install different version of the app by running the command:

```bash
./install.sh <<VERSION>>
```

## How to use

**Note** The tool was not tested in windows.

- Run to see the help:

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

## Collaboration

If you find this tool useful but missing some features, please open a github issue, or submit a PR with modifications.

### Note

The codebase is distributed under [the MIT copyright license](./LICENSE).

## Author

[Dmitry Kisler](https://www.dkisler.com)

Feel free to get in touch [here](https://www.linkedin.com/in/dkisler/) to chat over tech stuff :)
