package aldoutil

import (
	"database/sql"
	"time"

	"github.com/douglasmg7/currency"
)

var db *sql.DB

// func SetDb(dataBase *sql.DB) error {
// db = dataBase
// err := db.Ping()
// if err != nil {
// return err
// }
// return nil
// }

// Store product to create a new product on store.
type StoreProduct struct {
	// MongodbId                 string    `json:"_id"` // Identify product into store site using mongo _id.
	DealerName                        string    `json:"dealerName"`
	DealerProductId                   string    `json:"dealerProductId"`
	DealerProductTitle                string    `json:"dealerProductTitle"`
	DealerProductDesc                 string    `json:"dealerProductDesc"`
	DealerProductMaker                string    `json:"dealerProductMaker"`
	DealerProductCategory             string    `json:"dealerProductCategory"`
	DealerProductWarrantyDays         int       `json:"dealerProductWarrantyDays"`
	DealerProductDeep                 int       `json:"dealerProductDeep"`   // Deep (comprimento) in cm.
	DealerProductHeight               int       `json:"dealerProductHeight"` // Height in cm.
	DealerProductWidth                int       `json:"dealerProductWidth"`  // Width in cm.
	DealerProductWeight               int       `json:"dealerProductWeight"` // Weight in grams.
	DealerProductActive               bool      `json:"dealerProductActive"`
	DealerProductFinalPriceSuggestion int       `json:"dealerProductFinalPriceSuggestion"`
	DealerProductPrice                int       `json:"dealerProductPrice"`
	DealerProductLastUpdate           time.Time `json:"dealerProductLastUpdate"`
	DealerProductImagesLink           string    `json:"dealerProductImagesLink"` // Images link separated by "__,__".
}

// Aldo product.
type Product struct {
	MongodbId            string            `db:"mongodb_id"` // Store id from mongodb.
	Code                 string            `db:"code"`       // From dealer.
	Brand                string            `db:"brand"`
	Category             string            `db:"category"`
	Description          string            `db:"description"`
	DealerPrice          currency.Currency `db:"dealer_price"`
	SuggestionPrice      currency.Currency `db:"suggestion_price"`
	TechnicalDescription string            `db:"technical_description"`
	Availability         bool              `db:"availability"`
	Length               int               `db:"length"` // cm.
	Width                int               `db:"width"`  // cm.
	Height               int               `db:"height"` // cm.
	Weight               int               `db:"weight"` // grams.
	PictureLink          string            `db:"picture_link"`
	WarrantyPeriod       int               `db:"warranty_period"` // Months.
	RMAProcedure         string            `db:"rma_procedure"`
	CreatedAt            time.Time         `db:"created_at"`
	ChangedAt            time.Time         `db:"changed_at"`
	Changed              bool              `db:"changed"`
	New                  bool              `db:"new"`
	Removed              bool              `db:"removed"`
	StoreProductId       bool              `db:"store_product_id"`
}

// Define product status.
func (p *Product) Status(validDate time.Time) string {
	if !p.Availability {
		return "unavailable"
	}
	if p.Removed {
		return "removed"
	}
	if p.ChangedAt.Before(validDate) {
		return ""
	} else {
		if p.ChangedAt.Equal(p.CreatedAt) {
			if p.MongodbId == "" {
				return "new"
			}
			// Product alredy created at zunka site.
			return ""
		} else {
			return "changed"
		}
	}
}

// // FindByCode get product from db by code.
// func (p *Product) FindByCode(db *sql.DB, code string) error {
// return p.findByCode(db, code, false)
// }

// // FindHistoryByCode get product history from db by code.
// func (p *Product) FindHistoryByCode(db *sql.DB, code string) error {
// return p.findByCode(db, code, true)
// }

// // findByCode get product or product history from db by code.
// func (p *Product) findByCode(db *sql.DB, code string, history bool) error {
// var fieldsName []string
// var fieldsNameDb []string
// var fieldsInterface []interface{}
// // Table name.
// var tableName = "product"
// if history {
// tableName = "product_history"
// }

// val := reflect.ValueOf(p).Elem()
// for i := 0; i < val.NumField(); i++ {
// fieldType := val.Type().Field(i)
// fieldsName = append(fieldsName, fieldType.Name)
// fieldsNameDb = append(fieldsNameDb, fieldType.Tag.Get("db"))
// fieldsInterface = append(fieldsInterface, val.Field(i).Addr().Interface())
// }
// var buffer bytes.Buffer
// buffer.WriteString("SELECT ")
// buffer.WriteString(strings.Join(fieldsNameDb, ", "))
// buffer.WriteString(" FROM ")
// buffer.WriteString(tableName)
// buffer.WriteString(" WHERE code=?")

// err := db.QueryRow(buffer.String(), code).Scan(fieldsInterface...)
// return err
// }

// func FindAllProducts(db *sql.DB) ([]Product, error) {
// return findAllProducts(db, false)
// }

// func FindAllProductsHistory(db *sql.DB) ([]Product, error) {
// return findAllProducts(db, true)
// }

// func findAllProducts(db *sql.DB, history bool) (products []Product, err error) {
// p := &Product{}
// var fieldsName []string
// var fieldsNameDb []string
// var fieldsInterface []interface{}
// // Table name.
// var tableName = "product"
// if history {
// tableName = "product_history"
// }

