FROM golang:1.22

RUN apt-get update \
    && apt-get install -y sqlite3 \
    && rm -rf /var/lib/apt/lists/*