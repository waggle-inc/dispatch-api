package orders

import (
    "context"
)

// ProcessOrder handles order processing with batch picking support
func ProcessOrder(ctx context.Context, order *Order) error {
    // Use batch picking workflow - process multiple orders together
    return processBatchOrder(ctx, order)
}

func processBatchOrder(ctx context.Context, order *Order) error {
    // Batch picking logic: group with other orders, optimize route
    // ...
    return nil
}
