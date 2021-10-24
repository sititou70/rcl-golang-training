#!/bin/sh
set -eu
cd $(dirname $0)

# main
current_session_name=$(tmux list-sessions | grep attached | cut -d ":" -f 1)

# server
tmux send-key -t $current_session_name.1 \
  "go run server/main.go" C-m

# client 1
tmux split-window -v -t $current_session_name
tmux send-key -t $current_session_name.2 \
  "cd client" C-m \
  "sleep 1" C-m \
  "go run client.go" C-m

# client 2
tmux split-window -h -t $current_session_name
tmux send-key -t $current_session_name.3 \
  "cd client" C-m \
  "sleep 1" C-m \
  "go run client.go" C-m
