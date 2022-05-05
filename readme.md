### Curl Get
```
curl -X GET \
  'localhost:9997/bitcoin' \
  --header 'Accept: */*' \
  --header 'User-Agent: chayanin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "startDateTime":"2022-05-01T00:00:00Z",
    "endDateTime":"2022-05-01T00:00:00Z"
}'
```
### Curl Post
```
curl -X POST \
  'localhost:9997/bitcoin' \
  --header 'Accept: */*' \
  --header 'User-Agent: chayanin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "amount":101.1
}'
```

### Build Application
``` 
$ make build-app
```

### Run Application
```
$ make up
```

### Stop Application
```
$ <ctrl>C
$ make down
```