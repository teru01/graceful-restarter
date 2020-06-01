# graceful-restarter

This project is inspired by [server-starter](https://github.com/lestrrat-go/server-starter), but a bit different in that graceful-restarter provides graceful shutdown and restart on TCP layer.

# Installation

```
$ curl -L -O https://github.com/teru01/graceful-restarter/releases/download/1.0/grestarter
$ chmod +x grestarter
$ mv grestarter /usr/local/bin
```

# How to use

```
$ grestarter [-L IP:Port] program arg1 arg2 ...
```

`program` has to use "github.com/teru01/graceful-restarter/graceful-listener" as listener (see example directory).

# How it works

1. The `grestarter` command  makes a socket listening on the given IP:Port (default 127.0.0.1:0).

2. Then, `grestarter` passes the opened socket file descriptor to `program` by launching as a child process.

3. Incoming requests are handled by the child process.

3. If `grestarter` process got a `SIGHUP` signal, it sends `SIGTERM` to all its child processes.
