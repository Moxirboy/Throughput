package usecase

import (
	"database/sql"
	"fmt"
	"project/internal/controller/v1/dto"
)

type RequirementImpl struct {
	Requirement     dto.Requirements
	RequirementGood dto.RequirementGoods
}

func (r *RequirementImpl) GetId(DB *sql.DB) string {
	var goodsId string
	err := DB.QueryRow("select id from kirim.goods where name=?", r.RequirementGood.GoodName).Scan(&goodsId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("Goods ID:", goodsId)
	}
	return goodsId
}
func (r *RequirementImpl) InserterRequirementGoods(DB *sql.DB) error {
	var clientId string
	err := DB.QueryRow("select id from client where name=?", r.Requirement.NameClient).Scan(&clientId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("client ID:", clientId)
	}
	_, err = DB.Query("insert into requirement (date,client_id) values(?,?)", r.Requirement.Date, clientId)
	if err != nil {

		return err
	}
	var RequirementId string
	err = DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("Requirement ID:", RequirementId)
	}
	_, err = DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, r.GetId(DB), r.RequirementGood.Amount, r.RequirementGood.CostCell)
	if err != nil {
		return err
	}
	return nil
}
