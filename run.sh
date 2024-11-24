#!/bin/zsh

$(go env GOPATH)/bin/air -c .air.toml & npx tailwindcss -i 'web/static/css/main.css' -o 'web/static/css/style.css' --watch
