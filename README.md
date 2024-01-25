# Beaver admin

Beaver admin.

### Install project:

```bash
> nvm install 20
> nvm use 20
> cp .env.sample .env
> cp next-client/.env.sample next-client/.env
```

Add the missing info to the `.env` files.

```bash
> docker compose build
> docker compose --profile tools run migrate
> docker compose run api go run main.go migrate
```

### Start backend:

```bash
> docker compose up
```

### Start frontend

```bash
> cd next-client
> npm run dev
```
