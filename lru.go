package Algo

import "Algo/list"

type LRU struct {
	maxSize int64
	currSize int64
	ll *list.List
	cache map[string]list.Val
}

type ele struct {
	key string
	value interface{}
}

func NewLRU(max int64)*LRU{
	return &LRU{
		maxSize: max,
		currSize: 0,
		ll:    list.NewListWithHead(),
		cache: make(map[string]list.Val,0),
	}
}

func (l *LRU)Add(key string,value interface{}){
	if _,ok:= l.cache[key];ok{
		l.ll.MoveToFront(value)
	}else{
		l.ll.MoveToFront(value)
		l.currSize++
		l.cache[key]=list.Val{Value:value}
	}
	for l.maxSize!=0 && l.currSize>=l.maxSize{
		l.RemoveOld()
	}
}


func (l *LRU) Get(key string)interface{} {
	if v,ok:=l.cache[key];ok{
		l.ll.MoveToFront(v)
		v := v.Value.(ele)
		return v.value
	}
	return nil
}

func (l *LRU) RemoveOld() {
	back := l.ll.Back()
	if back!=nil{
		l.ll.DeleteElement(back)
		back := back.(ele)
		delete(l.cache,back.key)
		l.currSize--
	}

}

