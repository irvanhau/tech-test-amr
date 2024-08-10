package helpers

func FormatPaginate(totalPage, perPage, total, currentPage, lastPage, nextPage any) map[string]any {
	var paginate = map[string]any{}

	paginate["total_page"] = totalPage
	paginate["per_page"] = perPage
	paginate["total_data"] = total
	paginate["current_page"] = currentPage
	paginate["last_page"] = lastPage
	paginate["next_page"] = nextPage

	return paginate

}
