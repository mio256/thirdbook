# thirdbook

![mio256](https://avatars.githubusercontent.com/u/71450182)

## Overview

simple booking system

## Requirement

[go.mod](./go.mod)

## Usage

[Makefile](./Makefile)

```sh
docker compose -f docker-compose.dev.yaml up -d
make tools
make migrate-db
go get .
cp .env.sample .env
echo dotenv > .envrc
make local-server
```

## Deploy

```sh
gcloud builds submit --tag "tag_name"
```

## Author

[mio256](https://github.com/mio256)

## License

[MIT License](./LICENSE)
