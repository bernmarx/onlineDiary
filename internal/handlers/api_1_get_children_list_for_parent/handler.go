package api_1_get_children_list_for_parent

import (
	"github.com/bernmarx/onlineDiary/internal/handlers/adapter"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

const (
	parentsStudentsCollectionName = "parents_students"
	studentIdKey                  = "student_id"
)

type recordGetter interface {
	FindRecordsByExpr(string, ...dbx.Expression) ([]*models.Record, error)
}

type Handler struct {
	recordGetter recordGetter
}

type Input struct {
	ParentID string `json:"parentId"`
}

type Output struct {
	ChildrenIDs []string `json:"childrenIds"`
}

func NewHandler(recordGetter recordGetter) echo.HandlerFunc {
	handler := Handler{
		recordGetter: recordGetter,
	}

	return adapter.Query(handler.Handle)
}

func (h *Handler) Handle(c echo.Context, input Input) (Output, error) {
	records, err := h.recordGetter.FindRecordsByExpr(parentsStudentsCollectionName, dbx.HashExp{"parent_id": input.ParentID})
	if err != nil {
		return Output{}, err
	}

	res := Output{}

	for _, record := range records {
		res.ChildrenIDs = append(res.ChildrenIDs, record.GetString(studentIdKey))
	}

	return res, nil
}
