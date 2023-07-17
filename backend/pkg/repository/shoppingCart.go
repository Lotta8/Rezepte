package repository

import (
	"errors"
	"recipies/pkg/model"
	"sync"
)

type ShoppingCartInMemoryStorage struct {
	nextId           int
	shoppingCart     *model.ShoppingCart
	recipeRepository *RecipeInMemoryStorage
}

func NewShoppingCartInMemoryStorage() *ShoppingCartInMemoryStorage {
	return &ShoppingCartInMemoryStorage{
		nextId:           0,
		shoppingCart:     &model.ShoppingCart{Items: make([]*model.ShoppingCartItem, 0)},
		recipeRepository: NewRecipeInMemoryStorage(),
	}
}

func (s *ShoppingCartInMemoryStorage) Add(id int, count int) (bool, error) {
	recipe, err := s.recipeRepository.GetRecipe(id)
	if err != nil {
		return false, err
	}

	s.nextId++
	s.shoppingCart.Items = append(s.shoppingCart.Items, &model.ShoppingCartItem{
		Id:     s.nextId,
		Count:  count,
		Recipe: recipe,
	})

	return true, nil
}

func (s *ShoppingCartInMemoryStorage) Get() []*model.ShoppingCartItem {
	return s.shoppingCart.Items
}

// ----------------------------

type ShoppingCart struct {
	mutex        sync.RWMutex
	cartByUserID map[int][]string
}

func (s *ShoppingCart) DeleteEinkaufskorb(userID int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.cartByUserID[userID]
	if !ok {
		return errors.New("Einkaufskorb nicht gefunden")
	}

	delete(s.cartByUserID, userID)

	return nil
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		cartByUserID: make(map[int][]string),
	}
}

func (s *ShoppingCart) AddToCart(userID int, item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cartByUserID[userID] = append(s.cartByUserID[userID], item)
}

func (s *ShoppingCart) GetCart(userID int) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cart, ok := s.cartByUserID[userID]
	if !ok {
		return nil, errors.New("Shopping cart not found")
	}

	return cart, nil
}

func (s *ShoppingCart) RemoveFromCart(userID int, item string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, ok := s.cartByUserID[userID]
	if !ok {
		return errors.New("Shopping cart not found")
	}

	for i, v := range cart {
		if v == item {
			// Remove the item from the cart
			s.cartByUserID[userID] = append(cart[:i], cart[i+1:]...)
			return nil
		}
	}

	return errors.New("Item not found in the shopping cart")
}

func (s *ShoppingCart) GetEinkaufskorb(id int) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cart, ok := s.cartByUserID[id]
	if !ok {
		return nil, errors.New("Einkaufskorb nicht gefunden")
	}

	return cart, nil
}
