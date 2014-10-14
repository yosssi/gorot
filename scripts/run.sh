#!/bin/sh

# Check the arguments.
if [ $# -lt 1 ]; then
  echo "Invalid argument"
  exit 1
fi

# Set values to variables.
dir=./cmd/$1
bin=$dir/$1
bin_with_args=$dir/$@

# Build an executable file.
go build -o $bin $dir

# Remove the generaged directory.
rm -rf ./testgorot

# Execute the executable file.
$bin_with_args

# Get the exit code.
code=$?

# Remove the executable file.
rm $bin

# Exit with the exit code.
exit $code
