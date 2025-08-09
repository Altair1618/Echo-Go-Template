# Echo Go Template

This is repository for my implementation template for Echo Framework in Go, which is used to build REST APIs. It includes basic setup, routing, and API documentation generation using Swag.

## Table of Contents

- [Echo Go Template](#echo-go-template)
    - [Table of Contents](#table-of-contents)
    - [Getting Started](#getting-started)
        - [Prerequisites](#prerequisites)
        - [Installation](#installation)
    - [Usage](#usage)
    - [Acknowledgements](#acknowledgements)

## Getting Started

This project is a template for building REST APIs using the Echo framework in Go. It provides a basic structure and setup to help you get started quickly.

### Prerequisites

Before setting up the project, make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Makefile](https://www.gnu.org/software/make/) (optional, but recommended for running commands provided in the Makefile)
- [Air](https://github.com/air-verse/air) (optional, but recommended for live reloading during development). You can install it using the following command:
  ```bash
  go install github.com/air-verse/air@latest
  ```
- [Swag](https://github.com/swaggo/swag) (optional, but needed for regenerating the API documentation). You can install it using the following command:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

### Installation

1. Clone this repository to your local machine
2. Navigate to the project directory
3. Run the following command to install the dependencies:
   ```bash
   go mod tidy
   ```
4. Create a `.env` file in the root directory and set the required environment variables. You can use the provided `.env.example` as a reference.
5. To regenerate the API documentation (optional), run the following command:
   ```bash
   make generate-docs
   ```

## Usage

1. Run the following command to start the application:
   ```bash
   make all
   ```
   or if you have Air installed, you can use:
   ```bash
   air
   ```
2. The application will start running on the specified port (default is `8080`). You can access the API documentation at `http://localhost:8080/swagger/index.html`.

## Acknowledgements

This project is being developed by:

- [Farhan Nabil Suryono](https://github.com/Altair1618)
