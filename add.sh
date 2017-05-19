#!/bin/bash

COUNTER=1
while [ $COUNTER != 200 ]
do
		 COUNTER=$[$COUNTER +1]
		 ./client -add="$COUNTER, $COUNTER"
done
