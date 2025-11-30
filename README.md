[![Build Status](https://github.com/tischda/timer/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/timer/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/timer/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/timer/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/timer/badge.svg)](https://coveralls.io/r/tischda/timer)
[![Linter Status](https://github.com/tischda/timer/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/timer/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/timer)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/timer.svg)](https://github.com/tischda/timer/releases/latest)


Windows utility to measure the time between two events.
Timers are persisted in the Windows registry:

`HKEY_CURRENT_USER\Software\Tischer\timers`

Name  | Type      | Data
----  | ----      | ----
t1    | REG_QWORD | 13de77095f0a6014

Data is the number of nanoseconds elapsed since January 1, 1970 UTC.

## Install

~~~
go install github.com/tischda/timer
~~~

## Usage

~~~
Usage: timer [OPTION] exec task
       timer [OPTION] COMMAND timer-name

 COMMANDS:

  start: start named timer
  read:  read timer (elapsed time)
  stop:  read and then clear timer
  list:  list timers
  clear: clear named timer, remove from registry
  exec:  execute task and print elapsed time

OPTIONS:

  -q, --quiet
        hide process output
  -?, --help
        display this help message
  -v, --version
        print version and exit
~~~

## Examples

~~~
C:\>timer start t1
C:\>timer read t1
Elapsed time (t1): 5.9200225s

C:\>timer start t2
C:\>timer list
[t1 t2]

C:\>timer stop t1
Elapsed time (t1): 1m30.6471884s

C:\>timer clear

C:\>timer -quiet exec "dir /s"
Total time: 91.2001ms
~~~

## Other timers

* [clTimer](http://www.cylog.org/tools/cmdline.jsp)
* [utime](http://www.rohitab.com/discuss/topic/38678-unix-time-on-windows/)
