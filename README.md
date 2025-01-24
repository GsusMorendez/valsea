# Valsea REST API

In the `statement` folder, you can find the description of the problem this program solves. Essentially, it is a REST API with 6 endpoints that allow interaction with accounts and transactions.

---

## USAGE

To start the application easily, I’ve added a Makefile with basic options like running tests, building the application, or running it. To start the application, just navigate to the project root and execute the following command:

`make run`

A Postman collection has been added to easily test the API. Only the IDs need to be updated, as they are generated randomly each time the application is run, so they must be adjusted in the request.

`Valsea.postman_collection.json`

---

## General Information

- Code organization:
  I’ve implemented a simple and intuitive package structure that allows for quickly extending the application’s features.

- `data.go` layer:
  This layer contains in-memory data. I’ve isolated it and made it accessible only through the repository, which would make it easy to replace with a database in the future. I avoided using advanced structures like maps with O(1) access since this layer is simply for handling mock data and is not crucial for this test.

- Logger:
  I’ve added the `zap` logger, which can be safely used in production and provides advanced options such as adding logs to files, log rotation, etc. In this case, it is configured to display logs on the console.

- Dependency injection:
  I used the `wire` library to manage dependency injection. It is straightforward and very useful.

- HTTP Framework:
  I used the `chi` framework to manage the server easily.

- Configuration file:
  I’ve added a very simple configuration file. Adding more files would allow quick setup for different environments.

---

## Future Improvements

This is a quick and simple version that could be enhanced in many ways. For example, when adding accounts, if one account is invalid, the entire call is discarded, and no accounts are saved. Instead, it could save the valid accounts and return an error message for those that failed to create. This is just one example of many potential improvements.

Some ideas for future enhancements:

- Add tests for different layers of the application.
- Incorporate a database connection or a file reader to load account and transaction data.
- Manage service configuration through a database.
- Add a Dockerfile for building and deploying the application.
- And much more.

