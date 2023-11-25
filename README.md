# User Microservice
Microservice designed to manage user entities within the Terasehat system. Provides basic functionalities such as user registration, authentication, and user information management.

---

## Getting Started
To get started with the User Microservice, follow the steps below:

### Prerequisites
- Go >= v1.21
- MySQL
- Goswagger v2
- Dependencies

### Installation **TODO
1. Clone the repository: `git clone https://XXX/user-microservice.git`
2. Navigate to the project directory: `cd user_micro`
4. Re-generate goswagger by the makefile command: `make gen`
3. Install dependencies: `go mod tidy`
4. Configure the the YAML file to `app.yaml` in in the `config` directory for project **Environment**.
5. Run the microservice by the makefile command: `make run` or `make all-dev`

---

## Architecture Layer
```
ENDPOINT
└── routing
    └── handlers
        └── usecases
            └── repositories
```
- **Endpoint:** Provides API endpoints for interacting with this microservice.

  - **Routing:** Responsible for routing requests to the appropriate handlers.

    - **Handlers:** Handles HTTP requests and calls functions from the service layer.

      - **Usecases:** Contains business logic for user entities. Handles validation, data transformation, and interacts with the repository.

        - **Repositories:** Responsible for direct interaction with the data storage (database) to retrieve, store, and delete user data.

---

## Branching

Default branch:
```
├── main (production live)
├── staging (development live)
└── development (development local)
```

## Committing ***TODO

- Don't commit file contains some credentials!
- ...

---

## API Endpoints
Document the available API endpoints and their functionalities.
- You can check the documentation API from `swagger.yaml` file at `./api` via swagger UI.

---

## Contributing
If you'd like to contribute to the development of the User Microservice, please follow the guidelines and existing code styles.

### Basic Knowledge

Just make sure you need to following the existing code styles and the architecture.

- create new endpoint

  You can start to create a new endpoint by following this steps:
  
  - Define a new endpoint at `./api/swagger.yaml`.
  - Generate your new endpoint by the makefile command: `make gen`. And dont always remember **Don't Forget** to regenerate server and client goswagger After doing some changes at `swagger.yaml` file.
  - Define a new route at `./internal/routes`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - Define a new handler at `./internal/handlers`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - If Needed. Define a new usecase at `./internal/usecases`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - If Needed. Define a new repository at `./internal/repositories`. You can create new file for new module or feature (eg. registration, authentication, migration, integration etc).



---

## License **TODO
This project is licensed under the [Developer License](LICENSE).

## Acknowledgments
- Mention any contributors or libraries used in this project.
- ...
