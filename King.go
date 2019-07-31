package main

import "fmt"

func max(order *list,size *int64) int64
{	max := order[0]
	for i := 0; i < size-1; i++ 
	{	sample := order[i]
		if sample > max
		{	max = sample
		} 
	}
	fmt.Println(max)
	order.remove(max)
}

func main()
{	var citizen int64
    _, err: fmt.Scan(&citizen)
    var king int64
    _, err: fmt.Scan(&king)
    order := list.New()
   	var citizen_order int64
    for i := 0; i < citizen-1 ; i++ 
    {    _, err: fmt.Scan(&citizen_order)
		order.PushBack(citizen_order)    	
    }
    visit := list.New()
  	var king_visit int64
    for i := 0; i < king-1 ; i++ 
    {	_, err: fmt.Scan(&king_visit)
		visit.PushBack(king_visit)    	
    }
    for i := 0; i < king-1 ; i++ 
    {	arrival := visit[i]   	
	    wqrqerqe :=max(order,arrival)
    }
}