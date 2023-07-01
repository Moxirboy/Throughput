package mysql

import "project/internal/controller/v1/dto"

type RequirementImpl struct {
	Requirement     dto.Requirements
	RequirementGood dto.RequirementGoods
}
