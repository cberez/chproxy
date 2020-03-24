#!/bin/bash

# Stops the first process matching
ps aux | \
  grep "chromium --headless --disable-gpu --remote-debugging-port=9222 http://localhost" | \
  awk 'NR==1{print $2}' | \
  xargs kill
