package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nomankhokhar/GoTonyUrl/api/database"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest

	if err := c.ShouldBind(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return

	}
	shortid := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortid).Result()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found for the given ShortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		// if the data is not a JSON object, assume it as plain string
		data = make(map[string]interface{})
		data["data"] = val
	}

	// check if 'tags' filed already exists and it's slice of strings

	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	// check for dublicate tags
	for _, existingTags := range tags {
		if existingTags == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Tag Already Exists",
			})
			return
		}
	}

	// Add the new tag to the Tags slice

	tags = append(tags, tag)

	data["tags"] = tags

	// marshal the updated data back to JSON
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Marshal update data",
		})
		return
	}

	err = r.Set(database.Ctx, shortid, updatedData, 0).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Update the Database",
		})
		return
	}

	// Response with the updated data
	c.JSON(http.StatusOK, data)
}
