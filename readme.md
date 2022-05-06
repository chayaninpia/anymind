### Curl Get
```
curl -X GET \
  'localhost:9997/bitcoin' \
  --header 'Accept: */*' \
  --header 'User-Agent: chayanin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "startDateTime": "2022-05-04T14:15:39+07:00",
  "endDateTime": "2022-06-07T00:00:00+07:00"
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
    "date_time": null,
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
