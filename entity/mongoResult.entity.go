package entity

type MongoResult struct {
	Success    bool
	Message    string
	InsertedID string
	Count      int64
	Data       interface{}
}
