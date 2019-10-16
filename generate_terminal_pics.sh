#!/bin/bash

cd examples

for dir in *; do
    if [ -d "${dir}" ]; then
        echo "Generating image for ${dir}.."
        cd ${dir}
        termtosvg -c "go run main.go" -g 80x2 --template ../../assets/marker.term.theme ${dir}.svg
        mv ${dir}.svg ../../assets
        cd ../
    fi
done