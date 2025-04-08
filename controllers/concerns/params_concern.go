package concerns

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AfterIdParam(c *gin.Context) (string, error) {
	afterId, afterIdExists := c.GetQuery("after_id")

	if !afterIdExists {
		return "", nil
	}

	if err := uuid.Validate(afterId); err != nil {
		return "", errors.New("invalid after uuid")
	}

	return afterId, nil
}

func LimitParam(c *gin.Context) (int, error) {
	limitStr, limitStrExists := c.GetQuery("limit")

	if !limitStrExists {
		return DefaultLimit, nil
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return 0, err
	}

	if limit <= 0 || limit > 10 {
		return DefaultLimit, nil
	}

	return limit, nil
}

func OrderParam(c *gin.Context) (string, string, string) {
	orderBy, orderDirectionExists := c.GetQuery("order_by")
	direction, directionDirectionExists := c.GetQuery("direction")

	if !orderDirectionExists {
		orderBy = DefaultAttributeOrder
	}

	if !directionDirectionExists {
		direction = DefaultDirectionOrder
	}

	OrderExpression := fmt.Sprintf("%v %v", orderBy, direction)
	return OrderExpression, orderBy, direction
}
