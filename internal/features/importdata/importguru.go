package importdata

import (
	e "monitoring-guru/entities"
	"monitoring-guru/utils"
	"time"

	"monitoring-guru/internal/features/guru"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func UploadGuruHandler(c *fiber.Ctx) error {

	guruService := guru.GuruServ

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

	cols, err := excelFile.GetCols("guru")
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
		var nip, name string

		if i < len(columnA) {
			nip = columnA[i]
		}
	
		if i < len(columnB) {
			name = columnB[i]
		}
		
		if nip == "" || name == "" {
			errorData = append(errorData, []string{nip, name})
			continue
		}

		if guruService.IsGuruExist(nip, name) {
			errorData = append(errorData, []string{nip, name})
			continue
		}

		hashedPassword, _ := utils.HashPassword("default")
		err := guruService.CreateGuru(&e.Guru{
			Nip: nip,
			Name: name,
			Jabatan: "guru",
			Password: hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if err != nil {
			errorData = append(errorData, []string{columnA[i], columnB[i]})
		}

	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengimport file",
		"unsaved data (existed)": errorData,
	})
}