// val := reflect.ValueOf(p).Elem()
// for i := 0; i < val.NumField(); i++ {
// fieldType := val.Type().Field(i)
// fieldsName = append(fieldsName, fieldType.Name)
// fieldsNameDb = append(fieldsNameDb, fieldType.Tag.Get("db"))
// fieldsInterface = append(fieldsInterface, val.Field(i).Addr().Interface())
// }
// var buffer bytes.Buffer
// buffer.WriteString("SELECT ")
// buffer.WriteString(strings.Join(fieldsNameDb, ", "))
// buffer.WriteString(" FROM ")
// buffer.WriteString(tableName)
// // buffer.WriteString(" LIMIT 10")

// rows, err := db.Query(buffer.String())
// if err != nil {
// return
// }
// defer rows.Close()
// for rows.Next() {
// product := Product{}
// fieldsInterfaceRes := []interface{}{}
// val := reflect.ValueOf(product).Elem()
// for i := 0; i < val.NumField(); i++ {
// fieldsInterfaceRes = append(fieldsInterfaceRes, val.Field(i).Addr().Interface())
// }
// err = rows.Scan(fieldsInterface...)
// if err != nil {
// return
// }
// products = append(products, product)
// }
// return
// }

// // Save product to db.
// func (p *Product) Save(db *sql.DB) error {
// return p.save(db, false)
// }

// // Save product history to db.
// func (p *Product) SaveHistory(db *sql.DB) error {
// return p.save(db, true)
// }

// // Save  product or product history to db.
// func (p *Product) save(db *sql.DB, history bool) error {
// var fieldsName []string
// var fieldsNameDb []string
// var fieldsInterface []interface{}
// // Table name.
// var tableName = "product"
// if history {
// tableName = "product_history"
// }

// val := reflect.ValueOf(p).Elem()
// for i := 0; i < val.NumField(); i++ {
// fieldType := val.Type().Field(i)
// fieldsName = append(fieldsName, fieldType.Name)
// fieldsNameDb = append(fieldsNameDb, fieldType.Tag.Get("db"))
// fieldsInterface = append(fieldsInterface, val.Field(i).Addr().Interface())
// }
// var buffer bytes.Buffer
// buffer.WriteString("INSERT INTO ")
// buffer.WriteString(tableName)
// buffer.WriteString(` (`)
// buffer.WriteString(strings.Join(fieldsNameDb, ", "))
// buffer.WriteString(`) VALUES(?`)
// buffer.WriteString(strings.Repeat(`, ?`, len(fieldsNameDb)-1))
// buffer.WriteString(`)`)

// stmt, err := db.Prepare(buffer.String())
// if err != nil {
// return err
// }
// defer stmt.Close()
// _, err = stmt.Exec(fieldsInterface...)
// return err
// }

// // Update product from db.
// func (p *Product) Update(db *sql.DB) error {
// return p.update(db, false)
// }

// // UpdateHistory update product history db.
// func (p *Product) UpdateHistory(db *sql.DB) error {
// return p.update(db, true)
// }

// // Update product or product history db.
// func (p *Product) update(db *sql.DB, history bool) error {
// var fieldsNameSet []string
// // var fieldsNameDb []string
// var fieldsInterface []interface{}
// // Table name.
// var tableName = "product"
// if history {
// tableName = "product_history"
// }

// val := reflect.ValueOf(p).Elem()
// for i := 0; i < val.NumField(); i++ {
// fieldType := val.Type().Field(i)
// // fieldsName = append(fieldsName, fieldType.Name)
// fieldsNameSet = append(fieldsNameSet, fieldType.Tag.Get("db")+"=?")
// fieldsInterface = append(fieldsInterface, val.Field(i).Addr().Interface())
// }

// query := "UPDATE " + tableName + " SET " + strings.Join(fieldsNameSet, ", ") + " WHERE id=?"
// // fmt.Println(query)
// // fmt.Println("brand:", p.Brand)

// _, err := db.Exec(query, fieldsInterface...)
// return err
// }

// Diff check if products are different.
func (p *Product) Diff(pn *Product) bool {
	if p.Code != pn.Code {
		return true
	}
	if p.Brand != pn.Brand {
		return true
	}
	if p.Category != pn.Category {
		return true
	}
	if p.Description != pn.Description {
		return true
	}
	// if p.Unit != pn.Unit {
	// return true
	// }
	// if p.Multiple != pn.Multiple {
	// return true
	// }
	if p.DealerPrice != pn.DealerPrice {
		return true
	}
	if p.SuggestionPrice != pn.SuggestionPrice {
		return true
	}
	if p.TechnicalDescription != pn.TechnicalDescription {
		return true
	}
	if p.Availability != pn.Availability {
		return true
	}
	if p.Length != pn.Length {
		return true
	}
	if p.Width != pn.Width {
		return true
	}
	if p.Height != pn.Height {
		return true
	}
	if p.Weight != pn.Weight {
		return true
	}
	if p.PictureLink != pn.PictureLink {
		return true
	}
	if p.WarrantyPeriod != pn.WarrantyPeriod {
		return true
	}
	if p.RMAProcedure != pn.RMAProcedure {
		return true
	}
	return false
}
