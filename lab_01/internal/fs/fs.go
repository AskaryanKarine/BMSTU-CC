package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateGraphsDir() error {
	// Используем MkdirAll для рекурсивного создания директорий
	err := os.MkdirAll("graphs/emulate", 0755)
	if err != nil {
		return fmt.Errorf("ошибка создания директорий: %v", err)
	}
	return nil
}

func DeleteGraphsDir() error {
	// Используем RemoveAll для удаления директории и всех вложенных элементов
	err := os.RemoveAll("graphs")
	if err != nil {
		return fmt.Errorf("ошибка удаления директории: %v", err)
	}
	return nil
}

func RecreateEmulateDir() error {
	// Формируем полный путь к папке
	emulatePath := filepath.Join("graphs", "emulate")

	// Удаляем папку со всем содержимым
	if err := os.RemoveAll(emulatePath); err != nil {
		return fmt.Errorf("не удалось удалить папку: %v", err)
	}

	// Создаём новую пустую папку
	if err := os.MkdirAll(emulatePath, 0755); err != nil {
		return fmt.Errorf("не удалось создать папку: %v", err)
	}

	return nil
}
