package application

import (
	"replica/src/vivero/domain/entities"
	"replica/src/vivero/domain/repository"
)

type LongPollViverosUseCase struct {
	viveroRepo *repository.ViveroRepository // ✅ Se usa un puntero
}

func NewLongPollViverosUseCase(repo *repository.ViveroRepository) *LongPollViverosUseCase { // ✅ Se recibe un puntero
	return &LongPollViverosUseCase{viveroRepo: repo}
}

func (u *LongPollViverosUseCase) LongPoll(lastData []entities.Vivero) ([]entities.Vivero, error) {
	viveros, err := u.viveroRepo.GetAll() // ✅ Se accede correctamente al método del repositorio
	if err != nil {
		return nil, err
	}
	if len(viveros) != len(lastData) {
		return viveros, nil
	}
	return lastData, nil
}
