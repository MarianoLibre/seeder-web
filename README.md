# Seeder (Web)

Web tool for database seeding (melisprint), with a twist :blush:

## Instructions

Clone (ssh): 
```
git clone git@github.com:MarianoLibre/seeder-web.git
```

### Build the front:

```
$ cd web
$ npm i
...
$ npm run build
```

### Configure the seeder

You'll need a `.env` file on you `api` folder, with the following content:

```
USERNAME=
PASS=
DATABASE=
PORT=8080

GIN_MODE='debug'
```

This will fill the dataSource `user:pass@/database?parseTime=true` string so, use the name and pass you normally use to connect to mysql.
If you wanna test on a test_db first, just create a db on mysql (i.e. `test_db`) and use `DATABASE='test_db'`.

### Run the seeder

```
$ cd ../api
$ go run .
```

Check if it's working: `localhost:8080`

You'll see the routes list on the terminal. Use Insomnia or Postman to do the requests or, you can use my super cool front instead 😃 at `localhost:8080/foo`

If I missed something let me know 😉

Have fun 🤩 !
