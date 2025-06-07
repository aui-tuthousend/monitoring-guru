package kelas

import (
	e "monitoring-guru/entities"
	
	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateKelasRequestBody
// @Description Create kelas request body
type CreateKelasRequest struct {
	// @Description Name of the kelas
	// @Required true
	// @Example "XII RPL 1"
	Name string `json:"name"`
	// @Description Jurusan ID of the kelas
	// @Required true
	// @Example "123456789"
	JurusanID string `json:"jurusan_id"`
	// @Description Ketua Kelas ID of the kelas
	// @Required true
	// @Example "123456789"
	KetuaKelasID string `json:"ketua_kelas_id"`
	// @Description Wali Kelas ID of the kelas
	// @Required true
	// @Example "123456789"
	// WaliKelasID string `json:"wali_kelas_id"`
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

		ketuaID, err := utils.ParseUUID(req.KetuaKelasID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		// wakilID, err := utils.ParseUUID(req.WaliKelasID)
		// if err != nil {
		// 	return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		// }
		jurusanID, err := utils.ParseUUID(req.JurusanID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jurusan, err := h.JurusanService.GetJurusanByID(jurusanID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Jurusan not found", nil))
		}

		kelas := e.Kelas{
			ID:        uuid.New(),
			KetuaID:   ketuaID,
			// WakilID:   wakilID,
			JurusanID: jurusanID,
			Name:      req.Name,
			IsActive:  false,
		}

		if err := h.Service.CreateKelas(&kelas); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		ketua, err := h.KetuaKelasService.GetKetuaKelasByID(ketuaID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Ketua Kelas not found", nil))
		}

		res := KelasResponse{
			ID:   kelas.ID.String(),
			Name: kelas.Name,
			Jurusan: h.JurusanService.ResponseJurusanMapper(jurusan),
			KetuaKelas: h.KetuaKelasService.ResponseKetuaKelasMapper(ketua),
			IsActive: kelas.IsActive,
		}

		return c.JSON(e.SuccessResponse(&res))
	}
}
