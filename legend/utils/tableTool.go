package utils

/**
** 计算偏移量
 */
func CountOffset(page, limit int) int {
	if page > 0 {
		return (page - 1) * limit
	} else {
		return limit
	}
}
