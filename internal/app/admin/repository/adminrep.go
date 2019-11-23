package adminrepository

import "database/sql"
//import "github.com/golang-chat/internal/"

type AdminRepository struct {
	db *sql.DB
}


func NewAdminRepository(db *sql.DB) admin.Repository {
	return &AdminRepository{db}
}

func (r *AdminRepository) Create(company *model.Company) error {
	return r.db.QueryRow(
		"INSERT INTO companies (companyName, site, tagLine, description, country, city, address, phone) " +
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		company.CompanyName,
		company.Site,
		company.TagLine,
		company.Description,
		company.Country,
		company.City,
		company.Address,
		company.Phone,
	).Scan(&company.ID)
}
