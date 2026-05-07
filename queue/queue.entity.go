package queue

type queue struct {
	Id, NextFreeSlotNumber, CurrentSlotNumber int
	NameRus, NameKaz, ResponsibleUserUsername string
}
