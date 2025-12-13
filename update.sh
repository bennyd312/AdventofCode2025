#!/bin/bash
timestamp=$(date "+%Y-%m-%d %H:%M:%S")
git add .
git commit -m "$1" -m "$timestamp"
git push