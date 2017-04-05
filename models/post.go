package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

type Post struct {
    Id      string `gorm:"primary_key;size:16"`
    Title   string `gorm:"size:100;default:'<blank>'"`
    Content string `gorm:"type:text"`
    Created time.Time
    Updated time.Time
}

func (post *Post) BeforeCreate(scope *gorm.Scope) {
    scope.SetColumn("id", GenShortId(8))
}

func (post *Post) BeforeSave(scope *gorm.Scope) {
    if post.Created.IsZero() {
        scope.SetColumn("created", time.Now())
    } else {
        scope.SetColumn("created", post.Created)
    }
    scope.SetColumn("updated", time.Now())
}

func (post *Post) Save() {
    db.Save(post)
}

func GetPostById(id string) (*Post, error) {
    post := new(Post)
    err := db.Where("id = ?", id).First(post).Error
    return post, err
}

func GetPostsCount() uint {
    var count uint
    posts := new(Post)
    db.Model(posts).Count(&count)
    return count
}

func GetRecentPosts() *[]*Post {
    posts := new([]*Post)
    db.Order("updated desc").Limit(5).Find(posts)
    return posts
}

func GetAllPosts() *[]*Post {
    posts := new([]*Post)
    db.Order("created desc").Find(posts)
    return posts
}

func DeletePosts(ids []string) {
    posts := new(Post)
    db.Where("id in (?)", ids).Delete(posts)
}
