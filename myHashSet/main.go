package myHashSet

type MyHashSet struct {
	set []int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		set: []int{},
	}

}


func (this *MyHashSet) Add(key int)  {
	for i := 0; i < len(this.set); i++ {
		if this.set[i] == key {
			return
		}
	}
	this.set = append(this.set, key)
	return
}


func (this *MyHashSet) Remove(key int)  {
	for i := 0; i < len(this.set); i++ {
		if this.set[i]	== key {
			if i == len(this.set)-1 {
				this.set = this.set[:i]
				return
			}
			this.set = append(this.set[:i], this.set[i+1:]...)
			return
		}
	}
	return
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for i := 0; i < len(this.set); i++ {
		if this.set[i] == key {
			return true
		}
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
