Необходимо сделать gRPC обёртку над сайтом https://www.rusprofile.ru/

## API

Сервис должен реализовывать один метод, принимающий на вход ИНН компании, ищущий компанию на rusprofile, и возвращающий её ИНН, КПП, название, ФИО руководителя

## Технологии

* сервис должен быть написан на Go с использованием Go Modules
* предоставлять API через [gRPC](https://grpc.io/docs/languages/go/quickstart/)
* предоставлять API через HTTP с помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* предоставлять Swagger UI с документацией, сгенерированной из .proto файла с помощью protoc-gen-swagger
* быть упакован в docker контейнер

## Запуск
```shell
git clone https://github.com/syalsr/Rusprofile
make run
```
1. grpc server - :50051 
2. http server - localhost:8081, для получения данные о компании по ИНН, достаточно перейти на данный адрес - http://localhost:8081/INN/2112001817
3. swagger - localhost:8080, для доступа к документации достаточно перейти на данную страницу - http://localhost:8080/swagger/index.html