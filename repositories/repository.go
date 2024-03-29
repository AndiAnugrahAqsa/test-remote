package repositories

import "mini-project/models"

type BlogRepository interface {
	GetAll() []models.Blog
	GetByUserID(user_id int) []models.Blog
	GetByCategoryID(category_id int) []models.Blog
	GetByID(id int) models.Blog
	Create(blogRequest models.BlogRequest) models.Blog
	Update(id int, blogRequest models.BlogRequest) models.Blog
	Delete(id int) bool
}

type CategoryRepository interface {
	GetAll() []models.Category
	GetByID(id int) models.Category
	Create(categoryRequest models.CategoryRequest) models.Category
	Update(id int, categoryRequest models.CategoryRequest) models.Category
	Delete(id int) bool
}

type CommentRepository interface {
	GetAll() []models.Comment
	GetByBlogID(blog_id int) []models.Comment
	GetByID(id int) models.Comment
	Create(commentRequest models.CommentRequest) models.Comment
	Update(id int, commentRequest models.CommentRequest) models.Comment
	Delete(id int) bool
}

type LikeRepository interface {
	GetAll() []models.Like
	GetByBlogID(id int) []models.Like
	Create(likeRequest models.LikeRequest) models.Like
	Delete(id int) bool
}

type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) models.User
	Create(userRequest models.UserRequest) models.User
	Login(userRequest models.UserRequest) models.User
	Update(id int, userRequest models.UserRequest) models.User
	Delete(id int) bool
}
