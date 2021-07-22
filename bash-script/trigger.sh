#!/usr/bin/bash

echo "what is the enemy passcode?"

read loginpass

expect expect.exp $(arp-scan -l | grep Raspberry | awk '{print $1}') root $loginpass
