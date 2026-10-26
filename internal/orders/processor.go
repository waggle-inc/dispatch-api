package orders

  import (
      "context"
      "github.com/rollout/rox-go/v5/core"
  )

  var batchPickingMode = core.NewRoxFlag(false)

  func init() {
      core.Register("", batchPickingMode)
  }

  func ProcessOrder(ctx context.Context, order *Order) error {
      if batchPickingMode.IsEnabled(ctx) {
          return processBatchOrder(ctx, order)
      }
      return processSingleOrder(ctx, order)
  }

  func processBatchOrder(ctx context.Context, order *Order) error {
      return nil
  }

  func processSingleOrder(ctx context.Context, order *Order) error {
      return nil
  }
