package entity

const (
	// Ответы в чат
	RulesAnswer      = "Отправь в чат файл с расширением .torrent\n Дождись оповещения о том, что он скачан и следуй дальнейшим инструкциям"
	StartRulesAnswer = "Для того, чтобы начать общение отправь команду /start. А через пробел напиши секретную фразу.\n" +
		"Пример: /start umpalumpa"
	SuccessMovedFile           = "Файл %s успешно перемещен. Можно смотреть!"
	FileSuccessfullyDownloaded = "Файл успешно добавлен на закачку. Я скажу тебе, как он будет скачан"

	// Сообщения ошибки
	InvalidCommand     = "Я не знаю команду %s"
	InvalidSecret      = "Пу пу пу... Неверная секретная фраза, дружок"
	FailedMoveFile     = "Я не смог переместить файл. Вот ошибка %v"
	FailedDownloadFile = "Не смог загрузить файл. Вот ошибка %v"

	// Перечень комманд
	CommandStart = "start"
	CommandMove  = "move"
)
