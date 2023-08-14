package repository

import (
	"errors"
	"recipies/pkg/model"
)

type ShoppingCartInMemoryStorage struct {
	nextId           int
	shoppingCart     map[int]*model.ShoppingCartItem
	recipeRepository *RecipeInMemoryStorage
}

func NewShoppingCartInMemoryStorage() *ShoppingCartInMemoryStorage {
	return &ShoppingCartInMemoryStorage{
		nextId:           0,
		shoppingCart:     make(map[int]*model.ShoppingCartItem),
		recipeRepository: NewRecipeInMemoryStorage(),
	}
}

func (s *ShoppingCartInMemoryStorage) DeleteAll() (bool, error) {
	s.shoppingCart = make(map[int]*model.ShoppingCartItem)
	return true, nil
}

func (s *ShoppingCartInMemoryStorage) DeleteById(id int) (bool, error) {
	if _, exists := s.shoppingCart[id]; exists {
		delete(s.shoppingCart, id)
		return true, nil
	}
	return false, errors.New("Rezept nicht im Warenkorb gefunden")
}

func (s *ShoppingCartInMemoryStorage) Add(id int, count int) (bool, error) {
	recipe, err := s.recipeRepository.GetRecipe(id)
	if err != nil {
		return false, err
	}

	s.nextId++
	s.shoppingCart[s.nextId] = &model.ShoppingCartItem{
		Id:     s.nextId,
		Count:  count,
		Recipe: recipe,
	}

	return true, nil
}

func (s *ShoppingCartInMemoryStorage) Get() []*model.ShoppingCartItem {
	items := make([]*model.ShoppingCartItem, 0, len(s.shoppingCart))
	for _, item := range s.shoppingCart {
		items = append(items, item)
	}
	return items
}
