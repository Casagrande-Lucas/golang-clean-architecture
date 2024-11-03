# Go Clean Architecture API

![Go](https://img.shields.io/badge/Go-1.20-blue.svg)
![Docker](https://img.shields.io/badge/Docker-20.10.7-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Overview

Welcome to the **Go Clean Architecture API**! This project serves as a foundational template for building scalable and maintainable Go applications. It leverages **Clean Architecture** principles combined with **SOLID** design principles to ensure a robust and flexible codebase.

## Features

- **Clean Architecture & SOLID Principles:** Ensures separation of concerns, making the codebase easy to maintain and extend.
- **User CRUD Operations:** Create, Read, Update, and Delete functionalities for user management.
- **JWT Authentication:** Secure authentication mechanism using JSON Web Tokens.
- **Middleware:** Handles authentication and other cross-cutting concerns seamlessly.
- **Dockerized Setup:** Easily containerize the application with Docker and orchestrate services using Docker Compose.
- **Persistent Storage:** Utilizes Docker volumes for database persistence.

## Architecture

The project is structured following **Clean Architecture**, which divides the application into distinct layers:

1. **Domain Layer:** Contains business logic and entities.
2. **Use Case Layer:** Implements application-specific business rules.
3. **Interface Layer:** Handles interactions with external systems (e.g., databases, APIs).
4. **Frameworks & Drivers Layer:** Includes frameworks, tools, and external interfaces like the HTTP server.

This separation ensures that each layer is independent and can be modified without affecting others, promoting scalability and testability.

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [JWT](https://jwt.io/) - JSON Web Tokens for authentication
- [Docker](https://www.docker.com/) - Containerization
- [Docker Compose](https://docs.docker.com/compose/) - Orchestrating multi-container Docker applications
- [PostgreSQL](https://www.postgresql.org/) - Relational database

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go (1.22+)](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/downloads)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone git@github.com:Casagrande-Lucas/golang-clean-architecture.git
   cd golang-clean-architecture
   
2. **Install Dependencies**

    ```bash
    go mod download

## Configuration

1. **Copy the Example Environment File**
    ```bash
    cp .env.example .env

2. **Update the .env file and docker-compose.yml**
   <p>Open the .env file in your preferred text editor and update the variables as needed.</p> 
   <p>There are two pre-installed databases, one being MySQL and the other Postgres. In the .env file we have two variable blocks for both databases, postgres by default is commented out, if you want to change it, just comment out the mysql variables and uncomment the postgres ones.</p>
   <p>In the docker-compose.yml file we have the two databases again with separate volumes, and once again you will find the postgres information commented out, to change the database, just continue commenting one and uncommenting the other and changing the dependency (depends_on:) from "app" to the desired database.</p>

## Running the Application

1. **Build and Start Services**

   ```bash
   docker-compose up -d --build

2. **Accessing the API**

   Once the services are up and running, the API will be accessible at [http://localhost:8080](http://localhost:8080)

3. **Stopping the Services**
   
   ```bash
   docker-compose down
