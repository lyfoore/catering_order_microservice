package admin

import (
	"fmt"
	"github.com/lib/pq"

	"github.com/sunfmin/reflectutils"
	"github.com/theplant/myadmin/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/web/v3"
	. "github.com/qor5/x/v3/ui/vuetify"
	. "github.com/theplant/htmlgo"
)

func Initialize() *http.ServeMux {
	b := setupAdmin()
	mux := setupRouter(b)

	return mux
}

func setupAdmin() (b *presets.Builder) {
	db := ConnectDB()

	b = presets.New()

	b.URIPrefix("/admin").
		BrandTitle("Admin").
		DataOperator(gorm2op.DataOperator(db)).
		HomePageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
			r.Body = VContainer(
				H1("Home"),
				P().Text("Change your home page here"),
			)
			return
		})

	mb := b.Model(&models.Order{})

	mb.Listing("ID", "UserID", "Items", "Status", "StatusAI", "Message", "Response")
	//ListAll("ID", "UserID", "Items", "Status", "CreatedAt", "UpdatedAt")

	eb := mb.Editing()
	eb.Only("UserID", "Status", "Items", "Message", "Response")

	eb.Field("Items").
		ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
			var itemsStr string
			if items, err := reflectutils.Get(obj, "Items"); err == nil {
				if int64Arr, ok := items.(pq.Int64Array); ok {
					strItems := make([]string, len(int64Arr))
					for i, val := range int64Arr {
						strItems[i] = strconv.FormatInt(val, 10)
					}
					itemsStr = strings.Join(strItems, ",")
				}
			}

			return VTextField().
				Label("Items (comma-separated)").
				Attr(
					web.VField(field.FormKey, itemsStr)...,
				).
				Hint("Example: 1,2,3")
		}).
		SetterFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) error {
			value := ctx.R.FormValue(field.FormKey)
			strItems := strings.Split(value, ",")
			var intItems pq.Int64Array

			for _, s := range strItems {
				s = strings.TrimSpace(s)
				if s == "" {
					continue
				}
				i, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid item format: %v", err)
				}
				intItems = append(intItems, i)
			}

			return reflectutils.Set(obj, "Items", intItems)
		})

	return
}
