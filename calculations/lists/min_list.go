package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	// Merke das erste element als vorläufiges min
	min := nums[0]
	if len(nums) == 0 {
		return 0
		}
	// laufe durch die liste
	for i := 1, i < len(nums), i++ {
		
	// vergleiche jedes element mit min
		if nums[i] < min {
		
	// wenn das element kleiner, dann ersetze min
		min = nums[i]
		}
	}	
	
	return min
}
