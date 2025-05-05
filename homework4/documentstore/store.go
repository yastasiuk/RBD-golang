package documentstore

type Store struct {
	collections map[string]Collection
}

func NewStore() *Store {
	store := Store{
		collections: make(map[string]Collection),
	}

	return &store
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створення, то повертаємо `false` та nil
	if _, exists := s.GetCollection(name); exists {
		return false, nil
	}

	newCollection := NewCollection(cfg)
	s.collections[name] = *newCollection

	return true, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	if collection, ok := s.collections[name]; ok {
		return &collection, true
	}

	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	if _, exists := s.GetCollection(name); exists {
		delete(s.collections, name)
		return true
	}

	return false
}
