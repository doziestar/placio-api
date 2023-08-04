package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"placio-app/Dto"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/faq"
	"placio-app/utility"
)

type FAQService interface {
	CreateFAQ(ctx context.Context, businessID string, faqData *Dto.FAQDto) (*ent.FAQ, error)
	GetFAQ(ctx context.Context, faqID string) (*ent.FAQ, error)
	UpdateFAQ(ctx context.Context, faqID string, faqData map[string]interface{}) (*ent.FAQ, error)
	DeleteFAQ(ctx context.Context, faqID string) error
	AssociateFAQWithPlace(ctx context.Context, faqID, placeID string) error
	AssociateFAQWithEvent(ctx context.Context, faqID, eventID string) error
	GetFAQsByBusiness(ctx context.Context, businessID string) ([]*ent.FAQ, error)
}

type FAQServiceImpl struct {
	client *ent.Client
	cache  *utility.RedisClient
}

func NewFAQService(client *ent.Client, cache *utility.RedisClient) *FAQServiceImpl {
	return &FAQServiceImpl{client: client, cache: cache}
}

func (s *FAQServiceImpl) CreateFAQ(ctx context.Context, businessID string, faqData *Dto.FAQDto) (*ent.FAQ, error) {
	faq, err := s.client.FAQ.
		Create().
		SetID(uuid.New().String()).
		SetQuestion(faqData.Question).
		SetAnswer(faqData.Answer).
		SetBusinessID(businessID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Cache the result.
	go func() {
		data, err := json.Marshal(faq)
		if err == nil {
			s.cache.SetCache(ctx, fmt.Sprintf("faq:%s", faq.ID), data)
		}
	}()

	return faq, nil
}

func (s *FAQServiceImpl) GetFAQ(ctx context.Context, faqID string) (*ent.FAQ, error) {
	// Check the cache first.
	cacheKey := fmt.Sprintf("faq:%s", faqID)
	data, err := s.cache.GetCache(ctx, cacheKey)
	if err == nil {
		var faq ent.FAQ
		err = json.Unmarshal(data, &faq)
		if err == nil {
			return &faq, nil
		}
	}

	faq, err := s.client.FAQ.
		Get(ctx, faqID)
	if err != nil {
		return nil, err
	}

	// Cache the result.
	go func() {
		data, err := json.Marshal(faq)
		if err == nil {
			s.cache.SetCache(ctx, cacheKey, data)
		}
	}()

	return faq, nil
}

func (s *FAQServiceImpl) GetFAQsByBusiness(ctx context.Context, businessID string) ([]*ent.FAQ, error) {
	faqs, err := s.client.FAQ.
		Query().
		Where(faq.HasBusinessWith(business.IDEQ(businessID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return faqs, nil
}

func (s *FAQServiceImpl) UpdateFAQ(ctx context.Context, faqID string, faqData map[string]interface{}) (*ent.FAQ, error) {
	updater := s.client.FAQ.UpdateOneID(faqID)

	// Update fields dynamically based on faqData.
	for key, value := range faqData {
		switch key {
		case "question":
			updater.SetQuestion(value.(string))
		case "answer":
			updater.SetAnswer(value.(string))
			// Add more fields if needed.
		}
	}

	faq, err := updater.Save(ctx)
	if err != nil {
		return nil, err
	}

	// Cache the result.
	go func() {
		data, err := json.Marshal(faq)
		if err == nil {
			s.cache.SetCache(ctx, fmt.Sprintf("faq:%s", faq.ID), data)
		}
	}()

	return faq, nil
}

func (s *FAQServiceImpl) DeleteFAQ(ctx context.Context, faqID string) error {
	err := s.client.FAQ.
		DeleteOneID(faqID).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Remove from cache.
	go func() {
		s.cache.DeleteCache(ctx, fmt.Sprintf("faq:%s", faqID))
	}()

	return nil
}

func (s *FAQServiceImpl) AssociateFAQWithPlace(ctx context.Context, faqID, placeID string) error {
	// get faq
	faq, err := s.client.FAQ.
		Get(ctx, faqID)
	if err != nil {
		return err
	}

	err = s.client.Place.
		UpdateOneID(placeID).
		AddFaqs(faq).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Cache the result, if needed.
	// This is a complex operation, we should invalidate the cache of the associated objects (Place and FAQ).
	go func() {
		s.cache.DeleteCache(ctx, fmt.Sprintf("faq:%s", faqID))
		s.cache.DeleteCache(ctx, fmt.Sprintf("place:%s", placeID))
	}()

	return nil
}

func (s *FAQServiceImpl) AssociateFAQWithEvent(ctx context.Context, faqID, eventID string) error {
	// get faq
	faq, err := s.client.FAQ.
		Get(ctx, faqID)
	if err != nil {
		return err
	}

	err = s.client.Event.
		UpdateOneID(eventID).
		AddFaqs(faq).
		Exec(ctx)
	if err != nil {
		return err
	}

	go func() {
		s.cache.DeleteCache(ctx, fmt.Sprintf("faq:%s", faqID))
		s.cache.DeleteCache(ctx, fmt.Sprintf("event:%s", eventID))
	}()

	return nil
}
