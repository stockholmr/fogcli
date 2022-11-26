package fog

import (
	"github.com/stockholmr/logging"
)

func Hosts_List() (*[]Host, error) {

	var model []Host

	cmd := `SELECT hostID,hostName FROM hosts`

	rows, err := _db.Queryx(cmd)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entry Host
		err = rows.StructScan(&entry)
		if err != nil {
			logging.Error(err.Error())
			break
		}
		model = append(model, entry)
	}

	return &model, nil
}
