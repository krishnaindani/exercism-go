package elon

import "fmt"

//Drive ...
func (c *Car) Drive() {
	battery := c.battery - c.batteryDrain
	if battery < 0 {
		return
	}
	c.battery = battery
	c.distance = c.speed
}

//DisplayDistance ...
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

//DisplayBattery ...
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

//CanFinish ...
func (c Car) CanFinish(trackDistance int) bool {
	batteryNeeded := trackDistance * c.batteryDrain / c.speed
	return c.battery >= batteryNeeded
}
