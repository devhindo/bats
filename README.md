# bats


## Endpoints

| Endpoint | Description | Functionality |
|----------|-------------|---------------|
| /api/user/new | create new user | when making `POST` req a user registeration proccess should be made with an OTP validation that has a 1hr expiration date. |

## Development

- this project follows the community driven [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Deployment

`make up`
-- fire a docker-compose up command to start the server. It contains a mysql container and a go server container --