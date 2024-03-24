# golang-wallet-api

A simple wallet application that exposes RESTful API.  
This app is for demo purpose only, and is very basic.

## Available Endpoints

These endpoints are the basic endpoints needed to have a simple wallet API.

- `POST /accounts` Create an account
- `GET /accounts/{accountID}` Retrieve an account
- `POST /accounts/{accountID}/transfer` Transfer an amount from an account's wallet to another
- `GET /accounts/{accountID}/transactions` Retrieve transactions of an account

Swagger is available when application is run.

## Development

This application is still in its early stages of development. There is no production deployment pipeline yet, but there is an existing local development environment that is easy to setup.  

The local environment setup is deployed into a kubernetes cluster (preferrably on host machine) using Kustomize and orchestrated by [Skaffold](https://skaffold.dev/).  

The local env consists of 2 key deployments and 1 job. The application, which is the actual application that exposes the wallet API. The database, which is a postgresql. And 1 job that handles the database migration.  

All the basic commands to run the application is already available in the Makefile. In theory, you dont have to configure anything for the local environment to run, and just start developing.

### Requirements

> [!IMPORTANT]  
> This environment setup was only tested to run in MacOs (Sonoma) host machine

To successfully run the application on your host machine, you will need the following:

- Go >= 1.21.0
- Make
- Docker
- Kubernetes cluster & kubectl

For the Kubernetes cluster, I recommend using [Colima](https://github.com/abiosoft/colima) for easy setup. Or you can also use other ones like Minikube, etc.

#### Quick guide

```sh
# Homebrew
brew install go
brew install make
brew install docker
brew install kubectl
brew install colima

colima start --kubernetes
```

Make sure the context of your k8s cluster is set to `colima` to use it.

### Running the APP

```
make skaffold
```

This will initially download skaffold and then run `skaffold dev`.  

#### Accessing the API

After running the application, the base endpoint of the API is [http://localhost:3000](http://localhost:3000).  

Swagger API is also available via [http://localhost:3000/swagger/](http://localhost:3000/swagger/).  

There is also a Postman Collection (v2.1) in [postman/](postman/) that contains simple API calls that creates 2 accounts, transfers, and transactions list. Please refer to [Postman - Import the Postman collection](https://docs.tink.com/entries/articles/postman-collection-for-account-check#import-the-postman-collection) for instructions on how to import it to Postman app.  

If you prefer to do it the cURL way, here is an example workflow:

1. Create 2 accounts to use for the transaction
```sh
# first account assuming id will be 1
curl -X 'POST' \
  'http://localhost:3000/accounts' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "initial_balance": 300,
  "name": "Richmond Wang"
}'
# second account assuming id will be 2
curl -X 'POST' \
  'http://localhost:3000/accounts' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "initial_balance": 0,
  "name": "Skyler Chase"
}'
```

2. Transfer from one account to the other
```sh
curl -X 'POST' \
  'http://localhost:3000/accounts/1/transfer' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
      "account_id": 2,
      "amount": 100
}'
```

3. Check transactions list
```sh
# all transactions
curl -X 'GET' 'http://localhost:3000/accounts/1/transactions'
# incoming transactions
curl -X 'GET' 'http://localhost:3000/accounts/1/transactions?type=incoming'
# outgoing transactions
curl -X 'GET' 'http://localhost:3000/accounts/1/transactions?type=outgoing'
```

4. Check balance
```sh
curl -X 'GET' 'http://localhost:3000/accounts/1'
curl -X 'GET' 'http://localhost:3000/accounts/2'
```

#### Teardown

Just press C^ on the `make skaffold` running command and it will teardown everything for you. Except for the docker resources and the colima instance. For cleaning up docker, you can simply run `docker system prune`. And for colima, just run `colima stop && colima delete`.

### ORM

This application uses the [Ent.(entgo.io)](https://entgo.io/) framework which handles a lot if not all of the transactions against the database including the schema migrations.

The database schema is located at [api/ent/schema](api/ent/schema).

#### Updates to schema

After you update the schema, it is important to regenerate ent generated files to reflect the new schema on the models.

```sh
# command for re/generating entgo files from schema
make entgen
```

### Dockerfile

The application's [Dockerfile](api/Dockerfile) consists of 4 stages, for `build`, `dbmigrate`, `dev` (for local) and `dist` (for dev, stage, and prod envs). The `build` stage mostly just downloads the go mods and tests the application. The `dbmigrate` stage is for building the db schema migration tool. The `dev` stage is the one that builds the application, and runs `air` ([comstrek/air](github.com/cosmtrek/air)) which (supposedly *note in TODOs*) rebuilds the application in the container on runtime during file syncs done by skaffold. And the default stage (`dist`), just an alpine image with the application binary inside, which is inteded for use in higher environments like dev, stage, and prod.

### Tests

Currently, functional tests are written only for the handlers. In later development, it will be neccessary to add more tests like integration, performance, load testings, etc.  

Tests are located at [api/pkg/handlers/tests](api/pkg/handlers/tests).

#### Running the tests

Tests are run during the building of the container image when you run the `make skaffold` command. But you can also run them manually with:

```sh
# command for running go tests
make run-tests
```

#### Updating Mocks

When you have updates to the code and you want to generate or regenerate the mocks.

```sh
# command for re/generating test mocks
make mockery
```

### Configuration

You do not have to configure anything but if you have to, all the configuration and secret files for the local development environment are stored in [deploy/skaffold/api](deploy/skaffold/api).

Application:
 
- [api.config.env](deploy/skaffold/api/api.config.env) - This file will be stored in the k8s cluster as `configmap` and will be loaded into the deployment as Environment Variables
- [api.secrets.env](deploy/skaffold/api/api.secrets.env) - This file will be stored in the k8s cluster as `secret` and will be loaded into the application deployment as Environment Variables

Database:

- [database.env](deploy/skaffold/api/database.env) - This file will be stored in the k8s cluster as `secret` and will be loaded into the database deployment as Environment Variables

## Troubleshooting

- docker daemon not running after colima start, try setting the docker host `export DOCKER_HOST=unix:///$HOME/.colima/default/docker.sock`

## Notes / TODOs

During the development of this application, there were some issues that encountered that were not addressed fully due to time constraints.

- Add more tests. Add unit tests and application testing.
- Database model is returned from the API responses instead of its own models. The response schemas could be better.
- Schema migration can be done better. Currently migrations are done using a Job that is required by the deployments to complete before starting the application pods.
- Optimize skaffold to sync properly on specific set of files. It should only sync files and let `air` rebuild from the container instead of having skaffold rebuild the images.
- Working CICD pipeline. Although, the local env deployment manifests are just few updates away from an actual working pipeline that can deploy to a remote environment (e.g. stage/prod).
- Database persistence. The database is purge everytime skaffold is quit. This can be fixed if we attach a PV into the database deployment. This PV can be created separately before running skaffold.
- Swagger endpoint is still a bit flaky. There are some situations where the API hangs because of the swagger endpoint. Maybe better approach is to have another swagger deployment pod and load the json file into it to have it in a separate instance than the API.
- Add linters and formatters