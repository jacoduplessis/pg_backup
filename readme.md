# pg_backup

Requires `pg_dump` to be available on the system.

Optional database connection string can be provided via `DATABASE_URL` environment variable,
either in URI format or normal format.

The default is `host=/var/run/postgresql/ sslmode=disable` to connect
with peer authentication on a normal Ubuntu server. For usage without
configuration the program therefore needs to be run with the `postgres` user.

Write a custom format `pg_dump` file in the format `<name>_<date>.pg_dump`
of all databases in the current directory.

No error handling is done: fails when an error is encountered.

Will copy any output from `pg_dump` to stdout.






