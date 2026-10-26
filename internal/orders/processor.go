package orders

  import (
      "context"
  )

  func ProcessOrder(ctx context.Context, order *Order) error {
      return processSingleOrder(ctx, order)
  }

  func processSingleOrder(ctx context.Context, order *Order) error {
      return nil
  }
