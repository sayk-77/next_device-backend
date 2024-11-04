package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(userID uint, items []models.OrderItem, totalPrice float64) (*models.Order, error) {
	// Создаем заказ
	order := &models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     "pending",
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	for i := range items {
		items[i].OrderID = order.ID
	}

	if err := s.orderRepo.CreateOrderItems(items); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) DeleteOrder(orderId uint) error {
	err := s.orderRepo.DeleteOrder(orderId)
	if err != nil {
		return err
	}
	return nil
}
