package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int { "quarter_of_a_dozen": 3,
        					"half_of_a_dozen": 6,
        					"dozen": 12,
        					"small_gross": 120,
        					"gross": 144,	
        					"great_gross": 1728 }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; ok {
        if _, ok := bill[item]; ok {
            bill[item] += units[unit]
        } else {
    		bill[item] = units[unit]
        }
    	return true
    } else {
		return false
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    _, unitExists := units[unit]
    _, itemInBill := bill[item]
    
	if unitExists && itemInBill && bill[item] - units[unit] >= 0 {
    	if bill[item] - units[unit] == 0 {
            delete(bill, item)
        } else {
        	bill[item] -= units[unit]
        }
    	return true
    } else {
		return false
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; ok {
    	return bill[item], true
    } else {
        return 0, false
    }
}
