# Используем официальный образ PostgreSQL
FROM postgres:latest

# Установка переменных окружения для настройки базы данных
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=song_libraries


# Открытие порта
EXPOSE 5432