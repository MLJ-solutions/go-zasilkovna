package models

import "time"

const dateLayout = "2006-01-02"
const dateTimeLayout = "2006-01-02 15:00:00"

type ZasilkovnaDate time.Time

type ZasilkovnaDateTime time.Time