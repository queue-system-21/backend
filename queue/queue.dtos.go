package queue

type getDto struct {
	Id      int    `json:"id"`
	NameRus string `json:"nameRus"`
	NameKaz string `json:"nameKaz"`
}

type createDto struct {
	NameRus                 string `json:"nameRus"`
	NameKaz                 string `json:"nameKaz"`
	ResponsibleUserUsername string `json:"responsibleUserUsername"`
}

type infoDto struct {
	Number       int    `json:"number"`
	QueueNameRus string `json:"queueNameRus"`
	QueueNameKaz string `json:"queueNameKaz"`
}
