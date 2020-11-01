# emaild

![emaild-tests](https://github.com/4thel00z/emaild/workflows/Test/badge.svg)

## Description

emaild is a cool email daemon 😎 which can:

- provide multi vendor support (First vendor is GMail and tests are running)
- schedule email sending to a later time in the day and handle bulk requests. (Not done yet)

Furthermore you can find the `gmail-token` binary also in this repository.
It makes it easy to generate a token for the gmail module of emaild to work:

[![emaild in action](https://asciinema.org/a/WEMHnuuj7yXHf7pHEsNOt8gEB.svg)](https://asciinema.org/a/WEMHnuuj7yXHf7pHEsNOt8gEB)

## Installation

```
make
# in case you want emaild to be installed under /usr/local/bin
make install
```

## Configuration

The configuration of the daemon is still TBD. Have some patients n00b 😇

### Generate a Gmail Token

Go to the Go Gmail [getting started](https://developers.google.com/gmail/api/quickstart/go).
Click on "Enable Gmail API" and select Desktop Client.
Download the `credentials.json` and store somewhere securely.
Use the `gmail-token` binary:

With make:
```
make build-gmail-token
env ARGS="-config <path-to-the-credentials-json> -out token.json" make run-gmail-token 
```

If you have [just](https://github.com/casey/just) installed:
```
just build-gmail-token
just run-gmail-token -config <path-to-the-credentials-json> -out token.json
```

It will prompt you to go to websites. Do that. Copy the token into the console and press enter.
It will save the token to a file called `token.json` if you did not specify another path via the `-out` parameter.


## Todos

- Provide unix domain socket webserver
- Document the interface
- Provide a client


## Architecture

TBD

## Contribution guide

TBD
