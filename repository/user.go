package repository

import "github.com/lunghyun/CRUD_SERVER/types"

type UserRepository struct {
	userMap []*types.User
	// 나중에 db로 변경
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser) // 이번엔 이런 검증까지는 안할 생각
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}

func (u *UserRepository) Update(updatedUser *types.User) error {
	// name이 같은 user를 찾고, 수정
	isExist := false

	for idx, user := range u.userMap {
		if user.Name == updatedUser.Name {
			u.userMap[idx] = updatedUser
			isExist = true
		}
	}

	if !isExist {
		return nil // TOdo 에러처리
	}

	return nil
}

func (u *UserRepository) Delete(userName string) error {
	// name과 age에 해당하는 user 삭제
	isExist := false

	for idx, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:idx], u.userMap[idx+1:]...)
			isExist = true
			break
		}
	}

	if !isExist {
		return nil // Todo 에러코드작성해야함
	}

	return nil
}
