package _1801_number_of_orders_in_the_backlog

import (
	"container/heap"
	"github.com/babadro/algorithms-go/utils"
)

const (
	buy = 0
	d   = 1_000_000_007
)

// passed. tptl. best solution. medium
func getNumberOfBacklogOrders2(orders [][]int) int {
	buyOr, sellOr := &maxHeap{}, &minHeap{}

	for _, batch := range orders {
		if batch[2] == 0 {
			heap.Push(buyOr, batch)
		} else {
			heap.Push(sellOr, batch)
		}

		for buyOr.Len() > 0 && sellOr.Len() > 0 && (*sellOr)[0][0] <= (*buyOr)[0][0] {
			topSell, topBuy := (*sellOr)[0], (*buyOr)[0]

			count := utils.Min(topSell[1], topBuy[1])

			topSell[1] -= count
			topBuy[1] -= count

			if topSell[1] == 0 {
				heap.Pop(sellOr)
			}

			if topBuy[1] == 0 {
				heap.Pop(buyOr)
			}
		}
	}

	sum := 0
	for _, batch := range *buyOr {
		sum = (sum + batch[1]) % d
	}

	for _, batch := range *sellOr {
		sum = (sum + batch[1]) % d
	}

	return sum
}

// passed. too verbose
func getNumberOfBacklogOrders(orders [][]int) int {
	buyOr, sellOr := &maxHeap{}, &minHeap{}

	for _, batch := range orders {
		price, amount, typ := batch[0], batch[1], batch[2]

		if typ == buy {
			for sellOr.Len() > 0 && amount > 0 {
				minItem := heap.Pop(sellOr).([]int)
				sellPrice := minItem[0]
				if sellPrice > price {
					heap.Push(sellOr, minItem)

					break
				}

				itemAmount := minItem[1]
				if itemAmount <= amount {
					amount -= itemAmount

					continue
				}

				minItem[1] -= amount
				amount = 0
				heap.Push(sellOr, minItem)
			}

			if amount > 0 {
				heap.Push(buyOr, []int{price, amount})
			}
		} else {
			for buyOr.Len() > 0 && amount > 0 {
				maxItem := heap.Pop(buyOr).([]int)
				buyPrice := maxItem[0]
				if buyPrice < price {
					heap.Push(buyOr, maxItem)

					break
				}

				itemAmount := maxItem[1]
				if itemAmount <= amount {
					amount -= itemAmount

					continue
				}

				maxItem[1] -= amount
				amount = 0
				heap.Push(buyOr, maxItem)
			}

			if amount > 0 {
				heap.Push(sellOr, []int{price, amount})
			}
		}
	}

	sum := 0
	for _, batch := range *buyOr {
		sum += batch[1]
		sum %= d
	}

	for _, batch := range *sellOr {
		sum += batch[1]
		sum %= d
	}

	return sum
}

type minHeap [][]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]

	return res
}

type maxHeap [][]int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]

	return res
}
