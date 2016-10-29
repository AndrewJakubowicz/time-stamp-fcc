# time-stamp-fcc

Free Code Camp /challenges/timestamp-microservice app written in Go.

Running at `https://time-stamp-fcc.herokuapp.com`.

Can accept general date form:

[time-stamp-fcc.herokuapp.com/1 May, 2002](https://time-stamp-fcc.herokuapp.com/1 May, 2002)

Can accept Unix date form:

[time-stamp-fcc.herokuapp.com/1231231311](https://time-stamp-fcc.herokuapp.com/1231231311)


Returns null for garbage inputs or lack of input:

- [time-stamp-fcc.herokuapp.com/garbage](https://time-stamp-fcc.herokuapp.com/garbage)
- [time-stamp-fcc.herokuapp.com/1 1 1](https://time-stamp-fcc.herokuapp.com/1%201%201%20)
