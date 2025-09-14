Uber ( ride sharing system )

<!-- Ride booking -->

- The ride sharing service should allow passengers to request rides and drivers to accept and fulfill those ride requests.
- Riders can request rides by specifying pickup and drop locations and desired ride type (e.g., regular, premium).
- Riders should be able to cancel a ride before confirmation 

<!-- Driver matching -->
- The system should find the nearest available driver to the rider (there are many strategies to finding the nearest). Let's stick to the basic sqrt of distance formula. 
- Drivers should receive ride requests and have the option to accept or reject. 

The system should calculate the fare for each ride based on distance, time, and ride type.
The system should handle payments and process transactions between passengers and drivers.
The system should provide real-time tracking of ongoing rides and notify passengers and drivers about ride status updates.
The system should handle concurrent requests and ensure data consistency.