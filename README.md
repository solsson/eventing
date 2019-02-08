
# Kafka dispatcher

Standalone kafka<->REST bridge based on Knative Eventing source.

To run the PoC you currently need:

 * Kafka bootstrap on localhost:9094
 * kafkacat
 * node.js
 * go with with this source cloned to $GOPATH/src/github.com/knative/eventing

Start the dispatcher using this fork's patched main:

```
cd contrib/kafka/cmd/dispatcher
go build && ./dispatcher
```

Start the test service that logs requests:

```
cd kafka-dispatcher-testserver/
npm install --ignore-scripts
node ./index.js | bunyan
```

Produce an event (watch testserver's logs for the result):

```
echo '{"json":true}' | kafkacat -b localhost:9094 -P -t knative-eventing-channel.dummyns1.dummyservice-trigger
```

The service can also be used to produce to kafka through REST:

```
curl --data "{\"date\":\"$(date)\"}" -H 'Host: dummyservice.dummyns1.whatever' -v http://localhost:8080/
kafkacat -b localhost:9094 -C -t knative-eventing-channel.dummyns1.dummyservice
```

Make sure we can docker build the same main, for use as a sidecar:

```
docker build -t kafka-dispatcher-standalone .
```

Next step would be to investigate what kind of guarantees we can get.
See https://github.com/Yolean/kafka-transform-nodejs-runtime.

# Knative Eventing

[![GoDoc](https://godoc.org/github.com/knative/eventing?status.svg)](https://godoc.org/github.com/knative/eventing)
[![Go Report Card](https://goreportcard.com/badge/knative/eventing)](https://goreportcard.com/report/knative/eventing)

This repository contains a work-in-progress eventing system that is designed to
address a common need for cloud native development:

1. Services are loosely coupled during development and deployed independently
2. A producer can generate events before a consumer is listening, and a consumer
   can express an interest in an event or class of events that is not yet being
   produced.
3. Services can be connected to create new applications
   - without modifying producer or consumer, and
   - with the ability to select a specific subset of events from a particular
     producer.

For complete Knative Eventing documentation, see
[Knative eventing](https://github.com/knative/docs/tree/master/eventing) or
[Knative docs](https://github.com/knative/docs/) to learn about Knative.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md),
[DEVELOPMENT.md](./DEVELOPMENT.md) and
[Knative WORKING-GROUPS.md](https://github.com/knative/docs/blob/master/community/WORKING-GROUPS.md#events).

The planned project releases are covered in this
[roadmap](https://docs.google.com/document/d/1z0z412rL9FsBsF8kwKxG6w7sflJLKe9hIECkc7jWfOY/edit#).
