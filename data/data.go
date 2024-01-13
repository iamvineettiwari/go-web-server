package data

import "sync"

var DB *Store

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	locker sync.RWMutex
	data   map[int]*User
}

func GetDB() *Store {
	if DB == nil {
		DB = &Store{
			data:   make(map[int]*User),
			locker: sync.RWMutex{},
		}
	}

	return DB
}

func (s *Store) Get() []*User {
	users := []*User{}

	s.locker.RLock()
	defer s.locker.RUnlock()

	for _, user := range s.data {
		users = append(users, user)
	}

	return users
}

func (s *Store) GetById(id int) *User {
	s.locker.RLock()
	defer s.locker.RUnlock()

	user, isPresent := s.data[id]

	if !isPresent {
		return nil
	}

	return user
}

func (s *Store) Insert(user *User) *User {
	if user.Id == 0 {
		(*user).Id = len(s.data) + 1
	}

	s.locker.Lock()
	defer s.locker.Unlock()

	s.data[user.Id] = user
	return user
}

func (s *Store) Delete(id int) {
	s.locker.Lock()
	s.locker.Unlock()

	if _, isPresent := s.data[id]; isPresent {
		delete(s.data, id)
	}
}

func (s *Store) Update(user *User) *User {
	s.locker.Lock()
	s.locker.Unlock()

	s.data[user.Id] = user
	return user
}
