#!/bin/bash

set -e

host="$1"
shift
port="$1"
shift
cmd="$@"

until nc -z "$host" "$port"; do
  >&2 echo "Server is not available yet - waiting..."
  sleep 1
done

>&2 echo "Server is up - executing command: $cmd"
exec "$cmd"
