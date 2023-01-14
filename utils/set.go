package utils

// 声明元素类型
type Element comparable

// 空结构体做值 unsafe.Sizeof(struct{}) = 0
type Empty struct{}

/*
不要用空接口做值 空接口占用空间为 16

> go/src/reflect/value.go:195

type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}
*/

var empty Empty

// Set 结构体
type Set[E Element] struct {
	m map[E]Empty  // map 存储
}

func NewSet[E Element]() *Set[E] {
	return &Set[E]{m: map[E]Empty{}}
}

// 添加元素
func (s *Set[E]) Add(val E) {
	s.m[val] = empty // 使用一个empty单例作为所有键的值
}

// 删除元素
func (s *Set[E]) Remove(val E) {
	delete(s.m, val)
}

// 获取长度
func (s *Set[E]) Size() int {
	return len(s.m)
}

// 清空set
func (s *Set[E]) Clear() {
	s.m = make(map[E]Empty)
}

// 查看某个元素是否存在
func (s *Set[E]) Exist(val E) (ok bool) {
	_, ok = s.m[val]
	return
}

// 用于遍历
func (s *Set[E]) Elems() map[E]Empty {
	return s.m
}
