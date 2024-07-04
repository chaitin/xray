#!/bin/bash

targets=(
    "darwin amd64 evilpot_darwin_amd64"
    "darwin arm64 evilpot_darwin_arm64"
    "linux 386 evilpot_linux_386"
    "linux amd64 evilpot_linux_amd64"
    "linux arm64 evilpot_linux_arm64"
    "windows amd64 evilpot_windows_amd64.exe"
)

for target in "${targets[@]}"; do
    IFS=' ' read -r -a params <<< "$target"
    GOOS=${params[0]}
    GOARCH=${params[1]}
    OUTPUT=${params[2]}

    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT

    if [ $? -eq 0 ]; then
        echo "Successfully built $OUTPUT"
    else
        echo "Failed to build $OUTPUT"
    fi
done
