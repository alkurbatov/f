package f

// Broadcast sends event to list of channels in non-blocking manner.
func Broadcast[T any](subs []chan<- T, event T) {
	for _, sub := range subs {
		select {
		case sub <- event:
		default:
		}
	}
}
