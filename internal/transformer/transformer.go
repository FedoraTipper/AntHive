package transformer

import (
	"errors"
	"fmt"
	"github.com/FedoraTipper/AntHive/internal/constants"
	"github.com/FedoraTipper/AntHive/internal/models"
)

type Transformer interface {
	ConvertStatsPayloadToMiner(string, string, interface{}) (*models.Miner, error)
}

func GetTransformer(model string) (Transformer, error) {
	switch model {
	case constants.X19_MinerSeries:
		return &S19Transformer{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown model: %s", model))
	}
}
