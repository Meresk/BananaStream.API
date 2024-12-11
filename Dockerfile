# Dockerfile for API
FROM ubuntu:22.04

WORKDIR /app
COPY ../courseProj/api .
COPY wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it 

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 3000

CMD ["wait-for-it", "courseproj-db-1:3306", "--timeout=30", "--", "./bnnstrmAPI"]

