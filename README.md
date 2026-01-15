# Go_Learn
Репозитонрий для изучения Golang

Сайт с которого берётся материал для изучения: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals

Локальный просмотр документации по пакету: go doc NAME_OF_PACKAGE
Пример: go doc fmt
АЛЬТЕРНАТИВНО:
go install golang.org/x/pkgsite/cmd/pkgsite@latest
и пишем: pkgsite -open
вся документация теперь доступна по: http://localhost:8080/testing

# Naming for test:
(https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#writing-tests)

Написание теста похоже на написание функции, но с некоторыми правилами

-  Он должен находиться в файле с таким названием, как xxx_test.go

- Тестовая функция должна начинаться со слова Test

- Функция тестирования принимает только один аргумент t *testing.T

- Чтобы использовать тип *testing.T нужно выполнить в начале файла с тестами import "testing"