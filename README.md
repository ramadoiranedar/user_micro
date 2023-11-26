# User Microservice (v1.3.1)

Microservice designed to manage user entities within the Terasehat system. Provides basic functionalities such as user registration, authentication, and user information management.

## Contributing
If you'd like to contribute to the development of the User Microservice, please follow the guidelines and existing code styles.

### Basic Knowledge

First, make sure you need to following the existing code styles and the architecture.

- semantic versioning

  - Perform semantic versioning when you make changes (major/minor/patch) to the source code by updating several files or text in: ./api/swagger.yaml, ./config/app.yaml, ./README.md

  - And don't forget to update `git tag` versioning when you `push` your changes

- create new endpoint

  You can start to create a new endpoint by following this steps:
  
  - Define a new endpoint at `./api/swagger.yaml`.
    - make sure when define response you need to use `results` as object for consumer this endpoint
    - make sure when define request parameters and response you don't make a redundant parameter and must be useful information
    - make sure you when create function using DTO object from the handler layer and so on until the last layer repository. (eg. h.HealthCheckGetV1Health(params) using DTO health_check.GetV1HealthParams)
	  - and you can passing everything you need from route params above to handler
  - Generate your new endpoint by the makefile command: `make gen`. And dont always remember **Don't Forget** to regenerate server and client goswagger After doing some changes at `swagger.yaml` file.
  - Define a new route at `./internal/routes`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - Define a new handler at `./internal/handlers`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - If Needed. Define a new usecase at `./internal/usecases`. You can create new file for new module or feature (eg. registration, authentication, etc).
  - If Needed. Define a new repository at `./internal/repositories`. You can create new file for new module or feature (eg. registration, authentication, migration, integration etc).

---

## Getting Started
To get started with the User Microservice, follow the steps below:

### Prerequisites
- Go >= v1.21
- MySQL
- Goswagger v2
- GNU make (Makefile)
- Dependencies

### Installation **TODO
1. Clone the repository: `git clone https://XXX/user-microservice.git`
2. Navigate to the project directory: `cd user_micro`
4. Re-generate goswagger by the makefile command: `make gen`
3. Install dependencies: `go mod tidy`
4. Configure the the YAML file to `app.yaml` in in the `config` directory for project **Environment**.
5. Run the microservice by the makefile command: `make run` or `make all-dev`

### Makefile

```
make make-gen: make folder /gen
make remove-gen: remove folder /gen
make remove-pkg: remove folder /pkg
make remove-cmd: remove folder /cmd
make remove-executable-server: remove file executable server at root directory /
make validate: validation OpenAPI swagger yaml file
make new-server-default: [ONCE] init for the first time generate server default main.go file
make new-server-custom: [ONCE] init for the first time generate server custom main.go file from ./main.go.dist
make gen: generate server and client after run the new-server
make gen-server-configureapi: generate server excluding main.go file with regenerate configureapi
make gen-server: generate server excluding main.go file without regenerate configureapi
make gen-client: generate client after generate server for this app and you can add more than one client for this
make clean: go cleaning and remove existing executable file server
make build: generate build file executable server
make run: serve REST API by the executable file server
make run-dev: run the server for development without build the file executable server
make all-dev: for easy generate all and run server for development
make run-api-doc: serve UI of API documenations Swagger OpenAPI
```
---

## Branching

Default branch:
```
├── main (production live)
├── staging (development live)
└── development (development local)
```

## Committing

- Don't commit file contains some credentials!
- commit format: [WORK_BRANCH][YOUR_UPDATED_VERSION][FEATURE/BUG/PATCH][TICKET_NUMBER][COMMIT_DESCRIPTION]
- ...

## Versioning

Don't forget to always update `git tag` versioning when you `push` your changes:

- `git tag` to check latest version
- `git tag 1.2.0` to update your version
- `git push origin 1.2.0` to push your tag

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

## API Endpoints
Document the available API endpoints and their functionalities.
- You can check the documentation API from `swagger.yaml` file at `./api` via swagger UI.

---

## License **TODO
This project is licensed under the [Developer License](LICENSE).
