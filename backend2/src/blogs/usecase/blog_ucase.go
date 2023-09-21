package usecase

import (
	"github.com/y-maruyama1002/Techport/domain"
)

type blogUsecase struct {
	blogRepo domain.BlogRepository
}

func NewBlogUsecase(blg domain.BlogRepository) domain.BlogUsecase {
	return &blogUsecase{
		blogRepo: blg,
	}
}

func (u *blogUsecase) GetById(id int64) (res domain.Blog, err error) {
	res, err = u.blogRepo.GetById(id)
	if err != nil {
		return
	}
	return
}

func (u *blogUsecase) CreateBlog(blog *domain.CreateBlog) error {
	// titleで存在性のチェック

	err := u.blogRepo.CreateBlog(blog)
	return err
}