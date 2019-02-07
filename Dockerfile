FROM golang:1.11.5-stretch@sha256:3d60c5d680a8bc7ea701d7cd11e74202891f251709405f240729845e9fbce767

WORKDIR /go/src/github.com/knative/eventing

COPY vendor ./
#COPY Gopkg.* ./
#RUN dep ensure

COPY . .

WORKDIR /go/src/github.com/knative/eventing/contrib/kafka/cmd/dispatcher

RUN sed -i 's/zap.NewDevelopment()/zap.NewProduction()/' main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-w -extldflags "-static"'

FROM scratch

COPY --from=0 /go/src/github.com/knative/eventing/contrib/kafka/cmd/dispatcher/dispatcher /usr/local/bin/dispatcher

ENTRYPOINT ["dispatcher"]
