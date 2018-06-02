package services

/*func getBooks(c *gin.Context) {
	// Check if the article ID is valid
	if bookID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if book, err := books.GetByCategory(bookID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   "GetByCategory",
				"payload": book}, "book.html")

		} else {
			// If the book is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid category ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}    */