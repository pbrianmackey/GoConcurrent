#Message Queue

##options

- ZeroMQ (Open source)
- MSMQ
- Websphere MQ

Message queue over the cloud:  Azure service bus and AWS

##Messaging patterns

- Fire and forget: Client sends message and doesn't care what happens after that.  
- Request/Response:  Client sends message and receives a reply back from its dedicated response queue.
- Publish/Subscribe:  Publisher sends message to queue and queue forwards message onto all subscribers.
Good pattern for event driven processing.
