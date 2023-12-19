#!/bin/bash

if [ "$(docker ps -q -f name=travel-booking-portal-mongo-db)" ]; then

docker stop travel-booking-portal-mongo-db

docker rm travel-booking-portal-mongo-db

fi