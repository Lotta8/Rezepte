package repository

import (
	"errors"
	"fmt"
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
	for key, item := range s.shoppingCart {
		fmt.Printf("Checking item with ID %d\n", item.Recipe.ID)
		if item.Recipe.ID == id {
			delete(s.shoppingCart, key)
			fmt.Printf("Deleted item with ID %d\n", id)
			return true, nil
		}
	}
	return false, errors.New("Rezept nicht im Warenkorb gefunden")
}

func (s *ShoppingCartInMemoryStorage) Add(id int, count int) (bool, error) {
	recipe, err := s.recipeRepository.GetRecipe(id)
	if err != nil {
		return false, err
	}

	s.nextId++
	newItem := &model.ShoppingCartItem{
		Id:     s.nextId,
		Count:  count,
		Recipe: recipe,
	}
	s.shoppingCart[s.nextId] = newItem

	return true, nil
}

func (s *ShoppingCartInMemoryStorage) Get() []*model.ShoppingCartItem {
	items := make([]*model.ShoppingCartItem, 0, len(s.shoppingCart))
	for _, item := range s.shoppingCart {
		items = append(items, item)
	}
	return items
}
