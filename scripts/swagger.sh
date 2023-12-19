if [ ! -d "gen" ]; then
  mkdir -p gen gen/models gen/restapi
fi

swagger generate server -t gen --exclude-main -A travel-booking-portal
swagger generate client -t gen -A travel-booking-portal
