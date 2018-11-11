#!/bin/bash

echo "Building binary"

GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating zip file"

zip deployment.zip main

# aws s3 cp deployment.zip s3://sl-listings-scraper
aws lambda update-function-code --function-name sl-listings-scraper --zip-file fileb://./deployment.zip --region eu-west-1

echo "cleanup!"

rm main
rm deployment.zip
