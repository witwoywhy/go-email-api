# go-email-api

## Apis
### Notification
```
curl --location 'http://localhost:8080/v1/notification' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "userId": "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec",
        "email": "mail@gmail.com",
        "event": "E10001",
        "details": {
            "accountName":"Tester Test",
            "accountNo": "012345678"
        },
        "files": ["cat.jpg", "cat.pdf"]
}'
```

## Design Pattern
- reference from Arnon Keereena