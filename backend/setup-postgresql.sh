#!/bin/bash

# Start the service
service postgresql start

# Set up the database
sudo -u postgres psql --command "ALTER USER postgres WITH SUPERUSER PASSWORD '123456';"
sudo -u postgres createdb -O postgres postgres

# Start a bash shell
bash
