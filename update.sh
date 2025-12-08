#!/bin/bash
timestamp=$(date "+%Y-%m-%d %H:%M:%S")
git add .
git commit -m "Update from date $timestamp"
git push