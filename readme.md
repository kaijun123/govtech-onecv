## Govtech-OneCV

### How to run the code
- To run this code, you need to have docker installed and also make.
- Run in the home directory `make start`. If you do not have make, you can also run `docker compose up --quiet-pull --remove-orphans`

### How to run the test
- Run in the home directory `make start`
- Open another terminal. Run `go test -v` in the `test` directory
- While this is definitely not the right way to write the tests, there was insufficient time to debug the code properly to resolve some of the issues that I was facing.
- Unit tests cover most of the basic test cases

### Notes:
- Postgres container will be runing on port 5431 on the local machine
- Server will be running on port 8080
- Any api calls to the server should be directed to `http://localhost:8080`