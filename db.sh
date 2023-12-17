docker run \
    --name some-postgres \
    -e POSTGRES_PASSWORD=mustakrakish \
    -e POSTGRES_USER=admin \
    -e POSTGRES_DB=sm-system -p 5432:5432  -d postgres
