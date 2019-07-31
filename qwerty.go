	//reading an integer
    // var age int
    // fmt.Println("What is your age?")
    // _, err: fmt.Scan(&age)

    // func (l *List) Init() *List

    order := list.New()
    for i := 0; i < citizen-1 ; i++ 
    {
    	var citizen_order int64
   	    _, err: fmt.Scan(&citizen_order)
		co := l.PushBack(citizen_order)    	
    }


    visit := list.New()
    for i := 0; i < king-1 ; i++ 
    {
    	var king_visit int64
   	    _, err: fmt.Scan(&king_visit)
		kv := l.PushBack(king_visit)    	
    }

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

// remove removes e from its list, decrements l.len, and returns e.
func (l *List) remove(e *Element) *Element 
{
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}




for i := 0; i < king-1 ; i++ 
    {
    	arrival := visit.Next()    	
    }
    max(order,arrival)
}