package main

/*
 * @lc app=leetcode.cn id=1670 lang=golang
 *
 * [1670] 设计前中后队列
 */

// @lc code=start
type QueueVal struct {
	val  int
	pre  *QueueVal
	next *QueueVal
}

type FrontMiddleBackQueue struct {
	len    int
	head   *QueueVal
	end    *QueueVal
	mid    *QueueVal
	midIdx int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	queueVal := &QueueVal{
		val: val,
	}
	if this.len == 0 {
		this.head = queueVal
		this.end = queueVal
		this.mid = queueVal
		this.midIdx = 1
		this.len = 1
	} else {
		this.head.pre = queueVal
		queueVal.next = this.head
		this.head = queueVal
		this.len += 1
		this.midIdx += 1
		this.nodifyQueue(true)
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	queueVal := &QueueVal{
		val: val,
	}

	if this.len == 0 {
		this.head = queueVal
		this.end = queueVal
		this.mid = queueVal
		this.midIdx = 1
		this.len = 1
	} else {
		this.mid.next.pre = queueVal
		queueVal.next = this.mid.next
		this.mid.next = queueVal
		this.len += 1
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	queueVal := &QueueVal{
		val: val,
	}

	if this.len == 0 {
		this.head = queueVal
		this.end = queueVal
		this.mid = queueVal
		this.midIdx = 1
		this.len = 1
	} else {
		this.mid.pre = queueVal
		queueVal.next = this.mid
		this.mid = queueVal
		this.len += 1
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {

}

func (this *FrontMiddleBackQueue) PopMiddle() int {

}

func (this *FrontMiddleBackQueue) PopBack() int {

}

func (this *FrontMiddleBackQueue) nodifyQueue(add bool) bool {
	if add {

		this.len = this.len + 1
		if this.len < 3 {
			return false
		}

		if this.len%2 == 0 {
			return true
		}
		return false
	} else {
		this.len = this.len - 1
		if this.len < 3 {
			return false
		}

		if this.len%2 == 1 {
			return true
		}
		return false
	}

}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
// @lc code=end
