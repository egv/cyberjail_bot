FROM golang:1.8
WORKDIR /app
RUN go get "github.com/go-telegram-bot-api/telegram-bot-api" && go get "github.com/mstralenya/cyberjail_bot"
ENV SRC_DIR=/go/src/github.com/mstralenya/cyberjail_bot

# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o cyberjail; cp cyberjail /app/
ENTRYPOINT ["./cyberjail"]