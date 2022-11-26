package fog

import (
	"gopkg.in/guregu/null.v3"
)

type Host struct {
	ID       null.Int    `db:"hostID"`
	Hostname null.String `db:"hostName"`
}

type MacAddress struct {
	ID         null.Int    `db:"hmID"`
	HostID     null.Int    `db:"hmHostID"`
	MacAddress null.String `db:"hmMAC"`
	Primary    null.Bool   `db:"hmPrimary"`
	Pending    null.Bool   `db:"hmPending"`
}
