﻿timer
=====

Utility for Windows written in GO to measure the time between two events.
Timers are persisted in the Windows registry `HKCU\Software\Tischer` as follows: 

Name    Type        Data
----    ----        ----
t1      REG_QWORD   13de77095f0a6014

Data is the number of nanoseconds elapsed since January 1, 1970 UTC.

### Compile

Tested with GO 1.4.2. There are no dependencies.

~~~
go build
~~~

### Usage

~~~
Usage of timer:
  -clear=false: clear all timers
  -read="REQUIRED": read timer (elapsed time)
  -start="REQUIRED": start timer
  -stop="REQUIRED": stop timer and print elapsed time
~~~

### Example

~~~
U:\src\timer>timer -start t1
Starting timer t1

U:\src\timer>timer -read t1
Elapsed time (t1): 5.9200225s

U:\src\timer>timer -start t2 -read t1
Starting timer t2
Elapsed time (t1): 56.3191111s

U:\src\timer>timer -start t3 -read t1 -stop t2
Starting timer t3
Elapsed time (t1): 1m30.6471884s
Elapsed time (t2): 34.3280773s

U:\src\timer>timer -clear
All timers deleted
~~~
