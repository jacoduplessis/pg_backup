# pg_backup

Requires `pg_dump` to be available on the system.

Reads connection information from `DATABASE_URL` environment variable.

Write a custom format `pg_dump` file in the format `<name>_<date>.pg_dump`
of all databases in the current directory.

No error handling is done: fails when an error is encountered.

Will copy any output from `pg_dump` to stdout.






