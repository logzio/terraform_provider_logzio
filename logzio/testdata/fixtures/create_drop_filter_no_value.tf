resource "logzio_drop_filter" "%s" {
  log_type = "some_type"

  field_conditions {
    field_name = "some_field"
  }
}