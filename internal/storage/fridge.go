package storage

type Fridge struct {
	BaseStorage
}

func NewFridge(maxVolume float64) Fridge {
	return Fridge{
		BaseStorage: BaseStorage{
			MaxVolume: maxVolume,
		},
	}
}
