# Команды для запуска проекта

## Требования:
- Docker и Docker Compose
- Go 1.23+ (локально опционально)

## Запуск:

### 1. Остановить все контейнеры (если запущены):
```bash
docker-compose down
```

### 2. Собрать и запустить проект:
```bash
docker-compose up --build
```

### 3. Посмотреть логи:
```bash
docker-compose logs -f app
```

### 4. Остановить проект:
```bash
docker-compose down
```

### 5. Очистить всё (контейнеры, volumes):
```bash
docker-compose down -v
```

## Что делает проект:

1. Создает MinIO бакет `mybucket`
2. Загружает 3 файла:
   - `files/text/sample.txt` (TXT)
   - `files/json/sample.json` (JSON)  
   - `files/images/sample.png` (PNG)

## Доступ к MinIO:

- **API**: http://localhost:9000
- **Console**: http://localhost:9001
- **Login**: admin / admin12345

## Проверка результата:

После запуска проверьте файлы в MinIO Console:
1. Откройте http://localhost:9001
2. Войдите (admin/admin12345)
3. Выберите бакет `mybucket`
4. Увидите 3 загруженных файла
