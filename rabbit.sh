#!/bin/bash

case "$1" in
    start)
        echo "Starting RabbitMQ container..."
        docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 -p 61613:61613 rabbitmq-stomp
        ;;
    stop)
        echo "Stopping RabbitMQ container..."
        docker stop rabbitmq
        ;;
    logs)
        echo "Fetching logs for RabbitMQ container..."
        docker logs -f rabbitmq
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac