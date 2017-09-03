FROM laincloud/golang:1.8.3

RUN go get -u github.com/golang/dep/cmd/dep \
    && go get -u github.com/go-swagger/go-swagger/cmd/swagger

WORKDIR $GOPATH/src/github.com/laincloud/todomvc
COPY . .
# RUN dep ensure
RUN mkdir -p gen \
    && go install github.com/laincloud/todomvc/gen/cmd/todomvc-server

WORKDIR /lain/app
RUN mv $GOPATH/bin/todomvc-server /lain/app/ \
    && mv $GOPATH/src/github.com/laincloud/todomvc/local.json /lain/app/ \
    && mv $GOPATH/src/github.com/laincloud/todomvc/frontend/dist /lain/app/dist

EXPOSE 8080
CMD ["/lain/app/todomvc-server", "--host=0.0.0.0", "--port=8080", "--config=/lain/app/local.json"]