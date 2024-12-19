# Веб-сервис, для вычисления арифметических выражений
Данный проект вычисляет арифметические выражения. Пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.
## Установка
1. Установите язык программирования [Golang](https://go.dev/dl/).
2. Установите текстовый редактор [Visual Studio Code](https://code.visualstudio.com/).
3. Установите систему контроля версий [Git](https://git-scm.com/downloads).
4. Создайте папку и откройте ее в Visual Studio Code.
5. Создайте клон репозитория с GitHub. Для этого в терминале Visual Studio Code введите следующую команду:
```
git clone https://github.com/kingofhandsomes/calculation_go
```
## Использование
1. В терминале Visual Studio Code перейдите в папку calculation_go с помощью команды:
```
cd calculation_go
```
2. Введите команду ниже для установки пакета zap:
```
go get go.uber.org/zap
```
3. Запустите веб-сервис, введя следующую команду:
```
go run ./cmd/main.go
```
4. Откройте git bash и введите команду нижe. __После двоеточия после слова expression в двойных кавычках напишите свое выражение (поддерживаемые символы: (, ), +, -, *, /):__
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":""}'
```
Вам будет выводиться либо решение примера, либо ошибка.
Также при отправке запроса в Git Bash в терминале Visual Studio Code будут отображаться логи для отладки и дополнительной информации о запросах и длительности их обработки, об ответах.
## Пример успешного запроса в Git Bash
Пример успешного запроса:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"5+5*5"}'
```
Результат запроса:
```
{"result":30}
```
## Пример провальных запросов в Git Bash
Данный веб-сервис содержит обработку 7 ошибок, для их вывода можно написать следующие запросы:
1. Возвращает ошибку - ErrDivisionByZero - деление на 0 запрещено:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"2+3/0"}'
```
Результат запроса:
```
{"error":"division by zero"}
```
2. Возвращает ошибку - ErrNumberSearch - после операции умножения не стоит число:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"3-7**9"}'
```
Результат запроса:
```
{"error":"symbol was encountered instead of a number"}
```
3. Возвращает ошибку - ErrAmountBrackets - разное количество открывающих и закрывающих скобок:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"(1+3*8+(3*0)-6"}'
```
Результат запроса:
```
{"error":"different number of opening and closing brackets"}
```
4. Возвращает ошибку - ErrInvalidCharacter - в выражении присутствует недопустимый символ:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"a1+4"}'
```
Результат запроса:
```
{"error":"invalid character was encountered"}
```
5. Возвращает ошибку - ErrUnexpectedEnd - неожиданный конец выражения:
```
curl -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"100*"}'
```
Результат запроса:
```
{"error":"unexpected end of the expression"}
```
6. Возвращает ошибку - ErrInvalidMethod - неверный метод запроса:
```
curl -X GET -L 'localhost:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression":"1+2"}'
```
Результат запроса:
```
{"error":"request method is specified incorrectly"}
```
7. Возвращает ошибку - ErrInternalServer - ошибка со стороны сервера. Для данной ошибки нет примера запроса, потому что вышеперечисленные ошибки итак обрабатывают все варианты поломки кода.
Результат запроса, который должен вызывать ошибку ErrInternalServer:
```
{"error":"internal server error"}
```
## Запуск тестов в Visual Studio Code
В данном проекте есть 2 файла с тестами:
1. calculation_test.go - тесты, с помощью которых можно проверить правильность вычисления и написания выражения. Для запуска в терминале Visual Studio Code введите команду:
```
go test ./package/calculation/calculation_test.go
```
Варианты вывода:
- ok - тесты выполнились успешно;
- fail - некоторые тесты вернули ошибку.
2. application_test.go - тесты, с помощью которых можно проверить правильность вычисления, написания выражения и HTTP-запроса, HTTP-ответа. Для запуска в терминале Visual Studio Code введите команду:
```
go test ./internal/application/application_test.go
```
Варианты вывода:
- ok - тесты выполнились успешно;
- fail - некоторые тесты вернули ошибку.
## Обратная связь
Телеграмм: @KinGofHanDSomEs
