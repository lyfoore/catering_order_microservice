package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetOrdersTable(ctx *context.Context) table.Table {

	orders := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := orders.GetInfo().HideFilterArea()

	info.SetTable("public.orders").SetTitle("Orders").SetDescription("Orders")

	formList := orders.GetForm()

	formList.SetTable("public.orders").SetTitle("Orders").SetDescription("Orders")

	return orders
}
