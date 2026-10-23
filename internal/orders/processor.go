package orders

import (
    "context"
    "github.com/rollout/rox-go/v5/core"
)

// Define the flag
var batchPickingMode = core.NewRoxFlag(false) // default to false

// Register flags during initialization
func init() {
    core.Register("", batchPickingMode)
}

// ProcessOrder handles order processing with batch picking support
func ProcessOrder(ctx context.Context, order *Order) error {
    // Check if batch picking mode is enabled
    if batchPickingMode.IsEnabled(ctx) {
        // Use batch picking workflow - process multiple orders together
        return processBatchOrder(ctx, order)
    }
    
    // Use standard single-order picking workflow
    return processSingleOrder(ctx, order)
}

func processBatchOrder(ctx context.Context, order *Order) error {
    // Batch picking logic: group with other orders, optimize route
    // ...
    return nil
}

func processSingleOrder(ctx context.Context, order *Order) error {
    // Standard picking logic: process order individually
    // ...
    return nil
}
