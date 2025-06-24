package importdata

import (
	e "monitoring-guru/entities"
	"time"

	"monitoring-guru/internal/features/jurusan"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func UploadJurusanHandler(c *fiber.Ctx) error {

	jurusanServ := jurusan.JurusanServ

	fileHeader, err := c.FormFile("files")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "File tidak ditemukan"})
	}

	// Buka file dari form
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuka file"})
	}
	defer file.Close()

	// Baca dengan excelize langsung dari reader
	excelFile, err := excelize.OpenReader(file)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membaca file Excel"})
	}

	cols, err := excelFile.GetCols("_kompetensi")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membaca kolom dari sheet"})
	}

	columnA := cols[0]
	columnB := cols[1]

	// Gabungkan data dari dua kolom menjadi slice 2D
	var errorData [][]string
	maxLen := len(columnA)
	if len(columnB) > maxLen {
		maxLen = len(columnB)
	}

	for i := 1; i < maxLen; i++ {

		var kodeJurusan, name string

		if i < len(columnA) {
			kodeJurusan = columnA[i]
		}
	
		if i < len(columnB) {
			name = columnB[i]
		}
		
		if kodeJurusan == "" || name == "" {
			errorData = append(errorData, []string{kodeJurusan, name})
			continue
		}

		if jurusanServ.IsJurusanExist(kodeJurusan, name) {
			errorData = append(errorData, []string{kodeJurusan, name})
			continue
		}

		err := jurusanServ.CreateJurusan(&e.Jurusan{
				Name: name,
				KodeJurusan: kodeJurusan,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

		if err != nil {
			errorData = append(errorData, []string{columnA[i], columnB[i]})
		}

	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengimport file",
		"unsaved data": errorData,
	})
}
