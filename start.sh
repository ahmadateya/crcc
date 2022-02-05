#!/bin/bash


cp .env.example .env
cp frontend/.env.example frontend/.env
cp backend/config.yml.example backend/config.yml
sudo docker-compose up