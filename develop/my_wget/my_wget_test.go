package main

import "testing"

func TestDownloadSite(t *testing.T) {
	existingURL := "https://microsoft.com"
	err := downloadSite(existingURL)
	if err != nil {
		t.Errorf("Ожидалось успешное выполнение, получено: %v", err)
	}

	nonExistentURL := "https://www.thissitedoesnotexist.com"
	err = downloadSite(nonExistentURL)
	if err == nil {
		t.Error("Ожидалась ошибка при попытке загрузить несуществующий сайт")
	}
}
