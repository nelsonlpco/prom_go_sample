package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/nelsonlpco/classic_cc_problens/internal/api/models"
	"github.com/nelsonlpco/classic_cc_problens/internal/domain/gene"
	"github.com/nelsonlpco/classic_cc_problens/internal/shared"
)

const DNAPath = "/dna"

func (h Handler) CompressDNA(c echo.Context) error {
	span := shared.Tracer.Start(c.Request().Context(), "compress-dna")
	defer span.End()

	dna := new(models.DNAModel)
	if err := c.Bind(dna); err != nil {
		return err
	}

	_, err := gene.New(dna.DNA)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &models.DefaultResponseModel{Code: http.StatusOK, Message: "rna compactado com sucesso"})
}
