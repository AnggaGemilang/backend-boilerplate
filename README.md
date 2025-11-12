### Reference
https://github.com/golang-standards/project-layout/tree/master

### Boilerplate structure

```
├── cmd/
├── config/
│   ├── apisix/
├── docs/
├── internal/
│   ├── adapter/
│   │   ├── adapterinterface/
│   │   │   ├── common/
│   │   │   └── persistence/
│   │   ├── messaging/
│   │   │   └── nats
│   │   ├── persistence/
│   │   │   └── repository/
│   ├── app/
│   ├── config/
│   ├── controller/
│   │   ├── api/
│   │   │   ├── grpc/
│   │   │   ├── middleware/
│   │   │   └── rest/
│   │   └── handler/
│   ├── domain/
│   │   ├── dto/
│   │   │   ├── request/
│   │   │   └── response/
│   │   ├── model/
│   │   ├── presenter/
│   ├── service/
│   ├── shared/
│   │   ├── constant/
│   │   ├── environment/
│   │   ├── event/
│   │   ├── util/
│   │   │   ├── helper/
│   └   └   └── logger/
├── proto/
└── test/
```


