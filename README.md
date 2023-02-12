# Тестовое задание WB
Сервер на Golang для сохранения и выдачи JSON файлов. 
Проект состоит из нескольких компонентов: 
- Сервера
- Базы данных (файлы работы сервера с базой Postgres находятся в pkg/repository)
- nasts-subscriber (pkg/mynats)
- nasts-publisher (cmd/publisher)
- самого nats-streaming-server, который запускался отдельно (проект по ссылке в дополнительных материалах)
## Дополнительные материалы
Демонстрация работы: https://youtu.be/Z4PWYY9EVRA
База данных Postgres версии 13.3  
Nats-streaming: https://github.com/nats-io/nats-streaming-server.git  
