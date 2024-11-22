# 9ja-State API

The 9ja-State API is a lightweight RESTful API built using Golang and the Gin framework. It provides endpoints to access information about Nigeria's 36 states and their respective Local Government Areas (LGAs). This project is designed to make it easy for developers to integrate state and LGA data into their applications.


## Features

- Retrieve all Nigerian states and their LGAs.

- Search for specific states or LGAs.

- Organized with proper versioning (/api/v1).

- Built with the Gin framework for high performance.

- Simple and developer-friendly interface.



## Getting Started

#### Prerequisites

- Go 1.20+: Ensure you have Go installed on your system.

- Git: Clone the repository to your local machine.


#### Installation

1. Clone the repository:

git clone https://github.com/your-username/9ja-state.git
cd 9ja-state


2. Install dependencies:

go mod tidy


3. Run the application:

go run main.go


4. Access the API in your browser or API client (e.g., Postman) at:

http://localhost:8080


## Endpoints

Base URL

http://localhost:8080/api/v1

Endpoints

1. Get All States

URL: /states

Method: GET

Description: Retrieve all states and their LGAs.

Response:

{
  "states": [
    {
      "name": "Abia",
      "lgas": ["Aba North", "Aba South", ...]
    },
    ...
  ]
}


2. Search States or LGAs

URL: /search

Method: GET

Query Params:

q (string): The name of the state or LGA to search for.


Example: /search?q=lagos

Response:

{
  "results": [
    {
      "type": "state",
      "name": "Lagos",
      "lgas": ["Agege", "Ajeromi-Ifelodun", ...]
    }
  ]
}


3. Serve Static HTML

URL: /index.html

Method: GET

Description: Displays a static HTML page for project information or as a basic homepage.

## Contributing

We welcome contributions to the 9ja-State API! To contribute:

1. Fork the repository.


2. Create a new branch for your feature or bug fix.


3. Commit your changes and push them to your fork.


4. Submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.


## Author

Developed by [Abdul Console](https://abdulconsole.com.ng), a passionate software engineer with expertise in Golang and cybersecurity.

GitHub: [@AbdulConsole](https://github.com/AbdulConsole)

Twitter: [@abdulconsole](https://x.com/AbdulConsole)
