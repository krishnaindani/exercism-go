package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	value, present := units[unit]
	if !present {
		return false
	}

	//bill[item] += value
	//can replace below code with this one line
	//but not too readable

	itemsValue, itemPresent := bill[item]
	if !itemPresent {
		bill[item] = value
	} else {
		bill[item] = itemsValue + value
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	itemValue, itemPresent := bill[item]
	if !itemPresent {
		return false
	}

	unitValue, unitPresent := units[unit]
	if !unitPresent {
		return false
	}

	newQuantity := itemValue - unitValue
	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQuantity

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	itemValue, itemPresent := bill[item]
	if !itemPresent {
		return 0, false
	}

	return itemValue, true
}
