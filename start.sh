#!/bin/bash

python3 services/booking_reference_service/booking_reference_service.py &
cd services/train_data_service
python3 start_service.py

trap 'kill -9 $(jobs -p)' SIGINT SIGTERM EXIT

