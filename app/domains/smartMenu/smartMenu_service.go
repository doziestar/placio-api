package smartMenu

import "placio-app/ent"

type SmartMenuService struct {
	client *ent.Client
}

func NewSmartMenuService(client *ent.Client) *SmartMenuService {
	return &SmartMenuService{client: client}
}
