#!/bin/bash

"$@" &
EXE_PID=$!
open -a firefox -g http://localhost:${PORT}
FF_PID=$!
read -p "Press ENTER to EXIT" -s
echo
kill ${EXE_PID}
kill ${FF_PID}
echo Bye!
