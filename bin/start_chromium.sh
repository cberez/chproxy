#!/bin/bash

(chromium --headless --disable-gpu --remote-debugging-port=9222 http://localhost &) > /dev/null 2>&1
