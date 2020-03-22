# Start postgres

```
make build
make start
```
# Run migrations

```
make migrate
```

# create pgadmin server
```
docker run -p 5434:80 -e 'PGADMIN_DEFAULT_EMAIL=pg@admin.com' -e 'PGADMIN_DEFAULT_PASSWORD=password' --name pgadmin --network=host -d dpage/pgadmin4
```
- user: pg@admin.com
- password: password