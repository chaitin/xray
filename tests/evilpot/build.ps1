$targets = @(
    @{GOOS="darwin"; GOARCH="amd64"; Output="evilpot_darwin_amd64"},
    @{GOOS="darwin"; GOARCH="arm64"; Output="evilpot_darwin_arm64"},
    @{GOOS="linux"; GOARCH="386"; Output="evilpot_linux_386"},
    @{GOOS="linux"; GOARCH="amd64"; Output="evilpot_linux_amd64"},
    @{GOOS="linux"; GOARCH="arm64"; Output="evilpot_linux_arm64"},
    @{GOOS="windows"; GOARCH="amd64"; Output="evilpot_windows_amd64.exe"}
)

foreach ($target in $targets) {
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    $output = $target.Output

    Write-Host "Building for $($env:GOOS)/$($env:GOARCH)..."
    go build -o $output .

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Successfully built $output"
    } else {
        Write-Host "Failed to build $output"
    }
}

Remove-Item env:GOOS
Remove-Item env:GOARCH
