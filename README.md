# bank
A simple banking API implemented using event sourcing.

See the Makefile for commands.

To build the server run `make build`. This will produce a `main` executable.

To run the built server: `./main -server`
To run tests with the build server: `./main -test`

`-verbose` can also be passed.

In development use:

`make run` for the server
`make test` for the tests

If you modify the sql, run `make sqlc`.
If you add migrations, run `make migrate`.

To install deps locally, `make install`.
To open a database shell, `make dbshell`.

This project assumes you are using a postgres database.
It can also be developed/run in a GitHub Codespace.