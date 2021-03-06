package registry

import (
	"apiserver/pkg/storage/mysqld"
	"apiserver/pkg/util/jsonx"
)

var (
	db = mysqld.GetDB()
)

type Image struct {
	Name   string      `json:"name"`
	TagLen int         `json:"tagLen"`
	Tags   []string    `json:"tags"`
	Fest   []*Manifest `json:"manifest"`
}

type Manifest struct {
	UserName      string `json:"namespace"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Architecture  string `json:"architecture"`
	Os            string `json:"os"`
	Author        string `json:"author"`
	Id            string `json:"id"`
	ParentId      string `json:"parent"`
	Created       string `json:"created"`
	DockerVersion string `json:"docker_version"`
	Pull          string `json:"pull"`
}

func init() {
	db.SingularTable(true)
	db.CreateTable(&Manifest{})
}

func (manifest *Manifest) String() string {
	manifestStr := jsonx.ToJson(manifest)
	return manifestStr
}

func (manifest *Manifest) Insert() {
	db.Create(manifest)
}

func (manifest *Manifest) Delete() {
	db.Model(manifest).Delete(manifest)
}

func (manifest *Manifest) Update() {
	db.Model(manifest).Update(manifest)
}

func (manifest *Manifest) QueryOne() *Manifest {
	db.Model(manifest).First(manifest)
	return manifest
}

func (manifest *Manifest) QuerySet(where map[string]interface{}) (fests []*Manifest, total int64) {
	pageCnt := where["pageCnt"].(int)
	pageNum := where["pageNum"].(int)
	namespace := where["namespace"].(string)
	if where["name"].(string) != "" {
		name := where["name"].(string)
		db.Model(manifest).Where("name like ? and user_name=?", `%`+name+`%`, namespace).Offset(pageCnt * pageNum).Limit(pageCnt).Find(&fests)
		db.Model(manifest).Select("count(distinct name)").Where("name like ? and user_name=?", `%`+name+`%`, namespace).Count(&total)
	} else {
		db.Model(manifest).Where("user_name=?", namespace).Offset(pageCnt * pageNum).Limit(pageCnt).Order("name desc").Find(&fests)
		db.Model(manifest).Select("count(distinct name)").Count(&total)
	}
	return
}

func (manifest *Manifest) Exsit() bool {
	return !db.Model(manifest).First(manifest).RecordNotFound()
}
