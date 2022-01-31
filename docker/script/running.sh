#!/bin/sh
echo "Process pod start"
cd /application
exec -a btcwallet api
exit $?