
The Pub-Sub system should allow publishers to publish messages to specific topics.
Subscribers should be able to subscribe to topics of interest and receive messages published to those topics.
The system should support multiple publishers and subscribers.
Messages should be delivered to all subscribers of a topic in real-time.
The system should handle concurrent access and ensure thread safety.
The Pub-Sub system should be scalable and efficient in terms of message delivery.




Solution-
Topic maintains a list of subscriber. A subscriber is an interface with `onMessage` method. 
Whenever a topic receives a message from publisher, it calls `onMessage` method of all the subscribers 
in its subscribers list. 
Publisher also maintains a list of topics it has registered to. 


Entities:
Message, 
Topic,
Publisher
Subscriber 