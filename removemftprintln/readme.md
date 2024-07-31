# Removing fmt.Println Statements from code

So, my code is nearly always chock-full of fmt.Printlns that are there for debugging - the go equivalent of console.log()

When writing server apps, I try to make sure I use "Log" for the interesting log-worthy things, but then have to remember (or not) to clean up all those debugging messages.

It occurred to me that it would be nice to automate the removal of fmt.Printlns

### A bash script
```
#!/bin/bash

# This script removes all fmt.Println statements from Go source files.

find . -name "*.go" -type f -exec sed -i '/fmt\.Println/d' {} +

echo "Removed all fmt.Println statements"

```
This has to be chmodded so it's executable - 

```
chmod +x remove_fmt_println.sh
```

## How about a github action

What if we could automate this on github - Create a new file in the repository at .github/workflows/remove_fmt_println.yml:

```
name: Remove fmt.Println

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  remove-fmt-println:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: '^1.16'

      - name: Run script to remove fmt.Println
        run: |
          sudo apt-get update && sudo apt-get install -y findutils
          chmod +x ./remove_fmt_println.sh
          ./remove_fmt_println.sh

      - name: Commit changes
        run: |
          git config --global user.name 'github-actions'
          git config --global user.email 'github-actions@github.com'
          git add .
          git commit -m "Remove fmt.Println statements"
          git push
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```