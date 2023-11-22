# User Microservice
Microservice designed to manage user entities within the Terasehat system. Provides basic functionalities such as user registration, authentication, and user information management.

---

## Architecture Layer
```
ENDPOINT
└── routing
<<<<<<< Updated upstream
    └── handlers
        ├── usecases
        └── repositories
=======
    ├── handlers
    │   └── usecases
    │       └── repositories
    ├── constants
    └── utilities
>>>>>>> Stashed changes
```
- **Endpoint:** Provides API endpoints for interacting with this microservice.

  - **Routing:** Responsible for routing requests to the appropriate handlers.

    - **Handlers:** Handles HTTP requests and calls functions from the service layer.

      - **Usecases:** Contains business logic for user entities. Handles validation, data transformation, and interacts with the repository.

      - **Repositories:** Responsible for direct interaction with the data storage (database) to retrieve, store, and delete user data.

---

## Branching
```
├── main (production live)
├── staging (development live)
└── development (development local)
```
---

## Getting Started
To get started with the User Microservice, follow the steps below:

### Prerequisites **TODO
- Go >= v1.21
- MySQL
- Goswagger v2
- Dependencies

### Installation **TODO
1. Clone the repository: `git clone https://XXX/user-microservice.git`
2. Navigate to the project directory: `cd user-microservice`
3. Install dependencies: `go mod download`
4. Configure the database connection in the `config` directory.
5. Run the microservice: `go run main.go`

---

## API Endpoints **TODO
Document the available API endpoints and their functionalities.
- `POST /users`: Register a new user.
- `POST /login`: Authenticate and log in a user.
- ...

---

## Contributing **TODO
If you'd like to contribute to the development of the User Microservice, please follow the guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).

## License **TODO
This project is licensed under the [MIT License](LICENSE).

## Acknowledgments **TODO
- Mention any contributors or libraries used in this project.
- ...
