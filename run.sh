#!/bin/zsh

if lsof -i :8080 >/dev/null; then
  echo "Port 8080 is already in use. Killing process..."
  fuser -k 8080/tcp
fi

npx update-browserslist-db@latest --quiet || true

$(go env GOPATH)/bin/air -c .air.toml & npx tailwindcss -i 'web/static/css/main.css' -o 'web/static/css/style.css' --watch
