package fog

import (
	"errors"

	"github.com/stockholmr/logging"
)

func Hosts_List() (*[]Host, error) {
	if _db == nil {
		return nil, errors.New("database not connected")
	}

	var model = make([]Host, 0)

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

func Hosts_Find(hostname string) (*[]Host, error) {
    if _db == nil {
        return nil, errors.New("database not connected")
    }

    var model = make([]Host, 0)

    cmd := `SELECT hostID,hostName FROM hosts WHERE hostName like '%?%'`

    rows, err := _db.Queryx(cmd, hostname)
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