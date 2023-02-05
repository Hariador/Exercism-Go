package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := map[string]int{}
	unitMap["quarter_of_a_dozen"] = 3
	unitMap["half_of_a_dozen"] = 6
	unitMap["dozen"] = 12
	unitMap["small_gross"] = 120
	unitMap["gross"] = 144
	unitMap["great_gross"] = 1728
	return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if exists {
		_, exists = bill[item]
		if exists {
			bill[item] += units[unit]
		} else {
			bill[item] = units[unit]
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}
	_, exists = bill[item]
	if !exists {
		return false
	}
	if bill[item]-unitValue < 0 {
		return false
	}
	if bill[item]-unitValue == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = bill[item] - unitValue
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemCount, exists := bill[item]
	if !exists {
		return 0, false
	}
	return itemCount, true
}
