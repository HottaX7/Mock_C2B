package main

import (
	"io"
	"os"

	"espp-mock/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// 1. Создаем (или открываем) файл для логов
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось открыть файл логов")
	}
	defer logFile.Close()

	// 2. Настраиваем zerolog: выводим в консоль и в файл
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	multi := io.MultiWriter(consoleWriter, logFile)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	// 3. Уровень логирования
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// 4. Запуск приложения
	log.Info().Msg("Сервис запущен. Логирование начато.")
	cmd.Execute()
}
