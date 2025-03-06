build:
    docker build -t chat-app .

run:
    docker run -p 8080:8080 chat-app
