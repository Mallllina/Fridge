package storage

type Freezer struct {
	BaseStorage
}

func NewFreezer(maxVolume float64) Freezer {
	return Freezer{
		BaseStorage: BaseStorage{
			MaxVolume: maxVolume,
		},
	}
}
