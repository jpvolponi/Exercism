package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok { //Return `false` if the given `unit` is not in the `units` map.
		return ok
	} else if _, ok := bill[item]; !ok { //Otherwise add the item to the customer `bill`, indexed by the item name, then return `true`.
		bill[item] = units[unit]
		return true
	} else { // If the item is already present in the bill, increase its quantity by the amount that belongs to the provided `unit`.D
		bill[item] += units[unit]
		return ok
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return ok
	} else if sera, ok := bill[item]; !ok {
		return ok
	} else if sera < 0 {
		fmt.Println("Debug: ", sera)
		return ok
	} else if sera == 0 {
		delete(bill, item)
		fmt.Println("Debug: ", sera)
		return true
	} else {
		bill[item] -= units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
