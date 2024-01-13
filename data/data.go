package data

var DB *Store

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	data map[int]*User
}

func GetDB() *Store {
	if DB == nil {
		DB = &Store{
			data: make(map[int]*User),
		}
	}

	return DB
}

func (s *Store) Get() []*User {
	users := []*User{}

	for _, user := range s.data {
		users = append(users, user)
	}

	return users
}

func (s *Store) GetById(id int) *User {
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

	s.data[user.Id] = user
	return user
}

func (s *Store) Delete(id int) {
	if _, isPresent := s.data[id]; isPresent {
		delete(s.data, id)
	}
}

func (s *Store) Update(user *User) *User {
	s.data[user.Id] = user
	return user
}
