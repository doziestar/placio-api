package schema

//func updateRelevanceScoreHook() ent.Hook {
//	return func(next ent.Mutator) ent.Mutator {
//		return hook.On(next, func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
//			switch tp := m.Type(); tp {
//			case UserFollowUser.Type, UserFollowBusiness.Type, BusinessFollowUser.Type, BusinessFollowBusiness.Type:
//				if !m.Op().Is(ent.Add) {
//					break
//				}
//				followedID, exists := m.ID()
//				if !exists {
//					return nil, errors.New("followed id not provided")
//				}
//
//				// You need to get a handle of your client to make this update operation
//				client := yourClientFromSomewhere
//				if _, err := client.Tp.UpdateOneID(followedID).AddRelevanceScore(1).Save(ctx); err != nil {
//					return nil, err
//				}
//			}
//			return next.Mutate(ctx, m)
//		})
//	}
//}
//
//func updateSearchTextHook() ent.Hook {
//	return func(next ent.Mutator) ent.Mutator {
//		return hook.On(next, func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
//			switch tp := m.Type(); tp {
//			case User.Type, Place.Type, Business.Type:
//				if !m.Op().Is(ent.UpdateOne) {
//					break
//				}
//				name, exists := m.Value("name").(string)
//				if !exists {
//					return nil, errors.New("name not provided")
//				}
//				description, _ := m.Value("description").(string)
//				location, _ := m.Value("location").(string)
//				searchText := name + " " + description + " " + location
//				m.Set("search_text", searchText)
//			}
//			return next.Mutate(ctx, m)
//		})
//	}
//}
