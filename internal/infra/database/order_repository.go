package database

import (
	"database/sql"

	"github.com/devfullcycle/intensivo-go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("Insert into orders (id, price, tax, final_price) Values(?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (float64, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return float64(total), nil
}
