package models

import "time"

// Struct for `cohorts``
type Cohort struct {
  cohortid int32
  name string
  date time.Time
  at_sku int32
  at_id string
  status string
  zoom string
  game string
  level int8
  imported_students int8
  start_date time.Time
  end_date time.Time
  mapaccess int8
  record int8
  event_id string
}
