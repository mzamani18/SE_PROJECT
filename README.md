# TikOn


TikOn is a ticket reservation system is an golang application designed to provide customers with a personalized easy-to-utilize user experience for booking and purchasing tickets online. It stores customers' personal data records, scheduled routes, frequent trips, and other information. In TikOn every user has dedicated wallet.

### Postman Api Documentation:
* [PostmanDoc](https://documenter.getpostman.com/view/14995830/2s946fdCDM#52da8536-364f-40c8-adc1-483b326dd652)

### Config SetUp:
you must create .env file in TIKON directory and you can fill it like Config Sample.

### Config Sample:
```
PORT = 3000
DATABASE_URL = "root:pass@tcp(127.0.0.1:3306)/TikOn?parseTime=true"
SECRET_AUTH_KEY = "edwlefgkjrt"
```

### Run:
Use ``go run main.go`` to run in your local.

### Generate ORM And Migrate:
Use ``go run ./Migrate/Migrate.go`` to regenerate the entity package based on the current db schema.


