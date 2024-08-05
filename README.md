# nas-torrent-bot
## Бот для NAS. Для скачивания торрентов

### Подготовка на NAS
Я использую Synology и DownloadStation.

В настройках DownloadStation указать следующее:

Settings->Location->Destination - путь до папки, куда сохраняется скачанный торрент (WATCH_DIR)

Settings->Location->WatchedFolder - путь до папки, откуда автоматом будут ставиться торренты на закачку (DOWNLOAD_DIR)

### Сборка образа контейнера
Скачайте проект и выполните команду
```
make build
```
В корне проекта появится torrentbot.tar

### Запуск образа на NAS
В Synology есть Container Manager. В него надо загрузить torrentbot.tar
Далее необходимо загрузить и настроить образ.

Необходимо добавить следующие переменные окружения:

WATCH_DIR - папка, куда сохраняются скаченные торренты

DOWNLOAD_DIR - папка, которую отслеживает DownloadStation. Сюда бот будет загружать торренты

BOT_TOKEN - токен бота

SECRET - секретная фраза для команды /start. Сесюрити превыше всего

Также при настройке контейнера - не забудьте прокинуть папки.


### Что еще в планах
- Сейчас бот не показывает какие подпакпки есть. То есть при перемещении файла, надо помнить, как вы назвали папки на NAS
- В качестве хранилища для юзеров используется мапа. Любое падение бота - надо заново писать команду /start

