#!/bin/bash
if [ $# -ne 2 ]; then
  echo "init script to create go template for a given aoc2021 day including input"
  echo "usage: init <day> <one of go's basic types>"
  exit
fi

gofile="../go/day$1/input.go"
if [ -f $gofile ]; then
  read -r -p "input for day$1 already exists, overwrite? [y/n]" response
  if [[ "$response" =~ ^([Nn])$ ]] ; then
    exit 0
  fi
fi

tmpfile=$( mktemp )
token=$( cat aocsession.user )

echo "creating input.go"
curl --cookie "session=$token" https://adventofcode.com/2021/day/$1/input -o $tmpfile
mkdir -p ../go/day$1
cat <<EOT > $gofile
package main

var input = []$2{
EOT
if [ "$2" = "string" ]; then
  sed 's/.*/    "&",/' $tmpfile >> $gofile
else
  sed 's/.*/    &,/' $tmpfile >> $gofile
fi
echo '}' >> $gofile
rm "$tmpfile"
if [ ! -f "../go/day$1/main.go" ]; then
  echo "creating main.go"
  cat << EOT > ../go/day$1/main.go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(input)
}
EOT
fi

if [ ! -f "../go/day$1/go.mod" ]; then
  echo "creating go.mod"
  cat << EOT > ../go/day$1/go.mod
module main

go 1.17
EOT
fi
echo "done....."
