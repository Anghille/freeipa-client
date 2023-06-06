#!/bin/bash

# Logging everything
exec > >(tee -i run_script.log)
exec 2>&1

echo "Check for Python3 version..."

# Check Python3 version
py_version=$(python3 -c 'import sys; print(sys.version_info[:])' 2>&1)
if [[ "$py_version" < "(3, 4" ]]; then
    echo "Warning: Python version is below 3.4"
else
    echo "Python version is 3.4 or above"
fi

echo "Check for Go version..."

# Check Go version
go_version=$(go version | awk '{print $3}' 2>&1)
if [[ -z "$go_version" ]]; then
    echo "Error: Go is not installed"
    exit 1
elif [[ "$go_version" < "go1.18" ]]; then
    echo "Warning: Go version is below 1.18"
else
    echo "Go version is 1.18 or above"
fi

echo "Running Python script..."

# Run Python3 script
python3 dump-errors/dump.py

echo "Running Go commands..."

# Run Go commands
cd ./gen
go run .

# Return to root directory
cd ..

echo "Script execution finished."
