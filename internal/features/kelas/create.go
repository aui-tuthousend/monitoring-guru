package kelas

import (
	e "monitoring-guru/entities"
	"strconv"

	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateKelasRequestBody
// @Description Create kelas request body
type CreateKelasRequest struct {
	// @Description Grade of the kelas
	// @Required true
	// @Example "XII"
	Grade string `json:"grade"`
	// @Description Index of the kelas
	// @Required true
	// @Example "1"
	Index int `json:"index"`
	// @Description Jurusan ID of the kelas
	// @Required true
	// @Example "1"
	JurusanID string `json:"jurusan_id"`
	// @Description Ketua Kelas ID of the kelas
	// @Required true
	// @Example "123456789"
	KetuaKelasID string `json:"ketua_kelas_id"`

}

// CreateKelasRequest godoc
// @summary Create Kelas request body
// @Description	Create Kelas baru request body
// @Tags			Kelas
// @Security    BearerAuth
// @Accept			json
// @Produce		json
// @Param			request	body		CreateKelasRequest	true	"Create kelas request body"
// @Success		200		{object}	KelasResponseWrapper
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/api/kelas [post]
func (h *KelasHandler) CreateKelas() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKelasRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if req.Grade == "" || req.Index == 0 || req.JurusanID == "" || req.KetuaKelasID == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid request body", nil))
		}

		if req.Grade != "X" && req.Grade != "XI" && req.Grade != "XII" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid grade", nil))
		}

		if req.Index < 1 {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid index", nil))
		}

		ketuaID, err := utils.ParseUUID(req.KetuaKelasID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jurusanID, err := utils.ParseUUID(req.JurusanID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jurusan, err := h.JurusanService.GetJurusanByID(jurusanID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Jurusan not found", nil))
		}

		ketuaKelas, err := h.KetuaKelasService.GetKetuaKelas(ketuaID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Ketua Kelas not found", nil))
		}

		name := jurusan.KodeJurusan + " " + req.Grade + " " + strconv.Itoa(req.Index)

		kelas := e.Kelas{
			ID:        uuid.New(),
			KetuaID:   ketuaID,
			JurusanID: jurusanID,
			Name:      name,
			Grade:      req.Grade,
			Index:      req.Index,
			IsActive:  false,
		}

		statusKelas := e.StatusKelas{
			ID:        uuid.New(),
			KelasID: kelas.ID,
		}

		if err := h.Service.CreateKelas(&kelas, &statusKelas); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		ketuaKelas.KelasID = kelas.ID
		if err := h.KetuaKelasService.UpdateKetuaKelas(ketuaKelas); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		res := KelasResponse{
			ID:   kelas.ID.String(),
			Name: kelas.Name,
			Jurusan: h.JurusanService.ResponseJurusanMapper(jurusan),
			KetuaKelas: h.KetuaKelasService.ResponseKetuaKelasMapper(ketuaKelas),
			IsActive: kelas.IsActive,
		}

		return c.JSON(e.SuccessResponse(&res))
	}
}
