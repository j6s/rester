# Rester

A simple wrapper for [restic](https://github.com/restic/restic) that allows
backup options to be specified in a simple JSON file.

This wrapper intends to be a one-stop shop for creating new and deleting old backups.

## Installation

TODO

## Usage

```
rester backup.json
```

## Options

For an example configuration you can look at `example.json`

**General**

The following options are mandatory for a configuration file:

- `repository`: Location of the backup repository. Equivalent to the command line option `--repo` in the restic CLI.
- `sources`: Array of paths to backup. Currently these paths must be absolute.
- `password`: The password of the repository.

**Creating backups**

The following options may be used to configure the backup:

- `excludes`: Array of paths to exclude. Equivalent to the command line option `--exclude` in the restic CLI.

**Forgetting backups**

The following options may be used to configure how restic forgets old backups:

- `keepLast`
- `keepHourly`
- `keepDaily`
- `keepWeekly`
- `keepMonthly`
- `keepYearly`

For more information about these options, please refer to the [restic forget manual](https://restic.readthedocs.io/en/v0.7.3/manual.html#removing-snapshots-according-to-a-policy)


## TODO

- How to implement `make install`? Where to install binary to?
- How to structure repository? Are there go conventions?
- Allow for relative paths in sources
- Extract main functionality into public methods in order to use in other projects.
- Add all `restic backup` options to configuration file.
- Everything else that is mentioned as `TODO` comment in the source files

## Contributing

This is my first go project. I don't know the language, I don't know the environment, I am just trying to implement something that works.

Please open issues for every little thing that is wrong or could be done better. I would be super grateful for that.
