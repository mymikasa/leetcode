package leetcode

type StockSpanner struct {
	stock []sPair
}

func Constructor() StockSpanner {
	return StockSpanner{[]sPair{}}
}

type sPair struct {
	price int
	cnt   int
}

func (s *StockSpanner) Next(price int) int {
	cnt := 1

	for len(s.stock) > 0 && s.stock[len(s.stock)-1].price <= price {
		cnt += s.stock[len(s.stock)-1].cnt
		s.stock = s.stock[:len(s.stock)-1]
	}
	s.stock = append(s.stock, sPair{price: price, cnt: cnt})
	return cnt
}
