# hex-pulsar
Simple attempt to get pulsar working in hex/explicit architecture.

The producer and consumer assume a local pulsar is running on port 6650.

Run pulsar using docker:
```sh
$ docker run -it -p 6650:6650  -p 8080:8080 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:2.9.1 bin/pulsar standalone
```

Run the producer and consumer.

