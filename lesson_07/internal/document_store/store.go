package documentstore

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Store struct {
	Collections map[string]*Collection `json:"collections"`
}

func NewStore() *Store {
	store := Store{
		Collections: make(map[string]*Collection),
	}

	return &store
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	slog.Debug("CreateCollection", "name", name, "cfg", cfg)
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створення, то повертаємо `false` та nil
	if _, exists := s.GetCollection(name); exists {
		slog.Error("Cannot create collections. Not unique one", "name", name)
		return false, nil
	}

	newCollection := NewCollection(cfg)
	s.Collections[name] = newCollection

	return true, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	slog.Debug("GetCollection", "name", name)
	if collection, ok := s.Collections[name]; ok {
		return collection, true
	}

	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	slog.Debug("DeleteCollection", "name", name)
	if _, exists := s.GetCollection(name); exists {
		delete(s.Collections, name)
		return true
	}

	slog.Error("DeleteCollection: Collection does not exists:", "name", name)

	return false
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	slog.Debug("NewStoreFromDump")
	// Функція повинна створити та проініціалізувати новий `Store`
	// зі всіма колекціями та даними з вхідного дампу.
	store := &Store{}

	if err := json.Unmarshal(dump, store); err != nil {
		slog.Error("Failed to create store for dump:", "err", err)
		return nil, err
	}

	return store, nil
}

// Dump Методи повинен віддати дамп нашого стору в який включені дані про колекції та документ
func (s *Store) Dump() ([]byte, error) {
	slog.Debug("Dump store")

	jsonBytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		slog.Error("Failed to Dump() store:", "err", err)
		return nil, err
	}

	return jsonBytes, nil
}

// NewStoreFromFile Значення яке повертає метод `store.Dump()` має без помилок оброблятись функцією `NewStoreFromDump`
func NewStoreFromFile(filename string) (*Store, error) {
	slog.Debug("NewStoreFromFile", "filename", filename)
	// Робить те ж саме що і функція `NewStoreFromDump`, але сам дамп має діставатись з файлу
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Failed to read file:", "filename", filename, "err", err)
		return nil, err
	}

	store, err := NewStoreFromDump(data)
	if err != nil {
		slog.Error("Failed on store from dump creation:", "err", err)
		return nil, err
	}

	return store, nil
}

func (s *Store) DumpToFile(filename string) error {
	slog.Debug("DumpToFile", "filename", filename)
	// Робить те ж саме що і метод `Dump`, але записує у файл замість того щоб повертати сам дамп
	jsonBytes, err := s.Dump()
	if err != nil {
		slog.Error("Error on Dump()", "err", err)
		return err
	}

	err = os.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		slog.Error("Error on WriteFile()", "err", err)
		return err
	}

	return nil
}
