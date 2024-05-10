# ahab

_tail the whale!_

ahab is a quick tool to allow following the logs of multiple docker containers. it was created to provide similar functionality as [stern](https://github.com/stern/stern) for kubernetes.

## usage

```bash
NAME:
   ahab - tail the whale!

USAGE:
   ahab [global options] command [command options] [container IDs or names]

VERSION:
   0.0.2

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --since value, -s value  show logs since timestamp [e.g. "2013-01-02T13:23:37Z"] or relative [e.g. "42m" for 42 minutes] (default: "10m")
   --until value, -u value  show logs until timestamp [e.g. "2013-01-02T13:23:37Z"] or relative [e.g. "42m" for 42 minutes]
   --no-follow              don't follow after printing the tails of the given containers (default: false) [$AHAB_NO_FOLLOW]
   --no-timestamps          don't print timestamps with log entries (default: false) [$AHAB_NO_TIMESTAMPS]
   --help, -h               show help
   --version, -v            print the version
```

## demo

![ahab_demo](https://cobi.dev/static/img/github/gif/ahab-0.0.2.gif)
