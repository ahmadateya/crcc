#!/bin/bash

cd /app && go build -o goapp
chmod +x /app/goapp

./goapp