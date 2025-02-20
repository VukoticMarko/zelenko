# Zelenko an Ecology App with CRDT

A distributed map-based application that allows users to mark trash can locations so everyone knows where to dispose of garbage. This project leverages Conflict-free Replicated Data Types (CRDTs) to provide real-time, conflict-free updates across a distributed system.

## Overview

This is implementation of my diploma thesis "Оцењивање зелених објеката на мапама употребом реплицираних типова података без конфликта", this project demonstrates how to build a consistent distributed application using replicated data types. In this app, users can add, update, or remove green objects (such as trash cans or recycling containers) on a global map, while the system automatically resolves conflicts—even when multiple users interact concurrently.

Key concepts include:

- **Distributed Systems:** A network of interconnected nodes (servers) that work together to ensure high availability, fault tolerance, and low latency.
- **CRDTs:** Data structures (like the G-Counter) designed to reconcile concurrent operations without a central coordinator, ensuring that all replicas eventually converge to the same state.
- **Replication & Synchronization:** Data is continuously replicated among nodes using both state-based and operation-based techniques to maintain consistency despite network partitions or failures.

## Key Features

### Interactive Map Interface

- **Add, update, and remove trash can locations** with an easy-to-use map.
- **Display details** such as location name, coordinates, street, city, and country.

### Real-Time Conflict Resolution

- **Employs CRDTs (G-Counter)** to merge updates seamlessly.
- **Supports concurrent modifications** without the need for heavy synchronization protocols.

### Robust Distributed Architecture

- **Combines fast in-memory storage (Redis) with persistent storage (PostgreSQL)** for durability.
- **Designed for distributed environments** where nodes can be spread across different geographic locations.

### Scalability and Fault Tolerance

- **Uses replication strategies** (master-slave and Redis clustering) to ensure high availability.
- **Handles network partitions and node failures gracefully** while maintaining a consistent global view.

## Technologies Used

- **Backend:** Golang  
  Implements the core business logic, including CRDT operations and replication methods.
- **Frontend:** React  
  Delivers a responsive user interface that interacts with the backend API.
- **Containerization:** Docker  
  Facilitates deployment and scalability through containerized services.
- **Data Stores:**
  - **Redis:** In-memory database used for fast access to CRDT states and temporary data.
  - **PostgreSQL:** Persistent relational database for long-term storage and backup.
- **CRDT Implementation:**  
  Uses the G-Counter (Grow-only Counter) to manage the scoring of green objects. It supports basic operations such as Increment, Decrement (via PN-Counter variants), and value retrieval to ensure convergence regardless of the order of operations.

## System Architecture

The application is structured as a multi-layered distributed system:

### Decentralized Nodes

- Each node functions autonomously by processing local user interactions, updating its local CRDT state, and then propagating those changes to peer nodes. This approach minimizes latency and avoids a single point of failure.

### Hybrid Storage Strategy

- **Redis** is used as the primary store for real-time data, ensuring quick reads/writes.
- **PostgreSQL** serves as the long-term persistent store, with data periodically migrated from Redis.

### Replication & Synchronization

- The system uses both state-based and operation-based replication strategies.
- In a **master-slave replication** scenario, changes made in the master are propagated to all slave nodes.
- **Redis clustering** is configured so that multiple Redis instances work together, ensuring that even if one instance fails, others keep the system running.

## Implementation Details

### CRDT (G-Counter)

- **G-Counter:**  
  A grow-only counter that can be incremented (or combined with a decrement variant, such as the PN-Counter) to allow users to rate or assess the "green" quality of an object.
  - **Increment:** Increases the counter, reflecting positive votes.
  - **Decrement:** (In extended variants) Decreases the counter to represent negative feedback.
  - **Convergence:** Thanks to the mathematical properties of CRDTs, no matter the order or timing of operations across nodes, all replicas will eventually converge to the same value.

### Node Structure

The project follows a modular structure in the backend to separate concerns:

- **Main Package:** Initializes the router, repositories, services, and controllers.
- **Model Package:** Contains the domain models and business logic.
- **Controller Package:** Acts as the intermediary between HTTP requests and business logic.
- **Service Package:** Implements the core functionality and CRDT operations.
- **Repository Package:** Handles database interactions with Redis and PostgreSQL.
- **Router Package:** Defines the HTTP endpoints and request routing.

### Replication Strategy

#### Master-Slave Replication

- The Golang code establishes connections with the local Redis instance and its slave replicas. When an update occurs, a replication method is triggered to update all replicas.

#### Redis Cluster Configuration

- During setup, configuration files are adjusted (e.g., setting `cluster-enabled yes` and specifying `replicaof` parameters) so that multiple Redis instances form a cluster. For example, a command such as:

  ```bash
  redis-cli --cluster create 127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:6381 --cluster-replicas 1
  ```
  groups the instances together, ensuring efficient replication and fault tolerance.

## Getting Started

### Prerequisites
- Git
- Docker
- Docker Compose

### Installation

**Clone the Repository:**

```bash
git clone https://github.com/VukoticMarko/zelenko
cd zelenko
```

**Build and Run with Docker Compose:**
```
docker-compose up --build
```
This command builds and starts the backend (Golang API), frontend (React app), Redis, and PostgreSQL containers.

## Configuration

### Backend
Configuration files and environment variables are located in the `backend/config/` directory.

### Frontend
Modify settings in the `frontend/src/` directory as needed.

### Database
Ensure that PostgreSQL and Redis configurations in `docker-compose.yml` match your local environment if you are running outside Docker.

## Usage

### Access the App
Navigate to [http://localhost:3000](http://localhost:3000) (or your specified port) in your web browser.

### Interact with the Map
Use the map interface to add, update, or remove trash can locations.

### Automatic Synchronization
All updates are processed locally and then synchronized across nodes via CRDT, ensuring eventual consistency.

## Contributing

Contributions are welcome! To contribute:

1. **Fork the repository.**
2. **Create a new branch** for your feature or bug fix.
3. **Commit your changes** with clear commit messages.
4. **Open a pull request** detailing your modifications.

## Future Enhancements

- **Offline Support:** Enable local storage and automatic re-synchronization when connectivity is restored.
- **Enhanced Map Features:** Integrate additional map layers, filtering options, and GPS navigation.
- **Improved Security:** Add authentication, authorization, and user-specific preferences.
- **Mobile Application:** Develop a mobile version for on-the-go access.
- **Monetization:** Explore premium features and advertisement integration.
- **Friends:** Add friend lists and profile showcases with contribution to the map.

