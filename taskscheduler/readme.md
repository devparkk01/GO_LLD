Problem: Design a Distributed Task Queue System
Scenario: You're working on a large-scale e-commerce platform that needs to process various background tasks asynchronously,
such as sending emails, generating reports, and processing orders. Design a distributed task queue system that can handle a high 
volume of tasks, ensure reliability, and provide scalability.

Requirements: 
1. The system should be able to handle different types of tasks with varying priorities.
2. Tasks should be processed in a distributed manner to ensure scalability.
3. The system should be fault-tolerant and able to recover from worker node failures.
4. It should support task scheduling (both immediate and delayed execution).
5. The system should provide a way to monitor task status and retry failed tasks.
6. It should be able to handle a high throughput of tasks (thousands per minute).
