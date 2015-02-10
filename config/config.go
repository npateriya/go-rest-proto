package config

import (
  "shipped/datastore"
 )

var Context struct{
  DS *datastore.Datastore
}
