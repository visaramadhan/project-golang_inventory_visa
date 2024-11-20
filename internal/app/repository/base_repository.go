package repository

type BaseRepository[T any] interface {
	Create(payload *T) (*T, error)
	List() ([]*T, error)
	Get(id string) (*T, error)
	Update(id string, payload *T) (*T, error)
	Delete(id string) (*T, error)
}

// type BaseRepositoryPaging[T any] interface {
// 	Paging(requestPaging dto.PaginationParam, query ...string) ([]*T, *dto.Paging, error)
// }
